:: Yeah, Windows build script for Linux software.
:: It's the most convenient for me as I use Windows
:: computers most of the time.

@echo off

del /Q build

echo Building for Linux (AMD64)
set GOOS=linux
set GOARCH=amd64
go build -o build/roagen-amd64

echo Building for Linux (ARM)
set GOOS=linux
set GOARCH=arm
go build -o build/roagen-arm

set GOOS=
set GOARCH=