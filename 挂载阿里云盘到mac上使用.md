# 挂载阿里云盘到 mac 上使用

[TOC]

## 安装 alist

参考文档 https://alist.nn.ci/zh/guide/install/docker.html 进行安装



## 获取阿里云盘相关参数

参考文档 https://alist.nn.ci/zh/guide/drivers/aliyundrive.html

2022.08.23亲测以下方案成功，PC端浏览器登录这里https://passport.aliyundrive.com/mini_login.htm?lang=zh_cn&appName=aliyun_drive&appEntrance=web&styleType=auto&bizParams=&notLoadSsoView=false&notKeepLogin=false&isMobile=true&hidePhoneCode=true&rnd=0.9186864872885723（移动端登录地址）然后login.do的response返回值json里bizExt值，base64解码之后最底下refresh_token可以实现下载和在线播放了



## 添加磁盘

https://alist.nn.ci/zh/guide/drivers/common.html



## Mac 挂载 DAV

https://support.apple.com/zh-cn/guide/mac-help/mchlp1546/13.0/mac/13.0

服务器地址：127.0.0.1:5244/dav

账号密码填写 alist 的

