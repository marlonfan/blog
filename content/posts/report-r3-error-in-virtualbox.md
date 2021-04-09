---
title: virtualbox报R3相关错误
date: 2015-06-10T14:40:00+08:00
categories: [技术]
tags: [vagrant]
---

#### virtualbox R3问题

```bash
Unable to load R3 module D:\app\virtualbox/VBoxDD.DLL (VBoxDD): GetLastError=1790 (VERR_UNRESOLVED_ERROR).
```

<!--more-->

这是我的错误日志,是使用vagrant的过程中虚拟机打不开找到的.百思不得其解,google之~

发现这里出现的问题主要是因为windows7的三个主题破解文件引起的,在网上有专门针对恢复的一种工具**UniversalThemePatcher**! 前车之鉴...

``嗯.还有个小提示,用ssh提交git不要密码的,但是如果用的是http协议还是会要的.``
