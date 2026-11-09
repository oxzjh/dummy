package dummy

import "math"

type Audio struct {
	iWave []int16
	fWave []float32
	p     int
	n     int
}

func (a *Audio) Read(sample []int16) {
	for i := range sample {
		sample[i] = a.iWave[a.p]
		a.move()
	}
}

func (a *Audio) ReadFloat32(sample []float32) {
	for i := range sample {
		sample[i] = a.fWave[a.p]
		a.move()
	}
}

func (a *Audio) move() {
	a.p++
	if a.p >= a.n {
		a.p = 0
	}
}

func NewAudio(sampleRate, frequency int, volume float64) *Audio {
	n := sampleRate / frequency
	iWave := make([]int16, n)
	fWave := make([]float32, n)
	for i := 0; i < n; i++ {
		wave := math.Sin(2 * math.Pi * float64(i) / float64(n))
		iWave[i] = int16(wave * volume * 0x7fff)
		fWave[i] = float32(wave * volume)
	}
	return &Audio{iWave: iWave, fWave: fWave, n: n}
}

func NewDefaultAudio() *Audio {
	return NewAudio(16000, 400, 0.25)
}
