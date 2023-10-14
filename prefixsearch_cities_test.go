package prefixsearch_test

import (
	"testing"

	gocitiesjson "github.com/ringsaturn/go-cities.json"
	"github.com/ringsaturn/prefixsearch"
)

type City struct {
	Name    string
	Country string
	Lng     float64
	Lon     float64
}

func (c *City) LessThan(j City) bool { return c.Name < j.Name }
func (c *City) EqualTo(j City) bool  { return c.Name == j.Name }

func BenchmarkSearchTree_Cities(b *testing.B) {
	st := prefixsearch.New[City]()
	for _, c := range gocitiesjson.Cities {
		city := City{c.Name, c.Country, c.Lng, c.Lat}
		st.Add(city.Name, city)
	}
	words := []string{
		"New",
		"New York",
		"London",
		"Beijing",
		"Tokyo",
		"Shanghai",
		"Paris",
		"Berlin",
		"Madrid",
		"Rome",
	}
	for _, word := range words {
		b.ResetTimer()
		b.Run(word, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				st.Search(word)
			}
		})
	}
}
