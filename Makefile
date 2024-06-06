build-image:
	docker build -t ltwo/finance .

run-app:
	docker compose -f .devops/app.yaml up -d

unit-test:
	go test ./...