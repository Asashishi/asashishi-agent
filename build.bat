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
echo go build -ldflags="-s -w" -trimpath -o asashishi-agent.exe
go build -ldflags="-s -w" -trimpath -o asashishi-agent.exe
echo del /f /q app.syso
del /f /q app.syso
@REM https://upx.github.io/
upx --best --lzma asashishi-agent.exe
move asashishi-agent.exe build\
copy README.md build\
copy config.json build\
robocopy web build\web /E
powershell -command "Compress-Archive -Path build -DestinationPath build\win-x86-amd64.zip"
echo Complete! --