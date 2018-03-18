# golang-grpc-gin

## 简介

本程序是一个 简陋的 漫画 demo ,没有业务就只是为了打通rpc与gin 数据库操作,用来学习grpc,gin 框架所写。

grpc 提供服务，gin调用服务从而提供 RESTfull Api, demo还非常简陋,后面继续完善。


## 特性

- 1.rpc : 采用 grpc 框架 实现 rpc 服务调用

- 2.api : 采用 gin 框架 实现RESTfull api

- 3.db  : 采用 mysql 实现分库,主从

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
├── rpc.go                应用入口
├── service               要注册的服务
│   ├── register.go
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

   可以理解成gin中的 middleware ,其他语言中的 hook。实现 异常的捕获，恢复防止程序意外退出。
   
- 2.trace

   利用golang 标准库 trace 来实现 服务调用的跟踪，如果想更详细的trace可以结合 zipkin。
   
- 3.mysql

   实现通过主从可以操作多个数据库。
   

### api部分

- 1.middleware

  jwt 验证，recover 异常的捕获与恢复，防止程序退出。
  
- 2.client

  rpc调用客户端，实现了连接池(有人说gRPC基于HTTP/2标准设计,已经足够快了,不需要连接池)
  
  
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

## 截图

![](http://static.golangtab.com/images/2018-03/WX20180318-170345@2x.png)
![](http://static.golangtab.com/images/2018-03/WX20180318-170432@2x.png)
![](http://static.golangtab.com/images/2018-03/WX20180318-170524@2x.png)
![](http://static.golangtab.com/images/2018-03/WX20180318-170611@2x.png)
   


