GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

ifeq ($(GOOS),windows)
OUTPUT := audio-normalize_$(GOOS)_$(GOARCH).exe
else
OUTPUT := audio-normalize_$(GOOS)_$(GOARCH)
endif