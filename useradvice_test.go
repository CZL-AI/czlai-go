// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/czlai-go"
	"github.com/stainless-sdks/czlai-go/internal/testutil"
	"github.com/stainless-sdks/czlai-go/option"
)

func TestUserAdviceNew(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := czlai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	err := client.UserAdvices.New(context.TODO(), czlai.UserAdviceNewParams{
		AdviceType:  czlai.F("advice_type"),
		Description: czlai.F("description"),
		ImageList:   czlai.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
