package go_shopify

import "fmt"

// ShopBaseUrl Return the Shop's base url.
func ShopBaseUrl(name string) string {
	name = ShopFullName(name)
	return fmt.Sprintf("https://%s", name)
}
