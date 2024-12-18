package apitest

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/navyx/ai/maos/maos-core/api"
	"gitlab.com/navyx/ai/maos/maos-core/internal/fixture"
	"gitlab.com/navyx/ai/maos/maos-core/internal/testhelper"
)

func TestGetMCPServersEndpoint(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("Normal", func(t *testing.T) {
		server, ds, _ := SetupHttpTestWithDb(t, ctx)

		actor := fixture.InsertActor2(t, ctx, ds, "actor1", "agent", true, true, true, false, true) // Enabled, Deployable, MCP Enabled
		fixture.InsertActor2(t, ctx, ds, "actor2", "agent", true, true, true, false, false)         // Enabled, Deployable, MCP Disabled
		fixture.InsertActor2(t, ctx, ds, "actor3", "agent", false, true, true, false, true)         // Disabled, Deployable, MCP Enabled
		fixture.InsertActor2(t, ctx, ds, "actor4", "agent", true, false, true, false, true)         // Enabled, Not Deployable, MCP Enabled
		token := fixture.InsertToken(t, ctx, ds, "actor-token", actor.ID, []string{string(api.ReadMcp)})

		resp, resBody := GetHttp(t, server.URL+"/v1/mcp?trace_id=123", token.ID)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		resJson := testhelper.JsonToMap(t, resBody)
		require.Len(t, resJson["data"], 1)
		data := resJson["data"].([]any)[0].(map[string]any)
		require.Equal(t, strconv.FormatInt(actor.ID, 10), data["id"])
		require.Equal(t, actor.Name, data["name"])
	})

	t.Run("With wrong permission", func(t *testing.T) {
		server, ds, _ := SetupHttpTestWithDb(t, ctx)

		actor := fixture.InsertActor2(t, ctx, ds, "actor1", "agent", true, true, true, false, true) // Enabled, Deployable, MCP Enabled
		token := fixture.InsertToken(t, ctx, ds, "actor-token", actor.ID, []string{string(api.ReadInvocation)})

		resp, _ := GetHttp(t, server.URL+"/v1/mcp?trace_id=123", token.ID)
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("With wrong token", func(t *testing.T) {
		server, _, _ := SetupHttpTestWithDb(t, ctx)

		resp, _ := GetHttp(t, server.URL+"/v1/mcp?trace_id=123", "wrong-token")
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
