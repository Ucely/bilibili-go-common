// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/Banner.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BannerGetNewBannerReq struct {
	// 业务id，首页为0，web主站hove为1
	Platform int64 `protobuf:"varint,1,opt,name=platform,proto3" json:"platform"`
	// 第几帧，0表示取全部
	Position int64 `protobuf:"varint,2,opt,name=position,proto3" json:"position"`
	// 平台
	UserPlatform string `protobuf:"bytes,3,opt,name=userPlatform,proto3" json:"userPlatform"`
	// 设备
	UserDevice string `protobuf:"bytes,4,opt,name=userDevice,proto3" json:"userDevice"`
	// 版本号
	Build int64 `protobuf:"varint,5,opt,name=build,proto3" json:"build"`
}

func (m *BannerGetNewBannerReq) Reset()         { *m = BannerGetNewBannerReq{} }
func (m *BannerGetNewBannerReq) String() string { return proto.CompactTextString(m) }
func (*BannerGetNewBannerReq) ProtoMessage()    {}
func (*BannerGetNewBannerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_Banner_dd0f870b0c73e0e6, []int{0}
}
func (m *BannerGetNewBannerReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BannerGetNewBannerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BannerGetNewBannerReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BannerGetNewBannerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerGetNewBannerReq.Merge(dst, src)
}
func (m *BannerGetNewBannerReq) XXX_Size() int {
	return m.Size()
}
func (m *BannerGetNewBannerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerGetNewBannerReq.DiscardUnknown(m)
}

var xxx_messageInfo_BannerGetNewBannerReq proto.InternalMessageInfo

func (m *BannerGetNewBannerReq) GetPlatform() int64 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *BannerGetNewBannerReq) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *BannerGetNewBannerReq) GetUserPlatform() string {
	if m != nil {
		return m.UserPlatform
	}
	return ""
}

func (m *BannerGetNewBannerReq) GetUserDevice() string {
	if m != nil {
		return m.UserDevice
	}
	return ""
}

func (m *BannerGetNewBannerReq) GetBuild() int64 {
	if m != nil {
		return m.Build
	}
	return 0
}

type BannerGetNewBannerResp struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data []*BannerGetNewBannerResp_NewBanner `protobuf:"bytes,3,rep,name=data" json:"data"`
}

func (m *BannerGetNewBannerResp) Reset()         { *m = BannerGetNewBannerResp{} }
func (m *BannerGetNewBannerResp) String() string { return proto.CompactTextString(m) }
func (*BannerGetNewBannerResp) ProtoMessage()    {}
func (*BannerGetNewBannerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_Banner_dd0f870b0c73e0e6, []int{1}
}
func (m *BannerGetNewBannerResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BannerGetNewBannerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BannerGetNewBannerResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BannerGetNewBannerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerGetNewBannerResp.Merge(dst, src)
}
func (m *BannerGetNewBannerResp) XXX_Size() int {
	return m.Size()
}
func (m *BannerGetNewBannerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerGetNewBannerResp.DiscardUnknown(m)
}

var xxx_messageInfo_BannerGetNewBannerResp proto.InternalMessageInfo

func (m *BannerGetNewBannerResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BannerGetNewBannerResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BannerGetNewBannerResp) GetData() []*BannerGetNewBannerResp_NewBanner {
	if m != nil {
		return m.Data
	}
	return nil
}

type BannerGetNewBannerResp_NewBanner struct {
	// banner id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 图片地址
	Pic string `protobuf:"bytes,2,opt,name=pic,proto3" json:"pic"`
	// 图片地址
	Img string `protobuf:"bytes,3,opt,name=img,proto3" json:"img"`
	// 跳转链接
	Link string `protobuf:"bytes,4,opt,name=link,proto3" json:"link"`
	// 标题
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title"`
	// 第几帧
	Position string `protobuf:"bytes,6,opt,name=position,proto3" json:"position"`
	// 权重
	SortNum string `protobuf:"bytes,7,opt,name=sort_num,json=sortNum,proto3" json:"sort_num"`
	// 注释
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
}

func (m *BannerGetNewBannerResp_NewBanner) Reset()         { *m = BannerGetNewBannerResp_NewBanner{} }
func (m *BannerGetNewBannerResp_NewBanner) String() string { return proto.CompactTextString(m) }
func (*BannerGetNewBannerResp_NewBanner) ProtoMessage()    {}
func (*BannerGetNewBannerResp_NewBanner) Descriptor() ([]byte, []int) {
	return fileDescriptor_Banner_dd0f870b0c73e0e6, []int{1, 0}
}
func (m *BannerGetNewBannerResp_NewBanner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BannerGetNewBannerResp_NewBanner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BannerGetNewBannerResp_NewBanner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BannerGetNewBannerResp_NewBanner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerGetNewBannerResp_NewBanner.Merge(dst, src)
}
func (m *BannerGetNewBannerResp_NewBanner) XXX_Size() int {
	return m.Size()
}
func (m *BannerGetNewBannerResp_NewBanner) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerGetNewBannerResp_NewBanner.DiscardUnknown(m)
}

var xxx_messageInfo_BannerGetNewBannerResp_NewBanner proto.InternalMessageInfo

func (m *BannerGetNewBannerResp_NewBanner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BannerGetNewBannerResp_NewBanner) GetPic() string {
	if m != nil {
		return m.Pic
	}
	return ""
}

func (m *BannerGetNewBannerResp_NewBanner) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *BannerGetNewBannerResp_NewBanner) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *BannerGetNewBannerResp_NewBanner) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BannerGetNewBannerResp_NewBanner) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *BannerGetNewBannerResp_NewBanner) GetSortNum() string {
	if m != nil {
		return m.SortNum
	}
	return ""
}

func (m *BannerGetNewBannerResp_NewBanner) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*BannerGetNewBannerReq)(nil), "room_ex.v1.BannerGetNewBannerReq")
	proto.RegisterType((*BannerGetNewBannerResp)(nil), "room_ex.v1.BannerGetNewBannerResp")
	proto.RegisterType((*BannerGetNewBannerResp_NewBanner)(nil), "room_ex.v1.BannerGetNewBannerResp.NewBanner")
}
func (m *BannerGetNewBannerReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BannerGetNewBannerReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Platform != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBanner(dAtA, i, uint64(m.Platform))
	}
	if m.Position != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBanner(dAtA, i, uint64(m.Position))
	}
	if len(m.UserPlatform) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.UserPlatform)))
		i += copy(dAtA[i:], m.UserPlatform)
	}
	if len(m.UserDevice) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.UserDevice)))
		i += copy(dAtA[i:], m.UserDevice)
	}
	if m.Build != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintBanner(dAtA, i, uint64(m.Build))
	}
	return i, nil
}

func (m *BannerGetNewBannerResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BannerGetNewBannerResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBanner(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintBanner(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *BannerGetNewBannerResp_NewBanner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BannerGetNewBannerResp_NewBanner) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Pic) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Pic)))
		i += copy(dAtA[i:], m.Pic)
	}
	if len(m.Img) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Img)))
		i += copy(dAtA[i:], m.Img)
	}
	if len(m.Link) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Link)))
		i += copy(dAtA[i:], m.Link)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Position) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Position)))
		i += copy(dAtA[i:], m.Position)
	}
	if len(m.SortNum) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.SortNum)))
		i += copy(dAtA[i:], m.SortNum)
	}
	if len(m.Remark) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintBanner(dAtA, i, uint64(len(m.Remark)))
		i += copy(dAtA[i:], m.Remark)
	}
	return i, nil
}

func encodeVarintBanner(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BannerGetNewBannerReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Platform != 0 {
		n += 1 + sovBanner(uint64(m.Platform))
	}
	if m.Position != 0 {
		n += 1 + sovBanner(uint64(m.Position))
	}
	l = len(m.UserPlatform)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.UserDevice)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	if m.Build != 0 {
		n += 1 + sovBanner(uint64(m.Build))
	}
	return n
}

func (m *BannerGetNewBannerResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovBanner(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovBanner(uint64(l))
		}
	}
	return n
}

func (m *BannerGetNewBannerResp_NewBanner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.Pic)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.Img)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.Link)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.Position)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.SortNum)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	l = len(m.Remark)
	if l > 0 {
		n += 1 + l + sovBanner(uint64(l))
	}
	return n
}

func sovBanner(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBanner(x uint64) (n int) {
	return sovBanner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BannerGetNewBannerReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBanner
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BannerGetNewBannerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BannerGetNewBannerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			m.Platform = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Platform |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			m.Position = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Position |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserPlatform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserPlatform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserDevice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserDevice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Build", wireType)
			}
			m.Build = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Build |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBanner
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BannerGetNewBannerResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBanner
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BannerGetNewBannerResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BannerGetNewBannerResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &BannerGetNewBannerResp_NewBanner{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBanner
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BannerGetNewBannerResp_NewBanner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBanner
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NewBanner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewBanner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Img", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Img = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Link = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortNum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SortNum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remark", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Remark = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBanner
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBanner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBanner
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBanner
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBanner
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBanner
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBanner(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBanner = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBanner   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/Banner.proto", fileDescriptor_Banner_dd0f870b0c73e0e6) }

var fileDescriptor_Banner_dd0f870b0c73e0e6 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xed, 0xd4, 0x89, 0x87, 0x08, 0xd0, 0x4a, 0x54, 0x26, 0xaa, 0xec, 0x90, 0x0b, 0x39,
	0x80, 0xa3, 0x14, 0xbe, 0xc0, 0x42, 0x42, 0xe2, 0x50, 0xa1, 0x95, 0xb8, 0x70, 0x89, 0x1c, 0x7b,
	0x6b, 0x56, 0x8d, 0xbd, 0xae, 0xbd, 0x0e, 0xfc, 0x05, 0xfc, 0x09, 0xbf, 0xc1, 0xb1, 0x47, 0x4e,
	0x16, 0x4a, 0x6e, 0x3e, 0xf2, 0x05, 0x68, 0xc7, 0xae, 0xeb, 0x48, 0x45, 0xbd, 0xcc, 0xce, 0x7b,
	0xfb, 0xd6, 0x33, 0xf3, 0x46, 0x86, 0x27, 0xbb, 0xd5, 0xd2, 0x0f, 0xd2, 0x94, 0xe5, 0x5e, 0x96,
	0x0b, 0x29, 0x08, 0xe4, 0x42, 0x24, 0x6b, 0xf6, 0xcd, 0xdb, 0xad, 0xa6, 0xaf, 0x63, 0x2e, 0xbf,
	0x94, 0x1b, 0x2f, 0x14, 0xc9, 0x32, 0x16, 0xb1, 0x58, 0xa2, 0x64, 0x53, 0x5e, 0x22, 0x42, 0x80,
	0x59, 0xf3, 0x74, 0xfe, 0x57, 0x83, 0x67, 0xcd, 0xb7, 0xde, 0x33, 0x79, 0xc1, 0xbe, 0x36, 0x39,
	0x65, 0xd7, 0x64, 0x01, 0xe3, 0x6c, 0x1b, 0xc8, 0x4b, 0x91, 0x27, 0xb6, 0x36, 0xd3, 0x16, 0x86,
	0x3f, 0xa9, 0x2b, 0xb7, 0xe3, 0x68, 0x97, 0xa1, 0x52, 0x14, 0x5c, 0x72, 0x91, 0xda, 0x7a, 0x4f,
	0xd9, 0x72, 0xb4, 0xcb, 0xc8, 0x5b, 0x98, 0x94, 0x05, 0xcb, 0x3f, 0xde, 0x7e, 0xd7, 0x98, 0x69,
	0x0b, 0xcb, 0x7f, 0x5a, 0x57, 0xee, 0x11, 0x4f, 0x8f, 0x10, 0xf1, 0x00, 0x14, 0x7e, 0xc7, 0x76,
	0x3c, 0x64, 0xf6, 0x10, 0xdf, 0x3c, 0xae, 0x2b, 0xb7, 0xc7, 0xd2, 0x5e, 0x4e, 0x5c, 0x38, 0xd9,
	0x94, 0x7c, 0x1b, 0xd9, 0x27, 0xd8, 0x8c, 0x55, 0x57, 0x6e, 0x43, 0xd0, 0xe6, 0x98, 0xff, 0x34,
	0xe0, 0xf4, 0xbe, 0xa1, 0x8b, 0x8c, 0x9c, 0xc1, 0x30, 0x14, 0x11, 0x6b, 0x27, 0x1e, 0xd7, 0x95,
	0x8b, 0x98, 0x62, 0x24, 0xcf, 0xc1, 0x48, 0x8a, 0x18, 0x87, 0xb4, 0xfc, 0x51, 0x5d, 0xb9, 0x0a,
	0x52, 0x15, 0xc8, 0x07, 0x18, 0x46, 0x81, 0x0c, 0x6c, 0x63, 0x66, 0x2c, 0x1e, 0x9d, 0xbf, 0xf2,
	0xee, 0x56, 0xe2, 0xdd, 0x5f, 0xca, 0xeb, 0x50, 0x53, 0x46, 0xbd, 0xa6, 0x18, 0xa7, 0xdf, 0x75,
	0xb0, 0xba, 0x5b, 0x72, 0x0a, 0x3a, 0x8f, 0xb0, 0x21, 0xcb, 0x37, 0xeb, 0xca, 0xd5, 0x79, 0x44,
	0x75, 0x1e, 0xa9, 0x66, 0x32, 0x1e, 0xf6, 0x9b, 0xc9, 0x78, 0x48, 0x55, 0x50, 0x57, 0x3c, 0x89,
	0x5b, 0x7b, 0xf1, 0x8a, 0x27, 0x31, 0x55, 0x41, 0x0d, 0xb8, 0xe5, 0xe9, 0x55, 0x6b, 0x23, 0x56,
	0x56, 0x98, 0x62, 0x54, 0xd6, 0x49, 0x2e, 0xb7, 0x0c, 0xad, 0xb3, 0x1a, 0xeb, 0x90, 0xa0, 0xcd,
	0x71, 0xb4, 0x6b, 0x13, 0x35, 0xff, 0xdb, 0xf5, 0x4b, 0x18, 0x17, 0x22, 0x97, 0xeb, 0xb4, 0x4c,
	0xec, 0xd1, 0x9d, 0xf2, 0x96, 0xa3, 0x23, 0x95, 0x5d, 0x94, 0x09, 0x99, 0x83, 0x99, 0xb3, 0x24,
	0xc8, 0xaf, 0xec, 0x31, 0xca, 0xa0, 0xae, 0xdc, 0x96, 0xa1, 0xed, 0x79, 0xbe, 0x06, 0xb3, 0x75,
	0xe3, 0x13, 0x4c, 0xe2, 0x9e, 0x93, 0xe4, 0xc5, 0x43, 0x4e, 0x5f, 0x4f, 0xe7, 0x0f, 0x2f, 0xc3,
	0x3f, 0xfb, 0xb5, 0x77, 0xb4, 0x9b, 0xbd, 0xa3, 0xfd, 0xd9, 0x3b, 0xda, 0x8f, 0x83, 0x33, 0xb8,
	0x39, 0x38, 0x83, 0xdf, 0x07, 0x67, 0xf0, 0x59, 0xdf, 0xad, 0x36, 0x26, 0xfe, 0x2c, 0x6f, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x16, 0x9d, 0x49, 0x1e, 0x7a, 0x03, 0x00, 0x00,
}
