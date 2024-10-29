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

// AipicExoticService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAipicExoticService] method instead.
type AipicExoticService struct {
	Options []option.RequestOption
}

// NewAipicExoticService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAipicExoticService(opts ...option.RequestOption) (r *AipicExoticService) {
	r = &AipicExoticService{}
	r.Options = opts
	return
}

// 获取问题选项
func (r *AipicExoticService) Options(ctx context.Context, body AipicExoticOptionsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic-exotic/options"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取图片结果(流式)
func (r *AipicExoticService) PicResult(ctx context.Context, body AipicExoticPicResultParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic-exotic/pic-result"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取问题(流式)
func (r *AipicExoticService) Question(ctx context.Context, body AipicExoticQuestionParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic-exotic/question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AipicExoticService) Report(ctx context.Context, body AipicExoticReportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic-exotic/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AipicExoticService) ReportTask(ctx context.Context, body AipicExoticReportTaskParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic-exotic/report-task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取小结(流式)
func (r *AipicExoticService) Summary(ctx context.Context, body AipicExoticSummaryParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic-exotic/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 验证答案是否有效
func (r *AipicExoticService) Validate(ctx context.Context, body AipicExoticValidateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic-exotic/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AipicExoticOptionsParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 问题
	Question param.Field[string] `json:"question"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicExoticOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicExoticPicResultParams struct {
	// 图片归属(1:宠物品种分析、2:宠物表情分析)
	ImgBelong param.Field[int64] `json:"img_belong"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicExoticPicResultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicExoticQuestionParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicExoticQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicExoticReportParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 图片归属(1:宠物体态分析、2:宠物表情分析、3:牙齿分析)
	SubModuleType param.Field[int64] `json:"sub_module_type"`
}

func (r AipicExoticReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicExoticReportTaskParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 报告类型
	ReportType param.Field[int64] `json:"report_type"`
}

func (r AipicExoticReportTaskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicExoticSummaryParams struct {
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
	// 图片归属(1:宠物体态分析、2:宠物表情分析)
	SubModuleType param.Field[int64] `json:"sub_module_type"`
}

func (r AipicExoticSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicExoticValidateParams struct {
	// 用户回答
	Answer param.Field[string] `json:"answer"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 问题
	Question param.Field[string] `json:"question"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicExoticValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
