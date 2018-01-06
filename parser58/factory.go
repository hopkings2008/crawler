package parser58

import (
	//"fmt"

	"github.com/hopkings/crawler/parser"
)

type Factory58 struct {
	HostPrefix string
}

func (f5 *Factory58) Build() (parser.Parser, error) {
	p := &Parser58{}
	return p, nil
}

func init() {
	f := &Factory58{
		HostPrefix: "58",
	}
	parser.RegisterFactory(f.HostPrefix, f)
}
