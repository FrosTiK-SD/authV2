run local:
	nodemon --exec go run main.go --signal SIGTERM

serve:
	kubectl port-forward service/authv2 8080:8080

redis:
	kubectl port-forward service/redis 6379:6379