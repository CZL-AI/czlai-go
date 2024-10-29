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

// UserPointService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPointService] method instead.
type UserPointService struct {
	Options []option.RequestOption
}

// NewUserPointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserPointService(opts ...option.RequestOption) (r *UserPointService) {
	r = &UserPointService{}
	r.Options = opts
	return
}

// 获取用户积分
func (r *UserPointService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserPointGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user-point"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// 积分消耗报表
func (r *UserPointService) CostReport(ctx context.Context, body UserPointCostReportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "user-point/cost-report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserPointGetResponse struct {
	Data UserPointGetResponseData `json:"data"`
	JSON userPointGetResponseJSON `json:"-"`
}

// userPointGetResponseJSON contains the JSON metadata for the struct
// [UserPointGetResponse]
type userPointGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPointGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserPointGetResponseData struct {
	// 赠送的积分余额
	BonusPoint string `json:"bonus_point"`
	// 购买的积分余额
	PurchasePoint string `json:"purchase_point"`
	// 总积分余额
	TotalPoint string                       `json:"total_point"`
	JSON       userPointGetResponseDataJSON `json:"-"`
}

// userPointGetResponseDataJSON contains the JSON metadata for the struct
// [UserPointGetResponseData]
type userPointGetResponseDataJSON struct {
	BonusPoint    apijson.Field
	PurchasePoint apijson.Field
	TotalPoint    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserPointGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPointGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type UserPointCostReportParams struct {
	ItemKey param.Field[string] `json:"item_key"`
	// 病历 id
	MedicalRecordID param.Field[int64] `json:"medical_record_id"`
}

func (r UserPointCostReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
