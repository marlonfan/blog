---
title: laravel excel包使用
date: 2014-12-09T18:45:00+08:00
categories: [技术]
tags: [laravel]
---

## laravel中excel插件的安装


在composer中引入laravel excel的包
```json
    "maatwebsite/excel": "1.*"
```

<!--more-->

在位于laravel/app/config下编辑app.php文件,在providers数组中添加以下值
```php
    'Maatwebsite\Excel\ExcelServiceProvider',
```
在同文件中找到aliasses数组添加以下值
```php
    'Excel' => 'Maatwebsite\Excel\Facades\Excel',
```
执行`composer install` 或 `composer update`命令.

## laravel excel的配置

在位于`laravel/vendor/maatwebsite/excel/src/config`下一些对于插件的一些配置项

`config.php` > 对excel和表全局的一些设置
`csv.php` > 对导入导出csv文件的设置
`export.pho` > 对打印出文件内容的一些设置
`import.php` > 对导入excel文件的设置

## laravel excel的简单使用

在之前的准备工作都做好了以后我们就可以用excel插件了

> **导出excel**

```php
<?php
$rows = array( array( 'id' => 1, 'name' => 'marlon' ) );

Excel::create($name, function($excel) use ($rows) {
    $excel->sheet('当天报名', function($sheet) use ($rows) {
        $sheet->fromArray($rows);
    });
})->store('xls', storage_path('excel'));
```

由于在php闭包中无法拿到闭包外的变量,所以需要用use把`$rows`引入进去,在最后的链式调用的store中所传的参数就是所需excel的格式和要保存到服务器的位置,此为绝对路径.

**在这个地方`store()`方法为存储,相对应的还可以使用`download()`方法来直接下载,至于export方法笔者还没搞懂用处是什么**

> **导入excel**

```php
<?php
Excel::load(Input::file('excel'), function($reader) {
    //获取excel的第几张表
    $reader = $reader->getSheet(0);
    //获取表中的数据
    $results = $reader->toArray();
    //在这里的时候$results 已经是excel中的数据了,可以再这里对他进行操作,入库或者其他....
});
```



**END**

在最后感谢excel插件的作者,他的官方网站[maatwebsite](http://www.maatwebsite.nl/laravel-excel/).
本文中若有错误欢迎指正.
