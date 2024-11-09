package snmp

import (
	"testing"
)

func TestOID_Equals(t *testing.T) {
	cases := []struct {
		name     string
		oid1     OID
		oid2     OID
		expected bool
	}{
		{
			name:     "EqualOIDs",
			oid1:     OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			oid2:     OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			expected: true,
		},
		{
			name:     "NotEqualOIDs",
			oid1:     OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			oid2:     OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 2},
			expected: false,
		},
		{
			name:     "PartialEqualOIDs",
			oid1:     OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			oid2:     OID{1, 3, 6, 1, 2, 1, 2, 2, 1},
			expected: false,
		},
		{
			name:     "EmptyOIDs",
			oid1:     OID{},
			oid2:     OID{},
			expected: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.oid1.Equals(tc.oid2)
			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}

func TestOID_IsScalar(t *testing.T) {
	cases := []struct {
		name     string
		oid      OID
		expected bool
	}{
		{
			name:     "Scalar",
			oid:      OID{1, 3, 6, 1, 2, 0},
			expected: true,
		},
		{
			name:     "NonScalar",
			oid:      OID{1, 3, 6, 1, 2, 1},
			expected: false,
		},
		{
			name:     "Empty",
			oid:      OID{},
			expected: false,
		},
		{
			name:     "SingleSubIdentifier",
			oid:      OID{1},
			expected: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.oid.IsScalar()

			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}

func TestOID_IsRootOf(t *testing.T) {
	cases := []struct {
		name     string
		root     OID
		oid      OID
		expected bool
	}{
		{
			name:     "IsRootOfOID",
			root:     OID{1, 3, 6},
			oid:      OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			expected: true,
		},
		{
			name:     "IsNotRootOfOID",
			root:     OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			oid:      OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 2},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.root.IsRootOf(tc.oid)
			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}

func TestOID_IsChildOf(t *testing.T) {
	cases := []struct {
		name     string
		child    OID
		parent   OID
		expected bool
	}{
		{
			name:     "OIDIsChildOfParent",
			child:    OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			parent:   OID{1, 3, 6},
			expected: true,
		},
		{
			name:     "OIDIsNotChildOfParent",
			child:    OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 2},
			parent:   OID{1, 3, 6, 1, 2, 1, 2, 2, 1, 1},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.child.IsChildOf(tc.parent)
			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}

func TestOID_String(t *testing.T) {
	cases := []struct {
		name     string
		oid      OID
		expected string
	}{
		{
			name:     "OIDToString",
			oid:      OID{1, 3, 6, 1, 2, 1},
			expected: "1.3.6.1.2.1",
		},
		{
			name:     "EmptyOIDToString",
			oid:      OID{},
			expected: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.oid.String()
			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}

// defining test cases for ParseOID function
func TestParseOID(t *testing.T) {
	cases := []struct {
		name     string
		oidStr   string
		expected OID
	}{
		{
			name:     "Valid",
			oidStr:   "1.3.6.1.2.1",
			expected: OID{1, 3, 6, 1, 2, 1},
		},
		{
			name:     "Valid_LeadingDot",
			oidStr:   ".1.3",
			expected: OID{1, 3},
		},
		{
			name:     "Invalid_TrailingDot",
			oidStr:   "1.3.6.1.2.1.",
			expected: nil,
		},
		{
			name:     "Invalid_EmptyString",
			oidStr:   "",
			expected: nil,
		},
		{
			name:     "Invalid_SingleSubIdentifier",
			oidStr:   "1",
			expected: nil,
		},
		{
			name:     "Invalid_ConsecutiveDots",
			oidStr:   "1..3",
			expected: nil,
		},
		{
			name:     "Invalid_NegativeSubIdentifier",
			oidStr:   "1.-3",
			expected: nil,
		},
		{
			name:     "Invalid_TooLargeSubIdentifier",
			oidStr:   "1.4294967296",
			expected: nil,
		},
		{
			name:     "Invalid_WrongFirstSubIdentifier",
			oidStr:   "3.1.2",
			expected: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseOID(tc.oidStr)

			if tc.expected == nil {
				if err == nil {
					t.Errorf("want error for %q, got nil", tc.oidStr)
				}
			} else {
				if !tc.expected.Equals(got) {
					t.Errorf("got %v; want %v", got, tc.expected)
				}
			}
		})
	}
}

// defining test cases for IsValid() function
func TestOID_IsValid(t *testing.T) {
	cases := []struct {
		name     string
		oid      OID
		expected bool
	}{
		{
			name:     "ValidOID",
			oid:      OID{1, 3, 6, 1, 2, 1},
			expected: true,
		},
		{
			name:     "InvalidOID",
			oid:      OID{1},
			expected: false,
		},
		{
			name:     "EmptyOID",
			oid:      OID{},
			expected: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.oid.IsValid()
			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}
