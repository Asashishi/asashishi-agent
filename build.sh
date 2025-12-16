#!/bin/bash
set -e  # 出错立即退出

echo "-- Start"

echo "rm -rf build"
rm -rf build
mkdir -p build

echo "go mod tidy"
go mod tidy

# rsrc 工具在 Linux 下也能用，如果需要生成 .syso 文件可以保留
echo "rsrc -ico ./resources/app.ico -o app.syso"
rsrc -ico ./resources/app.ico -o app.syso

echo "go build -ldflags=\"-s -w\" -trimpath -o asashishi-agent"
go build -ldflags="-s -w" -trimpath -o asashishi-agent

echo "rm -f app.syso"
rm -f app.syso

# 压缩可执行文件
echo "upx --best --lzma asashishi-agent"
upx --best --lzma asashishi-agent

mv asashishi-agent build/
cp README.md build/
cp config.json build/

echo "zip -r build/linux-x86-amd64.zip build"
zip -r build/linux-x86-amd64.zip build

echo "Complete!"
