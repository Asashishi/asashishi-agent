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
echo "go build -ldflags=\"-s -w\" -trimpath -o asashishi-agent"
go build -ldflags="-s -w" -trimpath -o asashishi-agent
echo "rm -f app.syso"
rm -f app.syso
echo "upx --best --lzma asashishi-agent"
upx --best --lzma asashishi-agent
mv asashishi-agent build/
cp README.md build/
cp config.json build/
cp -r web build/
echo "tar -czf linux-x86-amd64.gz build"
tar -czf linux-x86-amd64.gz build
mv linux-x86-amd64.gz build/
echo "Complete!"