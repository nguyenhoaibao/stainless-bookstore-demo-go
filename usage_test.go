// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bookstore2_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/bookstore-2-go"
	"github.com/stainless-sdks/bookstore-2-go/internal/testutil"
	"github.com/stainless-sdks/bookstore-2-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bookstore2.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Echo(context.TODO(), bookstore2.EchoParams{
		Body: "Hello world!",
	})
	if err != nil {
		t.Error(err)
	}
}
