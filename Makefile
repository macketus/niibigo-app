BINARY_NAME=NiibigoApp

build:
	@go mod vendor
	@echo "Building Niibigo..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Niibigo built!"

run: build
	@echo "Starting Niibigo..."
	@./tmp/${BINARY_NAME} &
	@echo "Niibigo started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Niibigo..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Niibigo!"

restart: stop start