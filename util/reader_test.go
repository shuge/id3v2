package util

import (
	"bytes"
	"testing"
)

var (
	bs   = []byte{0, 11, 22, 33, 44, 55, 77, 88, 55, 55, 66, 77, 88}
	bsRd = bytes.NewReader(bs)
)

func TestReadBytes(t *testing.T) {
	bsRd.Reset(bs)
	bsReader := NewReader(bsRd)
	n := 5 // Read 5 elements

	read, err := bsReader.ReadBytes(n)
	if err != nil {
		t.Error(err)
	}
	if n != len(read) {
		t.Errorf("Expecting to read: %v, got: %v", n, read)
	}
	if !bytes.Equal(bs[:n], read) {
		t.Error("Expecting: %v, got: %v", bs[:n], len(read))
	}

	if bsReader.buf.Buffered() != len(bs)-n {
		t.Errorf("Expecting buffered: %v, got: %v", len(bs)-n, bsReader.buf.Buffered())
	}
}

func TestReadTillDelims(t *testing.T) {
	bsRd.Reset(bs)
	bsReader := NewReader(bsRd)
	delims := []byte{55, 66}
	expectedLen := 9

	read, err := bsReader.ReadTillDelims(delims)
	if err != nil {
		t.Error(err)
	}
	if expectedLen != len(read) {
		t.Errorf("Expecting to read: %v, got: %v", expectedLen, len(read))
	}
	if !bytes.Equal(bs[:expectedLen], read) {
		t.Error("Expecting: %v, got: %v", bs[:expectedLen], read)
	}
}