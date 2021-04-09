---
title: linux 自定义变量以及方法
date: 2015-04-02T14:37:00+08:00
categories: [技术]
tags: [bash]
---

#### 在自己home目录里可以找到相对应的.bash*文件

- .bashrc 这里可以添加一些自定义的sh命令,可以很方便的执行
- .bash_history 这个文件主要负责记录用户所操作的命令 对于记录的量是可以设置的
- .bash_profile 这个文件是系统的每个用户设置环境信息,当用户第一次登录时,该文件被执行.
- .bash_logout 这个文件在用户退出是会被执行

在使用时,应该**尽可能的把东西加在.bashrc里面而不是.bash_profile里面**,这样可以减少错误的发生.

<!--more-->
