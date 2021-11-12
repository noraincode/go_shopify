package go_shopify

import "fmt"

// MetafieldPathPrefix Return the prefix for a metafield path
func MetafieldPathPrefix(resource string, resourceID int64) string {
	prefix := "metafields"
	if resource != "" {
		prefix = fmt.Sprintf("%s/%d/metafields", resource, resourceID)
	}
	return prefix
}

// FulfillmentPathPrefix Return the prefix for a fulfillment path
func FulfillmentPathPrefix(resource string, resourceID int64) string {
	prefix := "fulfillments"
	if resource != "" {
		prefix = fmt.Sprintf("%s/%d/fulfillments", resource, resourceID)
	}
	return prefix
}
