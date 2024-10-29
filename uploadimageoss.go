// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/czlai-go/internal/apiform"
	"github.com/stainless-sdks/czlai-go/internal/apiquery"
	"github.com/stainless-sdks/czlai-go/internal/param"
	"github.com/stainless-sdks/czlai-go/internal/requestconfig"
	"github.com/stainless-sdks/czlai-go/option"
)

// UploadImageOssService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadImageOssService] method instead.
type UploadImageOssService struct {
	Options []option.RequestOption
}

// NewUploadImageOssService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUploadImageOssService(opts ...option.RequestOption) (r *UploadImageOssService) {
	r = &UploadImageOssService{}
	r.Options = opts
	return
}

// 允许用户上传一张图片
func (r *UploadImageOssService) New(ctx context.Context, params UploadImageOssNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "upload-image-oss"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type UploadImageOssNewParams struct {
	// 图片上传类型 1-头像 2-图片识别模块 3-表情包 4-其他
	UploadType param.Field[UploadImageOssNewParamsUploadType] `query:"upload_type,required"`
	// 要上传的图片文件
	Image param.Field[io.Reader] `json:"image,required" format:"binary"`
	// 是否上传到本地服务器
	UploadToLocal param.Field[int64] `query:"upload_to_local"`
}

func (r UploadImageOssNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [UploadImageOssNewParams]'s query parameters as
// `url.Values`.
func (r UploadImageOssNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// 图片上传类型 1-头像 2-图片识别模块 3-表情包 4-其他
type UploadImageOssNewParamsUploadType int64

const (
	UploadImageOssNewParamsUploadType1 UploadImageOssNewParamsUploadType = 1
	UploadImageOssNewParamsUploadType2 UploadImageOssNewParamsUploadType = 2
	UploadImageOssNewParamsUploadType3 UploadImageOssNewParamsUploadType = 3
	UploadImageOssNewParamsUploadType4 UploadImageOssNewParamsUploadType = 4
)

func (r UploadImageOssNewParamsUploadType) IsKnown() bool {
	switch r {
	case UploadImageOssNewParamsUploadType1, UploadImageOssNewParamsUploadType2, UploadImageOssNewParamsUploadType3, UploadImageOssNewParamsUploadType4:
		return true
	}
	return false
}
