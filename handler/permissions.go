package handler

import (
	"context"

	"github.com/samber/lo"
	"gitlab.com/navyx/ai/maos/maos-core/api"
	"gitlab.com/navyx/ai/maos/maos-core/middleware"
)

var (
	// Permissions is a map of operation id to the permissions they require.
	Permissions = map[string][]api.Permission{
		"CreateInvocationAsync":          {api.CreateInvocation},
		"CreateInvocationSync":           {api.CreateInvocation},
		"GetNextInvocation":              {api.ReadInvocation},
		"ReturnInvocationResponse":       {api.ReadInvocation},
		"ListEmbeddingModels":            {api.ReadCompletion},
		"CreateCompletion":               {api.CreateCompletion},
		"GetMCPServers":                  {api.ReadMcp},
		"AdminListActors":                {api.Admin},
		"AdminGetActors":                 {api.Admin},
		"AdminCreateActor":               {api.Admin},
		"AdminUpdateActor":               {api.Admin},
		"AdminDeleteActor":               {api.Admin},
		"AdminGetActorConfig":            {api.Admin},
		"AdminListApiTokens":             {api.Admin},
		"AdminCreateApiToken":            {api.Admin},
		"AdminDeleteApiToken":            {api.Admin},
		"AdminUpdateConfig":              {api.Admin},
		"AdminListDeployments":           {api.Admin},
		"AdminGetDeployment":             {api.Admin},
		"AdminGetDeploymentResult":       {api.Admin},
		"AdminCreateDeployment":          {api.Admin},
		"AdminUpdateDeployment":          {api.Admin},
		"AdminDeleteDeployment":          {api.Admin},
		"AdminSubmitDeployment":          {api.Admin},
		"AdminPublishDeployment":         {api.Admin},
		"AdminRejectDeployment":          {api.Admin},
		"AdminRestartDeployment":         {api.Admin},
		"AdminListPodMetrics":            {api.Admin},
		"AdminListReferenceConfigSuites": {api.Admin},
		"AdminSyncReferenceConfigSuites": {api.Admin},
		"AdminGetSetting":                {api.Admin},
		"AdminUpdateSetting":             {api.Admin},
		"AdminListSecrets":               {api.Admin},
		"AdminUpdateSecret":              {api.Admin},
		"AdminDeleteSecret":              {api.Admin},
	}
)

func GetContextToken(ctx context.Context) *middleware.Token {
	tokenValue := ctx.Value(middleware.TokenContextKey)
	if tokenValue == nil {
		return nil
	}

	return tokenValue.(*middleware.Token)
}

func ValidatePermissions(ctx context.Context, operationID string) *middleware.Token {
	token := GetContextToken(ctx)
	if token == nil {
		return nil
	}

	requiredPermissions, ok := Permissions[operationID]
	if !ok {
		return nil
	}

	for _, requiredPermission := range requiredPermissions {
		if !lo.Contains(token.Permissions, requiredPermission) {
			return nil
		}
	}

	return token
}
