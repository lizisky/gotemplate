# Golang Server Project Template


### build

```
make
```

### clean this project

```
make clean
```


### run server

```
copy config file to HOME directory
cp -r bin/.lizi_data/ ~

cd bin
./run -data_dir=<data/dir> 
如果没有指定 -data_dir 参数，则从默认的 $HOME/.lizi_data/ 目录下读取配置文件
```

### start mysql service in Docker

```
refer to: bin/boot_services_in_docker.sh
```


### build docker image

```
make docker
```

### run server inside docker
```
cd <path>/bin
docker-compose up -d

Note: please change local file path in file: <path>/bin/docker-compose.yml
```

### stop server inside docker
```
cd <path>/bin
docker-compose down
```

### More commands related with docker
```
1: export docker image:  docker save [image/name] -o [local/file/name]
2: load docker image:  docker load -i [local/file/name]
3: watch logs in docker container:  docker logs -f [container/id]
```
