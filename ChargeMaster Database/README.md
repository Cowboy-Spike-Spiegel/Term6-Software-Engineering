BUPT Software Engineering
===========================

##### BUPT软件工程大作业——充电桩系统



###### part: 数据库

###### contributors: 2020211376-马天成



### 数据库平台

1. Windows开发环境：mysql Ver 8.0.28 for Win64 on x86_64
2. Linux部署环境：mysql80-community-release-el7-7.noarch



### Linux部署步骤

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

