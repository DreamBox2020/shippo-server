# shippo-server

## build

dev

>go build -o out/main cmd/main.go

prod

>GOARCH=amd64 GOOS=linux go build -o out/main cmd/main.go

## run

dev

>GIN_MODE=release ./out/main

prod

>export GIN_MODE=release

>netstat -tunlp|grep ${port}

>kill ${id}

>sudo nohup ./main_${ver} > app_${ver}.log 2>&1 &
