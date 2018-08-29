GOOS=linux GOARCH=386 go build -ldflags "-linkmode external -extldflags -static" -o AAH-linux-386
GOOS=linux GOARCH=amd64 go build -ldflags "-linkmode external -extldflags -static" -o AAH-linux-amd64
GOOS=linux GOARCH=arm go build -ldflags "-linkmode external -extldflags -static" -o AAH-linux-arm
GOOS=linux GOARCH=arm64 go build -ldflags "-linkmode external -extldflags -static" -o AAH-linux-arm64
