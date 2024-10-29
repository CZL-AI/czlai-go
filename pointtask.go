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

// PointTaskService contains methods and other services that help with interacting
// with the czlai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPointTaskService] method instead.
type PointTaskService struct {
	Options []option.RequestOption
}

// NewPointTaskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPointTaskService(opts ...option.RequestOption) (r *PointTaskService) {
	r = &PointTaskService{}
	r.Options = opts
	return
}

// 获取积分任务列表
func (r *PointTaskService) List(ctx context.Context, opts ...option.RequestOption) (res *PointTaskListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "point-task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// 领取奖励
func (r *PointTaskService) Bonus(ctx context.Context, body PointTaskBonusParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "point-task/bonus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// 确认积分任务
func (r *PointTaskService) Confirm(ctx context.Context, body PointTaskConfirmParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "point-task/confirm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type PointTaskListResponse struct {
	Data []PointTaskListResponseData `json:"data"`
	// message
	Message string `json:"message"`
	// success
	Success bool                      `json:"success"`
	JSON    pointTaskListResponseJSON `json:"-"`
}

// pointTaskListResponseJSON contains the JSON metadata for the struct
// [PointTaskListResponse]
type pointTaskListResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PointTaskListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pointTaskListResponseJSON) RawJSON() string {
	return r.raw
}

type PointTaskListResponseData struct {
	// id
	ID int64 `json:"id"`
	// 可完成次数
	AchieveCount int64 `json:"achieve_count"`
	// 积分奖励
	BonusPoint string `json:"bonus_point"`
	// 完成条件次数
	ConditionCount int64 `json:"condition_count"`
	// 任务说明
	Description string `json:"description"`
	// 任务图标
	Icon string `json:"icon"`
	// 0-未开启 1-开启
	IsOpen int64 `json:"is_open"`
	// 关联模块
	RelatedModule string `json:"related_module"`
	// 1-未完成 2-未领取
	Status int64 `json:"status"`
	// 任务动作
	TaskAction string `json:"task_action"`
	// 任务名称
	TaskName string `json:"task_name"`
	// 跳转页面
	ToPage string                        `json:"to_page"`
	JSON   pointTaskListResponseDataJSON `json:"-"`
}

// pointTaskListResponseDataJSON contains the JSON metadata for the struct
// [PointTaskListResponseData]
type pointTaskListResponseDataJSON struct {
	ID             apijson.Field
	AchieveCount   apijson.Field
	BonusPoint     apijson.Field
	ConditionCount apijson.Field
	Description    apijson.Field
	Icon           apijson.Field
	IsOpen         apijson.Field
	RelatedModule  apijson.Field
	Status         apijson.Field
	TaskAction     apijson.Field
	TaskName       apijson.Field
	ToPage         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PointTaskListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pointTaskListResponseDataJSON) RawJSON() string {
	return r.raw
}

type PointTaskBonusParams struct {
	// 任务 id
	TaskID param.Field[int64] `json:"task_id"`
}

func (r PointTaskBonusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PointTaskConfirmParams struct {
	// 任务 id
	TaskID param.Field[int64] `json:"task_id"`
}

func (r PointTaskConfirmParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
