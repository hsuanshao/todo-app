package router

import (
	"github.com/blackhorseya/todo-app/internal/app/apis"
	"github.com/blackhorseya/todo-app/internal/app/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)

// ProviderSet is a router provider set
var ProviderSet = wire.NewSet(NewRouter)

// IRouter define a interface for router
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Router is a route management
type Router struct {
	C         *config.Config
	HealthAPI *apis.Health
}

// NewRouter constructor of Router
func NewRouter(c *config.Config, healthAPI *apis.Health) IRouter {
	return &Router{C: c, HealthAPI: healthAPI}
}

// Register register route to Gin engine
func (r *Router) Register(app *gin.Engine) error {
	r.RegisterAPI(app)
	return nil
}

// Prefixes is prefix list of route
func (r *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
