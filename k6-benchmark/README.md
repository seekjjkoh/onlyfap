# Benchmark with k6

## Result
Configuration is as below:
- Machine: Lenovo Thinkpad X1 Carbon (4th Gen) i7-6600U 16GB
- Docker: --cpus="0.1", --memory="128m"
- k6: --vus=100, --duration="30s"

|Language + Framework | #Requests | Avg | Median | P95 | P99.99 |
| --- | --- | --- | --- | --- | --- |
| Go | 43480 | 68.72ms | 88.84ms | 114.58ms | 350.14ms |
| Deno | 16044 | 186.91ms | 192.23ms | 300.46ms | 8.1s |
| Lua (Luvit) | 9626 | 316.42ms | 247.69ms | 699.96ms | 1.29s |
| Rust (Rocket) | 9468 | 315.74ms | 316.14ms | 501.3ms | 699.17ms |
| Erlang | 5391 | 527.66ms | 4.8ms | 97.94ms | 56.58s |
| Nodejs (Express) | 3089 | 985.14ms | 915.63ms | 1.4s | 2.6s |
| Python (Flask+Gunicorn) | 1809 | 1.7s | 1.7s | 2.29s | 5.59s |
| Scala (Play) | 201 | 12.96s | 2.79s | 53.16s | 59.26s |

Note that Scala (Play), Erlang metrics are over 30s.