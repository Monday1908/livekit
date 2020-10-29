// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rtc.proto

package livekit

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

type SignalRequest struct {
	// Types that are valid to be assigned to Message:
	//	*SignalRequest_Join
	//	*SignalRequest_Negotiate
	//	*SignalRequest_Trickle
	Message              isSignalRequest_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SignalRequest) Reset()         { *m = SignalRequest{} }
func (m *SignalRequest) String() string { return proto.CompactTextString(m) }
func (*SignalRequest) ProtoMessage()    {}
func (*SignalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bd91e80f756774, []int{0}
}

func (m *SignalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalRequest.Unmarshal(m, b)
}
func (m *SignalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalRequest.Marshal(b, m, deterministic)
}
func (m *SignalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalRequest.Merge(m, src)
}
func (m *SignalRequest) XXX_Size() int {
	return xxx_messageInfo_SignalRequest.Size(m)
}
func (m *SignalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignalRequest proto.InternalMessageInfo

type isSignalRequest_Message interface {
	isSignalRequest_Message()
}

type SignalRequest_Join struct {
	Join *JoinRequest `protobuf:"bytes,1,opt,name=join,proto3,oneof"`
}

type SignalRequest_Negotiate struct {
	Negotiate *SessionDescription `protobuf:"bytes,2,opt,name=negotiate,proto3,oneof"`
}

type SignalRequest_Trickle struct {
	Trickle *Trickle `protobuf:"bytes,3,opt,name=trickle,proto3,oneof"`
}

func (*SignalRequest_Join) isSignalRequest_Message() {}

func (*SignalRequest_Negotiate) isSignalRequest_Message() {}

func (*SignalRequest_Trickle) isSignalRequest_Message() {}

func (m *SignalRequest) GetMessage() isSignalRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignalRequest) GetJoin() *JoinRequest {
	if x, ok := m.GetMessage().(*SignalRequest_Join); ok {
		return x.Join
	}
	return nil
}

func (m *SignalRequest) GetNegotiate() *SessionDescription {
	if x, ok := m.GetMessage().(*SignalRequest_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (m *SignalRequest) GetTrickle() *Trickle {
	if x, ok := m.GetMessage().(*SignalRequest_Trickle); ok {
		return x.Trickle
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignalRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignalRequest_Join)(nil),
		(*SignalRequest_Negotiate)(nil),
		(*SignalRequest_Trickle)(nil),
	}
}

type SignalResponse struct {
	// Types that are valid to be assigned to Message:
	//	*SignalResponse_Jin
	//	*SignalResponse_Negotiate
	//	*SignalResponse_Trickle
	Message              isSignalResponse_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SignalResponse) Reset()         { *m = SignalResponse{} }
func (m *SignalResponse) String() string { return proto.CompactTextString(m) }
func (*SignalResponse) ProtoMessage()    {}
func (*SignalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bd91e80f756774, []int{1}
}

func (m *SignalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalResponse.Unmarshal(m, b)
}
func (m *SignalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalResponse.Marshal(b, m, deterministic)
}
func (m *SignalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalResponse.Merge(m, src)
}
func (m *SignalResponse) XXX_Size() int {
	return xxx_messageInfo_SignalResponse.Size(m)
}
func (m *SignalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignalResponse proto.InternalMessageInfo

type isSignalResponse_Message interface {
	isSignalResponse_Message()
}

type SignalResponse_Jin struct {
	Jin *JoinResponse `protobuf:"bytes,1,opt,name=jin,proto3,oneof"`
}

type SignalResponse_Negotiate struct {
	Negotiate *SessionDescription `protobuf:"bytes,2,opt,name=negotiate,proto3,oneof"`
}

type SignalResponse_Trickle struct {
	Trickle *Trickle `protobuf:"bytes,3,opt,name=trickle,proto3,oneof"`
}

func (*SignalResponse_Jin) isSignalResponse_Message() {}

func (*SignalResponse_Negotiate) isSignalResponse_Message() {}

func (*SignalResponse_Trickle) isSignalResponse_Message() {}

func (m *SignalResponse) GetMessage() isSignalResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignalResponse) GetJin() *JoinResponse {
	if x, ok := m.GetMessage().(*SignalResponse_Jin); ok {
		return x.Jin
	}
	return nil
}

func (m *SignalResponse) GetNegotiate() *SessionDescription {
	if x, ok := m.GetMessage().(*SignalResponse_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (m *SignalResponse) GetTrickle() *Trickle {
	if x, ok := m.GetMessage().(*SignalResponse_Trickle); ok {
		return x.Trickle
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignalResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignalResponse_Jin)(nil),
		(*SignalResponse_Negotiate)(nil),
		(*SignalResponse_Trickle)(nil),
	}
}

type JoinRequest struct {
	RoomId               string              `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Token                string              `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	PeerId               string              `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Offer                *SessionDescription `protobuf:"bytes,4,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bd91e80f756774, []int{2}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *JoinRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *JoinRequest) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *JoinRequest) GetOffer() *SessionDescription {
	if m != nil {
		return m.Offer
	}
	return nil
}

type JoinResponse struct {
	Answer               *SessionDescription `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bd91e80f756774, []int{3}
}

func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetAnswer() *SessionDescription {
	if m != nil {
		return m.Answer
	}
	return nil
}

type Trickle struct {
	// serialized JSON structure of candidate init
	CandidateInit        string   `protobuf:"bytes,1,opt,name=candidate_init,json=candidateInit,proto3" json:"candidate_init,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trickle) Reset()         { *m = Trickle{} }
func (m *Trickle) String() string { return proto.CompactTextString(m) }
func (*Trickle) ProtoMessage()    {}
func (*Trickle) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bd91e80f756774, []int{4}
}

func (m *Trickle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trickle.Unmarshal(m, b)
}
func (m *Trickle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trickle.Marshal(b, m, deterministic)
}
func (m *Trickle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trickle.Merge(m, src)
}
func (m *Trickle) XXX_Size() int {
	return xxx_messageInfo_Trickle.Size(m)
}
func (m *Trickle) XXX_DiscardUnknown() {
	xxx_messageInfo_Trickle.DiscardUnknown(m)
}

var xxx_messageInfo_Trickle proto.InternalMessageInfo

func (m *Trickle) GetCandidateInit() string {
	if m != nil {
		return m.CandidateInit
	}
	return ""
}

type SessionDescription struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sdp                  string   `protobuf:"bytes,2,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionDescription) Reset()         { *m = SessionDescription{} }
func (m *SessionDescription) String() string { return proto.CompactTextString(m) }
func (*SessionDescription) ProtoMessage()    {}
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bd91e80f756774, []int{5}
}

func (m *SessionDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDescription.Unmarshal(m, b)
}
func (m *SessionDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDescription.Marshal(b, m, deterministic)
}
func (m *SessionDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDescription.Merge(m, src)
}
func (m *SessionDescription) XXX_Size() int {
	return xxx_messageInfo_SessionDescription.Size(m)
}
func (m *SessionDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDescription proto.InternalMessageInfo

func (m *SessionDescription) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SessionDescription) GetSdp() string {
	if m != nil {
		return m.Sdp
	}
	return ""
}

func init() {
	proto.RegisterType((*SignalRequest)(nil), "livekit.SignalRequest")
	proto.RegisterType((*SignalResponse)(nil), "livekit.SignalResponse")
	proto.RegisterType((*JoinRequest)(nil), "livekit.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "livekit.JoinResponse")
	proto.RegisterType((*Trickle)(nil), "livekit.Trickle")
	proto.RegisterType((*SessionDescription)(nil), "livekit.SessionDescription")
}

func init() { proto.RegisterFile("rtc.proto", fileDescriptor_85bd91e80f756774) }

var fileDescriptor_85bd91e80f756774 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0x63, 0x92, 0xda, 0xf2, 0x94, 0x56, 0xd5, 0xa8, 0x50, 0x0b, 0x2e, 0xc8, 0x12, 0x52,
	0x41, 0xe0, 0xf4, 0xcf, 0x0d, 0xc4, 0xa5, 0xe5, 0xe0, 0xc0, 0x6d, 0xd3, 0x13, 0x97, 0xca, 0xb5,
	0xa7, 0x66, 0x9a, 0x64, 0xd7, 0xec, 0x6e, 0x83, 0x78, 0x03, 0x5e, 0x87, 0x37, 0x44, 0x5e, 0xaf,
	0xdd, 0x88, 0x48, 0x70, 0xe4, 0xe4, 0x9d, 0xd1, 0xf7, 0x5b, 0x7d, 0x33, 0xb6, 0x21, 0xd6, 0xb6,
	0xcc, 0x1a, 0xad, 0xac, 0xc2, 0x68, 0xc9, 0x6b, 0x5a, 0xb0, 0x4d, 0x7f, 0x05, 0xb0, 0x37, 0xe7,
	0x5a, 0x16, 0x4b, 0x41, 0xdf, 0xee, 0xc9, 0x58, 0x7c, 0x0d, 0x93, 0x3b, 0xc5, 0x32, 0x09, 0x5e,
	0x04, 0xc7, 0xbb, 0x67, 0x87, 0x99, 0x27, 0xb3, 0x4f, 0x8a, 0xa5, 0x67, 0xf2, 0x91, 0x70, 0x0c,
	0xbe, 0x87, 0x58, 0x52, 0xad, 0x2c, 0x17, 0x96, 0x92, 0x47, 0x2e, 0xf0, 0x7c, 0x08, 0xcc, 0xc9,
	0x18, 0x56, 0xf2, 0x23, 0x99, 0x52, 0x73, 0x63, 0x59, 0xc9, 0x7c, 0x24, 0x1e, 0x78, 0x7c, 0x03,
	0x91, 0xd5, 0x5c, 0x2e, 0x96, 0x94, 0x8c, 0x5d, 0xf4, 0x60, 0x88, 0x5e, 0x75, 0xfd, 0x7c, 0x24,
	0x7a, 0xe4, 0x22, 0x86, 0x68, 0x45, 0xc6, 0x14, 0x35, 0xb5, 0xce, 0xfb, 0xbd, 0xb3, 0x69, 0x94,
	0x34, 0x84, 0xaf, 0x60, 0x7c, 0x37, 0x38, 0x3f, 0xf9, 0xc3, 0xb9, 0x63, 0xf2, 0x91, 0x68, 0x99,
	0xff, 0xe4, 0xfc, 0x33, 0x80, 0xdd, 0x8d, 0x0d, 0xe2, 0x11, 0x44, 0x5a, 0xa9, 0xd5, 0x35, 0x57,
	0x4e, 0x3a, 0x16, 0x61, 0x5b, 0xce, 0x2a, 0x3c, 0x84, 0x1d, 0xab, 0x16, 0x24, 0x9d, 0x5a, 0x2c,
	0xba, 0xa2, 0xc5, 0x1b, 0x22, 0xdd, 0xe2, 0xe3, 0x0e, 0x6f, 0xcb, 0x59, 0x85, 0xa7, 0xb0, 0xa3,
	0x6e, 0x6f, 0x49, 0x27, 0x93, 0x7f, 0x4e, 0x22, 0x3a, 0x32, 0xbd, 0x84, 0xc7, 0x9b, 0x7b, 0xc1,
	0x73, 0x08, 0x0b, 0x69, 0xbe, 0x93, 0xf6, 0xeb, 0xfb, 0xeb, 0x1d, 0x1e, 0x4d, 0x4f, 0x20, 0xf2,
	0x03, 0xe3, 0x4b, 0xd8, 0x2f, 0x0b, 0x59, 0x71, 0x55, 0x58, 0xba, 0x66, 0xc9, 0xd6, 0x4f, 0xb4,
	0x37, 0x74, 0x67, 0x92, 0x6d, 0xfa, 0x0e, 0x70, 0xfb, 0x3e, 0x44, 0x98, 0xd8, 0x1f, 0x0d, 0xf9,
	0x88, 0x3b, 0xe3, 0x01, 0x8c, 0x4d, 0xd5, 0xf8, 0x05, 0xb4, 0xc7, 0xb3, 0xcf, 0x00, 0xe2, 0xea,
	0x72, 0x4e, 0x7a, 0xcd, 0x25, 0xe1, 0x07, 0x08, 0xbb, 0xd7, 0x8f, 0x4f, 0x1f, 0x54, 0x37, 0xbf,
	0xe1, 0x67, 0x47, 0x5b, 0xfd, 0x6e, 0xd6, 0xe3, 0xe0, 0x24, 0xb8, 0x38, 0xfd, 0x32, 0xad, 0xd9,
	0x7e, 0xbd, 0xbf, 0xc9, 0x4a, 0xb5, 0x9a, 0x7a, 0xb0, 0x7f, 0xbe, 0x35, 0xa4, 0xd7, 0xa4, 0xa7,
	0xee, 0x37, 0xe9, 0x9b, 0x37, 0xa1, 0x2b, 0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x3d,
	0xb9, 0xb1, 0x42, 0x03, 0x00, 0x00,
}
