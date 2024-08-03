# Impressive Linux commands cheat sheet cli
> Copy From https://github.com/chenjiandongx/pls
> And support offline mode

## How to build at Windows
```shell
go env -w GOOS=linux
go build -o pls
```

## How to use
1. cp pls to /usr/local/bin and chmod x to it
```shell
cp pls /usr/local/bin
chmod 777 pls
```

2. upgrade or offline
```shell
pls upgrade
# or
pls offline
```

3. search and show
```shell
pls search 压缩
pls show zip
```