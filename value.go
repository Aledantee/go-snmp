package snmp

import (
	"errors"
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"net"
	"strconv"
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

// decodeBERValue parses a BER-encoded value (Varbind) the given BER-encoded value.
// NoSuchObject, NoSuchInstance, and EndOfMibView are not valid values and return an error.
func decodeBERValue(b []byte) (Value, error) {
	if len(b) == 0 {
		return nil, fmt.Errorf("no data")
	}

	switch typeTag(b[0]) {
	case tagInteger, tagUTnteger32:
	case tagOctetString:
	case tagNull:
	case tagObjectIdentifier:
	case tagIpAddress:
	case tagCounter32:
	case tagGauge32:
	case tagTimeTicks:
	case tagOpaque:
	case tagCounter64:
	case tagNoSuchObject:
		return nil, errors.New("no such object")
	case tagNoSuchInstance:
		return nil, errors.New("no such instance")
	case tagEndOfMibView:
		return nil, errors.New("end of MIB view")
	}

	return nil, fmt.Errorf("unsupported type tag %d", b[0])
}
