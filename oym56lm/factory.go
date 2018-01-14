package oym56lm

import (
	//"fmt"

	"github.com/hopkings/crawler/parser"
)

type FactoryOym struct {
	HostPrefix string
}

func (f5 *FactoryOym) Build() (parser.Parser, error) {
	p := &OymParser{}
	return p, nil
}

func init() {
	f := &FactoryOym{
		HostPrefix: "oym56lm",
	}
	parser.RegisterFactory(f.HostPrefix, f)
}
