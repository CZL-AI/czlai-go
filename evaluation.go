// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"

	"github.com/CZL-AI/czlai-go/internal/apijson"
	"github.com/CZL-AI/czlai-go/internal/param"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// EvaluationService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvaluationService] method instead.
type EvaluationService struct {
	Options []option.RequestOption
}

// NewEvaluationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvaluationService(opts ...option.RequestOption) (r *EvaluationService) {
	r = &EvaluationService{}
	r.Options = opts
	return
}

// 修改评价
func (r *EvaluationService) PutEvaluation(ctx context.Context, body EvaluationPutEvaluationParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "evaluation/put_evaluation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type EvaluationPutEvaluationParams struct {
	// 文本内容
	Content param.Field[string] `json:"content,required"`
	// 评价 id
	Evaluation param.Field[int64] `json:"evaluation,required"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
}

func (r EvaluationPutEvaluationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
