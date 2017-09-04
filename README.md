# Todo MVC

一个用于演示的 LAIN 应用，展示了 LAIN 的用法和特色。

## lain.yaml

- appname
- build
- prepare
- release
- worker
- web
- secret file
- mysql-service
- 在不同集群使用不同的配置文件

## 架构

### 前端

yarn + webpack + vue.js

### 后端

go-swagger

[gen/](gen/) 文件夹里的文件由 go-swagger 生成，[handler/](handler/) 包含业务逻辑。

## 效果

## 部署与本地开发

### 部署

```
lain build
lain tag ${LAIN-cluster}
lain push ${LAIN-cluster}
lain deploy ${LAIN-cluster}
```

### 本地开发

```
go get -u dep
go get -u go-swagger
dep ensure
go-swagger generate server -A todomvc -T go-swagger-gen-templates -t gen
cd frontend && yarn install; cd ..
```

## TODO

- [ ] README
