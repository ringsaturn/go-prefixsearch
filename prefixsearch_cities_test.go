package prefixsearch_test

import (
	"testing"

	gocitiesjson "github.com/ringsaturn/go-cities.json"
	"github.com/ringsaturn/prefixsearch"
)

type city struct {
	Name    string
	Country string
	Lng     float64
	Lat     float64
}

func (c city) LessThan(j city) bool { return c.Name < j.Name }
func (c city) EqualTo(j city) bool  { return c.Name == j.Name }

var commonCities = []string{
	"New York",
	"London",
	"Beijing",
	"Haidian",
	"Tokyo",
	"Shanghai",
	"Paris",
	"Berlin",
	"Madrid",
	"Rome",
}

var commonInputs = map[string][]string{
	"New York": {
		"N",
		"Ne",
		"New",
		"New ",
		"New Y",
		"New Yo",
		"New Yor",
		"New York",
	},
	"London": {
		"L",
		"Lo",
		"Lon",
		"Lond",
		"Londo",
		"London",
	},
	"Beijing": {
		"B",
		"Be",
		"Bei",
		"Beij",
		"Beiji",
		"Beijin",
		"Beijing",
	},
	"Haidian": {
		"H",
		"Ha",
		"Hai",
		"Haid",
		"Haidi",
		"Haidia",
		"Haidian",
	},
	"Tokyo": {
		"T",
		"To",
		"Tok",
		"Toky",
		"Tokyo",
	},
	"Shanghai": {
		"S",
		"Sh",
		"Sha",
		"Shan",
		"Shang",
		"Shangh",
		"Shangha",
		"Shanghai",
	},
	"Paris": {
		"P",
		"Pa",
		"Par",
		"Paris",
	},
	"Berlin": {
		"B",
		"Be",
		"Ber",
		"Berl",
		"Berli",
		"Berlin",
	},
	"Madrid": {
		"M",
		"Ma",
		"Mad",
		"Madri",
		"Madrid",
	},
}

func newSearchTree_Cities() *prefixsearch.SearchTree[city] {
	st := prefixsearch.New[city]()
	for _, c := range gocitiesjson.Cities {
		city := city{
			Name:    c.Name,
			Country: c.Country,
			Lng:     c.Lng,
			Lat:     c.Lat,
		}
		st.Add(city.Name, city)
	}

	return st
}

func BenchmarkSearchTree_Search_Cities(b *testing.B) {
	st := newSearchTree_Cities()
	b.ResetTimer()

	for _, city := range commonCities {
		b.Run(city, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = st.Search(city)
			}
		})
	}
}

func BenchmarkSearchTree_AutoComplete_Cities(b *testing.B) {
	st := newSearchTree_Cities()
	b.ResetTimer()

	for cityName, inputs := range commonInputs {
		for _, input := range inputs {
			b.Run(cityName+"_"+input, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if len(st.AutoComplete(input)) == 0 {
						b.Fatal("no result")
					}
				}
			})
		}
	}
}
