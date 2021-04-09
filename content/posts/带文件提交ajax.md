---
title: 带文件提交ajax
date: 2014-09-09T19:42:00+08:00
categories: [技术]
tags: [javascript]
---

ajax方式提交带文件上传的表单


一般的表单都是通过ajax方式提交，所以碰到带文件上传的表单就比较麻烦。基本原理就是在页面增加一个隐藏iframe，然后通过ajax提交除文件之外的表单数据，在表单数据提交成功之后的回调函数中，通过form单独提交文件，而这个提交文件的form的target就指向前述隐藏的iframe。

<!--more-->

html 代码

```
＜html＞
＜body＞

＜form action="upload.jsp" id="form1" name="form1" encType="multipart/form-data"  method="post" target="hidden_frame" ＞
    ＜input type="file" id="file" name="file" style="width:450"＞
    ＜INPUT type="submit" value="上传文件"＞＜span id="msg"＞＜/span＞
    ＜br＞
    ＜font color="red"＞支持JPG,JPEG,GIF,BMP,SWF,RMVB,RM,AVI文件的上传＜/font＞
    ＜iframe name='hidden_frame' id="hidden_frame" style='display:none'＞＜/iframe＞
＜/form＞

＜/body＞
＜/html＞

＜script type="text/javascript"＞
function callback(msg)
{
    document.getElementByIdx_x_x("file").outerHTML = document.getElementByIdx_x_x("file").outerHTML;
    document.getElementByIdx_x_x("msg").innerHTML = "＜font color=red＞"+msg+"＜/font＞";
}
＜/script＞
```

index.html 中主要要做的就是写一个 form和 iframe ，并把 form 的 target 设为 iframe 的名字，注意要把 iframe设为不可见，其他的都是正常的文件上传的写法，这样刷新的页面就是这个隐藏的 Iframe ，而在 index.html中是不会有页面刷新的，js的 callback方法是回调方法。用于清空文件上传框和显示后台信息，注意清空文件上传框的方法，和普通方法有点不一样。



--upload.jsp9Dhjsp 代码

```
＜%@ page language="java" contentType="text/html; charset=gb2312" %＞
＜%@ page import="com.jspsmart.upload.SmartUpload"%＞

＜%
    //新建一个SmartUpload对象
    SmartUpload su = new SmartUpload();

    //上传初始化
    su.initialize(pageContext);

    // 设定上传限制
    //1.限制每个上传文件的最大长度。
    su.setMaxFileSize(10000000);

    //2.限制总上传数据的长度。
    su.setTotalMaxFileSize(20000000);

    //3.设定允许上传的文件（通过扩展名限制）,仅允许doc,txt文件。
    su.setAllowedFilesList("doc,txt,jpg,rar,mid,waw,mp3,gif");

    boolean sign = true;

    //4.设定禁止上传的文件（通过扩展名限制）,禁止上传带有exe,bat,jsp,htm,html扩展名的文件和没有扩展名的文件。
    try {
        su.setDeniedFilesList("exe,bat,jsp,htm,html");

        //上传文件
        su.upload();
        //将上传文件保存到指定目录
        su.save("c://");

    } catch (Exception e) {
        e.printStackTrace();
        sign = false;
    }
    if(sign==true)
    {
        out.println("＜script＞parent.callback('upload file success')＜/script＞");
    }else
    {
        out.println("＜script＞parent.callback('upload file error')＜/script＞");
    }
%＞
upload.jsp 中只要注意最后输出的格式就可以了。其实原理就是输出一段js代码到 iframe 中，然后在iframe中来控制它的父页面。
```

OK，至此一个无刷新的页面上传组件就做好了，不要忘了在 WEB-INF/lib 下加上必须的 jspSmartUpload.jar 包。

    需要说明的是使用Iframe来上传，状态栏还是会有刷新的，因为iframe 中的页面刷新了嘛，但是外部页面，就是你所看到的页面是没有刷新的，所以也可以说是类似Ajax上传。
