// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ranksrv.proto

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

// 通用排行榜定义
type RANK_TYPE int32

const (
	RANK_TYPE_DEFAULT RANK_TYPE = 0
	RANK_TYPE_PART    RANK_TYPE = 1
	RANK_TYPE_WHOLE   RANK_TYPE = 2
)

var RANK_TYPE_name = map[int32]string{
	0: "DEFAULT",
	1: "PART",
	2: "WHOLE",
}

var RANK_TYPE_value = map[string]int32{
	"DEFAULT": 0,
	"PART":    1,
	"WHOLE":   2,
}

func (x RANK_TYPE) String() string {
	return proto.EnumName(RANK_TYPE_name, int32(x))
}

func (RANK_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{0}
}

type RankInfo struct {
	Name                 string    `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type                 RANK_TYPE `protobuf:"varint,2,opt,name=Type,proto3,enum=pb.RANK_TYPE" json:"Type,omitempty"`
	TopNum               uint32    `protobuf:"varint,3,opt,name=TopNum,proto3" json:"TopNum,omitempty"`
	FlushInterval        uint32    `protobuf:"varint,4,opt,name=FlushInterval,proto3" json:"FlushInterval,omitempty"`
	ScoreMinLimit        uint64    `protobuf:"varint,5,opt,name=ScoreMinLimit,proto3" json:"ScoreMinLimit,omitempty"`
	ScoreNum             uint32    `protobuf:"varint,6,opt,name=ScoreNum,proto3" json:"ScoreNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RankInfo) Reset()         { *m = RankInfo{} }
func (m *RankInfo) String() string { return proto.CompactTextString(m) }
func (*RankInfo) ProtoMessage()    {}
func (*RankInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{0}
}

func (m *RankInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankInfo.Unmarshal(m, b)
}
func (m *RankInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankInfo.Marshal(b, m, deterministic)
}
func (m *RankInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankInfo.Merge(m, src)
}
func (m *RankInfo) XXX_Size() int {
	return xxx_messageInfo_RankInfo.Size(m)
}
func (m *RankInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RankInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RankInfo proto.InternalMessageInfo

func (m *RankInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankInfo) GetType() RANK_TYPE {
	if m != nil {
		return m.Type
	}
	return RANK_TYPE_DEFAULT
}

func (m *RankInfo) GetTopNum() uint32 {
	if m != nil {
		return m.TopNum
	}
	return 0
}

func (m *RankInfo) GetFlushInterval() uint32 {
	if m != nil {
		return m.FlushInterval
	}
	return 0
}

func (m *RankInfo) GetScoreMinLimit() uint64 {
	if m != nil {
		return m.ScoreMinLimit
	}
	return 0
}

func (m *RankInfo) GetScoreNum() uint32 {
	if m != nil {
		return m.ScoreNum
	}
	return 0
}

type RankUnitData struct {
	UniqueID             string   `protobuf:"bytes,1,opt,name=UniqueID,proto3" json:"UniqueID,omitempty"`
	Score                []uint64 `protobuf:"varint,2,rep,packed,name=Score,proto3" json:"Score,omitempty"`
	UpdateTime           uint32   `protobuf:"varint,3,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankUnitData) Reset()         { *m = RankUnitData{} }
func (m *RankUnitData) String() string { return proto.CompactTextString(m) }
func (*RankUnitData) ProtoMessage()    {}
func (*RankUnitData) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{1}
}

func (m *RankUnitData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankUnitData.Unmarshal(m, b)
}
func (m *RankUnitData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankUnitData.Marshal(b, m, deterministic)
}
func (m *RankUnitData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankUnitData.Merge(m, src)
}
func (m *RankUnitData) XXX_Size() int {
	return xxx_messageInfo_RankUnitData.Size(m)
}
func (m *RankUnitData) XXX_DiscardUnknown() {
	xxx_messageInfo_RankUnitData.DiscardUnknown(m)
}

var xxx_messageInfo_RankUnitData proto.InternalMessageInfo

func (m *RankUnitData) GetUniqueID() string {
	if m != nil {
		return m.UniqueID
	}
	return ""
}

func (m *RankUnitData) GetScore() []uint64 {
	if m != nil {
		return m.Score
	}
	return nil
}

func (m *RankUnitData) GetUpdateTime() uint32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type UnitRankInfo struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=Rank,proto3" json:"Rank,omitempty"`
	Score                []uint64 `protobuf:"varint,2,rep,packed,name=Score,proto3" json:"Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnitRankInfo) Reset()         { *m = UnitRankInfo{} }
func (m *UnitRankInfo) String() string { return proto.CompactTextString(m) }
func (*UnitRankInfo) ProtoMessage()    {}
func (*UnitRankInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{2}
}

func (m *UnitRankInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnitRankInfo.Unmarshal(m, b)
}
func (m *UnitRankInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnitRankInfo.Marshal(b, m, deterministic)
}
func (m *UnitRankInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnitRankInfo.Merge(m, src)
}
func (m *UnitRankInfo) XXX_Size() int {
	return xxx_messageInfo_UnitRankInfo.Size(m)
}
func (m *UnitRankInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UnitRankInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UnitRankInfo proto.InternalMessageInfo

func (m *UnitRankInfo) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *UnitRankInfo) GetScore() []uint64 {
	if m != nil {
		return m.Score
	}
	return nil
}

type RspHead struct {
	RetCode              uint32   `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspHead) Reset()         { *m = RspHead{} }
func (m *RspHead) String() string { return proto.CompactTextString(m) }
func (*RspHead) ProtoMessage()    {}
func (*RspHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{3}
}

func (m *RspHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspHead.Unmarshal(m, b)
}
func (m *RspHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspHead.Marshal(b, m, deterministic)
}
func (m *RspHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspHead.Merge(m, src)
}
func (m *RspHead) XXX_Size() int {
	return xxx_messageInfo_RspHead.Size(m)
}
func (m *RspHead) XXX_DiscardUnknown() {
	xxx_messageInfo_RspHead.DiscardUnknown(m)
}

var xxx_messageInfo_RspHead proto.InternalMessageInfo

func (m *RspHead) GetRetCode() uint32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

type RankInitReq struct {
	Info                 *RankInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RankInitReq) Reset()         { *m = RankInitReq{} }
func (m *RankInitReq) String() string { return proto.CompactTextString(m) }
func (*RankInitReq) ProtoMessage()    {}
func (*RankInitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{4}
}

func (m *RankInitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankInitReq.Unmarshal(m, b)
}
func (m *RankInitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankInitReq.Marshal(b, m, deterministic)
}
func (m *RankInitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankInitReq.Merge(m, src)
}
func (m *RankInitReq) XXX_Size() int {
	return xxx_messageInfo_RankInitReq.Size(m)
}
func (m *RankInitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankInitReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankInitReq proto.InternalMessageInfo

func (m *RankInitReq) GetInfo() *RankInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type RankInitRsp struct {
	Head                 *RspHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankInitRsp) Reset()         { *m = RankInitRsp{} }
func (m *RankInitRsp) String() string { return proto.CompactTextString(m) }
func (*RankInitRsp) ProtoMessage()    {}
func (*RankInitRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{5}
}

func (m *RankInitRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankInitRsp.Unmarshal(m, b)
}
func (m *RankInitRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankInitRsp.Marshal(b, m, deterministic)
}
func (m *RankInitRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankInitRsp.Merge(m, src)
}
func (m *RankInitRsp) XXX_Size() int {
	return xxx_messageInfo_RankInitRsp.Size(m)
}
func (m *RankInitRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankInitRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RankInitRsp proto.InternalMessageInfo

func (m *RankInitRsp) GetHead() *RspHead {
	if m != nil {
		return m.Head
	}
	return nil
}

type RankDataUpdateReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	UniqueID             string   `protobuf:"bytes,2,opt,name=UniqueID,proto3" json:"UniqueID,omitempty"`
	Score                []uint64 `protobuf:"varint,3,rep,packed,name=Score,proto3" json:"Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankDataUpdateReq) Reset()         { *m = RankDataUpdateReq{} }
func (m *RankDataUpdateReq) String() string { return proto.CompactTextString(m) }
func (*RankDataUpdateReq) ProtoMessage()    {}
func (*RankDataUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{6}
}

func (m *RankDataUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankDataUpdateReq.Unmarshal(m, b)
}
func (m *RankDataUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankDataUpdateReq.Marshal(b, m, deterministic)
}
func (m *RankDataUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankDataUpdateReq.Merge(m, src)
}
func (m *RankDataUpdateReq) XXX_Size() int {
	return xxx_messageInfo_RankDataUpdateReq.Size(m)
}
func (m *RankDataUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankDataUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankDataUpdateReq proto.InternalMessageInfo

func (m *RankDataUpdateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankDataUpdateReq) GetUniqueID() string {
	if m != nil {
		return m.UniqueID
	}
	return ""
}

func (m *RankDataUpdateReq) GetScore() []uint64 {
	if m != nil {
		return m.Score
	}
	return nil
}

type RankDataUpdateRsp struct {
	Head                 *RspHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankDataUpdateRsp) Reset()         { *m = RankDataUpdateRsp{} }
func (m *RankDataUpdateRsp) String() string { return proto.CompactTextString(m) }
func (*RankDataUpdateRsp) ProtoMessage()    {}
func (*RankDataUpdateRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{7}
}

func (m *RankDataUpdateRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankDataUpdateRsp.Unmarshal(m, b)
}
func (m *RankDataUpdateRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankDataUpdateRsp.Marshal(b, m, deterministic)
}
func (m *RankDataUpdateRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankDataUpdateRsp.Merge(m, src)
}
func (m *RankDataUpdateRsp) XXX_Size() int {
	return xxx_messageInfo_RankDataUpdateRsp.Size(m)
}
func (m *RankDataUpdateRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankDataUpdateRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RankDataUpdateRsp proto.InternalMessageInfo

func (m *RankDataUpdateRsp) GetHead() *RspHead {
	if m != nil {
		return m.Head
	}
	return nil
}

type RankQueryReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	UniqueID             string   `protobuf:"bytes,2,opt,name=UniqueID,proto3" json:"UniqueID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankQueryReq) Reset()         { *m = RankQueryReq{} }
func (m *RankQueryReq) String() string { return proto.CompactTextString(m) }
func (*RankQueryReq) ProtoMessage()    {}
func (*RankQueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{8}
}

func (m *RankQueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankQueryReq.Unmarshal(m, b)
}
func (m *RankQueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankQueryReq.Marshal(b, m, deterministic)
}
func (m *RankQueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankQueryReq.Merge(m, src)
}
func (m *RankQueryReq) XXX_Size() int {
	return xxx_messageInfo_RankQueryReq.Size(m)
}
func (m *RankQueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankQueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankQueryReq proto.InternalMessageInfo

func (m *RankQueryReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankQueryReq) GetUniqueID() string {
	if m != nil {
		return m.UniqueID
	}
	return ""
}

type RankQueryRsp struct {
	Head                 *RspHead      `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	RankInfo             *UnitRankInfo `protobuf:"bytes,2,opt,name=RankInfo,proto3" json:"RankInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RankQueryRsp) Reset()         { *m = RankQueryRsp{} }
func (m *RankQueryRsp) String() string { return proto.CompactTextString(m) }
func (*RankQueryRsp) ProtoMessage()    {}
func (*RankQueryRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{9}
}

func (m *RankQueryRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankQueryRsp.Unmarshal(m, b)
}
func (m *RankQueryRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankQueryRsp.Marshal(b, m, deterministic)
}
func (m *RankQueryRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankQueryRsp.Merge(m, src)
}
func (m *RankQueryRsp) XXX_Size() int {
	return xxx_messageInfo_RankQueryRsp.Size(m)
}
func (m *RankQueryRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankQueryRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RankQueryRsp proto.InternalMessageInfo

func (m *RankQueryRsp) GetHead() *RspHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *RankQueryRsp) GetRankInfo() *UnitRankInfo {
	if m != nil {
		return m.RankInfo
	}
	return nil
}

type RankTopQueryReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ReqNum               uint32   `protobuf:"varint,2,opt,name=ReqNum,proto3" json:"ReqNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankTopQueryReq) Reset()         { *m = RankTopQueryReq{} }
func (m *RankTopQueryReq) String() string { return proto.CompactTextString(m) }
func (*RankTopQueryReq) ProtoMessage()    {}
func (*RankTopQueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{10}
}

func (m *RankTopQueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankTopQueryReq.Unmarshal(m, b)
}
func (m *RankTopQueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankTopQueryReq.Marshal(b, m, deterministic)
}
func (m *RankTopQueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankTopQueryReq.Merge(m, src)
}
func (m *RankTopQueryReq) XXX_Size() int {
	return xxx_messageInfo_RankTopQueryReq.Size(m)
}
func (m *RankTopQueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankTopQueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankTopQueryReq proto.InternalMessageInfo

func (m *RankTopQueryReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankTopQueryReq) GetReqNum() uint32 {
	if m != nil {
		return m.ReqNum
	}
	return 0
}

type RankTopQueryRsp struct {
	Head                 *RspHead        `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	TopRank              []*RankUnitData `protobuf:"bytes,2,rep,name=TopRank,proto3" json:"TopRank,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RankTopQueryRsp) Reset()         { *m = RankTopQueryRsp{} }
func (m *RankTopQueryRsp) String() string { return proto.CompactTextString(m) }
func (*RankTopQueryRsp) ProtoMessage()    {}
func (*RankTopQueryRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{11}
}

func (m *RankTopQueryRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankTopQueryRsp.Unmarshal(m, b)
}
func (m *RankTopQueryRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankTopQueryRsp.Marshal(b, m, deterministic)
}
func (m *RankTopQueryRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankTopQueryRsp.Merge(m, src)
}
func (m *RankTopQueryRsp) XXX_Size() int {
	return xxx_messageInfo_RankTopQueryRsp.Size(m)
}
func (m *RankTopQueryRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankTopQueryRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RankTopQueryRsp proto.InternalMessageInfo

func (m *RankTopQueryRsp) GetHead() *RspHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *RankTopQueryRsp) GetTopRank() []*RankUnitData {
	if m != nil {
		return m.TopRank
	}
	return nil
}

type RankDataDeleteReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	UniqueID             string   `protobuf:"bytes,2,opt,name=UniqueID,proto3" json:"UniqueID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankDataDeleteReq) Reset()         { *m = RankDataDeleteReq{} }
func (m *RankDataDeleteReq) String() string { return proto.CompactTextString(m) }
func (*RankDataDeleteReq) ProtoMessage()    {}
func (*RankDataDeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{12}
}

func (m *RankDataDeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankDataDeleteReq.Unmarshal(m, b)
}
func (m *RankDataDeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankDataDeleteReq.Marshal(b, m, deterministic)
}
func (m *RankDataDeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankDataDeleteReq.Merge(m, src)
}
func (m *RankDataDeleteReq) XXX_Size() int {
	return xxx_messageInfo_RankDataDeleteReq.Size(m)
}
func (m *RankDataDeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankDataDeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankDataDeleteReq proto.InternalMessageInfo

func (m *RankDataDeleteReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankDataDeleteReq) GetUniqueID() string {
	if m != nil {
		return m.UniqueID
	}
	return ""
}

type RankDataDeleteRsp struct {
	Head                 *RspHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankDataDeleteRsp) Reset()         { *m = RankDataDeleteRsp{} }
func (m *RankDataDeleteRsp) String() string { return proto.CompactTextString(m) }
func (*RankDataDeleteRsp) ProtoMessage()    {}
func (*RankDataDeleteRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{13}
}

func (m *RankDataDeleteRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankDataDeleteRsp.Unmarshal(m, b)
}
func (m *RankDataDeleteRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankDataDeleteRsp.Marshal(b, m, deterministic)
}
func (m *RankDataDeleteRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankDataDeleteRsp.Merge(m, src)
}
func (m *RankDataDeleteRsp) XXX_Size() int {
	return xxx_messageInfo_RankDataDeleteRsp.Size(m)
}
func (m *RankDataDeleteRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankDataDeleteRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RankDataDeleteRsp proto.InternalMessageInfo

func (m *RankDataDeleteRsp) GetHead() *RspHead {
	if m != nil {
		return m.Head
	}
	return nil
}

type RankCloseReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankCloseReq) Reset()         { *m = RankCloseReq{} }
func (m *RankCloseReq) String() string { return proto.CompactTextString(m) }
func (*RankCloseReq) ProtoMessage()    {}
func (*RankCloseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{14}
}

func (m *RankCloseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankCloseReq.Unmarshal(m, b)
}
func (m *RankCloseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankCloseReq.Marshal(b, m, deterministic)
}
func (m *RankCloseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankCloseReq.Merge(m, src)
}
func (m *RankCloseReq) XXX_Size() int {
	return xxx_messageInfo_RankCloseReq.Size(m)
}
func (m *RankCloseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankCloseReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankCloseReq proto.InternalMessageInfo

func (m *RankCloseReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RankCloseRsp struct {
	Head                 *RspHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankCloseRsp) Reset()         { *m = RankCloseRsp{} }
func (m *RankCloseRsp) String() string { return proto.CompactTextString(m) }
func (*RankCloseRsp) ProtoMessage()    {}
func (*RankCloseRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{15}
}

func (m *RankCloseRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankCloseRsp.Unmarshal(m, b)
}
func (m *RankCloseRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankCloseRsp.Marshal(b, m, deterministic)
}
func (m *RankCloseRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankCloseRsp.Merge(m, src)
}
func (m *RankCloseRsp) XXX_Size() int {
	return xxx_messageInfo_RankCloseRsp.Size(m)
}
func (m *RankCloseRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankCloseRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RankCloseRsp proto.InternalMessageInfo

func (m *RankCloseRsp) GetHead() *RspHead {
	if m != nil {
		return m.Head
	}
	return nil
}

type RankQueryByScoreReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Score                []uint64 `protobuf:"varint,2,rep,packed,name=Score,proto3" json:"Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankQueryByScoreReq) Reset()         { *m = RankQueryByScoreReq{} }
func (m *RankQueryByScoreReq) String() string { return proto.CompactTextString(m) }
func (*RankQueryByScoreReq) ProtoMessage()    {}
func (*RankQueryByScoreReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{16}
}

func (m *RankQueryByScoreReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankQueryByScoreReq.Unmarshal(m, b)
}
func (m *RankQueryByScoreReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankQueryByScoreReq.Marshal(b, m, deterministic)
}
func (m *RankQueryByScoreReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankQueryByScoreReq.Merge(m, src)
}
func (m *RankQueryByScoreReq) XXX_Size() int {
	return xxx_messageInfo_RankQueryByScoreReq.Size(m)
}
func (m *RankQueryByScoreReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankQueryByScoreReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankQueryByScoreReq proto.InternalMessageInfo

func (m *RankQueryByScoreReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankQueryByScoreReq) GetScore() []uint64 {
	if m != nil {
		return m.Score
	}
	return nil
}

type RankQueryByScoreRsp struct {
	Head                 *RspHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	Ranking              uint32   `protobuf:"varint,2,opt,name=Ranking,proto3" json:"Ranking,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankQueryByScoreRsp) Reset()         { *m = RankQueryByScoreRsp{} }
func (m *RankQueryByScoreRsp) String() string { return proto.CompactTextString(m) }
func (*RankQueryByScoreRsp) ProtoMessage()    {}
func (*RankQueryByScoreRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_729be9a7727521cd, []int{17}
}

func (m *RankQueryByScoreRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankQueryByScoreRsp.Unmarshal(m, b)
}
func (m *RankQueryByScoreRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankQueryByScoreRsp.Marshal(b, m, deterministic)
}
func (m *RankQueryByScoreRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankQueryByScoreRsp.Merge(m, src)
}
func (m *RankQueryByScoreRsp) XXX_Size() int {
	return xxx_messageInfo_RankQueryByScoreRsp.Size(m)
}
func (m *RankQueryByScoreRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankQueryByScoreRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RankQueryByScoreRsp proto.InternalMessageInfo

func (m *RankQueryByScoreRsp) GetHead() *RspHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *RankQueryByScoreRsp) GetRanking() uint32 {
	if m != nil {
		return m.Ranking
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.RANK_TYPE", RANK_TYPE_name, RANK_TYPE_value)
	proto.RegisterType((*RankInfo)(nil), "pb.RankInfo")
	proto.RegisterType((*RankUnitData)(nil), "pb.RankUnitData")
	proto.RegisterType((*UnitRankInfo)(nil), "pb.UnitRankInfo")
	proto.RegisterType((*RspHead)(nil), "pb.RspHead")
	proto.RegisterType((*RankInitReq)(nil), "pb.RankInitReq")
	proto.RegisterType((*RankInitRsp)(nil), "pb.RankInitRsp")
	proto.RegisterType((*RankDataUpdateReq)(nil), "pb.RankDataUpdateReq")
	proto.RegisterType((*RankDataUpdateRsp)(nil), "pb.RankDataUpdateRsp")
	proto.RegisterType((*RankQueryReq)(nil), "pb.RankQueryReq")
	proto.RegisterType((*RankQueryRsp)(nil), "pb.RankQueryRsp")
	proto.RegisterType((*RankTopQueryReq)(nil), "pb.RankTopQueryReq")
	proto.RegisterType((*RankTopQueryRsp)(nil), "pb.RankTopQueryRsp")
	proto.RegisterType((*RankDataDeleteReq)(nil), "pb.RankDataDeleteReq")
	proto.RegisterType((*RankDataDeleteRsp)(nil), "pb.RankDataDeleteRsp")
	proto.RegisterType((*RankCloseReq)(nil), "pb.RankCloseReq")
	proto.RegisterType((*RankCloseRsp)(nil), "pb.RankCloseRsp")
	proto.RegisterType((*RankQueryByScoreReq)(nil), "pb.RankQueryByScoreReq")
	proto.RegisterType((*RankQueryByScoreRsp)(nil), "pb.RankQueryByScoreRsp")
}

func init() { proto.RegisterFile("ranksrv.proto", fileDescriptor_729be9a7727521cd) }

var fileDescriptor_729be9a7727521cd = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0xb5, 0xa5, 0x7c, 0x5d, 0xa8, 0xe2, 0x68, 0x36, 0x8d, 0x0f, 0x5a, 0x47, 0x1f, 0x9a, 0x8d,
	0xb2, 0x09, 0xfa, 0xe0, 0x8b, 0x1a, 0x04, 0x36, 0x4b, 0x44, 0xc4, 0xb1, 0xc4, 0xec, 0x83, 0x1f,
	0x45, 0x46, 0x6d, 0x80, 0xce, 0xd0, 0x96, 0x4d, 0xf8, 0x79, 0xfe, 0x33, 0x33, 0xd3, 0xa1, 0x3b,
	0xab, 0x60, 0x9a, 0x7d, 0x9b, 0x3b, 0x73, 0xef, 0x39, 0xf7, 0x9e, 0x73, 0x5b, 0xb0, 0xe3, 0x20,
	0x5a, 0x24, 0xf1, 0x45, 0x9b, 0xc7, 0x2c, 0x65, 0xc8, 0xe4, 0x33, 0xfc, 0xdb, 0x80, 0x1a, 0x09,
	0xa2, 0xc5, 0x30, 0xfa, 0xc1, 0x10, 0x02, 0x6b, 0x1c, 0xac, 0xa8, 0x63, 0xb8, 0x86, 0x57, 0x27,
	0xf2, 0x8c, 0x1e, 0x82, 0xe5, 0x6f, 0x39, 0x75, 0x4c, 0xd7, 0xf0, 0x6e, 0x76, 0xec, 0x36, 0x9f,
	0xb5, 0x49, 0x77, 0xfc, 0xf6, 0xab, 0x7f, 0x3e, 0x19, 0x10, 0xf9, 0x84, 0x8e, 0xa0, 0xe2, 0x33,
	0x3e, 0xde, 0xac, 0x9c, 0x92, 0x6b, 0x78, 0x36, 0x51, 0x11, 0x7a, 0x0c, 0xf6, 0xe9, 0x72, 0x93,
	0xfc, 0x1a, 0x46, 0x29, 0x8d, 0x2f, 0x82, 0xa5, 0x63, 0xc9, 0xe7, 0xab, 0x97, 0x22, 0xeb, 0xe3,
	0x77, 0x16, 0xd3, 0x77, 0x61, 0x34, 0x0a, 0x57, 0x61, 0xea, 0x94, 0x5d, 0xc3, 0xb3, 0xc8, 0xd5,
	0x4b, 0x74, 0x0f, 0x6a, 0xf2, 0x42, 0xb0, 0x54, 0x24, 0x4c, 0x1e, 0xe3, 0x6f, 0xd0, 0x14, 0x23,
	0x4c, 0xa3, 0x30, 0xed, 0x07, 0x69, 0x20, 0x72, 0xa7, 0x51, 0xb8, 0xde, 0xd0, 0x61, 0x5f, 0x8d,
	0x92, 0xc7, 0xe8, 0x2e, 0x94, 0x65, 0x9d, 0x63, 0xba, 0x25, 0xcf, 0x22, 0x59, 0x80, 0xee, 0x03,
	0x4c, 0xf9, 0x3c, 0x48, 0xa9, 0x1f, 0xae, 0xa8, 0x9a, 0x42, 0xbb, 0xc1, 0x2f, 0xa0, 0x29, 0xd0,
	0x75, 0xa1, 0xc4, 0x59, 0xa2, 0xdb, 0x44, 0x9e, 0xf7, 0x23, 0xe3, 0x47, 0x50, 0x25, 0x09, 0x3f,
	0xa3, 0xc1, 0x1c, 0x39, 0x50, 0x25, 0x34, 0xed, 0xb1, 0x39, 0x55, 0x75, 0xbb, 0x10, 0x9f, 0x40,
	0x23, 0x83, 0x0e, 0x53, 0x42, 0xd7, 0xc8, 0x05, 0x4b, 0xb0, 0xc8, 0xac, 0x46, 0xa7, 0x29, 0x25,
	0x57, 0xcc, 0x44, 0xbe, 0xe0, 0xb6, 0x56, 0x90, 0x70, 0xf4, 0x00, 0x2c, 0xc1, 0xa0, 0x0a, 0x1a,
	0xb2, 0x20, 0x23, 0x25, 0xf2, 0x01, 0x9f, 0xc3, 0x6d, 0x91, 0x2f, 0xd4, 0xc9, 0xa6, 0x12, 0x34,
	0xfb, 0xdc, 0xd6, 0xa5, 0x33, 0x0f, 0x49, 0x57, 0xd2, 0x07, 0x7c, 0xfe, 0x0f, 0x74, 0x91, 0x86,
	0x5e, 0x65, 0x96, 0x7d, 0xd8, 0xd0, 0x78, 0x7b, 0x8d, 0x5e, 0xf0, 0x67, 0xbd, 0xbe, 0x00, 0x21,
	0x7a, 0x72, 0xb9, 0xe6, 0x12, 0xac, 0xd1, 0x69, 0x89, 0x24, 0xdd, 0x55, 0x92, 0x67, 0xe0, 0x97,
	0x70, 0x4b, 0x9c, 0x7d, 0xc6, 0xff, 0xdb, 0xe1, 0x11, 0x54, 0x08, 0x5d, 0x8b, 0x95, 0x34, 0xb3,
	0xc5, 0xcf, 0x22, 0xfc, 0xe5, 0xaf, 0xf2, 0x22, 0x0d, 0x1e, 0x43, 0xd5, 0x67, 0x5c, 0x6e, 0x95,
	0x58, 0x20, 0xd5, 0x9f, 0xbe, 0xd7, 0x64, 0x97, 0x80, 0x7b, 0x97, 0x9a, 0xf7, 0xe9, 0x92, 0x5e,
	0xcb, 0x4e, 0xdd, 0x38, 0x05, 0x52, 0xc4, 0x38, 0x9c, 0x09, 0xdf, 0x5b, 0xb2, 0xe4, 0x10, 0x2b,
	0x3e, 0xd1, 0x73, 0x8a, 0x80, 0xbe, 0x86, 0x3b, 0xb9, 0x9b, 0x6f, 0xb6, 0x72, 0xaf, 0x0e, 0x4d,
	0xb4, 0xff, 0x2b, 0x9b, 0xec, 0x01, 0x28, 0x22, 0xba, 0xf8, 0x24, 0x83, 0x68, 0x11, 0x46, 0x3f,
	0x95, 0x83, 0xbb, 0xf0, 0xf8, 0x29, 0xd4, 0xf3, 0xdf, 0x1c, 0x6a, 0x40, 0xb5, 0x3f, 0x38, 0xed,
	0x4e, 0x47, 0x7e, 0xeb, 0x06, 0xaa, 0x81, 0x35, 0xe9, 0x12, 0xbf, 0x65, 0xa0, 0x3a, 0x94, 0x3f,
	0x9d, 0xbd, 0x1f, 0x0d, 0x5a, 0xe6, 0xac, 0x22, 0xff, 0xa8, 0xcf, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x79, 0x0f, 0xf8, 0xe8, 0x62, 0x05, 0x00, 0x00,
}
