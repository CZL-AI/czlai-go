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

// CheckPointService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckPointService] method instead.
type CheckPointService struct {
	Options []option.RequestOption
}

// NewCheckPointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckPointService(opts ...option.RequestOption) (r *CheckPointService) {
	r = &CheckPointService{}
	r.Options = opts
	return
}

// 埋点
func (r *CheckPointService) New(ctx context.Context, body CheckPointNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "check-point"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CheckPointNewParams struct {
	// 埋点动作
	Action param.Field[string] `json:"action"`
	// 微信 code
	Code param.Field[string] `json:"code"`
	// 页面路径
	PagePath param.Field[string] `json:"page_path"`
	// 关联 id
	RelatedID param.Field[string] `json:"related_id"`
}

func (r CheckPointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
