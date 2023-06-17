BUPT Software Engineering
===========================

##### BUPT软件工程大作业——充电桩系统



###### part: 客户端前端

###### contributors: 2020211377-孟宇航、2020211393-潘婷



#### 1. 环境依赖

TypeScript  v5.1.3

node.js  v18.14.2

vue3



#### 2. 开发工具

WebStorm  v2023.1.2



#### 3. 客户端前端目录结构描述

│  .gitignore
│  env.d.ts
│  index.html
│  package-lock.json
│  package.json
│  README.md
│  tsconfig.app.json
│  tsconfig.json
│  tsconfig.node.json
│  vite.config.ts
│
├─public
│      favicon.ico
│      reset.css
│
└─src
    │  App.vue
    │  list.txt
    │  main.ts
    │
    ├─assets
    │  │  base.css
    │  │  logo.svg
    │  │  main.css
    │  │
    │  └─img
    │          1.svg
    │          2.svg
    │
    ├─class
    │      CarMsg.ts
    │      Login.ts
    │      Register.ts
    │      UserMsg.ts
    │      UserOrder.ts
    │
    ├─components
    │  │  AddCar.vue
    │  │  AddOrder.vue
    │  │  CarMessageView.vue
    │  │  DeleteCurrentOrder.vue
    │  │  EditOrder.vue
    │  │  EditUserName.vue
    │  │  EditUserPassword.vue
    │  │  LoginForm.vue
    │  │  Map.vue
    │  │  OrderDetail.vue
    │  │  OrderMessageView.vue
    │  │  RegisterForm.vue
    │  │  Report.vue
    │  │  updateCarCurrent.vue
    │  │  UpdateCarName.vue
    │  │  updateCarPower.vue
    │  │  UpdateOrder.vue
    │  │  UserMessage.vue
    │  │
    │  └─icons
    │          IconCommunity.vue
    │          IconDocumentation.vue
    │          IconEcosystem.vue
    │          IconSupport.vue
    │          IconTooling.vue
    │
    ├─http
    │  └─index
    │          api.ts
    │          index.ts
    │          request.ts
    │
    ├─router
    │      index.ts
    │
    └─views
            HomeView.vue
            LoginRegister.vue



#### 4. 部署说明

使用WebStorm创建vue3项目，将工程文件夹覆盖当前项目，在终端使用npm run dev 运行

##### 4.1 Project Setup

```sh
npm install
```

  在启动程序之前首先进行依赖的安装，将原有的module文件夹删除，然后在终端上执行上述命令即可。

##### 4.2 Port Setting

```
target:'http://121.43.119.64:8848/',
```

  在vue.config.js这一文件中找到该行，然后修改对应的服务端ip地址即可

##### 4.3 Compile and Hot-Reload for Development

```sh
npm run dev
```

安装了对应依赖并配置正确的IP地址之后，利用上述命令进行执行即可。

##### 4.4 Type-Check, Compile and Minify for Production

```sh
npm run build
```







