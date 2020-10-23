package rng

import (
	"log"
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
	base := rand.Float64() * 1000     // 0 ~ 999.999999
	value := math.Floor(base)/100 + 1 // 0~999 / 100 = 0 ~ 9.99 , +1 -> 1 ~ 10.99
	d := time.Duration(value*1000) * time.Millisecond
	log.Println(value, d)
	return rocket.Bust{
		Value:    float32(value),
		Duration: d,
	}
}
