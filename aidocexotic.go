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

// AidocExoticService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAidocExoticService] method instead.
type AidocExoticService struct {
	Options []option.RequestOption
}

// NewAidocExoticService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAidocExoticService(opts ...option.RequestOption) (r *AidocExoticService) {
	r = &AidocExoticService{}
	r.Options = opts
	return
}

// 判断是否需要继续提问
func (r *AidocExoticService) AskContinue(ctx context.Context, body AidocExoticAskContinueParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "aidoc-exotic/if-continue-ask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 判断是否需要传图
func (r *AidocExoticService) IfNeedImage(ctx context.Context, body AidocExoticIfNeedImageParams, opts ...option.RequestOption) (res *AidocExoticIfNeedImageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "aidoc-exotic/if-need-image"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取关键词,科室
func (r *AidocExoticService) Keywords(ctx context.Context, body AidocExoticKeywordsParams, opts ...option.RequestOption) (res *AidocExoticKeywordsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "aidoc-exotic/keywords"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取问题选项
func (r *AidocExoticService) Options(ctx context.Context, body AidocExoticOptionsParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "aidoc-exotic/options"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取图片结果(流式)
func (r *AidocExoticService) PicResult(ctx context.Context, body AidocExoticPicResultParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aidoc-exotic/pic-result"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取问题(流式)
func (r *AidocExoticService) Question(ctx context.Context, body AidocExoticQuestionParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "aidoc-exotic/question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 发布获取诊断报告任务
func (r *AidocExoticService) Report(ctx context.Context, body AidocExoticReportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aidoc-exotic/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AidocExoticService) ReportTask(ctx context.Context, body AidocExoticReportTaskParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aidoc-exotic/report-task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取病情小结(流式)
func (r *AidocExoticService) Summarize(ctx context.Context, body AidocExoticSummarizeParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aidoc-exotic/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AidocExoticIfNeedImageResponse = interface{}

type AidocExoticKeywordsResponse struct {
	Data    AidocExoticKeywordsResponseData `json:"data"`
	Message string                          `json:"message"`
	Success bool                            `json:"success"`
	JSON    aidocExoticKeywordsResponseJSON `json:"-"`
}

// aidocExoticKeywordsResponseJSON contains the JSON metadata for the struct
// [AidocExoticKeywordsResponse]
type aidocExoticKeywordsResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AidocExoticKeywordsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aidocExoticKeywordsResponseJSON) RawJSON() string {
	return r.raw
}

type AidocExoticKeywordsResponseData struct {
	// 关键词
	Keywords string `json:"keywords"`
	// 科室
	Unit string                              `json:"unit"`
	JSON aidocExoticKeywordsResponseDataJSON `json:"-"`
}

// aidocExoticKeywordsResponseDataJSON contains the JSON metadata for the struct
// [AidocExoticKeywordsResponseData]
type aidocExoticKeywordsResponseDataJSON struct {
	Keywords    apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AidocExoticKeywordsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aidocExoticKeywordsResponseDataJSON) RawJSON() string {
	return r.raw
}

type AidocExoticAskContinueParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocExoticAskContinueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticIfNeedImageParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocExoticIfNeedImageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticKeywordsParams struct {
	// 用户主诉
	Content param.Field[string] `json:"content"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocExoticKeywordsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticOptionsParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 问题
	Question param.Field[string] `json:"question"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocExoticOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticPicResultParams struct {
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocExoticPicResultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticQuestionParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocExoticQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticReportParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
}

func (r AidocExoticReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticReportTaskParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 报告类型
	ReportType param.Field[int64] `json:"report_type"`
}

func (r AidocExoticReportTaskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocExoticSummarizeParams struct {
	// 图片地址
	ImageURL param.Field[string] `json:"image_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocExoticSummarizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
