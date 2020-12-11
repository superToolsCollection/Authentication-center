// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"authentication-center/internal/dao"
	"authentication-center/internal/server/grpc"
	"authentication-center/internal/server/http"
	"authentication-center/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
