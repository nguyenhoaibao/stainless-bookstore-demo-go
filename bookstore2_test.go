// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bookstore2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nguyenhoaibao/stainless-bookstore-demo-go"
	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/internal/testutil"
	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/option"
)

func TestBookstore2Echo(t *testing.T) {
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
		var apierr *bookstore2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
