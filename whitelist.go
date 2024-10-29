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

// WhitelistService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhitelistService] method instead.
type WhitelistService struct {
	Options []option.RequestOption
}

// NewWhitelistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhitelistService(opts ...option.RequestOption) (r *WhitelistService) {
	r = &WhitelistService{}
	r.Options = opts
	return
}

// 白名单数据过滤
func (r *WhitelistService) FilteringData(ctx context.Context, body WhitelistFilteringDataParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "whitelist/filtering_data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 白名单数据保存
func (r *WhitelistService) SaveData(ctx context.Context, body WhitelistSaveDataParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "whitelist/save_data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type WhitelistFilteringDataParams struct {
	// 过滤数据
	FilteringData param.Field[string] `json:"filtering_data"`
}

func (r WhitelistFilteringDataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhitelistSaveDataParams struct {
	// 保存数据
	SaveData param.Field[[]string] `json:"save_data"`
}

func (r WhitelistSaveDataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
