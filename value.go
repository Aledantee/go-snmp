package snmp

import "fmt"

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
