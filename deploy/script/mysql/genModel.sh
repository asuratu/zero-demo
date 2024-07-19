#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

# 生成的表名
tables=$3
# 表生成的 model 目录
modelDir=../../../app/$2/rpc/model
# 自定义模板目录
homeDir=../../goctl/1.6.6
# 是否使用缓存
isCache=$4

# 数据库配置
host=127.0.0.1
port=3306
dbname=$1
username=root
passwd=

echo "开始创建库：$dbname 的表：$3"

if [ "$isCache" = "true" ]; then
    goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modelDir}" --style=goZero --home="${homeDir}" -cache=true
else
    goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modelDir}" --style=goZero --home="${homeDir}"
fi

echo "创建库：$dbname 的表：$3 完成"

