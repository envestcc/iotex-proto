// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/blockchain.proto

package iotextypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// header of a block
type BlockHeader struct {
	Core                 *BlockHeaderCore `protobuf:"bytes,1,opt,name=core,proto3" json:"core,omitempty"`
	ProducerPubkey       []byte           `protobuf:"bytes,2,opt,name=producerPubkey,proto3" json:"producerPubkey,omitempty"`
	Signature            []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetCore() *BlockHeaderCore {
	if m != nil {
		return m.Core
	}
	return nil
}

func (m *BlockHeader) GetProducerPubkey() []byte {
	if m != nil {
		return m.ProducerPubkey
	}
	return nil
}

func (m *BlockHeader) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type BlockHeaderCore struct {
	Version              uint32               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Height               uint64               `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevBlockHash        []byte               `protobuf:"bytes,4,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	TxRoot               []byte               `protobuf:"bytes,5,opt,name=txRoot,proto3" json:"txRoot,omitempty"`
	DeltaStateDigest     []byte               `protobuf:"bytes,6,opt,name=deltaStateDigest,proto3" json:"deltaStateDigest,omitempty"`
	ReceiptRoot          []byte               `protobuf:"bytes,7,opt,name=receiptRoot,proto3" json:"receiptRoot,omitempty"`
	LogsBloom            []byte               `protobuf:"bytes,8,opt,name=logsBloom,proto3" json:"logsBloom,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockHeaderCore) Reset()         { *m = BlockHeaderCore{} }
func (m *BlockHeaderCore) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderCore) ProtoMessage()    {}
func (*BlockHeaderCore) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{1}
}

func (m *BlockHeaderCore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeaderCore.Unmarshal(m, b)
}
func (m *BlockHeaderCore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeaderCore.Marshal(b, m, deterministic)
}
func (m *BlockHeaderCore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderCore.Merge(m, src)
}
func (m *BlockHeaderCore) XXX_Size() int {
	return xxx_messageInfo_BlockHeaderCore.Size(m)
}
func (m *BlockHeaderCore) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderCore.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderCore proto.InternalMessageInfo

func (m *BlockHeaderCore) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeaderCore) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeaderCore) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BlockHeaderCore) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *BlockHeaderCore) GetTxRoot() []byte {
	if m != nil {
		return m.TxRoot
	}
	return nil
}

func (m *BlockHeaderCore) GetDeltaStateDigest() []byte {
	if m != nil {
		return m.DeltaStateDigest
	}
	return nil
}

func (m *BlockHeaderCore) GetReceiptRoot() []byte {
	if m != nil {
		return m.ReceiptRoot
	}
	return nil
}

func (m *BlockHeaderCore) GetLogsBloom() []byte {
	if m != nil {
		return m.LogsBloom
	}
	return nil
}

// footer of a block
type BlockFooter struct {
	Endorsements         []*Endorsement       `protobuf:"bytes,1,rep,name=endorsements,proto3" json:"endorsements,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockFooter) Reset()         { *m = BlockFooter{} }
func (m *BlockFooter) String() string { return proto.CompactTextString(m) }
func (*BlockFooter) ProtoMessage()    {}
func (*BlockFooter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{2}
}

func (m *BlockFooter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockFooter.Unmarshal(m, b)
}
func (m *BlockFooter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockFooter.Marshal(b, m, deterministic)
}
func (m *BlockFooter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockFooter.Merge(m, src)
}
func (m *BlockFooter) XXX_Size() int {
	return xxx_messageInfo_BlockFooter.Size(m)
}
func (m *BlockFooter) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockFooter.DiscardUnknown(m)
}

var xxx_messageInfo_BlockFooter proto.InternalMessageInfo

func (m *BlockFooter) GetEndorsements() []*Endorsement {
	if m != nil {
		return m.Endorsements
	}
	return nil
}

func (m *BlockFooter) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// body of a block
type BlockBody struct {
	Actions              []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BlockBody) Reset()         { *m = BlockBody{} }
func (m *BlockBody) String() string { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()    {}
func (*BlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{3}
}

func (m *BlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockBody.Unmarshal(m, b)
}
func (m *BlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockBody.Marshal(b, m, deterministic)
}
func (m *BlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockBody.Merge(m, src)
}
func (m *BlockBody) XXX_Size() int {
	return xxx_messageInfo_BlockBody.Size(m)
}
func (m *BlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BlockBody proto.InternalMessageInfo

func (m *BlockBody) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

// block consists of header followed by transactions
// hash of current block can be computed from header hence not stored
type Block struct {
	Header               *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body                 *BlockBody   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Footer               *BlockFooter `protobuf:"bytes,3,opt,name=footer,proto3" json:"footer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{4}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Block) GetFooter() *BlockFooter {
	if m != nil {
		return m.Footer
	}
	return nil
}

// Receipts consists of a collection of recepit
type Receipts struct {
	Receipts             []*Receipt `protobuf:"bytes,1,rep,name=receipts,proto3" json:"receipts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Receipts) Reset()         { *m = Receipts{} }
func (m *Receipts) String() string { return proto.CompactTextString(m) }
func (*Receipts) ProtoMessage()    {}
func (*Receipts) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{5}
}

func (m *Receipts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipts.Unmarshal(m, b)
}
func (m *Receipts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipts.Marshal(b, m, deterministic)
}
func (m *Receipts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipts.Merge(m, src)
}
func (m *Receipts) XXX_Size() int {
	return xxx_messageInfo_Receipts.Size(m)
}
func (m *Receipts) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipts.DiscardUnknown(m)
}

var xxx_messageInfo_Receipts proto.InternalMessageInfo

func (m *Receipts) GetReceipts() []*Receipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

type EpochData struct {
	Num                     uint64   `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Height                  uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	GravityChainStartHeight uint64   `protobuf:"varint,3,opt,name=gravityChainStartHeight,proto3" json:"gravityChainStartHeight,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *EpochData) Reset()         { *m = EpochData{} }
func (m *EpochData) String() string { return proto.CompactTextString(m) }
func (*EpochData) ProtoMessage()    {}
func (*EpochData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{6}
}

func (m *EpochData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EpochData.Unmarshal(m, b)
}
func (m *EpochData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EpochData.Marshal(b, m, deterministic)
}
func (m *EpochData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochData.Merge(m, src)
}
func (m *EpochData) XXX_Size() int {
	return xxx_messageInfo_EpochData.Size(m)
}
func (m *EpochData) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochData.DiscardUnknown(m)
}

var xxx_messageInfo_EpochData proto.InternalMessageInfo

func (m *EpochData) GetNum() uint64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *EpochData) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *EpochData) GetGravityChainStartHeight() uint64 {
	if m != nil {
		return m.GravityChainStartHeight
	}
	return 0
}

// Blockchain Metadata
type ChainMeta struct {
	Height               uint64     `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	NumActions           int64      `protobuf:"varint,2,opt,name=numActions,proto3" json:"numActions,omitempty"`
	Tps                  int64      `protobuf:"varint,3,opt,name=tps,proto3" json:"tps,omitempty"`
	Epoch                *EpochData `protobuf:"bytes,4,opt,name=epoch,proto3" json:"epoch,omitempty"`
	TpsFloat             float32    `protobuf:"fixed32,5,opt,name=tpsFloat,proto3" json:"tpsFloat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ChainMeta) Reset()         { *m = ChainMeta{} }
func (m *ChainMeta) String() string { return proto.CompactTextString(m) }
func (*ChainMeta) ProtoMessage()    {}
func (*ChainMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{7}
}

func (m *ChainMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainMeta.Unmarshal(m, b)
}
func (m *ChainMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainMeta.Marshal(b, m, deterministic)
}
func (m *ChainMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainMeta.Merge(m, src)
}
func (m *ChainMeta) XXX_Size() int {
	return xxx_messageInfo_ChainMeta.Size(m)
}
func (m *ChainMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ChainMeta proto.InternalMessageInfo

func (m *ChainMeta) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ChainMeta) GetNumActions() int64 {
	if m != nil {
		return m.NumActions
	}
	return 0
}

func (m *ChainMeta) GetTps() int64 {
	if m != nil {
		return m.Tps
	}
	return 0
}

func (m *ChainMeta) GetEpoch() *EpochData {
	if m != nil {
		return m.Epoch
	}
	return nil
}

func (m *ChainMeta) GetTpsFloat() float32 {
	if m != nil {
		return m.TpsFloat
	}
	return 0
}

// Block Metadata
type BlockMeta struct {
	Hash                 string               `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height               uint64               `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	NumActions           int64                `protobuf:"varint,4,opt,name=numActions,proto3" json:"numActions,omitempty"`
	ProducerAddress      string               `protobuf:"bytes,5,opt,name=producerAddress,proto3" json:"producerAddress,omitempty"`
	TransferAmount       string               `protobuf:"bytes,6,opt,name=transferAmount,proto3" json:"transferAmount,omitempty"`
	TxRoot               string               `protobuf:"bytes,7,opt,name=txRoot,proto3" json:"txRoot,omitempty"`
	ReceiptRoot          string               `protobuf:"bytes,8,opt,name=receiptRoot,proto3" json:"receiptRoot,omitempty"`
	DeltaStateDigest     string               `protobuf:"bytes,9,opt,name=deltaStateDigest,proto3" json:"deltaStateDigest,omitempty"`
	LogsBloom            []byte               `protobuf:"bytes,10,opt,name=logsBloom,proto3" json:"logsBloom,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockMeta) Reset()         { *m = BlockMeta{} }
func (m *BlockMeta) String() string { return proto.CompactTextString(m) }
func (*BlockMeta) ProtoMessage()    {}
func (*BlockMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{8}
}

func (m *BlockMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMeta.Unmarshal(m, b)
}
func (m *BlockMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMeta.Marshal(b, m, deterministic)
}
func (m *BlockMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMeta.Merge(m, src)
}
func (m *BlockMeta) XXX_Size() int {
	return xxx_messageInfo_BlockMeta.Size(m)
}
func (m *BlockMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMeta.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMeta proto.InternalMessageInfo

func (m *BlockMeta) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BlockMeta) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockMeta) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BlockMeta) GetNumActions() int64 {
	if m != nil {
		return m.NumActions
	}
	return 0
}

func (m *BlockMeta) GetProducerAddress() string {
	if m != nil {
		return m.ProducerAddress
	}
	return ""
}

func (m *BlockMeta) GetTransferAmount() string {
	if m != nil {
		return m.TransferAmount
	}
	return ""
}

func (m *BlockMeta) GetTxRoot() string {
	if m != nil {
		return m.TxRoot
	}
	return ""
}

func (m *BlockMeta) GetReceiptRoot() string {
	if m != nil {
		return m.ReceiptRoot
	}
	return ""
}

func (m *BlockMeta) GetDeltaStateDigest() string {
	if m != nil {
		return m.DeltaStateDigest
	}
	return ""
}

func (m *BlockMeta) GetLogsBloom() []byte {
	if m != nil {
		return m.LogsBloom
	}
	return nil
}

// Account Metadata
type AccountMeta struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance              string   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce                uint64   `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	PendingNonce         uint64   `protobuf:"varint,4,opt,name=pendingNonce,proto3" json:"pendingNonce,omitempty"`
	NumActions           uint64   `protobuf:"varint,5,opt,name=numActions,proto3" json:"numActions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountMeta) Reset()         { *m = AccountMeta{} }
func (m *AccountMeta) String() string { return proto.CompactTextString(m) }
func (*AccountMeta) ProtoMessage()    {}
func (*AccountMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e828f5966a7c29d, []int{9}
}

func (m *AccountMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountMeta.Unmarshal(m, b)
}
func (m *AccountMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountMeta.Marshal(b, m, deterministic)
}
func (m *AccountMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountMeta.Merge(m, src)
}
func (m *AccountMeta) XXX_Size() int {
	return xxx_messageInfo_AccountMeta.Size(m)
}
func (m *AccountMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountMeta.DiscardUnknown(m)
}

var xxx_messageInfo_AccountMeta proto.InternalMessageInfo

func (m *AccountMeta) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountMeta) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *AccountMeta) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *AccountMeta) GetPendingNonce() uint64 {
	if m != nil {
		return m.PendingNonce
	}
	return 0
}

func (m *AccountMeta) GetNumActions() uint64 {
	if m != nil {
		return m.NumActions
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "iotextypes.BlockHeader")
	proto.RegisterType((*BlockHeaderCore)(nil), "iotextypes.BlockHeaderCore")
	proto.RegisterType((*BlockFooter)(nil), "iotextypes.BlockFooter")
	proto.RegisterType((*BlockBody)(nil), "iotextypes.BlockBody")
	proto.RegisterType((*Block)(nil), "iotextypes.Block")
	proto.RegisterType((*Receipts)(nil), "iotextypes.Receipts")
	proto.RegisterType((*EpochData)(nil), "iotextypes.EpochData")
	proto.RegisterType((*ChainMeta)(nil), "iotextypes.ChainMeta")
	proto.RegisterType((*BlockMeta)(nil), "iotextypes.BlockMeta")
	proto.RegisterType((*AccountMeta)(nil), "iotextypes.AccountMeta")
}

func init() { proto.RegisterFile("proto/types/blockchain.proto", fileDescriptor_0e828f5966a7c29d) }

var fileDescriptor_0e828f5966a7c29d = []byte{
	// 768 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4b, 0x6f, 0xdb, 0x46,
	0x10, 0x06, 0x25, 0xca, 0x12, 0x47, 0x76, 0x6d, 0x6c, 0x1f, 0x26, 0x54, 0xb7, 0x15, 0x88, 0xa2,
	0x50, 0x5f, 0x22, 0xe0, 0xa2, 0x68, 0x0b, 0x9f, 0x24, 0x3f, 0xe0, 0x4b, 0x8b, 0x62, 0xdd, 0x53,
	0x6f, 0x2b, 0x72, 0x4d, 0xb1, 0x26, 0xb9, 0xc4, 0x72, 0x69, 0x58, 0xf7, 0xfe, 0x84, 0xdc, 0x02,
	0xe4, 0x9a, 0x9f, 0x96, 0xbf, 0x11, 0xec, 0x2c, 0x29, 0x91, 0x54, 0x1c, 0x20, 0x87, 0xdc, 0x38,
	0xdf, 0x7c, 0xda, 0x99, 0xf9, 0xe6, 0x21, 0x38, 0xcb, 0xa5, 0x50, 0xc2, 0x57, 0x9b, 0x9c, 0x17,
	0xfe, 0x2a, 0x11, 0xc1, 0x43, 0xb0, 0x66, 0x71, 0x36, 0x47, 0x98, 0x40, 0x2c, 0x14, 0x7f, 0x42,
	0xe7, 0xc4, 0x6d, 0x32, 0x59, 0xa0, 0x62, 0x51, 0xb1, 0x26, 0x5f, 0x35, 0x3d, 0x3c, 0x0b, 0x85,
	0x2c, 0x78, 0xca, 0x33, 0x55, 0xb9, 0xbf, 0x89, 0x84, 0x88, 0x12, 0xee, 0xa3, 0xb5, 0x2a, 0xef,
	0x7d, 0x15, 0xa7, 0xbc, 0x50, 0x2c, 0xcd, 0x0d, 0xc1, 0xfb, 0xdf, 0x82, 0xf1, 0x52, 0x87, 0xbe,
	0xe5, 0x2c, 0xe4, 0x92, 0xf8, 0x60, 0x07, 0x42, 0x72, 0xd7, 0x9a, 0x5a, 0xb3, 0xf1, 0xf9, 0x97,
	0xf3, 0x5d, 0x12, 0xf3, 0x06, 0xed, 0x52, 0x48, 0x4e, 0x91, 0x48, 0xbe, 0x83, 0x4f, 0x72, 0x29,
	0xc2, 0x32, 0xe0, 0xf2, 0xef, 0x72, 0xf5, 0xc0, 0x37, 0x6e, 0x6f, 0x6a, 0xcd, 0x0e, 0x69, 0x07,
	0x25, 0x67, 0xe0, 0x14, 0x71, 0x94, 0x31, 0x55, 0x4a, 0xee, 0xf6, 0x91, 0xb2, 0x03, 0xbc, 0xd7,
	0x3d, 0x38, 0xee, 0xbc, 0x4f, 0x5c, 0x18, 0x3e, 0x72, 0x59, 0xc4, 0x22, 0xc3, 0x6c, 0x8e, 0x68,
	0x6d, 0x92, 0x2f, 0xe0, 0x60, 0xcd, 0xe3, 0x68, 0xad, 0x30, 0x96, 0x4d, 0x2b, 0x8b, 0xfc, 0x0e,
	0xce, 0xb6, 0x3e, 0x8c, 0x31, 0x3e, 0x9f, 0xcc, 0x8d, 0x02, 0xf3, 0x5a, 0x81, 0xf9, 0x3f, 0x35,
	0x83, 0xee, 0xc8, 0xe4, 0x5b, 0x38, 0xca, 0x25, 0x7f, 0x34, 0x29, 0xb0, 0x62, 0xed, 0xda, 0x98,
	0x61, 0x1b, 0xd4, 0x71, 0xd5, 0x13, 0x15, 0x42, 0xb9, 0x03, 0x74, 0x57, 0x16, 0xf9, 0x01, 0x4e,
	0x42, 0x9e, 0x28, 0x76, 0xa7, 0x98, 0xe2, 0x57, 0x71, 0xc4, 0x0b, 0xe5, 0x1e, 0x20, 0x63, 0x0f,
	0x27, 0x53, 0x18, 0x4b, 0x1e, 0xf0, 0x38, 0x57, 0xf8, 0xd0, 0x10, 0x69, 0x4d, 0x48, 0x2b, 0x95,
	0x88, 0xa8, 0x58, 0x26, 0x42, 0xa4, 0xee, 0xc8, 0x28, 0xb5, 0x05, 0x76, 0x0d, 0xbb, 0x11, 0x42,
	0x71, 0x49, 0x2e, 0xe0, 0xb0, 0xd1, 0xf6, 0xc2, 0xb5, 0xa6, 0xfd, 0xd9, 0xf8, 0xfc, 0xb4, 0xd9,
	0xb8, 0xeb, 0x9d, 0x9f, 0xb6, 0xc8, 0x6d, 0xc1, 0x7a, 0x1f, 0x20, 0x98, 0xf7, 0x07, 0x38, 0x98,
	0xc5, 0x52, 0x84, 0x1b, 0xf2, 0x13, 0x0c, 0xcd, 0x50, 0xd6, 0xe1, 0x49, 0x33, 0xfc, 0x02, 0x5d,
	0xb4, 0xa6, 0x78, 0x2f, 0x2c, 0x18, 0xe0, 0x6f, 0x89, 0xaf, 0xfb, 0xa8, 0xfb, 0x5d, 0x8d, 0xdb,
	0xe9, 0x33, 0xe3, 0x46, 0x2b, 0x1a, 0xf9, 0x1e, 0xec, 0x95, 0x08, 0x37, 0x55, 0xaa, 0x9f, 0xef,
	0xd1, 0x75, 0x36, 0x14, 0x29, 0xfa, 0xed, 0x7b, 0x54, 0xa8, 0x1a, 0x84, 0xfd, 0xb7, 0x8d, 0x80,
	0xb4, 0xa2, 0x79, 0x17, 0x30, 0xa2, 0xa6, 0x0b, 0x05, 0xf1, 0x61, 0x54, 0x75, 0xa4, 0xae, 0xe8,
	0xd3, 0xe6, 0xcf, 0x2b, 0x1e, 0xdd, 0x92, 0x3c, 0x01, 0xce, 0x75, 0x2e, 0x82, 0xf5, 0x15, 0x53,
	0x8c, 0x9c, 0x40, 0x3f, 0x2b, 0x53, 0xac, 0xc9, 0xa6, 0xfa, 0xf3, 0x3d, 0x03, 0x7b, 0x1a, 0x49,
	0xf6, 0x18, 0xab, 0xcd, 0xa5, 0xde, 0xfc, 0x3b, 0xc5, 0xa4, 0xba, 0x35, 0xc4, 0x3e, 0x12, 0x9f,
	0x73, 0x7b, 0xaf, 0x2c, 0x70, 0x10, 0xfc, 0x93, 0x2b, 0xd6, 0x78, 0xdf, 0x6a, 0xbd, 0xff, 0x35,
	0x40, 0x56, 0xa6, 0x8b, 0xaa, 0x37, 0x3a, 0x76, 0x9f, 0x36, 0x10, 0x9d, 0xa9, 0xca, 0x0b, 0x8c,
	0xd5, 0xa7, 0xfa, 0x93, 0xfc, 0x08, 0x03, 0xae, 0x0b, 0xc1, 0x05, 0xe8, 0x48, 0xbc, 0xad, 0x90,
	0x1a, 0x0e, 0x99, 0xc0, 0x48, 0xe5, 0xc5, 0x4d, 0x22, 0x98, 0xd9, 0x88, 0x1e, 0xdd, 0xda, 0xde,
	0x9b, 0x5e, 0x35, 0x21, 0x98, 0x20, 0x01, 0x7b, 0xad, 0xd7, 0x4a, 0xa7, 0xe7, 0x50, 0xfc, 0xfe,
	0x08, 0x5b, 0xdc, 0x2e, 0xd7, 0xde, 0x2b, 0x77, 0x06, 0xc7, 0xf5, 0x55, 0x5a, 0x84, 0xa1, 0xe4,
	0x45, 0x81, 0x69, 0x3b, 0xb4, 0x0b, 0xeb, 0xab, 0xa6, 0x24, 0xcb, 0x8a, 0x7b, 0x2e, 0x17, 0xa9,
	0x28, 0x33, 0xb3, 0xcf, 0x0e, 0xed, 0xa0, 0x8d, 0x8b, 0x30, 0x44, 0x7f, 0x7d, 0x11, 0x3a, 0x5b,
	0x3e, 0x42, 0x67, 0x6b, 0xcb, 0xdf, 0x75, 0x33, 0x1c, 0xa4, 0xed, 0xdf, 0x8c, 0xd6, 0x45, 0x80,
	0xee, 0x45, 0x78, 0x69, 0xc1, 0x78, 0x11, 0x04, 0x3a, 0x1f, 0xd4, 0xda, 0x85, 0x21, 0xab, 0xaa,
	0x33, 0x72, 0xd7, 0xa6, 0xf6, 0xac, 0x58, 0xc2, 0xb2, 0x80, 0xa3, 0xe4, 0x0e, 0xad, 0x4d, 0xf2,
	0x19, 0x0c, 0x32, 0xa1, 0x71, 0x33, 0x76, 0xc6, 0x20, 0x1e, 0x1c, 0xe6, 0x3c, 0x0b, 0xe3, 0x2c,
	0xfa, 0x0b, 0x9d, 0x36, 0x3a, 0x5b, 0x58, 0x47, 0xf3, 0x01, 0x32, 0x1a, 0xc8, 0xf2, 0xb7, 0x7f,
	0x7f, 0x8d, 0x62, 0xb5, 0x2e, 0x57, 0xf3, 0x40, 0xa4, 0x3e, 0x4e, 0x53, 0x2e, 0xc5, 0x7f, 0x3c,
	0x50, 0xc6, 0xf8, 0xd9, 0xfc, 0x81, 0x45, 0x22, 0x61, 0x59, 0xe4, 0xef, 0xa6, 0x6d, 0x75, 0x80,
	0x8e, 0x5f, 0xde, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x91, 0xa3, 0x22, 0x26, 0x07, 0x00, 0x00,
}
