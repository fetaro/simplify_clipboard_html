rm -rf dist
mkdir -p dist/arm64
go build -o dist/arm64/slack_to_gdocs

mkdir -p dist/x86_64
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o slack_to_gdocs_x86
mv slack_to_gdocs_x86   dist/x86_64/slack_to_gdocs
