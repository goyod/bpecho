# How to deal with dependencies in Echo

~~~
Handler use log, http.Client, database, config
http.Client use config
http.Request use X-Request-ID, config
database use log, config
log use X-Request-ID
~~~