$env:GOOS = "linux"
$env:GOARCH = "amd64"
go mod download
go build -o code-executor
docker system prune -f
heroku container:login
heroku container:push web --recursive --app code-executor-golang
heroku container:release web --app code-executor-golang