package go_shopify

import "testing"

func TestShopBaseUrl(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "https://myshop.myshopify.com"},
		{"myshop.", "https://myshop.myshopify.com"},
		{" myshop", "https://myshop.myshopify.com"},
		{"myshop ", "https://myshop.myshopify.com"},
		{"myshop \n", "https://myshop.myshopify.com"},
		{"myshop.myshopify.com", "https://myshop.myshopify.com"},
	}

	for _, c := range cases {
		actual := ShopBaseUrl(c.in)
		if actual != c.expected {
			t.Errorf("ShopBaseUrl(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}
