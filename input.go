package simple_plugin

import "time"

type Input interface {

	// Gather takes in an accumulator and adds the metrics that the Input
	// gathers. This is called every "interval"
	Gather(Accumulator) error
}

type Accumulator interface {
	// AddFields adds a metric to the accumulator with the given measurement
	// name, fields, and tags (and timestamp). If a timestamp is not provided,
	// then the accumulator sets it to "now".
	AddFields(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddGauge is the same as AddFields, but will add the metric as a "Gauge" type
	AddGauge(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddCounter is the same as AddFields, but will add the metric as a "Counter" type
	AddCounter(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddSummary is the same as AddFields, but will add the metric as a "Summary" type
	AddSummary(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddHistogram is the same as AddFields, but will add the metric as a "Histogram" type
	AddHistogram(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// SetPrecision sets the timestamp rounding precision.  All metrics addeds
	// added to the accumulator will have their timestamp rounded to the
	// nearest multiple of precision.
	SetPrecision(precision time.Duration)

	// Report an error.
	AddError(err error)
}
