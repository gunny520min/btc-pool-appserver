# btc-pool-appserver
 BTC.com AppServer Demo

> version: 1.0
> date: 2021/03/09

--------------------------------
2021/03/11
搭建框架

--------------------------------
2021/03/12
搭建框架

注释了不少部分，基本跑通了流程了，
http请求没成功，用的假数据
写了个banner的 ： 测试 curl https://localhost:8080/api/public/home/index -X POST

TODO：
* 初始化很多步骤都注释了。。回头慢慢加吧
* config 只简单写了一个
* 简单配置了conf中文件，目前Load后没有映射到config下到结构体中，不知道为什么。。

--------------------------------
2021/03/15
```
    gin -> handler(controller) {
        service -> Get {
            client -> Get {
                dorequest
            }
        }
    }
```