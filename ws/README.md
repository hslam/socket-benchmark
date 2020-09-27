# ws

```
./client -addr=:9999 -total=100000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	100000
	Total time:	6.30s
	Requests per second:	15862.45
	Fastest time for request:	0.04ms
	Average time per request:	0.06ms
	Slowest time for request:	0.66ms

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
	99%	time for request:	0.12ms
	99.9%	time for request:	0.32ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	8121573.61 Byte/s (8.12 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	8121573.61 Byte/s (8.12 MByte/s)

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
	Total time:	5.11s
	Requests per second:	19573.11
	Fastest time for request:	0.03ms
	Average time per request:	0.05ms
	Slowest time for request:	0.51ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.04ms
	50%	time for request:	0.04ms
	75%	time for request:	0.05ms
	90%	time for request:	0.06ms
	95%	time for request:	0.07ms
	99%	time for request:	0.10ms
	99.9%	time for request:	0.25ms

Request:
	Total request body sizes:	51200000
	Average body size per request:	512.00 Byte
	Request rate per second:	10021430.59 Byte/s (10.02 MByte/s)

Response:
	Total response body sizes:	51200000
	Average body size per response:	512.00 Byte
	Response rate per second:	10021430.59 Byte/s (10.02 MByte/s)

Result:
	Response ok:	100000 (100.00%)
	Errors:	0 (0.00%)
```