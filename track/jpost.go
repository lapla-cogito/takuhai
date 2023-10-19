package track

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
	"strings"
)

const (
	jpostendpoint = "https://trackings.post.japanpost.jp/services/srv/search/direct"
	fieldMax      = 6
)

func removeConsecutiveSpace(str string) string {
	str = strings.TrimSpace(str)
	rep := regexp.MustCompile(`[\s　]+`)
	return rep.ReplaceAllString(str, " ")
}

func init() {
	companies["jpost"] = &company{
		companystat: &companystat{
			Key:     "jpost",
			Shorten: "j",
		},
		Trackfunc: JPosttrack,
	}
}

func JPosttrack(tn string) ([]Stat, error) {
	val := url.Values{}
	val.Add("searchKind", "S002")
	val.Add("locale", "ja")
	val.Add("reqCodeNo1", tn)

	doc, err := req(jpostendpoint, val)
	if err != nil {
		return nil, err
	}

	nimotus := []Stat{}
	field := [fieldMax]string{}

	doc.Find("[summary='履歴情報'] td").Each(func(i int, s *goquery.Selection) {
		if (i+1)%fieldMax == 0 {
			nimotus = append(nimotus, Stat{
				Date:   field[0],
				Status: field[1],
				Office: removeConsecutiveSpace(field[4] + " " + field[3]),
			})
		}
		field[i%fieldMax] = s.Text()
	})

	if len(nimotus) == 0 {
		return nil, errors.New("no such a record")
	}

	return nimotus, nil
}
