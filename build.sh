CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' ./src/main.go
scp ./main 185.154.53.78:/opt/jobhelper
