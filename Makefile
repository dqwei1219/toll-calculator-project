gpu-coordinate:
	@go build -o bin/gpu-coordinate gpu-coordinate/main.go
	@./bin/gpu-coordinate

receiver:
	@go build -o bin/receiver data-receiver/main.go
	@./bin/receiver

.PHONY: gpu-coordinate receiver