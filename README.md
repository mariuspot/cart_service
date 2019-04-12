# Cart Service

## Calls
The api provides 6 calls
* create cart 
* add item
* remote item
* remove all / empty
* get items
* convert cart to order

## Design notes and some possibly wrong assumtions :)

Ideally the configuration should use something like etcd instead of all the flags.

Since the line_items table has a refference to products I assumed it should be product aware (not just a generic cart service that stores id, quantity and price)
Since there are no refferences between users and carts I assumed that the calling side will handle linking users to carts.

If it's an internal service then using certs should be secure and the service should be limited to only accepting internal traffic, but if the service is going to be publically accessible the request should be authenticated.

Both line_items and products have a price field. I wasn't sure on the purpose for this. (Maybe to store the final price once the cart gets converted into an order)

## Setup

Script to create the database is included in scripts folder under `narnes_and_bobles.sql`

## Running the server
Starting the service with `-tls` flag and without providing a cert and key file will use the test certs provided by `google.golang.org/grpc/testdata`
```
$ ./nab_cart_service -tls 
2019/04/12 10:39:14 Starting metrics listener on 127.0.0.1:2112
2019/04/12 10:39:14 Starting server listening on 127.0.0.1:10000
```

The server requires access to a database created with the `narnes_and_bobles.sql` script in the `scripts` folder.

The database access can be provided with `-db_ip`, `-db_port`, `-db_username`, `-db_password` and `-db_name` flags.

```
$ ./nab_cart_service -help
Usage of ./nab_cart_service:
  -cert_file string
    	The TLS cert file
  -db_ip string
    	Database ip (default "127.0.0.1")
  -db_name string
    	Database name (default "narnes_and_boble")
  -db_password string
    	Database password (default "password")
  -db_port int
    	Database port (default 3306)
  -db_username string
    	Database username (default "username")
  -ip string
    	The ip to listen on for the server (default "127.0.0.1")
  -key_file string
    	The TLS key file
  -metric_ip string
    	Metrics ip (prometheus) (default "127.0.0.1")
  -metric_port int
    	Metrics port (prometheus) (default 2112)
  -port int
    	The port to listen on for the server (default 10000)
  -tls
    	Connection uses TLS if true, else plain TCP
```

## Running the example
An example is included in `examples/cart_api`. 

Running the example with `-tls` flag and without providing a ca file will use the test certs provided by `google.golang.org/grpc/testdata`
```
$ ./cart_api -tls
2019/04/12 10:39:45 Cart created, cart_id:24
2019/04/12 10:39:45 Items added to cart (3 of product 1)
2019/04/12 10:39:45 Item removed from cart (1 of product 1)
2019/04/12 10:39:45 Item all of product from cart (all of product 1)
2019/04/12 10:39:45 Expected, error converting cart to order (err: rpc error: code = Internal desc = Error converting cart to order (err: Cart is empty))
2019/04/12 10:39:45 Items added to cart (3 of product 1)
2019/04/12 10:39:45 Emptied cart (removed everything)
2019/04/12 10:39:45 Items added to cart (3 of product 1)
2019/04/12 10:39:45 Items added to cart (2 of product 2)
Item 12019/04/12 10:39:45 	Title: Beans
	Description: An idept look at beans and their uses
	ImageUrl:
	Quantity: 3
	Price: 14.95
Item 22019/04/12 10:39:45 	Title: Stuff
	Description: Some stuffs
	ImageUrl:
	Quantity: 2
	Price: 9.95
2019/04/12 10:39:45 Cart converted to order, order_id: 24
```

It includes example for 
* connecting to the cart service
* creating a cart
* adding items to a cart
* removing items from a cart
* emptying a cart
* listing items in a cart
* converting a cart into an order

```
$ ./cart_api -help
Usage of ./cart_api:
  -ca_file string
    	The file containning the CA root cert file
  -server_addr string
    	The server address in the format of host:port (default "127.0.0.1:10000")
  -server_host_override string
    	The server name use to verify the hostname returned by TLS handshake (default "x.test.youtube.com")
  -tls
    	Connection uses TLS if true, else plain TCP
```
## Metrics
Metrics for prometheus are by default export on `http://localhost:2112/metrics` but can be changed using the `-metric_ip` and `-metric_port` flags

There are 4 service specific metrics
* cart_service_query_counter: The number of api calls since the service was started
* cart_service_query_time_summary: A summary of the query times for the api (0.99, 0.9 and 0.5 percentiles)
* cart_service_database_query_counter: The number of database queries since the service was started
* cart_service_database_query_time_summary: A summary of the database query times
```
$ curl http://localhost:2112/metrics
# HELP cart_service_database_query_counter The total number of database queries
# TYPE cart_service_database_query_counter counter
cart_service_database_query_counter 1
# HELP cart_service_database_query_time_summary The database query times summary
# TYPE cart_service_database_query_time_summary summary
cart_service_database_query_time_summary{quantile="0.5"} NaN
cart_service_database_query_time_summary{quantile="0.9"} NaN
cart_service_database_query_time_summary{quantile="0.99"} NaN
cart_service_database_query_time_summary_sum 0
cart_service_database_query_time_summary_count 0
# HELP cart_service_request_time_summary The request times summary
# TYPE cart_service_request_time_summary summary
cart_service_request_time_summary{quantile="0.5"} NaN
cart_service_request_time_summary{quantile="0.9"} NaN
cart_service_request_time_summary{quantile="0.99"} NaN
cart_service_request_time_summary_sum 0
cart_service_request_time_summary_count 0
# HELP cart_service_requests The total number of requests
# TYPE cart_service_requests counter
cart_service_requests 0
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.12.1"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 640520
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 640520
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.443459e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 197
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 2.240512e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 640520
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.4700416e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.88416e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 2579
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.6584576e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 2776
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 27504
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 32768
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.313397e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 524288
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 524288
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.2155384e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 10
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
