package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/hupe1980/go-huggingface"
)

type HuggingFaceClient struct {
	client *huggingface.InferenceClient
	model  string
}

func NewHuggingFaceClient(model string) *HuggingFaceClient {
	token := os.Getenv("HUGGINGFACE_TOKEN")
	client := huggingface.NewInferenceClient(token)

	return &HuggingFaceClient{
		client: client,
		model:  model,
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func (c *HuggingFaceClient) Explain(input string) (string, error) {
	req := &huggingface.TextGenerationRequest{
		Model:  c.model,
		Inputs: input,
		Options: huggingface.Options{
			UseCache:     boolPtr(true),
			WaitForModel: boolPtr(true),
		},
	}

	resp, err := c.client.TextGeneration(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("Hugging Face API error: %w", err)
	}

	if len(resp) == 0 {
		return "", fmt.Errorf("no output returned from model")
	}

	return resp[0].GeneratedText, nil
}
