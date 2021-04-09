---
title: git在sublime text中的使用
date: 2014-07-20T17:28:00+08:00
categories: [技术]
tags: [git]
---

Sublime作为一个轻量级的代码编辑器，凭借出色的界面和丰富的插件，逐渐成为编码者（不只是程序员，还包括诸如前端工程师、部分文字工作者）钟爱的编辑器。而git又是截至目前版本管理软件的领先者，因此，如何在sublime中使用git自然也就成了编码者面对的问题。本文将结合网络上的资料和自己摸索的经验，介绍如何在sublime text 2中实现git插件的版本管理以及如何在git@oschina远程托管git项目。

一、git的下载及安装

msysgit是Windows版的Git，从http://msysgit.github.io/下载，然后按默认选项安装即可。安装完成后，在开始菜单里找到“Git”->“Git Bash”，蹦出一个类似命令行窗口的东西，就说明Git安装成功！除了上述方式外，建议在项目文件夹上按鼠标右键，选择git Bash快捷方式使用，这样可以节省跳转目录的长传命令（尤其是目录比较深，文件夹名称较长的情况）。

然后将git的bin/cmd目录设置到path环境变量中。我的是“D:\IDE\Git\cmd”。

二、git的全局配置

Git的配置主要包括用户名，邮箱的设置，以及生成SSH密钥公钥等。

首先运行一下的命令设置git提交代码时你自己的用户信息。

git config --global user.name "username"
git config --global user.email "username@email.com"

在Sublime Text中使用的时候还需要设置 push.default参数。使用命令行窗口的时候没有问题，在Sublime Text中用push命令的时候就提示需要设置这个参数。

push.default参数主要是设置在执行push命令是的策略，主要的选项有以下几个：

nothing : Do not push anything
matching : Push all matching branches (default)
tracking : Push the current branch to whatever it is tracking
current : Push the current branch

这里我们手动设置成默认值：

git config --global push.default matching

生成SSH key

到开始菜单，找到“Git Bash”，运行之，并执行以下命令：

$ ssh-keygen -t rsa

程序会提示您输入密钥的文件名，比如输入oschina，按回车即可。然后会要求你输入一个密码，将来在使用密钥的时候需要提供这个密码。可以输入，也可以不输入直接回车（无论输入还是不输入，都会要求你确认一次）。
确认完毕后，程序将生成一对密钥存放在以下文件夹：

C:\Users\Administrator[这里替换成你的用户名]\.ssh

密钥分成两个文件，一个私钥（github_rsa）、一个公钥（github_rsa.pub）。
私钥保存在您的电脑上，公钥交项目负责人添加到服务器上。用户必须拥有与服务器公钥所配对的私钥，才能访问服务器上的代码库。

三、安装sublime的git插件

使用Package Control组件（推荐），打开install package控制台后，直接输入git就可以安装git插件。

这个时候Sublime Text只是安装了git插件，但还不能使用git命令，需要在修改Sublimt Text针对git的配置文件 “Git.sublime-settings”，这个文件一般在你的账户目录下，如：

C:\Users\Owen\Git.sublime-settings， 如果没有则创建这个文件。

在这个文件中加入如下内容：

"git_command": "D:/IDE/Git/cmd/git.exe"

指向的是你的git程序中的git.exe文件。

四、在git@osc建立项目仓库

访问http://git.oschina.net/，注册帐号，创建一个仓库（私有、公开都可），然后访问http://git.oschina.net/keys，添加前面生成的ssh公钥。

五、在本地创建git项目

在sublime中使用“Ctrl+Shift+p”打开命令窗口，输入“Git:init”来初始化git化境。 ST2会让你选择需要初始化的Git目录，选择到你的工程目录即可，之后就可以正常的使用git命令了。

使用git:add将所有文件添加到本地git项目中。

六、提交本地项目到git@osc远程仓库

在项目文件夹上按右键，选择git bash，执行如下命令，便可增加https远程仓库地址，这一步骤只需设置一次，之后可以直接提交代码。

git remote add origin http://yourname:password@git.oschina.net/name/project.git

使用git:commit，来提交更改。Sublime Text会自动跳出一个文本文件，你可以在文件的最上方输入这次更改的comments，然后直接关闭这个文件，就会出发commit操作。并且将你输入的comments作为-m的参数。 这个是非常方便的，比用命令行运行commit 用-m参数添加评论的方式要方便很多，而且可以随便修改。ctrl+w关闭该文件的同时，commit操作自动触发。

如果没有自动提交到远程，可使用git bash，运行git push手工提交。

至此，首次提交完成，以后当代码发生变化时，只需执行git:add git:commit git:push即可提交本地代码到远程仓库。
