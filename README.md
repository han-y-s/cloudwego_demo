# student_demo
## 开启服务
1.开启consul
``` bash
consul agent -dev -server -bootstrap-expect 1 -data-dir /tmp/consul -node=node1
```
2.运行kitex
```
./kitex_demo
```
3.运行hertz
```
./hertz
```
## api使用
### 注册
```
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8887/add-student-info -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@pku.com"]}'
```
### 查询
```
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8887/query\?id\=100
```
### 热更新
```
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8887/ping
```
### 热更新后重新注册
```
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8887/add-student-info -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@pku.com"], "gender":"male"}'
```
