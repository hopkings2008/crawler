package icangbao

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/hopkings/crawler/parser"
)

type IcangbaoParser struct {
}

func (yp *IcangbaoParser) Parse(doc *goquery.Document) (*parser.WarehouseInfo, error) {
	//get the location
	whi := parser.CreateWarehouseInfo()
	whi.IsValid = 1
	doc.Find(".cbzx_tao6_r p").Each(func(i int, sel *goquery.Selection) {
		/*sel.Find(".").Each(func(ii int, addr *goquery.Selection) {
			for _, p := range addr.Nodes {
				n := p.FirstChild
				fmt.Printf("warehouse: 仓库地址: %s", n.Data)
			}
		})*/
		udata := mahonia.NewDecoder("gbk").ConvertString(sel.Text())
		udata = strings.Replace(udata, " ", "", -1)
		udata = strings.Replace(udata, "\n", " ", -1)
		//elems := strings.Split(udata, "\n")
		fmt.Printf("%s ", udata)
	})

	doc.Find(".cb210_c2").Each(func(i int, sel *goquery.Selection) {
		sel.Find("p").Each(func(ii int, elems *goquery.Selection) {
			udata := mahonia.NewDecoder("gbk").ConvertString(elems.Text())
			udata = strings.Replace(udata, "\n", "", -1)
			fmt.Printf("%s ", udata)
		})
	})

	fmt.Printf("\n")

	return whi, nil
}
