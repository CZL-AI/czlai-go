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

// AidocService contains methods and other services that help with interacting with
// the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAidocService] method instead.
type AidocService struct {
	Options []option.RequestOption
}

// NewAidocService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAidocService(opts ...option.RequestOption) (r *AidocService) {
	r = &AidocService{}
	r.Options = opts
	return
}

// 判断是否需要继续提问
func (r *AidocService) IfContinueAsk(ctx context.Context, body AidocIfContinueAskParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "aidoc/if-continue-ask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 判断是否需要传图
func (r *AidocService) IfNeedImage(ctx context.Context, body AidocIfNeedImageParams, opts ...option.RequestOption) (res *AidocIfNeedImageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "aidoc/if-need-image"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取图片结果(流式)
func (r *AidocService) PicResult(ctx context.Context, body AidocPicResultParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aidoc/pic-result"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 发布获取诊断报告任务
func (r *AidocService) Report(ctx context.Context, body AidocReportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aidoc/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *AidocService) ReportTask(ctx context.Context, body AidocReportTaskParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "aidoc/report-task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AidocIfNeedImageResponse = interface{}

type AidocIfContinueAskParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocIfContinueAskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocIfNeedImageParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocIfNeedImageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocPicResultParams struct {
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r AidocPicResultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocReportParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
}

func (r AidocReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AidocReportTaskParams struct {
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 报告类型
	ReportType param.Field[int64] `json:"report_type"`
}

func (r AidocReportTaskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
