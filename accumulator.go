package simple_plugin

import (
	"time"
)

type MyAcc struct {
	precision   time.Duration
	onAddFields addFields
}

func NewMyAcc(precision time.Duration, onAddFields addFields) *MyAcc {
	m := MyAcc{
		precision: precision, onAddFields: onAddFields,
	}

	return &m
}
func (ac *MyAcc) AddFields(
	measurement string,
	fields map[string]interface{},
	tags map[string]string,
	t ...time.Time,
) {
	rounded := ac.getTime(t)
	ac.onAddFields(measurement, tags, fields, rounded)
}

func (ac *MyAcc) AddGauge(
	measurement string,
	fields map[string]interface{},
	tags map[string]string,
	t ...time.Time,
) {
	rounded := ac.getTime(t)
	ac.onAddFields(measurement, tags, fields, rounded)
}

func (ac *MyAcc) AddCounter(
	measurement string,
	fields map[string]interface{},
	tags map[string]string,
	t ...time.Time,
) {
	rounded := ac.getTime(t)
	ac.onAddFields(measurement, tags, fields, rounded)
}

func (ac *MyAcc) AddSummary(
	measurement string,
	fields map[string]interface{},
	tags map[string]string,
	t ...time.Time,
) {
	rounded := ac.getTime(t)
	ac.onAddFields(measurement, tags, fields, rounded)
}

func (ac *MyAcc) AddHistogram(
	measurement string,
	fields map[string]interface{},
	tags map[string]string,
	t ...time.Time,
) {
	rounded := ac.getTime(t)
	ac.onAddFields(measurement, tags, fields, rounded)
}

// AddError passes a runtime error to the accumulator.
// The error will be tagged with the plugin name and written to the log.
func (ac *MyAcc) AddError(err error) {
	if err == nil {
		return
	}

}

func (ac *MyAcc) SetPrecision(precision time.Duration) {
	ac.precision = precision
}

func (ac *MyAcc) getTime(t []time.Time) time.Time {
	var timestamp time.Time
	if len(t) > 0 {
		timestamp = t[0]
	} else {
		timestamp = time.Now()
	}
	return timestamp.Round(ac.precision)
}
