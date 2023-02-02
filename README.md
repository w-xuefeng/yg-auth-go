# yg-auth-go
yg-auth 项目的 go 实现

- 1.统一登录
  - `POST /auth`
  - Request Body 结构如下：
    ````json
    {
      "stuid": "c3R1aWQ=",
      "password": "cGFzc3dvcmQ="
    }

  - Response 结构如下：
    ````json
    {
      "status": true,
      "resdata": {
        "token": "xxxxxxxxxx"
      }
    }

- 2.token 换取用户信息
  - `POST /auth`
  - Request Body 结构如下：
    ````json
    {
      "token": "xxxxxxxxxx",
    }

  - Response 结构如下：
    ````json
    {
       "status": true,
       "resdata": {
         "id": 1,
         "stuid": "111111",
         "name": "张三",
         "department": "开发部",
         "college": "信息工程学院",
         "majorclass": "计科2301",
         "email": "abc@def.com",
         "head": "/abc/def/head/123.png",
         "ifkey": 0,
         "online": 0,
         "birthday": "2023-01-01",
         "duty": "{\"week\": [2, 3], \"class\": [1, 2], \"state\": 0, \"dutydate\": []}",
         "loginip": "",
         "interviewform": "{}",
         "phone": "12345678911",
         "photo": "https://abc.def.com/photo/hi/jk/lmnopqrstuvwxyz.jpg",
         "position": 0,
         "positionName": "站员",
         "qq": "123456789",
         "qqid": "xxxx",
         "registerdate": "2023-01-01 14:57:24",
         "sex": "1",
         "signcount": 100,
         "ulevel": 10,
         "utype": 1,
         "wxid": "xxxxx",
         "token": "xxxxx"
       }
    }
    
- 3.校验注册码
  - `POST /auth`
  - Request Body 结构如下：
    ````json
    {
      "regcode": "cnJycnI="
    }

  - Response 结构如下：
    - 校验成功
      ````json
      {
        "status": true,
      }
      
    - 校验失败
      ```json
      {
        "status": false,
      }
