package go_shopify

import "testing"

func TestMetafieldPathPrefix(t *testing.T) {
	cases := []struct {
		resource   string
		resourceID int64
		expected   string
	}{
		{"", 0, "metafields"},
		{"products", 123, "products/123/metafields"},
	}

	for _, c := range cases {
		actual := MetafieldPathPrefix(c.resource, c.resourceID)
		if actual != c.expected {
			t.Errorf("MetafieldPathPrefix(%s, %d): expected %s, actual %s", c.resource, c.resourceID, c.expected, actual)
		}
	}
}

func TestFulfillmentPathPrefix(t *testing.T) {
	cases := []struct {
		resource   string
		resourceID int64
		expected   string
	}{
		{"", 0, "fulfillments"},
		{"orders", 123, "orders/123/fulfillments"},
	}

	for _, c := range cases {
		actual := FulfillmentPathPrefix(c.resource, c.resourceID)
		if actual != c.expected {
			t.Errorf("FulfillmentPathPrefix(%s, %d): expected %s, actual %s", c.resource, c.resourceID, c.expected, actual)
		}
	}
}
