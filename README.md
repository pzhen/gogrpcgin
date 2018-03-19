# golang-grpc-gin

## 简介

本程序是一个 简陋的 demo ,没有业务就只是为了打通rpc与gin 数据库操作,用来学习grpc,gin 框架所写。

grpc 提供服务，gin调用服务从而提供 RESTfull Api, demo还非常简陋,后面继续完善。


## 特性

- 1.rpc : 采用 grpc 框架 实现 rpc 服务调用

- 2.api : 采用 gin 框架 实现RESTfull api

- 3.db  : 采用 xorm 操作 mysql 实现分库,主从

- 4.package : 采用 glide  实现包管理


## 目录结构

```
.
├── README.md
├── api                   Api 应用
│   ├── api.go            应用入口
│   ├── caches
│   ├── client
│   ├── controllers
│   ├── logs
│   ├── models
│   └── routers
├── conf                  配置
│   ├── conf.go
│   ├── conf_dev.toml
│   └── conf_pro.toml
├── core                  module
│   └── model.go
├── glide.lock
├── glide.yaml
├── pb                    protobuf 描述文件目录
│   ├── srvcomic.pb.go
│   ├── srvcomic.proto
│   └── srvnews.proto
├── rpc.go                RPC应用入口
├── service               服务
│   ├── register.go       服务注册
│   ├── srv-comic         
│   └── srv-news
├── utils                 工具函数
│   ├── utils_log.go
│   └── utils_string.go
└── vendor                第三方包
```

## 封装的内容

### grpc部分 

- 1.拦截器 

   服务端: 可以理解成gin中的 middleware ,其他语言中的 hook。实现 异常的捕获，恢复防止程序意外退出。以及函数栈执行时间。
   
   客户端: 拦截器 执行时间, 连接池(虽然http2已经足够快保持长链接,这里也手动实现下)
   
- 2.trace

   利用golang 标准库 trace 来实现 服务调用的跟踪，如果想更详细的trace可以结合 zipkin。
   
- 3.mysql

   采用了xorm 实现通过主从可以操作多个数据库。
   
- 4.配置加载

   采用 toml 来对配置的解码,针对 开发环境 生产环境 加载不同配置
   

### api部分

- 1.middleware

  jwt 验证，recover 异常的捕获与恢复，防止程序退出。
  
  MVC 目录划分
  
  路由分组规划(私有以及公开路由)
  
  
## 数据表

只是简单的几个字段。

```

DROP TABLE IF EXISTS `comic_info`;

CREATE TABLE `comic_info` (
  `comic_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `comic_name` varchar(32) NOT NULL DEFAULT '' COMMENT '名字',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1:正常2:关闭',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '时间',
  PRIMARY KEY (`comic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='漫画';

LOCK TABLES `comic_info` WRITE;
/*!40000 ALTER TABLE `comic_info` DISABLE KEYS */;

INSERT INTO `comic_info` (`comic_id`, `comic_name`, `status`, `create_time`)
VALUES
	(1,'啊 啊啊啊',1,12121212),
	(2,'分水电费水电费',1,111110),
	(3,'吞吞吐吐',2,9999);

```

## 使用

前提是 对golang有所熟悉, 比如环境搭,工具链的使用。

1.<code>git clone https://github.com/pzhen/golang-grpc-gin.git  ./gogrpcgin </code>执行 glide up 将类包更新下来(需要vpn,你懂得)

2.将数据表导入mysql

3.然后修改配置文件中的数据库信息。

4.分别运行 go run rpc.go |  go run api.go

5.访问 http://localhost:8080/rpc/comic/test 可以看到数据,命令行查看调用栈时间

可以手动抛出 panic ,来验证拦截器。




   


