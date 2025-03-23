#!/bin/bash

# 服务器信息
SERVER_USER="root"
SERVER_HOST="42.192.105.150"
DEPLOY_DIR="/root/backend/fuck-the-world"
PROJECT_NAME="fuck-the-world"  # 可执行文件名称
LOG_FILE="./${PROJECT_NAME}.log.txt"
# 1. 编译项目
echo "🚀 正在编译 Go 项目..."
GOOS=linux GOARCH=amd64 go build -o $PROJECT_NAME

# 2. 上传可执行文件到服务器
echo "📤 正在上传可执行文件到服务器..."
scp ./$PROJECT_NAME $SERVER_USER@$SERVER_HOST:$DEPLOY_DIR/

# 3. 上传配置文件到服务器
echo "📤 正在上传配置文件到服务器..."
scp ./config.production.yaml $SERVER_USER@$SERVER_HOST:$DEPLOY_DIR/

# 4. 远程连接服务器，设置服务
echo "🔧 正在远程设置服务..."
ssh $SERVER_USER@$SERVER_HOST <<EOF
  # 进入部署目录
  cd $DEPLOY_DIR
  # 停止正在运行的服务
  pkill -f $PROJECT_NAME

  rm -r "$LOG_FILE"
  touch "$LOG_FILE"
  # 创建 systemd 服务文件（如果不存在）
  if [ ! -f /etc/systemd/system/$PROJECT_NAME.service ]; then
    echo "[Unit]
Description=Go Project Service
After=network.target

[Service]
Type=simple
ExecStart=$DEPLOY_DIR/$PROJECT_NAME
Restart=on-failure

[Install]
WantedBy=multi-user.target" | sudo tee /etc/systemd/system/$PROJECT_NAME.service
  fi

  # 设置环境变量并启动服务
  sudo chmod +x $PROJECT_NAME
  sudo systemctl start $PROJECT_NAME

  # 启动程序时设置环境变量
  echo "📈 启动服务..."
  export GIN_MODE=release
  export APP_ENV=production
  ./$PROJECT_NAME db migrate
  nohup ./"$PROJECT_NAME" server >> "$LOG_FILE" 2>&1 &

  echo "✅ 服务已启动！"
EOF

# 5. 清理本地二进制文件
rm -f ./$PROJECT_NAME
echo "🧹 本地清理完成！"

echo "🎉 Go 项目部署完成！"
