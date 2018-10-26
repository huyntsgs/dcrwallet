// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dicemix.proto

package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PeerInfo struct {
	PeerId               uint32   `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Pk                   []byte   `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	Sharedkey            []byte   `protobuf:"bytes,3,opt,name=sharedkey,proto3" json:"sharedkey,omitempty"`
	NumMsg               uint32   `protobuf:"varint,4,opt,name=num_msg,json=numMsg,proto3" json:"num_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerInfo) Reset()         { *m = PeerInfo{} }
func (m *PeerInfo) String() string { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()    {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{0}
}
func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerInfo.Unmarshal(m, b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
}
func (dst *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(dst, src)
}
func (m *PeerInfo) XXX_Size() int {
	return xxx_messageInfo_PeerInfo.Size(m)
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *PeerInfo) GetPk() []byte {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (m *PeerInfo) GetSharedkey() []byte {
	if m != nil {
		return m.Sharedkey
	}
	return nil
}

func (m *PeerInfo) GetNumMsg() uint32 {
	if m != nil {
		return m.NumMsg
	}
	return 0
}

type CoinJoinReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoinJoinReq) Reset()         { *m = CoinJoinReq{} }
func (m *CoinJoinReq) String() string { return proto.CompactTextString(m) }
func (*CoinJoinReq) ProtoMessage()    {}
func (*CoinJoinReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{1}
}
func (m *CoinJoinReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoinJoinReq.Unmarshal(m, b)
}
func (m *CoinJoinReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoinJoinReq.Marshal(b, m, deterministic)
}
func (dst *CoinJoinReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinJoinReq.Merge(dst, src)
}
func (m *CoinJoinReq) XXX_Size() int {
	return xxx_messageInfo_CoinJoinReq.Size(m)
}
func (m *CoinJoinReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinJoinReq.DiscardUnknown(m)
}

var xxx_messageInfo_CoinJoinReq proto.InternalMessageInfo

type CoinJoinRes struct {
	SessionId            uint32   `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	PeerId               uint32   `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	PeerIds              []uint64 `protobuf:"varint,4,rep,packed,name=peer_ids,json=peerIds,proto3" json:"peer_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoinJoinRes) Reset()         { *m = CoinJoinRes{} }
func (m *CoinJoinRes) String() string { return proto.CompactTextString(m) }
func (*CoinJoinRes) ProtoMessage()    {}
func (*CoinJoinRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{2}
}
func (m *CoinJoinRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoinJoinRes.Unmarshal(m, b)
}
func (m *CoinJoinRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoinJoinRes.Marshal(b, m, deterministic)
}
func (dst *CoinJoinRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinJoinRes.Merge(dst, src)
}
func (m *CoinJoinRes) XXX_Size() int {
	return xxx_messageInfo_CoinJoinRes.Size(m)
}
func (m *CoinJoinRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinJoinRes.DiscardUnknown(m)
}

var xxx_messageInfo_CoinJoinRes proto.InternalMessageInfo

func (m *CoinJoinRes) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CoinJoinRes) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *CoinJoinRes) GetPeerIds() []uint64 {
	if m != nil {
		return m.PeerIds
	}
	return nil
}

type KeyExchangeReq struct {
	SessionId            uint32   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	PeerId               uint32   `protobuf:"varint,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	NumMsg               uint32   `protobuf:"varint,3,opt,name=num_msg,json=numMsg,proto3" json:"num_msg,omitempty"`
	Pk                   []byte   `protobuf:"bytes,4,opt,name=pk,proto3" json:"pk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyExchangeReq) Reset()         { *m = KeyExchangeReq{} }
func (m *KeyExchangeReq) String() string { return proto.CompactTextString(m) }
func (*KeyExchangeReq) ProtoMessage()    {}
func (*KeyExchangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{3}
}
func (m *KeyExchangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExchangeReq.Unmarshal(m, b)
}
func (m *KeyExchangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExchangeReq.Marshal(b, m, deterministic)
}
func (dst *KeyExchangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExchangeReq.Merge(dst, src)
}
func (m *KeyExchangeReq) XXX_Size() int {
	return xxx_messageInfo_KeyExchangeReq.Size(m)
}
func (m *KeyExchangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExchangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExchangeReq proto.InternalMessageInfo

func (m *KeyExchangeReq) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *KeyExchangeReq) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *KeyExchangeReq) GetNumMsg() uint32 {
	if m != nil {
		return m.NumMsg
	}
	return 0
}

func (m *KeyExchangeReq) GetPk() []byte {
	if m != nil {
		return m.Pk
	}
	return nil
}

type KeyExchangeRes struct {
	SessionId            uint32      `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	PeerId               uint32      `protobuf:"varint,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Peers                []*PeerInfo `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *KeyExchangeRes) Reset()         { *m = KeyExchangeRes{} }
func (m *KeyExchangeRes) String() string { return proto.CompactTextString(m) }
func (*KeyExchangeRes) ProtoMessage()    {}
func (*KeyExchangeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{4}
}
func (m *KeyExchangeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExchangeRes.Unmarshal(m, b)
}
func (m *KeyExchangeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExchangeRes.Marshal(b, m, deterministic)
}
func (dst *KeyExchangeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExchangeRes.Merge(dst, src)
}
func (m *KeyExchangeRes) XXX_Size() int {
	return xxx_messageInfo_KeyExchangeRes.Size(m)
}
func (m *KeyExchangeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExchangeRes.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExchangeRes proto.InternalMessageInfo

func (m *KeyExchangeRes) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *KeyExchangeRes) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *KeyExchangeRes) GetPeers() []*PeerInfo {
	if m != nil {
		return m.Peers
	}
	return nil
}

type DcExpVector struct {
	PeerId               uint32   `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Len                  uint32   `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	Vector               []byte   `protobuf:"bytes,3,opt,name=vector,proto3" json:"vector,omitempty"`
	Commit               []byte   `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DcExpVector) Reset()         { *m = DcExpVector{} }
func (m *DcExpVector) String() string { return proto.CompactTextString(m) }
func (*DcExpVector) ProtoMessage()    {}
func (*DcExpVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{5}
}
func (m *DcExpVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcExpVector.Unmarshal(m, b)
}
func (m *DcExpVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcExpVector.Marshal(b, m, deterministic)
}
func (dst *DcExpVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcExpVector.Merge(dst, src)
}
func (m *DcExpVector) XXX_Size() int {
	return xxx_messageInfo_DcExpVector.Size(m)
}
func (m *DcExpVector) XXX_DiscardUnknown() {
	xxx_messageInfo_DcExpVector.DiscardUnknown(m)
}

var xxx_messageInfo_DcExpVector proto.InternalMessageInfo

func (m *DcExpVector) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *DcExpVector) GetLen() uint32 {
	if m != nil {
		return m.Len
	}
	return 0
}

func (m *DcExpVector) GetVector() []byte {
	if m != nil {
		return m.Vector
	}
	return nil
}

func (m *DcExpVector) GetCommit() []byte {
	if m != nil {
		return m.Commit
	}
	return nil
}

type DcXorVector struct {
	PeerId               uint32   `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Len                  uint32   `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	Vector               []byte   `protobuf:"bytes,3,opt,name=vector,proto3" json:"vector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DcXorVector) Reset()         { *m = DcXorVector{} }
func (m *DcXorVector) String() string { return proto.CompactTextString(m) }
func (*DcXorVector) ProtoMessage()    {}
func (*DcXorVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{6}
}
func (m *DcXorVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcXorVector.Unmarshal(m, b)
}
func (m *DcXorVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcXorVector.Marshal(b, m, deterministic)
}
func (dst *DcXorVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcXorVector.Merge(dst, src)
}
func (m *DcXorVector) XXX_Size() int {
	return xxx_messageInfo_DcXorVector.Size(m)
}
func (m *DcXorVector) XXX_DiscardUnknown() {
	xxx_messageInfo_DcXorVector.DiscardUnknown(m)
}

var xxx_messageInfo_DcXorVector proto.InternalMessageInfo

func (m *DcXorVector) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *DcXorVector) GetLen() uint32 {
	if m != nil {
		return m.Len
	}
	return 0
}

func (m *DcXorVector) GetVector() []byte {
	if m != nil {
		return m.Vector
	}
	return nil
}

type DcXorVectorResult struct {
	PeerIds              []uint32 `protobuf:"varint,1,rep,packed,name=peer_ids,json=peerIds,proto3" json:"peer_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DcXorVectorResult) Reset()         { *m = DcXorVectorResult{} }
func (m *DcXorVectorResult) String() string { return proto.CompactTextString(m) }
func (*DcXorVectorResult) ProtoMessage()    {}
func (*DcXorVectorResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{7}
}
func (m *DcXorVectorResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcXorVectorResult.Unmarshal(m, b)
}
func (m *DcXorVectorResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcXorVectorResult.Marshal(b, m, deterministic)
}
func (dst *DcXorVectorResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcXorVectorResult.Merge(dst, src)
}
func (m *DcXorVectorResult) XXX_Size() int {
	return xxx_messageInfo_DcXorVectorResult.Size(m)
}
func (m *DcXorVectorResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DcXorVectorResult.DiscardUnknown(m)
}

var xxx_messageInfo_DcXorVectorResult proto.InternalMessageInfo

func (m *DcXorVectorResult) GetPeerIds() []uint32 {
	if m != nil {
		return m.PeerIds
	}
	return nil
}

type AllMessages struct {
	Len                  uint32   `protobuf:"varint,1,opt,name=len,proto3" json:"len,omitempty"`
	Msgs                 []byte   `protobuf:"bytes,2,opt,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllMessages) Reset()         { *m = AllMessages{} }
func (m *AllMessages) String() string { return proto.CompactTextString(m) }
func (*AllMessages) ProtoMessage()    {}
func (*AllMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{8}
}
func (m *AllMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllMessages.Unmarshal(m, b)
}
func (m *AllMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllMessages.Marshal(b, m, deterministic)
}
func (dst *AllMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllMessages.Merge(dst, src)
}
func (m *AllMessages) XXX_Size() int {
	return xxx_messageInfo_AllMessages.Size(m)
}
func (m *AllMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_AllMessages.DiscardUnknown(m)
}

var xxx_messageInfo_AllMessages proto.InternalMessageInfo

func (m *AllMessages) GetLen() uint32 {
	if m != nil {
		return m.Len
	}
	return 0
}

func (m *AllMessages) GetMsgs() []byte {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type TxInputs struct {
	PeerId               uint32   `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	TicketPrice          int64    `protobuf:"varint,2,opt,name=ticket_price,json=ticketPrice,proto3" json:"ticket_price,omitempty"`
	Txins                []byte   `protobuf:"bytes,3,opt,name=txins,proto3" json:"txins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxInputs) Reset()         { *m = TxInputs{} }
func (m *TxInputs) String() string { return proto.CompactTextString(m) }
func (*TxInputs) ProtoMessage()    {}
func (*TxInputs) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{9}
}
func (m *TxInputs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxInputs.Unmarshal(m, b)
}
func (m *TxInputs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxInputs.Marshal(b, m, deterministic)
}
func (dst *TxInputs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInputs.Merge(dst, src)
}
func (m *TxInputs) XXX_Size() int {
	return xxx_messageInfo_TxInputs.Size(m)
}
func (m *TxInputs) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInputs.DiscardUnknown(m)
}

var xxx_messageInfo_TxInputs proto.InternalMessageInfo

func (m *TxInputs) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *TxInputs) GetTicketPrice() int64 {
	if m != nil {
		return m.TicketPrice
	}
	return 0
}

func (m *TxInputs) GetTxins() []byte {
	if m != nil {
		return m.Txins
	}
	return nil
}

type JoinTx struct {
	PeerId               uint32   `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Tx                   []byte   `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinTx) Reset()         { *m = JoinTx{} }
func (m *JoinTx) String() string { return proto.CompactTextString(m) }
func (*JoinTx) ProtoMessage()    {}
func (*JoinTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{10}
}
func (m *JoinTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinTx.Unmarshal(m, b)
}
func (m *JoinTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinTx.Marshal(b, m, deterministic)
}
func (dst *JoinTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinTx.Merge(dst, src)
}
func (m *JoinTx) XXX_Size() int {
	return xxx_messageInfo_JoinTx.Size(m)
}
func (m *JoinTx) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinTx.DiscardUnknown(m)
}

var xxx_messageInfo_JoinTx proto.InternalMessageInfo

func (m *JoinTx) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *JoinTx) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type PublishResult struct {
	PeerId               uint32   `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Tx                   []byte   `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResult) Reset()         { *m = PublishResult{} }
func (m *PublishResult) String() string { return proto.CompactTextString(m) }
func (*PublishResult) ProtoMessage()    {}
func (*PublishResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dicemix_6a80c36df978aca2, []int{11}
}
func (m *PublishResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResult.Unmarshal(m, b)
}
func (m *PublishResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResult.Marshal(b, m, deterministic)
}
func (dst *PublishResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResult.Merge(dst, src)
}
func (m *PublishResult) XXX_Size() int {
	return xxx_messageInfo_PublishResult.Size(m)
}
func (m *PublishResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResult.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResult proto.InternalMessageInfo

func (m *PublishResult) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *PublishResult) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerInfo)(nil), "messages.PeerInfo")
	proto.RegisterType((*CoinJoinReq)(nil), "messages.CoinJoinReq")
	proto.RegisterType((*CoinJoinRes)(nil), "messages.CoinJoinRes")
	proto.RegisterType((*KeyExchangeReq)(nil), "messages.KeyExchangeReq")
	proto.RegisterType((*KeyExchangeRes)(nil), "messages.KeyExchangeRes")
	proto.RegisterType((*DcExpVector)(nil), "messages.DcExpVector")
	proto.RegisterType((*DcXorVector)(nil), "messages.DcXorVector")
	proto.RegisterType((*DcXorVectorResult)(nil), "messages.DcXorVectorResult")
	proto.RegisterType((*AllMessages)(nil), "messages.AllMessages")
	proto.RegisterType((*TxInputs)(nil), "messages.TxInputs")
	proto.RegisterType((*JoinTx)(nil), "messages.JoinTx")
	proto.RegisterType((*PublishResult)(nil), "messages.PublishResult")
}

func init() { proto.RegisterFile("dicemix.proto", fileDescriptor_dicemix_6a80c36df978aca2) }

var fileDescriptor_dicemix_6a80c36df978aca2 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xc9, 0xe5, 0x7a, 0xbd, 0x4e, 0x9a, 0xa2, 0x8b, 0xe8, 0x09, 0x0a, 0xba, 0x4f, 0x79,
	0x0a, 0x68, 0x5f, 0x7c, 0x15, 0xed, 0xc3, 0x29, 0x85, 0x10, 0x8a, 0xf8, 0x20, 0x9c, 0xd7, 0x64,
	0x4d, 0xd6, 0x4b, 0x76, 0xd3, 0xcc, 0x46, 0xb6, 0xdf, 0xde, 0xdd, 0x64, 0xc3, 0xe5, 0x0a, 0x2d,
	0x22, 0x7d, 0x08, 0xcc, 0xcc, 0x66, 0xe7, 0x37, 0x33, 0xff, 0x1d, 0x08, 0x73, 0x9e, 0xb1, 0x9a,
	0xeb, 0xb8, 0x69, 0xa5, 0x92, 0x64, 0x59, 0x33, 0xc4, 0x6d, 0xc1, 0x90, 0xfe, 0x86, 0x65, 0xc2,
	0x58, 0xbb, 0x16, 0xbf, 0x24, 0x79, 0x01, 0xc7, 0x8d, 0xb1, 0x37, 0x3c, 0x5f, 0x79, 0x6f, 0xbc,
	0x28, 0x4c, 0x17, 0xd6, 0x5d, 0xe7, 0xe4, 0x0c, 0x66, 0xcd, 0x6e, 0x35, 0x33, 0xb1, 0xd3, 0xd4,
	0x58, 0xe4, 0x15, 0x9c, 0x60, 0xb9, 0x6d, 0x59, 0xbe, 0x63, 0xb7, 0x2b, 0xbf, 0x0f, 0xef, 0x03,
	0x36, 0x8d, 0xe8, 0xea, 0x4d, 0x8d, 0xc5, 0x6a, 0x3e, 0xa4, 0x31, 0xee, 0x25, 0x16, 0x34, 0x84,
	0xe0, 0x93, 0xe4, 0xe2, 0x8b, 0xf9, 0x52, 0x76, 0x43, 0x7f, 0x4e, 0x5d, 0x24, 0xaf, 0x01, 0xd0,
	0x54, 0xc5, 0xa5, 0xb0, 0x05, 0xcc, 0xfa, 0x9b, 0x27, 0x2e, 0x62, 0x6a, 0x98, 0x14, 0xe7, 0x1f,
	0x14, 0xf7, 0x12, 0x96, 0xee, 0x00, 0x0d, 0xcf, 0x8f, 0xe6, 0xe9, 0xf1, 0x70, 0x82, 0xf4, 0x06,
	0xce, 0xbe, 0xb2, 0xdb, 0x0b, 0x9d, 0x95, 0x5b, 0x51, 0x30, 0xc3, 0xbc, 0x03, 0xf1, 0x1e, 0x80,
	0xcc, 0x0e, 0x20, 0x93, 0x9e, 0xfc, 0x69, 0x4f, 0x6e, 0x34, 0xf3, 0x71, 0x34, 0xb4, 0xbd, 0x83,
	0xc4, 0xff, 0x46, 0x46, 0x70, 0x64, 0x2d, 0x34, 0x40, 0x3f, 0x0a, 0xde, 0x93, 0x78, 0xd4, 0x2c,
	0x1e, 0x05, 0x4b, 0x87, 0x1f, 0x68, 0x09, 0xc1, 0xe7, 0xec, 0x42, 0x37, 0xdf, 0x58, 0xa6, 0x64,
	0x7b, 0xbf, 0x8c, 0x4f, 0xc0, 0xaf, 0x98, 0x70, 0x18, 0x6b, 0x92, 0xe7, 0xb0, 0xf8, 0xd3, 0x5f,
	0x72, 0x2a, 0x3a, 0xcf, 0xc6, 0x33, 0x59, 0xd7, 0x5c, 0xb9, 0xce, 0x9c, 0x47, 0x13, 0x4b, 0xfa,
	0x2e, 0xdb, 0x47, 0x23, 0xd1, 0x18, 0x9e, 0x4e, 0x32, 0x9a, 0x79, 0x75, 0x95, 0x3a, 0x90, 0xd4,
	0x33, 0xdd, 0x87, 0x7b, 0x49, 0xcf, 0x21, 0xf8, 0x58, 0x55, 0x97, 0x6e, 0x14, 0x23, 0xc8, 0xdb,
	0x83, 0x08, 0xcc, 0x8d, 0x4a, 0xe8, 0x5e, 0x6b, 0x6f, 0xd3, 0x1f, 0xb0, 0xbc, 0xd2, 0x6b, 0xd1,
	0x74, 0x0a, 0xef, 0xaf, 0xf9, 0x2d, 0x9c, 0x2a, 0x9e, 0xed, 0x98, 0xda, 0x34, 0xad, 0x59, 0x96,
	0x3e, 0x81, 0x9f, 0x06, 0x43, 0x2c, 0xb1, 0x21, 0xf2, 0x0c, 0x8e, 0x94, 0xe6, 0x02, 0x5d, 0x0f,
	0x83, 0x43, 0xdf, 0xc1, 0xc2, 0xbe, 0xe1, 0x2b, 0xfd, 0xe0, 0x02, 0x29, 0x3d, 0x2e, 0x90, 0xd2,
	0xf4, 0x03, 0x84, 0x49, 0x77, 0x5d, 0x71, 0x2c, 0x5d, 0xc7, 0xff, 0x7a, 0xf3, 0x7a, 0xd1, 0x2f,
	0xf0, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x39, 0xe6, 0xb2, 0xd1, 0x03, 0x00, 0x00,
}