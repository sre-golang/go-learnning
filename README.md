# go-learnning

下载地址

Go官网下载地址：https://golang.org/dl/
Go官方镜像站（推荐）：https://golang.google.cn/dl/

一、mac 下go环境配置
GOROOT = /usr/local/go       //go安装路径
GOBIN = /usr/local/go/bin    //go命令路径
GOPATH = /Users/yangguoqiang/Documents/project/      //go项目路径

# vim .bash_profile
  export GOPATH=/Users/yangguoqiang/Documents/project/
  export GOBIN=/usr/local/go/bin
  export PATH=$PATH:$GOBIN
  export GOROOT=/usr/local/go/
  export PATH=${PATH}:/usr/local/mysql/bin
# source .bash_profile

二、go项目文件详解
# cd /Users/yangguoqiang/Documents/project/src
# ls -al
  drwxr-xr-x   7 sf  staff   224 10 25  2018 .
  drwx------+ 13 sf  staff   416  4 18 11:30 ..
  -rw-r--r--@  1 sf  staff  6148  9 15  2018 .DS_Store 
  drwxr-xr-x   7 sf  staff   224 10 25  2018 .idea                 //ide环境配置
  drwxr-xr-x   2 sf  staff    64  8 21  2018 bin                   //存放每个项目编译后的二进制文件
  drwxr-xr-x   3 sf  staff    96  8 21  2018 pkg                   //存放编译后的库文件         
  drwxr-xr-x  16 sf  staff   512  4 26 14:40 src                   //存放项目源码文件


