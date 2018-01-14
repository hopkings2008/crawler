package china

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/hopkings/crawler/parser"
)

type ChinaParser struct {
}

func (yp *ChinaParser) Parse(doc *goquery.Document) (*parser.WarehouseInfo, error) {
	//get the location
	whi := parser.CreateWarehouseInfo()
	whi.IsValid = 1

	doc.Find(".column_r_box .category_an .text").Each(func(i int, sel *goquery.Selection) {
		sel.Find("tr").Each(func(ii int, tr *goquery.Selection) {
			/*for _, p := range addr.Nodes {
				n := p.FirstChild
				fmt.Printf("warehouse: 仓库地址: %s", n.Data)
			}*/

			tr.Find("td").Each(func(itr int, td *goquery.Selection) {
				td.Find("span").Each(func(ispan int, span *goquery.Selection) {
					udata := mahonia.NewDecoder("gbk").ConvertString(span.Text())
					//udata = strings.Replace(udata, " ", "", -1)
					//udata = strings.Replace(udata, "\n", " ", -1)
					strings.Split(udata, "\n")
					fmt.Printf("%s ", udata)
				})
			})
		})
	})

	fmt.Printf("\n")

	return whi, nil
}
