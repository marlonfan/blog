---
title: Express中间件加载
date: 2015-10-01T14:31:00+08:00
categories: [技术]
tags: [express]
---

**middleware**, express中几乎所有的东西都是通过中间件来完成的.大量采用第三方的中间件.但是它是怎样来工作的呢,让我们来一步一步揭开它.

从代码上看上去第一眼就是`app.use()`了.在使用`express-generator`构建好一个express应用后,在`app.js`文件下可以看到以下代码。(因代码较多,有节选).

<!--more-->

```javascript
var express = require('express');
var bodyParser = require('body-parser');
var routes = require('./routes/index');

var app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', routes);
```

从上面代码中很显然可以看出`app.use()`是用来加载这些中间件的.翻文档看,`use()`接受两个参数,第一个`path`为可选项,第二个就是中间件的函数了.

那我们就试着自己写一个东西调试一下能不能执行到了.

```javascript
app.use(express.static(path.join(__dirname, 'public')));

app.use(function(req, res, next) {
  console.log('11');
  next();
})

app.use(function(req, res, next) {
  console.log(22);
  next();
})

app.use('/', routes);
```

在这样的情况下,我们看到的响应还是正确的,并且在console里log出了`11`,`22`.

那我们试着把`console.log('11')`里面的next()取消掉呢,这时候命令行里console出了11后之后的全部不执行了,连路由都没有响应了.

ok,到这里的时候,我们搞明白了,`app.use()`里面的东西是依次执行的,那么我们可以把它想象成为一个tasks,在请求过来的时候依次通过`next()`调用之前注册好的中间件过滤一遍.使用 app.use() “定义的”中间件的顺序非常重要，它们将会顺序执行，use的先后顺序决定了中间件的优先级。 比如说通常 express.logger()是最先使用的一个组件，纪录每一个请求.

而在没有next()下,是无法继续传递下去的.

那么我们搞明白这个中间件是怎么加载的了,但是这个`next()`又是怎么实现的呢?我们试着一步一步自己实现一个简单的`next()`来看看.

```javascript
var http = require('http');

function express() {
	var fns = [];

	var expr = function(req, res) {
		var i = 0;
		function next() {
			var task = fns[i++];
			if ( ! task) {
				return;
			}
			task(req, res, next);
		}
		next();
	}

	expr.use = function(fn) {
		fns.push(fn);
	}

	return expr;
}

var app = express();

app.use(function(req, res, next) {
	console.log('req');
	next();
})

app.use(function(req, res, next) {
	console.log('done');
	next();
})

app.use(function(req, res, next) {
	res.end('over');
	next();
})

http.createServer(app).listen('3000');
```

这样就可以在`use`里添加你想添加的各种中间件啦~~~
