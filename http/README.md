# http

```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	5.91s
	Requests per second:	16930.75
	Fastest time for request:	0.03ms
	Average time per request:	0.06ms
	Slowest time for request:	2.46ms

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
	99.9%	time for request:	0.27ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	8668545.30 Byte/s (8.67 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	8668545.30 Byte/s (8.67 MByte/s)

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
	Total time:	4.82s
	Requests per second:	20762.67
	Fastest time for request:	0.03ms
	Average time per request:	0.05ms
	Slowest time for request:	0.61ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.04ms
	50%	time for request:	0.04ms
	75%	time for request:	0.05ms
	90%	time for request:	0.05ms
	95%	time for request:	0.06ms
	99%	time for request:	0.09ms
	99.9%	time for request:	0.24ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	10630487.57 Byte/s (10.63 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	10630487.57 Byte/s (10.63 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)
```