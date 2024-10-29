// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/czlai-go/internal/apijson"
	"github.com/stainless-sdks/czlai-go/internal/param"
	"github.com/stainless-sdks/czlai-go/internal/requestconfig"
	"github.com/stainless-sdks/czlai-go/option"
)

// AITrialService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITrialService] method instead.
type AITrialService struct {
	Options []option.RequestOption
}

// NewAITrialService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAITrialService(opts ...option.RequestOption) (r *AITrialService) {
	r = &AITrialService{}
	r.Options = opts
	return
}

// 获取问题选项
func (r *AITrialService) Options(ctx context.Context, body AITrialOptionsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-trial/options"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取问题(流式)
func (r *AITrialService) Question(ctx context.Context, body AITrialQuestionParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-trial/question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AITrialOptionsParams struct {
	// 问题
	Question param.Field[string] `json:"question"`
	// 1-猫狗 2-异宠
	ServiceType param.Field[int64] `json:"service_type"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AITrialOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AITrialQuestionParams struct {
	// 1-猫狗 2-异宠
	ServiceType param.Field[int64] `json:"service_type"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AITrialQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
