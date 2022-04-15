package sinecosine

import (
	"github.com/hyeongseoknam/simple_plugin"

	"fmt"
	"math"

	"github.com/hyeongseoknam/simple_plugin/plugin/inputs"
)

type SineCosine struct {
	Magnitude float64
	Inc       float64
	x         float64
}

func (sc *SineCosine) Gather(acc simple_plugin.Accumulator) error {
	fields := make(map[string]interface{})
	tags := make(map[string]string)

	tags["Magnitude"] = fmt.Sprint(sc.Magnitude)
	tags["Inc"] = fmt.Sprint(sc.Inc)
	fields["X"] = sc.x
	fields["sine"] = sc.Magnitude * math.Sin(sc.x)
	fields["cosine"] = math.Cos(sc.x)

	acc.AddFields("sinecosine", fields, tags)

	sc.x += sc.Inc

	return nil
}

func init() {
	inputs.Add("sinecosine", func() simple_plugin.Input {
		return &SineCosine{
			Magnitude: 0,
			Inc:       0.1,
		}
	})
}
