# ===========================
# my-gin 项目统一 Makefile
# ===========================

# ---------- Gin App ----------
APP_NAME=my-gin
SRC_FILE=cmd/http/main.go
GO_CMD=go
PID_FILE=./log/app.pid

# ---------- Docker ----------
DC=docker\ compose    # 新版 Docker CLI
MYSQL_SERVICE=mysql

.PHONY: all build start stop restart ps logs docker-start docker-stop docker-restart docker-logs up

# -----------------------------
# Gin build / run
# -----------------------------

build:
	@echo "🏗️  Building $(SRC_FILE) ..."
	@$(GO_CMD) build -o $(APP_NAME) $(SRC_FILE)
	@echo "✅ Build finished"

start: build
	@echo "🚀 Starting $(APP_NAME) in background..."
	@bash -c 'nohup ./$(APP_NAME) > ./log/app.log 2>&1 & echo $$! > $(PID_FILE)'
	@echo "✅ Started with PID `cat $(PID_FILE)`"

stop:
	@echo "🛑 Stopping $(APP_NAME)..."
	@if [ -f $(PID_FILE) ]; then \
		PID=`cat $(PID_FILE)`; \
		if [ ! -z "$$PID" ] && ps -p $$PID > /dev/null; then \
			kill -9 $$PID; \
			echo "✅ Stopped PID $$PID"; \
		else \
			echo "❌ PID $$PID not running"; \
		fi; \
		rm -f $(PID_FILE); \
	else \
		echo "❌ PID file not found"; \
	fi

ps:
	@if [ -f $(PID_FILE) ]; then \
		PID=`cat $(PID_FILE)`; \
		if [ ! -z "$$PID" ] && ps -p $$PID > /dev/null; then \
			echo "✅ Running with PID $$PID"; \
		else \
			echo "❌ Not running"; \
		fi; \
	else \
		echo "❌ Not running"; \
	fi

restart: stop start

# -----------------------------
# Docker / MySQL 管理
# -----------------------------

docker-start:
	@echo "🚀 Starting MySQL container..."
	@docker compose up -d mysql

docker-stop:
	@echo "🛑 Stopping MySQL container..."
	@docker compose stop mysql
	@docker compose rm -f mysql

docker-restart: docker-stop docker-start

docker-logs:
	@docker compose logs -f mysql

db:
	@sudo docker exec -it my_gin_mysql mysql -uroot -pgin123

# -----------------------------
# 一条命令启动整个项目
# -----------------------------
up: docker-start gateway-start start
	@echo "✅ All services are up!"


# Gateway
GATEWAY_SRC=cmd/gateway/main.go
GATEWAY_BIN=gateway

gateway-start:
	@echo "🚀 Starting Gateway..."
	@go build -o $(GATEWAY_BIN) $(GATEWAY_SRC)
	@bash -c 'nohup ./$(GATEWAY_BIN) > ./log/gateway.log 2>&1 & echo $$! > ./log/gateway.pid'
	@echo "✅ Gateway started with PID `cat ./log/gateway.pid`"

gateway-stop:
	@echo "🛑 Stopping Gateway..."
	@if [ -f gateway.pid ]; then \
		PID=`cat gateway.pid`; \
		if ps -p $$PID > /dev/null; then \
			kill -9 $$PID; \
			echo "✅ Stopped Gateway PID $$PID"; \
			fi; \
			rm -f gateway.pid; \
	fi


