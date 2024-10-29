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

// AIConvService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConvService] method instead.
type AIConvService struct {
	Options []option.RequestOption
}

// NewAIConvService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIConvService(opts ...option.RequestOption) (r *AIConvService) {
	r = &AIConvService{}
	r.Options = opts
	return
}

// AI 聊天
func (r *AIConvService) Answer(ctx context.Context, body AIConvAnswerParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "ai-conv/answer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取联想
func (r *AIConvService) Relation(ctx context.Context, body AIConvRelationParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-conv/relation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// AI 聊天校验
func (r *AIConvService) Validate(ctx context.Context, body AIConvValidateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-conv/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AIConvAnswerParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
	// 用户输入
	UserInput param.Field[string] `json:"user_input"`
}

func (r AIConvAnswerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIConvRelationParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
	// 用户输入
	UserInput param.Field[string] `json:"user_input"`
}

func (r AIConvRelationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIConvValidateParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
	// 用户输入
	UserInput param.Field[string] `json:"user_input"`
}

func (r AIConvValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
