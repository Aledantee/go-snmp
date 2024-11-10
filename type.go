package snmp

import "fmt"

// Type is the interface implemented by all supported SNMP types. Sealed.
type Type interface {
	// Stringer is implemented by all SNMP types.
	// The implementation should return a string representation of the type, including constraints if the type
	// is a subtype of a base type.
	fmt.Stringer

	// Tag returns the ASN.1 tag of the type.
	Tag() int

	// Name returns the name of the type.
	Name() string

	// Base returns the base type of the type.
	// If the type is a base type, the method returns itself.
	Base() Type

	// ValidateValue validates the value against the type to ensure that the value is of the correct type and
	// satisfies any constraints imposed by the type.
	ValidateValue(v Value) error

	// sealed is a marker method to prevent implementation of the interface outside of this package.
	// Does not do anything.
	sealed()
}

type BooleanType struct{}

func (t BooleanType) String() string {
	return "BOOLEAN"
}
func (t BooleanType) Tag() int {
	return int(tagBoolean)
}
func (t BooleanType) Name() string {
	return "BOOLEAN"
}
func (t BooleanType) Base() Type {
	return t
}
func (t BooleanType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t BooleanType) sealed() {}

type IntegerType struct{}

func (t IntegerType) String() string {
	return "INTEGER"
}
func (t IntegerType) Tag() int {
	return int(tagInteger)
}
func (t IntegerType) Name() string {
	return "INTEGER"
}
func (t IntegerType) Base() Type {
	return t
}
func (t IntegerType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t IntegerType) sealed() {}

type BitStringType struct{}

func (t BitStringType) String() string {
	return "BIT STRING"
}
func (t BitStringType) Tag() int {
	return int(tagBitString)
}
func (t BitStringType) Name() string {
	return "BIT STRING"
}
func (t BitStringType) Base() Type {
	return t
}
func (t BitStringType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t BitStringType) sealed() {}

type OctetStringType struct{}

func (t OctetStringType) String() string {
	return "OCTET STRING"
}
func (t OctetStringType) Tag() int {
	return int(tagOctetString)
}
func (t OctetStringType) Name() string {
	return "OCTET STRING"
}
func (t OctetStringType) Base() Type {
	return t
}
func (t OctetStringType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t OctetStringType) sealed() {}

type ObjectIdentifierType struct{}

func (t ObjectIdentifierType) String() string {
	return "OBJECT IDENTIFIER"
}
func (t ObjectIdentifierType) Tag() int {
	return int(tagObjectIdentifier)
}
func (t ObjectIdentifierType) Name() string {
	return "OCTET IDENTIFIER"
}
func (t ObjectIdentifierType) Base() Type {
	return t
}
func (t ObjectIdentifierType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t ObjectIdentifierType) sealed() {}

type ObjectDescriptionType struct{}

func (t ObjectDescriptionType) String() string {
	return "OBJECT DESCRIPTION"
}
func (t ObjectDescriptionType) Tag() int {
	return int(tagObjectDescription)
}
func (t ObjectDescriptionType) Name() string {
	return "OCTET DESCRIPTION"
}
func (t ObjectDescriptionType) Base() Type {
	return t
}
func (t ObjectDescriptionType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t ObjectDescriptionType) sealed() {}

type IPAddressType struct{}

func (t IPAddressType) String() string {
	return "IP ADDRESS"
}
func (t IPAddressType) Tag() int {
	return int(tagIpAddress)
}
func (t IPAddressType) Name() string {
	return "IP ADDRESS"
}
func (t IPAddressType) Base() Type {
	return t
}
func (t IPAddressType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t IPAddressType) sealed() {}

type Counter32Type struct{}

func (t Counter32Type) String() string {
	return "COUNTER32"
}
func (t Counter32Type) Tag() int {
	return int(tagIpAddress)
}
func (t Counter32Type) Name() string {
	return "COUNTER32"
}
func (t Counter32Type) Base() Type {
	return t
}
func (t Counter32Type) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t Counter32Type) sealed() {}

type Gauge32Type struct{}

func (t Gauge32Type) String() string {
	return "GAUGE32"
}
func (t Gauge32Type) Tag() int {
	return int(tagGauge32)
}
func (t Gauge32Type) Name() string {
	return "GAUGE32"
}
func (t Gauge32Type) Base() Type {
	return t
}
func (t Gauge32Type) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t Gauge32Type) sealed() {}

type OpaqueType struct{}

func (t OpaqueType) String() string {
	return "OPAQUE"
}
func (t OpaqueType) Tag() int {
	return int(tagOpaque)
}
func (t OpaqueType) Name() string {
	return "OPAQUE"
}
func (t OpaqueType) Base() Type {
	return t
}
func (t OpaqueType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t OpaqueType) sealed() {}

type NsapAddressType struct{}

func (t NsapAddressType) String() string {
	return "NSAP ADDRESS"
}
func (t NsapAddressType) Tag() int {
	return int(tagNsapAddress)
}
func (t NsapAddressType) Name() string {
	return "NSAP ADDRESS"
}
func (t NsapAddressType) Base() Type {
	return t
}
func (t NsapAddressType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t NsapAddressType) sealed() {}

type Counter64Type struct{}

func (t Counter64Type) String() string {
	return "COUNTER64"
}
func (t Counter64Type) Tag() int {
	return int(tagCounter64)
}
func (t Counter64Type) Name() string {
	return "COUNTER64"
}
func (t Counter64Type) Base() Type {
	return t
}
func (t Counter64Type) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t Counter64Type) sealed() {}

type UInteger32Type struct{}

func (t UInteger32Type) String() string {
	return "UINTEGER32"
}
func (t UInteger32Type) Tag() int {
	return int(tagUTnteger32)
}
func (t UInteger32Type) Name() string {
	return "UINTEGER32"
}
func (t UInteger32Type) Base() Type {
	return t
}
func (t UInteger32Type) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t UInteger32Type) sealed() {}

type OpaqueFloatType struct{}

func (t OpaqueFloatType) String() string {
	return "OPAQUE FLOAT"
}
func (t OpaqueFloatType) Tag() int {
	return int(tagOpaqueFloat)
}
func (t OpaqueFloatType) Name() string {
	return "OPAQUE FLOAT"
}
func (t OpaqueFloatType) Base() Type {
	return t
}
func (t OpaqueFloatType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t OpaqueFloatType) sealed() {}

type OpaqueDoubleType struct{}

func (t OpaqueDoubleType) String() string {
	return "OPAQUE DOUBLE"
}
func (t OpaqueDoubleType) Tag() int {
	return int(tagOpaqueDouble)
}
func (t OpaqueDoubleType) Name() string {
	return "OPAQUE DOUBLE"
}
func (t OpaqueDoubleType) Base() Type {
	return t
}
func (t OpaqueDoubleType) ValidateValue(v Value) error {
	return assertSameType(v.Type(), t)
}
func (t OpaqueDoubleType) sealed() {}

type NullType struct{}

func (n NullType) String() string {
	return "NULL"
}
func (n NullType) Tag() int {
	return int(tagNull)
}
func (n NullType) Name() string {
	return "NULL"
}
func (n NullType) Base() Type {
	return n
}

func (n NullType) ValidateValue(v Value) error {
	//TODO implement me
	panic("implement me")
}

func (n NullType) sealed() {
	//TODO implement me
	panic("implement me")
}

type typeTag int

const (
	tagEndOfContents     typeTag = 0x00
	tagUnknownType       typeTag = 0x00
	tagBoolean           typeTag = 0x01
	tagInteger           typeTag = 0x02
	tagBitString         typeTag = 0x03
	tagOctetString       typeTag = 0x04
	tagNull              typeTag = 0x05
	tagObjectIdentifier  typeTag = 0x06
	tagObjectDescription typeTag = 0x07
	tagIpAddress         typeTag = 0x40
	tagCounter32         typeTag = 0x41
	tagGauge32           typeTag = 0x42
	tagTimeTicks         typeTag = 0x43
	tagOpaque            typeTag = 0x44
	tagNsapAddress       typeTag = 0x45
	tagCounter64         typeTag = 0x46
	tagUTnteger32        typeTag = 0x47
	tagOpaqueFloat       typeTag = 0x78
	tagOpaqueDouble      typeTag = 0x79
	tagNoSuchObject      typeTag = 0x80
	tagNoSuchInstance    typeTag = 0x81
	tagEndOfMibView      typeTag = 0x82
)

func assertSameType(a, b Type) error {
	// We can't compare the types directly because they are interfaces which would only do a pointer comparison.
	// Instead, we compare the tags, which should be unique for each type either way.
	if a.Tag() != b.Tag() {
		return fmt.Errorf("type mismatch: %q (tag: %d) does not match %q (tag: %d)", a, a.Tag(), b, b.Tag())
	}

	return nil
}
