echo off

echo Checking required build folders...
if not exist ".tmp" mkdir .tmp
if not exist ".tmp\linux" mkdir .tmp\linux

echo Setting OS and Architecture for linux / amd64...
set GOOS=linux
set GOARCH=amd64


echo Building: BDAYQUOTE...
go build -o .tmp\linux\bdayq.exe .\cmd\bdayquote

echo Creating archive for version %*...
7z a -tzip .\.dist\birthday-quote-getter--linux-amd64--%*.zip .\.tmp\linux\*.exe readme.md

echo All done!
