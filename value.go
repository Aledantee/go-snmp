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
	Value bool
}

// NewBoolean creates a new Boolean value.
func NewBoolean(value bool) Boolean {
	return Boolean{Value: value}
}

func (b Boolean) String() string {
	return strconv.FormatBool(b.Value)
}
func (b Boolean) Type() Type {
	return BooleanType{}
}
func (b Boolean) RawValue() any {
	return b.Value
}
func (b Boolean) sealed() {}

type Integer struct {
	Value int
}

// NewInteger creates a new Integer value.
func NewInteger(value int) Integer {
	return Integer{Value: value}
}

func (i Integer) String() string {
	return strconv.FormatInt(int64(i.Value), 10)
}
func (i Integer) Type() Type {
	return IntegerType{}
}
func (i Integer) RawValue() any {
	return i.Value
}
func (i Integer) sealed() {

}

type BitString struct {
	Value *bitset.BitSet
}

// NewBitString creates a new BitString value.
func NewBitString(value *bitset.BitSet) BitString {
	return BitString{Value: value}
}

func (b BitString) String() string {
	return b.Value.String()
}
func (b BitString) Type() Type {
	return BitStringType{}
}
func (b BitString) RawValue() any {
	return b.Value
}
func (b BitString) sealed() {}

type OctetString struct {
	Value []byte
}

// NewOctetString creates a new OctetString value.
func NewOctetString(value []byte) OctetString {
	return OctetString{Value: value}
}

func (o OctetString) String() string {
	return string(o.Value)
}
func (o OctetString) Type() Type {
	return OctetStringType{}
}
func (o OctetString) RawValue() any {
	return o.Value
}
func (o OctetString) sealed() {}

type ObjectIdentifier struct {
	Value OID
}

// NewObjectIdentifier creates a new ObjectIdentifier value.
func NewObjectIdentifier(value OID) ObjectIdentifier {
	return ObjectIdentifier{Value: value}
}

func (o ObjectIdentifier) String() string {
	return o.Value.String()
}
func (o ObjectIdentifier) Type() Type {
	return ObjectIdentifierType{}
}
func (o ObjectIdentifier) RawValue() any {
	return o.Value
}
func (o ObjectIdentifier) sealed() {}

type ObjectDescription struct {
	Value string
}

// NewObjectDescription creates a new ObjectDescription value.
func NewObjectDescription(value string) ObjectDescription {
	return ObjectDescription{Value: value}
}

func (o ObjectDescription) String() string {
	return o.Value
}
func (o ObjectDescription) Type() Type {
	return ObjectDescriptionType{}
}
func (o ObjectDescription) RawValue() any {
	return o.Value
}
func (o ObjectDescription) sealed() {}

type IPAddress struct {
	Value net.IP
}

// NewIPAddress creates a new IPAddress value.
func NewIPAddress(value net.IP) IPAddress {
	return IPAddress{Value: value}
}

func (i IPAddress) String() string {
	return i.Value.String()
}
func (i IPAddress) Type() Type {
	return IPAddressType{}
}
func (i IPAddress) RawValue() any {
	return i.Value
}
func (i IPAddress) sealed() {}

type Counter32 struct {
	Value uint32
}

// NewCounter32 creates a new Counter32 value.
func NewCounter32(value uint32) Counter32 {
	return Counter32{Value: value}
}

func (c Counter32) String() string {
	return strconv.FormatUint(uint64(c.Value), 10)
}
func (c Counter32) Type() Type {
	return Counter32Type{}
}
func (c Counter32) RawValue() any {
	return c.Value
}
func (c Counter32) sealed() {}

type Opaque struct {
	Value []byte
}

// NewOpaque creates a new Opaque value.
func NewOpaque(value []byte) Opaque {
	return Opaque{Value: value}
}

func (o Opaque) String() string {
	return fmt.Sprintf("%x", o.Value)
}
func (o Opaque) Type() Type {
	return OpaqueType{}
}
func (o Opaque) RawValue() any {
	return o.Value
}
func (o Opaque) sealed() {}

type NsapAddress struct {
	// FIXME: Find better representation
	Value []byte
}

// NewNsapAddress creates a new NsapAddress value.
func NewNsapAddress(value []byte) NsapAddress {
	return NsapAddress{Value: value}
}

func (n NsapAddress) String() string {
	return fmt.Sprintf("%x", n.Value)
}
func (n NsapAddress) Type() Type {
	return NsapAddressType{}
}
func (n NsapAddress) RawValue() any {
	return n.Value
}
func (n NsapAddress) sealed() {}

type Counter64 struct {
	Value uint64
}

// NewCounter64 creates a new Counter64 value.
func NewCounter64(value uint64) Counter64 {
	return Counter64{Value: value}
}

func (c Counter64) String() string {
	return strconv.FormatUint(c.Value, 10)
}
func (c Counter64) Type() Type {
	return Counter64Type{}
}
func (c Counter64) RawValue() any {
	return c.Value
}
func (c Counter64) sealed() {}

type UInteger32 struct {
	Value uint32
}

// NewUInteger32 creates a new UInteger32 value.
func NewUInteger32(value uint32) UInteger32 {
	return UInteger32{Value: value}
}

func (u UInteger32) String() string {
	return strconv.FormatUint(uint64(u.Value), 10)
}
func (u UInteger32) Type() Type {
	return UInteger32Type{}
}
func (u UInteger32) RawValue() any {
	return u.Value
}
func (u UInteger32) sealed() {}

type OpaqueFloat struct {
	Value float32
}

// NewOpaqueFloat creates a new OpaqueFloat value.
func NewOpaqueFloat(value float32) OpaqueFloat {
	return OpaqueFloat{Value: value}
}

func (o OpaqueFloat) String() string {
	return strconv.FormatFloat(float64(o.Value), 'f', -1, 32)
}
func (o OpaqueFloat) Type() Type {
	return OpaqueFloatType{}
}
func (o OpaqueFloat) RawValue() any {
	return o.Value
}
func (o OpaqueFloat) sealed() {}

type OpaqueDouble struct {
	Value float64
}

// NewOpaqueDouble creates a new OpaqueDouble value.
func NewOpaqueDouble(value float64) OpaqueDouble {
	return OpaqueDouble{Value: value}
}

func (o OpaqueDouble) String() string {
	return strconv.FormatFloat(o.Value, 'f', -1, 64)
}
func (o OpaqueDouble) Type() Type {
	return OpaqueDoubleType{}
}
func (o OpaqueDouble) RawValue() any {
	return o.Value
}
func (o OpaqueDouble) sealed() {}

type TimeTicks struct {
	Value time.Duration
}

// NewTimeTicks creates a new TimeTicks value.
func NewTimeTicks(value time.Duration) TimeTicks {
	return TimeTicks{Value: value}
}

// NewTimeTicksFromHundredths creates a new TimeTicks value from the given value in hundredths of a second, i.e.
// the representation used in SNMP.
func NewTimeTicksFromHundredths(value uint32) TimeTicks {
	return NewTimeTicks(time.Duration(value) * 100 * time.Millisecond)
}

func (t TimeTicks) String() string {
	return t.Value.String()
}
func (t TimeTicks) Type() Type {
	return TimeTicksType{}
}
func (t TimeTicks) RawValue() any { return t.Value }
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
func decodeBERValue(b []byte) (Value, error) {
	if len(b) == 0 {
		return nil, fmt.Errorf("no data")
	}

	switch typeTag(b[0]) {
	case tagInteger:
		// Integer is internally represented as int32, so we decode it as int32 and convert it to int.
		return tryDecode(b[1:], decodeBERInt32, func(t int32) Integer {
			return NewInteger(int(t))
		})
	case tagUTnteger32:
		return tryDecode(b[1:], decodeBERUint32, NewUInteger32)
	case tagOctetString:
		return tryDecode(b[1:], decodeBERBytes, NewOctetString)
	case tagObjectIdentifier:
		return tryDecode(b[1:], decodeBERObjectIdentifier, NewObjectIdentifier)
	case tagIpAddress:
		return tryDecode(b[1:], decodeBERIpAddress, NewIPAddress)
	case tagCounter32:
		return tryDecode(b[1:], decodeBERUint32, NewCounter32)
	case tagGauge32:
		return tryDecode(b[1:], decodeBERUint32, NewUInteger32)
	case tagTimeTicks:
		return tryDecode(b[1:], decodeBERUint32, NewTimeTicksFromHundredths)
	case tagOpaque:
		return tryDecode(b[1:], decodeBERBytes, NewOpaque)
	case tagCounter64:
		return tryDecode(b[1:], decodeBERUint64, NewCounter64)
	case tagNull:
		return Null{}, nil
	case tagNoSuchObject:
		return NoSuchObject{}, nil
	case tagNoSuchInstance:
		return NoSuchInstance{}, nil
	case tagEndOfMibView:
		return EndOfMibView{}, nil
	}

	return nil, fmt.Errorf("unsupported type tag %d", b[0])
}

// tryDecode tries to decode the given BER-encoded value using the given decode function and creates a new Value
// using the given value constructor.
// This is a helper function to reduce code duplication in decodeBERValue.
func tryDecode[T any, V Value](b []byte, decodeFn func([]byte) (T, error), valueConstr func(T) V) (Value, error) {
	v, err := decodeFn(b)
	if err != nil {
		return nil, err
	}

	return valueConstr(v), nil
}

// decodeBERLength decodes the length of a BER-encoded value.
// It is assumed that the first byte of the input is the first byte of the length field.
//
// The length is encoded as a variable-length integer, which may be in one of the following forms:
//   - Short form: The length is encoded in the lower 7 bits of the byte. The most significant bit is 0.
//   - Long form: The lower 7 bits of the byte are the number of bytes used to encode the length. The most significant
//     bit is 1. The following bytes encode the length as a big-endian integer.
//
// Returns the length and the number of bytes used to encode the length as an offset to facilitate skipping over the
// length field.
// Adapted from https://github.com/gosnmp/gosnmp
func decodeBERLength(b []byte) (length int, offset int, _ error) {
	if len(b) == 0 {
		return 0, 0, fmt.Errorf("no data")
	}

	// Short form - interpret the byte as the length and return it
	if b[0] <= 0x7F {
		return int(b[0]), 1, nil
	}

	// Long form - interpret the byte as the number of bytes used to encode the length
	// Then read the following bytes as the length
	lengthLength := int(b[0] & 0x7f)
	if lengthLength > len(b)-1 {
		return 0, 0, fmt.Errorf("length is longer than the remaining data: %d > %d", lengthLength, len(b)-1)
	}

	for i := 0; i < lengthLength; i++ {
		// Left shift the length by 8 bits to  and add the next byte
		length = length<<8 | int(b[i+1])
	}

	// Check for overflow and too long length
	if length < 0 {
		return 0, 0, fmt.Errorf("negative length (overflow): %d", length)
	} else if length > len(b)-1-lengthLength {
		return 0, 0, fmt.Errorf("length is longer than the remaining data: %d > %d", length, len(b)-1-lengthLength)
	}

	return length, lengthLength + 1, nil // +1 for the first byte indicating the long form
}

func decodeBERInt32(b []byte) (int32, error) {
	return 0, errors.New("not implemented")
}

func decodeBERUint32(b []byte) (uint32, error) {
	return 0, errors.New("not implemented")
}

func decodeBERUint64(b []byte) (uint64, error) {
	return 0, errors.New("not implemented")
}

func decodeBERBytes(b []byte) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func decodeBERObjectIdentifier(b []byte) (OID, error) {
	return nil, errors.New("not implemented")
}

func decodeBERIpAddress(b []byte) (net.IP, error) {
	return nil, errors.New("not implemented")
}
