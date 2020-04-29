# test

Run in container:

docker build . -t test
docker run -p 8080:8080 test

ab -n 10000 -c 10 localhost:8080/request

Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        fasthttp
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /request
Document Length:        2 bytes

Concurrency Level:      10
Time taken for tests:   8.691 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1550000 bytes
HTML transferred:       20000 bytes
Requests per second:    1150.57 [#/sec] (mean)
Time per request:       8.691 [ms] (mean)
Time per request:       0.869 [ms] (mean, across all concurrent requests)
Transfer rate:          174.16 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.8      1      37
Processing:     2    8   2.7      7      43
Waiting:        2    7   2.6      7      41
Total:          3    8   2.8      8      45

Percentage of the requests served within a certain time (ms)
  50%      8
  66%      9
  75%     10
  80%     10
  90%     12
  95%     13
  98%     15
  99%     17
 100%     45 (longest request)


Run local:

go build
./test

ab -n 10000 -c 10 localhost:8080/request

Server Software:        fasthttp
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /request
Document Length:        2 bytes

Concurrency Level:      10
Time taken for tests:   2.663 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1550000 bytes
HTML transferred:       20000 bytes
Requests per second:    3754.53 [#/sec] (mean)
Time per request:       2.663 [ms] (mean)
Time per request:       0.266 [ms] (mean, across all concurrent requests)
Transfer rate:          568.31 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   7.3      1     338
Processing:     0    1   8.8      1     338
Waiting:        0    1   7.0      1     338
Total:          0    2  11.4      2     340

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      2
  75%      3
  80%      3
  90%      3
  95%      3
  98%      4
  99%      4
 100%    340 (longest request)
