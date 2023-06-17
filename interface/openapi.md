---
title: 充电桩调度接口 v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# 充电桩调度接口

> v1.0.0

Base URLs:

Shared URL: https://apifox.com/apidoc/shared-c62b25a8-c763-4a4e-9baa-05a12aff450e

# 管理员组

## GET 查询充电站信息

GET /admin/index/manage

展示充电桩表单，每行后面跟开关
最上边显示全部开/关

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 17,
  "msg": "ex minim ullamco quis reprehenderit",
  "data": {
    "station": {
      "id": "21",
      "charge_array": [
        {
          "id": 50,
          "total_charge_duration": 95,
          "total_charge_fee": 97,
          "total_charge_times": 644311563206,
          "state": false,
          "mode": "Duis dolore dolore fugiat",
          "price": 73,
          "car_blocks": [
            {
              "name": "众但马约机很",
              "power_capicity": 43,
              "car_id": "53",
              "state": "ut irure dolore commodo",
              "power_current": 15,
              "user_id": "41"
            },
            {
              "name": "别从放任大色",
              "power_capicity": 5,
              "car_id": "62",
              "state": "sint ut dolor veniam",
              "power_current": 98,
              "user_id": "10"
            },
            {
              "name": "金第查物油",
              "power_capicity": 56,
              "car_id": "92",
              "state": "aliquip culpa",
              "power_current": 75,
              "user_id": "33"
            },
            {
              "name": "术技军南应",
              "power_capicity": 81,
              "car_id": "47",
              "state": "ut sit mollit exercitation labore",
              "power_current": 45,
              "user_id": "60"
            }
          ],
          "total_charge": 95,
          "total_charge_service": 32,
          "total_fee": 36,
          "name": "使强该住"
        },
        {
          "id": 69,
          "total_charge_duration": 44,
          "total_charge_fee": 53,
          "total_charge_times": 1596488664661,
          "state": false,
          "mode": "exercitation",
          "price": 39,
          "car_blocks": [
            {
              "name": "二提认指斗",
              "power_capicity": 15,
              "car_id": "2",
              "state": "nisi ut",
              "power_current": 50,
              "user_id": "33"
            }
          ],
          "total_charge": 24,
          "total_charge_service": 25,
          "total_fee": 11,
          "name": "明海具美"
        }
      ],
      "wait_array": [
        {
          "name": "只劳民只",
          "power_capicity": 64,
          "car_id": "19",
          "state": "exercitation Excepteur laboris deserunt",
          "power_current": 61,
          "user_id": "14"
        },
        {
          "name": "后听号",
          "power_capicity": 16,
          "car_id": "33",
          "state": "esse cillum",
          "power_current": 50,
          "user_id": "32"
        },
        {
          "name": "界便土厂工",
          "power_capicity": 68,
          "car_id": "16",
          "state": "in anim",
          "power_current": 61,
          "user_id": "96"
        },
        {
          "name": "长专型快说",
          "power_capicity": 54,
          "car_id": "73",
          "state": "voluptate ut minim",
          "power_current": 97,
          "user_id": "94"
        },
        {
          "name": "面体难成真",
          "power_capicity": 49,
          "car_id": "62",
          "state": "ipsum fugiat dolor qui",
          "power_current": 31,
          "user_id": "43"
        }
      ],
      "name": "传观前用火光"
    }
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|返回状态码|可自定义|
|» msg|string|true|none||none|
|» data|object¦null|true|none||这里面的数据一定是系统启动后的|
|»» id|string|true|none|id标识|none|
|»» charge_array|[[Charge](#schemacharge)]|true|none|充电桩数组|充电站下的充电桩|
|»»» 充电桩|[Charge](#schemacharge)|false|none|充电桩|none|
|»»»» id|string|true|none|充电桩id|唯一分配|
|»»»» name|string|true|none|名字|none|
|»»»» mode|string|true|none|充电模式|F- 快充 T- 慢充|
|»»»» state|string|true|none|工作状态|OFF - 关机状态<br />WORK - 充电状态<br />REST - 空闲状态<br />FAULT - 故障状态|
|»»»» total_charge|number|true|none|总充电总量|none|
|»»»» total_charge_times|number|true|none|总充电次数|none|
|»»»» total_charge_duration|number|true|none|总充电时长|单位为小时|
|»»»» total_charge_fee|number|true|none|总充电费用|单位为RMB|
|»»»» total_charge_service|number|true|none|总服务费用|单位为RMB|
|»»»» total_fee|number|true|none|总费用|单位为RMB|
|»»»» car_blocks|[[CarBlock](#schemacarblock)]|true|none|充电车位数组|空格隔开，大小为2，第一个位置和第二个位置|
|»»»»» 车位信息|[CarBlock](#schemacarblock)|false|none|车位信息|none|
|»»»»»» number|string|true|none|此车位的车辆分配的number|none|
|»»»»»» car_id|string|true|none|车辆ID|none|
|»»»»»» user_id|string|true|none|用户ID|none|
|»»»»»» name|string|false|none||none|
|»»»»»» power_capacity|number|true|none|总电量|none|
|»»»»»» power_current|number|true|none|当前电量|none|
|»»»»»» apply_kwh|number|true|none|申请的充电量|none|
|»»»»»» wait_time|number|true|none|等待时间|单位为小时|
|»»»»»» state|string|true|none|车位内的车辆状态|WAITING - 在等待区等待DISPATCHED - 在充电桩内等待CHARGING - 充电中|
|»» wait_array|[[CarBlock](#schemacarblock)]|true|none|等待车位数组|内容是车位|
|»»» 车位信息|[CarBlock](#schemacarblock)|false|none|车位信息|none|
|»»»» number|string|true|none|此车位的车辆分配的number|none|
|»»»» car_id|string|true|none|车辆ID|none|
|»»»» user_id|string|true|none|用户ID|none|
|»»»» name|string|false|none||none|
|»»»» power_capacity|number|true|none|总电量|none|
|»»»» power_current|number|true|none|当前电量|none|
|»»»» apply_kwh|number|true|none|申请的充电量|none|
|»»»» wait_time|number|true|none|等待时间|单位为小时|
|»»»» state|string|true|none|车位内的车辆状态|WAITING - 在等待区等待DISPATCHED - 在充电桩内等待CHARGING - 充电中|

## GET 查询单个充电桩报表展示

GET /admin/index/report

查看充电桩单个的状态，可展示历史信息

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|charge_id|query|string| 否 |none|
|token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "id": "string",
    "name": "string",
    "mode": "string",
    "state": "string",
    "total_charge": 0,
    "total_charge_times": 0,
    "total_charge_duration": 0,
    "total_charge_fee": 0,
    "total_charge_service": 0,
    "total_fee": 0,
    "time": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|
|»» id|string|true|none|充电桩id|唯一分配|
|»» name|string|true|none|名字|none|
|»» mode|string|true|none|充电模式|F- 快充 T- 慢充|
|»» state|string|true|none|工作状态|OFF - 关机状态<br />WORK - 充电状态<br />REST - 空闲状态<br />FAULT - 故障状态|
|»» total_charge|number|true|none|总充电总量|none|
|»» total_charge_times|number|true|none|总充电次数|none|
|»» total_charge_duration|number|true|none|总充电时长|单位为小时|
|»» total_charge_fee|number|true|none|总充电费用|单位为RMB|
|»» total_charge_service|number|true|none|总服务费用|单位为RMB|
|»» total_fee|number|true|none|总费用|单位为RMB|
|»» time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|

## POST 管理员登录

POST /admin/login

前端页面发送登录请求的路径

> Body 请求参数

```json
{
  "account": "u.zomc@oeoeh.mz",
  "password": "thisIsPass",
  "session": ""
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» account|body|string| 是 | 账户名|注册的用户名（邮箱格式）|
|» password|body|string| 是 | 密码|限制长度32|

#### 详细说明

**» account**: 注册的用户名（邮箱格式）
参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$

**» password**: 限制长度32
参考正则式：^.{1,32}$

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "token": "asfjecnnasiws",
    "account": "_Test_2020211226"
  }
}
```

```json
{
  "code": -1,
  "data": {
    "token": "",
    "account": ""
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|0表示正常，不为0表示注册失败|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|
|»» token|string|true|none|用户token|none|

## GET 查询充电桩服务的车辆信息

GET /admin/index/charge

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|charge_id|query|string| 否 ||none|
|token|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "id": "string",
    "name": "string",
    "mode": "string",
    "state": "string",
    "car_blocks": [
      {
        "number": "string",
        "car_id": "string",
        "user_id": "string",
        "name": "string",
        "power_capacity": 0,
        "power_current": 0,
        "apply_kwh": 0,
        "wait_time": 0,
        "state": "string"
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none||none|
|» data|object¦null|true|none||none|
|»» id|string|true|none|充电桩id|唯一分配|
|»» name|string|true|none|名字|none|
|»» mode|string|true|none|充电模式|F- 快充 T- 慢充|
|»» state|string|true|none|工作状态|OFF - 关机状态<br />WORK - 充电状态<br />REST - 空闲状态<br />FAULT - 故障状态|
|»» car_blocks|[[CarBlock](#schemacarblock)]|true|none|充电车位数组|空格隔开，大小为2，第一个位置和第二个位置|
|»»» 车位信息|[CarBlock](#schemacarblock)|false|none|车位信息|none|
|»»»» number|string|true|none|此车位的车辆分配的number|none|
|»»»» car_id|string|true|none|车辆ID|none|
|»»»» user_id|string|true|none|用户ID|none|
|»»»» name|string|false|none||none|
|»»»» power_capacity|number|true|none|总电量|none|
|»»»» power_current|number|true|none|当前电量|none|
|»»»» apply_kwh|number|true|none|申请的充电量|none|
|»»»» wait_time|number|true|none|等待时间|单位为小时|
|»»»» state|string|true|none|车位内的车辆状态|WAITING - 在等待区等待DISPATCHED - 在充电桩内等待CHARGING - 充电中|

## GET 管理员首页

GET /admin/index

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 客户组

## POST 用户登录

POST /client/login

前端页面发送登录请求的路径

> Body 请求参数

```json
{
  "account": "n.ykfcdp@dpoyx.kh",
  "password": "dolore sint deserunt ipsum cupidatat"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» account|body|string| 是 | 账户名|注册的用户名（邮箱格式）|
|» password|body|string| 否 | 密码|^.{1,32}$|

#### 详细说明

**» account**: 注册的用户名（邮箱格式）
参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "nostrud exercitation",
  "data": {
    "token": "veniam Lorem"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|0表示成功登录，非0表示账号或密码错误|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|
|»» token|string¦null|true|none||none|

## GET 查询订单信息

GET /client/index/order

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|mode|query|string| 是 ||两个值，HISTORY 和 CURRENT，表示历史订单和当前订单|
|token|header|string| 是 ||none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "commodo veniam",
  "data": [
    {
      "front_cars": 17,
      "id": "27",
      "user_id": "48",
      "car_id": "6",
      "mode": "id irure laboris dolor amet",
      "apply_kwh": 28,
      "charge_kwh": 38,
      "state": "cupidatat veniam esse",
      "charge_price": 54,
      "service_price": 53,
      "fee": 80,
      "create_time": "nisi quis ut sed amet",
      "dispatch_time": "ullamco sit",
      "charge_id": "94",
      "start_time": "veniam dolore sint",
      "finish_time": "sunt Lorem adipisicing",
      "number": 99
    },
    {
      "front_cars": 16,
      "id": "76",
      "user_id": "99",
      "car_id": "27",
      "mode": "nisi elit pariatur Excepteur",
      "apply_kwh": 73,
      "charge_kwh": 96,
      "state": "irure",
      "charge_price": 92,
      "service_price": 1,
      "fee": 39,
      "create_time": "mollit eiusmod deserunt",
      "dispatch_time": "Duis eiusmod",
      "charge_id": "86",
      "start_time": "Excepteur ad laboris tempor dolore",
      "finish_time": "sunt aute",
      "number": 90
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|[object]¦null|true|none||none|
|»» id|string|true|none|订单ID|none|
|»» user_id|string|true|none|用户ID|none|
|»» car_id|string|true|none|车辆ID|none|
|»» mode|string|true|none|充电模式|F快充，T慢充|
|»» apply_kwh|number|true|none|申请度数|可能不是最后充电的度数！|
|»» charge_kwh|number|true|none|实际充电度数|确切的充电度数，由时间算出|
|»» state|string|true|none|订单状态|WAITING - 在等待区等待<br />DISPATCHED - 在充电桩内等待<br />CHARGING - 充电中<br />FINISHED - 结束|
|»» charge_price|number|true|none|充电计价|平均充电计价|
|»» service_price|number|true|none|服务计价|服务计价|
|»» fee|number|true|none|订单最终花费|none|
|»» create_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|
|»» dispatch_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|
|»» charge_id|string|true|none|调度到的充电桩id|无为空串|
|»» start_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|
|»» finish_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|
|»» number|integer|false|none|排号号码|类似叫餐服务，订单当日递增|
|»» front_cars|integer|true|none|查看前面有多少车|排号号码-当前已进入充电区号码-1|

## POST 创建当前订单

POST /client/index/order

用户创建订单发出请求的路径

> Body 请求参数

```json
{
  "car_id": "string",
  "mode": "string",
  "apply_kwh": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» car_id|body|string| 是 | 车辆ID|none|
|» mode|body|string| 是 | 充电模式|none|
|» apply_kwh|body|number| 是 | 申请度数|none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "order": {
    "orderID": 56,
    "account": "eu_ut_reprehenderit_cillum",
    "price": 13,
    "kwh": 53,
    "createTime": "2012-01-29 01:46:55",
    "finishTime": "1986-07-16 08:27:13"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|成功返回200，失败返回-1|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|

## PATCH 更新当前订单

PATCH /client/index/order

前端需要检验是否修改，不修改则传参为空（不是空串）！！！

> Body 请求参数

```json
{
  "id": "string",
  "mode": "string",
  "apply_kwh": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 | 订单ID|none|
|» mode|body|string| 是 | 充电模式|F快充，T慢充|
|» apply_kwh|body|number| 是 | 申请度数|none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "oldOrder": {
    "orderID": 26,
    "account": "Excepteur_",
    "price": 61,
    "kwh": 3,
    "createTime": "1988-10-20 10:08:45",
    "finishTime": "1980-02-21 13:42:59"
  },
  "newOrder": {
    "orderID": 25,
    "account": "velit_ullamco",
    "price": 6,
    "kwh": 91,
    "createTime": "1990-12-27 00:47:42",
    "finishTime": "2005-03-28 07:04:56"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|

## DELETE 结束当前订单

DELETE /client/index/order

> Body 请求参数

```json
{
  "id": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 | 订单id|none|

> 返回示例

> 200 Response

```json
{
  "code": "string",
  "msg": "string",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|string|true|none||none|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|

## POST 开关充电桩

POST /admin/index/manage

> Body 请求参数

```json
{
  "mode": "string",
  "switch_array": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» mode|body|string| 是 | 开关模式|SWITCHON - 打开|
|» switch_array|body|[string]| 是 | 开关的充电站数组|none|

#### 详细说明

**» mode**: SWITCHON - 打开
SWITCHOFF - 关闭

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||状态码|
|» msg|string|true|none||none|
|» data|object¦null|true|none||none|

## POST 用户注册

POST /client/register

前端页面发出注册请求的路径
注册完默认返回一个token，认为登录，可跳转到index

> Body 请求参数

```json
{
  "account": "o.iylh@cdojfza.gf",
  "password": "qui sed ea",
  "name": "处风转国道"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» account|body|string| 是 | 账户名|注册的用户名（邮箱格式）|
|» password|body|string| 是 | 密码|^.{1,32}$|
|» name|body|string| 是 | 用户名|用户昵称|

#### 详细说明

**» account**: 注册的用户名（邮箱格式）
参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "token": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|0表示正常，不为0表示注册失败|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|
|»» token|string|true|none||none|

## GET 查询用户信息

GET /client/index/information

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "id": "string",
    "account": "string",
    "name": "string",
    "car_array": [
      {
        "car_id": "string",
        "user_id": "string",
        "name": "string",
        "power_capacity": 0,
        "power_current": 0
      }
    ],
    "balance": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|
|»» id|string|true|none|用户id|none|
|»» account|string|true|none|账户名|注册的用户名（邮箱格式）<br />参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$|
|»» name|string|true|none|用户名|none|
|»» car_array|[[Car](#schemacar)]|true|none|car_id数组|注册时初始化为大小为0|
|»»» 车辆信息|[Car](#schemacar)|false|none|车辆信息|none|
|»»»» car_id|string|true|none|车辆ID|none|
|»»»» user_id|string|true|none|用户ID|none|
|»»»» name|string|false|none||none|
|»»»» power_capacity|number|true|none|总电量|none|
|»»»» power_current|number|true|none|当前电量|none|
|»» balance|number|true|none|用户余额|none|

## POST 更新用户信息

POST /client/index/information

前端需要检验是否修改，不修改则传参为空（不是空串）！！！

> Body 请求参数

```json
{
  "old_password": "et Excepteur exercitation in",
  "name": "门联深物计",
  "new_password": "dolore est"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» old_password|body|string| 是 | 要更新的密码|none|
|» new_password|body|string| 是 ||none|
|» name|body|string| 是 | 要更新的名称|none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none|返回信息|失败的时候返回报错信息，成功的时候返回成功信息|
|» data|object¦null|true|none||none|

## GET 查询用户车辆信息

GET /client/index/car

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "nostrud",
  "data": [
    {
      "car_id": "32",
      "power_capacity": 33,
      "power_current": 12,
      "name": "车1"
    },
    {
      "car_id": "26",
      "power_capacity": 66,
      "power_current": 12,
      "name": "车2"
    },
    {
      "car_id": "58",
      "power_capacity": 63,
      "power_current": 40,
      "name": "车3"
    },
    {
      "car_id": "91",
      "power_capacity": 92,
      "power_current": 84,
      "name": "车4"
    },
    {
      "car_id": "55",
      "power_capacity": 71,
      "power_current": 17,
      "name": "车5"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none||none|
|» data|[object]¦null|true|none||none|
|»» car_id|string|true|none|车辆ID|none|
|»» name|string|false|none||none|
|»» power_capacity|number|true|none|总电量|none|
|»» power_current|number|true|none|当前电量|none|

## POST 新增用户车辆信息

POST /client/index/car

> Body 请求参数

```json
{
  "name": "string",
  "power_capacity": 0,
  "power_current": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» name|body|string¦null| 是 | 车辆名称|none|
|» power_capacity|body|number¦null| 是 | 总电量|none|
|» power_current|body|number¦null| 是 | 当前电量|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "添加成功!",
  "data": {}
}
```

```json
{
  "code": 1,
  "msg": "添加失败，报错信息：err",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none|返回信息|none|
|» data|object¦null|true|none|数据|none|

## PATCH 更新用户车辆信息

PATCH /client/index/car

前端需要检验是否修改，不修改则传参为空（不是空串）！！！

> Body 请求参数

```json
{
  "car_id": "string",
  "name": "string",
  "power_capacity": 0,
  "power_current": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» car_id|body|string| 是 | 车辆ID|none|
|» name|body|string| 是 | 车辆名称|none|
|» power_capacity|body|number| 是 | 总电量|none|
|» power_current|body|number| 是 | 当前电量|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "修改成功",
  "data": {}
}
```

```json
{
  "code": 1,
  "msg": "修改失败",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none|返回信息|none|
|» data|object¦null|true|none||none|

## DELETE 删除用户车辆信息

DELETE /client/index/car

可以不实现！！！

> Body 请求参数

```json
{
  "car_id": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» car_id|body|string| 是 | 车辆ID|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "删除成功",
  "data": {}
}
```

```json
{
  "code": 1,
  "msg": "删除失败",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none|状态码|none|
|» msg|string|true|none||none|
|» data|object¦null|true|none||none|

## POST 充值

POST /client/index/recharge

可以不实现！！！

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|amount|query|string| 否 ||单位人名币，精确到小数点后两文|
|token|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object¦null|true|none||none|

## GET 用户首页

GET /client/index

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|token|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_Car">Car</h2>

<a id="schemacar"></a>
<a id="schema_Car"></a>
<a id="tocScar"></a>
<a id="tocscar"></a>

```json
{
  "car_id": "string",
  "user_id": "string",
  "name": "string",
  "power_capacity": 0,
  "power_current": 0
}

```

车辆信息

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|car_id|string|true|none|车辆ID|none|
|user_id|string|true|none|用户ID|none|
|name|string|false|none||none|
|power_capacity|number|true|none|总电量|none|
|power_current|number|true|none|当前电量|none|

<h2 id="tocS_User">User</h2>

<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": "string",
  "account": "string",
  "password": "string",
  "name": "string",
  "car_array": [
    {
      "car_id": "string",
      "user_id": "string",
      "name": "string",
      "power_capacity": 0,
      "power_current": 0
    }
  ],
  "balance": 0
}

```

用户

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none|用户id|none|
|account|string|true|none|账户名|注册的用户名（邮箱格式）<br />参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$|
|password|string|true|none|密码|^.{1,32}$|
|name|string|true|none|用户名|none|
|car_array|[[Car](#schemacar)]|true|none|car_id数组|注册时初始化为大小为0|
|balance|number|true|none|用户余额|none|

<h2 id="tocS_Station">Station</h2>

<a id="schemastation"></a>
<a id="schema_Station"></a>
<a id="tocSstation"></a>
<a id="tocsstation"></a>

```json
{
  "id": "string",
  "name": "string",
  "charge_array": [
    {
      "id": "string",
      "name": "string",
      "mode": "string",
      "state": "string",
      "total_charge": 0,
      "total_charge_times": 0,
      "total_charge_duration": 0,
      "total_charge_fee": 0,
      "total_charge_service": 0,
      "total_fee": 0,
      "car_blocks": [
        {
          "number": "string",
          "car_id": "string",
          "user_id": "string",
          "name": "string",
          "power_capacity": 0,
          "power_current": 0,
          "apply_kwh": 0,
          "wait_time": 0,
          "state": "string"
        }
      ]
    }
  ],
  "wait_array": [
    {
      "number": "string",
      "car_id": "string",
      "user_id": "string",
      "name": "string",
      "power_capacity": 0,
      "power_current": 0,
      "apply_kwh": 0,
      "wait_time": 0,
      "state": "string"
    }
  ]
}

```

充电站

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none|id标识|none|
|name|string|false|none||none|
|charge_array|[[Charge](#schemacharge)]|true|none|充电桩数组|充电站下的充电桩|
|wait_array|[[CarBlock](#schemacarblock)]|true|none|等待车位数组|内容是车位|

<h2 id="tocS_CarBlock">CarBlock</h2>

<a id="schemacarblock"></a>
<a id="schema_CarBlock"></a>
<a id="tocScarblock"></a>
<a id="tocscarblock"></a>

```json
{
  "number": "string",
  "car_id": "string",
  "user_id": "string",
  "name": "string",
  "power_capacity": 0,
  "power_current": 0,
  "apply_kwh": 0,
  "wait_time": 0,
  "state": "string"
}

```

车位信息

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|number|string|true|none|此车位的车辆分配的number|none|
|car_id|string|true|none|车辆ID|none|
|user_id|string|true|none|用户ID|none|
|name|string|false|none||none|
|power_capacity|number|true|none|总电量|none|
|power_current|number|true|none|当前电量|none|
|apply_kwh|number|true|none|申请的充电量|none|
|wait_time|number|true|none|等待时间|单位为小时|
|state|string|true|none|车位内的车辆状态|WAITING - 在等待区等待DISPATCHED - 在充电桩内等待CHARGING - 充电中|

<h2 id="tocS_Charge">Charge</h2>

<a id="schemacharge"></a>
<a id="schema_Charge"></a>
<a id="tocScharge"></a>
<a id="tocscharge"></a>

```json
{
  "id": "string",
  "name": "string",
  "mode": "string",
  "state": "string",
  "total_charge": 0,
  "total_charge_times": 0,
  "total_charge_duration": 0,
  "total_charge_fee": 0,
  "total_charge_service": 0,
  "total_fee": 0,
  "car_blocks": [
    {
      "number": "string",
      "car_id": "string",
      "user_id": "string",
      "name": "string",
      "power_capacity": 0,
      "power_current": 0,
      "apply_kwh": 0,
      "wait_time": 0,
      "state": "string"
    }
  ]
}

```

充电桩

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none|充电桩id|唯一分配|
|name|string|true|none|名字|none|
|mode|string|true|none|充电模式|F- 快充 T- 慢充|
|state|string|true|none|工作状态|OFF - 关机状态<br />WORK - 充电状态<br />REST - 空闲状态<br />FAULT - 故障状态|
|total_charge|number|true|none|总充电总量|none|
|total_charge_times|number|true|none|总充电次数|none|
|total_charge_duration|number|true|none|总充电时长|单位为小时|
|total_charge_fee|number|true|none|总充电费用|单位为RMB|
|total_charge_service|number|true|none|总服务费用|单位为RMB|
|total_fee|number|true|none|总费用|单位为RMB|
|car_blocks|[[CarBlock](#schemacarblock)]|true|none|充电车位数组|空格隔开，大小为2，第一个位置和第二个位置|

<h2 id="tocS_Time">Time</h2>

<a id="schematime"></a>
<a id="schema_Time"></a>
<a id="tocStime"></a>
<a id="tocstime"></a>

```json
"string"

```

时间

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|时间|string|false|none|时间|xxxx-xx-xx xx:xx:xx|

<h2 id="tocS_Order">Order</h2>

<a id="schemaorder"></a>
<a id="schema_Order"></a>
<a id="tocSorder"></a>
<a id="tocsorder"></a>

```json
{
  "id": "string",
  "user_id": "string",
  "car_id": "string",
  "mode": "string",
  "apply_kwh": 0,
  "charge_kwh": 0,
  "state": "string",
  "charge_price": 0,
  "service_price": 0,
  "fee": 0,
  "create_time": "string",
  "dispatch_time": "string",
  "charge_id": "string",
  "start_time": "string",
  "finish_time": "string"
}

```

订单信息

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none|订单ID|none|
|user_id|string|true|none|用户ID|none|
|car_id|string|true|none|车辆ID|none|
|mode|string|true|none|充电模式|F快充，T慢充|
|apply_kwh|number|true|none|申请度数|可能不是最后充电的度数！|
|charge_kwh|number|true|none|实际充电度数|确切的充电度数，由时间算出|
|state|string|true|none|订单状态|WAITING - 在等待区等待<br />DISPATCHED - 在充电桩内等待<br />CHARGING - 充电中<br />FINISHED - 结束|
|charge_price|number|true|none|充电计价|平均充电计价|
|service_price|number|true|none|服务计价|服务计价|
|fee|number|true|none|订单最终花费|none|
|create_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|
|dispatch_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|
|charge_id|string|true|none|调度到的充电桩id|无为空串|
|start_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|
|finish_time|[Time](#schematime)|true|none|创建时间|xxxx-xx-xx xx:xx:xx|

