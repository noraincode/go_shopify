package go_shopify

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

// errReader can be used to simulate a failed call to response.Body.Read
type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("test-error")
}

func (errReader) Close() error {
	return nil
}

func TestResponseErrorStructError(t *testing.T) {
	res := ResponseError{
		Status:  400,
		Message: "invalid email",
		Errors:  []string{"invalid email"},
	}

	expected := ResponseError{
		Status:  res.GetStatus(),
		Message: res.GetMessage(),
		Errors:  res.GetErrors(),
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("ResponseError returned  %+v, expected %+v", res, expected)
	}
}

func TestResponseErrorError(t *testing.T) {
	cases := []struct {
		err      ResponseError
		expected string
	}{
		{
			ResponseError{Message: "oh no"},
			"oh no",
		},
		{
			ResponseError{},
			"Unknown Error",
		},
		{
			ResponseError{Errors: []string{"title: not a valid title"}},
			"title: not a valid title",
		},
		{
			ResponseError{Errors: []string{
				"not a valid title",
				"not a valid description",
			}},
			// The strings are sorted description comes first
			"not a valid description, not a valid title",
		},
	}

	for _, c := range cases {
		actual := fmt.Sprint(c.err)
		if actual != c.expected {
			t.Errorf("ResponseError.Error(): expected %s, actual %s", c.expected, actual)
		}
	}
}

func TestCheckResponseError(t *testing.T) {
	cases := []struct {
		resp     *http.Response
		expected error
	}{
		{
			httpmock.NewStringResponse(200, `{"foo": "bar"}`),
			nil,
		},
		{
			httpmock.NewStringResponse(299, `{"foo": "bar"}`),
			nil,
		},
		{
			httpmock.NewStringResponse(400, `{"error": "bad request"}`),
			ResponseError{Status: 400, Message: "bad request"},
		},
		{
			httpmock.NewStringResponse(500, `{"error": "terrible error"}`),
			ResponseError{Status: 500, Message: "terrible error"},
		},
		{
			httpmock.NewStringResponse(500, `{"errors": "This action requires read_customers scope"}`),
			ResponseError{Status: 500, Message: "This action requires read_customers scope"},
		},
		{
			httpmock.NewStringResponse(500, `{"errors": ["not", "very good"]}`),
			ResponseError{Status: 500, Message: "not, very good", Errors: []string{"not", "very good"}},
		},
		{
			httpmock.NewStringResponse(400, `{"errors": { "order": ["order is wrong"] }}`),
			ResponseError{Status: 400, Message: "order: order is wrong", Errors: []string{"order: order is wrong"}},
		},
		{
			httpmock.NewStringResponse(400, `{"errors": { "collection_id": "collection_id is wrong" }}`),
			ResponseError{Status: 400, Message: "collection_id: collection_id is wrong", Errors: []string{"collection_id: collection_id is wrong"}},
		},
		{
			httpmock.NewStringResponse(400, `{error:bad request}`),
			errors.New("invalid character 'e' looking for beginning of object key string"),
		},
		{
			&http.Response{StatusCode: 400, Body: errReader{}},
			errors.New("test-error"),
		},
	}

	for _, c := range cases {
		actual := CheckResponseError(c.resp)
		if fmt.Sprint(actual) != fmt.Sprint(c.expected) {
			t.Errorf("CheckResponseError(): expected %v, actual %v", c.expected, actual)
		}
	}
}
