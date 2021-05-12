package randomnumber

import (
	"math/rand"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Ranno struct {
	ran int
}

func (s *Ranno) SampleConfig() string {
	return ""
}

func (s *Ranno) Description() string {
	return "Inserts randomnumber for Poc"
}

func (s *Ranno) Gather(acc telegraf.Accumulator) error {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	ran := y1.Intn(time.Now().Hour())

	fields := make(map[string]interface{})
	fields["rand"] = ran

	tags := make(map[string]string)

	acc.AddFields("randomno", fields, tags)

	return nil
}

func init() {
	inputs.Add("randomnumber", func() telegraf.Input { return &Ranno{ran: 0} })
}
