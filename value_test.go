package snmp

import (
	"bytes"
	"testing"
)

func TestDecodeBERLength(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expLen   int
		expTail  []byte
		expError bool
	}{
		{
			name:    "Valid Short Form",
			input:   []byte{0x7F, 0x01},
			expLen:  127,
			expTail: []byte{0x01},
		},
		{
			name:    "Valid Long Form",
			input:   []byte{0x81, 0x01, 0x02},
			expLen:  1,
			expTail: []byte{0x02},
		},
		{
			name:    "Valid Indefinite Form",
			input:   []byte{0x80, 0x01, 0x02},
			expLen:  -1,
			expTail: []byte{0x01, 0x02},
		},
		{
			name:     "Invalid Empty",
			input:    []byte{},
			expError: true,
		},
		{
			name:     "Invalid Long Form Truncated Length",
			input:    []byte{0x82, 0x01},
			expError: true,
		},
		{
			name:     "Invalid Long Form Overflow",
			input:    []byte{0x88, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01},
			expError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length, tail, err := decodeBERLength(tt.input)

			if !tt.expError {
				if length != tt.expLen {
					t.Errorf("expected length %d, got %d", tt.expLen, length)
				}
				if !bytes.Equal(tail, tt.expTail) {
					t.Errorf("expected tail %v, got %v", tt.expTail, tail)
				}
			} else if err == nil {
				t.Errorf("expected error, got no error and length %d with tail %v", length, tail)
			}
		})
	}
}

func TestDecodeBERInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expValue int32
		expTail  []byte
		expError bool
	}{
		{
			name:     "Valid Positive",
			input:    []byte{0x01, 0x7F, 0x75},
			expValue: 127,
			expTail:  []byte{0x75},
		},
		{
			name:     "Valid Negative",
			input:    []byte{0x01, 0xFF},
			expValue: -1,
			expTail:  []byte{},
		},
		{
			name:     "Invalid Truncated",
			input:    []byte{0x01},
			expError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotTail, err := decodeBERInt32(tt.input)

			if !tt.expError {
				if gotValue != tt.expValue {
					t.Errorf("expected value %d, got %d", tt.expValue, gotValue)
				}
				if !bytes.Equal(gotTail, tt.expTail) {
					t.Errorf("expected tail %v, got %v", tt.expTail, gotTail)
				}
			} else if err == nil {
				t.Errorf("expected error, got no error and value %d with tail %v", gotValue, gotTail)
			}
		})
	}
}
