package randomNumber

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type random struct {
}

func (s *random) SampleConfig() string {
	return ""
}

func (s *random) Description() {
	return ""
}

func (s *random) Gather(acc telegraf.Accumulator) error {
	return nil
}
func init() {
	inputs.Add("random", func() telegraf.Input { return &random{} })
}
