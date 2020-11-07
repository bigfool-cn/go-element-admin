## 简介

[go-element-admin](https://go-vue.usenav.com) 是一个后台前端解决方案，它基于 [vue](https://github.com/vuejs/vue) 、 [element-ui](https://github.com/ElemeFE/element)、[golang](https://golang.org/)、[gorm](https://gorm.io/)、[casbin](https://casbin.org)实现。它使用了最新的前端技术栈，内置了动态路由，权限验证以及接口验证，并提供了简单的后端API服务，开箱即用，帮助你快速构建前后端分离动态路由和权限模型。

本项目通过[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)精简改造，移除了大部分组件，如需要其他组件，可前往搬运。

- [在线预览](https://go-vue.usenav.com)

- [github仓库](https://gitgub.com/bigfool-cn/go-element-admin)

- [gitee仓库](https://gitee.com/bigfool-cn/go-element-admin)


## 前序准备

你需要在本地安装 [node](http://nodejs.org/) 、[git](https://git-scm.com/) 和 [mysql](https://www.mysql.com/)。本项目技术栈基于 [ES2015+](http://es6.ruanyifeng.com/)、[vue](https://cn.vuejs.org/index.html)、[vuex](https://vuex.vuejs.org/zh-cn/)、[vue-router](https://router.vuejs.org/zh-cn/) 、[vue-cli](https://github.com/vuejs/vue-cli) 、[axios](https://github.com/axios/axios) 、 [element-ui](https://github.com/ElemeFE/element) 和 [golang](https://golang.org/)，提前了解和学习这些知识会对使用本项目有很大的帮助。


## 功能

```
- 登录 / 注销

- 系统管理
  - 用户管理
  - 菜单管理
  - 接口管理
  - 角色管理
  - 登录日志
- 系统工具
  - 接口文档
  - 表单生成器
```



## 服务端
>环境

* Golang版本：1.13.8

* 包管理器：[go module](https://github.com/golang/go/wiki/Modules)

* MySQL: 5.7
#### 导入数据
新建go-element-admin数据库，然后把go-element-admin/server/go-element-admin.sql文件导入到数据库中。

#### 启动服务
```bash
cd go-element-admin/server
go build main.go
main.exe
```
需要热更新可以使用[bee](https://github.com/beego/bee)启动。


#### 生产环境打包
```bash
SET GOARCH=amd64
SET GOOD=linux
go build main.go
```
生成的main文件放到服务器项目目录下，添加执行权限:chmod +x main后，执行./main即可启动服务，或者用[pm2](https://github.com/Unitech/pm2)启动。

#### 生产环境nginx部分配置
```bash
location /api/ {
        proxy_pass http://127.0.0.1:8001/;
        proxy_set_header Host $host:$server_port;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
}
location ^~/chat {
        ://github.com/Unitech/pm2roxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection upgrade;
        proxy_set_header Host $http_host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Real-Ip $remote_addr;
        proxy_set_header X-NginX-Proxy true;
        proxy_pass http://127.0.0.1:8001/chat;
        proxy_redirect off;

}
location /swagger/ {
        proxy_pass http://127.0.0.1:8001;
}

```
#### pm2启动服务配置
```bash
{
  "apps" : [{
    "name": "go-element-admin-api",
    "script": "./main",
    "cwd": "./",
    "error_file":"./log.txt",
    "out_file":"./log.txt",
    "merge_logs": true,
    "instances": 1,
    "autorestart": true,
    "watch": [
        "configs",
    ],
    "max_memory_restart": "1G",
  }]
}
```
接口地址：http://127.0.0.1:8001

## 前台

```bash
# 克隆项目
git clone https://github.com/bigfool-cn/go-element-admin.git

# 进入项目目录
cd go-element-admin

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动服务
npm run dev
```

浏览器访问 http://localhost:9527

## 发布

```bash
# 构建测试环境
npm run build:stage

# 构建生产环境
npm run build:prod
```

## 其它

```bash
# 预览发布环境效果
npm run preview

# 预览发布环境效果 + 静态资源分析
npm run preview -- --report

# 代码格式检查
npm run lint

# 代码格式检查并自动修复
npm run lint -- --fix
```

## 打赏
如果你觉得这个项目帮助到了你，你可以请作者吃一杯冰阔乐表示鼓励🍹:
![打赏](https://usenav-1256191445.cos.ap-guangzhou.myqcloud.com/pays.png)

## 支持浏览器版本

Modern browsers and Internet Explorer 10+.

| IE / Edge | Firefox | Chrome | Safari |
| --------- | --------- | --------- | --------- |
| IE10, IE11, Edge| last 2 versions| last 2 versions| last 2 versions

## License

[MIT](https://github.com/bigfool-cn/go-element-admin/blob/master/LICENSE)

Copyright (c) 2020-present bigfool-cn
