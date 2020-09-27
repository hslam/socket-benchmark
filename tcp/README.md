# tcp

```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	5.87s
	Requests per second:	17026.40
	Fastest time for request:	0.03ms
	Average time per request:	0.06ms
	Slowest time for request:	0.70ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.05ms
	50%	time for request:	0.05ms
	75%	time for request:	0.06ms
	90%	time for request:	0.07ms
	95%	time for request:	0.08ms
	99%	time for request:	0.11ms
	99.9%	time for request:	0.28ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	8717517.03 Byte/s (8.72 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	8717517.03 Byte/s (8.72 MByte/s)

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
	Total time:	4.87s
	Requests per second:	20547.94
	Fastest time for request:	0.03ms
	Average time per request:	0.05ms
	Slowest time for request:	0.56ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.04ms
	50%	time for request:	0.04ms
	75%	time for request:	0.05ms
	90%	time for request:	0.06ms
	95%	time for request:	0.06ms
	99%	time for request:	0.10ms
	99.9%	time for request:	0.26ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	10520545.06 Byte/s (10.52 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	10520545.06 Byte/s (10.52 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)
```