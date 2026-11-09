package dummy

import (
	"image"
	"math/rand"
)

type Video struct {
	width  int
	height int
	wEnd   int
	hEnd   int
	dst    *image.YCbCr
	b      []byte
}

func (v *Video) Read() image.Image {
	for y := v.hEnd; y < v.height; y++ {
		yi := v.width * y
		for x := v.wEnd; x < v.width; x++ {
			if rand.Intn(2) == 0 {
				v.dst.Y[yi+x] = 0
			} else {
				v.dst.Y[yi+x] = 255
			}
		}
	}
	return v.dst
}

func (v *Video) ReadBytes() []byte {
	v.Read()
	copy(v.b, v.dst.Y)
	return v.b
}

func NewVideo(width, height int) *Video {
	colors := [][3]byte{
		{235, 128, 128},
		{210, 16, 146},
		{170, 166, 16},
		{145, 54, 34},
		{107, 202, 222},
		{82, 90, 240},
		{41, 240, 110},
	}

	yi := width * height
	ci := (width / 2) * (height / 2)
	m := &image.YCbCr{
		Y:              make([]uint8, yi),
		Cb:             make([]uint8, ci),
		Cr:             make([]uint8, ci),
		YStride:        width,
		CStride:        width / 2,
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect:           image.Rect(0, 0, width, height),
	}

	hColorBarEnd := height * 3 / 4
	wGradationEnd := width * 5 / 7
	for y := 0; y < hColorBarEnd; y++ {
		yi := width * y
		ci := (width / 2) * (y / 2)
		// Color bar
		for x := 0; x < width; x++ {
			c := x * 7 / width
			m.Y[yi+x] = uint8(uint16(colors[c][0]) * 75 / 100)
			m.Cb[ci+x/2] = colors[c][1]
			m.Cr[ci+x/2] = colors[c][2]
		}
	}
	for y := hColorBarEnd; y < height; y++ {
		yi := width * y
		ci := (width / 2) * (y / 2)
		for x := 0; x < wGradationEnd; x++ {
			// Gray gradation
			m.Y[yi+x] = uint8(x * 255 / wGradationEnd)
			m.Cb[ci+x/2] = 128
			m.Cr[ci+x/2] = 128
		}
		for x := wGradationEnd; x < width; x++ {
			// Noise area
			m.Cb[ci+x/2] = 128
			m.Cr[ci+x/2] = 128
		}
	}

	b := make([]byte, len(m.Y)+len(m.Cb)+len(m.Cr))
	copy(b[len(m.Y):], m.Cb)
	copy(b[len(m.Y)+len(m.Cb):], m.Cr)
	return &Video{width, height, wGradationEnd, hColorBarEnd, m, b}
}
