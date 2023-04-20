#!/bin/bash

# 获取当前脚本运行的目录
rootDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
whitelist=("kafka_job" "my_todo")
serviceName="$1"

if [ "$serviceName" = "all" ]; then
  for service in "${whitelist[@]}"; do
    echo "Running build $service..."
    go build -o "${rootDir}/output/${service}" "$rootDir"/srv/"$service"/
  done
  exit 0
fi

found=0
for service in "${whitelist[@]}"; do
  if [ "$service" = "$serviceName" ]; then
    found=1
    break
  fi
done

# 如果 serviceName 不在白名单中，则输出错误信息并退出脚本
if [ $found -eq 0 ]; then
  echo "Invalid service name: $serviceName"
  echo "Allowed services: ${whitelist[*]}"
  exit 1
fi
echo "Running build $service..."
go build -o "${rootDir}/output/${serviceName}" "$rootDir"/srv/"$serviceName"/
exit 0


