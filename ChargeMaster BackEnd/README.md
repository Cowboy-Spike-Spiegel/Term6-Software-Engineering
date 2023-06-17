

BUPT Software Engineering
===========================

##### BUPT软件工程大作业——充电桩系统



###### part: 后端

###### contributors: 2020211376-马天成



#### 1. 环境依赖

1. windows开发环境：

	1. go 1.19.3
	2. mysql Ver 8.0.28 for Win64 on x86_64

2. Linux部署环境：

	1. centos 7.9
	2. mysql80-community-release-el7-7.noarch



#### 2. 开发工具

GoLand 2022.2.4



#### 3. 后端目录结构描述

├── README.md
|
├── main.go                 			        // 源程序入口
├── src
│   ├── consts              			        // 常量
│   │   └── consts.go
│   ├── model              				    // 接口模型
│   │   ├── object.go							  // request model
│   │   └── response.go						// response model
│   ├── middleware    				    // 中间件
│   │   ├── cros.go								// 跨域中间件
│   │   └── auth.go								// 鉴权中间件
│   ├── router              					// 路由层
│   ├── controller         				   // 控制层
│   ├── service             					// 服务层
│   │   ├── crud_station.go					// 调用system.operate
│   │   ├── crud_car.go
│   │   ├── crud_order.go
│   │   └── crud_user.go
│   ├── system            				     // 业务系统层
│   │   ├── operate.go							// system对service的开放接口
│   │   ├── system.go
│   │   ├── run.go
│   │   └── dispatch.go
│   ├── dao                					 // 数据持久化层
│   ├── output             				        // 输出管理
│   │   ├── print.go
│   │   └── log.go
│   └── utils                				     // 小工具
│        ├── jwt.go									// 鉴权
│        ├── time.go
│        ├── type.go
│        └── yaml.go
├── go.mod
├── go.sum
|
├── log							    	 	// 日志输出
├── bin							    	 	// windows本地开发打包文件
│   ├── release.exe        					   // 发行版
│   └── test.exe              					   // 测试版（相比发行版增加调试日志）
|
├── ChargeBackEnd-v1 			// 打包后的linux可执行二进制文件
|
└──  configs								// 配置文件&设计模型
  ├── project_chargemaster.sql
  └── design.jpg



#### 4. 后端数据库说明

在项目configs目录下，为project_chargemaster.sql文件（包含数据和结构）。



#### 5. Linux部署说明

1. 数据库部署

	1. 上传sql文件到linux（假定上传到/home/charge目录）

		```sql
		/home/charge/project_chargemaster.sql
		```

	2. 登录mysql，新建数据库

		```sql
		CREATE DATABASE project_chargemaster;
		```

	3. 执行sql文件

		```sql
		source /home/charge/project_chargemaster.sql
		```

	

2. windows环境后端项目打包

	1. 进入项目根目录

	2. 设定目标平台

		```shell
		set GOOS=linux
		```

	3. 设定目标架构

		```shell
		set GOARCH=amd64
		```

	4. 打包

		```shell
		go build -o "ChargeBackEnd-v1"
		```

		

3. linux环境后端部署

	1. 上传二进制文件到linux（假定目录为/home/charge）

	2. 设定执行权限

		```shell
		chmod +x /home/charge/ChargeBackEnd-v1
		```

	3. 后台执行并输出打印信息到日志（重定向stderr并入stdout并输出内容到文件charge.log）

		```shell
		nohup ./home/charge/ChargeBackEnd-v1 release > charge.log 2>&1 &
		```

	

4. 服务器端口配置

	1. 防火墙端口开放
	2. 服务器安全组端口开放
	3. 适用http协议
