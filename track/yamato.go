package track

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"time"
)

const (
	yamatoendpoint = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"
)

func init() {
	companies["yamato"] = &company{
		companystat: &companystat{
			Key:     "yamato",
			Shorten: "y",
		},
		Trackfunc: Yamatotrack,
	}
}

func Yamatotrack(tn string) ([]Stat, error) {
	val := url.Values{}
	val.Add("number00", "1")
	val.Add("number01", tn)

	res, err := req(yamatoendpoint, val)
	if err != nil {
		return nil, err
	}

	nimotus := []Stat{}

	table := res.Find("div .tracking-invoice-block-detail li")

	if table == nil {
		return nil, errors.New("no such a record")
	}

	table.Each(func(i int, s *goquery.Selection) {
		item := s.Find("div .item").Text()
		date := s.Find("div .date").Text()
		name := s.Find("div .name").Text()

		pt, _ := time.Parse("01月02日 15:04", date)
		date = pt.Format("01/02 15:04")

		nimotus = append(nimotus, Stat{
			Date:   date,
			Status: item,
			Office: name,
		})
	})

	if len(nimotus) == 0 {
		return nil, errors.New("no such a record")
	}

	return nimotus, nil
}
