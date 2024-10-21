#!/bin/bash

# 编译前端应用
echo "Building front end app..."
cd web
rm -rf ../offline/dist
rm -rf dist
pnpm build
cd ..
mv web/dist/ offline/dist/
echo "Build front end app success!"

# 编译 go 应用
echo "Building go app..."
go env -w GOOS=linux
go build -o bin/pls main.go
echo "Build go app success!"