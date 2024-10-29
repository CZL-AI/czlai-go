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

// UserSummaryService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSummaryService] method instead.
type UserSummaryService struct {
	Options []option.RequestOption
}

// NewUserSummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserSummaryService(opts ...option.RequestOption) (r *UserSummaryService) {
	r = &UserSummaryService{}
	r.Options = opts
	return
}

// 合作交流信息保存
func (r *UserSummaryService) New(ctx context.Context, body UserSummaryNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserSummaryNewParams struct {
	SessionID param.Field[string] `json:"session_id"`
}

func (r UserSummaryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
