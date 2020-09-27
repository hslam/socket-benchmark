# unix

```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	3.11s
	Requests per second:	32146.80
	Fastest time for request:	0.01ms
	Average time per request:	0.03ms
	Slowest time for request:	0.48ms

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
	Request rate per second:	16459159.10 Byte/s (16.46 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	16459159.10 Byte/s (16.46 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)

```
poll
```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	2.76s
	Requests per second:	36281.96
	Fastest time for request:	0.01ms
	Average time per request:	0.02ms
	Slowest time for request:	0.45ms

Time:
	0.1%	time for request:	0.01ms
	1%	time for request:	0.01ms
	5%	time for request:	0.01ms
	10%	time for request:	0.02ms
	25%	time for request:	0.02ms
	50%	time for request:	0.02ms
	75%	time for request:	0.03ms
	90%	time for request:	0.03ms
	95%	time for request:	0.04ms
	99%	time for request:	0.06ms
	99.9%	time for request:	0.16ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	18576361.36 Byte/s (18.58 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	18576361.36 Byte/s (18.58 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)
```