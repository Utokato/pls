#!/bin/bash

# 编译 go 应用
echo "Building go app..."
go env -w GOOS=linux
go build -o bin/pls main.go
echo "Build go app success!"