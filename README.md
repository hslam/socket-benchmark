# socket-benchmark

## tcp
```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	5.77s
	Requests per second:	17333.81
	Fastest time for request:	0.03ms
	Average time per request:	0.05ms
	Slowest time for request:	0.96ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.05ms
	50%	time for request:	0.05ms
	75%	time for request:	0.06ms
	90%	time for request:	0.07ms
	95%	time for request:	0.07ms
	99%	time for request:	0.10ms
	99.9%	time for request:	0.20ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	8874912.44 Byte/s (8.87 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	8874912.44 Byte/s (8.87 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=64
Summary:
	Clients:	64
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	9.78s
	Requests per second:	102273.92
	Fastest time for request:	0.03ms
	Average time per request:	0.62ms
	Slowest time for request:	17.49ms

Time:
	0.1%	time for request:	0.07ms
	1%	time for request:	0.31ms
	5%	time for request:	0.45ms
	10%	time for request:	0.51ms
	25%	time for request:	0.58ms
	50%	time for request:	0.62ms
	75%	time for request:	0.65ms
	90%	time for request:	0.70ms
	95%	time for request:	0.79ms
	99%	time for request:	1.20ms
	99.9%	time for request:	2.28ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	52364245.69 Byte/s (52.36 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	52364245.69 Byte/s (52.36 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## http
```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	5.78s
	Requests per second:	17307.40
	Fastest time for request:	0.03ms
	Average time per request:	0.05ms
	Slowest time for request:	0.90ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.05ms
	50%	time for request:	0.05ms
	75%	time for request:	0.06ms
	90%	time for request:	0.07ms
	95%	time for request:	0.07ms
	99%	time for request:	0.10ms
	99.9%	time for request:	0.22ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	8861390.89 Byte/s (8.86 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	8861390.89 Byte/s (8.86 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=64
Summary:
	Clients:	64
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	9.84s
	Requests per second:	101642.65
	Fastest time for request:	0.03ms
	Average time per request:	0.63ms
	Slowest time for request:	22.48ms

Time:
	0.1%	time for request:	0.09ms
	1%	time for request:	0.29ms
	5%	time for request:	0.42ms
	10%	time for request:	0.49ms
	25%	time for request:	0.57ms
	50%	time for request:	0.61ms
	75%	time for request:	0.65ms
	90%	time for request:	0.72ms
	95%	time for request:	0.85ms
	99%	time for request:	1.38ms
	99.9%	time for request:	2.78ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	52041035.17 Byte/s (52.04 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	52041035.17 Byte/s (52.04 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
## ipc
```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	3.14s
	Requests per second:	31845.37
	Fastest time for request:	0.01ms
	Average time per request:	0.03ms
	Slowest time for request:	8.72ms
Time:
	0.1%	time for request:	0.01ms
	1%	time for request:	0.01ms
	5%	time for request:	0.02ms
	10%	time for request:	0.02ms
	25%	time for request:	0.02ms
	50%	time for request:	0.03ms
	75%	time for request:	0.03ms
	90%	time for request:	0.04ms
	95%	time for request:	0.04ms
	99%	time for request:	0.07ms
	99.9%	time for request:	0.16ms
Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	16304828.97 Byte/s (16.30 MByte/s)
Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=64
Summary:
	Clients:	64
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	4.22s
	Requests per second:	236859.51
	Fastest time for request:	0.01ms
	Average time per request:	0.27ms
	Slowest time for request:	36.74ms
Time:
	0.1%	time for request:	0.02ms
	1%	time for request:	0.06ms
	5%	time for request:	0.11ms
	10%	time for request:	0.14ms
	25%	time for request:	0.19ms
	50%	time for request:	0.24ms
	75%	time for request:	0.32ms
	90%	time for request:	0.39ms
	95%	time for request:	0.46ms
	99%	time for request:	0.85ms
	99.9%	time for request:	1.81ms
Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	121272068.20 Byte/s (121.27 MByte/s)
Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## ws
```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	6.32s
	Requests per second:	15813.38
	Fastest time for request:	0.04ms
	Average time per request:	0.06ms
	Slowest time for request:	1.97ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.05ms
	10%	time for request:	0.05ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.06ms
	90%	time for request:	0.07ms
	95%	time for request:	0.08ms
	99%	time for request:	0.11ms
	99.9%	time for request:	0.28ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	8096450.23 Byte/s (8.10 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	8096450.23 Byte/s (8.10 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=64
Summary:
	Clients:	64
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	10.13s
	Requests per second:	98761.86
	Fastest time for request:	0.03ms
	Average time per request:	0.65ms
	Slowest time for request:	118.09ms

Time:
	0.1%	time for request:	0.05ms
	1%	time for request:	0.10ms
	5%	time for request:	0.31ms
	10%	time for request:	0.40ms
	25%	time for request:	0.51ms
	50%	time for request:	0.59ms
	75%	time for request:	0.69ms
	90%	time for request:	0.83ms
	95%	time for request:	1.00ms
	99%	time for request:	2.02ms
	99.9%	time for request:	7.48ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	50566073.36 Byte/s (50.57 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	50566073.36 Byte/s (50.57 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
