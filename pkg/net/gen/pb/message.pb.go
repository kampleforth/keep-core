// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/message.proto

package pb

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	bytes "bytes"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// BroadcastNetworkMessage represents a network message used by broadcast
// channels.
type BroadcastNetworkMessage struct {
	// The PublicKey of the sender.
	Sender []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// A marshaled Protocol Message.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Type of the message as registered by the protocol.
	Type []byte `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Sequence number of the message. Retransmissions have the same sequence
	// number as the original message.
	SequenceNumber uint64 `protobuf:"varint,4,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
}

func (m *BroadcastNetworkMessage) Reset()                    { *m = BroadcastNetworkMessage{} }
func (*BroadcastNetworkMessage) ProtoMessage()               {}
func (*BroadcastNetworkMessage) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

func (m *BroadcastNetworkMessage) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *BroadcastNetworkMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *BroadcastNetworkMessage) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *BroadcastNetworkMessage) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

// UnicastNetworkMessage represents a network message used by unicast
// channels.
type UnicastNetworkMessage struct {
	// The PublicKey of the sender.
	Sender []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// A marshaled Protocol Message.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Type of the message as registered by the protocol.
	Type []byte `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Message signature.
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *UnicastNetworkMessage) Reset()                    { *m = UnicastNetworkMessage{} }
func (*UnicastNetworkMessage) ProtoMessage()               {}
func (*UnicastNetworkMessage) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func (m *UnicastNetworkMessage) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *UnicastNetworkMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UnicastNetworkMessage) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *UnicastNetworkMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Identity struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{2} }

func (m *Identity) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func init() {
	proto.RegisterType((*BroadcastNetworkMessage)(nil), "net.BroadcastNetworkMessage")
	proto.RegisterType((*UnicastNetworkMessage)(nil), "net.UnicastNetworkMessage")
	proto.RegisterType((*Identity)(nil), "net.Identity")
}
func (this *BroadcastNetworkMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BroadcastNetworkMessage)
	if !ok {
		that2, ok := that.(BroadcastNetworkMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Sender, that1.Sender) {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if !bytes.Equal(this.Type, that1.Type) {
		return false
	}
	if this.SequenceNumber != that1.SequenceNumber {
		return false
	}
	return true
}
func (this *UnicastNetworkMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnicastNetworkMessage)
	if !ok {
		that2, ok := that.(UnicastNetworkMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Sender, that1.Sender) {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if !bytes.Equal(this.Type, that1.Type) {
		return false
	}
	if !bytes.Equal(this.Signature, that1.Signature) {
		return false
	}
	return true
}
func (this *Identity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Identity)
	if !ok {
		that2, ok := that.(Identity)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.PubKey, that1.PubKey) {
		return false
	}
	return true
}
func (this *BroadcastNetworkMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.BroadcastNetworkMessage{")
	s = append(s, "Sender: "+fmt.Sprintf("%#v", this.Sender)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "SequenceNumber: "+fmt.Sprintf("%#v", this.SequenceNumber)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UnicastNetworkMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.UnicastNetworkMessage{")
	s = append(s, "Sender: "+fmt.Sprintf("%#v", this.Sender)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Signature: "+fmt.Sprintf("%#v", this.Signature)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Identity) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.Identity{")
	s = append(s, "PubKey: "+fmt.Sprintf("%#v", this.PubKey)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *BroadcastNetworkMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BroadcastNetworkMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.Payload) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.SequenceNumber != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.SequenceNumber))
	}
	return i, nil
}

func (m *UnicastNetworkMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnicastNetworkMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.Payload) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	return i, nil
}

func (m *Identity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Identity) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.PubKey)))
		i += copy(dAtA[i:], m.PubKey)
	}
	return i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BroadcastNetworkMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.SequenceNumber != 0 {
		n += 1 + sovMessage(uint64(m.SequenceNumber))
	}
	return n
}

func (m *UnicastNetworkMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *Identity) Size() (n int) {
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BroadcastNetworkMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BroadcastNetworkMessage{`,
		`Sender:` + fmt.Sprintf("%v", this.Sender) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`SequenceNumber:` + fmt.Sprintf("%v", this.SequenceNumber) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UnicastNetworkMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnicastNetworkMessage{`,
		`Sender:` + fmt.Sprintf("%v", this.Sender) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Signature:` + fmt.Sprintf("%v", this.Signature) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Identity) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Identity{`,
		`PubKey:` + fmt.Sprintf("%v", this.PubKey) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BroadcastNetworkMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: BroadcastNetworkMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BroadcastNetworkMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = append(m.Type[:0], dAtA[iNdEx:postIndex]...)
			if m.Type == nil {
				m.Type = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *UnicastNetworkMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: UnicastNetworkMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnicastNetworkMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = append(m.Type[:0], dAtA[iNdEx:postIndex]...)
			if m.Type == nil {
				m.Type = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Identity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Identity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(dAtA[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4b,
	0x2d, 0x51, 0x6a, 0x67, 0xe4, 0x12, 0x77, 0x2a, 0xca, 0x4f, 0x4c, 0x49, 0x4e, 0x2c, 0x2e, 0xf1,
	0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xf6, 0x85, 0x28, 0x13, 0x12, 0xe3, 0x62, 0x2b, 0x4e, 0xcd,
	0x4b, 0x49, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0xf2, 0x84, 0x24, 0xb8, 0xd8,
	0x0b, 0x12, 0x2b, 0x73, 0xf2, 0x13, 0x53, 0x24, 0x98, 0xc0, 0x12, 0x30, 0xae, 0x90, 0x10, 0x17,
	0x4b, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x33, 0x58, 0x18, 0xcc, 0x16, 0x52, 0xe3, 0xe2, 0x2b, 0x4e,
	0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0x4e, 0xf5, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x92, 0x60, 0x51, 0x60,
	0xd4, 0x60, 0x09, 0x42, 0x13, 0x55, 0xaa, 0xe6, 0x12, 0x0d, 0xcd, 0xcb, 0xa4, 0x99, 0x33, 0x64,
	0xb8, 0x38, 0x8b, 0x33, 0xd3, 0xf3, 0x12, 0x4b, 0x4a, 0x8b, 0x52, 0xc1, 0x2e, 0xe0, 0x09, 0x42,
	0x08, 0x28, 0x29, 0x73, 0x71, 0x78, 0xa6, 0xa4, 0xe6, 0x95, 0x64, 0x96, 0x54, 0x0a, 0x89, 0x73,
	0xb1, 0x17, 0x94, 0x26, 0xc5, 0x67, 0xa7, 0x56, 0xc2, 0x2c, 0x2c, 0x28, 0x4d, 0xf2, 0x4e, 0xad,
	0x74, 0x32, 0xb8, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x1b,
	0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x51, 0x4c, 0x05, 0x49, 0x49, 0x6c, 0xe0, 0x90, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0x82, 0x1f, 0xda, 0x7d, 0x01, 0x00, 0x00,
}
