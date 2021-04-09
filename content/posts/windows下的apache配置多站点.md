---
title: windows下的apache配置多站点
date: 2014-07-22T17:50:00+08:00
categories: [技术]
tags: [apache]
---

apache下配置多个网站：

在apache的配置文件中找到httpd.conf里找到

#Include conf/extra/httpd-vhosts.conf

去掉前面的#注释

在conf\extra\httpd-vhosts.conf文件里配置如下：
```
#

# Virtual Hosts

#

# If you want to maintain multiple domains/hostnames on your

# machine you can setup VirtualHost containers for them. Most configurations

# use only name-based virtual hosts so the server doesn't need to worry about

# IP addresses. This is indicated by the asterisks in the directives below.

#

# Please see the documentation at

# <URL:http://httpd.apache.org/docs/2.2/vhosts/>

# for further details before you try to setup virtual hosts.

#

# You may use the command line option '-S' to verify your virtual host

# configuration.



#

# Use name-based virtual hosting.

#

NameVirtualHost *:80



#

# VirtualHost example:

# Almost any Apache directive may go into a VirtualHost container.

# The first VirtualHost section is used for all requests that do not

# match a ServerName or ServerAlias in any <VirtualHost> block.

#

<VirtualHost *:80>

    ServerAdmin admin.chinablackhat.org

    DocumentRoot "e:/www"         (网站的位置)

    ServerName www.chinablackhat.org       (域名)

    ServerAlias www.chinablackhat.org

    ErrorLog "logs/dummy-host.x-error.log"

    CustomLog "logs/dummy-host.x-Access.log" common

</VirtualHost>



<VirtualHost *:80>

    ServerAdmin admin.chinablackhat.org

    DocumentRoot "e:/www/test"    (网站位置)

    ServerName test.chinablackhat.org      (网站域名)

    ErrorLog "logs/dummy-host2.x-error.log"

    CustomLog "logs/dummy-host2.x-Access.log" common

</VirtualHost>
```
