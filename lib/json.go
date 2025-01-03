package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseJson[T any](filePath string) (T, error) {
	var result T

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return result, fmt.Errorf("failed to read file %q: %w", filePath, err)
	}

	if err := json.Unmarshal(fileContent, &result); err != nil {
		return result, fmt.Errorf("failed to parse JSON from file %q: %w", filePath, err)
	}

	return result, nil
}
