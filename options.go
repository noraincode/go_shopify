package go_shopify

import (
	"fmt"
	"net/http"
	"time"
)

// Option is used to configure client with options
type Option func(c *Client)

// ListOptions General list options that can be used for most collections of entities.
type ListOptions struct {

	// PageInfo is used with new pagination search.
	PageInfo string `url:"page_info,omitempty"`

	// Page is used to specify a specific page to load.
	// It is the deprecated way to do pagination.
	Page         int       `url:"page,omitempty"`
	Limit        int       `url:"limit,omitempty"`
	SinceID      int64     `url:"since_id,omitempty"`
	CreatedAtMin time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax time.Time `url:"updated_at_max,omitempty"`
	Order        string    `url:"order,omitempty"`
	Fields       string    `url:"fields,omitempty"`
	Vendor       string    `url:"vendor,omitempty"`
	IDs          []int64   `url:"ids,omitempty,comma"`
}

// CountOptions General count options that can be used for most collection counts.
type CountOptions struct {
	CreatedAtMin time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax time.Time `url:"updated_at_max,omitempty"`
}

// WithVersion optionally sets the api-version if the passed string is valid
func WithVersion(apiVersion string) Option {
	return func(c *Client) {
		pathPrefix := defaultApiPathPrefix
		if len(apiVersion) > 0 && (apiVersionRegex.MatchString(apiVersion) || apiVersion == UnstableApiVersion) {
			pathPrefix = fmt.Sprintf("admin/api/%s", apiVersion)
		}
		c.apiVersion = apiVersion
		c.pathPrefix = pathPrefix
	}
}

func WithRetry(retries int) Option {
	return func(c *Client) {
		c.retries = retries
	}
}

func WithLogger(logger LeveledLoggerInterface) Option {
	return func(c *Client) {
		c.log = logger
	}
}

// WithHTTPClient is used to set a custom http client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.Client = client
	}
}
