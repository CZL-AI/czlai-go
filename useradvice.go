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

// UserAdviceService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserAdviceService] method instead.
type UserAdviceService struct {
	Options []option.RequestOption
}

// NewUserAdviceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAdviceService(opts ...option.RequestOption) (r *UserAdviceService) {
	r = &UserAdviceService{}
	r.Options = opts
	return
}

// 用户反馈
func (r *UserAdviceService) New(ctx context.Context, body UserAdviceNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "user-advice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserAdviceNewParams struct {
	// 反馈类型
	AdviceType param.Field[string] `json:"advice_type,required"`
	// 反馈内容
	Description param.Field[string]   `json:"description,required"`
	ImageList   param.Field[[]string] `json:"image_list,required"`
}

func (r UserAdviceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
