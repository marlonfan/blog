---
title: windows安装上mysql以后空密码的处理方法
date: 2014-08-15T17:38:00+08:00
categories: [技术]
tags: [mysql]
---

WAMP安装好后，mysql密码是为空的，那么要如何修改呢？其实很简单，通过几条指令就行了，下面我就一步步来操作。

首先，通过WAMP打开mysql控制台。

提示输入密码，因为现在是空，所以直接按回车。

然后输入“use mysql”，意思是使用mysql这个数据库，提示“Database changed”就行。

<!--more-->

然后输入要修改的密码的sql语句“update user set password=PASSWORD('hooray') where user='root';”，注意，sql语句结尾的分号不能少，提示什么什么OK就行了。

最后输入“flush privileges;”，不输入这个的话，修改密码的操作不会生效的。

然后输入“quit”退出。

另外，很多人说通过phpmyadmin直接修改mysql表里的密码就行，原理上应该是没错，但是我发现修改后mysql整个库都不见了，害的我重装了WAMP，最终还是通过命令行去修改的。

大家可以摸索下，其实操作并不困难，因为我发现同事电脑上的mysql都是空密码，这以后要是配服务器，也弄个空密码还不完蛋。
