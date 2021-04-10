---
title: 自动切换 MacOS 主题颜色
date: 2021-04-10T02:04:01+08:00
tags: [vim]
---

最近白天在公司使用 ``Dark`` 模式的主题眼睛刺痛. 但是又比较喜欢暗色系主题. 想到 ``macos`` 下 部分软件是支持主题颜色跟随系统切换的, 那么白天就可以使用相对护眼的浅色系主题, 晚上再换回来~ 探索下如何在自己常用的几个工具中做对应设置:

- VSCode
- Goland
- Iterm2
- Neovim

目前初步达到想要的效果了~
<!--more-->

## VSCode

vscode本身支持了设置:

```json
{
    "workbench.colorTheme": "Atom One Dark",
    "workbench.preferredDarkColorTheme": "Atom One Dark",
    "workbench.preferredLightColorTheme": "Solarized Light"
}
```

设置对应模式下的模板即可

## Goland

Goland 需要安装一个插件 ``Auto Dark Mode``, 然后在 ``Plugin`` 的设置中设置对应模式主题即可.

![设置](https://tva1.sinaimg.cn/large/008eGmZEgy1gpe1svnnt5j31200u0wir.jpg)

## Iterm2

Iterm2可以借助 ``python`` 插件的能力, 来监听系统的 ``Dark`` 模式切换.

[设置教程](https://gist.github.com/FradSer/de1ca0989a9d615bd15dc6eaf712eb93)

## Neovim

Neovim 的设置依赖一个vim插件 [dark_notify](https://github.com/cormacrelf/dark-notify)

通过 vim 插件安装后, 在 ``init.vim`` 中增加以下设置

```vim
" dark-notify
:lua <<EOF
local dn = require('dark_notify')
dn.run({
    schemes = {
        dark = "gruvbox",
        light = "solarized8",
    },
})
EOF
```

就万事大吉啦~


## 最终效果

![播放效果](marlon-storage.oss-cn-shenzhen.aliyuncs.com/uPic/darkmode.gif)


## 参考

- [Automatic dark mode for terminal applications](https://arslan.io/2021/02/15/automatic-dark-mode-for-terminal-applications/)
- [iterm2_switch_automatic.md](https://gist.github.com/FradSer/de1ca0989a9d615bd15dc6eaf712eb93)
