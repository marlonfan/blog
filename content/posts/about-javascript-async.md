---
title: javascript异步理解
date: 2015-08-13T14:40:00+08:00
categories: [技术]
tags: [javascript]
---

废话不说,直接扔代码

<!--more-->

```javascript
interface TaskFunction {
	(done: (result?: any) => void): void;
}

function all(taskFns: TaskFunction[], callback: (results: any[]) => void): void {
	var results: string[] = [];

	var pending = taskFns.length;

	taskFns.forEach((taskFn, index) => {
		taskFn(result => {
			if (index in results) {
				return;
			}

			results[index] = result;

			if (--pending == 0) {
				callback(results);
			}
		});
	});
}

all([
	done => {
		done('hello');
	},
	done => {
		setTimeout(() => {
			done(', ');
		}, 100);
	},
	done => {
		setInterval(() => {
			done('world');
		}, 1000);
	},
	done => {
		done('!');
	}
], results => {
	console.log(results.join('')); // 输出 hello, world!
})
```
