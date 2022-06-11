package api

import (
	"github.com/BackendTest/models"
	"github.com/BackendTest/util"
)

var Cache = make([]*models.Tick, 0, 10)

// initialize the cache with 10 symbols
func InitializeCache() {
	for i := 0; i < 10; i++ {
		symbol := util.RandomString(util.SymbolLength)
		tick := models.NewTick(symbol)
		Cache = append(Cache, tick)
	}
}
