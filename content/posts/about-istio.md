---
title: 关于istio
toc: true
date: 2018-06-09T21:26:00+08:00
categories: [技术]
tags: [架构,微服务,k8s]
---

## 前言

在微服务大行其道的当下,随着服务数量的增多,语言的多样化,服务治理慢慢变成了一个大部分公司都在头痛的难题. 在这样的背景下,**istio**应运而生.

<!--more-->

Istio是由Google/IBM/Lyft共同开发的新一代Service Mesh开源项目, 继承了服务的发现,负载,限流和熔断,链路追踪,网关等微服务必不可少的功能,把开发者从不断的维护开发环境,语言SDK中解救出来.

在研究过程中,笔者发现当前有一些热心的作者,已经翻译出了[istio中文手册](http://istio.doczh.cn/),但是由于操作API的变更,可能无法让对应配置生效. 但还是非常感谢他们的付出.

> 在当前时间(ps:2018-06-09), istio发布了``0.8``版本,  同时更换了新一版的alpha3操作API,老一版本的alpha2版本的api已经不建议使用,并且会在下一版本中删除.
> 官网中的提示:
> This task uses the new v1alpha3 traffic management API. The old API has been deprecated and will be removed in the next Istio release. If you need to use the old version, follow the docs here.

这里推荐英文基础不好的同学可以结合两个版本一起来看,比较容易理解.

## 初探

### k8s安装

首先,想要使用istio,我们需要先选择一个与他配合的基础设施. 由于istio对k8s的支持最为强大,所以这里我们选择了k8s为基础设施.

但是在安装k8s的时候发现这个集群搭起来还是有点费劲的. 想偷懒的同学可以去使用一个叫做[minikube](https://github.com/kubernetes/minikube)的开源软件,可以直接启动一个k8s集群.

由于在安装minikube的过程中发现由于一些GWF的不可描述原因,导致google的docker镜像库里面的镜像拉不回来, 有两个选择:

1. 装一个不可描述的软件,来越过GWF,下载原汁原味的minikube
1. 找到阿里的同学封装的一个国内定制版本: [传送门](https://yq.aliyun.com/articles/221687)


pppppppps: 写到这里发现其实需要的篇幅还是蛮长的,抽空来写,撸代码去....
