Tested on default GAE instances(F1):

Memory - 128MB, CPU - 600MHz

#
Python:
```
➜ wrk -t8 -c100 -d10s https://python-dot-wixgamma.appspot.com/benchmark
Running 10s test @ https://python-dot-wixgamma.appspot.com/benchmark
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   190.40ms  107.29ms   1.26s    95.83%
    Req/Sec     9.95      6.15    30.00     71.47%
  446 requests in 10.10s, 193.39KB read
  Socket errors: connect 0, read 0, write 0, timeout 38
Requests/sec:     44.15
Transfer/sec:     19.14KB
```
#
NodeJS:
```
➜ wrk -t8 -c100 -d10s https://nodejs-dot-wixgamma.appspot.com/benchmark
Running 10s test @ https://nodejs-dot-wixgamma.appspot.com/benchmark
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   201.11ms  154.96ms   1.73s    97.98%
    Req/Sec    13.05      7.54    40.00     78.24%
  890 requests in 10.08s, 385.90KB read
  Socket errors: connect 0, read 0, write 0, timeout 48
Requests/sec:     88.26
Transfer/sec:     38.27KB
```
