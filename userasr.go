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

// UserAsrService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserAsrService] method instead.
type UserAsrService struct {
	Options []option.RequestOption
}

// NewUserAsrService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserAsrService(opts ...option.RequestOption) (r *UserAsrService) {
	r = &UserAsrService{}
	r.Options = opts
	return
}

// 语音识别
func (r *UserAsrService) New(ctx context.Context, body UserAsrNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "asr"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserAsrNewParams struct {
	Length param.Field[int64]  `json:"length"`
	Speech param.Field[string] `json:"speech"`
}

func (r UserAsrNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
