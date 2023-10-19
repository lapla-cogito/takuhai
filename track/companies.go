package track

type Companyname string

type companystat struct {
	Key     string
	Shorten string
}

type company struct {
	*companystat
	Trackfunc func(string) ([]Stat, error)
}

var companies = map[string]*company{}
