package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/nais/device/device-agent/apiserver"
	"github.com/nais/device/device-agent/config"
	"github.com/nais/device/device-agent/filesystem"
	"github.com/nais/device/device-agent/runtimeconfig"
	log "github.com/sirupsen/logrus"
)

type ProgramState int

const (
	StateDisconnected ProgramState = iota
	StateBootstrapping
	StateConnecting
	StateConnected
	StateDisconnecting
	StateQuitting
	StateSavingConfiguration
)

type GuiState struct {
	ProgramState ProgramState
	Gateways     apiserver.Gateways
}

func (g GuiState) String() string {
	switch g.ProgramState {
	case StateDisconnected:
		return "Disconnected"
	case StateBootstrapping:
		return "Bootstrapping..."
	case StateConnecting:
		return "Connecting..."
	case StateSavingConfiguration:
		fallthrough
	case StateConnected:
		return fmt.Sprintf("Connected since %s", connectedTime.Format(time.Kitchen))
	case StateDisconnecting:
		return "Disconnecting..."
	case StateQuitting:
		return "Quitting..."
	default:
		return "Unknown state!!!"
	}
}

var (
	cfg           = config.DefaultConfig()
	state         = StateDisconnected
	newstate      = make(chan ProgramState, 1)
	connectedTime = time.Now()
)

func notify(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args)
	err := beeep.Notify("NAIS device", message, "")
	log.Infof("sending message to notification centre: %s", message)
	if err != nil {
		log.Errorf("failed sending message due to error: %s", err)
	}
}

// read in external gui events
func guiloop(mConnect, mQuit *systray.MenuItem, interrupt chan os.Signal) {
	for {
		select {
		case <-mConnect.ClickedCh:
			if state == StateDisconnected {
				newstate <- StateConnecting
			} else {
				newstate <- StateDisconnecting
			}
		case <-mQuit.ClickedCh:
			newstate <- StateQuitting
		case <-interrupt:
			log.Info("Received interrupt, shutting down gracefully.")
			newstate <- StateQuitting
		}
	}
}

func mainloop(updateGUI func(guiState GuiState)) {
	var rc *runtimeconfig.RuntimeConfig
	var err error

	stop := make(chan interface{}, 1)

	for st := range newstate {
		state = st

		//noinspection GoNilness
		updateGUI(GuiState{
			ProgramState: state,
			Gateways:     rc.GetGateways(),
		})

		switch state {
		case StateDisconnected:
		case StateBootstrapping:
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
			rc, err = runtimeconfig.New(cfg, ctx)
			cancel()
			if err != nil {
				notify(err.Error())
				newstate <- StateDisconnected
				continue
			}
			err = WriteConfigFile(rc.Config.WireGuardConfigPath, *rc)
			if err != nil {
				err = fmt.Errorf("unable to write WireGuard configuration file: %w", err)
				notify(err.Error())
				newstate <- StateDisconnected
				continue
			}
			newstate <- StateConnecting

		case StateConnecting:
			if rc == nil {
				newstate <- StateBootstrapping
				continue
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			err := ensureAuth(ctx, rc)
			cancel()

			if err == nil {
				newstate <- StateConnected
				notify("connected")
				go connectedLoop(stop, rc)
				connectedTime = time.Now()
			} else {
				newstate <- StateDisconnected
				notify(err.Error())
			}

		case StateConnected:

		case StateDisconnecting:
			stop <- new(interface{})
			err := TruncateConfigFile(rc.Config.WireGuardConfigPath)
			if err != nil {
				notify("error synchronizing WireGuard config: %s", err)
			}
			newstate <- StateDisconnected

		case StateQuitting:
			// stop <- new(interface{})
			systray.Quit()

		case StateSavingConfiguration:
			// TODO: Bør vi egentlig skrive fila på nytt hvert 10 sekund om det ikke er endringer?
			err = WriteConfigFile(rc.Config.WireGuardConfigPath, *rc)
			if err != nil {
				err = fmt.Errorf("unable to write WireGuard configuration file: %w", err)
				notify(err.Error())
				return
			}
			newstate <- StateConnected
		}
	}
}

func connectedLoop(stop chan interface{}, rc *runtimeconfig.RuntimeConfig) {
	// Sleeping 2 seconds whilst waiting for API-server connection; waiting for wireguard to sync configuration
	time.Sleep(2 * time.Second)
	interval := 10 * time.Second
	ticker := time.NewTicker(interval)
	ctx, cancel := context.WithTimeout(context.Background(), interval)
	fetchConfig(ctx, rc)
	cancel()

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), interval)
			fetchConfig(ctx, rc)
			cancel()

		case <-stop:
			return
		}
	}
}

func fetchConfig(ctx context.Context, rc *runtimeconfig.RuntimeConfig) {
	gateways, err := apiserver.GetGateways(rc.SessionInfo.Key, rc.Config.APIServer, ctx)

	if err != nil {
		log.Errorf("unable to get gateway config: %v", err)
		return
	}

	if ue, ok := err.(*apiserver.UnauthorizedError); ok {
		newstate <- StateDisconnecting
		log.Errorf("unauthorized access from apiserver: %v", ue)
		log.Errorf("assuming invalid session; disconnecting.")
		return
	}

	for _, gw := range gateways {

		err := ping(gw.IP)
		if err == nil {
			gw.Healthy = true
		} else {
			gw.Healthy = false
			log.Errorf("unable to ping host %s: %v", gw.IP, err)
		}
	}

	rc.Gateways = gateways

	newstate <- StateSavingConfiguration
}

func readIcon(color string) []byte {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	iconPath := filepath.Join(currentDir, "..", "..", "assets", fmt.Sprintf("nais-logo-%s.png", color))
	icon, err := ioutil.ReadFile(iconPath)
	if err != nil {
		log.Errorf("finding icon: %v", err)
	}
	return icon

}

func onReady() {

	systray.SetIcon(readIcon("blue"))
	cfg.SetDefaults()
	if err := filesystem.EnsurePrerequisites(&cfg); err != nil {
		notify(fmt.Sprintf("Missing prerequisites: %s", err))
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	mState := systray.AddMenuItem("", "State")
	mState.Disable()
	systray.AddSeparator()
	mConnect := systray.AddMenuItem("Connect", "Bootstrap the nais device")
	mQuit := systray.AddMenuItem("Quit", "exit the application")
	systray.AddSeparator()
	mCurrentGateways := make(map[string]*systray.MenuItem)

	updateGUI := func(st GuiState) {
		mState.SetTitle("Status: " + st.String())
		switch st.ProgramState {
		case StateDisconnected:
			mConnect.SetTitle("Connect")
			systray.SetIcon(readIcon("red"))
			mConnect.Enable()
		case StateConnected:
			mConnect.SetTitle("Disconnect")
			systray.SetIcon(readIcon("green"))
			mConnect.Enable()
			mConnect.Enable()
		case StateSavingConfiguration:
			for _, gateway := range st.Gateways {
				if _, ok := mCurrentGateways[gateway.Endpoint]; !ok {
					mCurrentGateways[gateway.Endpoint] = systray.AddMenuItem(gateway.Name, gateway.Endpoint)
					mCurrentGateways[gateway.Endpoint].Disable()
				}
				if gateway.Healthy {
					mCurrentGateways[gateway.Endpoint].Check()
				} else {
					mCurrentGateways[gateway.Endpoint].Uncheck()
				}
			}

		default:
			mConnect.Disable()
		}
	}

	go guiloop(mConnect, mQuit, interrupt)
	newstate <- StateDisconnected
	mainloop(updateGUI)
}

func onExit() {
	// This is where we clean up
}

func ensureAuth(ctx context.Context, rc *runtimeconfig.RuntimeConfig) error {
	var err error
	if !rc.SessionInfo.Expired() {
		return nil
	}
	rc.SessionInfo, err = ensureValidSessionInfo(cfg.APIServer, cfg.Platform, rc.Serial, ctx)
	if err != nil {
		return fmt.Errorf("ensuring valid session key: %v", err)
	}
	return nil
}

func ping(addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	cmd := exec.CommandContext(ctx, "ping", "-c", "1", "-t", "1", addr)
	defer cancel()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running command %v: %w", cmd, err)
	}
	return nil
}
