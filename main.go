package main

import (
	"errors"
	"os"

	"go.uber.org/zap"
)

const dirPermissions = 0755

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func readTextFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Add(a, b int) int {
	return a + b
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			panic("failed to sync logger: " + err.Error())
		}
	}()

	result, err := divide(10, 2)
	if err != nil {
		logger.Error("division error", zap.Error(err))
	}
	logger.Info("Result", zap.Int("value", result))

	if err := os.Mkdir("testdir", dirPermissions); err != nil {
		logger.Error("mkdir error", zap.Error(err))
	}

	content, err := readTextFile("nonexistent.txt")
	if err != nil {
		logger.Error("read error", zap.Error(err))
	}
	logger.Info("Content", zap.String("value", content))
}
