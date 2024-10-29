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

// UserModuleUsageService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserModuleUsageService] method instead.
type UserModuleUsageService struct {
	Options     []option.RequestOption
	IsAddWecome *UserModuleUsageIsAddWecomeService
}

// NewUserModuleUsageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserModuleUsageService(opts ...option.RequestOption) (r *UserModuleUsageService) {
	r = &UserModuleUsageService{}
	r.Options = opts
	r.IsAddWecome = NewUserModuleUsageIsAddWecomeService(opts...)
	return
}

// 签到
func (r *UserModuleUsageService) Checkin(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "user-module-usage/checkin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// 领取添加企微的奖励
func (r *UserModuleUsageService) GetAddWecomeBonus(ctx context.Context, body UserModuleUsageGetAddWecomeBonusParams, opts ...option.RequestOption) (res *UserModuleUsageGetAddWecomeBonusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user-module-usage/get-add-wecome-bonus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取微信小程序二维码
func (r *UserModuleUsageService) GetWechatMiniQrcode(ctx context.Context, opts ...option.RequestOption) (res *UserModuleUsageGetWechatMiniQrcodeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user-module-usage/get-wechat-mini-qrcode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type UserModuleUsageGetAddWecomeBonusResponse struct {
	Data    UserModuleUsageGetAddWecomeBonusResponseData `json:"data"`
	Message string                                       `json:"message"`
	Success bool                                         `json:"success"`
	JSON    userModuleUsageGetAddWecomeBonusResponseJSON `json:"-"`
}

// userModuleUsageGetAddWecomeBonusResponseJSON contains the JSON metadata for the
// struct [UserModuleUsageGetAddWecomeBonusResponse]
type userModuleUsageGetAddWecomeBonusResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserModuleUsageGetAddWecomeBonusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userModuleUsageGetAddWecomeBonusResponseJSON) RawJSON() string {
	return r.raw
}

type UserModuleUsageGetAddWecomeBonusResponseData struct {
	IsAddWecome int64                                            `json:"is_add_wecome"`
	JSON        userModuleUsageGetAddWecomeBonusResponseDataJSON `json:"-"`
}

// userModuleUsageGetAddWecomeBonusResponseDataJSON contains the JSON metadata for
// the struct [UserModuleUsageGetAddWecomeBonusResponseData]
type userModuleUsageGetAddWecomeBonusResponseDataJSON struct {
	IsAddWecome apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserModuleUsageGetAddWecomeBonusResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userModuleUsageGetAddWecomeBonusResponseDataJSON) RawJSON() string {
	return r.raw
}

type UserModuleUsageGetWechatMiniQrcodeResponse struct {
	Data    UserModuleUsageGetWechatMiniQrcodeResponseData `json:"data"`
	Message string                                         `json:"message"`
	Success bool                                           `json:"success"`
	JSON    userModuleUsageGetWechatMiniQrcodeResponseJSON `json:"-"`
}

// userModuleUsageGetWechatMiniQrcodeResponseJSON contains the JSON metadata for
// the struct [UserModuleUsageGetWechatMiniQrcodeResponse]
type userModuleUsageGetWechatMiniQrcodeResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserModuleUsageGetWechatMiniQrcodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userModuleUsageGetWechatMiniQrcodeResponseJSON) RawJSON() string {
	return r.raw
}

type UserModuleUsageGetWechatMiniQrcodeResponseData struct {
	CodeURL string                                             `json:"code_url"`
	JSON    userModuleUsageGetWechatMiniQrcodeResponseDataJSON `json:"-"`
}

// userModuleUsageGetWechatMiniQrcodeResponseDataJSON contains the JSON metadata
// for the struct [UserModuleUsageGetWechatMiniQrcodeResponseData]
type userModuleUsageGetWechatMiniQrcodeResponseDataJSON struct {
	CodeURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserModuleUsageGetWechatMiniQrcodeResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userModuleUsageGetWechatMiniQrcodeResponseDataJSON) RawJSON() string {
	return r.raw
}

type UserModuleUsageGetAddWecomeBonusParams struct {
	// 1-智能问诊 2-健康检测 3-用药分析 4-知识问答 5-图像识别
	ModuleType param.Field[int64] `json:"module_type"`
}

func (r UserModuleUsageGetAddWecomeBonusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
