// Code generated by Kitex v0.7.2. DO NOT EDIT.

package projectservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	project "github.com/li1553770945/sheepim-route-service/kitex_gen/project"
)

func serviceInfo() *kitex.ServiceInfo {
	return projectServiceServiceInfo
}

var projectServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ProjectService"
	handlerType := (*project.ProjectService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetProjects":   kitex.NewMethodInfo(getProjectsHandler, newProjectServiceGetProjectsArgs, newProjectServiceGetProjectsResult, false),
		"GetProjectNum": kitex.NewMethodInfo(getProjectNumHandler, newProjectServiceGetProjectNumArgs, newProjectServiceGetProjectNumResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "project",
		"ServiceFilePath": `idl/project.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func getProjectsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*project.ProjectServiceGetProjectsArgs)
	realResult := result.(*project.ProjectServiceGetProjectsResult)
	success, err := handler.(project.ProjectService).GetProjects(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProjectServiceGetProjectsArgs() interface{} {
	return project.NewProjectServiceGetProjectsArgs()
}

func newProjectServiceGetProjectsResult() interface{} {
	return project.NewProjectServiceGetProjectsResult()
}

func getProjectNumHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*project.ProjectServiceGetProjectNumResult)
	success, err := handler.(project.ProjectService).GetProjectNum(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProjectServiceGetProjectNumArgs() interface{} {
	return project.NewProjectServiceGetProjectNumArgs()
}

func newProjectServiceGetProjectNumResult() interface{} {
	return project.NewProjectServiceGetProjectNumResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetProjects(ctx context.Context, req *project.ProjectsReq) (r *project.ProjectsResp, err error) {
	var _args project.ProjectServiceGetProjectsArgs
	_args.Req = req
	var _result project.ProjectServiceGetProjectsResult
	if err = p.c.Call(ctx, "GetProjects", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetProjectNum(ctx context.Context) (r *project.ProjectNumResp, err error) {
	var _args project.ProjectServiceGetProjectNumArgs
	var _result project.ProjectServiceGetProjectNumResult
	if err = p.c.Call(ctx, "GetProjectNum", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
