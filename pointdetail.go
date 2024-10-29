// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/czlai-go/internal/apijson"
	"github.com/stainless-sdks/czlai-go/internal/apiquery"
	"github.com/stainless-sdks/czlai-go/internal/param"
	"github.com/stainless-sdks/czlai-go/internal/requestconfig"
	"github.com/stainless-sdks/czlai-go/option"
)

// PointDetailService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPointDetailService] method instead.
type PointDetailService struct {
	Options []option.RequestOption
}

// NewPointDetailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPointDetailService(opts ...option.RequestOption) (r *PointDetailService) {
	r = &PointDetailService{}
	r.Options = opts
	return
}

// 获取积分明细
func (r *PointDetailService) Get(ctx context.Context, query PointDetailGetParams, opts ...option.RequestOption) (res *PointDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "point-detail"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PointDetailGetResponse struct {
	Data []PointDetailGetResponseData `json:"data"`
	JSON pointDetailGetResponseJSON   `json:"-"`
}

// pointDetailGetResponseJSON contains the JSON metadata for the struct
// [PointDetailGetResponse]
type pointDetailGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PointDetailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pointDetailGetResponseJSON) RawJSON() string {
	return r.raw
}

type PointDetailGetResponseData struct {
	// id
	ID int64 `json:"id"`
	// 明细说明
	Description string `json:"description"`
	// 明细类型 1-购买增加积分 2-活动增加积分 3-模块核销积分
	DetailType int64 `json:"detail_type"`
	// 0-减少 1-增加
	IsAdd PointDetailGetResponseDataIsAdd `json:"is_add"`
	// 0-非购买积分 1-购买积分
	IsPurchasePoint int64 `json:"is_purchase_point"`
	// 积分数量
	PointNum string `json:"point_num"`
	// 会话 id
	SessionID string `json:"session_id"`
	// 用户 uuid
	UserUuid string                         `json:"user_uuid"`
	JSON     pointDetailGetResponseDataJSON `json:"-"`
}

// pointDetailGetResponseDataJSON contains the JSON metadata for the struct
// [PointDetailGetResponseData]
type pointDetailGetResponseDataJSON struct {
	ID              apijson.Field
	Description     apijson.Field
	DetailType      apijson.Field
	IsAdd           apijson.Field
	IsPurchasePoint apijson.Field
	PointNum        apijson.Field
	SessionID       apijson.Field
	UserUuid        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PointDetailGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pointDetailGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// 0-减少 1-增加
type PointDetailGetResponseDataIsAdd int64

const (
	PointDetailGetResponseDataIsAdd0 PointDetailGetResponseDataIsAdd = 0
	PointDetailGetResponseDataIsAdd1 PointDetailGetResponseDataIsAdd = 1
)

func (r PointDetailGetResponseDataIsAdd) IsKnown() bool {
	switch r {
	case PointDetailGetResponseDataIsAdd0, PointDetailGetResponseDataIsAdd1:
		return true
	}
	return false
}

type PointDetailGetParams struct {
	// 0-支出 1-收入 2-全部
	IsAdd param.Field[PointDetailGetParamsIsAdd] `query:"is_add"`
	// 每页数量
	Limit param.Field[int64] `query:"limit"`
	// 页数
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [PointDetailGetParams]'s query parameters as `url.Values`.
func (r PointDetailGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// 0-支出 1-收入 2-全部
type PointDetailGetParamsIsAdd int64

const (
	PointDetailGetParamsIsAdd0 PointDetailGetParamsIsAdd = 0
	PointDetailGetParamsIsAdd1 PointDetailGetParamsIsAdd = 1
	PointDetailGetParamsIsAdd2 PointDetailGetParamsIsAdd = 2
)

func (r PointDetailGetParamsIsAdd) IsKnown() bool {
	switch r {
	case PointDetailGetParamsIsAdd0, PointDetailGetParamsIsAdd1, PointDetailGetParamsIsAdd2:
		return true
	}
	return false
}
