// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"
	"net/url"

	"github.com/CZL-AI/czlai-go/internal/apiquery"
	"github.com/CZL-AI/czlai-go/internal/param"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// UserUserInfoService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserUserInfoService] method instead.
type UserUserInfoService struct {
	Options []option.RequestOption
}

// NewUserUserInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserUserInfoService(opts ...option.RequestOption) (r *UserUserInfoService) {
	r = &UserUserInfoService{}
	r.Options = opts
	return
}

// 获取用户信息
func (r *UserUserInfoService) Get(ctx context.Context, query UserUserInfoGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "user-info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type UserUserInfoGetParams struct {
	// 用户 UUID
	Uuid param.Field[string] `query:"uuid,required"`
}

// URLQuery serializes [UserUserInfoGetParams]'s query parameters as `url.Values`.
func (r UserUserInfoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
