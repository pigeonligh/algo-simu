init: 
	mkdir -p _output/bin

build:
	go build -o _output/bin/algo-simu ./main.go
