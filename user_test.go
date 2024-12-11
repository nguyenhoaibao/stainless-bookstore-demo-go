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

func TestUserGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Get(
		context.TODO(),
		"username",
		bookstore2.UserGetParams{
			PrettyPrint: bookstore2.F(true),
			WithEmail:   bookstore2.F(true),
		},
	)
	if err != nil {
		var apierr *bookstore2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Users.Update(
		context.TODO(),
		"username",
		bookstore2.UserUpdateParams{
			User: bookstore2.UserParam{
				Email:     bookstore2.F("john.smith@example.com"),
				FirstName: bookstore2.F("John"),
				LastName:  bookstore2.F("Smith"),
				Username:  bookstore2.F("John78"),
			},
			PrettyPrint: bookstore2.F(true),
		},
	)
	if err != nil {
		var apierr *bookstore2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
