package mcp_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/navyx/ai/maos/maos-core/api"
	"gitlab.com/navyx/ai/maos/maos-core/internal/fixture"
	"gitlab.com/navyx/ai/maos/maos-core/internal/testhelper"
	"gitlab.com/navyx/ai/maos/maos-core/mcp"
)

func TestGetMCPServersWithDB(t *testing.T) {
	t.Parallel()
	logger := testhelper.Logger(t)
	ctx := context.Background()

	t.Run("Successful listing of MCP agents", func(t *testing.T) {
		t.Parallel()
		dbPool := testhelper.TestDB(ctx, t)
		defer dbPool.Close()

		// Setup agents (actors with agent role)
		agent1 := fixture.InsertActor2(t, ctx, dbPool, "agent1", "agent", true, true, true, false, true)
		fixture.InsertActor2(t, ctx, dbPool, "agent2", "agent", true, true, true, false, false)

		response, err := mcp.GetMCPServers(ctx, logger, dbPool)

		assert.NoError(t, err)
		require.IsType(t, api.GetMCPServers200JSONResponse{}, response)

		jsonResponse := response.(api.GetMCPServers200JSONResponse)
		assert.NotNil(t, jsonResponse.Data)
		assert.Len(t, jsonResponse.Data, 1)

		// Verify agent details
		for _, agent := range jsonResponse.Data {
			assert.NotEmpty(t, agent.Id)
			assert.NotEmpty(t, agent.Name)
			assert.Contains(t, []string{agent1.Name}, agent.Name)
		}
	})

	// t.Run("Empty list when no agents exist", func(t *testing.T) {
	// 	t.Parallel()
	// 	dbPool := testhelper.TestDB(ctx, t)
	// 	defer dbPool.Close()

	// 	request := api.McpListAgentsRequestObject{}

	// 	response, err := mcp.ListAgents(ctx, logger, dbPool, request)

	// 	assert.NoError(t, err)
	// 	require.IsType(t, api.McpListAgents200JSONResponse{}, response)

	// 	jsonResponse := response.(api.McpListAgents200JSONResponse)
	// 	assert.NotNil(t, jsonResponse.Data)
	// 	assert.Empty(t, jsonResponse.Data)
	// })

	// t.Run("Database error", func(t *testing.T) {
	// 	t.Parallel()
	// 	dbPool := testhelper.TestDB(ctx, t)

	// 	// Setup an agent
	// 	fixture.InsertActor(t, ctx, dbPool, "agent1")

	// 	dbPool.Close() // Simulate database error

	// 	request := api.McpListAgentsRequestObject{}

	// 	response, err := mcp.ListAgents(ctx, logger, dbPool, request)

	// 	assert.NoError(t, err)
	// 	assert.IsType(t, api.McpListAgents500JSONResponse{}, response)
	// 	errorResponse := response.(api.McpListAgents500JSONResponse)
	// 	assert.Contains(t, errorResponse.Error, "closed pool")
	// })

	// t.Run("Only list agent role actors", func(t *testing.T) {
	// 	t.Parallel()
	// 	dbPool := testhelper.TestDB(ctx, t)
	// 	defer dbPool.Close()

	// 	// Setup both agent and non-agent actors
	// 	fixture.InsertActor(t, ctx, dbPool, "agent1")
	// 	_, err := dbPool.Exec(ctx, "INSERT INTO actors (name, role) VALUES ($1, $2)", "portal1", "portal")
	// 	require.NoError(t, err)

	// 	request := api.McpListAgentsRequestObject{}

	// 	response, err := mcp.ListAgents(ctx, logger, dbPool, request)

	// 	assert.NoError(t, err)
	// 	require.IsType(t, api.McpListAgents200JSONResponse{}, response)

	// 	jsonResponse := response.(api.McpListAgents200JSONResponse)
	// 	assert.NotNil(t, jsonResponse.Data)
	// 	assert.Len(t, jsonResponse.Data, 1)
	// 	assert.Equal(t, "agent1", jsonResponse.Data[0].Name)
	// })
}
