---
title: win10后homestead报错问题解决
date: 2015-10-24T23:32:00+08:00
categories: [技术]
tags: [vagrant]
---

升级win10以后突然发现homestead不能用了！惶恐~  解决之~

非常感谢@Hoang

<!--more-->

## 报错
```
There was an error while executing `VBoxManage`, a CLI used by Vagrant
for controlling VirtualBox. The command and stderr is shown below.

Command: ["hostonlyif", "create"]

Stderr: 0%...
Progress state: E_FAIL
VBoxManage.exe: error: Failed to create the host-only adapter
VBoxManage.exe: error: Code E_FAIL (0x80004005) - Unspecified error (extended info not available)
VBoxManage.exe: error: Context: "int __cdecl handleCreate(struct HandlerArg *,int,int *)" at line 66 of file VBoxManageHostonly.cpp
```

## 解决方案

### 2015年11月1日更新

出现这个问题的时候帮朋友解决掉了,但是发现其实网络还是没有通的,后来发现它是用一个类似来宾用户登录的. 然后给到一些权限连接的virtualbox里面的机子.

但是发现只适用于**部分用户**!

然后win7同样出现了这样的问题. 百思不得其解. 然后猜测是无权限创建与虚拟机外的主机的共享文件夹. 在这样的猜测下,试着,试着将homestead.yaml里的共享文件夹下面的nfs.png关掉以后,顺利开启.

那么这个nfs是个什么东西呢？下面的百度给出的解释:
```
NFS（Network File System）即网络文件系统，是FreeBSD支持的文件系统中的一种，它允许网络中的计算机之间通过TCP/IP网络共享资源。在NFS的应用中，本地NFS的客户端应用可以透明地读写位于远端NFS服务器上的文件，就像访问本地文件一样。
```

这样看来确实是vagrant的配置在virtualbox与外部主机进行通信的时候发生问题了呢...
然而对这块不大明白,还没找到个所以然.. 如有疏忽的地方,欢迎大牛指正~

---

### 原解决方法,创建于2015年10月24号
安装Virtualbox 5 和 Vagrant 1.7.4

注释homestead.rb第17行:
```
config.vm.network :private_network, ip: settings["ip"] ||= "192.168.10.10"
```

注释路径为:
``HashiCorp\Vagrant\embedded\gems\gems\vagrant1.7.4\plugins\providers\virtualbox\action.rb``第64行: **b.use ClearNetworkInterfaces**


若是第一次安装homestead 用``homestead init`` 若是因为更新win10而导致的问题，运行``homestead up``

然后运行``homestead halt``关闭box

打开virtualbox**管理->全局设定->网络->仅主机（host-only）网络->随便修改一个虚拟网卡，**Ip为：192.168.10.10

设置Homestead虚拟机,网络->网卡2->选中第6步设置的host-only的网卡保存

运行 ``homestead up0``

你的homestead 又复活了

## 参考链接
- laracast.com
- Hoang Stark's Blog
