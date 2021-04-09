---
title: ubuntu下运行laravel提示缺少mcrypt扩展
date: 2014-09-14T12:50:00+08:00
categories: [技术]
tags: [laravel]
---

```bash
sudo apt-get install php5-mcrypt

sudo mv -i /etc/php5/conf.d/mcrypt.ini /etc/php5/mods-available

sudo php5enmod mcrypt

service apache2 restart
```

<!--more-->

---------------------------------------------------------------------
上面的方法不知道为什么有时候不管用,用下面的这个
---------------------------------------------------------------------

cd /etc/php5/mods-available/conf.d/
sudo ln -s ../../conf.d/mcrypt.ini ./20-mcrypt.ini
sudo service apache2 restart
