package go_shopify

import (
	"fmt"
	"io/ioutil"
)

func loadFixture(filename string) []byte {
	f, err := ioutil.ReadFile("__fixtures__/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	return f
}
