.PHONY: gen
gen:
	 @cwgo server  --type HTTP --server_name ${svc} --module github.com/ocybers-dev/ocybers  --idl ./idl/${svc}.proto
	 @echo "generate ${svc} server success"

.PHONY: db
db:
	@cwgo  model --db_type mysql --dsn "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	@echo "generate db model success"

# 构建docker镜像命令
# 在使用该命令之前，确保已经登录到github的docker仓库
# docker login ghcr.io -u 用户名 -p <token>
.PHONY: build
build:
	@docker build -f ./deploy/dockerfile.backend -t ghcr.io/ocybers-dev/ocybers/ocybers:latest .
	@docker push ghcr.io/ocybers-dev/ocybers/ocybers:latest





