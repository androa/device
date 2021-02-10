// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/pb/protobuf-api.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AgentState int32

const (
	AgentState_Disconnected   AgentState = 0
	AgentState_Bootstrapping  AgentState = 1
	AgentState_Connected      AgentState = 2
	AgentState_Disconnecting  AgentState = 3
	AgentState_Unhealthy      AgentState = 4
	AgentState_Quitting       AgentState = 5
	AgentState_Authenticating AgentState = 6
	AgentState_SyncConfig     AgentState = 7
	AgentState_HealthCheck    AgentState = 8
)

// Enum value maps for AgentState.
var (
	AgentState_name = map[int32]string{
		0: "Disconnected",
		1: "Bootstrapping",
		2: "Connected",
		3: "Disconnecting",
		4: "Unhealthy",
		5: "Quitting",
		6: "Authenticating",
		7: "SyncConfig",
		8: "HealthCheck",
	}
	AgentState_value = map[string]int32{
		"Disconnected":   0,
		"Bootstrapping":  1,
		"Connected":      2,
		"Disconnecting":  3,
		"Unhealthy":      4,
		"Quitting":       5,
		"Authenticating": 6,
		"SyncConfig":     7,
		"HealthCheck":    8,
	}
)

func (x AgentState) Enum() *AgentState {
	p := new(AgentState)
	*p = x
	return p
}

func (x AgentState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentState) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_protobuf_api_proto_enumTypes[0].Descriptor()
}

func (AgentState) Type() protoreflect.EnumType {
	return &file_pkg_pb_protobuf_api_proto_enumTypes[0]
}

func (x AgentState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentState.Descriptor instead.
func (AgentState) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{0}
}

type ConfigureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigureResponse) Reset() {
	*x = ConfigureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureResponse) ProtoMessage() {}

func (x *ConfigureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureResponse.ProtoReflect.Descriptor instead.
func (*ConfigureResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{0}
}

type UpgradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpgradeResponse) Reset() {
	*x = UpgradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeResponse) ProtoMessage() {}

func (x *UpgradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeResponse.ProtoReflect.Descriptor instead.
func (*UpgradeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{1}
}

type ConfigureJITAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigureJITAResponse) Reset() {
	*x = ConfigureJITAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureJITAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureJITAResponse) ProtoMessage() {}

func (x *ConfigureJITAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureJITAResponse.ProtoReflect.Descriptor instead.
func (*ConfigureJITAResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{2}
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{3}
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{4}
}

type UpgradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpgradeRequest) Reset() {
	*x = UpgradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeRequest) ProtoMessage() {}

func (x *UpgradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeRequest.ProtoReflect.Descriptor instead.
func (*UpgradeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{5}
}

type ConfigureJITARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateway *Gateway `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *ConfigureJITARequest) Reset() {
	*x = ConfigureJITARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureJITARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureJITARequest) ProtoMessage() {}

func (x *ConfigureJITARequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureJITARequest.ProtoReflect.Descriptor instead.
func (*ConfigureJITARequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigureJITARequest) GetGateway() *Gateway {
	if x != nil {
		return x.Gateway
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{7}
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{8}
}

type AgentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeepConnectionOnComplete bool `protobuf:"varint,1,opt,name=keepConnectionOnComplete,proto3" json:"keepConnectionOnComplete,omitempty"`
}

func (x *AgentStatusRequest) Reset() {
	*x = AgentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentStatusRequest) ProtoMessage() {}

func (x *AgentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentStatusRequest.ProtoReflect.Descriptor instead.
func (*AgentStatusRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{9}
}

func (x *AgentStatusRequest) GetKeepConnectionOnComplete() bool {
	if x != nil {
		return x.KeepConnectionOnComplete
	}
	return false
}

type AgentStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionState     AgentState             `protobuf:"varint,1,opt,name=connectionState,proto3,enum=naisdevice.AgentState" json:"connectionState,omitempty"`
	ConnectedSince      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=connectedSince,proto3" json:"connectedSince,omitempty"`
	NewVersionAvailable bool                   `protobuf:"varint,3,opt,name=newVersionAvailable,proto3" json:"newVersionAvailable,omitempty"`
	Gateways            []*Gateway             `protobuf:"bytes,4,rep,name=Gateways,proto3" json:"Gateways,omitempty"`
}

func (x *AgentStatus) Reset() {
	*x = AgentStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentStatus) ProtoMessage() {}

func (x *AgentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentStatus.ProtoReflect.Descriptor instead.
func (*AgentStatus) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{10}
}

func (x *AgentStatus) GetConnectionState() AgentState {
	if x != nil {
		return x.ConnectionState
	}
	return AgentState_Disconnected
}

func (x *AgentStatus) GetConnectedSince() *timestamppb.Timestamp {
	if x != nil {
		return x.ConnectedSince
	}
	return nil
}

func (x *AgentStatus) GetNewVersionAvailable() bool {
	if x != nil {
		return x.NewVersionAvailable
	}
	return false
}

func (x *AgentStatus) GetGateways() []*Gateway {
	if x != nil {
		return x.Gateways
	}
	return nil
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey string     `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	DeviceIP   string     `protobuf:"bytes,2,opt,name=deviceIP,proto3" json:"deviceIP,omitempty"`
	Gateways   []*Gateway `protobuf:"bytes,3,rep,name=Gateways,proto3" json:"Gateways,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{11}
}

func (x *Configuration) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *Configuration) GetDeviceIP() string {
	if x != nil {
		return x.DeviceIP
	}
	return ""
}

func (x *Configuration) GetGateways() []*Gateway {
	if x != nil {
		return x.Gateways
	}
	return nil
}

type Gateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Healthy                  bool     `protobuf:"varint,2,opt,name=healthy,proto3" json:"healthy,omitempty"`
	PublicKey                string   `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Endpoint                 string   `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Ip                       string   `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	Routes                   []string `protobuf:"bytes,6,rep,name=routes,proto3" json:"routes,omitempty"`
	RequiresPrivilegedAccess bool     `protobuf:"varint,7,opt,name=requiresPrivilegedAccess,json=requires_privileged_access,proto3" json:"requiresPrivilegedAccess,omitempty"`
	AccessGroupIDs           []string `protobuf:"bytes,8,rep,name=accessGroupIDs,proto3" json:"accessGroupIDs,omitempty"`
}

func (x *Gateway) Reset() {
	*x = Gateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gateway) ProtoMessage() {}

func (x *Gateway) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gateway.ProtoReflect.Descriptor instead.
func (*Gateway) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{12}
}

func (x *Gateway) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Gateway) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *Gateway) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Gateway) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Gateway) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Gateway) GetRoutes() []string {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *Gateway) GetRequiresPrivilegedAccess() bool {
	if x != nil {
		return x.RequiresPrivilegedAccess
	}
	return false
}

func (x *Gateway) GetAccessGroupIDs() []string {
	if x != nil {
		return x.AccessGroupIDs
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protobuf_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protobuf_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protobuf_api_proto_rawDescGZIP(), []int{13}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_pb_protobuf_api_proto protoreflect.FileDescriptor

var file_pkg_pb_protobuf_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x61, 0x69,
	0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a,
	0x0f, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4a, 0x49, 0x54,
	0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45,
	0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4a, 0x49, 0x54, 0x41, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x07, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x18,
	0x6b, 0x65, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18,
	0x6b, 0x65, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x0b, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x30,
	0x0a, 0x13, 0x6e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6e, 0x65, 0x77,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x2f, 0x0a, 0x08, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x08, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x73, 0x22, 0x7c, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x50, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x50, 0x12, 0x2f,
	0x0a, 0x08, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x08, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x22,
	0xff, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x18, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x73, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2a, 0xa5, 0x01, 0x0a, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x51, 0x75, 0x69,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x10, 0x08, 0x32, 0x9f, 0x01, 0x0a,
	0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x49, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x19, 0x2e, 0x6e, 0x61, 0x69,
	0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x44, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xaf,
	0x02, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x45,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x4a, 0x49, 0x54, 0x41, 0x12, 0x20, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4a, 0x49, 0x54,
	0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4a,
	0x49, 0x54, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x19, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x61, 0x69, 0x73, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x61, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_protobuf_api_proto_rawDescOnce sync.Once
	file_pkg_pb_protobuf_api_proto_rawDescData = file_pkg_pb_protobuf_api_proto_rawDesc
)

func file_pkg_pb_protobuf_api_proto_rawDescGZIP() []byte {
	file_pkg_pb_protobuf_api_proto_rawDescOnce.Do(func() {
		file_pkg_pb_protobuf_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_protobuf_api_proto_rawDescData)
	})
	return file_pkg_pb_protobuf_api_proto_rawDescData
}

var file_pkg_pb_protobuf_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_pb_protobuf_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pkg_pb_protobuf_api_proto_goTypes = []interface{}{
	(AgentState)(0),               // 0: naisdevice.AgentState
	(*ConfigureResponse)(nil),     // 1: naisdevice.ConfigureResponse
	(*UpgradeResponse)(nil),       // 2: naisdevice.UpgradeResponse
	(*ConfigureJITAResponse)(nil), // 3: naisdevice.ConfigureJITAResponse
	(*LoginResponse)(nil),         // 4: naisdevice.LoginResponse
	(*LogoutResponse)(nil),        // 5: naisdevice.LogoutResponse
	(*UpgradeRequest)(nil),        // 6: naisdevice.UpgradeRequest
	(*ConfigureJITARequest)(nil),  // 7: naisdevice.ConfigureJITARequest
	(*LoginRequest)(nil),          // 8: naisdevice.LoginRequest
	(*LogoutRequest)(nil),         // 9: naisdevice.LogoutRequest
	(*AgentStatusRequest)(nil),    // 10: naisdevice.AgentStatusRequest
	(*AgentStatus)(nil),           // 11: naisdevice.AgentStatus
	(*Configuration)(nil),         // 12: naisdevice.Configuration
	(*Gateway)(nil),               // 13: naisdevice.Gateway
	(*Error)(nil),                 // 14: naisdevice.Error
	(*timestamppb.Timestamp)(nil), // 15: google.protobuf.Timestamp
}
var file_pkg_pb_protobuf_api_proto_depIdxs = []int32{
	13, // 0: naisdevice.ConfigureJITARequest.gateway:type_name -> naisdevice.Gateway
	0,  // 1: naisdevice.AgentStatus.connectionState:type_name -> naisdevice.AgentState
	15, // 2: naisdevice.AgentStatus.connectedSince:type_name -> google.protobuf.Timestamp
	13, // 3: naisdevice.AgentStatus.Gateways:type_name -> naisdevice.Gateway
	13, // 4: naisdevice.Configuration.Gateways:type_name -> naisdevice.Gateway
	12, // 5: naisdevice.DeviceHelper.Configure:input_type -> naisdevice.Configuration
	6,  // 6: naisdevice.DeviceHelper.Upgrade:input_type -> naisdevice.UpgradeRequest
	10, // 7: naisdevice.DeviceAgent.Status:input_type -> naisdevice.AgentStatusRequest
	7,  // 8: naisdevice.DeviceAgent.ConfigureJITA:input_type -> naisdevice.ConfigureJITARequest
	8,  // 9: naisdevice.DeviceAgent.Login:input_type -> naisdevice.LoginRequest
	9,  // 10: naisdevice.DeviceAgent.Logout:input_type -> naisdevice.LogoutRequest
	1,  // 11: naisdevice.DeviceHelper.Configure:output_type -> naisdevice.ConfigureResponse
	2,  // 12: naisdevice.DeviceHelper.Upgrade:output_type -> naisdevice.UpgradeResponse
	11, // 13: naisdevice.DeviceAgent.Status:output_type -> naisdevice.AgentStatus
	3,  // 14: naisdevice.DeviceAgent.ConfigureJITA:output_type -> naisdevice.ConfigureJITAResponse
	4,  // 15: naisdevice.DeviceAgent.Login:output_type -> naisdevice.LoginResponse
	5,  // 16: naisdevice.DeviceAgent.Logout:output_type -> naisdevice.LogoutResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_pb_protobuf_api_proto_init() }
func file_pkg_pb_protobuf_api_proto_init() {
	if File_pkg_pb_protobuf_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_protobuf_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureJITAResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureJITARequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gateway); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_protobuf_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_protobuf_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_pb_protobuf_api_proto_goTypes,
		DependencyIndexes: file_pkg_pb_protobuf_api_proto_depIdxs,
		EnumInfos:         file_pkg_pb_protobuf_api_proto_enumTypes,
		MessageInfos:      file_pkg_pb_protobuf_api_proto_msgTypes,
	}.Build()
	File_pkg_pb_protobuf_api_proto = out.File
	file_pkg_pb_protobuf_api_proto_rawDesc = nil
	file_pkg_pb_protobuf_api_proto_goTypes = nil
	file_pkg_pb_protobuf_api_proto_depIdxs = nil
}
