package models

import (
	"time"

	"github.com/BackendTest/util"
	"github.com/mcuadros/go-defaults"
)

type Tick struct {
	Time   time.Time `json:"time"`
	Symbol string    `json:"symbol"`
	Open   float32   `json:"open" default:"100.00"`
	High   float32   `json:"high" default:"100.00"`
	Low    float32   `json:"low" default:"100.00"`
	Close  float32   `json:"close" default:"100.00"`
	Volume uint32    `json:"volume" default:"10000"`
}

// creates a new tick
func NewTick(symbol string) *Tick {
	t := &Tick{
		Symbol: symbol,
		Time:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	defaults.SetDefaults(t)
	return t
}

// updates the tick
func (tick *Tick) UpdateTick() {
	tick.Time = time.Now().UTC()
	closePriceMin := int(tick.Close - (tick.Close * .1))
	closePriceMax := int(tick.Close + (tick.Close * .1))
	tick.Close = tick.Close + float32(util.RandomInt(closePriceMin, closePriceMax))

	if tick.Close > tick.High {
		tick.High = tick.Close
	}

	if tick.Close < tick.Low {
		tick.Low = tick.Close
	}

	tick.Volume = tick.Volume + uint32(util.RandomInt(0, 1000))

}
