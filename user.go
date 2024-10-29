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

// UserService contains methods and other services that help with interacting with
// the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options  []option.RequestOption
	UserInfo *UserUserInfoService
	Contact  *UserContactService
	Summary  *UserSummaryService
	Asr      *UserAsrService
	Industry *UserIndustryService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.UserInfo = NewUserUserInfoService(opts...)
	r.Contact = NewUserContactService(opts...)
	r.Summary = NewUserSummaryService(opts...)
	r.Asr = NewUserAsrService(opts...)
	r.Industry = NewUserIndustryService(opts...)
	return
}

// AI 图片聊天
func (r *UserService) ChatV(ctx context.Context, body UserChatVParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "chat-v"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Logs out the user
func (r *UserService) Logout(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// 发验证短信
func (r *UserService) SendSMS(ctx context.Context, body UserSendSMSParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "send-sms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserChatVParams struct {
	Content   param.Field[string] `json:"content"`
	SessionID param.Field[string] `json:"session_id"`
}

func (r UserChatVParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSendSMSParams struct {
	Phone param.Field[string] `json:"phone"`
}

func (r UserSendSMSParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
