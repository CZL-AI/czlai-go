// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/CZL-AI/czlai-go/internal/apiform"
	"github.com/CZL-AI/czlai-go/internal/apiquery"
	"github.com/CZL-AI/czlai-go/internal/param"
	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// UploadImageService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadImageService] method instead.
type UploadImageService struct {
	Options []option.RequestOption
}

// NewUploadImageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUploadImageService(opts ...option.RequestOption) (r *UploadImageService) {
	r = &UploadImageService{}
	r.Options = opts
	return
}

// 允许用户上传一张图片
func (r *UploadImageService) New(ctx context.Context, params UploadImageNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "upload-image"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type UploadImageNewParams struct {
	// 要上传的图片文件
	Image param.Field[io.Reader] `json:"image,required" format:"binary"`
	// 是否上传到图床
	IsToCloud param.Field[int64] `query:"is_to_cloud"`
	// 图片上传类型 1-头像 2-图片识别模块 3-表情包 4-其他
	UploadType param.Field[UploadImageNewParamsUploadType] `query:"upload_type"`
}

func (r UploadImageNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [UploadImageNewParams]'s query parameters as `url.Values`.
func (r UploadImageNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// 图片上传类型 1-头像 2-图片识别模块 3-表情包 4-其他
type UploadImageNewParamsUploadType int64

const (
	UploadImageNewParamsUploadType1 UploadImageNewParamsUploadType = 1
	UploadImageNewParamsUploadType2 UploadImageNewParamsUploadType = 2
	UploadImageNewParamsUploadType3 UploadImageNewParamsUploadType = 3
	UploadImageNewParamsUploadType4 UploadImageNewParamsUploadType = 4
)

func (r UploadImageNewParamsUploadType) IsKnown() bool {
	switch r {
	case UploadImageNewParamsUploadType1, UploadImageNewParamsUploadType2, UploadImageNewParamsUploadType3, UploadImageNewParamsUploadType4:
		return true
	}
	return false
}
