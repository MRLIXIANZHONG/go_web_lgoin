## 环境

基于 gin +gorm+govendor版本控制 搭建的简单web后台，适合于初学者

## 知识点

go web服务的搭建

静态路由

路由分组

session使用

gorm数据库操作

简单CURD

登录

注册

ini配置

## 开始

## Go 语言环境安装

Go 语言支持以下系统：

- Linux
- FreeBSD
- Mac OS X（也称为 Darwin）
- Window

安装包下载地址为：<https://golang.org/dl/>。

各个系统对应的包名：

| 操作系统 | 包名                           |
| -------- | ------------------------------ |
| Windows  | go1.4.windows-amd64.msi        |
| Linux    | go1.4.linux-amd64.tar.gz       |
| Mac      | go1.4.darwin-amd64-osx10.8.pkg |
| FreeBSD  | go1.4.freebsd-amd64.tar.gz     |

## 配置

1.配置环境变量$GOROOT为你安装go的bin目录

默认为：C:\go\

2.配置环境变量$GOPATH为你的工作区目录（项目目录）一般$GOPATH目录下需要建src  pkg bin三个子文件夹

把你的下载下来的项目放在src文件夹下面

> 如果环境变量配置不起作用可以用go env查看，最好重启一下电脑

3.下载项目

```
git clone https://github.com/MRLIXIANZHONG/go_web_lgoin.git
```

4.govendor第三方包依赖管理

```
go get -u github.com/kardianos/govendor
```

5.govendor使用时，必须保证你的工程项目放在`GOPATH/src`目录下

6.切换项目目录，初始化包依赖

```
govendor sync 			#一句本地vendor/verdor.json文件指定的包机器版本信息从远程资源库拉取资源
```

7.完成

