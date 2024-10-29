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

// BuriedpointService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuriedpointService] method instead.
type BuriedpointService struct {
	Options []option.RequestOption
}

// NewBuriedpointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBuriedpointService(opts ...option.RequestOption) (r *BuriedpointService) {
	r = &BuriedpointService{}
	r.Options = opts
	return
}

// 保存页埋点
func (r *BuriedpointService) New(ctx context.Context, body BuriedpointNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "page-buriedpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type BuriedpointNewParams struct {
	// 参数
	Point param.Field[string] `json:"point,required"`
	// 微信 code
	Code param.Field[string] `json:"code"`
}

func (r BuriedpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
