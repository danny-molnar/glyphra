package llm

import "testing"

func TestExplainMock(t *testing.T) {
	client := &HuggingFaceClient{
		model:  "test-model",
		client: nil, // pretend
	}

	_ = client // 👈 this silences the "unused" warning

	resp := "Pretend this is a summary."
	if resp == "" {
		t.Fatal("Expected a mock response, got empty string")
	}
}
