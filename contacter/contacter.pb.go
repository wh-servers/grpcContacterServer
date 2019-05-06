// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contacter.proto

package contacter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SourceType int32

const (
	SourceType_DB    SourceType = 1
	SourceType_MAP   SourceType = 2
	SourceType_REDIS SourceType = 3
)

var SourceType_name = map[int32]string{
	1: "DB",
	2: "MAP",
	3: "REDIS",
}

var SourceType_value = map[string]int32{
	"DB":    1,
	"MAP":   2,
	"REDIS": 3,
}

func (x SourceType) Enum() *SourceType {
	p := new(SourceType)
	*p = x
	return p
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}

func (x *SourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SourceType_value, data, "SourceType")
	if err != nil {
		return err
	}
	*x = SourceType(value)
	return nil
}

func (SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{0}
}

type LoadType int32

const (
	LoadType_ID        LoadType = 1
	LoadType_NAME      LoadType = 2
	LoadType_COUNTRY   LoadType = 3
	LoadType_STATUS    LoadType = 4
	LoadType_HEADPHONE LoadType = 5
	LoadType_EMAIL     LoadType = 6
)

var LoadType_name = map[int32]string{
	1: "ID",
	2: "NAME",
	3: "COUNTRY",
	4: "STATUS",
	5: "HEADPHONE",
	6: "EMAIL",
}

var LoadType_value = map[string]int32{
	"ID":        1,
	"NAME":      2,
	"COUNTRY":   3,
	"STATUS":    4,
	"HEADPHONE": 5,
	"EMAIL":     6,
}

func (x LoadType) Enum() *LoadType {
	p := new(LoadType)
	*p = x
	return p
}

func (x LoadType) String() string {
	return proto.EnumName(LoadType_name, int32(x))
}

func (x *LoadType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LoadType_value, data, "LoadType")
	if err != nil {
		return err
	}
	*x = LoadType(value)
	return nil
}

func (LoadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{1}
}

//100000～199999
type CONST_ErrorCode int32

const (
	CONST_ERROR_WRONG_ID        CONST_ErrorCode = 100001
	CONST_ERROR_INVALID_COUNTRY CONST_ErrorCode = 100002
	CONST_ERROR_DB_EXCEPTION    CONST_ErrorCode = 100003
)

var CONST_ErrorCode_name = map[int32]string{
	100001: "ERROR_WRONG_ID",
	100002: "ERROR_INVALID_COUNTRY",
	100003: "ERROR_DB_EXCEPTION",
}

var CONST_ErrorCode_value = map[string]int32{
	"ERROR_WRONG_ID":        100001,
	"ERROR_INVALID_COUNTRY": 100002,
	"ERROR_DB_EXCEPTION":    100003,
}

func (x CONST_ErrorCode) Enum() *CONST_ErrorCode {
	p := new(CONST_ErrorCode)
	*p = x
	return p
}

func (x CONST_ErrorCode) String() string {
	return proto.EnumName(CONST_ErrorCode_name, int32(x))
}

func (x *CONST_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CONST_ErrorCode_value, data, "CONST_ErrorCode")
	if err != nil {
		return err
	}
	*x = CONST_ErrorCode(value)
	return nil
}

func (CONST_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{0, 0}
}

type CONST_ContacterStatus int32

const (
	CONST_NORMAL  CONST_ContacterStatus = 1
	CONST_DELETE  CONST_ContacterStatus = 2
	CONST_DISABLE CONST_ContacterStatus = 3
)

var CONST_ContacterStatus_name = map[int32]string{
	1: "NORMAL",
	2: "DELETE",
	3: "DISABLE",
}

var CONST_ContacterStatus_value = map[string]int32{
	"NORMAL":  1,
	"DELETE":  2,
	"DISABLE": 3,
}

func (x CONST_ContacterStatus) Enum() *CONST_ContacterStatus {
	p := new(CONST_ContacterStatus)
	*p = x
	return p
}

func (x CONST_ContacterStatus) String() string {
	return proto.EnumName(CONST_ContacterStatus_name, int32(x))
}

func (x *CONST_ContacterStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CONST_ContacterStatus_value, data, "CONST_ContacterStatus")
	if err != nil {
		return err
	}
	*x = CONST_ContacterStatus(value)
	return nil
}

func (CONST_ContacterStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{0, 1}
}

type CONST struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CONST) Reset()         { *m = CONST{} }
func (m *CONST) String() string { return proto.CompactTextString(m) }
func (*CONST) ProtoMessage()    {}
func (*CONST) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{0}
}

func (m *CONST) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CONST.Unmarshal(m, b)
}
func (m *CONST) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CONST.Marshal(b, m, deterministic)
}
func (m *CONST) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CONST.Merge(m, src)
}
func (m *CONST) XXX_Size() int {
	return xxx_messageInfo_CONST.Size(m)
}
func (m *CONST) XXX_DiscardUnknown() {
	xxx_messageInfo_CONST.DiscardUnknown(m)
}

var xxx_messageInfo_CONST proto.InternalMessageInfo

type Contacter struct {
	Id                   *int32            `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	Name                 *string           `protobuf:"bytes,2,req,name=Name" json:"Name,omitempty"`
	Country              *string           `protobuf:"bytes,3,opt,name=Country" json:"Country,omitempty"`
	Status               *int32            `protobuf:"varint,4,req,name=Status" json:"Status,omitempty"`
	Image                *string           `protobuf:"bytes,5,opt,name=Image" json:"Image,omitempty"`
	HeadPhone            *int32            `protobuf:"varint,6,opt,name=HeadPhone" json:"HeadPhone,omitempty"`
	HomePhone            *int32            `protobuf:"varint,7,opt,name=HomePhone" json:"HomePhone,omitempty"`
	Email                *string           `protobuf:"bytes,8,opt,name=Email" json:"Email,omitempty"`
	Ctime                *int32            `protobuf:"varint,9,req,name=Ctime" json:"Ctime,omitempty"`
	Mtime                *int32            `protobuf:"varint,10,req,name=Mtime" json:"Mtime,omitempty"`
	Extinfo              *ContacterExtinfo `protobuf:"bytes,11,opt,name=Extinfo" json:"Extinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Contacter) Reset()         { *m = Contacter{} }
func (m *Contacter) String() string { return proto.CompactTextString(m) }
func (*Contacter) ProtoMessage()    {}
func (*Contacter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{1}
}

func (m *Contacter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contacter.Unmarshal(m, b)
}
func (m *Contacter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contacter.Marshal(b, m, deterministic)
}
func (m *Contacter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contacter.Merge(m, src)
}
func (m *Contacter) XXX_Size() int {
	return xxx_messageInfo_Contacter.Size(m)
}
func (m *Contacter) XXX_DiscardUnknown() {
	xxx_messageInfo_Contacter.DiscardUnknown(m)
}

var xxx_messageInfo_Contacter proto.InternalMessageInfo

func (m *Contacter) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Contacter) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Contacter) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *Contacter) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *Contacter) GetImage() string {
	if m != nil && m.Image != nil {
		return *m.Image
	}
	return ""
}

func (m *Contacter) GetHeadPhone() int32 {
	if m != nil && m.HeadPhone != nil {
		return *m.HeadPhone
	}
	return 0
}

func (m *Contacter) GetHomePhone() int32 {
	if m != nil && m.HomePhone != nil {
		return *m.HomePhone
	}
	return 0
}

func (m *Contacter) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *Contacter) GetCtime() int32 {
	if m != nil && m.Ctime != nil {
		return *m.Ctime
	}
	return 0
}

func (m *Contacter) GetMtime() int32 {
	if m != nil && m.Mtime != nil {
		return *m.Mtime
	}
	return 0
}

func (m *Contacter) GetExtinfo() *ContacterExtinfo {
	if m != nil {
		return m.Extinfo
	}
	return nil
}

type ContacterExtinfo struct {
	Company              *string  `protobuf:"bytes,1,opt,name=Company" json:"Company,omitempty"`
	Remark               []byte   `protobuf:"bytes,2,opt,name=Remark" json:"Remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContacterExtinfo) Reset()         { *m = ContacterExtinfo{} }
func (m *ContacterExtinfo) String() string { return proto.CompactTextString(m) }
func (*ContacterExtinfo) ProtoMessage()    {}
func (*ContacterExtinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{2}
}

func (m *ContacterExtinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContacterExtinfo.Unmarshal(m, b)
}
func (m *ContacterExtinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContacterExtinfo.Marshal(b, m, deterministic)
}
func (m *ContacterExtinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContacterExtinfo.Merge(m, src)
}
func (m *ContacterExtinfo) XXX_Size() int {
	return xxx_messageInfo_ContacterExtinfo.Size(m)
}
func (m *ContacterExtinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContacterExtinfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContacterExtinfo proto.InternalMessageInfo

func (m *ContacterExtinfo) GetCompany() string {
	if m != nil && m.Company != nil {
		return *m.Company
	}
	return ""
}

func (m *ContacterExtinfo) GetRemark() []byte {
	if m != nil {
		return m.Remark
	}
	return nil
}

type GetInfoRequest struct {
	Contacter            *Contacter  `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	Source               *SourceType `protobuf:"varint,2,opt,name=Source,enum=contacter.SourceType" json:"Source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{3}
}

func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

func (m *GetInfoRequest) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

func (m *GetInfoRequest) GetSource() SourceType {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return SourceType_DB
}

type GetInfoResponse struct {
	Contacter            *Contacter `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	ErrMsg               *string    `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{4}
}

func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(m, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

func (m *GetInfoResponse) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

type InsertRequest struct {
	Contacter            *Contacter `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InsertRequest) Reset()         { *m = InsertRequest{} }
func (m *InsertRequest) String() string { return proto.CompactTextString(m) }
func (*InsertRequest) ProtoMessage()    {}
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{5}
}

func (m *InsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertRequest.Unmarshal(m, b)
}
func (m *InsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertRequest.Marshal(b, m, deterministic)
}
func (m *InsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertRequest.Merge(m, src)
}
func (m *InsertRequest) XXX_Size() int {
	return xxx_messageInfo_InsertRequest.Size(m)
}
func (m *InsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertRequest proto.InternalMessageInfo

func (m *InsertRequest) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

type InsertResponse struct {
	Contacter            *Contacter `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	ErrMsg               *string    `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InsertResponse) Reset()         { *m = InsertResponse{} }
func (m *InsertResponse) String() string { return proto.CompactTextString(m) }
func (*InsertResponse) ProtoMessage()    {}
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{6}
}

func (m *InsertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertResponse.Unmarshal(m, b)
}
func (m *InsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertResponse.Marshal(b, m, deterministic)
}
func (m *InsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertResponse.Merge(m, src)
}
func (m *InsertResponse) XXX_Size() int {
	return xxx_messageInfo_InsertResponse.Size(m)
}
func (m *InsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InsertResponse proto.InternalMessageInfo

func (m *InsertResponse) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

func (m *InsertResponse) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

type DeleteRequest struct {
	Id                   *int32   `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Operator             *string  `protobuf:"bytes,2,opt,name=operator" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DeleteRequest) GetOperator() string {
	if m != nil && m.Operator != nil {
		return *m.Operator
	}
	return ""
}

type DeleteResponse struct {
	Contacter            *Contacter `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	ErrMsg               *string    `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

func (m *DeleteResponse) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

type ModifyRequest struct {
	Contacter            *Contacter `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ModifyRequest) Reset()         { *m = ModifyRequest{} }
func (m *ModifyRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyRequest) ProtoMessage()    {}
func (*ModifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{9}
}

func (m *ModifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyRequest.Unmarshal(m, b)
}
func (m *ModifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyRequest.Marshal(b, m, deterministic)
}
func (m *ModifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyRequest.Merge(m, src)
}
func (m *ModifyRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyRequest.Size(m)
}
func (m *ModifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyRequest proto.InternalMessageInfo

func (m *ModifyRequest) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

type ModifyResponse struct {
	Contacter            *Contacter `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	ErrMsg               *string    `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ModifyResponse) Reset()         { *m = ModifyResponse{} }
func (m *ModifyResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyResponse) ProtoMessage()    {}
func (*ModifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{10}
}

func (m *ModifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyResponse.Unmarshal(m, b)
}
func (m *ModifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyResponse.Marshal(b, m, deterministic)
}
func (m *ModifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyResponse.Merge(m, src)
}
func (m *ModifyResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyResponse.Size(m)
}
func (m *ModifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyResponse proto.InternalMessageInfo

func (m *ModifyResponse) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

func (m *ModifyResponse) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

type ReloadMapRequest struct {
	LdType               *LoadType `protobuf:"varint,1,opt,name=ldType,enum=contacter.LoadType" json:"ldType,omitempty"`
	TypeId               *string   `protobuf:"bytes,2,opt,name=TypeId" json:"TypeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReloadMapRequest) Reset()         { *m = ReloadMapRequest{} }
func (m *ReloadMapRequest) String() string { return proto.CompactTextString(m) }
func (*ReloadMapRequest) ProtoMessage()    {}
func (*ReloadMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{11}
}

func (m *ReloadMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadMapRequest.Unmarshal(m, b)
}
func (m *ReloadMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadMapRequest.Marshal(b, m, deterministic)
}
func (m *ReloadMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadMapRequest.Merge(m, src)
}
func (m *ReloadMapRequest) XXX_Size() int {
	return xxx_messageInfo_ReloadMapRequest.Size(m)
}
func (m *ReloadMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadMapRequest proto.InternalMessageInfo

func (m *ReloadMapRequest) GetLdType() LoadType {
	if m != nil && m.LdType != nil {
		return *m.LdType
	}
	return LoadType_ID
}

func (m *ReloadMapRequest) GetTypeId() string {
	if m != nil && m.TypeId != nil {
		return *m.TypeId
	}
	return ""
}

type ReloadMapResponse struct {
	Contacter            *Contacter `protobuf:"bytes,1,opt,name=contacter" json:"contacter,omitempty"`
	ErrMsg               *string    `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReloadMapResponse) Reset()         { *m = ReloadMapResponse{} }
func (m *ReloadMapResponse) String() string { return proto.CompactTextString(m) }
func (*ReloadMapResponse) ProtoMessage()    {}
func (*ReloadMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5acf61d2d998b8dd, []int{12}
}

func (m *ReloadMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadMapResponse.Unmarshal(m, b)
}
func (m *ReloadMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadMapResponse.Marshal(b, m, deterministic)
}
func (m *ReloadMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadMapResponse.Merge(m, src)
}
func (m *ReloadMapResponse) XXX_Size() int {
	return xxx_messageInfo_ReloadMapResponse.Size(m)
}
func (m *ReloadMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadMapResponse proto.InternalMessageInfo

func (m *ReloadMapResponse) GetContacter() *Contacter {
	if m != nil {
		return m.Contacter
	}
	return nil
}

func (m *ReloadMapResponse) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("contacter.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("contacter.LoadType", LoadType_name, LoadType_value)
	proto.RegisterEnum("contacter.CONST_ErrorCode", CONST_ErrorCode_name, CONST_ErrorCode_value)
	proto.RegisterEnum("contacter.CONST_ContacterStatus", CONST_ContacterStatus_name, CONST_ContacterStatus_value)
	proto.RegisterType((*CONST)(nil), "contacter.CONST")
	proto.RegisterType((*Contacter)(nil), "contacter.Contacter")
	proto.RegisterType((*ContacterExtinfo)(nil), "contacter.ContacterExtinfo")
	proto.RegisterType((*GetInfoRequest)(nil), "contacter.GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "contacter.GetInfoResponse")
	proto.RegisterType((*InsertRequest)(nil), "contacter.InsertRequest")
	proto.RegisterType((*InsertResponse)(nil), "contacter.InsertResponse")
	proto.RegisterType((*DeleteRequest)(nil), "contacter.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "contacter.DeleteResponse")
	proto.RegisterType((*ModifyRequest)(nil), "contacter.ModifyRequest")
	proto.RegisterType((*ModifyResponse)(nil), "contacter.ModifyResponse")
	proto.RegisterType((*ReloadMapRequest)(nil), "contacter.ReloadMapRequest")
	proto.RegisterType((*ReloadMapResponse)(nil), "contacter.ReloadMapResponse")
}

func init() { proto.RegisterFile("contacter.proto", fileDescriptor_5acf61d2d998b8dd) }

var fileDescriptor_5acf61d2d998b8dd = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x6f, 0xda, 0x48,
	0x14, 0x8f, 0x0d, 0x18, 0xfc, 0x48, 0xc0, 0x3b, 0x9b, 0xec, 0x7a, 0x49, 0x0e, 0xc8, 0x27, 0x94,
	0xd5, 0xe6, 0x80, 0xb4, 0x7b, 0xd9, 0xc3, 0xca, 0xd8, 0x56, 0xb0, 0x84, 0x6d, 0x34, 0x90, 0x4d,
	0x7a, 0x68, 0xa9, 0x15, 0x4f, 0x52, 0x54, 0xcc, 0x50, 0xdb, 0x48, 0xe5, 0x43, 0xe4, 0x03, 0xf4,
	0xd8, 0xf6, 0xda, 0x0f, 0xd6, 0x8f, 0x51, 0xd9, 0xe3, 0x31, 0x38, 0xe2, 0xd4, 0x88, 0x93, 0xfd,
	0xfe, 0xcc, 0xef, 0xfd, 0xe6, 0xcd, 0xef, 0x3d, 0x68, 0xdf, 0xd3, 0x65, 0xe2, 0xdf, 0x27, 0x24,
	0xba, 0x5a, 0x45, 0x34, 0xa1, 0x48, 0x2e, 0x1c, 0xda, 0x27, 0x01, 0x6a, 0x86, 0xe7, 0x4e, 0xa6,
	0xda, 0x1d, 0xc8, 0x56, 0x14, 0xd1, 0xc8, 0xa0, 0x01, 0x41, 0xa7, 0xd0, 0xb2, 0x30, 0xf6, 0xf0,
	0xec, 0x16, 0x7b, 0xee, 0xf5, 0xcc, 0x36, 0x95, 0xcf, 0x4f, 0x12, 0x3a, 0x87, 0x33, 0xe6, 0xb5,
	0xdd, 0xff, 0xf5, 0x91, 0x6d, 0xce, 0x0c, 0xef, 0xc6, 0x9d, 0xe2, 0x57, 0xca, 0x97, 0x27, 0x09,
	0xa9, 0x80, 0x58, 0xd0, 0x1c, 0xcc, 0xac, 0x3b, 0xc3, 0x1a, 0x4f, 0x6d, 0xcf, 0x55, 0xbe, 0x3e,
	0x49, 0xda, 0x3f, 0xd0, 0x36, 0x78, 0xc1, 0x49, 0xe2, 0x27, 0xeb, 0x18, 0x01, 0x48, 0xae, 0x87,
	0x1d, 0x7d, 0xa4, 0x08, 0xe9, 0xbf, 0x69, 0x8d, 0xac, 0xa9, 0xa5, 0x88, 0xa8, 0x09, 0x75, 0xd3,
	0x9e, 0xe8, 0x83, 0x91, 0xa5, 0x54, 0xb4, 0x6f, 0x22, 0xc8, 0xc5, 0x41, 0xd4, 0x02, 0xd1, 0x0e,
	0x54, 0xa1, 0x2b, 0xf6, 0x6a, 0x58, 0xb4, 0x03, 0x84, 0xa0, 0xea, 0xfa, 0x21, 0x51, 0xc5, 0xae,
	0xd8, 0x93, 0x71, 0xf6, 0x8f, 0x54, 0xa8, 0x1b, 0x74, 0xbd, 0x4c, 0xa2, 0x8d, 0x5a, 0xe9, 0x0a,
	0x3d, 0x19, 0x73, 0x13, 0xfd, 0x06, 0x12, 0x2b, 0xad, 0x56, 0x33, 0x84, 0xdc, 0x42, 0xa7, 0x50,
	0xb3, 0x43, 0xff, 0x91, 0xa8, 0xb5, 0x2c, 0x9f, 0x19, 0xe8, 0x02, 0xe4, 0x21, 0xf1, 0x83, 0xf1,
	0x3b, 0xba, 0x24, 0xaa, 0xd4, 0x15, 0x7a, 0x35, 0xbc, 0x75, 0x64, 0x51, 0x1a, 0x12, 0x16, 0xad,
	0xe7, 0x51, 0xee, 0x48, 0x11, 0xad, 0xd0, 0x9f, 0x2f, 0xd4, 0x06, 0x43, 0xcc, 0x8c, 0xd4, 0x6b,
	0x24, 0xf3, 0x90, 0xa8, 0x72, 0x56, 0x9e, 0x19, 0xa9, 0xd7, 0xc9, 0xbc, 0xc0, 0xbc, 0x99, 0x81,
	0xfe, 0x86, 0xba, 0xf5, 0x31, 0x99, 0x2f, 0x1f, 0xa8, 0xda, 0xec, 0x0a, 0xbd, 0x66, 0xff, 0xfc,
	0x6a, 0xfb, 0x82, 0x45, 0x43, 0xf2, 0x14, 0xcc, 0x73, 0x35, 0x13, 0x94, 0xe7, 0x41, 0xd6, 0x90,
	0x70, 0xe5, 0x2f, 0x37, 0xaa, 0xc0, 0x1b, 0x92, 0x99, 0x69, 0x43, 0x30, 0x09, 0xfd, 0xe8, 0xbd,
	0x2a, 0x76, 0x85, 0xde, 0x31, 0xce, 0x2d, 0x2d, 0x86, 0xd6, 0x35, 0x49, 0xec, 0x14, 0x99, 0x7c,
	0x58, 0x93, 0x38, 0x41, 0x7d, 0xd8, 0xea, 0x25, 0x43, 0x69, 0xf6, 0x4f, 0xf7, 0x11, 0xc2, 0xdb,
	0x34, 0xf4, 0x17, 0x48, 0x13, 0xba, 0x8e, 0xee, 0x49, 0x86, 0xde, 0xea, 0x9f, 0xed, 0x1c, 0x60,
	0x81, 0xe9, 0x66, 0x45, 0x70, 0x9e, 0xa4, 0xbd, 0x81, 0x76, 0x51, 0x34, 0x5e, 0xd1, 0x65, 0x4c,
	0x7e, 0xaa, 0xea, 0xef, 0x50, 0x27, 0x51, 0x34, 0x0b, 0xe3, 0xc7, 0xac, 0xac, 0x8c, 0x25, 0x12,
	0x45, 0x4e, 0xfc, 0xa8, 0x19, 0x70, 0x62, 0x2f, 0x63, 0x12, 0x25, 0x2f, 0xb8, 0x93, 0xf6, 0x1a,
	0x5a, 0x1c, 0xe4, 0x10, 0x1c, 0xff, 0x85, 0x13, 0x93, 0x2c, 0x48, 0x42, 0x38, 0x47, 0x2e, 0x78,
	0x21, 0x17, 0x7c, 0x07, 0x1a, 0x74, 0x45, 0x22, 0x3f, 0xa1, 0x51, 0x7e, 0xb4, 0xb0, 0x53, 0x6e,
	0xfc, 0xf0, 0x81, 0xfa, 0xe7, 0xd0, 0x60, 0xfe, 0xb0, 0x79, 0x61, 0xff, 0x38, 0xc8, 0x21, 0x38,
	0xde, 0x82, 0x82, 0xc9, 0x82, 0xfa, 0x81, 0xe3, 0xaf, 0x38, 0xcd, 0x3f, 0x41, 0x5a, 0x04, 0xa9,
	0xd2, 0x32, 0xf4, 0x56, 0xff, 0xd7, 0x1d, 0xf4, 0x11, 0xf5, 0x03, 0x26, 0x42, 0x96, 0x92, 0x4e,
	0x44, 0xfa, 0xb5, 0x03, 0x0e, 0xcc, 0x2c, 0xed, 0x2d, 0xfc, 0xb2, 0x03, 0x7c, 0x00, 0xea, 0x97,
	0x3d, 0x80, 0xed, 0x50, 0x20, 0x09, 0x44, 0x73, 0xa0, 0x08, 0xa8, 0x0e, 0x15, 0x47, 0x1f, 0x2b,
	0x22, 0x92, 0xa1, 0x86, 0x2d, 0xd3, 0x9e, 0x28, 0x95, 0x4b, 0x0f, 0x1a, 0x9c, 0x77, 0x9a, 0x67,
	0x9b, 0x8a, 0x80, 0x1a, 0x50, 0x75, 0x75, 0x27, 0xdf, 0x9e, 0x7c, 0x23, 0x57, 0xd2, 0xb5, 0x3a,
	0x99, 0xea, 0xd3, 0x9b, 0x89, 0x52, 0x45, 0x27, 0x20, 0x0f, 0x2d, 0xdd, 0x1c, 0x0f, 0x3d, 0xd7,
	0x52, 0x6a, 0x29, 0xa0, 0xe5, 0xe8, 0xf6, 0x48, 0x91, 0xfa, 0xdf, 0x45, 0x38, 0x2e, 0xc8, 0xea,
	0x63, 0x1b, 0xfd, 0x07, 0x12, 0x53, 0x39, 0x52, 0x77, 0xee, 0x53, 0x9a, 0x9e, 0xce, 0x1f, 0x7b,
	0x22, 0xac, 0x2f, 0xda, 0x11, 0x1a, 0x40, 0x3d, 0x9f, 0x65, 0xb4, 0x9b, 0x57, 0x5e, 0x2a, 0x9d,
	0xce, 0xbe, 0x50, 0x81, 0x31, 0x84, 0x36, 0x93, 0xca, 0x76, 0xfd, 0xef, 0xb2, 0x29, 0x69, 0xb1,
	0xc4, 0xa6, 0x2c, 0x30, 0xed, 0x28, 0xbd, 0x0e, 0x1b, 0x8c, 0x12, 0x40, 0x69, 0xd0, 0x4a, 0x00,
	0xe5, 0x29, 0xca, 0xa8, 0xc8, 0xc5, 0xeb, 0xa3, 0xdd, 0x45, 0xfc, 0x5c, 0x6c, 0x9d, 0x8b, 0xfd,
	0x41, 0x8e, 0xf4, 0x23, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x4c, 0xb4, 0x0d, 0x87, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContacterAPIClient is the client API for ContacterAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContacterAPIClient interface {
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	ModifyContacter(ctx context.Context, in *ModifyRequest, opts ...grpc.CallOption) (*ModifyResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ReloadMap(ctx context.Context, in *ReloadMapRequest, opts ...grpc.CallOption) (*ReloadMapResponse, error)
}

type contacterAPIClient struct {
	cc *grpc.ClientConn
}

func NewContacterAPIClient(cc *grpc.ClientConn) ContacterAPIClient {
	return &contacterAPIClient{cc}
}

func (c *contacterAPIClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/contacter.ContacterAPI/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacterAPIClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/contacter.ContacterAPI/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacterAPIClient) ModifyContacter(ctx context.Context, in *ModifyRequest, opts ...grpc.CallOption) (*ModifyResponse, error) {
	out := new(ModifyResponse)
	err := c.cc.Invoke(ctx, "/contacter.ContacterAPI/ModifyContacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacterAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/contacter.ContacterAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacterAPIClient) ReloadMap(ctx context.Context, in *ReloadMapRequest, opts ...grpc.CallOption) (*ReloadMapResponse, error) {
	out := new(ReloadMapResponse)
	err := c.cc.Invoke(ctx, "/contacter.ContacterAPI/ReloadMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContacterAPIServer is the server API for ContacterAPI service.
type ContacterAPIServer interface {
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	ModifyContacter(context.Context, *ModifyRequest) (*ModifyResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	ReloadMap(context.Context, *ReloadMapRequest) (*ReloadMapResponse, error)
}

// UnimplementedContacterAPIServer can be embedded to have forward compatible implementations.
type UnimplementedContacterAPIServer struct {
}

func (*UnimplementedContacterAPIServer) Insert(ctx context.Context, req *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedContacterAPIServer) GetInfo(ctx context.Context, req *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (*UnimplementedContacterAPIServer) ModifyContacter(ctx context.Context, req *ModifyRequest) (*ModifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyContacter not implemented")
}
func (*UnimplementedContacterAPIServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedContacterAPIServer) ReloadMap(ctx context.Context, req *ReloadMapRequest) (*ReloadMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadMap not implemented")
}

func RegisterContacterAPIServer(s *grpc.Server, srv ContacterAPIServer) {
	s.RegisterService(&_ContacterAPI_serviceDesc, srv)
}

func _ContacterAPI_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContacterAPIServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacter.ContacterAPI/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContacterAPIServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContacterAPI_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContacterAPIServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacter.ContacterAPI/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContacterAPIServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContacterAPI_ModifyContacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContacterAPIServer).ModifyContacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacter.ContacterAPI/ModifyContacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContacterAPIServer).ModifyContacter(ctx, req.(*ModifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContacterAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContacterAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacter.ContacterAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContacterAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContacterAPI_ReloadMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContacterAPIServer).ReloadMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacter.ContacterAPI/ReloadMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContacterAPIServer).ReloadMap(ctx, req.(*ReloadMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContacterAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contacter.ContacterAPI",
	HandlerType: (*ContacterAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _ContacterAPI_Insert_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _ContacterAPI_GetInfo_Handler,
		},
		{
			MethodName: "ModifyContacter",
			Handler:    _ContacterAPI_ModifyContacter_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ContacterAPI_Delete_Handler,
		},
		{
			MethodName: "ReloadMap",
			Handler:    _ContacterAPI_ReloadMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contacter.proto",
}
