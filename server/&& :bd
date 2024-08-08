pid=$(lsof -t -i:8080)

if [ -n "$pid" ]; then 
	kill -TERM $pid
fi

go run main.go
