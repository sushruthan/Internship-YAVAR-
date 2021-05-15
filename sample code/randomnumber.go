package randnumber

import (
	"math/rand"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type ran struct {
	x int
}

var ranConfig = `
  ## set sample configuration value to 10.0
  value = 10.0
`

func (s *ran) SampleConfig() string {
	return ranConfig
}
func (s *ran) Description() string {
	return "inputs random number"
}

func (s *ran) Gather(acc telegraf.Accumulator) error {
	r := rand.Float64()

	fields := make(map[string]interface{})
	fields["rand"] = r

	tags := make(map[string]string)

	acc.AddFields("random-number", fields, tags)

	return nil
}

func init() {
	inputs.Add("randnumber", func() telegraf.Input {
		return &ran{
			x: 0,
		}
	})
}
