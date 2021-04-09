---
title: netbeans 在linux下抗锯齿
date: 2014-08-25T15:15:00+08:00
categories: [技术]
tags: [ide]
---

Ubuntu下NetBeans消除字体锯齿的方法
在netbeans.conf 文件的netbeans_default_options的最后添加 -J-Dswing.aatext=true -J-Dawt.useSystemAAFontSettings=lcd

<!--more-->
