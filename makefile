build:
	go build -o bin/ChangeInspector
run:
	make build
	./bin/ChangeInspector $(path)
