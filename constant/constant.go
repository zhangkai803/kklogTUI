package constant

import "kklogTUI/dto"

var (
	// 状态机
	StateChooseNul dto.State = 0
	StateChooseEnv dto.State = 1
	StateChooseDep dto.State = 2
	StateChoosePod dto.State = 3
	StateChooseNsp dto.State = 4
	StateDispLog   dto.State = 5

	// 空选项列表
	ChoicesEmpty   dto.Choices = dto.Choices{}
)

// 环境区分
var Envs = []*dto.Env{
	{Name:  "dev", Alias: "测试环境"},
	{Name:  "prod", Alias: "生产环境"},
}

var (
	// 测试环境命名空间
	DevNsSlice = []*dto.Namespace{
		{Env: Envs[0], Name: "sit",  Alias: "默认测试环境"},
		{Env: Envs[0], Name: "dev1", Alias: "开发测试环境1"},
		{Env: Envs[0], Name: "dev2", Alias: "开发测试环境2"},
		{Env: Envs[0], Name: "dev3", Alias: "开发测试环境3"},
		{Env: Envs[0], Name: "dev4", Alias: "开发测试环境4"},
		{Env: Envs[0], Name: "dev5", Alias: "开发测试环境5"},
	}

	// 生产命名空间
	prodNsCore = 		&dto.Namespace{Env: Envs[1], Name: "core", 		 Alias: "core"}
	prodNsIProd = 		&dto.Namespace{Env: Envs[1], Name: "iprod", 	 Alias: "iprod"}
	prodNsProductioin = &dto.Namespace{Env: Envs[1], Name: "production", Alias: "production"}
)

// Deployment 项目列表
var (
	depWTM = &dto.Deployment{ProdNamespace: prodNsIProd, 	Name: "wk-tag-manage", 		Alias: "标签管理系统"}
	depCMS = &dto.Deployment{ProdNamespace: prodNsIProd, 	Name: "wk-miniprogram-cms",	Alias: "抖快小程序"}
	depTIC = &dto.Deployment{ProdNamespace: prodNsCore, 	Name: "wk-tic", 			Alias: "视频直播项目"}
)

// Pod 服务类型
var (
	podTypeAPI 		dto.PodType = "api"
	podTypeScript	dto.PodType = "script"
)

// Pod 服务列表
var Pods = []*dto.Pod{
	{Type: podTypeAPI, 		Deployment: depWTM, Name: "wk-tag-manage", 							Alias: "API服务"},
	{Type: podTypeScript, 	Deployment: depWTM, Name: "wk-tag-manage-tag-record-subscriber", 	Alias: "打标签记录Kafka消费脚本"},
	{Type: podTypeAPI, 		Deployment: depCMS, Name: "wk-miniprogram-cms", 					Alias: "API服务"},
	{Type: podTypeScript, 	Deployment: depCMS, Name: "wk-miniprogram-cms-async-task", 			Alias: "异步任务消费脚本"},
}
