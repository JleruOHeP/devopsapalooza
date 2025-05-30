https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html

go mod tidy

GOOS=linux GOARCH=amd64 go build -o bootstrap main.go