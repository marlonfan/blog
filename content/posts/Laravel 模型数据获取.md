---
title: Laravel 模型数据获取
date: 2015-02-12T14:38:00+08:00
categories: [技术]
tags: [laravel]
---

#### laravel中总模型可以直接转数组

在今天立哥review代码的时候提出来的.

<!--more-->

原代码:
```php
$bmlist  = Appointment::all();
$baoming = [];

foreach ($bmlist as $bm) {
    $baoming[] = $bm->toArray();
}

return $baoming;
```

改进后的代码:

```php
$baoming = Appointment::all()->toArray();
return $baoming;
```

我只想深深说句 fuck!.
