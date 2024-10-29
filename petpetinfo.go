// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"
	"net/url"

	"github.com/CZL-AI/czlai-go/internal/apiquery"
	"github.com/CZL-AI/czlai-go/internal/param"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// PetPetInfoService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetPetInfoService] method instead.
type PetPetInfoService struct {
	Options []option.RequestOption
}

// NewPetPetInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPetPetInfoService(opts ...option.RequestOption) (r *PetPetInfoService) {
	r = &PetPetInfoService{}
	r.Options = opts
	return
}

// 宠物数据
func (r *PetPetInfoService) Get(ctx context.Context, query PetPetInfoGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "pets/pet-info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type PetPetInfoGetParams struct {
	// dog cat
	PetsType param.Field[PetPetInfoGetParamsPetsType] `query:"pets_type,required"`
	// 0-分组 1-不分组
	IsSort param.Field[PetPetInfoGetParamsIsSort] `query:"is_sort"`
}

// URLQuery serializes [PetPetInfoGetParams]'s query parameters as `url.Values`.
func (r PetPetInfoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// dog cat
type PetPetInfoGetParamsPetsType string

const (
	PetPetInfoGetParamsPetsTypeDog PetPetInfoGetParamsPetsType = "dog"
	PetPetInfoGetParamsPetsTypeCat PetPetInfoGetParamsPetsType = "cat"
)

func (r PetPetInfoGetParamsPetsType) IsKnown() bool {
	switch r {
	case PetPetInfoGetParamsPetsTypeDog, PetPetInfoGetParamsPetsTypeCat:
		return true
	}
	return false
}

// 0-分组 1-不分组
type PetPetInfoGetParamsIsSort int64

const (
	PetPetInfoGetParamsIsSort0 PetPetInfoGetParamsIsSort = 0
	PetPetInfoGetParamsIsSort1 PetPetInfoGetParamsIsSort = 1
)

func (r PetPetInfoGetParamsIsSort) IsKnown() bool {
	switch r {
	case PetPetInfoGetParamsIsSort0, PetPetInfoGetParamsIsSort1:
		return true
	}
	return false
}
