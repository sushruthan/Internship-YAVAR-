package number

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type ran struct {
	x float64
}

var ranConfig = `
  ## Set the value
  value = 10.0
`

func (s *ran) SampleConfig() string {
	return ranConfig
}
func (s *ran) Description() string {
	return "inputs random number"
}

func (s *ran) Gather(acc telegraf.Accumulator) error {
	//r := rand.Float64()

	fields := make(map[string]interface{})

	fields["rand"] = 10.0

	tags := make(map[string]string)
	tags["random number"] = "number"

	acc.AddFields("random-number", fields, tags)

	return nil
}

func init() {
	inputs.Add("number", func() telegraf.Input {
		return &ran{
			x: 0.0,
		}
	})
}





