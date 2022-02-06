broker:
	docker-compose up

server:
	reflex -r '\.go' -s -- sh -c "go run server.go"

test:
	curl -X POST -H "Content-Type: application/json" \
    -d '{"title": "My Task", "description": "My Task Description"}' \
    http://127.0.0.1:3000/api/v1/sendtask

worker:
	go run worker/worker.go