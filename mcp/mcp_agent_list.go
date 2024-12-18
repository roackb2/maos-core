package mcp

import (
	"context"
	"fmt"
	"log/slog"

	"gitlab.com/navyx/ai/maos/maos-core/api"
	"gitlab.com/navyx/ai/maos/maos-core/dbaccess"
)

func GetMCPServers(ctx context.Context, logger *slog.Logger, ds dbaccess.DataSource) (api.GetMCPServersResponseObject, error) {
	logger.Info("GetMCPServers")

	agents, err := querier.ActorFindByMCPEnabled(ctx, ds)
	if err != nil {
		logger.Error("Cannot list agents", "error", err)
		return api.GetMCPServers500JSONResponse{
			N500JSONResponse: api.N500JSONResponse{Error: fmt.Sprintf("Cannot list agents: %v", err)},
		}, nil
	}

	response := api.GetMCPServers200JSONResponse{
		Data: make([]struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}, len(agents)),
	}

	// Populate the data using a loop
	for i, agent := range agents {
		response.Data[i].Id = agent.ID
		response.Data[i].Name = agent.Name
	}

	return response, nil
}
