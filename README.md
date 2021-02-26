# ScfWatchDog-gin-demo
ScfWatchDog 支持下部署gin服务

## 操作流程

- 去[watchdog](https://github.com/TencentCloudMiddleWare/ScfWatchDog)下载编译好的watchdog到`publish`目录下
- 编译gin项目，命令为`GOOS=linux GOARCH=amd64 go build`，然后将产出的文件，ScfWatchDog-gin-demo放到publish目录下
- 执行`sls deploy`命令
- 在新版编辑器给`publish`目录下的文件赋予可执行权限,`cd src`接着`chmod +x *`，接着点一下部署按钮

> 有问题提工单或者issue反馈即可
