package fuzz

import (
	"bytes"
	"testing"
)

func TestPanicNegativeLength(t *testing.T) {
	raw := []byte("\x10\a@\x01\x11\x01'(\xfc0000000000000000000")
	if _, err := Parse(raw); err == nil {
		t.Fatal("Parsed invalid input")
	}
}

func TestPanicTooShort(t *testing.T) {
	raw := []byte("\x00")
	if _, err := Parse(raw); err == nil {
		t.Fatal("Parsed invalid input")
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		entries []Entry
	}{
		{
			[]Entry{
				{Data: []byte("This")},
				{Data: []byte("is")},
				{Data: []byte("a")},
				{Data: []byte("test")},
			},
		},
	}
	for _, test := range tests {
		l := &List{Entries: test.entries}
		packed, err := l.Pack()
		if err != nil {
			t.Fatal("Unable to pack", err)
		}
		parsed, err := Parse(packed)
		if err != nil {
			t.Fatal("Unable to parse")
		}
		for i, entry := range test.entries {
			if !bytes.Equal(entry.Data, parsed.Entries[i].Data) {
				t.Fatal("Unequal")
			}
		}
	}
}
