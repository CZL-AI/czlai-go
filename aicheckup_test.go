// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/CZL-AI/czlai-go"
	"github.com/CZL-AI/czlai-go/internal/testutil"
	"github.com/CZL-AI/czlai-go/option"
)

func TestAICheckupIsFirst(t *testing.T) {
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
	_, err := client.AICheckup.IsFirst(context.TODO(), czlai.AICheckupIsFirstParams{
		PetProfileID: czlai.F(int64(0)),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupPicResult(t *testing.T) {
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
	err := client.AICheckup.PicResult(context.TODO(), czlai.AICheckupPicResultParams{
		ImgURL:       czlai.F("img_url"),
		PetProfileID: czlai.F(int64(0)),
		SessionID:    czlai.F("session_id"),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupQuestion(t *testing.T) {
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
	err := client.AICheckup.Question(context.TODO(), czlai.AICheckupQuestionParams{
		PetProfileID: czlai.F(int64(0)),
		SessionID:    czlai.F("session_id"),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupQuestionResult(t *testing.T) {
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
	err := client.AICheckup.QuestionResult(context.TODO(), czlai.AICheckupQuestionResultParams{
		Index:        czlai.F(int64(0)),
		PetProfileID: czlai.F(int64(0)),
		QuestionID:   czlai.F("question_id"),
		Round:        czlai.F("round"),
		SessionID:    czlai.F("session_id"),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupReport(t *testing.T) {
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
	err := client.AICheckup.Report(context.TODO(), czlai.AICheckupReportParams{
		PetProfileID: czlai.F(int64(0)),
		SessionID:    czlai.F("session_id"),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupReportTaskWithOptionalParams(t *testing.T) {
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
	err := client.AICheckup.ReportTask(context.TODO(), czlai.AICheckupReportTaskParams{
		SessionID:  czlai.F("session_id"),
		ImgURL:     czlai.F("img_url"),
		ReportType: czlai.F(int64(0)),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupSessionStart(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := czlai.NewClient(
		option.WithBaseURL("https://plaform-ms-ai.chongzhiling.com/api/v1.0/ai-b/"),
		option.WithBearerToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjEtMTExLTVkM2MyMzA2MzU5NzJjOGQyZThkZTQ5Yjc0Y2NkMTA0IiwiYXBpcyI6WyIvYWktY2hlY2t1cC9zZXNzaW9uLXN0YXJ0IiwiL3BldC1wcm9maWxlIiwiL3VwbG9hZC1pbWFnZS1vc3MiLCIvcGV0LXByb2ZpbGUvdmFyaWV0eSIsIi9wZXQtcHJvZmlsZXMiLCIvcGV0cy9wZXQtaW5mbyIsIi9wZXQtcHJvZmlsZS9leC1pbmZvIiwiL3Nlc3Npb24tcmVjb3JkL3Nlc3Npb24tc3RhcnQiLCIvc2Vzc2lvbi1yZWNvcmQiLCIvYWlkb2Mvb3B0aW9ucyIsIi9haWRvYy9pZi1jb250aW51ZS1hc2siLCIvYWlkb2Mvc3VtbWFyeSIsIi9haWRvYy9yZXBvcnQiLCIvYWlkb2MvcmVwb3J0LXRhc2siLCIvbWVkaWNhbC1yZWNvcmQiXSwidGltZXMiOjEwMCwiZXhwaXJlc19pbiI6ODY0MDAuMH0sImV4cCI6MTczMDQ0ODAyOCwianRpIjoiNTE0ODZhYjktYTY0MC00NDYwLThlNTMtMzE1YzRlMzNkZGQzIn0.G66BvcESqJRFcwFEFLoFs4eQNSvGOSTTKCRhU1ba4F0"),
	)
	r, err := client.AICheckup.SessionStart(context.TODO())
	fmt.Println(r.Data.SessionID)
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupSummaryWithOptionalParams(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints with content type text/event-stream, Prism mock server will fail")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := czlai.NewClient(
		option.WithBaseURL("https://plaform-ms-ai.chongzhiling.com/api/v1.0/ai-b"),
		option.WithBearerToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOnsidXNlcl9pZCI6IjEtMTExLTVkM2MyMzA2MzU5NzJjOGQyZThkZTQ5Yjc0Y2NkMTA0IiwiYXBpcyI6WyIvYWktY2hlY2t1cC9zZXNzaW9uLXN0YXJ0IiwiL2FpLWNoZWNrdXAvc2Vzc2lvbi1zdGFydCIsIi9wZXQtcHJvZmlsZSIsIi91cGxvYWQtaW1hZ2Utb3NzIiwiL3BldC1wcm9maWxlL3ZhcmlldHkiLCIvcGV0LXByb2ZpbGVzIiwiL3BldHMvcGV0LWluZm8iLCIvcGV0LXByb2ZpbGUvZXgtaW5mbyIsIi9zZXNzaW9uLXJlY29yZC9zZXNzaW9uLXN0YXJ0IiwiL3Nlc3Npb24tcmVjb3JkIiwiL2FpZG9jL29wdGlvbnMiLCIvYWlkb2MvaWYtY29udGludWUtYXNrIiwiL2FpZG9jL3N1bW1hcnkiLCIvYWlkb2MvcmVwb3J0IiwiL2FpZG9jL3JlcG9ydC10YXNrIiwiL21lZGljYWwtcmVjb3JkIl0sInRpbWVzIjoxMDAsImV4cGlyZXNfaW4iOjg2NDAwLjB9LCJleHAiOjE3MzA0NDQ4MDcsImp0aSI6IjRmZmE3MTBhLTg0Y2MtNGNkZS1iN2FhLTA0NzYyYzk2YTU1MSJ9.Bb3duFbM9NEyAZlnE0QcajNPQz-we8uGydpm-MNWlKk"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.AICheckup.Summary(context.TODO(), czlai.AICheckupSummaryParams{
		PetProfileID: czlai.F(int64(0)),
		SessionID:    czlai.F("session_id"),
	})

	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAICheckupUpdateProfileWithOptionalParams(t *testing.T) {
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
	_, err := client.AICheckup.UpdateProfile(context.TODO(), czlai.AICheckupUpdateProfileParams{
		PetProfileID: czlai.F(int64(0)),
		SessionID:    czlai.F("session_id"),
		UpdateType:   czlai.F(int64(0)),
	})
	if err != nil {
		var apierr *czlai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
