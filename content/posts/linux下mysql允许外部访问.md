---
title: linux下mysql允许外部访问
date: 2014-12-15T18:37:00+08:00
categories: [技术]
tags: [mysql]
---

> 设置mysql 的配置文件

``/etc/mysql/my.cnf``找到``bind-address  =127.0.0.1``将其注释掉；//作用是使得不再只允许本地访问；
重启mysql：``/etc/init.d/mysql restart;``


> 进入mysql 数据库

``mysql -u  root -p``,``mysql>grant all privileges on  *.*  to root@'%'  identifies  by ' xxxx';``其中 第一个``*``表示数据库名；第二个``*``表示该数据库的表名；如果像上面那样 ``*.*``的话表示所有到数据库下到所有表都允许访问；``‘%’``：表示允许访问到mysql的ip地址；当然你也可以配置为具体到ip名称；``%``表示所有ip均可以访问；后面到``‘xxxx’``为root 用户的password；

<!--more-->
