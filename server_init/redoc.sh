#!/bin/bash

# Параметры
REDOC_DIR="/opt/redoc"
SERVICE_FILE="/etc/systemd/system/redoc.service"
NGINX_CONF="/etc/nginx/sites-available/redoc"
DOMAIN_OR_IP="64.227.168.9"  # Замените на ваш домен или IP-адрес

# Шаг 1: Скачивание и настройка Redoc
sudo mkdir -p $REDOC_DIR
cd $REDOC_DIR
sudo wget https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js

# Создание HTML файла для Redoc
cat <<EOL | sudo tee $REDOC_DIR/index.html
<!DOCTYPE html>
<html>
  <head>
    <title>API Documentation</title>
    <meta charset="UTF-8">
  </head>
  <body>
    <redoc spec-url="/opt/openapi.yaml"></redoc>
    <script src="redoc.standalone.js"></script>
  </body>
</html>
EOL

# Шаг 2: Создание systemd сервиса
cat <<EOL | sudo tee $SERVICE_FILE
[Unit]
Description=Redoc Service
After=network.target

[Service]
Type=simple
User=deploy
WorkingDirectory=$REDOC_DIR
ExecStart=/usr/bin/python3 -m http.server 8444
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOL

# Перезагрузка конфигурации systemd и запуск сервиса
sudo systemctl daemon-reload
sudo systemctl enable redoc.service
sudo systemctl start redoc.service

# Шаг 3: Установка и настройка Nginx
sudo apt update
sudo apt install -y nginx

# Создание конфигурации для Nginx
cat <<EOL | sudo tee $NGINX_CONF
server {
    listen 80;
    server_name $DOMAIN_OR_IP;

    location / {
        proxy_pass http://localhost:8000;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
}
EOL

# Активирование конфигурации и перезагрузка Nginx
sudo ln -s $NGINX_CONF /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx

echo "Redoc service and Nginx setup complete. Access your API documentation at http://$DOMAIN_OR_IP"
