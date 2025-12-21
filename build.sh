#!/bin/bash
set -e
echo "-- Start"
echo "rm -rf build"
rm -rf build
mkdir -p build
echo "go mod tidy"
go mod tidy
echo "rsrc -ico ./resources/app.ico -o app.syso"
rsrc -ico ./resources/app.ico -o app.syso
echo "go build -ldflags=\"-s -w\" -trimpath -o AsashishiAgent"
go build -ldflags="-s -w" -trimpath -o AsashishiAgent
echo "rm -f app.syso"
rm -f app.syso
echo "upx --best --lzma AsashishiAgent"
upx --best --lzma AsashishiAgent
mv AsashishiAgent build/
cp README.md build/
cp config.json build/
npm run build &
BUILD_PID=$!
wait $BUILD_PID
echo "tar -czf linux-x86-amd64.gz build"
tar -czf linux-x86-amd64.gz build
mv linux-x86-amd64.gz build/
echo "Complete!"