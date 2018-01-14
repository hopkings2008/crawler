package icangbao

import (
	//"fmt"

	"github.com/hopkings/crawler/parser"
)

type FactoryIcangbao struct {
	HostPrefix string
}

func (f5 *FactoryIcangbao) Build() (parser.Parser, error) {
	p := &IcangbaoParser{}
	return p, nil
}

func init() {
	f := &FactoryIcangbao{
		HostPrefix: "icangbao",
	}
	parser.RegisterFactory(f.HostPrefix, f)
}
