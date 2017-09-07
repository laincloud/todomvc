# Todo MVC

一个用于演示的 LAIN 应用，展示了 LAIN 的用法和特色。

> [在线演示](http://todomvc.demo.laincloud.com)

## 架构

![架构][architecture]

上图为 todomvc 的架构：当用户通过浏览器访问 `http://todomvc.${LAIN-domain}` 时，请求会先到达
LAIN 的 webrouter 组件；然后 webrouter 根据 URL 将请求根据 mountpoint 分发到不同的 proc，比如
`http://todomvc.${LAIN-domain}/` 会分发到 todomvc.web.web，`http://todomvc.${LAIN-domain}/api`
会分发到 todomvc.web.api；todomvc.web.web 是前端服务，返回 index.html 和 bundle.js
等静态文件，todomvc.web.api 是后端服务，实现了数据持久化，并支持对 todo 事项的增删查改。

## lain.yaml

lain.yaml 是 LAIN 应用的核心，它定义了 LAIN 应用的编译步骤和部署信息等。下面将以本应用的 [lain.yaml](lain.yaml)
为例，说明各个字段的用法。

### appname

[lain.yaml](lain.yaml) 第一行的 `appname` 定义了 LAIN 应用的名字，需要保证在集群内唯一。

### build

`build.script` 字段定义了 todomvc 的编译步骤：
- `cp -rf . $GOPATH/src/github.com/laincloud/todomvc` 将工程文件从 `/lain/app` 复制到 `$GOPATH/src/github.com/laincloud/todomvc`
- `cd $GOPATH/src/github.com/laincloud/todomvc/ && swagger generate server -A todomvc -f ./swagger.yml -T go-swagger-gen-templates/ -t gen`
  用 go-swagger 生成 api 的外围代码，包括序列化反序列化和参数校验等功能
- `go install github.com/laincloud/todomvc/gen/cmd/todomvc-server` 编译生成 todomvc-server
- `cd $GOPATH/src/github.com/laincloud/todomvc/frontend && yarn install && yarn run build`
  编译生成前端代码

`build.prepare.script` 字段用来缓存下载结果，可以提高编译速度。

> 如果想更新 prepare 镜像，请增大 build.prepare.version

### release

`release.dest_base` 定义了发布镜像。之所以不使用默认的 build 镜像有 2 个原因：
- 我们不需要在发布镜像里包含编译依赖，比如 golang 和 node 等 
- 我们需要 nginx 作为静态文件服务器

`release.copy` 字段可以帮我们把编译结果 `todomvc-server` 和 `frontend/dist`
从编译镜像复制到发布镜像。

### web.web

`web.web` 是前端 proc，用 nginx 作为静态文件服务器返回 `index.html` 和 `bundle.js` 等。
`mountpoint` 将 `web.web` proc 挂载到了 `http://todomvc.${LAIN-domain}/`，即可以通过
`http://todomvc.${LAIN-domain}/` 访问。

> 另外，我们把 nginx 的日志打到了标准输出，然后通过 s6-log 重定向到 `/lain/logs/default/current`。
> 这么配置是为了利用 s6-log 进行自动 rotate。

### web.api

`web.api` 是后端 proc，`cmd` 指定了启动命令。此外，还配置了：
- `healthcheck`：健康检查 API，用于监控报警和负载迁移
- `secret_files`：秘密文件，用于保存敏感信息
- `mountpoint`：挂载点，`/api` 与 `web.web` 的默认挂载点 `/` 不同，以保证 webrouter 对流量的正确分发

> web.api 的日志也由 s6-log 从标准输出重定向到了 `/lain/logs/default/current`，并且会自动 rotate。

### use_services

`use_services` 定义了需要使用的 service。这里我们使用了 LAIN 提供的开箱即用的
[mysql-service](https://github.com/laincloud/mysql-service)。

## 代码细节

### 前端

`todomvc` 使用了 yarn 管理依赖，使用 webpack 编译打包，使用 vue.js 进行数据绑定。代码实现参考了
[vue.js 的 todomvc 示例](https://github.com/vuejs/vue/tree/dev/examples/todomvc)。

### 后端

- 对 todo 事项的增产查改等业务逻辑包含在 [handler/](handler/) 文件夹内
- 序列化、反序列化和参数验证的接口代码使用 [go-swagger](https://github.com/go-swagger/go-swagger)
  根据 [swagger.yml](swagger.yml) 生成
- go-swagger 的自定义模板在 [go-swagger-gen-templates](go-swagger-gen-templates) 文件夹里，
  主要添加了对 [xorm](https://github.com/go-xorm/xorm) 和 [zap](https://github.com/uber-go/zap) 的支持

## 部署

```
lain reposit ${LAIN-cluster}
lain secret add ${LAIN-cluster} api /lain/app/prod.json -f prod.json  # prod.json 是秘密文件，并没有上传到 github，请参照 demo.json 的格式填写
lain build
lain tag ${LAIN-cluster}
lain push ${LAIN-cluster}
lain deploy ${LAIN-cluster}
```

[architecture]: images/architecture.gv.png
