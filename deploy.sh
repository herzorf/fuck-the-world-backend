#!/bin/bash

# 服务器信息
SERVER_USER="root"
SERVER_HOST="42.192.105.150"
DEPLOY_DIR="/root/backend/fuck-the-world"
PROJECT_NAME="fuck-the-world"  # 可执行文件名称
LOG_FILE="${DEPLOY_DIR}/${PROJECT_NAME}.log.txt"

# 1. 编译项目
echo "🚀 正在编译 Go 项目..."
GOOS=linux GOARCH=amd64 go build -o "$PROJECT_NAME"

# 2. 上传可执行文件和配置文件到服务器
echo "📤 正在上传可执行文件到服务器..."
ssh $SERVER_USER@$SERVER_HOST "rm -f $DEPLOY_DIR/$PROJECT_NAME"
scp "$PROJECT_NAME" "$SERVER_USER@$SERVER_HOST:$DEPLOY_DIR/"
scp ./config.production.yaml "$SERVER_USER@$SERVER_HOST:$DEPLOY_DIR/"

# 3. 远程连接服务器，进行部署
echo "🔧 正在远程部署服务..."
ssh "$SERVER_USER@$SERVER_HOST" <<EOF
 # 进入部署目录
  cd $DEPLOY_DIR

  # 停止正在运行的服务
  echo "🛑 终止旧进程..."
  pkill -f "$PROJECT_NAME" || true

  # 赋予可执行权限
  chmod +x "$PROJECT_NAME"

  # 清理日志
  echo "🧹 清理日志..."
  [ -f "$LOG_FILE" ] && > "$LOG_FILE" || touch "$LOG_FILE"

  # 设置环境变量
  export APP_ENV=production
  export GIN_MODE=release

  # 启动新进程
  echo "📈 启动服务..."
  nohup ./"$PROJECT_NAME" server >> "$LOG_FILE" 2>&1 &

  echo "✅ 服务已启动！"
EOF

# 5. 清理本地二进制文件
rm -f ./$PROJECT_NAME
echo "🧹 本地清理完成！"

echo "🎉 Go 项目部署完成！"