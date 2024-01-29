run local:
	nodemon --exec go run main.go --signal SIGTERM

dev:
	find -name "*.go" | entr -r go run .

build:
	docker build -t authv2 . -f Dockerfile.production && docker run --env-file .env -dp 8081:8080 authv2

serve:
	kubectl port-forward service/authv2 8080:8080

redis:
	kubectl port-forward service/redis 6379:6379