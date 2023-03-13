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
var envDevelopemt 	= &dto.Env{Name:  "dev", Alias: "测试环境"}
var envProduction 	= &dto.Env{Name:  "prod", Alias: "生产环境"}
var Envs			= []*dto.Env{envDevelopemt, envProduction}

var (
	// 测试环境命名空间 组成 slice 供选择
	DevNsSlice = []*dto.Namespace{
		{Env: envDevelopemt, Name: "sit",  Alias: "默认测试环境"},
		{Env: envDevelopemt, Name: "dev1", Alias: "开发测试环境1"},
		{Env: envDevelopemt, Name: "dev2", Alias: "开发测试环境2"},
		{Env: envDevelopemt, Name: "dev3", Alias: "开发测试环境3"},
		{Env: envDevelopemt, Name: "dev4", Alias: "开发测试环境4"},
		{Env: envDevelopemt, Name: "dev5", Alias: "开发测试环境5"},
	}

	// 生产命名空间 单独声明绑定到 deployment 上 不需要选择
	prodNsCore			= &dto.Namespace{Env: envProduction, Name: "core", 		Alias: "core"}
	prodNsIProd			= &dto.Namespace{Env: envProduction, Name: "iprod", 	 	Alias: "iprod"}
	prodNsProductioin	= &dto.Namespace{Env: envProduction, Name: "production", 	Alias: "production"}
)

// Deployment 项目列表
var (
	depWTM = &dto.Deployment{ProdNamespace: prodNsIProd, 	Name: "wk-tag-manage", 		Alias: "标签管理系统"}
	depCMS = &dto.Deployment{ProdNamespace: prodNsIProd,	Name: "wk-miniprogram-cms",	Alias: "抖快小程序"}
	depTIC = &dto.Deployment{ProdNamespace: prodNsCore,		Name: "wk-tic", 			Alias: "视频直播项目"}
	depWCA = &dto.Deployment{ProdNamespace: prodNsCore,		Name: "wk-content-apis", 	Alias: "【旧】【外部】内容管理系统"}
	depWCB = &dto.Deployment{ProdNamespace: prodNsCore,		Name: "wk-ms-callback", 	Alias: "回调系统"}
	depWF  = &dto.Deployment{ProdNamespace: prodNsIProd,	Name: "wk-form",			Alias: "表单系统"}
	depWR  = &dto.Deployment{ProdNamespace: prodNsIProd,	Name: "wk-risk",			Alias: "风控系统"}
)

// Pod 服务类型
var (
	podTypeAPI 		dto.PodType = "api"
	podTypeScript	dto.PodType = "script"
)

// Pod 服务列表
var Pods = []*dto.Pod{
	{Type: podTypeAPI, 		Deployment: depWTM, Name: depWTM.Name, 								Alias: "API服务"},
	{Type: podTypeScript, 	Deployment: depWTM, Name: "wk-tag-manage-tag-record-subscriber",	Alias: "打标签记录Kafka消费脚本"},
	{Type: podTypeAPI, 		Deployment: depCMS, Name: depCMS.Name, 								Alias: "API服务"},
	{Type: podTypeScript, 	Deployment: depCMS, Name: "wk-miniprogram-cms-async-task", 			Alias: "异步任务消费脚本"},
	{Type: podTypeAPI, 		Deployment: depWCA, Name: depWCA.Name, 								Alias: "API服务"},
	{Type: podTypeAPI,		Deployment: depWCB, Name: depWCB.Name, 								Alias: "API服务"},
	{Type: podTypeAPI,		Deployment: depWF,	Name: depWF.Name, 								Alias: "API服务"},
	{Type: podTypeAPI,		Deployment: depWR,	Name: depWR.Name, 								Alias: "API服务"},
}
