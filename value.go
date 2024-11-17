package snmp

import (
	"errors"
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"net"
	"strconv"
	"time"
)

// Value is the interface implemented by all SNMP value types. Sealed.
type Value interface {
	// Stringer is implemented by all SNMP value types.
	// The implementation should return a string representation of the value without the type.
	fmt.Stringer

	// Type returns the type of the value.
	Type() Type
	// RawValue returns the raw value of the value.
	// Note that the concrete types are not guaranteed to be stable across versions. A type-switch on the value
	// is the recommended way to handle the raw value as the implementations expose the value with the concrete type.
	RawValue() any
	// sealed is a marker method to prevent implementation of the interface outside of this package.
	// Does not do anything.
	sealed()
}

type Boolean struct {
	Bool bool
}

// NewBoolean creates a new Boolean value.
func NewBoolean(value bool) Boolean {
	return Boolean{Bool: value}
}

func (b Boolean) String() string {
	return strconv.FormatBool(b.Bool)
}
func (b Boolean) Type() Type {
	return BooleanType{}
}
func (b Boolean) RawValue() any {
	return b.Bool
}
func (b Boolean) sealed() {}

type Integer struct {
	Integer int
}

// NewInteger creates a new Integer value.
func NewInteger(value int) Integer {
	return Integer{Integer: value}
}

func (i Integer) String() string {
	return strconv.FormatInt(int64(i.Integer), 10)
}
func (i Integer) Type() Type {
	return IntegerType{}
}
func (i Integer) RawValue() any {
	return i.Integer
}
func (i Integer) sealed() {

}

type BitString struct {
	BitSet *bitset.BitSet
}

// NewBitString creates a new BitString value.
func NewBitString(value *bitset.BitSet) BitString {
	return BitString{BitSet: value}
}

func (b BitString) String() string {
	return b.BitSet.String()
}
func (b BitString) Type() Type {
	return BitStringType{}
}
func (b BitString) RawValue() any {
	return b.BitSet
}
func (b BitString) sealed() {}

type OctetString struct {
	Bytes []byte
}

// NewOctetString creates a new OctetString value.
func NewOctetString(value []byte) OctetString {
	return OctetString{Bytes: value}
}

func (o OctetString) String() string {
	return string(o.Bytes)
}
func (o OctetString) Type() Type {
	return OctetStringType{}
}
func (o OctetString) RawValue() any {
	return o.Bytes
}
func (o OctetString) sealed() {}

type ObjectIdentifier struct {
	OId OID
}

// NewObjectIdentifier creates a new ObjectIdentifier value.
func NewObjectIdentifier(value OID) ObjectIdentifier {
	return ObjectIdentifier{OId: value}
}

func (o ObjectIdentifier) String() string {
	return o.OId.String()
}
func (o ObjectIdentifier) Type() Type {
	return ObjectIdentifierType{}
}
func (o ObjectIdentifier) RawValue() any {
	return o.OId
}
func (o ObjectIdentifier) sealed() {}

type ObjectDescription struct {
	Str string
}

// NewObjectDescription creates a new ObjectDescription value.
func NewObjectDescription(value string) ObjectDescription {
	return ObjectDescription{Str: value}
}

func (o ObjectDescription) String() string {
	return o.Str
}
func (o ObjectDescription) Type() Type {
	return ObjectDescriptionType{}
}
func (o ObjectDescription) RawValue() any {
	return o.Str
}
func (o ObjectDescription) sealed() {}

type IPAddress struct {
	IP net.IP
}

// NewIPAddress creates a new IPAddress value.
func NewIPAddress(value net.IP) IPAddress {
	return IPAddress{IP: value}
}

func (i IPAddress) String() string {
	return i.IP.String()
}
func (i IPAddress) Type() Type {
	return IPAddressType{}
}
func (i IPAddress) RawValue() any {
	return i.IP
}
func (i IPAddress) sealed() {}

type Counter32 struct {
	UInt32 uint32
}

// NewCounter32 creates a new Counter32 value.
func NewCounter32(value uint32) Counter32 {
	return Counter32{UInt32: value}
}

func (c Counter32) String() string {
	return strconv.FormatUint(uint64(c.UInt32), 10)
}
func (c Counter32) Type() Type {
	return Counter32Type{}
}
func (c Counter32) RawValue() any {
	return c.UInt32
}
func (c Counter32) sealed() {}

type Opaque struct {
	Bytes []byte
}

// NewOpaque creates a new Opaque value.
func NewOpaque(value []byte) Opaque {
	return Opaque{Bytes: value}
}

func (o Opaque) String() string {
	return fmt.Sprintf("%x", o.Bytes)
}
func (o Opaque) Type() Type {
	return OpaqueType{}
}
func (o Opaque) RawValue() any {
	return o.Bytes
}
func (o Opaque) sealed() {}

type NsapAddress struct {
	// FIXME: Find better representation
	Bytes []byte
}

// NewNsapAddress creates a new NsapAddress value.
func NewNsapAddress(value []byte) NsapAddress {
	return NsapAddress{Bytes: value}
}

func (n NsapAddress) String() string {
	return fmt.Sprintf("%x", n.Bytes)
}
func (n NsapAddress) Type() Type {
	return NsapAddressType{}
}
func (n NsapAddress) RawValue() any {
	return n.Bytes
}
func (n NsapAddress) sealed() {}

type Counter64 struct {
	Int64 uint64
}

// NewCounter64 creates a new Counter64 value.
func NewCounter64(value uint64) Counter64 {
	return Counter64{Int64: value}
}

func (c Counter64) String() string {
	return strconv.FormatUint(c.Int64, 10)
}
func (c Counter64) Type() Type {
	return Counter64Type{}
}
func (c Counter64) RawValue() any {
	return c.Int64
}
func (c Counter64) sealed() {}

type UInteger32 struct {
	UInt32 uint32
}

// NewUInteger32 creates a new UInteger32 value.
func NewUInteger32(value uint32) UInteger32 {
	return UInteger32{UInt32: value}
}

func (u UInteger32) String() string {
	return strconv.FormatUint(uint64(u.UInt32), 10)
}
func (u UInteger32) Type() Type {
	return UInteger32Type{}
}
func (u UInteger32) RawValue() any {
	return u.UInt32
}
func (u UInteger32) sealed() {}

type OpaqueFloat struct {
	Float32 float32
}

// NewOpaqueFloat creates a new OpaqueFloat value.
func NewOpaqueFloat(value float32) OpaqueFloat {
	return OpaqueFloat{Float32: value}
}

func (o OpaqueFloat) String() string {
	return strconv.FormatFloat(float64(o.Float32), 'f', -1, 32)
}
func (o OpaqueFloat) Type() Type {
	return OpaqueFloatType{}
}
func (o OpaqueFloat) RawValue() any {
	return o.Float32
}
func (o OpaqueFloat) sealed() {}

type OpaqueDouble struct {
	Float64 float64
}

// NewOpaqueDouble creates a new OpaqueDouble value.
func NewOpaqueDouble(value float64) OpaqueDouble {
	return OpaqueDouble{Float64: value}
}

func (o OpaqueDouble) String() string {
	return strconv.FormatFloat(o.Float64, 'f', -1, 64)
}
func (o OpaqueDouble) Type() Type {
	return OpaqueDoubleType{}
}
func (o OpaqueDouble) RawValue() any {
	return o.Float64
}
func (o OpaqueDouble) sealed() {}

type TimeTicks struct {
	Duration time.Duration
}

// NewTimeTicks creates a new TimeTicks value.
func NewTimeTicks(value time.Duration) TimeTicks {
	return TimeTicks{Duration: value}
}

// NewTimeTicksFromHundredths creates a new TimeTicks value from the given value in hundredths of a second, i.e.
// the representation used in SNMP.
func NewTimeTicksFromHundredths(value uint32) TimeTicks {
	return NewTimeTicks(time.Duration(value) * 100 * time.Millisecond)
}

func (t TimeTicks) String() string {
	return t.Duration.String()
}
func (t TimeTicks) Type() Type {
	return TimeTicksType{}
}
func (t TimeTicks) RawValue() any { return t.Duration }
func (t TimeTicks) sealed()       {}

type Null struct{}

func (n Null) String() string {
	return "NULL"
}
func (n Null) Type() Type {
	return NullType{}
}
func (n Null) RawValue() any {
	return nil
}
func (n Null) sealed() {}

// NoSuchObject is a value indicating that the object does not exist on the SNMP agent.
// Should be interpreted as an error (and will be returned as such by the client).
type NoSuchObject struct{}

func (n NoSuchObject) String() string {
	return "NO SUCH OBJECT"
}
func (n NoSuchObject) Type() Type {
	return NoSuchObjectType{}
}
func (n NoSuchObject) RawValue() any {
	return nil
}
func (n NoSuchObject) sealed() {}

// NoSuchInstance is a value indicating that the instance does not exist on the SNMP agent.
// Should be interpreted as an error (and will be returned as such by the client).
type NoSuchInstance struct{}

func (n NoSuchInstance) String() string { return "NO SUCH INSTANCE" }
func (n NoSuchInstance) Type() Type     { return NoSuchInstanceType{} }
func (n NoSuchInstance) RawValue() any  { return nil }
func (n NoSuchInstance) sealed()        {}

// EndOfMibView is a value indicating the end of the MIB view.
// Should be interpreted as an error (and will be returned as such by the client).
type EndOfMibView struct{}

func (e EndOfMibView) String() string { return "END OF MIB VIEW" }
func (e EndOfMibView) Type() Type     { return EndOfMibViewType{} }
func (e EndOfMibView) RawValue() any  { return nil }
func (e EndOfMibView) sealed()        {}

// decodeBERValue parses a BER-encoded value (Varbind) the given BER-encoded value.
// The input is assumed to be the value part of a BER-encoded TLV.
// The returned value is the decoded value, the input slice offset by the number of bytes used to decode the value and an error, if any.
// The number of bytes may be used as an offset to skip over the value in the input.
func decodeBERValue(b []byte) (Value, []byte, error) {
	if len(b) == 0 {
		return nil, b, fmt.Errorf("no data")
	}

	switch typeTag(b[0]) {
	case tagInteger:
		// Integer is internally represented as int32, so we decode it as int32 and convert it to int.
		return tryDecode(b[1:], decodeBERInt32, func(t int32) Integer {
			return NewInteger(int(t))
		})
	case tagUTnteger32:
		return tryDecode(b[1:], decodeBERUInt32, NewUInteger32)
	case tagOctetString:
		return tryDecode(b[1:], decodeBERBytes, NewOctetString)
	case tagObjectIdentifier:
		return tryDecode(b[1:], decodeBERObjectIdentifier, NewObjectIdentifier)
	case tagIpAddress:
		return tryDecode(b[1:], decodeBERIpAddress, NewIPAddress)
	case tagCounter32:
		return tryDecode(b[1:], decodeBERUInt32, NewCounter32)
	case tagGauge32:
		return tryDecode(b[1:], decodeBERUInt32, NewUInteger32)
	case tagTimeTicks:
		return tryDecode(b[1:], decodeBERUInt32, NewTimeTicksFromHundredths)
	case tagOpaque:
		return tryDecode(b[1:], decodeBERBytes, NewOpaque)
	case tagCounter64:
		return tryDecode(b[1:], decodeBERUInt64, NewCounter64)
	case tagNull:
		return Null{}, b[1:], nil
	case tagNoSuchObject:
		return NoSuchObject{}, b[1:], nil
	case tagNoSuchInstance:
		return NoSuchInstance{}, b[1:], nil
	case tagEndOfMibView:
		return EndOfMibView{}, b[1:], nil
	}

	return nil, b, fmt.Errorf("unsupported type tag %d", b[0])
}

// tryDecode tries to decode the given BER-encoded value using the given decode function and creates a new Value
// using the given value constructor.
// This is a helper function to reduce code duplication in decodeBERValue.
func tryDecode[T any, V Value](b []byte, decodeFn func([]byte) (T, []byte, error), valueConstr func(T) V) (Value, []byte, error) {
	v, o, err := decodeFn(b)
	if err != nil {
		return nil, o, err
	}

	return valueConstr(v), o, nil
}

// decodeBERLength decodes the length of a BER-encoded value.
// It is assumed that the first byte of the input is the first byte of the length field.
//
// The length is encoded as a variable-length integer, which may be in one of the following forms:
//   - Short form: The length is encoded in the lower 7 bits of the byte. The most significant bit is 0.
//   - Long form: The lower 7 bits of the byte are the number of bytes used to encode the length. The most significant
//     bit is 1. The following bytes encode the length as a big-endian integer.
//   - Indefinite: The length byte is 0x80. The length is not specified and the value is terminated by a special
//     end-of-contents byte.
//
// Returns the length and the input slice offset by the number of bytes used to decode the length.
// If indefinite length is encountered, the length is returned as -1.
// Adapted from https://github.com/gosnmp/gosnmp
func decodeBERLength(b []byte) (length int, tail []byte, _ error) {
	// Special case for empty octet strings deliberately omitted
	// That case should be handled by the octet string decoder itself

	if len(b) == 0 {
		return 0, b, fmt.Errorf("no data")
	}

	// Indefinite length - return -1
	if b[0] == 0x80 {
		return -1, b[1:], nil
	}

	// Short form - interpret the byte as the length and return it
	// No check for negative values is needed as the first bit is by definition 0 for short form
	if b[0] <= 0x7F { // 0x7F = 127 = 0111 1111
		return int(b[0]), b[1:], nil
	}

	// Long form - interpret the byte as the number of bytes used to encode the length
	// Then read the following bytes as the length
	lengthLength := int(b[0] & 0x7f) // 0x7f = 0111 1111
	if lengthLength > len(b)-1 {
		return 0, b, fmt.Errorf("length of length is longer than the remaining data: %d > %d", lengthLength, len(b)-1)
	}

	for i := 0; i < lengthLength; i++ {
		// Left shift the length by 8 bits to  and add the next byte
		length = length<<8 | int(b[i+1])
	}

	// Check for overflow
	if length < 0 {
		return 0, b, fmt.Errorf("negative length (overflow): %d", length)
	}

	return length, b[lengthLength+1:], nil // +1 for the first byte indicating the long form
}

// decodeBERUint decodes a BER-encoded unsigned integer with the given size.
func decodeBERUInt(b []byte, size int) (uint64, []byte, error) {
	length, b, err := decodeBERLength(b)
	if err != nil {
		return 0, b, err
	}

	if length == 0 {
		return 0, b, errors.New("zero-length integer")
	}
	if length > size {
		return 0, b, fmt.Errorf("integer too large (%d bytes) to fit into integer of bit-size %d", length, size*8)
	}
	if length > len(b) {
		return 0, b, fmt.Errorf("truncated integer: %d bytes, expected %d bytes", length, len(b))
	}

	var value uint64
	for i := 0; i < length; i++ {
		// Left shift the value by 8 bits and add the next byte
		value = value<<8 | uint64(b[i])
	}

	return value, b[length:], nil
}

// decodeBERInt decodes a BER-encoded integer with the given size.
func decodeBERInt(b []byte, size int) (int64, []byte, error) {
	length, b, err := decodeBERLength(b)
	if err != nil {
		return 0, b, err
	}

	if length == 0 {
		return 0, b, errors.New("zero-length integer")
	}
	if length > size {
		return 0, b, fmt.Errorf("integer too large (%d bytes) to fit into integer of bit-size %d", length, size*8)
	}
	if length > len(b) {
		return 0, b, fmt.Errorf("truncated integer: %d bytes, expected %d bytes", length, len(b))
	}

	var value int64
	for i := 0; i < length; i++ {
		// Left shift the value by 8 bits and add the next byte
		value = value<<8 | int64(b[i])
	}

	// Sign-extend the value if necessary
	if value&(1<<(8*length-1)) != 0 {
		value |= -1 << (8 * length)
	}

	return value, b[length:], nil
}

func decodeBERInt32(b []byte) (int32, []byte, error) {
	v, b, err := decodeBERInt(b, 4)
	if err != nil {
		return 0, b, err
	}

	return int32(v), b, nil
}

func decodeBERUInt32(b []byte) (uint32, []byte, error) {
	v, b, err := decodeBERUInt(b, 4)
	if err != nil {
		return 0, b, err
	}

	return uint32(v), b, nil
}

func decodeBERUInt64(b []byte) (uint64, []byte, error) {
	return decodeBERUInt(b, 8)
}

func decodeBERBytes(b []byte) ([]byte, []byte, error) {
	length, b, err := decodeBERLength(b)
	if err != nil {
		return nil, b, err
	}

	if length > len(b) {
		return nil, b, fmt.Errorf("truncated bytes: %d bytes, expected %d bytes", length, len(b))
	}

	return b[:length], b[length:], nil
}

func decodeBERObjectIdentifier(b []byte) (OID, []byte, error) {
	return nil, b, errors.New("not implemented")
}

func decodeBERIpAddress(b []byte) (net.IP, []byte, error) {
	return nil, b, errors.New("not implemented")
}
