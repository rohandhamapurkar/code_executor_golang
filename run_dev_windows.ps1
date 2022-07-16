# docker-compose -f docker-compose.dev.yml down
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go mod download
go build -o code-executor
docker-compose -f docker-compose.dev.yml build
docker system prune -f
docker-compose -f docker-compose.dev.yml up -d