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

// AipicService contains methods and other services that help with interacting with
// the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAipicService] method instead.
type AipicService struct {
	Options []option.RequestOption
}

// NewAipicService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAipicService(opts ...option.RequestOption) (r *AipicService) {
	r = &AipicService{}
	r.Options = opts
	return
}

// 获取问题选项
func (r *AipicService) Options(ctx context.Context, body AipicOptionsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic/options"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取图片结果(流式)
func (r *AipicService) PicResult(ctx context.Context, body AipicPicResultParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic/pic-result"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取问题(流式)
func (r *AipicService) Question(ctx context.Context, body AipicQuestionParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic/question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AipicService) Report(ctx context.Context, body AipicReportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AipicService) ReportTask(ctx context.Context, body AipicReportTaskParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic/report-task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 验证答案是否有效
func (r *AipicService) Validate(ctx context.Context, body AipicValidateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aipic/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AipicOptionsParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 问题
	Question param.Field[string] `json:"question"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicPicResultParams struct {
	// 图片归属(1:宠物体态分析、2:宠物表情分析)
	ImgBelong param.Field[int64] `json:"img_belong"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicPicResultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicQuestionParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicReportParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 图片归属(1:宠物体态分析、2:宠物表情分析、3:牙齿分析)
	SubModuleType param.Field[int64] `json:"sub_module_type"`
}

func (r AipicReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicReportTaskParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 报告类型
	ReportType param.Field[int64] `json:"report_type"`
}

func (r AipicReportTaskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AipicValidateParams struct {
	// 用户回答
	Answer param.Field[string] `json:"answer"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 问题
	Question param.Field[string] `json:"question"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AipicValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
