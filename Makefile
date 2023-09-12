gpu-coordinate:
	@go build -o bin/gpu-coordinate gpu-coordinate/main.go
	@./bin/gpu-coordinate

# need to build the whole folder to get the binary if
# there is main dependent on other files
receiver:
	@go build -o bin/receiver ./data-receiver 
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance-calculator
	@./bin/calculator

aggregator:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

.PHONY: gpu-coordinate receiver calculator aggregator