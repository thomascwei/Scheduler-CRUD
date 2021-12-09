# scheduler-CRUD

> 1. 從dbdiagram.io建立table與create語法
>
> 2. 在mysql建立table
>
> 3. 以sqlc建CRUD function與Transaction
>
> 4. 以gin開發REST API
> 
> 5. get_one_schedule & all_commands,此二api加入grpc
>

![image-20211128070316337](./Time_Schedule.png)


```shell
docker build -t thomaswei/schedule-crud . --no-cache
# 加入環境變量從container訪問host db
docker run -d --name schedule-crud -p 9567:9567 -e DB_HOST=host.docker.internal thomaswei/schedule-crud 
```

產生grpc命令, 在第一層目錄下執行
```shell
protoc -I=. --go_out=plugins=grpc:pkg/pb proto/*.proto
```
