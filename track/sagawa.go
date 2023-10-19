package track

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

const (
	sagawaendpoint = "https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do"
)

func init() {
	companies["sagawa"] = &company{
		companystat: &companystat{
			Key:     "sagawa",
			Shorten: "s",
		},
		Trackfunc: Sagawatrack,
	}
}

func Sagawatrack(tn string) ([]Stat, error) {
	val := url.Values{}
	val.Add("okurijoNo", tn)
	res, _ := req(sagawaendpoint, val)

	nimotus := []Stat{}

	table := res.Find(".table_okurijo_detail2").Eq(1)

	if table == nil {
		return nil, errors.New("no such a record")
	}

	table.Find("tr").Each(
		func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			td := s.Find("td").Map(
				func(_ int, s *goquery.Selection) string {
					return s.Text()
				})

			date := td[1]
			date = strings.Replace(date, "\t", "", -1)
			date = strings.Replace(date, "\n", "", -1)

			message := td[0]
			message = strings.Replace(message, "\t", "", -1)
			message = strings.Replace(message, "\n", "", -1)

			office := td[2]
			office = strings.Replace(office, "\t", "", -1)
			office = strings.Replace(office, "\n", "", -1)

			nimotus = append(nimotus, Stat{
				Date:   date,
				Status: message[3:],
				Office: office,
			})
		})

	if len(nimotus) == 0 {
		return nil, errors.New("no such a record")
	}

	return nimotus, nil
}
