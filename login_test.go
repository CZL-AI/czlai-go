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

func TestLoginSMSWithOptionalParams(t *testing.T) {
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
	_, err := client.Logins.SMS(context.TODO(), czlai.LoginSMSParams{
		Code:      czlai.F("123456"),
		Phone:     czlai.F("phone"),
		LoginFrom: czlai.F(int64(0)),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoginWechatWithOptionalParams(t *testing.T) {
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
	err := client.Logins.Wechat(context.TODO(), czlai.LoginWechatParams{
		WechatCode:    czlai.F("wechat_code"),
		EncryptedData: czlai.F("encryptedData"),
		Iv:            czlai.F("iv"),
		ModuleType:    czlai.F(int64(0)),
		PhoneNumber:   czlai.F("phone_number"),
		SpreadID:      czlai.F(int64(0)),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
