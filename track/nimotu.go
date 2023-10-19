package track

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Stat struct {
	Date   string
	Status string
	Office string
}

type Pkg struct {
	Company  string
	TN       string
	Alias    string
	Timeline []Stat
}

type List struct {
	Pkgs []Pkg
}

func (l *List) Addtoli(packages *Pkg) error {
	for _, p := range l.Pkgs {
		if p.Alias == packages.Alias {
			return errors.New("There is already exist the same name")
		}
	}
	l.Pkgs = append(l.Pkgs, *packages)
	return nil
}

func New() *List {
	return &List{
		Pkgs: []Pkg{},
	}
}

func (l *List) Save(path string) error {
	buf, err := yaml.Marshal(l.Pkgs)
	if err != nil {
		return err
	}

	if path == "" {
		path = filepath.Join(os.Getenv("HOME"), ".takuhai", "list.yml")
	}

	dirpath := filepath.Dir(path)

	if IsExist(path) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return err
		}
		if err := CreateFile(path); err != nil {
			return err
		}
	}

	return os.WriteFile(path, buf, os.ModePerm)
}

func (l *List) Load() error {
	dirpath := filepath.Join(os.Getenv("HOME"), ".takuhai")
	path := filepath.Join(os.Getenv("HOME"), ".takuhai", "list.yml")
	if IsExist(path) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return err
		}
		if err := CreateFile(path); err != nil {
			return err
		}
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buf, &l.Pkgs)
}

func (l *List) Merge(ymlData interface{}) error {
	for _, v := range ymlData.([]interface{}) {
		pkg := v.(map[interface{}]interface{})
		company := pkg["company"].(string)
		tn := pkg["tn"].(string)
		alias := pkg["alias"].(string)
		timelineRaw := pkg["timeline"].([]interface{})

		var timeline []Stat
		for _, t := range timelineRaw {
			timelineItem := t.(map[interface{}]interface{})
			date := timelineItem["date"].(string)
			status := timelineItem["status"].(string)
			office := timelineItem["office"].(string)
			timeline = append(timeline, Stat{Date: date, Status: status, Office: office})
		}

		p := &Pkg{
			Company:  company,
			TN:       tn,
			Alias:    alias,
			Timeline: timeline,
		}

		if err := l.Addtoli(p); err != nil {
			return err
		}
	}
	return nil
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func CreateFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return f.Close()
}

func (l *List) Gettn(alias string) string {
	for _, p := range l.Pkgs {
		if p.Alias == alias {
			return p.TN
		}
	}
	return ""
}

func (p *Pkg) Tracknimotu() error {
	c := companies[p.Company]

	stat, err := c.Trackfunc(p.TN)

	if err != nil {
		return err
	}

	p.Timeline = stat
	return nil
}
