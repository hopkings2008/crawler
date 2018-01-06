package ycparser

import (
	//"fmt"

	"github.com/hopkings/crawler/parser"
)

type FactoryYc struct {
	HostPrefix string
}

func (f5 *FactoryYc) Build() (parser.Parser, error) {
	p := &YcParser{}
	return p, nil
}

func init() {
	f := &FactoryYc{
		HostPrefix: "50yc",
	}
	parser.RegisterFactory(f.HostPrefix, f)
}
