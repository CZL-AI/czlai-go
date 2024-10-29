// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/czlai-go/internal/apijson"
	"github.com/stainless-sdks/czlai-go/internal/apiquery"
	"github.com/stainless-sdks/czlai-go/internal/param"
	"github.com/stainless-sdks/czlai-go/internal/requestconfig"
	"github.com/stainless-sdks/czlai-go/option"
)

// PetProfileService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetProfileService] method instead.
type PetProfileService struct {
	Options []option.RequestOption
}

// NewPetProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPetProfileService(opts ...option.RequestOption) (r *PetProfileService) {
	r = &PetProfileService{}
	r.Options = opts
	return
}

// 创建宠物档案
func (r *PetProfileService) New(ctx context.Context, body PetProfileNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "pet-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取宠物档案
func (r *PetProfileService) Get(ctx context.Context, query PetProfileGetParams, opts ...option.RequestOption) (res *PetProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "pet-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// 更新宠物档案
func (r *PetProfileService) Update(ctx context.Context, params PetProfileUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "pet-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// 获取宠物档案列表
func (r *PetProfileService) List(ctx context.Context, opts ...option.RequestOption) (res *PetProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "pet-profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// 删除宠物档案
func (r *PetProfileService) Delete(ctx context.Context, body PetProfileDeleteParams, opts ...option.RequestOption) (res *PetProfileDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "pet-profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// 获取宠物档案扩展信息
func (r *PetProfileService) ExInfo(ctx context.Context, body PetProfileExInfoParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "pet-profile/ex-info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取宠物品种
func (r *PetProfileService) Variety(ctx context.Context, body PetProfileVarietyParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "pet-profile/variety"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PetProfile struct {
	ID int64 `json:"id"`
	// 过敏史
	AllergyHistory string `json:"allergy_history,nullable"`
	// 头像
	Avatar string `json:"avatar,nullable"`
	// 生日
	Birthday  string    `json:"birthday"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// 疾病史
	DiseaseHistory string `json:"disease_history,nullable"`
	// 家族史
	FamilyHistory string `json:"family_history,nullable"`
	// 性别 1-公 2-母
	Gender int64 `json:"gender"`
	// 是否已绝育 0-否 1-是
	IsNeutered int64 `json:"is_neutered,nullable"`
	// 是否已接种疫苗 0-否 1-是
	IsVaccination int64 `json:"is_vaccination"`
	// 宠物名字
	Name string `json:"name"`
	// 宠物类型
	PetType int64 `json:"pet_type"`
	// 宠物品种
	PetVariety string    `json:"pet_variety,nullable"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// 疫苗时间
	VaccinationDate string `json:"vaccination_date,nullable"`
	// 重量
	Weight string         `json:"weight,nullable"`
	JSON   petProfileJSON `json:"-"`
}

// petProfileJSON contains the JSON metadata for the struct [PetProfile]
type petProfileJSON struct {
	ID              apijson.Field
	AllergyHistory  apijson.Field
	Avatar          apijson.Field
	Birthday        apijson.Field
	CreatedAt       apijson.Field
	DiseaseHistory  apijson.Field
	FamilyHistory   apijson.Field
	Gender          apijson.Field
	IsNeutered      apijson.Field
	IsVaccination   apijson.Field
	Name            apijson.Field
	PetType         apijson.Field
	PetVariety      apijson.Field
	UpdatedAt       apijson.Field
	VaccinationDate apijson.Field
	Weight          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PetProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petProfileJSON) RawJSON() string {
	return r.raw
}

type PetProfileParam struct {
	// 过敏史
	AllergyHistory param.Field[string] `json:"allergy_history"`
	// 头像
	Avatar param.Field[string] `json:"avatar"`
	// 生日
	Birthday param.Field[string] `json:"birthday"`
	// 疾病史
	DiseaseHistory param.Field[string] `json:"disease_history"`
	// 家族史
	FamilyHistory param.Field[string] `json:"family_history"`
	// 性别 1-公 2-母
	Gender param.Field[int64] `json:"gender"`
	// 是否已绝育 0-否 1-是
	IsNeutered param.Field[int64] `json:"is_neutered"`
	// 是否已接种疫苗 0-否 1-是
	IsVaccination param.Field[int64] `json:"is_vaccination"`
	// 宠物名字
	Name param.Field[string] `json:"name"`
	// 宠物类型
	PetType param.Field[int64] `json:"pet_type"`
	// 宠物品种
	PetVariety param.Field[string] `json:"pet_variety"`
	// 疫苗时间
	VaccinationDate param.Field[string] `json:"vaccination_date"`
	// 重量
	Weight param.Field[string] `json:"weight"`
}

func (r PetProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetProfileGetResponse struct {
	Data    PetProfile                `json:"data"`
	Message string                    `json:"message"`
	Success bool                      `json:"success"`
	JSON    petProfileGetResponseJSON `json:"-"`
}

// petProfileGetResponseJSON contains the JSON metadata for the struct
// [PetProfileGetResponse]
type petProfileGetResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

type PetProfileListResponse struct {
	Data    []PetProfile               `json:"data"`
	Message string                     `json:"message"`
	Success bool                       `json:"success"`
	JSON    petProfileListResponseJSON `json:"-"`
}

// petProfileListResponseJSON contains the JSON metadata for the struct
// [PetProfileListResponse]
type petProfileListResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petProfileListResponseJSON) RawJSON() string {
	return r.raw
}

type PetProfileDeleteResponse struct {
	Data    string                       `json:"data"`
	Message string                       `json:"message"`
	Success bool                         `json:"success"`
	JSON    petProfileDeleteResponseJSON `json:"-"`
}

// petProfileDeleteResponseJSON contains the JSON metadata for the struct
// [PetProfileDeleteResponse]
type petProfileDeleteResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetProfileDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petProfileDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PetProfileNewParams struct {
	PetProfile PetProfileParam `json:"pet_profile,required"`
}

func (r PetProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PetProfile)
}

type PetProfileGetParams struct {
	PetProfileID param.Field[int64] `query:"pet_profile_id,required"`
}

// URLQuery serializes [PetProfileGetParams]'s query parameters as `url.Values`.
func (r PetProfileGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetProfileUpdateParams struct {
	PetProfileID param.Field[int64] `query:"pet_profile_id,required"`
	PetProfile   PetProfileParam    `json:"pet_profile,required"`
}

func (r PetProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PetProfile)
}

// URLQuery serializes [PetProfileUpdateParams]'s query parameters as `url.Values`.
func (r PetProfileUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetProfileDeleteParams struct {
	PetProfileID param.Field[int64] `query:"pet_profile_id,required"`
}

// URLQuery serializes [PetProfileDeleteParams]'s query parameters as `url.Values`.
func (r PetProfileDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetProfileExInfoParams struct {
	PetProfileID param.Field[int64] `json:"pet_profile_id"`
}

func (r PetProfileExInfoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetProfileVarietyParams struct {
	UserInput param.Field[string] `json:"user_input"`
}

func (r PetProfileVarietyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
