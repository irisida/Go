run:
	go run -race main.go

build:
	go build -o goRaceConditions main.go

RUN_TARGETS := run-race
run-race: ARGS=-race
$(RUN_TARGETS): run
	go run $(ARGS) main.go