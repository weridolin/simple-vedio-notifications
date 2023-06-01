// Code generated by goctl. DO NOT EDIT.
package types

type UserInfo struct {
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Avatar       string   `json:"avatar"`
	Role         []string `json:"role"`
	IsSuperAdmin bool     `json:"is_super_admin"`
	Age          int      `json:"age"`
	Gender       int8     `json:"gender"`
}

type Task struct {
	ID          int                    `json:"id"`
	User        UserInfo               `json:"user"`
	PlatForm    string                 `json:"platform"`
	Ups         map[string]interface{} `json:"ups" desc:"up主"`
	Schedulers  []Scheduler            `json:"schedulers"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
}

type CreateTaskReq struct {
	Platform    string                 `json:"platform" desc:"平台"`
	Ups         map[string]interface{} `json:"ups" desc:"up主"`
	Schedulers  []int                  `json:"schedulers" desc:"绑定的运行计划列表"`
	Name        string                 `json:"name" desc:"任务名称"`
	Description string                 `json:"description" desc:"任务描述"`
}

type CreateTaskResp struct {
	BaseResponse
	Data Task `json:"data"`
}

type UpdateTaskReq struct {
	CreateTaskReq
}

type UpdateTaskResp struct {
	CreateTaskResp
}

type DeleteTaskReq struct {
	ID int `json:"id"`
}

type DeleteTaskResp struct {
	BaseResponse
	Data Task `json:"data,optional"`
}

type QueryTaskReq struct {
	PaginationParams
}

type QueryTaskResp struct {
	BaseResponse
	Data []Task `json:"data"`
}

type BindSchedulerReq struct {
	SchedulerID int   `json:"scheduler_id"`
	TaskID      []int `json:"task_id"`
}

type BindSchedulerResp struct {
	BaseResponse
	Data []Task `json:"data,optional"`
}

type Scheduler struct {
	ID          int      `json:"id"`
	Period      string   `json:"period"`
	Active      bool     `json:"active"`
	User        UserInfo `json:"user"`
	Tasks       []Task   `json:"tasks"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
}

type CreateSchedulerReq struct {
	Period      string `json:"period" desc:"运行周期"`
	Platform    string `json:"platform" desc:"平台"`
	Name        string `json:"name" desc:"计划名称"`
	Description string `json:"description" desc:"计划描述"`
}

type CreateSchedulerResp struct {
	BaseResponse
	Data Scheduler `json:"data"`
}

type UpdateSchedulerReq struct {
	CreateSchedulerReq
}

type UpdateSchedulerResp struct {
	CreateSchedulerResp
}

type DeleteSchedulerReq struct {
	Id int `path:"id"`
}

type DeleteSchedulerResp struct {
	BaseResponse
}

type QuerySchedulerReq struct {
	PaginationParams
}

type QuerySchedulerResp struct {
	BaseResponse
	Data []Scheduler `json:"data"`
}

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PaginationParams struct {
	Page int `query:"page" validate:"required,min=1"`
	Size int `query:"size" validate:"required,min=1,max=1000"`
}
