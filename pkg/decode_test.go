package pkg

import "testing"

func TestDecodeRejectsTruncatedData(t *testing.T) {
	for _, size := range []int{0, 1, encodedSize - 1} {
		if _, err := Decode(make([]byte, size)); err == nil {
			t.Fatalf("expected error for %d bytes", size)
		}
	}
}

func FuzzDecodeTruncated(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{1, 2, 3})
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) >= encodedSize {
			t.Skip()
		}
		if _, err := Decode(data); err == nil {
			t.Fatal("truncated data decoded")
		}
	})
}
