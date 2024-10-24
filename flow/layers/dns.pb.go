// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: flow/layers/dns.proto

package layers

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DNSMX specifies the mail server responsible for accepting email messages on behalf of a domain name.
type DNSMX struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Preference uint32 `protobuf:"varint,1,opt,name=preference,proto3" json:"preference,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DNSMX) Reset() {
	*x = DNSMX{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_layers_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSMX) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSMX) ProtoMessage() {}

func (x *DNSMX) ProtoReflect() protoreflect.Message {
	mi := &file_flow_layers_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSMX.ProtoReflect.Descriptor instead.
func (*DNSMX) Descriptor() ([]byte, []int) {
	return file_flow_layers_dns_proto_rawDescGZIP(), []int{0}
}

func (x *DNSMX) GetPreference() uint32 {
	if x != nil {
		return x.Preference
	}
	return 0
}

func (x *DNSMX) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DNSQuestion stores the question for a DNS query
type DNSQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Class string `protobuf:"bytes,3,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *DNSQuestion) Reset() {
	*x = DNSQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_layers_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSQuestion) ProtoMessage() {}

func (x *DNSQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_flow_layers_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSQuestion.ProtoReflect.Descriptor instead.
func (*DNSQuestion) Descriptor() ([]byte, []int) {
	return file_flow_layers_dns_proto_rawDescGZIP(), []int{1}
}

func (x *DNSQuestion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSQuestion) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DNSQuestion) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

// DNSResourceRecord holds resource records for DNS structs
type DNSResourceRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Class      string    `protobuf:"bytes,3,opt,name=class,proto3" json:"class,omitempty"`
	TTL        uint32    `protobuf:"varint,4,opt,name=TTL,proto3" json:"TTL,omitempty"`
	DataLength uint32    `protobuf:"varint,5,opt,name=data_length,json=dataLength,proto3" json:"data_length,omitempty"`
	IP         string    `protobuf:"bytes,6,opt,name=IP,proto3" json:"IP,omitempty"`
	NS         string    `protobuf:"bytes,7,opt,name=NS,proto3" json:"NS,omitempty"`
	CName      string    `protobuf:"bytes,8,opt,name=c_name,json=CName,proto3" json:"c_name,omitempty"`
	PTR        string    `protobuf:"bytes,9,opt,name=PTR,proto3" json:"PTR,omitempty"`
	TXTs       []string  `protobuf:"bytes,10,rep,name=TXTs,proto3" json:"TXTs,omitempty"`
	SOA        *DNSSOA   `protobuf:"bytes,11,opt,name=SOA,proto3" json:"SOA,omitempty"`
	SRV        *DNSSRV   `protobuf:"bytes,12,opt,name=SRV,proto3" json:"SRV,omitempty"`
	MX         *DNSMX    `protobuf:"bytes,13,opt,name=MX,proto3" json:"MX,omitempty"`
	OPT        []*DNSOPT `protobuf:"bytes,14,rep,name=OPT,proto3" json:"OPT,omitempty"`
}

func (x *DNSResourceRecord) Reset() {
	*x = DNSResourceRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_layers_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSResourceRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSResourceRecord) ProtoMessage() {}

func (x *DNSResourceRecord) ProtoReflect() protoreflect.Message {
	mi := &file_flow_layers_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSResourceRecord.ProtoReflect.Descriptor instead.
func (*DNSResourceRecord) Descriptor() ([]byte, []int) {
	return file_flow_layers_dns_proto_rawDescGZIP(), []int{2}
}

func (x *DNSResourceRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSResourceRecord) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DNSResourceRecord) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *DNSResourceRecord) GetTTL() uint32 {
	if x != nil {
		return x.TTL
	}
	return 0
}

func (x *DNSResourceRecord) GetDataLength() uint32 {
	if x != nil {
		return x.DataLength
	}
	return 0
}

func (x *DNSResourceRecord) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *DNSResourceRecord) GetNS() string {
	if x != nil {
		return x.NS
	}
	return ""
}

func (x *DNSResourceRecord) GetCName() string {
	if x != nil {
		return x.CName
	}
	return ""
}

func (x *DNSResourceRecord) GetPTR() string {
	if x != nil {
		return x.PTR
	}
	return ""
}

func (x *DNSResourceRecord) GetTXTs() []string {
	if x != nil {
		return x.TXTs
	}
	return nil
}

func (x *DNSResourceRecord) GetSOA() *DNSSOA {
	if x != nil {
		return x.SOA
	}
	return nil
}

func (x *DNSResourceRecord) GetSRV() *DNSSRV {
	if x != nil {
		return x.SRV
	}
	return nil
}

func (x *DNSResourceRecord) GetMX() *DNSMX {
	if x != nil {
		return x.MX
	}
	return nil
}

func (x *DNSResourceRecord) GetOPT() []*DNSOPT {
	if x != nil {
		return x.OPT
	}
	return nil
}

// DNSSOA is a type of resource record in the Domain Name System (DNS) containing administrative information
// about the zone, especially regarding zone transfers.
type DNSSOA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MName   string `protobuf:"bytes,1,opt,name=m_name,json=mName,proto3" json:"m_name,omitempty"`
	RName   string `protobuf:"bytes,2,opt,name=r_name,json=rName,proto3" json:"r_name,omitempty"`
	Serial  uint32 `protobuf:"varint,3,opt,name=serial,proto3" json:"serial,omitempty"`
	Refresh uint32 `protobuf:"varint,4,opt,name=refresh,proto3" json:"refresh,omitempty"`
	Retry   uint32 `protobuf:"varint,5,opt,name=retry,proto3" json:"retry,omitempty"`
	Expire  uint32 `protobuf:"varint,6,opt,name=expire,proto3" json:"expire,omitempty"`
	Minimum uint32 `protobuf:"varint,7,opt,name=minimum,proto3" json:"minimum,omitempty"`
}

func (x *DNSSOA) Reset() {
	*x = DNSSOA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_layers_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSSOA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSSOA) ProtoMessage() {}

func (x *DNSSOA) ProtoReflect() protoreflect.Message {
	mi := &file_flow_layers_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSSOA.ProtoReflect.Descriptor instead.
func (*DNSSOA) Descriptor() ([]byte, []int) {
	return file_flow_layers_dns_proto_rawDescGZIP(), []int{3}
}

func (x *DNSSOA) GetMName() string {
	if x != nil {
		return x.MName
	}
	return ""
}

func (x *DNSSOA) GetRName() string {
	if x != nil {
		return x.RName
	}
	return ""
}

func (x *DNSSOA) GetSerial() uint32 {
	if x != nil {
		return x.Serial
	}
	return 0
}

func (x *DNSSOA) GetRefresh() uint32 {
	if x != nil {
		return x.Refresh
	}
	return 0
}

func (x *DNSSOA) GetRetry() uint32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

func (x *DNSSOA) GetExpire() uint32 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *DNSSOA) GetMinimum() uint32 {
	if x != nil {
		return x.Minimum
	}
	return 0
}

// DNSSRV is a specification of data in the Domain Name System defining the
// location, i.e. the hostname and port number, of servers for specified services.
type DNSSRV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority uint32 `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	Weight   uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Port     uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DNSSRV) Reset() {
	*x = DNSSRV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_layers_dns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSSRV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSSRV) ProtoMessage() {}

func (x *DNSSRV) ProtoReflect() protoreflect.Message {
	mi := &file_flow_layers_dns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSSRV.ProtoReflect.Descriptor instead.
func (*DNSSRV) Descriptor() ([]byte, []int) {
	return file_flow_layers_dns_proto_rawDescGZIP(), []int{4}
}

func (x *DNSSRV) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *DNSSRV) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *DNSSRV) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *DNSSRV) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DNSOPT is a DNS Option, see RFC6891, section 6.1.2
type DNSOPT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DNSOPT) Reset() {
	*x = DNSOPT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_layers_dns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSOPT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSOPT) ProtoMessage() {}

func (x *DNSOPT) ProtoReflect() protoreflect.Message {
	mi := &file_flow_layers_dns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSOPT.ProtoReflect.Descriptor instead.
func (*DNSOPT) Descriptor() ([]byte, []int) {
	return file_flow_layers_dns_proto_rawDescGZIP(), []int{5}
}

func (x *DNSOPT) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DNSOPT) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// LayerDNS wrapper to generate extra layer
type DNS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	QR           bool                 `protobuf:"varint,2,opt,name=QR,proto3" json:"QR,omitempty"`
	OpCode       string               `protobuf:"bytes,3,opt,name=op_code,json=opCode,proto3" json:"op_code,omitempty"`
	AA           bool                 `protobuf:"varint,4,opt,name=AA,proto3" json:"AA,omitempty"`
	TC           bool                 `protobuf:"varint,5,opt,name=TC,proto3" json:"TC,omitempty"`
	RD           bool                 `protobuf:"varint,6,opt,name=RD,proto3" json:"RD,omitempty"`
	RA           bool                 `protobuf:"varint,7,opt,name=RA,proto3" json:"RA,omitempty"`
	Z            uint32               `protobuf:"varint,8,opt,name=z,proto3" json:"z,omitempty"`
	ResponseCode string               `protobuf:"bytes,9,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	QDCount      uint32               `protobuf:"varint,10,opt,name=q_d_count,json=qDCount,proto3" json:"q_d_count,omitempty"`
	ANCount      uint32               `protobuf:"varint,11,opt,name=a_n_count,json=aNCount,proto3" json:"a_n_count,omitempty"`
	NSCount      uint32               `protobuf:"varint,12,opt,name=n_s_count,json=nSCount,proto3" json:"n_s_count,omitempty"`
	ARCount      uint32               `protobuf:"varint,13,opt,name=a_r_count,json=aRCount,proto3" json:"a_r_count,omitempty"`
	Questions    []*DNSQuestion       `protobuf:"bytes,14,rep,name=questions,proto3" json:"questions,omitempty"`
	Answers      []*DNSResourceRecord `protobuf:"bytes,15,rep,name=answers,proto3" json:"answers,omitempty"`
	Authorities  []*DNSResourceRecord `protobuf:"bytes,16,rep,name=authorities,proto3" json:"authorities,omitempty"`
	Additionals  []*DNSResourceRecord `protobuf:"bytes,17,rep,name=additionals,proto3" json:"additionals,omitempty"`
	Timestamp    *timestamp.Timestamp `protobuf:"bytes,18,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DNS) Reset() {
	*x = DNS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_layers_dns_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNS) ProtoMessage() {}

func (x *DNS) ProtoReflect() protoreflect.Message {
	mi := &file_flow_layers_dns_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNS.ProtoReflect.Descriptor instead.
func (*DNS) Descriptor() ([]byte, []int) {
	return file_flow_layers_dns_proto_rawDescGZIP(), []int{6}
}

func (x *DNS) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DNS) GetQR() bool {
	if x != nil {
		return x.QR
	}
	return false
}

func (x *DNS) GetOpCode() string {
	if x != nil {
		return x.OpCode
	}
	return ""
}

func (x *DNS) GetAA() bool {
	if x != nil {
		return x.AA
	}
	return false
}

func (x *DNS) GetTC() bool {
	if x != nil {
		return x.TC
	}
	return false
}

func (x *DNS) GetRD() bool {
	if x != nil {
		return x.RD
	}
	return false
}

func (x *DNS) GetRA() bool {
	if x != nil {
		return x.RA
	}
	return false
}

func (x *DNS) GetZ() uint32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *DNS) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

func (x *DNS) GetQDCount() uint32 {
	if x != nil {
		return x.QDCount
	}
	return 0
}

func (x *DNS) GetANCount() uint32 {
	if x != nil {
		return x.ANCount
	}
	return 0
}

func (x *DNS) GetNSCount() uint32 {
	if x != nil {
		return x.NSCount
	}
	return 0
}

func (x *DNS) GetARCount() uint32 {
	if x != nil {
		return x.ARCount
	}
	return 0
}

func (x *DNS) GetQuestions() []*DNSQuestion {
	if x != nil {
		return x.Questions
	}
	return nil
}

func (x *DNS) GetAnswers() []*DNSResourceRecord {
	if x != nil {
		return x.Answers
	}
	return nil
}

func (x *DNS) GetAuthorities() []*DNSResourceRecord {
	if x != nil {
		return x.Authorities
	}
	return nil
}

func (x *DNS) GetAdditionals() []*DNSResourceRecord {
	if x != nil {
		return x.Additionals
	}
	return nil
}

func (x *DNS) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_flow_layers_dns_proto protoreflect.FileDescriptor

var file_flow_layers_dns_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3b, 0x0a, 0x05, 0x44, 0x4e, 0x53, 0x4d, 0x58, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a,
	0x0b, 0x44, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0xe6, 0x02, 0x0a, 0x11, 0x44,
	0x4e, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x54, 0x54, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x54, 0x54, 0x4c,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x50, 0x12, 0x0e, 0x0a, 0x02, 0x4e, 0x53, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4e,
	0x53, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x43, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x54, 0x52, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x54, 0x52, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x58,
	0x54, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x58, 0x54, 0x73, 0x12, 0x20,
	0x0a, 0x03, 0x53, 0x4f, 0x41, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x4e, 0x53, 0x53, 0x4f, 0x41, 0x52, 0x03, 0x53, 0x4f, 0x41,
	0x12, 0x20, 0x0a, 0x03, 0x53, 0x52, 0x56, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x4e, 0x53, 0x53, 0x52, 0x56, 0x52, 0x03, 0x53,
	0x52, 0x56, 0x12, 0x1d, 0x0a, 0x02, 0x4d, 0x58, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x4e, 0x53, 0x4d, 0x58, 0x52, 0x02, 0x4d,
	0x58, 0x12, 0x20, 0x0a, 0x03, 0x4f, 0x50, 0x54, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x4e, 0x53, 0x4f, 0x50, 0x54, 0x52, 0x03,
	0x4f, 0x50, 0x54, 0x22, 0xb0, 0x01, 0x0a, 0x06, 0x44, 0x4e, 0x53, 0x53, 0x4f, 0x41, 0x12, 0x15,
	0x0a, 0x06, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x22, 0x64, 0x0a, 0x06, 0x44, 0x4e, 0x53, 0x53, 0x52, 0x56,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x06,
	0x44, 0x4e, 0x53, 0x4f, 0x50, 0x54, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbd,
	0x04, 0x0a, 0x03, 0x44, 0x4e, 0x53, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x51, 0x52, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x51, 0x52, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x41, 0x41, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x41, 0x41, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x43, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x54, 0x43, 0x12,
	0x0e, 0x0a, 0x02, 0x52, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x52, 0x44, 0x12,
	0x0e, 0x0a, 0x02, 0x52, 0x41, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x52, 0x41, 0x12,
	0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x7a, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x09, 0x71, 0x5f, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x44, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x09, 0x61, 0x5f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x61, 0x4e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x09, 0x6e, 0x5f,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e,
	0x53, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x5f, 0x72, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x52, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x44,
	0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e,
	0x44, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x79,
	0x64, 0x69, 0x76, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x6b, 0x79,
	0x64, 0x69, 0x76, 0x65, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_layers_dns_proto_rawDescOnce sync.Once
	file_flow_layers_dns_proto_rawDescData = file_flow_layers_dns_proto_rawDesc
)

func file_flow_layers_dns_proto_rawDescGZIP() []byte {
	file_flow_layers_dns_proto_rawDescOnce.Do(func() {
		file_flow_layers_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_layers_dns_proto_rawDescData)
	})
	return file_flow_layers_dns_proto_rawDescData
}

var file_flow_layers_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_flow_layers_dns_proto_goTypes = []interface{}{
	(*DNSMX)(nil),               // 0: layers.DNSMX
	(*DNSQuestion)(nil),         // 1: layers.DNSQuestion
	(*DNSResourceRecord)(nil),   // 2: layers.DNSResourceRecord
	(*DNSSOA)(nil),              // 3: layers.DNSSOA
	(*DNSSRV)(nil),              // 4: layers.DNSSRV
	(*DNSOPT)(nil),              // 5: layers.DNSOPT
	(*DNS)(nil),                 // 6: layers.DNS
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_flow_layers_dns_proto_depIdxs = []int32{
	3, // 0: layers.DNSResourceRecord.SOA:type_name -> layers.DNSSOA
	4, // 1: layers.DNSResourceRecord.SRV:type_name -> layers.DNSSRV
	0, // 2: layers.DNSResourceRecord.MX:type_name -> layers.DNSMX
	5, // 3: layers.DNSResourceRecord.OPT:type_name -> layers.DNSOPT
	1, // 4: layers.DNS.questions:type_name -> layers.DNSQuestion
	2, // 5: layers.DNS.answers:type_name -> layers.DNSResourceRecord
	2, // 6: layers.DNS.authorities:type_name -> layers.DNSResourceRecord
	2, // 7: layers.DNS.additionals:type_name -> layers.DNSResourceRecord
	7, // 8: layers.DNS.timestamp:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_flow_layers_dns_proto_init() }
func file_flow_layers_dns_proto_init() {
	if File_flow_layers_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_layers_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSMX); i {
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
		file_flow_layers_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSQuestion); i {
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
		file_flow_layers_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSResourceRecord); i {
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
		file_flow_layers_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSSOA); i {
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
		file_flow_layers_dns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSSRV); i {
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
		file_flow_layers_dns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSOPT); i {
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
		file_flow_layers_dns_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNS); i {
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
			RawDescriptor: file_flow_layers_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_layers_dns_proto_goTypes,
		DependencyIndexes: file_flow_layers_dns_proto_depIdxs,
		MessageInfos:      file_flow_layers_dns_proto_msgTypes,
	}.Build()
	File_flow_layers_dns_proto = out.File
	file_flow_layers_dns_proto_rawDesc = nil
	file_flow_layers_dns_proto_goTypes = nil
	file_flow_layers_dns_proto_depIdxs = nil
}
