package china

import (
	//"fmt"

	"github.com/hopkings/crawler/parser"
)

type FactoryChina struct {
	HostPrefix string
}

func (f5 *FactoryChina) Build() (parser.Parser, error) {
	p := &ChinaParser{}
	return p, nil
}

func init() {
	f := &FactoryChina{
		HostPrefix: "china",
	}
	parser.RegisterFactory(f.HostPrefix, f)
}
