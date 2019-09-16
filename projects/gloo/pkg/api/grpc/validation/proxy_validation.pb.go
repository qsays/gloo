// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/grpc/validation/proxy_validation.proto

package validation

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListenerReport_Error_Type int32

const (
	ListenerReport_Error_NameNotUniqueError     ListenerReport_Error_Type = 0
	ListenerReport_Error_BindPortNotUniqueError ListenerReport_Error_Type = 1
	ListenerReport_Error_SSLConfigError         ListenerReport_Error_Type = 2
	ListenerReport_Error_ProcessingError        ListenerReport_Error_Type = 3
)

var ListenerReport_Error_Type_name = map[int32]string{
	0: "NameNotUniqueError",
	1: "BindPortNotUniqueError",
	2: "SSLConfigError",
	3: "ProcessingError",
}

var ListenerReport_Error_Type_value = map[string]int32{
	"NameNotUniqueError":     0,
	"BindPortNotUniqueError": 1,
	"SSLConfigError":         2,
	"ProcessingError":        3,
}

func (x ListenerReport_Error_Type) String() string {
	return proto.EnumName(ListenerReport_Error_Type_name, int32(x))
}

func (ListenerReport_Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{1, 0, 0}
}

type HttpListenerReport_Error_Type int32

const (
	HttpListenerReport_Error_ProcessingError HttpListenerReport_Error_Type = 0
)

var HttpListenerReport_Error_Type_name = map[int32]string{
	0: "ProcessingError",
}

var HttpListenerReport_Error_Type_value = map[string]int32{
	"ProcessingError": 0,
}

func (x HttpListenerReport_Error_Type) String() string {
	return proto.EnumName(HttpListenerReport_Error_Type_name, int32(x))
}

func (HttpListenerReport_Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{2, 0, 0}
}

type VirtualHostReport_Error_Type int32

const (
	VirtualHostReport_Error_NameNotUniqueError    VirtualHostReport_Error_Type = 0
	VirtualHostReport_Error_DomainsNotUniqueError VirtualHostReport_Error_Type = 1
	VirtualHostReport_Error_ProcessingError       VirtualHostReport_Error_Type = 2
)

var VirtualHostReport_Error_Type_name = map[int32]string{
	0: "NameNotUniqueError",
	1: "DomainsNotUniqueError",
	2: "ProcessingError",
}

var VirtualHostReport_Error_Type_value = map[string]int32{
	"NameNotUniqueError":    0,
	"DomainsNotUniqueError": 1,
	"ProcessingError":       2,
}

func (x VirtualHostReport_Error_Type) String() string {
	return proto.EnumName(VirtualHostReport_Error_Type_name, int32(x))
}

func (VirtualHostReport_Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{3, 0, 0}
}

type RouteReport_Error_Type int32

const (
	RouteReport_Error_InvalidMatcherError     RouteReport_Error_Type = 0
	RouteReport_Error_InvalidDestinationError RouteReport_Error_Type = 1
	RouteReport_Error_ProcessingError         RouteReport_Error_Type = 2
)

var RouteReport_Error_Type_name = map[int32]string{
	0: "InvalidMatcherError",
	1: "InvalidDestinationError",
	2: "ProcessingError",
}

var RouteReport_Error_Type_value = map[string]int32{
	"InvalidMatcherError":     0,
	"InvalidDestinationError": 1,
	"ProcessingError":         2,
}

func (x RouteReport_Error_Type) String() string {
	return proto.EnumName(RouteReport_Error_Type_name, int32(x))
}

func (RouteReport_Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{4, 0, 0}
}

type TcpListenerReport_Error_Type int32

const (
	TcpListenerReport_Error_NameNotUniqueError     TcpListenerReport_Error_Type = 0
	TcpListenerReport_Error_BindPortNotUniqueError TcpListenerReport_Error_Type = 1
	TcpListenerReport_Error_SSLConfigError         TcpListenerReport_Error_Type = 2
	TcpListenerReport_Error_ProcessingError        TcpListenerReport_Error_Type = 3
)

var TcpListenerReport_Error_Type_name = map[int32]string{
	0: "NameNotUniqueError",
	1: "BindPortNotUniqueError",
	2: "SSLConfigError",
	3: "ProcessingError",
}

var TcpListenerReport_Error_Type_value = map[string]int32{
	"NameNotUniqueError":     0,
	"BindPortNotUniqueError": 1,
	"SSLConfigError":         2,
	"ProcessingError":        3,
}

func (x TcpListenerReport_Error_Type) String() string {
	return proto.EnumName(TcpListenerReport_Error_Type_name, int32(x))
}

func (TcpListenerReport_Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{5, 0, 0}
}

type TcpHostReport_Error_Type int32

const (
	TcpHostReport_Error_NameNotUniqueError      TcpHostReport_Error_Type = 0
	TcpHostReport_Error_InvalidDestinationError TcpHostReport_Error_Type = 1
	TcpHostReport_Error_ProcessingError         TcpHostReport_Error_Type = 2
)

var TcpHostReport_Error_Type_name = map[int32]string{
	0: "NameNotUniqueError",
	1: "InvalidDestinationError",
	2: "ProcessingError",
}

var TcpHostReport_Error_Type_value = map[string]int32{
	"NameNotUniqueError":      0,
	"InvalidDestinationError": 1,
	"ProcessingError":         2,
}

func (x TcpHostReport_Error_Type) String() string {
	return proto.EnumName(TcpHostReport_Error_Type_name, int32(x))
}

func (TcpHostReport_Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{6, 0, 0}
}

//
// The Proxy Report should contain one report for each sub-resource of the Proxy
// E.g., each listener will have a corresponding report. Within each listener report is
// a route report corresponding to each route on the listener.
//
// If the report contains no errors, the (sub-)resource is valid.
type ProxyReport struct {
	ListenerReports      []*ListenerReport `protobuf:"bytes,1,rep,name=listener_reports,json=listenerReports,proto3" json:"listener_reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProxyReport) Reset()         { *m = ProxyReport{} }
func (m *ProxyReport) String() string { return proto.CompactTextString(m) }
func (*ProxyReport) ProtoMessage()    {}
func (*ProxyReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{0}
}
func (m *ProxyReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyReport.Unmarshal(m, b)
}
func (m *ProxyReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyReport.Marshal(b, m, deterministic)
}
func (m *ProxyReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyReport.Merge(m, src)
}
func (m *ProxyReport) XXX_Size() int {
	return xxx_messageInfo_ProxyReport.Size(m)
}
func (m *ProxyReport) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyReport.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyReport proto.InternalMessageInfo

func (m *ProxyReport) GetListenerReports() []*ListenerReport {
	if m != nil {
		return m.ListenerReports
	}
	return nil
}

type ListenerReport struct {
	// errors on top-level config of the listener
	Errors []*ListenerReport_Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// Types that are valid to be assigned to ListenerTypeReport:
	//	*ListenerReport_HttpListenerReport
	//	*ListenerReport_TcpListenerReport
	ListenerTypeReport   isListenerReport_ListenerTypeReport `protobuf_oneof:"listener_type_report"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ListenerReport) Reset()         { *m = ListenerReport{} }
func (m *ListenerReport) String() string { return proto.CompactTextString(m) }
func (*ListenerReport) ProtoMessage()    {}
func (*ListenerReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{1}
}
func (m *ListenerReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerReport.Unmarshal(m, b)
}
func (m *ListenerReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerReport.Marshal(b, m, deterministic)
}
func (m *ListenerReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerReport.Merge(m, src)
}
func (m *ListenerReport) XXX_Size() int {
	return xxx_messageInfo_ListenerReport.Size(m)
}
func (m *ListenerReport) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerReport.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerReport proto.InternalMessageInfo

type isListenerReport_ListenerTypeReport interface {
	isListenerReport_ListenerTypeReport()
}

type ListenerReport_HttpListenerReport struct {
	HttpListenerReport *HttpListenerReport `protobuf:"bytes,3,opt,name=http_listener_report,json=httpListenerReport,proto3,oneof" json:"http_listener_report,omitempty"`
}
type ListenerReport_TcpListenerReport struct {
	TcpListenerReport *TcpListenerReport `protobuf:"bytes,4,opt,name=tcp_listener_report,json=tcpListenerReport,proto3,oneof" json:"tcp_listener_report,omitempty"`
}

func (*ListenerReport_HttpListenerReport) isListenerReport_ListenerTypeReport() {}
func (*ListenerReport_TcpListenerReport) isListenerReport_ListenerTypeReport()  {}

func (m *ListenerReport) GetListenerTypeReport() isListenerReport_ListenerTypeReport {
	if m != nil {
		return m.ListenerTypeReport
	}
	return nil
}

func (m *ListenerReport) GetErrors() []*ListenerReport_Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *ListenerReport) GetHttpListenerReport() *HttpListenerReport {
	if x, ok := m.GetListenerTypeReport().(*ListenerReport_HttpListenerReport); ok {
		return x.HttpListenerReport
	}
	return nil
}

func (m *ListenerReport) GetTcpListenerReport() *TcpListenerReport {
	if x, ok := m.GetListenerTypeReport().(*ListenerReport_TcpListenerReport); ok {
		return x.TcpListenerReport
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerReport) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerReport_HttpListenerReport)(nil),
		(*ListenerReport_TcpListenerReport)(nil),
	}
}

// error types for top-level listener config
type ListenerReport_Error struct {
	// the type of the error
	Type ListenerReport_Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=gloo.solo.io.ListenerReport_Error_Type" json:"type,omitempty"`
	// any extra info as a string
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenerReport_Error) Reset()         { *m = ListenerReport_Error{} }
func (m *ListenerReport_Error) String() string { return proto.CompactTextString(m) }
func (*ListenerReport_Error) ProtoMessage()    {}
func (*ListenerReport_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{1, 0}
}
func (m *ListenerReport_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerReport_Error.Unmarshal(m, b)
}
func (m *ListenerReport_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerReport_Error.Marshal(b, m, deterministic)
}
func (m *ListenerReport_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerReport_Error.Merge(m, src)
}
func (m *ListenerReport_Error) XXX_Size() int {
	return xxx_messageInfo_ListenerReport_Error.Size(m)
}
func (m *ListenerReport_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerReport_Error.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerReport_Error proto.InternalMessageInfo

func (m *ListenerReport_Error) GetType() ListenerReport_Error_Type {
	if m != nil {
		return m.Type
	}
	return ListenerReport_Error_NameNotUniqueError
}

func (m *ListenerReport_Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type HttpListenerReport struct {
	Errors []*HttpListenerReport_Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	// report for nested virtual hosts
	VirtualHostReports   []*VirtualHostReport `protobuf:"bytes,2,rep,name=virtual_host_reports,json=virtualHostReports,proto3" json:"virtual_host_reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HttpListenerReport) Reset()         { *m = HttpListenerReport{} }
func (m *HttpListenerReport) String() string { return proto.CompactTextString(m) }
func (*HttpListenerReport) ProtoMessage()    {}
func (*HttpListenerReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{2}
}
func (m *HttpListenerReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpListenerReport.Unmarshal(m, b)
}
func (m *HttpListenerReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpListenerReport.Marshal(b, m, deterministic)
}
func (m *HttpListenerReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpListenerReport.Merge(m, src)
}
func (m *HttpListenerReport) XXX_Size() int {
	return xxx_messageInfo_HttpListenerReport.Size(m)
}
func (m *HttpListenerReport) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpListenerReport.DiscardUnknown(m)
}

var xxx_messageInfo_HttpListenerReport proto.InternalMessageInfo

func (m *HttpListenerReport) GetErrors() []*HttpListenerReport_Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *HttpListenerReport) GetVirtualHostReports() []*VirtualHostReport {
	if m != nil {
		return m.VirtualHostReports
	}
	return nil
}

// error types for top-level http listener config
type HttpListenerReport_Error struct {
	// the type of the error
	Type HttpListenerReport_Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=gloo.solo.io.HttpListenerReport_Error_Type" json:"type,omitempty"`
	// any extra info as a string
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpListenerReport_Error) Reset()         { *m = HttpListenerReport_Error{} }
func (m *HttpListenerReport_Error) String() string { return proto.CompactTextString(m) }
func (*HttpListenerReport_Error) ProtoMessage()    {}
func (*HttpListenerReport_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{2, 0}
}
func (m *HttpListenerReport_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpListenerReport_Error.Unmarshal(m, b)
}
func (m *HttpListenerReport_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpListenerReport_Error.Marshal(b, m, deterministic)
}
func (m *HttpListenerReport_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpListenerReport_Error.Merge(m, src)
}
func (m *HttpListenerReport_Error) XXX_Size() int {
	return xxx_messageInfo_HttpListenerReport_Error.Size(m)
}
func (m *HttpListenerReport_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpListenerReport_Error.DiscardUnknown(m)
}

var xxx_messageInfo_HttpListenerReport_Error proto.InternalMessageInfo

func (m *HttpListenerReport_Error) GetType() HttpListenerReport_Error_Type {
	if m != nil {
		return m.Type
	}
	return HttpListenerReport_Error_ProcessingError
}

func (m *HttpListenerReport_Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type VirtualHostReport struct {
	// errors on top-level config of the virtual host
	Errors               []*VirtualHostReport_Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	RouteReports         []*RouteReport             `protobuf:"bytes,2,rep,name=route_reports,json=routeReports,proto3" json:"route_reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *VirtualHostReport) Reset()         { *m = VirtualHostReport{} }
func (m *VirtualHostReport) String() string { return proto.CompactTextString(m) }
func (*VirtualHostReport) ProtoMessage()    {}
func (*VirtualHostReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{3}
}
func (m *VirtualHostReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHostReport.Unmarshal(m, b)
}
func (m *VirtualHostReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHostReport.Marshal(b, m, deterministic)
}
func (m *VirtualHostReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHostReport.Merge(m, src)
}
func (m *VirtualHostReport) XXX_Size() int {
	return xxx_messageInfo_VirtualHostReport.Size(m)
}
func (m *VirtualHostReport) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHostReport.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHostReport proto.InternalMessageInfo

func (m *VirtualHostReport) GetErrors() []*VirtualHostReport_Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *VirtualHostReport) GetRouteReports() []*RouteReport {
	if m != nil {
		return m.RouteReports
	}
	return nil
}

// error types for top-level virtual host config
type VirtualHostReport_Error struct {
	// the type of the error
	Type VirtualHostReport_Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=gloo.solo.io.VirtualHostReport_Error_Type" json:"type,omitempty"`
	// any extra info as a string
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualHostReport_Error) Reset()         { *m = VirtualHostReport_Error{} }
func (m *VirtualHostReport_Error) String() string { return proto.CompactTextString(m) }
func (*VirtualHostReport_Error) ProtoMessage()    {}
func (*VirtualHostReport_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{3, 0}
}
func (m *VirtualHostReport_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHostReport_Error.Unmarshal(m, b)
}
func (m *VirtualHostReport_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHostReport_Error.Marshal(b, m, deterministic)
}
func (m *VirtualHostReport_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHostReport_Error.Merge(m, src)
}
func (m *VirtualHostReport_Error) XXX_Size() int {
	return xxx_messageInfo_VirtualHostReport_Error.Size(m)
}
func (m *VirtualHostReport_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHostReport_Error.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHostReport_Error proto.InternalMessageInfo

func (m *VirtualHostReport_Error) GetType() VirtualHostReport_Error_Type {
	if m != nil {
		return m.Type
	}
	return VirtualHostReport_Error_NameNotUniqueError
}

func (m *VirtualHostReport_Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type RouteReport struct {
	// errors on the config of the route
	Errors               []*RouteReport_Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RouteReport) Reset()         { *m = RouteReport{} }
func (m *RouteReport) String() string { return proto.CompactTextString(m) }
func (*RouteReport) ProtoMessage()    {}
func (*RouteReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{4}
}
func (m *RouteReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteReport.Unmarshal(m, b)
}
func (m *RouteReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteReport.Marshal(b, m, deterministic)
}
func (m *RouteReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteReport.Merge(m, src)
}
func (m *RouteReport) XXX_Size() int {
	return xxx_messageInfo_RouteReport.Size(m)
}
func (m *RouteReport) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteReport.DiscardUnknown(m)
}

var xxx_messageInfo_RouteReport proto.InternalMessageInfo

func (m *RouteReport) GetErrors() []*RouteReport_Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

// error types for the given route config
type RouteReport_Error struct {
	// the type of the error
	Type RouteReport_Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=gloo.solo.io.RouteReport_Error_Type" json:"type,omitempty"`
	// any extra info as a string
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteReport_Error) Reset()         { *m = RouteReport_Error{} }
func (m *RouteReport_Error) String() string { return proto.CompactTextString(m) }
func (*RouteReport_Error) ProtoMessage()    {}
func (*RouteReport_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{4, 0}
}
func (m *RouteReport_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteReport_Error.Unmarshal(m, b)
}
func (m *RouteReport_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteReport_Error.Marshal(b, m, deterministic)
}
func (m *RouteReport_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteReport_Error.Merge(m, src)
}
func (m *RouteReport_Error) XXX_Size() int {
	return xxx_messageInfo_RouteReport_Error.Size(m)
}
func (m *RouteReport_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteReport_Error.DiscardUnknown(m)
}

var xxx_messageInfo_RouteReport_Error proto.InternalMessageInfo

func (m *RouteReport_Error) GetType() RouteReport_Error_Type {
	if m != nil {
		return m.Type
	}
	return RouteReport_Error_InvalidMatcherError
}

func (m *RouteReport_Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type TcpListenerReport struct {
	// errors on top-level config of the listener
	Errors               []*TcpListenerReport_Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	TcpHostReports       []*TcpHostReport           `protobuf:"bytes,2,rep,name=tcp_host_reports,json=tcpHostReports,proto3" json:"tcp_host_reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpListenerReport) Reset()         { *m = TcpListenerReport{} }
func (m *TcpListenerReport) String() string { return proto.CompactTextString(m) }
func (*TcpListenerReport) ProtoMessage()    {}
func (*TcpListenerReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{5}
}
func (m *TcpListenerReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpListenerReport.Unmarshal(m, b)
}
func (m *TcpListenerReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpListenerReport.Marshal(b, m, deterministic)
}
func (m *TcpListenerReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpListenerReport.Merge(m, src)
}
func (m *TcpListenerReport) XXX_Size() int {
	return xxx_messageInfo_TcpListenerReport.Size(m)
}
func (m *TcpListenerReport) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpListenerReport.DiscardUnknown(m)
}

var xxx_messageInfo_TcpListenerReport proto.InternalMessageInfo

func (m *TcpListenerReport) GetErrors() []*TcpListenerReport_Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *TcpListenerReport) GetTcpHostReports() []*TcpHostReport {
	if m != nil {
		return m.TcpHostReports
	}
	return nil
}

// error types for top-level tcp listener config
type TcpListenerReport_Error struct {
	// the type of the error
	Type TcpListenerReport_Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=gloo.solo.io.TcpListenerReport_Error_Type" json:"type,omitempty"`
	// any extra info as a string
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpListenerReport_Error) Reset()         { *m = TcpListenerReport_Error{} }
func (m *TcpListenerReport_Error) String() string { return proto.CompactTextString(m) }
func (*TcpListenerReport_Error) ProtoMessage()    {}
func (*TcpListenerReport_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{5, 0}
}
func (m *TcpListenerReport_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpListenerReport_Error.Unmarshal(m, b)
}
func (m *TcpListenerReport_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpListenerReport_Error.Marshal(b, m, deterministic)
}
func (m *TcpListenerReport_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpListenerReport_Error.Merge(m, src)
}
func (m *TcpListenerReport_Error) XXX_Size() int {
	return xxx_messageInfo_TcpListenerReport_Error.Size(m)
}
func (m *TcpListenerReport_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpListenerReport_Error.DiscardUnknown(m)
}

var xxx_messageInfo_TcpListenerReport_Error proto.InternalMessageInfo

func (m *TcpListenerReport_Error) GetType() TcpListenerReport_Error_Type {
	if m != nil {
		return m.Type
	}
	return TcpListenerReport_Error_NameNotUniqueError
}

func (m *TcpListenerReport_Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type TcpHostReport struct {
	// errors on the tcp host
	Errors               []*TcpHostReport_Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TcpHostReport) Reset()         { *m = TcpHostReport{} }
func (m *TcpHostReport) String() string { return proto.CompactTextString(m) }
func (*TcpHostReport) ProtoMessage()    {}
func (*TcpHostReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{6}
}
func (m *TcpHostReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpHostReport.Unmarshal(m, b)
}
func (m *TcpHostReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpHostReport.Marshal(b, m, deterministic)
}
func (m *TcpHostReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpHostReport.Merge(m, src)
}
func (m *TcpHostReport) XXX_Size() int {
	return xxx_messageInfo_TcpHostReport.Size(m)
}
func (m *TcpHostReport) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpHostReport.DiscardUnknown(m)
}

var xxx_messageInfo_TcpHostReport proto.InternalMessageInfo

func (m *TcpHostReport) GetErrors() []*TcpHostReport_Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

// error types for tcp host config
type TcpHostReport_Error struct {
	// the type of the error
	Type TcpHostReport_Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=gloo.solo.io.TcpHostReport_Error_Type" json:"type,omitempty"`
	// any extra info as a string
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpHostReport_Error) Reset()         { *m = TcpHostReport_Error{} }
func (m *TcpHostReport_Error) String() string { return proto.CompactTextString(m) }
func (*TcpHostReport_Error) ProtoMessage()    {}
func (*TcpHostReport_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4537dae4069b18, []int{6, 0}
}
func (m *TcpHostReport_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpHostReport_Error.Unmarshal(m, b)
}
func (m *TcpHostReport_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpHostReport_Error.Marshal(b, m, deterministic)
}
func (m *TcpHostReport_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpHostReport_Error.Merge(m, src)
}
func (m *TcpHostReport_Error) XXX_Size() int {
	return xxx_messageInfo_TcpHostReport_Error.Size(m)
}
func (m *TcpHostReport_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpHostReport_Error.DiscardUnknown(m)
}

var xxx_messageInfo_TcpHostReport_Error proto.InternalMessageInfo

func (m *TcpHostReport_Error) GetType() TcpHostReport_Error_Type {
	if m != nil {
		return m.Type
	}
	return TcpHostReport_Error_NameNotUniqueError
}

func (m *TcpHostReport_Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterEnum("gloo.solo.io.ListenerReport_Error_Type", ListenerReport_Error_Type_name, ListenerReport_Error_Type_value)
	proto.RegisterEnum("gloo.solo.io.HttpListenerReport_Error_Type", HttpListenerReport_Error_Type_name, HttpListenerReport_Error_Type_value)
	proto.RegisterEnum("gloo.solo.io.VirtualHostReport_Error_Type", VirtualHostReport_Error_Type_name, VirtualHostReport_Error_Type_value)
	proto.RegisterEnum("gloo.solo.io.RouteReport_Error_Type", RouteReport_Error_Type_name, RouteReport_Error_Type_value)
	proto.RegisterEnum("gloo.solo.io.TcpListenerReport_Error_Type", TcpListenerReport_Error_Type_name, TcpListenerReport_Error_Type_value)
	proto.RegisterEnum("gloo.solo.io.TcpHostReport_Error_Type", TcpHostReport_Error_Type_name, TcpHostReport_Error_Type_value)
	proto.RegisterType((*ProxyReport)(nil), "gloo.solo.io.ProxyReport")
	proto.RegisterType((*ListenerReport)(nil), "gloo.solo.io.ListenerReport")
	proto.RegisterType((*ListenerReport_Error)(nil), "gloo.solo.io.ListenerReport.Error")
	proto.RegisterType((*HttpListenerReport)(nil), "gloo.solo.io.HttpListenerReport")
	proto.RegisterType((*HttpListenerReport_Error)(nil), "gloo.solo.io.HttpListenerReport.Error")
	proto.RegisterType((*VirtualHostReport)(nil), "gloo.solo.io.VirtualHostReport")
	proto.RegisterType((*VirtualHostReport_Error)(nil), "gloo.solo.io.VirtualHostReport.Error")
	proto.RegisterType((*RouteReport)(nil), "gloo.solo.io.RouteReport")
	proto.RegisterType((*RouteReport_Error)(nil), "gloo.solo.io.RouteReport.Error")
	proto.RegisterType((*TcpListenerReport)(nil), "gloo.solo.io.TcpListenerReport")
	proto.RegisterType((*TcpListenerReport_Error)(nil), "gloo.solo.io.TcpListenerReport.Error")
	proto.RegisterType((*TcpHostReport)(nil), "gloo.solo.io.TcpHostReport")
	proto.RegisterType((*TcpHostReport_Error)(nil), "gloo.solo.io.TcpHostReport.Error")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/grpc/validation/proxy_validation.proto", fileDescriptor_8f4537dae4069b18)
}

var fileDescriptor_8f4537dae4069b18 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xee, 0x6e, 0x91, 0xc4, 0x53, 0x28, 0x30, 0xc5, 0x52, 0x16, 0x13, 0xea, 0x46, 0xb1, 0xd1,
	0xb8, 0x8d, 0xf5, 0x42, 0xc4, 0x00, 0x11, 0x21, 0x62, 0x82, 0xa4, 0x2c, 0x95, 0x0b, 0xbd, 0x68,
	0x96, 0x65, 0x6c, 0x47, 0xcb, 0xce, 0x3a, 0x33, 0x6d, 0xe4, 0xc2, 0xd7, 0xf1, 0x4a, 0xe3, 0x33,
	0x98, 0xf0, 0x02, 0xbe, 0x81, 0x2f, 0xe1, 0x8d, 0x57, 0x66, 0x7f, 0x58, 0xba, 0xb3, 0xa5, 0xbb,
	0xde, 0x79, 0x79, 0xce, 0x9c, 0xf9, 0x7a, 0xbe, 0xef, 0x7c, 0x3d, 0x3b, 0xd0, 0xec, 0x10, 0xd1,
	0xed, 0x1f, 0x1b, 0x36, 0x3d, 0xad, 0x73, 0xda, 0xa3, 0x0f, 0x08, 0xad, 0x77, 0x7a, 0x94, 0xd6,
	0x5d, 0x46, 0xdf, 0x63, 0x5b, 0xf0, 0x20, 0xb2, 0x5c, 0x52, 0xef, 0x30, 0xd7, 0xae, 0x0f, 0xac,
	0x1e, 0x39, 0xb1, 0x04, 0xa1, 0x8e, 0x57, 0xf1, 0xe9, 0xac, 0x7d, 0x99, 0x30, 0x5c, 0x46, 0x05,
	0x45, 0x53, 0xde, 0x05, 0xc3, 0xc3, 0x32, 0x08, 0xd5, 0x56, 0xb3, 0xe3, 0x0f, 0x1e, 0x06, 0x90,
	0x01, 0x8e, 0x7e, 0x04, 0x85, 0xa6, 0x17, 0x9a, 0xd8, 0xa5, 0x4c, 0xa0, 0x17, 0x30, 0xdb, 0x23,
	0x5c, 0x60, 0x07, 0xb3, 0x36, 0xf3, 0x53, 0xbc, 0xa2, 0x54, 0xf3, 0xb5, 0x42, 0xe3, 0xa6, 0x31,
	0xfc, 0x8b, 0xc6, 0x5e, 0x58, 0x15, 0xdc, 0x33, 0x67, 0x7a, 0xb1, 0x98, 0xeb, 0x3f, 0xf3, 0x50,
	0x8c, 0xd7, 0xa0, 0x35, 0x98, 0xc4, 0x8c, 0x51, 0xc6, 0x2b, 0xaa, 0x8f, 0xa8, 0x8f, 0x43, 0x34,
	0x76, 0xbc, 0x52, 0x33, 0xbc, 0x81, 0x5a, 0x30, 0xdf, 0x15, 0xc2, 0x6d, 0x4b, 0xcd, 0x55, 0xf2,
	0x55, 0xa5, 0x56, 0x68, 0x54, 0xe3, 0x48, 0xbb, 0x42, 0xb8, 0x71, 0xb4, 0xdd, 0x9c, 0x89, 0xba,
	0x89, 0x2c, 0x3a, 0x80, 0x92, 0xb0, 0x93, 0xa0, 0x13, 0x3e, 0xe8, 0x72, 0x1c, 0xb4, 0x65, 0x27,
	0x31, 0xe7, 0x84, 0x9c, 0xd4, 0x7e, 0x28, 0x70, 0xcd, 0x6f, 0x1d, 0x3d, 0x85, 0x09, 0x71, 0xe6,
	0xe2, 0x8a, 0x52, 0x55, 0x6a, 0xc5, 0xc6, 0xdd, 0x74, 0xb2, 0x46, 0xeb, 0xcc, 0xc5, 0xa6, 0x7f,
	0x09, 0x95, 0x61, 0x92, 0x61, 0x8b, 0x53, 0xa7, 0xa2, 0x56, 0x95, 0xda, 0x75, 0x33, 0x8c, 0x74,
	0x1b, 0x26, 0x5a, 0xc1, 0x39, 0xda, 0xb7, 0x4e, 0xf1, 0x3e, 0x15, 0xaf, 0x1d, 0xf2, 0xb1, 0x8f,
	0x7d, 0x80, 0xd9, 0x1c, 0xd2, 0xa0, 0xbc, 0x45, 0x9c, 0x93, 0x26, 0x65, 0x42, 0x3a, 0x53, 0x10,
	0x82, 0xe2, 0xe1, 0xe1, 0xde, 0x73, 0xea, 0xbc, 0x23, 0x9d, 0x20, 0xa7, 0xa2, 0x12, 0xcc, 0x34,
	0x19, 0xb5, 0x31, 0xe7, 0xc4, 0x09, 0x93, 0xf9, 0xad, 0x32, 0xcc, 0x47, 0x92, 0x78, 0xdd, 0x84,
	0xba, 0xe8, 0x5f, 0x54, 0x40, 0x49, 0x6d, 0xd1, 0x46, 0x34, 0xd7, 0xc0, 0x29, 0x2b, 0x69, 0xd3,
	0x90, 0x66, 0x7b, 0x00, 0xf3, 0x03, 0xc2, 0x44, 0xdf, 0xea, 0xb5, 0xbb, 0x94, 0x8b, 0xc8, 0x77,
	0x81, 0x4b, 0xa4, 0x31, 0x1c, 0x05, 0x95, 0xbb, 0x94, 0x8b, 0xd0, 0x7a, 0x68, 0x20, 0xa7, 0xb8,
	0xf6, 0xf9, 0x62, 0x08, 0x9b, 0xb1, 0x21, 0xdc, 0xcf, 0xd6, 0x59, 0x96, 0x41, 0x2c, 0x85, 0x83,
	0x18, 0x21, 0x60, 0x4e, 0x3f, 0x57, 0x61, 0x2e, 0xd1, 0x28, 0x5a, 0x97, 0x74, 0xba, 0x93, 0xc2,
	0x4c, 0x92, 0x69, 0x03, 0xa6, 0x19, 0xed, 0x0b, 0x2c, 0xe9, 0xb3, 0x18, 0x47, 0x31, 0xbd, 0x92,
	0x50, 0x99, 0x29, 0x76, 0x19, 0x70, 0xed, 0x7b, 0xe4, 0xcc, 0x8d, 0x98, 0x28, 0xf7, 0x32, 0xb5,
	0x91, 0x45, 0x93, 0xfd, 0x14, 0x73, 0x2e, 0xc2, 0x8d, 0x6d, 0x7a, 0x6a, 0x11, 0x87, 0x27, 0xbc,
	0x39, 0x42, 0x46, 0x55, 0xff, 0xad, 0x40, 0x61, 0x88, 0x0f, 0x7a, 0x2c, 0x09, 0xb8, 0x7c, 0x25,
	0xf5, 0xb8, 0x74, 0xda, 0xd7, 0x88, 0xfa, 0x6a, 0x8c, 0xfa, 0xed, 0x14, 0x80, 0x2c, 0xa4, 0x0f,
	0x42, 0xd2, 0x0b, 0x50, 0x7a, 0xe9, 0xf8, 0x6b, 0xfa, 0x95, 0x25, 0xec, 0x2e, 0x66, 0x17, 0xac,
	0x97, 0x60, 0x21, 0x3c, 0xd8, 0xc6, 0x5c, 0x10, 0xc7, 0xdf, 0xe2, 0x63, 0x79, 0xff, 0x52, 0x61,
	0x2e, 0xb1, 0x6e, 0xd2, 0xec, 0x93, 0xb8, 0x20, 0xd9, 0x67, 0x07, 0x66, 0xbd, 0x5d, 0x37, 0xe2,
	0x1f, 0xb6, 0x94, 0x00, 0x1a, 0xfa, 0x77, 0x15, 0xc5, 0x70, 0xc8, 0xb5, 0xf3, 0x6c, 0x2e, 0xba,
	0xa2, 0x9b, 0xff, 0x65, 0xc5, 0xe9, 0x7f, 0x14, 0x98, 0x8e, 0x11, 0x45, 0x4f, 0xa4, 0xaf, 0xd3,
	0xad, 0x31, 0xaa, 0x48, 0xf6, 0xfa, 0x16, 0x69, 0xb2, 0x16, 0xd3, 0x64, 0x25, 0x15, 0x22, 0x8b,
	0x1e, 0xcd, 0x14, 0x3d, 0xfe, 0xd9, 0x5f, 0x8d, 0xb7, 0x50, 0xf6, 0xbf, 0xf9, 0x47, 0xd1, 0xa3,
	0xe2, 0x10, 0xb3, 0x01, 0xb1, 0x31, 0x7a, 0x06, 0xd3, 0x61, 0x12, 0xfb, 0x15, 0xa8, 0x14, 0xa7,
	0xe0, 0x27, 0xb5, 0xc5, 0x11, 0xc9, 0x80, 0x95, 0x9e, 0xdb, 0xda, 0x7c, 0xb3, 0x9e, 0xed, 0x31,
	0xe2, 0x7e, 0xe8, 0x8c, 0x7a, 0xf0, 0x1c, 0x4f, 0xfa, 0x0f, 0x93, 0x47, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x1b, 0x3e, 0x96, 0x34, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyValidationServiceClient is the client API for ProxyValidationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyValidationServiceClient interface {
	ValidateProxy(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (*ProxyReport, error)
}

type proxyValidationServiceClient struct {
	cc *grpc.ClientConn
}

func NewProxyValidationServiceClient(cc *grpc.ClientConn) ProxyValidationServiceClient {
	return &proxyValidationServiceClient{cc}
}

func (c *proxyValidationServiceClient) ValidateProxy(ctx context.Context, in *v1.Proxy, opts ...grpc.CallOption) (*ProxyReport, error) {
	out := new(ProxyReport)
	err := c.cc.Invoke(ctx, "/gloo.solo.io.ProxyValidationService/ValidateProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyValidationServiceServer is the server API for ProxyValidationService service.
type ProxyValidationServiceServer interface {
	ValidateProxy(context.Context, *v1.Proxy) (*ProxyReport, error)
}

// UnimplementedProxyValidationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProxyValidationServiceServer struct {
}

func (*UnimplementedProxyValidationServiceServer) ValidateProxy(ctx context.Context, req *v1.Proxy) (*ProxyReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateProxy not implemented")
}

func RegisterProxyValidationServiceServer(s *grpc.Server, srv ProxyValidationServiceServer) {
	s.RegisterService(&_ProxyValidationService_serviceDesc, srv)
}

func _ProxyValidationService_ValidateProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Proxy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyValidationServiceServer).ValidateProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gloo.solo.io.ProxyValidationService/ValidateProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyValidationServiceServer).ValidateProxy(ctx, req.(*v1.Proxy))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxyValidationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gloo.solo.io.ProxyValidationService",
	HandlerType: (*ProxyValidationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateProxy",
			Handler:    _ProxyValidationService_ValidateProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/gloo/projects/gloo/api/grpc/validation/proxy_validation.proto",
}