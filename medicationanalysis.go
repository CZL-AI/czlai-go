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

// MedicationAnalysisService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMedicationAnalysisService] method instead.
type MedicationAnalysisService struct {
	Options []option.RequestOption
}

// NewMedicationAnalysisService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMedicationAnalysisService(opts ...option.RequestOption) (r *MedicationAnalysisService) {
	r = &MedicationAnalysisService{}
	r.Options = opts
	return
}

// 获取图片结果(流式)
func (r *MedicationAnalysisService) PicResult(ctx context.Context, body MedicationAnalysisPicResultParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "medication_analysis/pic-result"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取诊断报告
func (r *MedicationAnalysisService) Report(ctx context.Context, body MedicationAnalysisReportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "medication_analysis/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取病情小结(流式)
func (r *MedicationAnalysisService) Summary(ctx context.Context, body MedicationAnalysisSummaryParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "medication_analysis/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type MedicationAnalysisPicResultParams struct {
	// 图片归属(1:宠物体态分析、2:宠物表情分析、3:牙齿分析)
	ImgBelong param.Field[int64] `json:"img_belong"`
	// 图片 url
	ImgURL param.Field[string] `json:"img_url"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r MedicationAnalysisPicResultParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MedicationAnalysisReportParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r MedicationAnalysisReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MedicationAnalysisSummaryParams struct {
	// 宠物档案 id
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id"`
}

func (r MedicationAnalysisSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
