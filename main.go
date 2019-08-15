package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/francoispqt/gojay"
)

// ===== [ MAIN ] ===============

func main() {
	getLatest().dump()
}

// ===== [ XKCD ] ===============

// Comic is a struct that stores all information relevant to a comic
type Comic struct {
	Day   string
	Month string
	Year  string
	Num   int
	Title string
	Alt   string
	Img   string
}

func (c *Comic) NKeys() int {
	return 7
}

func (c *Comic) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "day":
		return dec.String(&c.Day)
	case "month":
		return dec.String(&c.Month)
	case "year":
		return dec.String(&c.Year)
	case "num":
		return dec.Int(&c.Num)
	case "title":
		return dec.String(&c.Title)
	case "alt":
		return dec.String(&c.Alt)
	case "img":
		return dec.String(&c.Img)
	}
	return nil
}

// ===== [ GET COMIC ] ==========

func get(apiURL string) (c Comic) {
	// Get data from XKCD api endpoint
	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Read response body into a []byte
	body, err := ioutil.ReadAll(resp.Body)

	// Create a comic struct and populate it with xkcd data
	c = Comic{}
	err = gojay.UnmarshalJSONObject(body, &c)
	if err != nil {
		panic(err)
	}

	return c
}

func getLatest() (c Comic) {
	return get("https://xkcd.com/info.0.json")
}

// ===== [ DUMP COMIC ] ==========

func (c Comic) dump() {
	fmt.Printf("XKCD #%d published on %s/%s/%s\n", c.Num, c.Month, c.Day, c.Year)
	fmt.Println(c.Title)
	fmt.Println(c.Alt)
	fmt.Println(c.Img)
}

// func (c Comic) dump() {
// 	fmt.Printf("XKCD #%d published on %s/%s/%s\n", c.Num, c.Month, c.Day, c.Year)
// 	fmt.Println(c.Title)
// 	fmt.Println(c.Alt)
// 	fmt.Println(c.Img)
// }
