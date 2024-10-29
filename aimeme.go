// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"
	"net/url"

	"github.com/CZL-AI/czlai-go/internal/apijson"
	"github.com/CZL-AI/czlai-go/internal/apiquery"
	"github.com/CZL-AI/czlai-go/internal/param"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// AIMemeService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMemeService] method instead.
type AIMemeService struct {
	Options []option.RequestOption
}

// NewAIMemeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIMemeService(opts ...option.RequestOption) (r *AIMemeService) {
	r = &AIMemeService{}
	r.Options = opts
	return
}

// 获取表情包数据
func (r *AIMemeService) New(ctx context.Context, body AIMemeNewParams, opts ...option.RequestOption) (res *AIMemeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai-meme"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取用户历史表情包数据列表
func (r *AIMemeService) Get(ctx context.Context, query AIMemeGetParams, opts ...option.RequestOption) (res *AIMemeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai-meme"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// 获取 AI 表情包
func (r *AIMemeService) Generate(ctx context.Context, body AIMemeGenerateParams, opts ...option.RequestOption) (res *AIMemeGenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai-meme/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIMeme struct {
	ID        int64         `json:"id"`
	Context   AIMemeContext `json:"context"`
	CreatedAt string        `json:"created_at"`
	// 表情包 url 列表
	MemeImage interface{} `json:"meme_image"`
	// 原始图片列表
	OriginImage string     `json:"origin_image"`
	UpdatedAt   string     `json:"updated_at"`
	JSON        aiMemeJSON `json:"-"`
}

// aiMemeJSON contains the JSON metadata for the struct [AIMeme]
type aiMemeJSON struct {
	ID          apijson.Field
	Context     apijson.Field
	CreatedAt   apijson.Field
	MemeImage   apijson.Field
	OriginImage apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMeme) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMemeJSON) RawJSON() string {
	return r.raw
}

type AIMemeContext struct {
	CaptionList []string `json:"caption_list"`
	// 1 表示是猫或狗，2 表示非猫狗的动植物，0 表示不是动植物。
	IsCatOrDog int64 `json:"is_cat_or_dog"`
	// 1 表示是合法的，0 表示不合法。
	IsLegal int64             `json:"is_legal"`
	JSON    aiMemeContextJSON `json:"-"`
}

// aiMemeContextJSON contains the JSON metadata for the struct [AIMemeContext]
type aiMemeContextJSON struct {
	CaptionList apijson.Field
	IsCatOrDog  apijson.Field
	IsLegal     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMemeContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMemeContextJSON) RawJSON() string {
	return r.raw
}

type AIMemeNewResponse struct {
	Data    AIMeme                `json:"data"`
	Message string                `json:"message"`
	Success bool                  `json:"success"`
	JSON    aiMemeNewResponseJSON `json:"-"`
}

// aiMemeNewResponseJSON contains the JSON metadata for the struct
// [AIMemeNewResponse]
type aiMemeNewResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMemeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMemeNewResponseJSON) RawJSON() string {
	return r.raw
}

type AIMemeGetResponse struct {
	Data    []AIMeme              `json:"data"`
	Message string                `json:"message"`
	Success bool                  `json:"success"`
	JSON    aiMemeGetResponseJSON `json:"-"`
}

// aiMemeGetResponseJSON contains the JSON metadata for the struct
// [AIMemeGetResponse]
type aiMemeGetResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMemeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMemeGetResponseJSON) RawJSON() string {
	return r.raw
}

type AIMemeGenerateResponse struct {
	Data    AIMemeGenerateResponseData `json:"data"`
	Message string                     `json:"message"`
	Success bool                       `json:"success"`
	JSON    aiMemeGenerateResponseJSON `json:"-"`
}

// aiMemeGenerateResponseJSON contains the JSON metadata for the struct
// [AIMemeGenerateResponse]
type aiMemeGenerateResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMemeGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMemeGenerateResponseJSON) RawJSON() string {
	return r.raw
}

type AIMemeGenerateResponseData struct {
	// 表情包地址
	MemeURL string                         `json:"meme_url"`
	JSON    aiMemeGenerateResponseDataJSON `json:"-"`
}

// aiMemeGenerateResponseDataJSON contains the JSON metadata for the struct
// [AIMemeGenerateResponseData]
type aiMemeGenerateResponseDataJSON struct {
	MemeURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMemeGenerateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiMemeGenerateResponseDataJSON) RawJSON() string {
	return r.raw
}

type AIMemeNewParams struct {
	// 图片地址
	ImageURL param.Field[string] `json:"image_url"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AIMemeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIMemeGetParams struct {
	// 每页数量
	Limit param.Field[int64] `query:"limit"`
	// 页数
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [AIMemeGetParams]'s query parameters as `url.Values`.
func (r AIMemeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMemeGenerateParams struct {
	// 文案序号
	ContextIndex param.Field[int64] `json:"context_index"`
	// 表情包 id
	MemeID param.Field[int64] `json:"meme_id"`
}

func (r AIMemeGenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
