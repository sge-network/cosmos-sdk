// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gov/v1/tmpgov.proto

package v1

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Proposal defines the core field members of a governance proposal.
// TODO: to be removed after SGE chain upgrade.
type TmpProposal struct {
	Id       uint64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Messages []*types.Any   `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Status   ProposalStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cosmos.gov.v1.ProposalStatus" json:"status,omitempty"`
	// final_tally_result is the final tally result of the proposal. When
	// querying a proposal via gRPC, this field is not populated until the
	// proposal's voting period has ended.
	FinalTallyResult *TallyResult  `protobuf:"bytes,4,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result,omitempty"`
	SubmitTime       *time.Time    `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3,stdtime" json:"submit_time,omitempty"`
	DepositEndTime   *time.Time    `protobuf:"bytes,6,opt,name=deposit_end_time,json=depositEndTime,proto3,stdtime" json:"deposit_end_time,omitempty"`
	TotalDeposit     []types1.Coin `protobuf:"bytes,7,rep,name=total_deposit,json=totalDeposit,proto3" json:"total_deposit"`
	VotingStartTime  *time.Time    `protobuf:"bytes,8,opt,name=voting_start_time,json=votingStartTime,proto3,stdtime" json:"voting_start_time,omitempty"`
	VotingEndTime    *time.Time    `protobuf:"bytes,9,opt,name=voting_end_time,json=votingEndTime,proto3,stdtime" json:"voting_end_time,omitempty"`
	// metadata is any arbitrary metadata attached to the proposal.
	Metadata    string `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	IsExpedited bool   `protobuf:"varint,11,opt,name=is_expedited,json=isExpedited,proto3" json:"is_expedited,omitempty"`
}

func (m *TmpProposal) Reset()         { *m = TmpProposal{} }
func (m *TmpProposal) String() string { return proto.CompactTextString(m) }
func (*TmpProposal) ProtoMessage()    {}
func (*TmpProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fc266f82120e475, []int{0}
}
func (m *TmpProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TmpProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TmpProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TmpProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TmpProposal.Merge(m, src)
}
func (m *TmpProposal) XXX_Size() int {
	return m.Size()
}
func (m *TmpProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TmpProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TmpProposal proto.InternalMessageInfo

func (m *TmpProposal) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TmpProposal) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *TmpProposal) GetStatus() ProposalStatus {
	if m != nil {
		return m.Status
	}
	return ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED
}

func (m *TmpProposal) GetFinalTallyResult() *TallyResult {
	if m != nil {
		return m.FinalTallyResult
	}
	return nil
}

func (m *TmpProposal) GetSubmitTime() *time.Time {
	if m != nil {
		return m.SubmitTime
	}
	return nil
}

func (m *TmpProposal) GetDepositEndTime() *time.Time {
	if m != nil {
		return m.DepositEndTime
	}
	return nil
}

func (m *TmpProposal) GetTotalDeposit() []types1.Coin {
	if m != nil {
		return m.TotalDeposit
	}
	return nil
}

func (m *TmpProposal) GetVotingStartTime() *time.Time {
	if m != nil {
		return m.VotingStartTime
	}
	return nil
}

func (m *TmpProposal) GetVotingEndTime() *time.Time {
	if m != nil {
		return m.VotingEndTime
	}
	return nil
}

func (m *TmpProposal) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *TmpProposal) GetIsExpedited() bool {
	if m != nil {
		return m.IsExpedited
	}
	return false
}

func init() {
	proto.RegisterType((*TmpProposal)(nil), "cosmos.gov.v1.TmpProposal")
}

func init() { proto.RegisterFile("cosmos/gov/v1/tmpgov.proto", fileDescriptor_9fc266f82120e475) }

var fileDescriptor_9fc266f82120e475 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xba, 0x52, 0x3a, 0x67, 0x2d, 0xc3, 0x9a, 0x44, 0x56, 0x89, 0x34, 0x70, 0x8a, 0x84,
	0xb0, 0x69, 0x11, 0x3f, 0x60, 0x65, 0x95, 0x26, 0xc4, 0x01, 0x65, 0x3d, 0x71, 0x89, 0xdc, 0xc6,
	0x0b, 0x16, 0x49, 0x1c, 0xd5, 0x2f, 0xd1, 0xfa, 0x2f, 0xf6, 0xb3, 0x76, 0xdc, 0x91, 0x13, 0xa0,
	0xf6, 0x6f, 0x70, 0x40, 0x71, 0x9c, 0xc2, 0x7a, 0xea, 0x29, 0xf6, 0xe7, 0xef, 0xfb, 0xde, 0xf7,
	0x9e, 0x5e, 0xd0, 0x70, 0x29, 0x55, 0x2a, 0x15, 0x8d, 0x65, 0x49, 0xcb, 0x31, 0x85, 0x34, 0x8f,
	0x65, 0x49, 0xf2, 0x95, 0x04, 0x89, 0xfb, 0xf5, 0x1b, 0xa9, 0x90, 0x72, 0x3c, 0x74, 0x0d, 0x75,
	0xc1, 0x14, 0xa7, 0xe5, 0x78, 0xc1, 0x81, 0x8d, 0xe9, 0x52, 0x8a, 0xac, 0xa6, 0x0f, 0xcf, 0x62,
	0x19, 0x4b, 0x7d, 0xa4, 0xd5, 0xc9, 0xa0, 0xa3, 0x58, 0xca, 0x38, 0xe1, 0x54, 0xdf, 0x16, 0xc5,
	0x0d, 0x05, 0x91, 0x72, 0x05, 0x2c, 0xcd, 0x0d, 0xe1, 0x7c, 0x9f, 0xc0, 0xb2, 0xb5, 0x79, 0x7a,
	0xf1, 0x38, 0xdc, 0x2e, 0xd9, 0xeb, 0x3f, 0x1d, 0x64, 0xcf, 0xd3, 0xfc, 0xcb, 0x4a, 0xe6, 0x52,
	0xb1, 0x04, 0x0f, 0x50, 0x5b, 0x44, 0x8e, 0xe5, 0x59, 0x7e, 0x27, 0x68, 0x8b, 0x08, 0xbf, 0x43,
	0xbd, 0x94, 0x2b, 0xc5, 0x62, 0xae, 0x9c, 0xb6, 0x77, 0xe4, 0xdb, 0x93, 0x33, 0x52, 0x97, 0x21,
	0x4d, 0x19, 0x72, 0x91, 0xad, 0x83, 0x1d, 0x0b, 0x7f, 0x40, 0x5d, 0x05, 0x0c, 0x0a, 0xe5, 0x1c,
	0x79, 0x96, 0x3f, 0x98, 0xbc, 0x24, 0x8f, 0x9a, 0x27, 0x4d, 0xa9, 0x6b, 0x4d, 0x0a, 0x0c, 0x19,
	0x5f, 0x21, 0x7c, 0x23, 0x32, 0x96, 0x84, 0xc0, 0x92, 0x64, 0x1d, 0xae, 0xb8, 0x2a, 0x12, 0x70,
	0x3a, 0x9e, 0xe5, 0xdb, 0x93, 0xe1, 0x9e, 0xc5, 0xbc, 0xa2, 0x04, 0x9a, 0x11, 0x9c, 0x6a, 0xd5,
	0x7f, 0x08, 0xbe, 0x40, 0xb6, 0x2a, 0x16, 0xa9, 0x80, 0xb0, 0x1a, 0x90, 0xf3, 0xc4, 0x58, 0xec,
	0xa7, 0x9e, 0x37, 0xd3, 0x9b, 0x76, 0xee, 0x7e, 0x8d, 0xac, 0x00, 0xd5, 0xa2, 0x0a, 0xc6, 0x9f,
	0xd0, 0x69, 0xc4, 0x73, 0xa9, 0x04, 0x84, 0x3c, 0x8b, 0x6a, 0x9f, 0xee, 0x81, 0x3e, 0x03, 0xa3,
	0x9c, 0x65, 0x91, 0xf6, 0xba, 0x44, 0x7d, 0x90, 0xc0, 0x92, 0xd0, 0xe0, 0xce, 0x53, 0x3d, 0xc6,
	0xf3, 0xa6, 0xa7, 0x6a, 0x09, 0x88, 0x59, 0x02, 0xf2, 0x51, 0x8a, 0x6c, 0xda, 0xb9, 0xff, 0x39,
	0x6a, 0x05, 0x27, 0x5a, 0x75, 0x59, 0x8b, 0xf0, 0x67, 0xf4, 0xbc, 0x94, 0x20, 0xb2, 0x38, 0x54,
	0xc0, 0x56, 0xa6, 0xb5, 0xde, 0x81, 0x91, 0x9e, 0xd5, 0xd2, 0xeb, 0x4a, 0xa9, 0x33, 0x5d, 0x21,
	0x03, 0xfd, 0x6b, 0xef, 0xf8, 0x40, 0xaf, 0x7e, 0x2d, 0x6c, 0xba, 0x1b, 0x56, 0xfb, 0x01, 0x2c,
	0x62, 0xc0, 0x1c, 0xe4, 0x59, 0xfe, 0x71, 0xb0, 0xbb, 0xe3, 0x57, 0xe8, 0x44, 0xa8, 0x90, 0xdf,
	0xe6, 0x3c, 0x12, 0xc0, 0x23, 0xc7, 0xf6, 0x2c, 0xbf, 0x17, 0xd8, 0x42, 0xcd, 0x1a, 0x68, 0x3a,
	0xbb, 0xdf, 0xb8, 0xd6, 0xc3, 0xc6, 0xb5, 0x7e, 0x6f, 0x5c, 0xeb, 0x6e, 0xeb, 0xb6, 0x1e, 0xb6,
	0x6e, 0xeb, 0xc7, 0xd6, 0x6d, 0x7d, 0x7d, 0x13, 0x0b, 0xf8, 0x56, 0x2c, 0xc8, 0x52, 0xa6, 0xd4,
	0x2c, 0x6f, 0xfd, 0x79, 0xab, 0xa2, 0xef, 0xf4, 0x56, 0x6f, 0x32, 0xac, 0x73, 0xae, 0xaa, 0x1f,
	0xa8, 0xab, 0xe3, 0xbe, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x2b, 0x37, 0xb1, 0x84, 0x03,
	0x00, 0x00,
}

func (m *TmpProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TmpProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TmpProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsExpedited {
		i--
		if m.IsExpedited {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTmpgov(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x52
	}
	if m.VotingEndTime != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.VotingEndTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.VotingEndTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTmpgov(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x4a
	}
	if m.VotingStartTime != nil {
		n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.VotingStartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.VotingStartTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintTmpgov(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x42
	}
	if len(m.TotalDeposit) > 0 {
		for iNdEx := len(m.TotalDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTmpgov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.DepositEndTime != nil {
		n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.DepositEndTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.DepositEndTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintTmpgov(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x32
	}
	if m.SubmitTime != nil {
		n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.SubmitTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.SubmitTime):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintTmpgov(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0x2a
	}
	if m.FinalTallyResult != nil {
		{
			size, err := m.FinalTallyResult.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTmpgov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Status != 0 {
		i = encodeVarintTmpgov(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTmpgov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != 0 {
		i = encodeVarintTmpgov(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTmpgov(dAtA []byte, offset int, v uint64) int {
	offset -= sovTmpgov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TmpProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTmpgov(uint64(m.Id))
	}
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTmpgov(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTmpgov(uint64(m.Status))
	}
	if m.FinalTallyResult != nil {
		l = m.FinalTallyResult.Size()
		n += 1 + l + sovTmpgov(uint64(l))
	}
	if m.SubmitTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.SubmitTime)
		n += 1 + l + sovTmpgov(uint64(l))
	}
	if m.DepositEndTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.DepositEndTime)
		n += 1 + l + sovTmpgov(uint64(l))
	}
	if len(m.TotalDeposit) > 0 {
		for _, e := range m.TotalDeposit {
			l = e.Size()
			n += 1 + l + sovTmpgov(uint64(l))
		}
	}
	if m.VotingStartTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.VotingStartTime)
		n += 1 + l + sovTmpgov(uint64(l))
	}
	if m.VotingEndTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.VotingEndTime)
		n += 1 + l + sovTmpgov(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTmpgov(uint64(l))
	}
	if m.IsExpedited {
		n += 2
	}
	return n
}

func sovTmpgov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTmpgov(x uint64) (n int) {
	return sovTmpgov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TmpProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmpgov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TmpProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TmpProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalTallyResult", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FinalTallyResult == nil {
				m.FinalTallyResult = &TallyResult{}
			}
			if err := m.FinalTallyResult.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubmitTime == nil {
				m.SubmitTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.SubmitTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DepositEndTime == nil {
				m.DepositEndTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.DepositEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalDeposit = append(m.TotalDeposit, types1.Coin{})
			if err := m.TotalDeposit[len(m.TotalDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VotingStartTime == nil {
				m.VotingStartTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.VotingStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VotingEndTime == nil {
				m.VotingEndTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.VotingEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTmpgov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTmpgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsExpedited", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsExpedited = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTmpgov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTmpgov
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
func skipTmpgov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTmpgov
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
					return 0, ErrIntOverflowTmpgov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTmpgov
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
			if length < 0 {
				return 0, ErrInvalidLengthTmpgov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTmpgov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTmpgov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTmpgov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTmpgov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTmpgov = fmt.Errorf("proto: unexpected end of group")
)