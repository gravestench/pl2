package pl2

import (
	"image/color"
	"io"

	"github.com/gravestench/pl2/pkg"
)

type (
	PL2       = pkg.PL2
	Transform = pkg.Transform
)

func FromBytes(data []byte) (*PL2, error) {
	return pkg.FromBytes(data)
}

func ToBytes(pl2 *PL2) ([]byte, error) {
	return pkg.ToBytes(pl2)
}

func Decode(rs io.ReadSeeker) (*PL2, error) {
	return pkg.Decode(rs)
}

func EncodePalette(p color.Palette) ([]byte, error) {
	return pkg.EncodePalette(p)
}
