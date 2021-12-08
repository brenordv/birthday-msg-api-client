echo off

echo Checking required build folders...
if not exist ".tmp" mkdir .tmp
if not exist ".tmp\osx" mkdir .tmp\osx

echo Setting OS and Architecture for OSX / amd64...
set GOOS=darwin
set GOARCH=amd64


echo Building: BDAYQUOTE...
go build -o .tmp\osx\bdayq.exe .\cmd\bdayquote

echo Creating archive for version %*...
7z a -tzip .\.dist\birthday-quote-getter--osx-amd64--%*.zip .\.tmp\osx\*.exe readme.md

echo All done!
