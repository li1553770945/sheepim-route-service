//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/li1553770945/sheepim-route-service/biz/infra/config"
	"github.com/li1553770945/sheepim-route-service/biz/infra/database"
	"github.com/li1553770945/sheepim-route-service/biz/infra/log"
	"github.com/li1553770945/sheepim-route-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-route-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-route-service/biz/internal/service"
)

func GetContainer(env string) *Container {
	panic(wire.Build(

		//infra
		config.GetConfig,
		log.InitLog,
		trace.InitTrace,

		//repo
		repo.NewRepository,
		database.NewDatabase,

		//service
		project.NewProjectService,

		NewContainer,
	))
}
