# Alphaquark API
_알파쿼크(AQT) 정보 조회 API_


<img src="https://img.shields.io/badge/-Golang-000000?style=flat&logo=Go">  <img src="https://img.shields.io/badge/-Python-000000?style=flat&logo=Python"> <img src="https://img.shields.io/badge/-Fastapi-000000?style=flat&logo=Fastapi">  <img src="https://img.shields.io/badge/-Redis-000000?style=flat&logo=Redis">  <img src="https://img.shields.io/badge/-Docker-000000?style=flat&logo=Docker">  <img src="https://img.shields.io/badge/-AWS EC2-000000?style=flat&logo=Amazon AWS">

## API URL  
GET https://openapi.alphaquark.io/api/aqt/info

<br/>

## Notice
- Rate limit : 100
- 5분 동안 허용되는 단일 IP 주소의 최대 요청 수
- 지속적으로 평가되며 이 제한에 도달하면 요청이 차단 ( return 404 forbidden )
- IP 주소가 제한 이하로 떨어지면 자동으로 차단이 해제

### Simple Benchmark (AB, Apache HTTP server benchmarking tool)

```bash
Document Path:          /api/aqt/info
Document Length:        1190 bytes

Concurrency Level:      10
Time taken for tests:   0.394 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      133600 bytes
HTML transferred:       119000 bytes
Requests per second:    253.94 [#/sec] (mean)
Time per request:       39.379 [ms] (mean)
Time per request:       3.938 [ms] (mean, across all concurrent requests)
Transfer rate:          331.31 [Kbytes/sec] received
```

<br/>

## More
- **API documentation**   
https://openapi.alphaquark.io/docs
