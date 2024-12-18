package routers

import (
	"go/go-backend-api/internal/routers/manage"
	"go/go-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin manage.AdminRouterGroup
}

var AppRouter = new(RouterGroup)
