package metrixunits

type MetricType string

const (
	Gauge   MetricType = "gauge"
	Counter MetricType = "counter"
)

type MetricUnit struct {
	Type   MetricType
	Name   string
	ValueI *int64
	ValueF *float64
}
