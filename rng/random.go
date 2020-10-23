package rng

import (
	"math"
	"math/rand"
	"time"

	"rocket"
)

func New() *Rng {
	rand.Seed(time.Now().UnixNano())
	return &Rng{}
}

type Rng struct{}

func (Rng) GenerateBust() rocket.Bust {
	base := math.Floor(rand.Float64() * 1000) // 0 ~ 999
	value := base*0.01 + 1                    // 0~999 => 0 ~ 9.99 => 1 ~ 10.99
	d := time.Duration(value*1000) * time.Millisecond
	return rocket.Bust{
		Value:    float32(value),
		Duration: d,
	}
}
