// Code generated by protoc-gen-go. DO NOT EDIT.
// source: syncer.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	syncer.proto

It has these top-level messages:
	PullRequest
	SyncData
	SyncService
	SyncInstance
	Expansion
	HealthCheck
	MappingEntry
*/
package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto1 "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion3 // please upgrade the proto package

type SyncService_Status int32

const (
	SyncService_UNKNOWN SyncService_Status = 0
	SyncService_UP      SyncService_Status = 1
	SyncService_DOWN    SyncService_Status = 2
)

var SyncService_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "UP",
	2: "DOWN",
}

var SyncService_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"UP":      1,
	"DOWN":    2,
}

func (x SyncService_Status) String() string {
	return proto1.EnumName(SyncService_Status_name, int32(x))
}

func (SyncService_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{5, 0}
}

type SyncInstance_Status int32

const (
	SyncInstance_UNKNOWN      SyncInstance_Status = 0
	SyncInstance_UP           SyncInstance_Status = 1
	SyncInstance_DOWN         SyncInstance_Status = 2
	SyncInstance_STARTING     SyncInstance_Status = 3
	SyncInstance_OUTOFSERVICE SyncInstance_Status = 4
)

var SyncInstance_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "UP",
	2: "DOWN",
	3: "STARTING",
	4: "OUTOFSERVICE",
}

var SyncInstance_Status_value = map[string]int32{
	"UNKNOWN":      0,
	"UP":           1,
	"DOWN":         2,
	"STARTING":     3,
	"OUTOFSERVICE": 4,
}

func (x SyncInstance_Status) String() string {
	return proto1.EnumName(SyncInstance_Status_name, int32(x))
}

func (SyncInstance_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{6, 0}
}

type HealthCheck_Modes int32

const (
	HealthCheck_UNKNOWN HealthCheck_Modes = 0
	HealthCheck_PUSH    HealthCheck_Modes = 1
	HealthCheck_PULL    HealthCheck_Modes = 2
)

var HealthCheck_Modes_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUSH",
	2: "PULL",
}

var HealthCheck_Modes_value = map[string]int32{
	"UNKNOWN": 0,
	"PUSH":    1,
	"PULL":    2,
}

func (x HealthCheck_Modes) String() string {
	return proto1.EnumName(HealthCheck_Modes_name, int32(x))
}

func (HealthCheck_Modes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{8, 0}
}

type PullRequest struct {
	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Options     string `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	Time        string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (m *PullRequest) Reset()         { *m = PullRequest{} }
func (m *PullRequest) String() string { return proto1.CompactTextString(m) }
func (*PullRequest) ProtoMessage()    {}
func (*PullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{0}
}

func (m *PullRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *PullRequest) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *PullRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type IncrementPullRequest struct {
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *IncrementPullRequest) Reset()         { *m = IncrementPullRequest{} }
func (m *IncrementPullRequest) String() string { return proto1.CompactTextString(m) }
func (*IncrementPullRequest) ProtoMessage()    {}
func (*IncrementPullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{1}
}

func (m *IncrementPullRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type DeclareRequest struct {
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *DeclareRequest) Reset()         { *m = DeclareRequest{} }
func (m *DeclareRequest) String() string { return proto1.CompactTextString(m) }
func (*DeclareRequest) ProtoMessage()    {}
func (*DeclareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{2}
}

func (m *DeclareRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type DeclareResponse struct {
	SyncDataLength int64 `protobuf:"varint,1,opt,name=syncDataLength,proto3" json:"syncDataLength,omitempty"`
}

func (m *DeclareResponse) Reset()         { *m = DeclareResponse{} }
func (m *DeclareResponse) String() string { return proto1.CompactTextString(m) }
func (*DeclareResponse) ProtoMessage()    {}
func (*DeclareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{3}
}

func (m *DeclareResponse) GetSyncDataLength() int64 {
	if m != nil {
		return m.SyncDataLength
	}
	return 0
}

type SyncData struct {
	Services  []*SyncService  `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	Instances []*SyncInstance `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
}

func (m *SyncData) Reset()         { *m = SyncData{} }
func (m *SyncData) String() string { return proto1.CompactTextString(m) }
func (*SyncData) ProtoMessage()    {}
func (*SyncData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{4}
}

func (m *SyncData) GetServices() []*SyncService {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *SyncData) GetInstances() []*SyncInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type SyncService struct {
	ServiceId     string             `protobuf:"bytes,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	App           string             `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	Name          string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version       string             `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Status        SyncService_Status `protobuf:"varint,5,opt,name=status,proto3,enum=proto1.SyncService_Status" json:"status,omitempty"`
	DomainProject string             `protobuf:"bytes,6,opt,name=domainProject,proto3" json:"domainProject,omitempty"`
	Environment   string             `protobuf:"bytes,7,opt,name=environment,proto3" json:"environment,omitempty"`
	PluginName    string             `protobuf:"bytes,8,opt,name=pluginName,proto3" json:"pluginName,omitempty"`
	Expansions    []*Expansion       `protobuf:"bytes,9,rep,name=expansions,proto3" json:"expansions,omitempty"`
}

func (m *SyncService) Reset()         { *m = SyncService{} }
func (m *SyncService) String() string { return proto1.CompactTextString(m) }
func (*SyncService) ProtoMessage()    {}
func (*SyncService) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{5}
}

func (m *SyncService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *SyncService) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *SyncService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SyncService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SyncService) GetStatus() SyncService_Status {
	if m != nil {
		return m.Status
	}
	return SyncService_UNKNOWN
}

func (m *SyncService) GetDomainProject() string {
	if m != nil {
		return m.DomainProject
	}
	return ""
}

func (m *SyncService) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *SyncService) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *SyncService) GetExpansions() []*Expansion {
	if m != nil {
		return m.Expansions
	}
	return nil
}

type SyncInstance struct {
	InstanceId  string              `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	ServiceId   string              `protobuf:"bytes,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Endpoints   []string            `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	HostName    string              `protobuf:"bytes,4,opt,name=hostName,proto3" json:"hostName,omitempty"`
	Status      SyncInstance_Status `protobuf:"varint,5,opt,name=status,proto3,enum=proto1.SyncInstance_Status" json:"status,omitempty"`
	HealthCheck *HealthCheck        `protobuf:"bytes,6,opt,name=healthCheck,proto3" json:"healthCheck,omitempty"`
	Version     string              `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	PluginName  string              `protobuf:"bytes,8,opt,name=pluginName,proto3" json:"pluginName,omitempty"`
	Expansions  []*Expansion        `protobuf:"bytes,9,rep,name=expansions,proto3" json:"expansions,omitempty"`
}

func (m *SyncInstance) Reset()         { *m = SyncInstance{} }
func (m *SyncInstance) String() string { return proto1.CompactTextString(m) }
func (*SyncInstance) ProtoMessage()    {}
func (*SyncInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{6}
}

func (m *SyncInstance) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *SyncInstance) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *SyncInstance) GetEndpoints() []string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *SyncInstance) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *SyncInstance) GetStatus() SyncInstance_Status {
	if m != nil {
		return m.Status
	}
	return SyncInstance_UNKNOWN
}

func (m *SyncInstance) GetHealthCheck() *HealthCheck {
	if m != nil {
		return m.HealthCheck
	}
	return nil
}

func (m *SyncInstance) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SyncInstance) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *SyncInstance) GetExpansions() []*Expansion {
	if m != nil {
		return m.Expansions
	}
	return nil
}

type Expansion struct {
	Kind   string            `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Bytes  []byte            `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Expansion) Reset()         { *m = Expansion{} }
func (m *Expansion) String() string { return proto1.CompactTextString(m) }
func (*Expansion) ProtoMessage()    {}
func (*Expansion) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{7}
}

func (m *Expansion) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Expansion) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *Expansion) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type HealthCheck struct {
	Mode     HealthCheck_Modes `protobuf:"varint,1,opt,name=mode,proto3,enum=proto1.HealthCheck_Modes" json:"mode,omitempty"`
	Port     int32             `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Interval int32             `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
	Times    int32             `protobuf:"varint,4,opt,name=times,proto3" json:"times,omitempty"`
	Url      string            `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto1.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{8}
}

func (m *HealthCheck) GetMode() HealthCheck_Modes {
	if m != nil {
		return m.Mode
	}
	return HealthCheck_UNKNOWN
}

func (m *HealthCheck) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HealthCheck) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *HealthCheck) GetTimes() int32 {
	if m != nil {
		return m.Times
	}
	return 0
}

func (m *HealthCheck) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type MappingEntry struct {
	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	//    Tenant tenant = 2;
	DomainProject string `protobuf:"bytes,2,opt,name=domainProject,proto3" json:"domainProject,omitempty"`
	OrgServiceID  string `protobuf:"bytes,3,opt,name=orgServiceID,proto3" json:"orgServiceID,omitempty"`
	OrgInstanceID string `protobuf:"bytes,4,opt,name=orgInstanceID,proto3" json:"orgInstanceID,omitempty"`
	CurServiceID  string `protobuf:"bytes,5,opt,name=curServiceID,proto3" json:"curServiceID,omitempty"`
	CurInstanceID string `protobuf:"bytes,6,opt,name=curInstanceID,proto3" json:"curInstanceID,omitempty"`
}

func (m *MappingEntry) Reset()         { *m = MappingEntry{} }
func (m *MappingEntry) String() string { return proto1.CompactTextString(m) }
func (*MappingEntry) ProtoMessage()    {}
func (*MappingEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9577b640f2aab197, []int{9}
}

func (m *MappingEntry) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *MappingEntry) GetDomainProject() string {
	if m != nil {
		return m.DomainProject
	}
	return ""
}

func (m *MappingEntry) GetOrgServiceID() string {
	if m != nil {
		return m.OrgServiceID
	}
	return ""
}

func (m *MappingEntry) GetOrgInstanceID() string {
	if m != nil {
		return m.OrgInstanceID
	}
	return ""
}

func (m *MappingEntry) GetCurServiceID() string {
	if m != nil {
		return m.CurServiceID
	}
	return ""
}

func (m *MappingEntry) GetCurInstanceID() string {
	if m != nil {
		return m.CurInstanceID
	}
	return ""
}

func init() {
	proto1.RegisterEnum("proto.SyncService_Status", SyncService_Status_name, SyncService_Status_value)
	proto1.RegisterEnum("proto.SyncInstance_Status", SyncInstance_Status_name, SyncInstance_Status_value)
	proto1.RegisterEnum("proto.HealthCheck_Modes", HealthCheck_Modes_name, HealthCheck_Modes_value)
	proto1.RegisterType((*PullRequest)(nil), "proto.PullRequest")
	proto1.RegisterType((*IncrementPullRequest)(nil), "proto.IncrementPullRequest")
	proto1.RegisterType((*DeclareRequest)(nil), "proto.DeclareRequest")
	proto1.RegisterType((*DeclareResponse)(nil), "proto.DeclareResponse")
	proto1.RegisterType((*SyncData)(nil), "proto.SyncData")
	proto1.RegisterType((*SyncService)(nil), "proto.SyncService")
	proto1.RegisterType((*SyncInstance)(nil), "proto.SyncInstance")
	proto1.RegisterType((*Expansion)(nil), "proto.Expansion")
	proto1.RegisterMapType((map[string]string)(nil), "proto.Expansion.LabelsEntry")
	proto1.RegisterType((*HealthCheck)(nil), "proto.HealthCheck")
	proto1.RegisterType((*MappingEntry)(nil), "proto.MappingEntry")
}

func init() { proto1.RegisterFile("syncer.proto1", fileDescriptor_9577b640f2aab197) }

var fileDescriptor_9577b640f2aab197 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x8e, 0xe3, 0x34,
	0x14, 0x9e, 0xa4, 0xe9, 0xdf, 0x49, 0x77, 0xb6, 0x1c, 0x16, 0x14, 0xca, 0x08, 0x55, 0xd1, 0x0a,
	0x46, 0x08, 0x2a, 0xb6, 0xec, 0x05, 0x0b, 0x17, 0x08, 0xb6, 0xc3, 0x6e, 0xc5, 0x4c, 0xa7, 0x72,
	0xa7, 0x20, 0x21, 0x71, 0x91, 0x49, 0xad, 0x36, 0x4c, 0xea, 0x04, 0xdb, 0xa9, 0x98, 0x07, 0x82,
	0xa7, 0x80, 0x87, 0xe1, 0x96, 0xa7, 0x40, 0x76, 0xdc, 0xd6, 0xe9, 0x54, 0x88, 0x8b, 0xbd, 0x8a,
	0xfd, 0xf9, 0xcb, 0xf1, 0xf9, 0xf9, 0x7c, 0x0e, 0x74, 0xc4, 0x3d, 0x8b, 0x29, 0x1f, 0xe4, 0x3c,
	0x93, 0x19, 0xd6, 0xf5, 0x27, 0xfc, 0x19, 0xfc, 0x69, 0x91, 0xa6, 0x84, 0xfe, 0x5a, 0x50, 0x21,
	0xb1, 0x0f, 0xbe, 0xa0, 0x7c, 0x93, 0xc4, 0x74, 0x12, 0xad, 0x69, 0xe0, 0xf4, 0x9d, 0xf3, 0x36,
	0xb1, 0x21, 0x0c, 0xa0, 0x99, 0xe5, 0x32, 0xc9, 0x98, 0x08, 0x5c, 0x7d, 0xba, 0xdd, 0x22, 0x82,
	0x27, 0x93, 0x35, 0x0d, 0x6a, 0x1a, 0xd6, 0xeb, 0xf0, 0x63, 0x78, 0x32, 0x66, 0x31, 0xa7, 0x6b,
	0xca, 0xa4, 0x7d, 0x0f, 0x82, 0x17, 0x2d, 0x16, 0xdc, 0x5c, 0xa0, 0xd7, 0xe1, 0x53, 0x38, 0x1d,
	0xd1, 0x38, 0x8d, 0x38, 0xfd, 0x2f, 0xd6, 0x0b, 0x78, 0xbc, 0x63, 0x89, 0x3c, 0x63, 0x82, 0xe2,
	0x87, 0x70, 0xaa, 0x42, 0x1b, 0x45, 0x32, 0xba, 0xa4, 0x6c, 0x29, 0x57, 0xfa, 0x87, 0x1a, 0x39,
	0x40, 0xc3, 0x35, 0xb4, 0x66, 0x06, 0xc1, 0x01, 0xb4, 0x4c, 0x54, 0x22, 0x70, 0xfa, 0xb5, 0x73,
	0x7f, 0x88, 0x65, 0x62, 0x06, 0x8a, 0x32, 0x2b, 0x8f, 0xc8, 0x8e, 0x83, 0xcf, 0xa0, 0x9d, 0x30,
	0x21, 0x23, 0xa6, 0x7e, 0x70, 0xf5, 0x0f, 0x6f, 0x5b, 0x3f, 0x8c, 0xcd, 0x19, 0xd9, 0xb3, 0xc2,
	0xbf, 0x5d, 0xf0, 0x2d, 0x63, 0x78, 0x06, 0x6d, 0x63, 0x6e, 0xbc, 0x30, 0x21, 0xed, 0x01, 0xec,
	0x42, 0x2d, 0xca, 0x73, 0x93, 0x53, 0xb5, 0x54, 0xd1, 0xb3, 0x68, 0x9f, 0x4f, 0x66, 0xb2, 0xbf,
	0xa1, 0x5c, 0x24, 0x19, 0x0b, 0xbc, 0x32, 0xfb, 0x66, 0x8b, 0xcf, 0xa0, 0x21, 0x64, 0x24, 0x0b,
	0x11, 0xd4, 0xfb, 0xce, 0xf9, 0xe9, 0xf0, 0xbd, 0x87, 0xe1, 0x0c, 0x66, 0x9a, 0x40, 0x0c, 0x11,
	0x9f, 0xc2, 0xa3, 0x45, 0xb6, 0x8e, 0x12, 0x36, 0xe5, 0xd9, 0x2f, 0x34, 0x96, 0x41, 0x43, 0x9b,
	0xac, 0x82, 0x4a, 0x12, 0x94, 0x6d, 0x12, 0x9e, 0x31, 0x55, 0xc4, 0xa0, 0x59, 0x4a, 0xc2, 0x82,
	0xf0, 0x03, 0x80, 0x3c, 0x2d, 0x96, 0x09, 0xd3, 0x9a, 0x69, 0x69, 0x82, 0x85, 0xe0, 0x67, 0x00,
	0xf4, 0xb7, 0x3c, 0x62, 0x42, 0xab, 0xa6, 0xad, 0x93, 0xd7, 0x35, 0xee, 0x5d, 0x6c, 0x0f, 0x88,
	0xc5, 0x09, 0x3f, 0x82, 0x46, 0xe9, 0x2b, 0xfa, 0xd0, 0x9c, 0x4f, 0xbe, 0x9f, 0x5c, 0xff, 0x38,
	0xe9, 0x9e, 0x60, 0x03, 0xdc, 0xf9, 0xb4, 0xeb, 0x60, 0x0b, 0xbc, 0x91, 0x42, 0xdc, 0xf0, 0xf7,
	0x1a, 0x74, 0xec, 0xfc, 0x2b, 0x5f, 0xb6, 0x15, 0xd8, 0x65, 0xd9, 0x42, 0xaa, 0x45, 0x70, 0x0f,
	0x8b, 0x70, 0x06, 0x6d, 0xca, 0x16, 0x79, 0x96, 0x30, 0x29, 0x82, 0x5a, 0xbf, 0xa6, 0x4e, 0x77,
	0x00, 0xf6, 0xa0, 0xb5, 0xca, 0x84, 0xd4, 0x51, 0x96, 0xd9, 0xdf, 0xed, 0x71, 0x78, 0x90, 0xfe,
	0xde, 0x11, 0x71, 0x1c, 0xe6, 0xff, 0x39, 0xf8, 0x2b, 0x1a, 0xa5, 0x72, 0xf5, 0x72, 0x45, 0xe3,
	0x3b, 0x9d, 0xfd, 0xbd, 0x0c, 0x5f, 0xef, 0x4f, 0x88, 0x4d, 0xb3, 0x25, 0xd0, 0xac, 0x4a, 0xe0,
	0xcd, 0xd7, 0xe1, 0xd5, 0xff, 0xac, 0x03, 0x76, 0xa0, 0x35, 0xbb, 0xf9, 0x86, 0xdc, 0x8c, 0x27,
	0xaf, 0xba, 0x35, 0xec, 0x42, 0xe7, 0x7a, 0x7e, 0x73, 0xfd, 0xdd, 0xec, 0x82, 0xfc, 0x30, 0x7e,
	0x79, 0xd1, 0xf5, 0xc2, 0x3f, 0x1c, 0x68, 0xef, 0xae, 0x50, 0xca, 0xbe, 0x4b, 0xd8, 0xb6, 0x3c,
	0x7a, 0x8d, 0x4f, 0xa0, 0x7e, 0x7b, 0x2f, 0x69, 0xd9, 0x55, 0x3a, 0xa4, 0xdc, 0xe0, 0x73, 0x68,
	0xa4, 0xd1, 0x2d, 0x4d, 0xcb, 0x6a, 0xf8, 0xc3, 0xb3, 0x43, 0x77, 0x07, 0x97, 0xfa, 0xf8, 0x82,
	0x49, 0x7e, 0x4f, 0x0c, 0xb7, 0xf7, 0x02, 0x7c, 0x0b, 0x56, 0x4f, 0xeb, 0x8e, 0xde, 0x9b, 0xdb,
	0xd4, 0x52, 0x5d, 0xb6, 0x89, 0xd2, 0x82, 0x1a, 0x05, 0x94, 0x9b, 0x2f, 0xdd, 0x2f, 0x9c, 0xf0,
	0x2f, 0x07, 0x7c, 0x2b, 0xf5, 0xf8, 0x09, 0x78, 0xeb, 0x6c, 0x51, 0x76, 0xc2, 0xd3, 0x61, 0xf0,
	0xb0, 0x38, 0x83, 0xab, 0x6c, 0x41, 0x05, 0xd1, 0x2c, 0x15, 0x58, 0x9e, 0x71, 0xa9, 0xcd, 0xd6,
	0x89, 0x5e, 0x2b, 0xd5, 0x24, 0x4c, 0x52, 0xbe, 0x89, 0x52, 0xfd, 0x94, 0xeb, 0x64, 0xb7, 0x57,
	0x7e, 0xa8, 0x36, 0x29, 0xb4, 0x9c, 0xea, 0xa4, 0xdc, 0x28, 0x7f, 0x0b, 0x9e, 0x6a, 0x21, 0xb5,
	0x89, 0x5a, 0x86, 0xe7, 0x50, 0xd7, 0xd7, 0x54, 0xcb, 0xd0, 0x02, 0x6f, 0x3a, 0x9f, 0xbd, 0x2e,
	0x0b, 0x31, 0x9d, 0x5f, 0x5e, 0x76, 0xdd, 0xf0, 0x1f, 0x07, 0x3a, 0x57, 0x51, 0x9e, 0x27, 0x6c,
	0x59, 0x06, 0xdf, 0x07, 0x3f, 0x4e, 0x0b, 0x21, 0x29, 0xb7, 0x3b, 0xba, 0x05, 0x3d, 0x6c, 0x03,
	0xee, 0xb1, 0x36, 0x10, 0x42, 0x27, 0xe3, 0x4b, 0xd3, 0x49, 0xc6, 0x23, 0xd3, 0x95, 0x2a, 0x98,
	0xb2, 0x94, 0xf1, 0xe5, 0x56, 0xee, 0xe3, 0x91, 0x79, 0x25, 0x55, 0x50, 0x59, 0x8a, 0x0b, 0xbe,
	0xb7, 0x54, 0xc6, 0x59, 0xc1, 0x94, 0xa5, 0xb8, 0xe0, 0x96, 0x25, 0xd3, 0x9a, 0x2a, 0xe0, 0xf0,
	0x4f, 0x07, 0x3c, 0xf5, 0xc0, 0xf0, 0x53, 0xf0, 0xd4, 0x74, 0xc1, 0xed, 0xe3, 0xb1, 0x46, 0x4d,
	0xef, 0xb1, 0xf5, 0x12, 0x55, 0xeb, 0x0f, 0x4f, 0x70, 0x04, 0x6f, 0x99, 0x19, 0xb2, 0x9f, 0x0e,
	0xf8, 0x8e, 0xe1, 0x55, 0x67, 0x50, 0xef, 0xdd, 0x43, 0xb8, 0x1c, 0x3a, 0xe1, 0x09, 0x7e, 0x0d,
	0x8f, 0x2a, 0xb3, 0x0d, 0xdf, 0x37, 0xd4, 0x63, 0x13, 0xef, 0x88, 0x1b, 0xdf, 0xb6, 0x7f, 0x6a,
	0x0e, 0xbe, 0xd2, 0xe8, 0x6d, 0x43, 0x7f, 0x3e, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0xed, 0xcc,
	0x45, 0x4e, 0xa4, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncClient is the client API for Sync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncClient interface {
	Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*SyncData, error)
	DeclareDataLength(ctx context.Context, in *DeclareRequest, opts ...grpc.CallOption) (*DeclareResponse, error)
	IncrementPull(ctx context.Context, in *IncrementPullRequest, opts ...grpc.CallOption) (*SyncData, error)
}

type syncClient struct {
	cc *grpc.ClientConn
}

func NewSyncClient(cc *grpc.ClientConn) SyncClient {
	return &syncClient{cc}
}

func (c *syncClient) Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*SyncData, error) {
	out := new(SyncData)
	err := c.cc.Invoke(ctx, "/proto.Sync/Pull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncClient) DeclareDataLength(ctx context.Context, in *DeclareRequest, opts ...grpc.CallOption) (*DeclareResponse, error) {
	out := new(DeclareResponse)
	err := c.cc.Invoke(ctx, "/proto.Sync/DeclareDataLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncClient) IncrementPull(ctx context.Context, in *IncrementPullRequest, opts ...grpc.CallOption) (*SyncData, error) {
	out := new(SyncData)
	err := c.cc.Invoke(ctx, "/proto.Sync/IncrementPull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServer is the server API for Sync service.
type SyncServer interface {
	Pull(context.Context, *PullRequest) (*SyncData, error)
	DeclareDataLength(context.Context, *DeclareRequest) (*DeclareResponse, error)
	IncrementPull(context.Context, *IncrementPullRequest) (*SyncData, error)
}

func RegisterSyncServer(s *grpc.Server, srv SyncServer) {
	s.RegisterService(&_Sync_serviceDesc, srv)
}

func _Sync_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sync/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServer).Pull(ctx, req.(*PullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sync_DeclareDataLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServer).DeclareDataLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sync/DeclareDataLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServer).DeclareDataLength(ctx, req.(*DeclareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sync_IncrementPull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementPullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServer).IncrementPull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sync/IncrementPull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServer).IncrementPull(ctx, req.(*IncrementPullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sync",
	HandlerType: (*SyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pull",
			Handler:    _Sync_Pull_Handler,
		},
		{
			MethodName: "DeclareDataLength",
			Handler:    _Sync_DeclareDataLength_Handler,
		},
		{
			MethodName: "IncrementPull",
			Handler:    _Sync_IncrementPull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncer.proto",
}

func init() { proto1.RegisterFile("syncer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xdb, 0x48,
	0x0c, 0x8e, 0x64, 0xf9, 0x8f, 0xf2, 0x66, 0x85, 0xd9, 0x3d, 0x68, 0x8d, 0x60, 0x61, 0x08, 0x0b,
	0xac, 0x0f, 0xbb, 0x46, 0xe3, 0xa6, 0x40, 0xdb, 0x5b, 0x11, 0xbb, 0x89, 0xd1, 0xc4, 0x31, 0xc6,
	0x71, 0x7b, 0xea, 0x41, 0x91, 0x07, 0xb6, 0x1a, 0x79, 0xa4, 0x6a, 0x46, 0x46, 0xfd, 0x40, 0xed,
	0x5b, 0xf4, 0x61, 0x7a, 0xed, 0x53, 0x14, 0xf3, 0x63, 0x7b, 0xe4, 0xe4, 0xd0, 0x43, 0x4f, 0x26,
	0x3f, 0x52, 0x1c, 0x92, 0x1f, 0x49, 0x43, 0x8b, 0x6d, 0x68, 0x44, 0xf2, 0x5e, 0x96, 0xa7, 0x3c,
	0x45, 0x55, 0xf9, 0x13, 0xbc, 0x07, 0x77, 0x52, 0x24, 0x09, 0x26, 0x1f, 0x0b, 0xc2, 0x38, 0xea,
	0x80, 0xcb, 0x48, 0xbe, 0x8e, 0x23, 0x32, 0x0e, 0x57, 0xc4, 0xb7, 0x3a, 0x56, 0xb7, 0x89, 0x4d,
	0x08, 0xf9, 0x50, 0x4f, 0x33, 0x1e, 0xa7, 0x94, 0xf9, 0xb6, 0xb4, 0x6e, 0x55, 0x84, 0xc0, 0xe1,
	0xf1, 0x8a, 0xf8, 0x15, 0x09, 0x4b, 0x39, 0x58, 0x41, 0x63, 0xba, 0xa1, 0xd1, 0x20, 0xe4, 0x21,
	0xea, 0x41, 0x43, 0x07, 0x62, 0xbe, 0xd5, 0xa9, 0x74, 0xdd, 0x3e, 0x52, 0xb9, 0xf4, 0x84, 0xcb,
	0x54, 0x99, 0xf0, 0xce, 0x07, 0x9d, 0x42, 0x33, 0xa6, 0x8c, 0x87, 0x54, 0x7c, 0x60, 0xcb, 0x0f,
	0xfe, 0x30, 0x3e, 0x18, 0x69, 0x1b, 0xde, 0x7b, 0x05, 0xdf, 0x6c, 0x70, 0x8d, 0x60, 0xe8, 0x04,
	0x9a, 0x3a, 0xdc, 0x68, 0xae, 0x8b, 0xd9, 0x03, 0xc8, 0x83, 0x4a, 0x98, 0x65, 0xba, 0x0c, 0x21,
	0x8a, 0x12, 0x68, 0xb8, 0x2f, 0x81, 0xea, 0x82, 0xd7, 0x24, 0x67, 0x71, 0x4a, 0x7d, 0x47, 0x15,
	0xac, 0x55, 0x74, 0x0a, 0x35, 0xc6, 0x43, 0x5e, 0x30, 0xbf, 0xda, 0xb1, 0xba, 0xc7, 0xfd, 0xbf,
	0x1e, 0x96, 0xd3, 0x9b, 0x4a, 0x07, 0xac, 0x1d, 0xd1, 0x3f, 0xf0, 0xdb, 0x3c, 0x5d, 0x85, 0x31,
	0x9d, 0xe4, 0xe9, 0x07, 0x12, 0x71, 0xbf, 0x26, 0x43, 0x96, 0x41, 0xc1, 0x02, 0xa1, 0xeb, 0x38,
	0x4f, 0xe9, 0x8a, 0x50, 0xee, 0xd7, 0x15, 0x0b, 0x06, 0x84, 0xfe, 0x06, 0xc8, 0x92, 0x62, 0x11,
	0x53, 0x49, 0x53, 0x43, 0x3a, 0x18, 0x08, 0x7a, 0x02, 0x40, 0x3e, 0x65, 0x21, 0x65, 0x92, 0xa8,
	0xa6, 0x6c, 0x9e, 0xa7, 0xd3, 0x1b, 0x6e, 0x0d, 0xd8, 0xf0, 0x09, 0xfe, 0x85, 0x9a, 0xca, 0x15,
	0xb9, 0x50, 0x9f, 0x8d, 0xdf, 0x8c, 0x6f, 0xde, 0x8d, 0xbd, 0x23, 0x54, 0x03, 0x7b, 0x36, 0xf1,
	0x2c, 0xd4, 0x00, 0x67, 0x20, 0x10, 0x3b, 0xf8, 0x5c, 0x81, 0x96, 0xd9, 0x7f, 0x91, 0xcb, 0x96,
	0x81, 0x5d, 0x97, 0x0d, 0xa4, 0x4c, 0x82, 0x7d, 0x48, 0xc2, 0x09, 0x34, 0x09, 0x9d, 0x67, 0x69,
	0x4c, 0x39, 0xf3, 0x2b, 0x9d, 0x8a, 0xb0, 0xee, 0x00, 0xd4, 0x86, 0xc6, 0x32, 0x65, 0x5c, 0x56,
	0xa9, 0xba, 0xbf, 0xd3, 0x51, 0xff, 0xa0, 0xfd, 0xed, 0x47, 0x86, 0xe3, 0xb0, 0xff, 0x67, 0xe0,
	0x2e, 0x49, 0x98, 0xf0, 0xe5, 0xf9, 0x92, 0x44, 0xf7, 0xb2, 0xfb, 0xfb, 0x31, 0xbc, 0xdc, 0x5b,
	0xb0, 0xe9, 0x66, 0x8e, 0x40, 0xbd, 0x3c, 0x02, 0xbf, 0x9e, 0x87, 0x8b, 0x9f, 0xe4, 0x01, 0xb5,
	0xa0, 0x31, 0xbd, 0x7d, 0x85, 0x6f, 0x47, 0xe3, 0x0b, 0xaf, 0x82, 0x3c, 0x68, 0xdd, 0xcc, 0x6e,
	0x6f, 0x5e, 0x4f, 0x87, 0xf8, 0xed, 0xe8, 0x7c, 0xe8, 0x39, 0xc1, 0x17, 0x0b, 0x9a, 0xbb, 0x27,
	0xc4, 0x64, 0xdf, 0xc7, 0x74, 0x4b, 0x8f, 0x94, 0xd1, 0x9f, 0x50, 0xbd, 0xdb, 0x70, 0xa2, 0x16,
	0xb9, 0x85, 0x95, 0x82, 0xce, 0xa0, 0x96, 0x84, 0x77, 0x24, 0x51, 0x6c, 0xb8, 0xfd, 0x93, 0xc3,
	0x74, 0x7b, 0x57, 0xd2, 0x3c, 0xa4, 0x3c, 0xdf, 0x60, 0xed, 0xdb, 0x7e, 0x01, 0xae, 0x01, 0x8b,
	0xd5, 0xba, 0x27, 0x1b, 0xfd, 0x9a, 0x10, 0xc5, 0x63, 0xeb, 0x30, 0x29, 0x88, 0x9e, 0x00, 0xa5,
	0xbc, 0xb4, 0x9f, 0x5b, 0xc1, 0x57, 0x0b, 0x5c, 0xa3, 0xf5, 0xe8, 0x3f, 0x70, 0x56, 0xe9, 0x5c,
	0x1d, 0x9f, 0xe3, 0xbe, 0xff, 0x90, 0x9c, 0xde, 0x75, 0x3a, 0x27, 0x0c, 0x4b, 0x2f, 0x51, 0x58,
	0x96, 0xe6, 0x5c, 0x86, 0xad, 0x62, 0x29, 0x8b, 0xa9, 0x89, 0x29, 0x27, 0xf9, 0x3a, 0x4c, 0xe4,
	0x2a, 0x57, 0xf1, 0x4e, 0x17, 0x79, 0x88, 0xcb, 0xc4, 0xe4, 0x38, 0x55, 0xb1, 0x52, 0x44, 0xbe,
	0x45, 0x9e, 0xc8, 0x41, 0x6a, 0x62, 0x21, 0x06, 0x5d, 0xa8, 0xca, 0x67, 0xca, 0x34, 0x34, 0xc0,
	0x99, 0xcc, 0xa6, 0x97, 0x8a, 0x88, 0xc9, 0xec, 0xea, 0xca, 0xb3, 0x83, 0xef, 0x16, 0xb4, 0xae,
	0xc3, 0x2c, 0x8b, 0xe9, 0x42, 0x15, 0xdf, 0x01, 0x37, 0x4a, 0x0a, 0xc6, 0x49, 0x6e, 0x1e, 0x51,
	0x03, 0x7a, 0x78, 0x06, 0xec, 0xc7, 0xce, 0x40, 0x00, 0xad, 0x34, 0x5f, 0xe8, 0x4b, 0x32, 0x1a,
	0xe8, 0xab, 0x54, 0xc2, 0x44, 0xa4, 0x34, 0x5f, 0x6c, 0xc7, 0x7d, 0x34, 0xd0, 0x5b, 0x52, 0x06,
	0x45, 0xa4, 0xa8, 0xc8, 0xf7, 0x91, 0x54, 0x9d, 0x25, 0x4c, 0x44, 0x8a, 0x8a, 0xdc, 0x88, 0xa4,
	0x4f, 0x53, 0x09, 0xec, 0x3f, 0x03, 0x47, 0xec, 0x17, 0xfa, 0x1f, 0x1c, 0xf1, 0xbf, 0x81, 0xb6,
	0xbb, 0x63, 0xfc, 0x89, 0xb4, 0x7f, 0x37, 0x16, 0x51, 0x5c, 0xfe, 0xe0, 0xe8, 0xae, 0x26, 0x91,
	0xa7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0xe0, 0xda, 0xd0, 0x84, 0x06, 0x00, 0x00,
}
