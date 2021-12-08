echo off

echo Checking required build folders...
if not exist ".tmp" mkdir .tmp
if not exist ".tmp\win64" mkdir .tmp\win64

echo Setting OS and Architecture for windows / amd64...
set GOOS=windows
set GOARCH=amd64


echo Building: BDAYQUOTE...
go build -o .tmp\win64\bdayq.exe .\cmd\bdayquote

echo Creating archive for version %*...
7z a -tzip .\.dist\birthday-quote-getter--win64-amd64--%*.zip .\.tmp\win64\*.exe readme.md

echo All done!
