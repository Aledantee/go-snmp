package snmp

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// OID is an SMIv2 Object Identifier.
type OID []int

// Equals returns true if the OID is equal to the other OID.
// Two OIDs are equal if they have the same length and the same values at each sub-identifier.
func (o OID) Equals(other OID) bool {
	if len(o) != len(other) {
		return false
	}

	for i, v := range o {
		if v != other[i] {
			return false
		}
	}

	return true
}

// IsScalar returns true if the OID is a scalar OID.
// A scalar OID is an OID with at least two sub-identifiers where the last sub-identifier is 0.
func (o OID) IsScalar() bool {
	return len(o) > 1 && o[len(o)-1] == 0
}

// IsRootOf returns true iff the OID is the root of the other OID.
// An OID is the root of another OID if it is a prefix of the other OID.
func (o OID) IsRootOf(other OID) bool {
	if len(o) > len(other) {
		return false
	}

	for i, v := range o {
		if v != other[i] {
			return false
		}
	}

	return true
}

// IsChildOf returns true iff the OID is a child of the other OID.
// An OID is a child of another OID if the other OID is a prefix of the OID.
// Shorthand for other.RootOf(o).
func (o OID) IsChildOf(other OID) bool {
	return other.IsRootOf(o)
}

// IsValid returns true iff the OID is a valid OID.
func (o OID) IsValid() bool {
	return o.Validate() == nil
}

// Validate returns an error if the OID is not a valid OID.
// If the OID contains multiple errors, a multi-error is returned containing all errors.
func (o OID) Validate() error {
	if len(o) < 2 {
		return errors.New("must have at least two sub-identifiers")
	}

	var errs []error
	for i, v := range o {
		switch {
		case i == 0 && v > 2:
			errs = append(errs, errors.New("first sub-identifier must be 0, 1, or 2"))
		case v < 0:
			errs = append(errs, fmt.Errorf("sub-identifier %d must be non-negative", i))
		case v > math.MaxUint32:
			errs = append(errs, fmt.Errorf("sub-identifier %d must be less than or equal to %d", i, math.MaxUint32))
		}
	}

	return errors.Join(errs...)
}

// String returns the string representation of the OID.
func (o OID) String() string {
	var sb strings.Builder
	for i, v := range o {
		if i > 0 {
			sb.WriteRune('.')
		}

		sb.WriteString(strconv.Itoa(v))
	}

	return sb.String()
}

// ParseOID attempts to parse the string as an OID.
// The string must be a sequence of sub-identifiers separated by dots.
func ParseOID(s string) (OID, error) {
	var (
		oid OID
		v   int
	)
	for i, c := range s {
		switch {
		case c == '.':
			if i == len(s)-1 {
				return nil, fmt.Errorf("trailing dot at character index %d", i)
			}

			// Ignore leading dot
			if i > 0 {
				if s[i-1] == '.' {
					return nil, fmt.Errorf("consecutive dot at character index %d", i)
				}

				oid = append(oid, v)
				v = 0
			}
		case c >= '0' && c <= '9':
			// Add digit to value by multiplying by 10 and adding the digit
			v = v*10 + int(c-'0')
		default:
			return nil, fmt.Errorf("invalid character at position %d", i)
		}
	}

	// Last sub-identifier is not added in the loop, so add it here
	oid = append(oid, v)
	return oid, oid.Validate()
}
