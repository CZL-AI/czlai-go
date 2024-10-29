// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"

	"github.com/CZL-AI/czlai-go/internal/apijson"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// UserModuleUsageIsAddWecomeService contains methods and other services that help
// with interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserModuleUsageIsAddWecomeService] method instead.
type UserModuleUsageIsAddWecomeService struct {
	Options []option.RequestOption
}

// NewUserModuleUsageIsAddWecomeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserModuleUsageIsAddWecomeService(opts ...option.RequestOption) (r *UserModuleUsageIsAddWecomeService) {
	r = &UserModuleUsageIsAddWecomeService{}
	r.Options = opts
	return
}

// 是否领取过添加企微的奖励
func (r *UserModuleUsageIsAddWecomeService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserModuleUsageIsAddWecomeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user-module-usage/is-add-wecome"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserModuleUsageIsAddWecomeGetResponse struct {
	Data    UserModuleUsageIsAddWecomeGetResponseData `json:"data"`
	Message string                                    `json:"message"`
	Success bool                                      `json:"success"`
	JSON    userModuleUsageIsAddWecomeGetResponseJSON `json:"-"`
}

// userModuleUsageIsAddWecomeGetResponseJSON contains the JSON metadata for the
// struct [UserModuleUsageIsAddWecomeGetResponse]
type userModuleUsageIsAddWecomeGetResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserModuleUsageIsAddWecomeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userModuleUsageIsAddWecomeGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserModuleUsageIsAddWecomeGetResponseData struct {
	IsAddWecome int64                                         `json:"is_add_wecome"`
	JSON        userModuleUsageIsAddWecomeGetResponseDataJSON `json:"-"`
}

// userModuleUsageIsAddWecomeGetResponseDataJSON contains the JSON metadata for the
// struct [UserModuleUsageIsAddWecomeGetResponseData]
type userModuleUsageIsAddWecomeGetResponseDataJSON struct {
	IsAddWecome apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserModuleUsageIsAddWecomeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userModuleUsageIsAddWecomeGetResponseDataJSON) RawJSON() string {
	return r.raw
}
