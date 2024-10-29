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

// PolicyService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyService] method instead.
type PolicyService struct {
	Options []option.RequestOption
}

// NewPolicyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPolicyService(opts ...option.RequestOption) (r *PolicyService) {
	r = &PolicyService{}
	r.Options = opts
	return
}

// 政策类型
func (r *PolicyService) PolicyInfo(ctx context.Context, body PolicyPolicyInfoParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "policy/policy_info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type PolicyPolicyInfoParams struct {
	// 密钥
	Keys param.Field[string] `json:"keys"`
	// 1-用户协议 2-免责条款 3-隐私政策
	PolicyType param.Field[int64] `json:"policy_type"`
}

func (r PolicyPolicyInfoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
