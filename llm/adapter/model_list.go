package adapter

import (
	"fmt"

	"gitlab.com/navyx/ai/maos/maos-core/llm"
)

const (
	PROVIDER_AZURE     = "Azure"
	PROVIDER_OPENAI    = "OpenAI"
	PROVIDER_ANTHROPIC = "Anthropic"
	PROVIDER_VOYAGE    = "VoyageAI"
)

var modelList = []llm.Model{
	{
		ID:       "5a265146-4e05-4cd7-a0a9-9adda7bf7a38-azure-gpt4o",
		Provider: PROVIDER_AZURE,
		Name:     "Azure gpt-4o",
	},
	{
		ID:       "bdf5c21b-ad28-4096-9bca-667927b5c742-azure-gpt4",
		Provider: PROVIDER_AZURE,
		Name:     "Azure gpt-4",
	},
	{
		ID:       "4b7b4d5c-7b6a-4d4e-8b0b-4b2b3c4d5e6f-openai-o1",
		Provider: PROVIDER_OPENAI,
		Name:     "OpenAI o1",
	},
	{
		ID:       "5c1a5c30-876a-48d4-9378-a2501fc6b92d-openai-o1-mini",
		Provider: PROVIDER_OPENAI,
		Name:     "OpenAI o1 mini",
	},
	{
		ID:       "be28aa22-c1c4-49e0-bc05-dddb6b7edb7b-openai-4o",
		Provider: PROVIDER_OPENAI,
		Name:     "OpenAI 4o",
	},
	{
		ID:       "7b8ffb04-4a8d-4e4a-b4b1-9ae90613d902-openai-4o-mini",
		Provider: PROVIDER_OPENAI,
		Name:     "OpenAI 4o mini",
	},
	{
		ID:       "3db6db92-a091-4944-9f7e-9d43e70218d3-anthropic-claude-3-opus-20240229",
		Provider: PROVIDER_ANTHROPIC,
		Name:     "Anthropic Claude 3 Opus 20240229",
	},
	{
		ID:       "93d07ee3-c9fb-4f0e-9fc1-df1a7af10b6c-anthropic-claude-3.5-sonnet-20240620",
		Provider: PROVIDER_ANTHROPIC,
		Name:     "Anthropic Claude 3.5 Sonnet 20240620",
	},
}

var modelMap = map[string]llm.Model{}

func init() {
	modelMap = make(map[string]llm.Model)
	for _, model := range modelList {
		modelMap[model.ID] = model
	}
}

func GetModelByID(id string) (llm.Model, bool) {
	model, ok := modelMap[id]
	return model, ok
}

func GetModelList() []llm.Model {
	return modelList
}

type AdapterCredentials struct {
	AOAIEndpoint    string
	AOAIAPIKey      string
	AnthropicAPIKey string
	OpenAIAPIKey    string
}

// CreateAdapter creates an adapter for the given model ID
// This is a variable so we can inject it for testing
var CreateAdapter = func(modelId string, credentials AdapterCredentials) (LLMAdapter, error) {
	model, ok := GetModelByID(modelId)
	if !ok {
		return nil, fmt.Errorf("model %s not found", modelId)
	}

	switch model.Provider {
	case PROVIDER_AZURE:
		return NewAzureAdapter(credentials.AOAIEndpoint, credentials.AOAIAPIKey, false)
	case PROVIDER_ANTHROPIC:
		return NewAnthropicAdapter(credentials.AnthropicAPIKey), nil
	case PROVIDER_OPENAI:
		return NewAzureAdapter("https://api.openai.com/v1", credentials.OpenAIAPIKey, true)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", model.Provider)
	}
}
