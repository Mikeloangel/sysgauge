package memstorage

import (
	"errors"

	"github.com/Mikeloangel/sysgauge/internal/metrixunits"
)

type Memstorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

var mem *Memstorage

func InitMemstorage() {
	if mem != nil {
		return
	}

	mem = &Memstorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func Update(mu metrixunits.MetricUnit) {
	switch mu.Type {
	case metrixunits.Counter:
		updateCounter(mu.Name, *mu.ValueI)
	case metrixunits.Gauge:
		updateGauge(mu.Name, *mu.ValueF)
	}
}

func updateCounter(key string, value int64) {
	_, ok := mem.counter[key]
	if !ok {
		mem.counter[key] = value
	} else {
		mem.counter[key] += value
	}
}

func updateGauge(key string, value float64) {
	mem.gauge[key] = value
}

func GetCounter(key string) (int64, error) {
	value, ok := mem.counter[key]
	if !ok {
		return 0, errors.New("key is not set")
	}

	return value, nil
}

func GetGauge(key string) (float64, error) {
	value, ok := mem.gauge[key]
	if !ok {
		return 0, errors.New("key is not set")
	}

	return value, nil
}
