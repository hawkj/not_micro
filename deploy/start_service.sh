#!/bin/bash

# 获取当前脚本运行的目录
rootDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
outputDir=$rootDir/../output

# 读取白名单
whitelist=()
while IFS= read -r line; do
    whitelist+=("$line")
done < "$rootDir/service_white_list"

# 循环白名单并启动服务
for serviceName in "${whitelist[@]}"; do
  # 检查服务名是否以 "_job" 结尾，是则跳过本次循环
  if [[ "$serviceName" == *_job ]]; then
    echo "Skipping service ${serviceName} because it ends with '_job'"
    continue
  fi
  # 检查服务是否已经运行，如果是则先停止
  pid=$(pgrep -f "${outputDir}/${serviceName}")
  if [ -n "$pid" ]; then
    echo "Service ${serviceName} is already running (PID: ${pid}), stopping it now..."
    kill "$pid"
  fi

  # 启动服务
  if [ -f "${outputDir}/${serviceName}" ]; then
    echo "Starting service ${serviceName}..."
    "${outputDir}/${serviceName}" &
    echo "Starting service ${serviceName}... done"
  else
    echo "Error: Binary file for service ${serviceName} not found"
  fi
done
