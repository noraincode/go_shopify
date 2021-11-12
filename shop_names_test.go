package go_shopify

import "testing"

func TestShopFullName(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "myshop.myshopify.com"},
		{"myshop.", "myshop.myshopify.com"},
		{" myshop", "myshop.myshopify.com"},
		{"myshop ", "myshop.myshopify.com"},
		{"myshop \n", "myshop.myshopify.com"},
		{"myshop.myshopify.com", "myshop.myshopify.com"},
	}

	for _, c := range cases {
		actual := ShopFullName(c.in)
		if actual != c.expected {
			t.Errorf("ShopFullName(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}

func TestShopShortName(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "myshop"},
		{"myshop.", "myshop"},
		{" myshop", "myshop"},
		{"myshop ", "myshop"},
		{"myshop \n", "myshop"},
		{"myshop.myshopify.com", "myshop"},
		{".myshop.myshopify.com.", "myshop"},
	}

	for _, c := range cases {
		actual := ShopShortName(c.in)
		if actual != c.expected {
			t.Errorf("ShopShortName(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}
