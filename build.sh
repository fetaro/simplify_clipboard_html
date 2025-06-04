rm -rf dist
mkdir -p dist/arm64
go build -o dist/arm64/simplify_clipboard_html

mkdir -p dist/x86_64
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o simplify_clipboard_html_x86
mv simplify_clipboard_html_x86   dist/x86_64/simplify_clipboard_html
