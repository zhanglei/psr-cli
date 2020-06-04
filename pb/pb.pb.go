// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BaseCommand_Type int32

const (
	BaseCommand_CONNECT                       BaseCommand_Type = 2
	BaseCommand_CONNECTED                     BaseCommand_Type = 3
	BaseCommand_PRODUCER                      BaseCommand_Type = 5
	BaseCommand_PRODUCER_SUCCESS              BaseCommand_Type = 17
	BaseCommand_PARTITIONED_METADATA          BaseCommand_Type = 21
	BaseCommand_PARTITIONED_METADATA_RESPONSE BaseCommand_Type = 22
	BaseCommand_Lookup                        BaseCommand_Type = 23
	BaseCommand_LookupResponse                BaseCommand_Type = 24
)

var BaseCommand_Type_name = map[int32]string{
	2:  "CONNECT",
	3:  "CONNECTED",
	5:  "PRODUCER",
	17: "PRODUCER_SUCCESS",
	21: "PARTITIONED_METADATA",
	22: "PARTITIONED_METADATA_RESPONSE",
	23: "Lookup",
	24: "LookupResponse",
}

var BaseCommand_Type_value = map[string]int32{
	"CONNECT":                       2,
	"CONNECTED":                     3,
	"PRODUCER":                      5,
	"PRODUCER_SUCCESS":              17,
	"PARTITIONED_METADATA":          21,
	"PARTITIONED_METADATA_RESPONSE": 22,
	"Lookup":                        23,
	"LookupResponse":                24,
}

func (x BaseCommand_Type) Enum() *BaseCommand_Type {
	p := new(BaseCommand_Type)
	*p = x
	return p
}

func (x BaseCommand_Type) String() string {
	return proto.EnumName(BaseCommand_Type_name, int32(x))
}

func (x *BaseCommand_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BaseCommand_Type_value, data, "BaseCommand_Type")
	if err != nil {
		return err
	}
	*x = BaseCommand_Type(value)
	return nil
}

func (BaseCommand_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0, 0}
}

type CommandPartitionedTopicMetadataResponse_LookupType int32

const (
	CommandPartitionedTopicMetadataResponse_Success CommandPartitionedTopicMetadataResponse_LookupType = 0
	CommandPartitionedTopicMetadataResponse_Failed  CommandPartitionedTopicMetadataResponse_LookupType = 1
)

var CommandPartitionedTopicMetadataResponse_LookupType_name = map[int32]string{
	0: "Success",
	1: "Failed",
}

var CommandPartitionedTopicMetadataResponse_LookupType_value = map[string]int32{
	"Success": 0,
	"Failed":  1,
}

func (x CommandPartitionedTopicMetadataResponse_LookupType) Enum() *CommandPartitionedTopicMetadataResponse_LookupType {
	p := new(CommandPartitionedTopicMetadataResponse_LookupType)
	*p = x
	return p
}

func (x CommandPartitionedTopicMetadataResponse_LookupType) String() string {
	return proto.EnumName(CommandPartitionedTopicMetadataResponse_LookupType_name, int32(x))
}

func (x *CommandPartitionedTopicMetadataResponse_LookupType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandPartitionedTopicMetadataResponse_LookupType_value, data, "CommandPartitionedTopicMetadataResponse_LookupType")
	if err != nil {
		return err
	}
	*x = CommandPartitionedTopicMetadataResponse_LookupType(value)
	return nil
}

func (CommandPartitionedTopicMetadataResponse_LookupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{4, 0}
}

type CommandLookupTopicResponse_LookupType int32

const (
	CommandLookupTopicResponse_Redirect CommandLookupTopicResponse_LookupType = 0
	CommandLookupTopicResponse_Connect  CommandLookupTopicResponse_LookupType = 1
	CommandLookupTopicResponse_Failed   CommandLookupTopicResponse_LookupType = 2
)

var CommandLookupTopicResponse_LookupType_name = map[int32]string{
	0: "Redirect",
	1: "Connect",
	2: "Failed",
}

var CommandLookupTopicResponse_LookupType_value = map[string]int32{
	"Redirect": 0,
	"Connect":  1,
	"Failed":   2,
}

func (x CommandLookupTopicResponse_LookupType) Enum() *CommandLookupTopicResponse_LookupType {
	p := new(CommandLookupTopicResponse_LookupType)
	*p = x
	return p
}

func (x CommandLookupTopicResponse_LookupType) String() string {
	return proto.EnumName(CommandLookupTopicResponse_LookupType_name, int32(x))
}

func (x *CommandLookupTopicResponse_LookupType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandLookupTopicResponse_LookupType_value, data, "CommandLookupTopicResponse_LookupType")
	if err != nil {
		return err
	}
	*x = CommandLookupTopicResponse_LookupType(value)
	return nil
}

func (CommandLookupTopicResponse_LookupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{6, 0}
}

//
// package protocol
type BaseCommand struct {
	Type                      *BaseCommand_Type                        `protobuf:"varint,1,req,name=type,enum=BaseCommand_Type" json:"type,omitempty"`
	Connect                   *CommandConnect                          `protobuf:"bytes,2,opt,name=connect" json:"connect,omitempty"`
	Connected                 *CommandConnected                        `protobuf:"bytes,3,opt,name=connected" json:"connected,omitempty"`
	Producer                  *CommandProducer                         `protobuf:"bytes,5,opt,name=producer" json:"producer,omitempty"`
	ProducerSuccess           *CommandProducerSuccess                  `protobuf:"bytes,17,opt,name=producer_success,json=producerSuccess" json:"producer_success,omitempty"`
	PartitionMetadata         *CommandPartitionedTopicMetadata         `protobuf:"bytes,21,opt,name=partitionMetadata" json:"partitionMetadata,omitempty"`
	PartitionMetadataResponse *CommandPartitionedTopicMetadataResponse `protobuf:"bytes,22,opt,name=partitionMetadataResponse" json:"partitionMetadataResponse,omitempty"`
	LookupTopic               *CommandLookupTopic                      `protobuf:"bytes,23,opt,name=lookupTopic" json:"lookupTopic,omitempty"`
	LookupTopicResponse       *CommandLookupTopicResponse              `protobuf:"bytes,24,opt,name=lookupTopicResponse" json:"lookupTopicResponse,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                 `json:"-"`
	XXX_unrecognized          []byte                                   `json:"-"`
	XXX_sizecache             int32                                    `json:"-"`
}

func (m *BaseCommand) Reset()         { *m = BaseCommand{} }
func (m *BaseCommand) String() string { return proto.CompactTextString(m) }
func (*BaseCommand) ProtoMessage()    {}
func (*BaseCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}

func (m *BaseCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseCommand.Unmarshal(m, b)
}
func (m *BaseCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseCommand.Marshal(b, m, deterministic)
}
func (m *BaseCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseCommand.Merge(m, src)
}
func (m *BaseCommand) XXX_Size() int {
	return xxx_messageInfo_BaseCommand.Size(m)
}
func (m *BaseCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseCommand.DiscardUnknown(m)
}

var xxx_messageInfo_BaseCommand proto.InternalMessageInfo

func (m *BaseCommand) GetType() BaseCommand_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return BaseCommand_CONNECT
}

func (m *BaseCommand) GetConnect() *CommandConnect {
	if m != nil {
		return m.Connect
	}
	return nil
}

func (m *BaseCommand) GetConnected() *CommandConnected {
	if m != nil {
		return m.Connected
	}
	return nil
}

func (m *BaseCommand) GetProducer() *CommandProducer {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *BaseCommand) GetProducerSuccess() *CommandProducerSuccess {
	if m != nil {
		return m.ProducerSuccess
	}
	return nil
}

func (m *BaseCommand) GetPartitionMetadata() *CommandPartitionedTopicMetadata {
	if m != nil {
		return m.PartitionMetadata
	}
	return nil
}

func (m *BaseCommand) GetPartitionMetadataResponse() *CommandPartitionedTopicMetadataResponse {
	if m != nil {
		return m.PartitionMetadataResponse
	}
	return nil
}

func (m *BaseCommand) GetLookupTopic() *CommandLookupTopic {
	if m != nil {
		return m.LookupTopic
	}
	return nil
}

func (m *BaseCommand) GetLookupTopicResponse() *CommandLookupTopicResponse {
	if m != nil {
		return m.LookupTopicResponse
	}
	return nil
}

// handshake logic connection
type CommandConnect struct {
	ClientVersion        *string  `protobuf:"bytes,1,req,name=client_version,json=clientVersion" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandConnect) Reset()         { *m = CommandConnect{} }
func (m *CommandConnect) String() string { return proto.CompactTextString(m) }
func (*CommandConnect) ProtoMessage()    {}
func (*CommandConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}

func (m *CommandConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandConnect.Unmarshal(m, b)
}
func (m *CommandConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandConnect.Marshal(b, m, deterministic)
}
func (m *CommandConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandConnect.Merge(m, src)
}
func (m *CommandConnect) XXX_Size() int {
	return xxx_messageInfo_CommandConnect.Size(m)
}
func (m *CommandConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandConnect.DiscardUnknown(m)
}

var xxx_messageInfo_CommandConnect proto.InternalMessageInfo

func (m *CommandConnect) GetClientVersion() string {
	if m != nil && m.ClientVersion != nil {
		return *m.ClientVersion
	}
	return ""
}

type CommandConnected struct {
	ServerVersion        *string  `protobuf:"bytes,1,req,name=server_version,json=serverVersion" json:"server_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandConnected) Reset()         { *m = CommandConnected{} }
func (m *CommandConnected) String() string { return proto.CompactTextString(m) }
func (*CommandConnected) ProtoMessage()    {}
func (*CommandConnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{2}
}

func (m *CommandConnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandConnected.Unmarshal(m, b)
}
func (m *CommandConnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandConnected.Marshal(b, m, deterministic)
}
func (m *CommandConnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandConnected.Merge(m, src)
}
func (m *CommandConnected) XXX_Size() int {
	return xxx_messageInfo_CommandConnected.Size(m)
}
func (m *CommandConnected) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandConnected.DiscardUnknown(m)
}

var xxx_messageInfo_CommandConnected proto.InternalMessageInfo

func (m *CommandConnected) GetServerVersion() string {
	if m != nil && m.ServerVersion != nil {
		return *m.ServerVersion
	}
	return ""
}

// partitioned topic metadata
type CommandPartitionedTopicMetadata struct {
	Topic                *string  `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	RequestId            *uint64  `protobuf:"varint,2,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandPartitionedTopicMetadata) Reset()         { *m = CommandPartitionedTopicMetadata{} }
func (m *CommandPartitionedTopicMetadata) String() string { return proto.CompactTextString(m) }
func (*CommandPartitionedTopicMetadata) ProtoMessage()    {}
func (*CommandPartitionedTopicMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{3}
}

func (m *CommandPartitionedTopicMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandPartitionedTopicMetadata.Unmarshal(m, b)
}
func (m *CommandPartitionedTopicMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandPartitionedTopicMetadata.Marshal(b, m, deterministic)
}
func (m *CommandPartitionedTopicMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandPartitionedTopicMetadata.Merge(m, src)
}
func (m *CommandPartitionedTopicMetadata) XXX_Size() int {
	return xxx_messageInfo_CommandPartitionedTopicMetadata.Size(m)
}
func (m *CommandPartitionedTopicMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandPartitionedTopicMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CommandPartitionedTopicMetadata proto.InternalMessageInfo

func (m *CommandPartitionedTopicMetadata) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *CommandPartitionedTopicMetadata) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

type CommandPartitionedTopicMetadataResponse struct {
	Partitions           *uint32                                             `protobuf:"varint,1,opt,name=partitions" json:"partitions,omitempty"`
	RequestId            *int64                                              `protobuf:"varint,2,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	Response             *CommandPartitionedTopicMetadataResponse_LookupType `protobuf:"varint,3,opt,name=response,enum=CommandPartitionedTopicMetadataResponse_LookupType" json:"response,omitempty"`
	Message              *string                                             `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *CommandPartitionedTopicMetadataResponse) Reset() {
	*m = CommandPartitionedTopicMetadataResponse{}
}
func (m *CommandPartitionedTopicMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*CommandPartitionedTopicMetadataResponse) ProtoMessage()    {}
func (*CommandPartitionedTopicMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{4}
}

func (m *CommandPartitionedTopicMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandPartitionedTopicMetadataResponse.Unmarshal(m, b)
}
func (m *CommandPartitionedTopicMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandPartitionedTopicMetadataResponse.Marshal(b, m, deterministic)
}
func (m *CommandPartitionedTopicMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandPartitionedTopicMetadataResponse.Merge(m, src)
}
func (m *CommandPartitionedTopicMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_CommandPartitionedTopicMetadataResponse.Size(m)
}
func (m *CommandPartitionedTopicMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandPartitionedTopicMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandPartitionedTopicMetadataResponse proto.InternalMessageInfo

func (m *CommandPartitionedTopicMetadataResponse) GetPartitions() uint32 {
	if m != nil && m.Partitions != nil {
		return *m.Partitions
	}
	return 0
}

func (m *CommandPartitionedTopicMetadataResponse) GetRequestId() int64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *CommandPartitionedTopicMetadataResponse) GetResponse() CommandPartitionedTopicMetadataResponse_LookupType {
	if m != nil && m.Response != nil {
		return *m.Response
	}
	return CommandPartitionedTopicMetadataResponse_Success
}

func (m *CommandPartitionedTopicMetadataResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// lookup topic broker
type CommandLookupTopic struct {
	Topic                *string  `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	RequestId            *uint64  `protobuf:"varint,2,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandLookupTopic) Reset()         { *m = CommandLookupTopic{} }
func (m *CommandLookupTopic) String() string { return proto.CompactTextString(m) }
func (*CommandLookupTopic) ProtoMessage()    {}
func (*CommandLookupTopic) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{5}
}

func (m *CommandLookupTopic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLookupTopic.Unmarshal(m, b)
}
func (m *CommandLookupTopic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLookupTopic.Marshal(b, m, deterministic)
}
func (m *CommandLookupTopic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLookupTopic.Merge(m, src)
}
func (m *CommandLookupTopic) XXX_Size() int {
	return xxx_messageInfo_CommandLookupTopic.Size(m)
}
func (m *CommandLookupTopic) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLookupTopic.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLookupTopic proto.InternalMessageInfo

func (m *CommandLookupTopic) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *CommandLookupTopic) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

type CommandLookupTopicResponse struct {
	BrokerServiceUrl     *string                                `protobuf:"bytes,1,opt,name=brokerServiceUrl" json:"brokerServiceUrl,omitempty"`
	BrokerServiceUrlTls  *string                                `protobuf:"bytes,2,opt,name=brokerServiceUrlTls" json:"brokerServiceUrlTls,omitempty"`
	Response             *CommandLookupTopicResponse_LookupType `protobuf:"varint,3,opt,name=response,enum=CommandLookupTopicResponse_LookupType" json:"response,omitempty"`
	RequestId            *uint64                                `protobuf:"varint,4,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	Message              *string                                `protobuf:"bytes,7,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CommandLookupTopicResponse) Reset()         { *m = CommandLookupTopicResponse{} }
func (m *CommandLookupTopicResponse) String() string { return proto.CompactTextString(m) }
func (*CommandLookupTopicResponse) ProtoMessage()    {}
func (*CommandLookupTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{6}
}

func (m *CommandLookupTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLookupTopicResponse.Unmarshal(m, b)
}
func (m *CommandLookupTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLookupTopicResponse.Marshal(b, m, deterministic)
}
func (m *CommandLookupTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLookupTopicResponse.Merge(m, src)
}
func (m *CommandLookupTopicResponse) XXX_Size() int {
	return xxx_messageInfo_CommandLookupTopicResponse.Size(m)
}
func (m *CommandLookupTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLookupTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLookupTopicResponse proto.InternalMessageInfo

func (m *CommandLookupTopicResponse) GetBrokerServiceUrl() string {
	if m != nil && m.BrokerServiceUrl != nil {
		return *m.BrokerServiceUrl
	}
	return ""
}

func (m *CommandLookupTopicResponse) GetBrokerServiceUrlTls() string {
	if m != nil && m.BrokerServiceUrlTls != nil {
		return *m.BrokerServiceUrlTls
	}
	return ""
}

func (m *CommandLookupTopicResponse) GetResponse() CommandLookupTopicResponse_LookupType {
	if m != nil && m.Response != nil {
		return *m.Response
	}
	return CommandLookupTopicResponse_Redirect
}

func (m *CommandLookupTopicResponse) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *CommandLookupTopicResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// register producer
type CommandProducer struct {
	Topic                *string  `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	ProducerId           *uint64  `protobuf:"varint,2,req,name=producer_id,json=producerId" json:"producer_id,omitempty"`
	RequestId            *uint64  `protobuf:"varint,3,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	ProducerName         *string  `protobuf:"bytes,4,opt,name=producer_name,json=producerName" json:"producer_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandProducer) Reset()         { *m = CommandProducer{} }
func (m *CommandProducer) String() string { return proto.CompactTextString(m) }
func (*CommandProducer) ProtoMessage()    {}
func (*CommandProducer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{7}
}

func (m *CommandProducer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandProducer.Unmarshal(m, b)
}
func (m *CommandProducer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandProducer.Marshal(b, m, deterministic)
}
func (m *CommandProducer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandProducer.Merge(m, src)
}
func (m *CommandProducer) XXX_Size() int {
	return xxx_messageInfo_CommandProducer.Size(m)
}
func (m *CommandProducer) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandProducer.DiscardUnknown(m)
}

var xxx_messageInfo_CommandProducer proto.InternalMessageInfo

func (m *CommandProducer) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *CommandProducer) GetProducerId() uint64 {
	if m != nil && m.ProducerId != nil {
		return *m.ProducerId
	}
	return 0
}

func (m *CommandProducer) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *CommandProducer) GetProducerName() string {
	if m != nil && m.ProducerName != nil {
		return *m.ProducerName
	}
	return ""
}

type CommandProducerSuccess struct {
	RequestId            *uint64  `protobuf:"varint,1,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	ProducerName         *string  `protobuf:"bytes,2,req,name=producer_name,json=producerName" json:"producer_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandProducerSuccess) Reset()         { *m = CommandProducerSuccess{} }
func (m *CommandProducerSuccess) String() string { return proto.CompactTextString(m) }
func (*CommandProducerSuccess) ProtoMessage()    {}
func (*CommandProducerSuccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{8}
}

func (m *CommandProducerSuccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandProducerSuccess.Unmarshal(m, b)
}
func (m *CommandProducerSuccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandProducerSuccess.Marshal(b, m, deterministic)
}
func (m *CommandProducerSuccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandProducerSuccess.Merge(m, src)
}
func (m *CommandProducerSuccess) XXX_Size() int {
	return xxx_messageInfo_CommandProducerSuccess.Size(m)
}
func (m *CommandProducerSuccess) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandProducerSuccess.DiscardUnknown(m)
}

var xxx_messageInfo_CommandProducerSuccess proto.InternalMessageInfo

func (m *CommandProducerSuccess) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *CommandProducerSuccess) GetProducerName() string {
	if m != nil && m.ProducerName != nil {
		return *m.ProducerName
	}
	return ""
}

func init() {
	proto.RegisterEnum("BaseCommand_Type", BaseCommand_Type_name, BaseCommand_Type_value)
	proto.RegisterEnum("CommandPartitionedTopicMetadataResponse_LookupType", CommandPartitionedTopicMetadataResponse_LookupType_name, CommandPartitionedTopicMetadataResponse_LookupType_value)
	proto.RegisterEnum("CommandLookupTopicResponse_LookupType", CommandLookupTopicResponse_LookupType_name, CommandLookupTopicResponse_LookupType_value)
	proto.RegisterType((*BaseCommand)(nil), "BaseCommand")
	proto.RegisterType((*CommandConnect)(nil), "CommandConnect")
	proto.RegisterType((*CommandConnected)(nil), "CommandConnected")
	proto.RegisterType((*CommandPartitionedTopicMetadata)(nil), "CommandPartitionedTopicMetadata")
	proto.RegisterType((*CommandPartitionedTopicMetadataResponse)(nil), "CommandPartitionedTopicMetadataResponse")
	proto.RegisterType((*CommandLookupTopic)(nil), "CommandLookupTopic")
	proto.RegisterType((*CommandLookupTopicResponse)(nil), "CommandLookupTopicResponse")
	proto.RegisterType((*CommandProducer)(nil), "CommandProducer")
	proto.RegisterType((*CommandProducerSuccess)(nil), "CommandProducerSuccess")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0x8d, 0x0d, 0x7c, 0xc0, 0x25, 0x90, 0x61, 0xf2, 0x37, 0x5f, 0xab, 0x34, 0xd4, 0x55, 0x5a,
	0x5a, 0x55, 0xb4, 0x4a, 0x54, 0x55, 0x5d, 0x12, 0x70, 0x25, 0xa4, 0x06, 0xd0, 0xe0, 0x64, 0x55,
	0x09, 0x39, 0xf6, 0x6d, 0x65, 0x05, 0x6c, 0xd7, 0x36, 0x91, 0xf2, 0x0e, 0x7d, 0x8b, 0x6e, 0xfb,
	0x5c, 0x7d, 0x88, 0xae, 0x2a, 0xc6, 0x1e, 0x07, 0x0c, 0x69, 0xa2, 0xee, 0x66, 0xce, 0x3d, 0xe7,
	0xf8, 0xfa, 0xce, 0x99, 0x81, 0x92, 0x7f, 0xd9, 0xf2, 0x03, 0x2f, 0xf2, 0xb4, 0x5f, 0x05, 0xa8,
	0x9c, 0x9a, 0x21, 0x76, 0xbc, 0xe9, 0xd4, 0x74, 0x6d, 0x7a, 0x04, 0xf9, 0xe8, 0xc6, 0x47, 0xa6,
	0x34, 0xd4, 0x66, 0xed, 0xb8, 0xde, 0x5a, 0xa8, 0xb5, 0x8c, 0x1b, 0x1f, 0xb9, 0x28, 0xd3, 0x97,
	0x50, 0xb4, 0x3c, 0xd7, 0x45, 0x2b, 0x62, 0x6a, 0x43, 0x69, 0x56, 0x8e, 0xb7, 0x5a, 0x09, 0xab,
	0x13, 0xc3, 0x5c, 0xd6, 0xe9, 0x1b, 0x28, 0x27, 0x4b, 0xb4, 0x59, 0x4e, 0x90, 0xeb, 0x19, 0x32,
	0xda, 0xfc, 0x96, 0x43, 0x5f, 0x43, 0xc9, 0x0f, 0x3c, 0x7b, 0x66, 0x61, 0xc0, 0x0a, 0x82, 0x4f,
	0x24, 0x7f, 0x98, 0xe0, 0x3c, 0x65, 0xd0, 0x53, 0x20, 0x72, 0x3d, 0x0e, 0x67, 0x96, 0x85, 0x61,
	0xc8, 0xea, 0x42, 0xb5, 0x9f, 0x55, 0x8d, 0xe2, 0x32, 0xdf, 0xf2, 0x97, 0x01, 0xda, 0x87, 0xba,
	0x6f, 0x06, 0x91, 0x13, 0x39, 0x9e, 0x7b, 0x86, 0x91, 0x69, 0x9b, 0x91, 0xc9, 0x76, 0x85, 0x49,
	0x23, 0x35, 0x91, 0x04, 0xb4, 0x0d, 0xcf, 0x77, 0x2c, 0xc9, 0xe3, 0xab, 0x52, 0xfa, 0x05, 0xfe,
	0x5f, 0x01, 0x39, 0x86, 0xbe, 0xe7, 0x86, 0xc8, 0xf6, 0x84, 0x6f, 0xf3, 0x5e, 0xdf, 0x84, 0xcf,
	0xef, 0xb6, 0xa2, 0xef, 0xa0, 0x32, 0xf1, 0xbc, 0xab, 0x99, 0x2f, 0x94, 0x6c, 0x5f, 0x38, 0x6f,
	0x4b, 0xe7, 0x4f, 0xb7, 0x25, 0xbe, 0xc8, 0xa3, 0x67, 0xb0, 0xbd, 0xb0, 0x4d, 0x1b, 0x63, 0x42,
	0xfe, 0x78, 0x9d, 0x5c, 0xf6, 0xb2, 0x4e, 0xa7, 0xfd, 0x50, 0x20, 0x3f, 0x8f, 0x06, 0xad, 0x40,
	0xb1, 0x33, 0xe8, 0xf7, 0xf5, 0x8e, 0x41, 0x54, 0x5a, 0x85, 0x72, 0xb2, 0xd1, 0xbb, 0x24, 0x47,
	0x37, 0xa1, 0x34, 0xe4, 0x83, 0xee, 0x79, 0x47, 0xe7, 0xa4, 0x40, 0x77, 0x80, 0xc8, 0xdd, 0x78,
	0x74, 0xde, 0xe9, 0xe8, 0xa3, 0x11, 0xa9, 0x53, 0x06, 0x3b, 0xc3, 0x36, 0x37, 0x7a, 0x46, 0x6f,
	0xd0, 0xd7, 0xbb, 0xe3, 0x33, 0xdd, 0x68, 0x77, 0xdb, 0x46, 0x9b, 0xec, 0xd2, 0xa7, 0x70, 0xb0,
	0xae, 0x32, 0xe6, 0xfa, 0x68, 0x38, 0xe8, 0x8f, 0x74, 0xb2, 0x47, 0x01, 0xfe, 0x8b, 0x3b, 0x26,
	0xfb, 0x94, 0x42, 0x2d, 0x5e, 0xcb, 0x1e, 0x09, 0xd3, 0xde, 0x43, 0x6d, 0x39, 0x74, 0xf4, 0x08,
	0x6a, 0xd6, 0xc4, 0x41, 0x37, 0x1a, 0x5f, 0x63, 0x10, 0x3a, 0x9e, 0x2b, 0x42, 0x5f, 0xe6, 0xd5,
	0x18, 0xbd, 0x88, 0x41, 0xed, 0x03, 0x90, 0x6c, 0x5a, 0xe7, 0xd2, 0x10, 0x83, 0x6b, 0x0c, 0xb2,
	0xd2, 0x18, 0x95, 0xd2, 0x0b, 0x38, 0xbc, 0xe7, 0x94, 0xe9, 0x0e, 0x14, 0x22, 0x71, 0x78, 0xb1,
	0x41, 0xbc, 0xa1, 0x07, 0x00, 0x01, 0x7e, 0x9b, 0x61, 0x18, 0x8d, 0x1d, 0x9b, 0xa9, 0x0d, 0xb5,
	0x99, 0xe7, 0xe5, 0x04, 0xe9, 0xd9, 0xda, 0x6f, 0x05, 0x5e, 0x3c, 0x30, 0x3e, 0xf4, 0x09, 0x40,
	0x1a, 0xa0, 0x90, 0x29, 0x0d, 0xa5, 0x59, 0xe5, 0x0b, 0xc8, 0x9a, 0x4f, 0xe5, 0x16, 0x3e, 0x45,
	0x07, 0x50, 0x0a, 0x64, 0x40, 0xe6, 0x97, 0xb7, 0x76, 0x7c, 0xf2, 0xd0, 0xe4, 0xb6, 0x92, 0x04,
	0xcd, 0x5f, 0x8d, 0xd4, 0x84, 0x32, 0x28, 0x4e, 0x31, 0x0c, 0xcd, 0xaf, 0x28, 0x2e, 0x77, 0x99,
	0xcb, 0xad, 0x76, 0x04, 0x70, 0xab, 0x98, 0x87, 0x29, 0xb9, 0x9e, 0x64, 0x63, 0x7e, 0xb8, 0x1f,
	0x4d, 0x67, 0x82, 0x36, 0x51, 0xb4, 0x1e, 0xd0, 0xd5, 0x84, 0xfe, 0xdb, 0x1c, 0x7f, 0xaa, 0xf0,
	0xe8, 0xee, 0xb4, 0xd3, 0x57, 0x40, 0x2e, 0x03, 0xef, 0x0a, 0x83, 0x11, 0x06, 0xd7, 0x8e, 0x85,
	0xe7, 0xc1, 0x44, 0x0c, 0xb0, 0xcc, 0x57, 0x70, 0xfa, 0x16, 0xb6, 0xb3, 0x98, 0x31, 0x09, 0xc5,
	0xe3, 0x58, 0xe6, 0xeb, 0x4a, 0xf4, 0x74, 0x65, 0xb2, 0xcf, 0xff, 0x72, 0xf5, 0xd6, 0x0f, 0x73,
	0xf9, 0xff, 0xf2, 0x99, 0xff, 0x5b, 0x9c, 0x75, 0x71, 0x79, 0xd6, 0x27, 0x4b, 0xb3, 0xde, 0x84,
	0x12, 0x47, 0xdb, 0x09, 0xd0, 0x8a, 0xc8, 0x86, 0xb8, 0xc6, 0x71, 0xd2, 0x89, 0xb2, 0x30, 0x79,
	0x55, 0xfb, 0xae, 0xc0, 0x56, 0xe6, 0x49, 0xbd, 0x63, 0xee, 0x87, 0x50, 0x49, 0x1f, 0xe5, 0x74,
	0xf0, 0x20, 0xa1, 0x9e, 0x9d, 0x69, 0x3c, 0x97, 0x6d, 0xfc, 0x19, 0x54, 0x53, 0xbd, 0x6b, 0x4e,
	0x91, 0xe5, 0x45, 0xfb, 0x9b, 0x12, 0xec, 0x9b, 0x53, 0xd4, 0x3e, 0xc3, 0xde, 0xfa, 0x07, 0x3e,
	0xe3, 0xae, 0xdc, 0xeb, 0xae, 0x8a, 0xde, 0x97, 0xdc, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x33,
	0x12, 0xa4, 0xb5, 0x23, 0x07, 0x00, 0x00,
}
