// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"

	"github.com/CZL-AI/czlai-go/internal/apijson"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// PointPriceService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPointPriceService] method instead.
type PointPriceService struct {
	Options []option.RequestOption
}

// NewPointPriceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPointPriceService(opts ...option.RequestOption) (r *PointPriceService) {
	r = &PointPriceService{}
	r.Options = opts
	return
}

// 获取积分价格
func (r *PointPriceService) Get(ctx context.Context, opts ...option.RequestOption) (res *PointPriceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "point-price"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PointPriceGetResponse struct {
	Data PointPriceGetResponseData `json:"data"`
	JSON pointPriceGetResponseJSON `json:"-"`
}

// pointPriceGetResponseJSON contains the JSON metadata for the struct
// [PointPriceGetResponse]
type pointPriceGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PointPriceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pointPriceGetResponseJSON) RawJSON() string {
	return r.raw
}

type PointPriceGetResponseData struct {
	// 是否为模块项目
	IsModuleItem int64 `json:"is_module_item"`
	// 积分类型
	ItemKey string `json:"item_key"`
	// 项目名
	ItemName string `json:"item_name"`
	// 项目 key 名
	Price string `json:"price"`
	// 关联模块
	RelatedModule string                        `json:"related_module"`
	JSON          pointPriceGetResponseDataJSON `json:"-"`
}

// pointPriceGetResponseDataJSON contains the JSON metadata for the struct
// [PointPriceGetResponseData]
type pointPriceGetResponseDataJSON struct {
	IsModuleItem  apijson.Field
	ItemKey       apijson.Field
	ItemName      apijson.Field
	Price         apijson.Field
	RelatedModule apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PointPriceGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pointPriceGetResponseDataJSON) RawJSON() string {
	return r.raw
}
