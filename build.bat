@echo off
echo -- Start
echo rmdir /s /q build
rmdir /s /q build
mkdir build
echo go mod tidy
go mod tidy
@REM go install github.com/akavel/rsrc@latest
echo rsrc -ico .\resources\app.ico -o app.syso
rsrc -ico .\resources\app.ico -o app.syso
echo go build -ldflags="-s -w" -trimpath -o AsashishiAgent.exe
go build -ldflags="-s -w" -trimpath -o AsashishiAgent.exe
echo del /f /q app.syso
del /f /q app.syso
@REM https://upx.github.io/
upx --best --lzma AsashishiAgent.exe
move AsashishiAgent.exe build\
copy README.md build\
copy config.json build\
powershell -command "npm run build"
powershell -command "Compress-Archive -Path build -DestinationPath build\\win-x86-amd64.zip"
echo Complete! --