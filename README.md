# BlueSky

![](https://img.shields.io/badge/Go-1.16-green.svg)
![](https://img.shields.io/badge/AndroidQQ-8.4.8-red.svg)
![](https://img.shields.io/badge/GCC-11-green.svg)

Go版本号：1.16

协议类型：android

协议版本：8.4.8

[ --> 点我查看BlueSky的开发进度](https://github.com/zhangshikj/BlueSky/tree/main/plans)

[ --> 羁绊框架-WhiteCloud(WeChat) ](https://github.com/zhangshikj/WhiteCloud)

# 参与人员

 - 伏秋洛 （协议开发与指导）

# 部署环境

 - 安装Golang并设置基础环境变量（GOROOT,GOPATH）
 
 - 安装GCC编译器并设置环境变量

## 需要安装的Go模块

 - 导入项目设置好GOPATH参数后，请执行init.sh文件

### 常见问题

 1. 某Go模块未找到
    </br>方案：正确设置项目src所在路径为GOPATH
    
 2. 某Go模块无法下载
    </br>方案：设置GoMod代理为中国境内源

#### 所有平台使用BlueSky

安装Golang到系统配置GOPATH直接运行即可！

#### Android平台打包核心

如果需要编译Android平台的CShare库请安装**NDK**并配置好环境变量

#### Windows平台打包核心

编译项目为Dll库文件，使用易语言或者其它语言调用合成框架

#### Linux平台打包核心

未通过方案！

## 本协议仅提供学习与交流，禁止使用进行商业用途! 使用该项目请遵守**《MPL开源协议》**！
