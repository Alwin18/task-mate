package rest

import (
	taskmate "task-mate"
	"task-mate/api/handler"
	"task-mate/gen/api/operations"
	"task-mate/gen/api/operations/health"
	"task-mate/gen/models"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *taskmate.Runtime, api *operations.TaskMateAPI, apiHandler handler.Handler) {
	// health
	api.HealthHealthHandler = health.HealthHandlerFunc(func(hp health.HealthParams) middleware.Responder {
		return health.NewHealthOK().WithPayload(&models.Success{
			Message: "Server up",
		})
	})
}
