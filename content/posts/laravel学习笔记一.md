---
title: laravel学习笔记一
date: 2014-07-18T18:38:00+08:00
categories: [技术]
tags: [laravel]
---

我在网上找啊找,关于laravel的资料少的可怜.
无奈之下就只能勉勉强强自己看着文档自学了.
想着自己学的也一点一点记录下来,对之后象我一样摸不着头脑的孩子们一点经验吧.

**laravel的安装**

1.下面我们首先下载好laravel的压缩包,解压到项目文件夹下,假设为wwwroot目录.
2.在下载composer安装好,安装过程中他需要让你选择php.exe的路径.一般在环境目录下的php文件下.
3.安装好后运行cmd,调出窗口.
4.在dos窗口下进入你的wwwroot目录,例子：
    原本打开cmd后应该是这样的：
        C:\Users\Administrator> _
    而我本机的wwwroot目录在D盘,那么我现在需要做的就是:
        C:\Users\Administrator>d:
        D:\>cd wwwroot
        D:\wwwroot>composer install
5.这时候他要自己下载安装东西,这个过程需要等待几分钟,等安装完以后你的laravel就搭建好了,尽情琢磨去吧~要注意,他的访问目录是public哦~之间访问站点是访问不到的,如果想直接访问域名应该还需要对apache进行一些设置,至此laravel安装完毕

