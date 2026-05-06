package router

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/hcd233/aris-api-tmpl/internal/handler"
)

func initHealthRouter(healthGroup huma.API, pingHandler handler.PingHandler) {
	huma.Register(healthGroup, huma.Operation{
		OperationID: "healthCheck",
		Method:      http.MethodGet,
		Path:        "/health",
		Summary:     "HealthCheck",
		Description: "Check the server health",
		Tags:        []string{"Health"},
	}, pingHandler.HandlePing)

	huma.Register(healthGroup, huma.Operation{
		OperationID: "sseHealthCheck",
		Method:      http.MethodGet,
		Path:        "/ssehealth",
		Summary:     "SSEHealthCheck",
		Description: "Check the server health",
		Tags:        []string{"Health"},
	}, pingHandler.HandleSSEPing)
}
