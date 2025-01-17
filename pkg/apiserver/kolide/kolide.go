package kolide

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/nais/device/pkg/apiserver/database"
	"github.com/nais/device/pkg/pb"

	kolidepb "github.com/nais/kolide-event-handler/pkg/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientInterceptor struct {
	RequireTLS bool
	Token      string
}

func (c *ClientInterceptor) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.Token,
	}, nil
}

func (c *ClientInterceptor) RequireTransportSecurity() bool {
	return c.RequireTLS
}

const (
	eventStreamBackoff = 10 * time.Second
)

func DeviceEventStreamer(ctx context.Context, grpcAddress, grpcToken string, grpcSecure bool, stream chan<- *kolidepb.DeviceEvent) error {
	interceptor := &ClientInterceptor{
		RequireTLS: grpcSecure,
		Token:      grpcToken,
	}

	dialOpts := make([]grpc.DialOption, 0)

	if grpcSecure {
		cred := credentials.NewTLS(&tls.Config{})
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(cred))
		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(interceptor))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	conn, err := grpc.DialContext(
		ctx,
		grpcAddress,
		dialOpts...,
	)

	if err != nil {
		return err
	}

	defer conn.Close()

	s := kolidepb.NewKolideEventHandlerClient(conn)

	for ctx.Err() == nil {
		events, err := s.Events(ctx, &kolidepb.EventsRequest{})

		if err != nil {
			log.Errorf("Start Kolide event stream: %v", err)
			log.Warnf("Restarting event stream in %s...", eventStreamBackoff)
			time.Sleep(eventStreamBackoff)
			continue
		}

		log.Infof("Started Kolide event stream to %v", conn.Target())

		for {
			event, err := events.Recv()
			if err != nil {
				log.Errorf("Receive Kolide event: %v", err)
				log.Warnf("Restarting event stream in %s...", eventStreamBackoff)
				time.Sleep(eventStreamBackoff)
				break
			}

			stream <- event
		}
	}

	return ctx.Err()
}

func LookupDevice(ctx context.Context, db database.APIServer, event *kolidepb.DeviceEvent) (*pb.Device, error) {
	platform := func(platform string) string {
		switch strings.ToLower(platform) {
		case "darwin":
			return "darwin"
		case "windows":
			return "windows"
		default:
			return "linux"
		}
	}

	p := platform(event.GetPlatform())

	device, err := db.ReadDeviceBySerialPlatform(ctx, event.GetSerial(), p)
	if err != nil {
		return nil, fmt.Errorf("read device with serial=%s platform=%s: %w", event.GetSerial(), p, err)
	}

	return device, nil
}
