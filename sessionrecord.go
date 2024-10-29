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

// SessionRecordService contains methods and other services that help with
// interacting with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionRecordService] method instead.
type SessionRecordService struct {
	Options []option.RequestOption
}

// NewSessionRecordService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSessionRecordService(opts ...option.RequestOption) (r *SessionRecordService) {
	r = &SessionRecordService{}
	r.Options = opts
	return
}

// 保存聊天记录
func (r *SessionRecordService) History(ctx context.Context, body SessionRecordHistoryParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "session-record/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type SessionRecordHistoryParams struct {
	// 内容
	Content param.Field[string] `json:"content,required"`
	// 1-智能问诊 2-健康检测 3-用药分析 4-知识问答 5-图像识别
	ModuleType param.Field[int64] `json:"module_type,required"`
	// 角色, 取值为其中之一 ==>[user, ai]
	Role param.Field[string] `json:"role,required"`
	// 会话 id
	SessionID param.Field[string] `json:"session_id,required"`
	// 1-文字 2-图文
	ContentType param.Field[int64] `json:"content_type"`
	// 1-用户主诉 2-用户回答 3-AI 提问 4-AI 病情小结 5-AI 病例报告 6-AI 输出 7-用户补充
	Stage param.Field[int64] `json:"stage"`
}

func (r SessionRecordHistoryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
