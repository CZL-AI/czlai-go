// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/CZL-AI/czlai-go/internal/apijson"
	"github.com/CZL-AI/czlai-go/internal/apiquery"
	"github.com/CZL-AI/czlai-go/internal/param"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// MedicalRecordService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMedicalRecordService] method instead.
type MedicalRecordService struct {
	Options []option.RequestOption
}

// NewMedicalRecordService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMedicalRecordService(opts ...option.RequestOption) (r *MedicalRecordService) {
	r = &MedicalRecordService{}
	r.Options = opts
	return
}

// 获取单个病例报告
func (r *MedicalRecordService) Get(ctx context.Context, query MedicalRecordGetParams, opts ...option.RequestOption) (res *MedicalRecordGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "medical-record"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// 修改状态和记录对话进行阶段
func (r *MedicalRecordService) Update(ctx context.Context, body MedicalRecordUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "medical-record"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// 获取医疗报告列表
func (r *MedicalRecordService) NewList(ctx context.Context, body MedicalRecordNewListParams, opts ...option.RequestOption) (res *MedicalRecordNewListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "medical-record-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 获取进行中的医疗报告
func (r *MedicalRecordService) OngoingRecord(ctx context.Context, query MedicalRecordOngoingRecordParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "medical-record/ongoing-record"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type MedicalRecord struct {
	// 主键 ID
	ID int64 `json:"id"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// 模块 1-智能问诊 2-健康检测
	ModuleType int64 `json:"module_type,nullable"`
	// 报告
	Report string `json:"report,nullable"`
	// 对应的 session_id
	SessionID string `json:"session_id,nullable"`
	// 小结
	Summary string `json:"summary,nullable"`
	// 用户 uuid
	UserUuid string            `json:"user_uuid,nullable"`
	JSON     medicalRecordJSON `json:"-"`
}

// medicalRecordJSON contains the JSON metadata for the struct [MedicalRecord]
type medicalRecordJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ModuleType  apijson.Field
	Report      apijson.Field
	SessionID   apijson.Field
	Summary     apijson.Field
	UserUuid    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalRecordJSON) RawJSON() string {
	return r.raw
}

type MedicalRecordGetResponse struct {
	Data    MedicalRecord                `json:"data"`
	Message string                       `json:"message"`
	Success bool                         `json:"success"`
	JSON    medicalRecordGetResponseJSON `json:"-"`
}

// medicalRecordGetResponseJSON contains the JSON metadata for the struct
// [MedicalRecordGetResponse]
type medicalRecordGetResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalRecordGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalRecordGetResponseJSON) RawJSON() string {
	return r.raw
}

type MedicalRecordNewListResponse struct {
	Data    []MedicalRecord                  `json:"data"`
	Message string                           `json:"message"`
	Success bool                             `json:"success"`
	JSON    medicalRecordNewListResponseJSON `json:"-"`
}

// medicalRecordNewListResponseJSON contains the JSON metadata for the struct
// [MedicalRecordNewListResponse]
type medicalRecordNewListResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalRecordNewListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalRecordNewListResponseJSON) RawJSON() string {
	return r.raw
}

type MedicalRecordGetParams struct {
	// 报告 ID
	ReportID param.Field[int64] `query:"report_id,required"`
}

// URLQuery serializes [MedicalRecordGetParams]'s query parameters as `url.Values`.
func (r MedicalRecordGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MedicalRecordUpdateParams struct {
	IsRead param.Field[int64] `json:"is_read"`
	// 医疗报告 ID
	ReportID   param.Field[int64]     `json:"report_id"`
	ReportTime param.Field[time.Time] `json:"report_time" format:"date-time"`
	// 阶段存档
	Stage param.Field[string] `json:"stage"`
	// 1-进行中 2-已完成 3-手动结束 4-自动结束
	Status param.Field[int64] `json:"status"`
}

func (r MedicalRecordUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MedicalRecordNewListParams struct {
	// 每页数量
	Limit      param.Field[int64]   `json:"limit"`
	ModuleType param.Field[[]int64] `json:"module_type"`
	// 页码
	Page param.Field[int64] `json:"page"`
	// 宠物档案 id
	PetProfileID param.Field[[]int64] `json:"pet_profile_id"`
}

func (r MedicalRecordNewListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MedicalRecordOngoingRecordParams struct {
	// 模块类型 1-智能问诊 2-健康检测 3-用药分析 4-知识问答 5-图像识别
	ModuleType param.Field[int64] `query:"module_type,required"`
	// 宠物档案 id
	PetProfileID param.Field[int64] `query:"pet_profile_id,required"`
}

// URLQuery serializes [MedicalRecordOngoingRecordParams]'s query parameters as
// `url.Values`.
func (r MedicalRecordOngoingRecordParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
