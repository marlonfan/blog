---
title: Consul快览
toc: true
date: 2018-06-02T21:26:00+08:00
categories: [技术]
tags: [golang]
---

## 介绍

因为团队目前正在考虑服务化部署,所以了解下相关的技术堆栈。 作为微服务架构里重中之重的``服务发现``和``集群一致性KV存储``当然是首先要了解的.

现在市面上有3种比较常见的服务发现服务:

- [zookeeper](https://zookeeper.apache.org/)
- [etcd](https://github.com/coreos/etcd)
- [consul](https://www.consul.io/)

在综合团队内部情况下,决定采用``consul``作为我们的服务发现中心和配置管理.

<!--more-->

## Consul快速入门

### 功能

1. 服务发现
1. Key->Value存储
1. 分布式锁
1. Watch变更
1. 属于颜控的一个好看的UI

### 环境搭建

建议使用docker快速部署:
```
###
# -d 后台运行
# --name 为容器起一个名字
# -p 映射docker容器内端口和宿主机器
# -server consul的命令,意思是以server端启动,ps: 还有个client端
# -ui 可以打开一个自带的UI可视化界面
# -bootstrap-expect 最小启动节点数量
###

docker run -d --name=consul -p 8500:8500/tcp consul agent -server -ui -bootstrap-expect=1 -client=0.0.0.0
```

### golang使用

这里提供一个golang的官方sdk: [sdk](https://github.com/hashicorp/consul/api)

官方已经有一个详细的 example, 这里就不多说了,但是研究Watch的时候发现有一点坑. Api模块里并没有Watch相关功能操作,所以看了下consul的源码,在consul/watch目录下看到了大概的实现

consul 的 watch有几种实现方式

1. 变更时执行shell脚本
1. 变更时通过http触发
1. 在源码中发现有一个hook接口,挂在handler上就会自动执行

这里因为避免部署上的操作,所以在应用启动中代码直接启动Watch. 下面是consul相关使用的实例代码:

```golang
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"

    "github.com/hashicorp/consul/api" // 操作consul的api模块
    "github.com/hashicorp/consul/watch" // api中未提及的watch模块
    "github.com/marlonfan/go-library/util" // 一个自定义的常用的package,在github可以找到
)

var consulClient *api.Client

var (
    serverID   = flag.String("server-id", "nomarl", "service's id")
    serverPort = flag.String("port", ":8080", "service's port")
)

func init() {
    flag.Parse()
    var err error
    consulClient, err = api.NewClient(api.DefaultConfig())
    util.CheckError(err)
}

func main() {
    go registorService()
    go waitToUnRegistService()
    go startWatch()

    // 健康检查
    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("check status.")
        fmt.Fprint(w, "status ok!")
    })

    fmt.Println("start listen...")
    err := http.ListenAndServe(*serverPort, nil)
    util.CheckError(err)
    select {}
}

// watch功能实现
func startWatch() {
    watchConfig := make(map[string]interface{})

    watchConfig["type"] = "service"
    watchConfig["service"] = "redis"
    watchConfig["handler_type"] = "script"
    watchPlan, err := watch.Parse(watchConfig)
    util.CheckError(err)
    watchPlan.Handler = func(lastIndex uint64, result interface{}) {
        services := result.([]*api.ServiceEntry)
        str, err := json.Marshal(services)
        util.CheckError(err)
        fmt.Println(string(str))
    }
    if err := watchPlan.Run("192.168.2.159:8500"); err != nil {
        log.Fatalf("start watch error, error message: %s", err.Error())
    }
}

// 注册服务实现
func registorService() {
    var err error
    client, err := api.NewClient(api.DefaultConfig())
    util.CheckError(err)

    service := &api.AgentServiceRegistration{
        ID:   *serverID,
        Name: "redis",
        Port: 12311,
        Check: &api.AgentServiceCheck{
            HTTP:     "http://192.168.2.159" + *serverPort + "/status",
            Interval: "1s",
            Timeout:  "1s",
        }}

    if err = client.Agent().ServiceRegister(service); err != nil {
        log.Fatalf("registor failed, error message: %s", err.Error())
    }
}


// 收到退出信号后的服务下线处理,如果不做,服务会显示健康检查失败,同时已经下线的服务还会在控制面板看到
func waitToUnRegistService() {
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, os.Kill)
    <-quit

    if consulClient == nil {
        return
    }

    if err := consulClient.Agent().ServiceDeregister(*serverID); err != nil {
        log.Fatal(err)
    }
    os.Exit(0)
}
```

全文完!
