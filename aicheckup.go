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

// AICheckupService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAICheckupService] method instead.
type AICheckupService struct {
	Options []option.RequestOption
}

// NewAICheckupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAICheckupService(opts ...option.RequestOption) (r *AICheckupService) {
	r = &AICheckupService{}
	r.Options = opts
	return
}

// 检查是否为当月首检
func (r *AICheckupService) IsFirst(ctx context.Context, query AICheckupIsFirstParams, opts ...option.RequestOption) (res *AICheckupIsFirstResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai-checkup/is-first"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// 获取图片结果
func (r *AICheckupService) PicResult(ctx context.Context, body AICheckupPicResultParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-checkup/pic-result"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取问题
func (r *AICheckupService) Question(ctx context.Context, body AICheckupQuestionParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-checkup/question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取问题
func (r *AICheckupService) QuestionResult(ctx context.Context, body AICheckupQuestionResultParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-checkup/question-result"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AICheckupService) Report(ctx context.Context, body AICheckupReportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-checkup/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AICheckupService) ReportTask(ctx context.Context, body AICheckupReportTaskParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ai-checkup/report-task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 开始一个新的会话
func (r *AICheckupService) SessionStart(ctx context.Context, opts ...option.RequestOption) (res *AICheckupSessionStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai-checkup/session-start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// 生成总结
func (r *AICheckupService) Summary(ctx context.Context, body AICheckupSummaryParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "ai-checkup/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 更新宠物档案的体检参数
func (r *AICheckupService) UpdateProfile(ctx context.Context, body AICheckupUpdateProfileParams, opts ...option.RequestOption) (res *AICheckupUpdateProfileResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai-checkup/update-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AICheckupIsFirstResponse struct {
	Data    AICheckupIsFirstResponseData `json:"data"`
	Message string                       `json:"message"`
	Success bool                         `json:"success"`
	JSON    aiCheckupIsFirstResponseJSON `json:"-"`
}

// aiCheckupIsFirstResponseJSON contains the JSON metadata for the struct
// [AICheckupIsFirstResponse]
type aiCheckupIsFirstResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AICheckupIsFirstResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiCheckupIsFirstResponseJSON) RawJSON() string {
	return r.raw
}

type AICheckupIsFirstResponseData struct {
	// 是否为当月首检
	IsFirst bool                             `json:"is_first"`
	JSON    aiCheckupIsFirstResponseDataJSON `json:"-"`
}

// aiCheckupIsFirstResponseDataJSON contains the JSON metadata for the struct
// [AICheckupIsFirstResponseData]
type aiCheckupIsFirstResponseDataJSON struct {
	IsFirst     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AICheckupIsFirstResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiCheckupIsFirstResponseDataJSON) RawJSON() string {
	return r.raw
}

type AICheckupSessionStartResponse struct {
	Data    AICheckupSessionStartResponseData `json:"data"`
	Message string                            `json:"message"`
	Success bool                              `json:"success"`
	JSON    aiCheckupSessionStartResponseJSON `json:"-"`
}

// aiCheckupSessionStartResponseJSON contains the JSON metadata for the struct
// [AICheckupSessionStartResponse]
type aiCheckupSessionStartResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AICheckupSessionStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiCheckupSessionStartResponseJSON) RawJSON() string {
	return r.raw
}

type AICheckupSessionStartResponseData struct {
	SessionID string                                `json:"session_id"`
	JSON      aiCheckupSessionStartResponseDataJSON `json:"-"`
}

// aiCheckupSessionStartResponseDataJSON contains the JSON metadata for the struct
// [AICheckupSessionStartResponseData]
type aiCheckupSessionStartResponseDataJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AICheckupSessionStartResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiCheckupSessionStartResponseDataJSON) RawJSON() string {
	return r.raw
}

type AICheckupUpdateProfileResponse struct {
	Data    interface{}                        `json:"data"`
	Message string                             `json:"message"`
	Success bool                               `json:"success"`
	JSON    aiCheckupUpdateProfileResponseJSON `json:"-"`
}

// aiCheckupUpdateProfileResponseJSON contains the JSON metadata for the struct
// [AICheckupUpdateProfileResponse]
type aiCheckupUpdateProfileResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AICheckupUpdateProfileResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiCheckupUpdateProfileResponseJSON) RawJSON() string {
	return r.raw
}

type AICheckupIsFirstParams struct {
	// 宠物档案 ID
	PetProfileID param.Field[int64] `query:"pet_profile_id,required"`
}

// URLQuery serializes [AICheckupIsFirstParams]'s query parameters as `url.Values`.
func (r AICheckupIsFirstParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AICheckupPicResultParams struct {
	// 图片 url
	ImgURL param.Field[string] `json:"img_url,required"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id,required"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
}

func (r AICheckupPicResultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AICheckupQuestionParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id,required"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
}

func (r AICheckupQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AICheckupQuestionResultParams struct {
	// 宠物档案 id
	Index param.Field[int64] `json:"index,required"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id,required"`
	// 回答 id
	QuestionID param.Field[string] `json:"question_id,required"`
	// 题目轮次
	Round param.Field[string] `json:"round,required"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
}

func (r AICheckupQuestionResultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AICheckupReportParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id,required"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
}

func (r AICheckupReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AICheckupReportTaskParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 报告类型
	ReportType param.Field[int64] `json:"report_type"`
}

func (r AICheckupReportTaskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AICheckupSummaryParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AICheckupSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AICheckupUpdateProfileParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
	// 更新类型, 可选 1,2,3
	UpdateType param.Field[int64] `json:"update_type"`
}

func (r AICheckupUpdateProfileParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
