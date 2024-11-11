package snmp

import (
	"testing"
)

func TestDecodeBERLength(t *testing.T) {
	var tests = []struct {
		name       string
		args       []byte
		wantLength int
		wantOffset int
		wantErr    bool
	}{
		{"Valid_Short", []byte{0x05, 0x00, 0x00, 0x00, 0x00, 0x00}, 5, 1, false},
		{"Valid_Long", []byte{0x81, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00}, 5, 2, false},
		{"Invalid_NoData", []byte{}, 0, 0, true},
		{"Invalid_LongTruncated", []byte{0x81, 0x05, 0x00}, 0, 0, true},
		{"Invalid_LongNegative", []byte{0x81, 0x05, 0x81, 0x00}, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLength, gotOffset, err := decodeBERLength(tt.args)

			if tt.wantErr && err == nil {
				t.Errorf("want error, but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("want nil, but got error: %s", err)
			}
			if gotLength != tt.wantLength {
				t.Errorf("want length %d, but got %d", tt.wantLength, gotLength)
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("want offset %d, but got %d", tt.wantOffset, gotOffset)
			}
		})
	}
}

func TestDecodeBERUint(t *testing.T) {
	var tests = []struct {
		name    string
		args    []byte
		size    int
		want    uint64
		wantErr bool
	}{
		{"Valid_Short", []byte{0x01, 0x05}, 1, 5, false},
		{"Valid_Long", []byte{0x81, 0x05, 0x57, 0xB4, 0xCB, 0xE8, 0xF0}, 5, 376695417072, false},
		{"Invalid_ZeroLength", []byte{0x00}, 1, 0, true},
		{"Invalid_LargeInteger", []byte{0x04, 0x80, 0x80, 0x80, 0x80}, 4, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeBERUint(tt.args, tt.size)
			if tt.wantErr && (err == nil) {
				t.Errorf("decodeBERUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (err != nil) {
				t.Errorf("decodeBERUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeBERUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
