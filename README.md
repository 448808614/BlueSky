# BlueSky

Go版本号：1.16

协议类型：android

协议版本：8.4.8

# 参与人员

 - 伏秋洛 （协议开发与指导）

# 部署环境

 - 安装Golang并设置基础环境变量（GOROOT,GOPATH）
 
 - 安装GCC编译器并设置环境变量

 - 导入项目后请按照文章安装Protobuf
   </br>**https://blog.csdn.net/JunChow520/article/details/115345622**

## 需要安装的Go模块

 - github.com/buger/jsonparser


### 常见问题

 1. 某Go模块未找到
    </br>方案：正确设置项目src所在路径为GOPATH
    
 2. 某Go模块无法下载
    </br>方案：设置GoMod代理为中国境内源

#### Android平台打包核心

如果需要编译Android平台的CShare库请安装**NDK**并配置好环境变量

#### Windows平台打包核心

暂时没有提供方案