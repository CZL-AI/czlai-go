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

// MagnumKeyService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMagnumKeyService] method instead.
type MagnumKeyService struct {
	Options []option.RequestOption
}

// NewMagnumKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMagnumKeyService(opts ...option.RequestOption) (r *MagnumKeyService) {
	r = &MagnumKeyService{}
	r.Options = opts
	return
}

// 获取 key_usage_id
func (r *MagnumKeyService) GetKey(ctx context.Context, body MagnumKeyGetKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "magnumkey/get-key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取图片选项
func (r *MagnumKeyService) PictureChoice(ctx context.Context, body MagnumKeyPictureChoiceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "magnumkey/picture-choice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取图片问题
func (r *MagnumKeyService) PictureQuestion(ctx context.Context, body MagnumKeyPictureQuestionParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "magnumkey/picture-question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取声音选项
func (r *MagnumKeyService) VoiceChoice(ctx context.Context, body MagnumKeyVoiceChoiceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "magnumkey/voice-choice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 获取声音问题
func (r *MagnumKeyService) VoiceQuestion(ctx context.Context, body MagnumKeyVoiceQuestionParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "magnumkey/voice-question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type MagnumKeyGetKeyParams struct {
	// 文本
	Context param.Field[string] `json:"context"`
}

func (r MagnumKeyGetKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagnumKeyPictureChoiceParams struct {
	// 图片 url
	ImgURL param.Field[string] `json:"img_url,required"`
	// 会话 id
	KeyUsageID param.Field[string] `json:"key_usage_id"`
	// 用户 uuid
	UserUuid param.Field[string] `json:"user_uuid"`
}

func (r MagnumKeyPictureChoiceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagnumKeyPictureQuestionParams struct {
	// 图片 url
	ImgURL param.Field[string] `json:"img_url,required"`
	// 会话 id
	KeyUsageID param.Field[string] `json:"key_usage_id"`
	// 用户 uuid
	UserUuid param.Field[string] `json:"user_uuid"`
}

func (r MagnumKeyPictureQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagnumKeyVoiceChoiceParams struct {
	// 获取声音选项
	Message param.Field[string] `json:"message,required"`
	// 会话 id
	KeyUsageID param.Field[string] `json:"key_usage_id"`
	// 用户 uuid
	UserUuid param.Field[string] `json:"user_uuid"`
}

func (r MagnumKeyVoiceChoiceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagnumKeyVoiceQuestionParams struct {
	// 语音文本
	Message param.Field[string] `json:"message,required"`
	// 会话 id
	KeyUsageID param.Field[string] `json:"key_usage_id"`
	// 宠物 id
	PetID param.Field[int64] `json:"pet_id"`
	// 用户 uuid
	UserUuid param.Field[string] `json:"user_uuid"`
}

func (r MagnumKeyVoiceQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
