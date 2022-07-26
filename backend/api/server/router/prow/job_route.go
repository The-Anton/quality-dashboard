package prow

import (
	"github.com/redhat-appstudio/quality-studio/api/server/router"
	"github.com/redhat-appstudio/quality-studio/pkg/logger"
	"github.com/redhat-appstudio/quality-studio/pkg/storage"
	"go.uber.org/zap"
)

// systemRouter provides information about the server version.
type jobRouter struct {
	routes  []router.Route
	Storage storage.Storage
	Logger  zap.Logger
}

// NewRouter initializes a new system router
func NewRouter(storage storage.Storage) router.Router {
	r := &jobRouter{}
	logger, _ := logger.InitZap("info")
	r.Logger = *logger
	r.Storage = storage
	r.routes = []router.Route{
		router.NewPostRoute("/prow/results/post", r.createProwCIResults),
		router.NewGetRoute("/prow/results/get", r.getProwJobs),
		router.NewGetRoute("/prow/results/latest/get", r.getLatestSuitesExecution),
	}

	return r
}

// Routes returns all the API routes dedicated to the server system
func (s *jobRouter) Routes() []router.Route {
	return s.routes
}
