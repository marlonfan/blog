---
title: 关于gitlab-ci的一些探索
date: 2016-09-07T15:00:00+08:00
categories: [技术]
tags: [自动化]
---

因为团队中想使用PHPCS来检查代码风格,所以考虑了两个方案.一个是基于``GIT HOOKS``的行为触发,一个是基于``GITLAB-CI``的一系列部署.考虑再三使用了``gitlab-ci``的解决方案.

<!--more-->

## 了解gitlab-ci的运行机制.

要想使用``gitlab-ci``,首先要明白它的组成. 这个东西有两个东西来支撑:

1. gitlab-ci server
2. gitlab-ci-runner

``gitlab-ci server``负责调度、触发Runner,以及获取返回结果. 而``gitlab-ci-runner``则是主要负责来跑自动化``CI``的一个宿主机子.

那么我们总结一下流程,其实是这个样子的:

![2016-07-22_13:48:39.jpg](https://static.marlon.life/2019-05-03-140450.jpg)

## Runner的概念

``runner``可以想象成一个守护进程,来守护你注册好的``service``和``gitlab-ci``绑定. 一个宿主机里的``runner``可以维护多个不同的``service``. 而``gitlab-ci``在收到需要build的请求时,会通知``service``执行你在``.gitlab-ci.yml``里面指定好的脚本,然后根据命令行的返回结果来决定这次``build``的成功还是失败.

**在了解完了这些概念以后我们就可以很轻松的搭建一个runner了.**

## GITLAB-CI搭配Runner的使用

#### 安装Runner

- 首先要找一台服务器来创建Runner，因为是要跟你的``gitlab``服务关联，所以服务器要可以访问你的``gitlab``服务。

- 安装gitlab-CI-multi-runner

- gitlab-ci-multi-runner是CI runner的运行程序，这里有多种安装方式（[见这里](https://gitlab.com/gitlab-org/gitlab-ci-multi-runner#installation/)），这里我们使用了第一种：在linux中安装软件。

```
友情提示: 在安装时由于网络问题不好安装,可以直接找到deb包下载后安装.
```

#### gitlab-ci-multi-runner命令介绍

- 执行``gitlab-ci-multi-runner help``可以看到所有命令的简介，在每个命令加``--help``可以看到更加具体的参数，比如``gitlab-ci-multi-runner start --help``,命令的执行顺序为：``register(注册runner)-->install(安装服务)-->start(运行服务)``.

#### GITLAB-CI配置

打开网址(比如你的gitlab服务地址是: ``http://gitlab.your.company/``，那gitlab CI的地址就是:``http://gitlab.your.company/ci``),找到想要配置CI的项目,点击后面的按钮``Add project to CI``
,给项目配置CI功能.进入CI项目，进入``Runners``标签页面，可以看到CI的url和token，这2个值是待会用命令注册runner时所需要的。

在``runner``的服务器上注册``runner``，执行命令``gitlab-ci-multi-runner register --user="你的用户"``,下面是执行命令后的交互信息。

```
友情提示:如果你用的是docker的执行方式，可以先把对应的docker的image下载下来，不然第一次执行CI会比较慢。
```

- 安装服务

执行命令``gitlab-ci-multi-runner install -n "服务名"``,后面的服务名是自己定义的名称，用来后面启动命名使用，与其相对的命令是``uninstall``.

启动服务，执行命令``gitlab-ci-multi-runner start -n "服务名"``,与其相类似的命令有``stop``
和``restart``.

验证``runner``，执行``gitlab-ci-multi-runner verify``,可以看到runner的运行情况.

```
root@cloudeye:~# gitlab-ci-multi-runner verify

aliveINFO[0000] 79bf814a Veryfing runner... is
aliveINFO[0000] 207a4b34 Veryfing runner... is
aliveINFO[0000] 20f849f7 Veryfing runner... is
aliveINFO[0000] 6e07e13a Veryfing runner... is
aliveINFO[0000] 23be6deb Veryfing runner... is
aliveINFO[0000] 4e348964 Veryfing runner... is
```

启动服务后，可以在刚才的``CI runners``页面看到已经有``runner``出现了。

#### gitlab-ci.yaml文件

配置好了``runner``，要让``CI``跑起来，还需要在项目根目录放一个``.gitlab-ci.yml``文件,在这个文件里面可以定制CI的任务,下面是简单的示例文件，更多的用法可以看[官方文档](http://doc.gitlab.com/ci/yaml/README.html)

```yaml
jobName:
  script:
    - ls
    - php command.php
    ...(比喻,使用直接删掉本行)
  only:
    develp
```

> 自己磕磕碰碰总结出来的,有不对的地方希望可以指正.
