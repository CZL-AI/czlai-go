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

// LoginService contains methods and other services that help with interacting with
// the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoginService] method instead.
type LoginService struct {
	Options []option.RequestOption
}

// NewLoginService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLoginService(opts ...option.RequestOption) (r *LoginService) {
	r = &LoginService{}
	r.Options = opts
	return
}

// 短信登录
func (r *LoginService) SMS(ctx context.Context, body LoginSMSParams, opts ...option.RequestOption) (res *LoginResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sms-login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Logs in the user by WeChat
func (r *LoginService) Wechat(ctx context.Context, body LoginWechatParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "wechat-login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type LoginResponse struct {
	AuthTokens LoginResponseAuthTokens `json:"authTokens"`
	Message    string                  `json:"message"`
	Success    bool                    `json:"success"`
	JSON       loginResponseJSON       `json:"-"`
}

// loginResponseJSON contains the JSON metadata for the struct [LoginResponse]
type loginResponseJSON struct {
	AuthTokens  apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginResponseJSON) RawJSON() string {
	return r.raw
}

type LoginResponseAuthTokens struct {
	AccessToken  string                      `json:"accessToken,required"`
	RefreshToken string                      `json:"refreshToken,required"`
	JSON         loginResponseAuthTokensJSON `json:"-"`
}

// loginResponseAuthTokensJSON contains the JSON metadata for the struct
// [LoginResponseAuthTokens]
type loginResponseAuthTokensJSON struct {
	AccessToken  apijson.Field
	RefreshToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoginResponseAuthTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginResponseAuthTokensJSON) RawJSON() string {
	return r.raw
}

type LoginSMSParams struct {
	Code  param.Field[string] `json:"code,required"`
	Phone param.Field[string] `json:"phone,required"`
	// 1-微信小程序 2-安卓 APP 3-IOS APP
	LoginFrom param.Field[int64] `json:"login_from"`
}

func (r LoginSMSParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoginWechatParams struct {
	// 会话 id
	WechatCode param.Field[string] `json:"wechat_code,required"`
	// 加密数据
	EncryptedData param.Field[string] `json:"encryptedData"`
	// 加密初始向量
	Iv param.Field[string] `json:"iv"`
	// 模块类型 1-智能问诊 2-健康检测 3-用药分析 4-知识问答 5-图片识别
	ModuleType param.Field[int64] `json:"module_type"`
	// 手机号
	PhoneNumber param.Field[string] `json:"phone_number"`
	// 推广人 sid
	SpreadID param.Field[int64] `json:"spread_id"`
}

func (r LoginWechatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
