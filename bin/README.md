# About run liziskyd

### docker-compose.yml

这个文件用于启动在docker中的liziskyd

### boot_services_in_docker.sh

这个文件用于启动在docker中的mysql。为了要让mysql的数据在本地持久化，请修改这个脚本文件中的本地目录

-v /work/data/data_mysql:/var/lib/mysql
