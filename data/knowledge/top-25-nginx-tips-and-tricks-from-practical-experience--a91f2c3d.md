---
title: "Top 25 Nginx Tips and Tricks From Practical Experience"
notion_id: a91f2c3d-8ed4-4190-bd86-b891cd6154b2
notion_url: https://app.notion.com/p/Top-25-Nginx-Tips-and-Tricks-From-Practical-Experience-a91f2c3d8ed44190bd86b891cd6154b2
last_edited: 2023-01-13T16:03:00.000Z
source_url: https://hackernoon.com/top-25-nginx-tips-and-tricks-from-practical-experience
tags: ["English", "Hosting", "Web Development", "SysAdmin", "Article", "Hackernoon"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## How NGINX works

NGINX is built on a non-blocking, event-driven architecture. It has one primary process, several worker_processes, and two cache management processes. To see this, you just need to run the _ps command_ with the switches:

```plain text
ps -ef --forest | grep nginx

```

The primary process runs as root and performs operations that require elevation, such as:

- 

reading config;

- 

opening ports;

- 

creation of new processes.

Workers do all the work:

- handle network connections;
- read and write data to disk;
- communicate with backend servers.

Usually, the number of workers equals the number of CPU cores. That allows you to use system resources as efficiently as possible.

Worker_processes listen on two kinds of sockets:

- listener - for new client events (connection requests);
- connection - for events that require processing.

The worker_processes wait for an event using two types of API: _epoll_ or _kqueue_. When receiving a new client event on the listener socket, the worker_process creates a new socket connection.

Having received an event on the connection socket, the worker_process executes this request. After processing the event, it does not get blocked, but it goes into standby mode and thus minimizes the number of context switches of an operating system.

With the above architecture, updating the NGINX configuration is as simple as sending a SIGHUP signal to the primary process. With NGINX adequately configured, each worker process is capable of processing hundreds of thousands of connections at the same time.

After that, the worker_processes will stop accepting new connections, and current connections, having processed the event, will be closed (without waiting for keep-alive). As soon as all connections are closed, the worker_process will end. That allows you to update the server configuration without losing connections, idle resources, or interrupting client service.

## Tips And Tricks

**1. Check the correctness of the written configuration before applying it**

After changing or creating a configuration, you need to check the syntax and the correctness of writing the config file. Nginx allows you to check the correctness of writing the configuration with the following command:

```plain text
nginx -t

```

**2. Reload Nginx without restarting the service**

It is a good practice to apply settings without restarting the service. This does not terminate current connections and eliminates the waiting period when loading the service, i.e. there is no downtime:

```plain text
nginx -s reload

```

or

```plain text
/etc/init.d/nginx reload

```

or

```plain text
service nginx reload

```

**3. Disable Nginx server_tokens**

By default, the server_tokens directive in Nginx displays the version number of Nginx. It is directly visible on all automatically generated error pages and in all HTTP responses in the server header.

This can lead to information disclosure - an unauthorized user can find out about your Nginx version. It would help if you disabled the `server_tokens` the directive in the Nginx config file by setting:

```plain text
server_tokens off;

```

**4. Disable legacy SSL/TLS protocols**

```plain text
ssl_protocols TLSv1.2 TLSv1.3;

```

**5. Disable any undesirable HTTP methods**

Disable any HTTP methods that will not be used and do not need to be implemented on the web server.

If you add the following condition to the location block of the Nginx virtual host configuration file, the server will only allow `GET, HEAD`, and `POST` methods and will filter out methods such as `DELETE` and `TRACE`:

```plain text
location / {
limit_except GET HEAD POST { deny all; }
}

```

Another approach is to add the following condition to the server section (or server block). It can be considered more generic, but you must be careful with [if statements](https://www.nginx.com/resources/wiki/start/topics/depth/ifisevil/?ref=hackernoon.com) in the location context:

```plain text
if ($request_method !~ ^(GET|HEAD|POST)$ ) {
    return 444; }

```

**6. Setup and configure Nginx access and error logs**

The Nginx access and error logs are enabled by default and located in `logs/error.log` and `logs/access.log,` respectively.

You can use the error_log directive in the Nginx configuration file if you want to change the location. You can also use this directive to specify the logs to be logged according to their severity level.

For example, critical severity will cause Nginx to log critical issues and all issues with a higher severity level than critical ones. To set the severity level to critical, set the error_log directive as follows:

```plain text
error_log logs/error.log crit;

```

You can find a complete list of error_log severity levels in the official Nginx documentation.

You can also change the access_log directive in the Nginx configuration file to specify a non-standard location for access logs.

Finally, you can use the log_format directive to customize the format of logged messages, as described in the Nginx documentation.

**7. Nginx worker connections**

One crucial setting we configure is the number of worker processes and the number of worker connections in the Nginx configuration file `/etc/nginx/nginx.conf`.

We will gradually adjust the workflow and worker connections to a higher or lower value to handle DDoS attacks:

This setting allows each worker process to handle up to 20,000 connections.

**8. Request rate limit**

Among the many valuable tactics to prevent DDoS attacks, one of the simplest and most effective is limiting the rate of incoming traffic.

For example, you can specify the rate at which NGINX accepts incoming requests to the average rate for your service from a specific client IP address for a specified period.

We set the limit_req_zone directive in the Nginx configuration file to limit the number of requests:

```plain text
limit_req_zone $binary_remote_addr zone=req_per_ip:1m rate=30r/m;

server {
    location /login {
        limit_req zone=req_per_ip;
    }
}

```

This example creates a storage area named one that can hold up to 16,000 (1m) unique IP addresses, and 30r/m means only 30 requests per minute are allowed.

We then use the limit_req directive to limit the speed of connections to a specific location or file, in this case, log in.

**9. Limiting the number of connections**

You can limit the number of connections that can open from a single client IP address, and here we are setting the limit_conn_zone and `limit_conn` directives to restrict the number of connections per IP address:

```plain text
limit_conn_zone $binary_remote_addr zone=conn_per_ip:1m;

server {
    location / {
        limit_conn conn_per_ip 10;
    }
}

```

This example creates a storage zone named `conn_per_ip` to store requests for the specified key, in this case, the client's IP address, $binary_remote_addr. The `limit_conn` the directive then sets ten connections from each client's IP address.

**10. Timeout options**

Slow connections may represent an attempt to keep connections open for a long time. As a result, the server cannot accept new connections:

```plain text
server {
    client_body_timeout 5s;
    client_header_timeout 5s;
}

```

In this example, the `client_body_timeout` directive specifies how long Nginx waits between client body entries and `client_header_timeout` specifies how long Nginx waits between client header entries. Both are set to 5 seconds.

**11. Limit the size of requests**

Similarly, large buffer values or large HTTP request sizes make DDoS attacks easier. So, we limit the following buffer values in the Nginx configuration file to mitigate DDoS attacks:

```plain text
client_body_buffer_size 200K;
client_header_buffer_size 2k;
client_max_body_size 200k;
large_client_header_buffers 3 1k;

```

Where,

`client_body_buffer_size 1k` — (default 8k or 16k) The directive specifies the size of the client request body buffer;

`client_header_buffer_size 1k` - directive sets the header buffer size for the request header from the client. A buffer size of 1Kb is sufficient for the vast majority of requests.

Increase this value if you have a custom header or a giant cookie sent by a client (like a wap client);

`client_max_body_size 1k` - the directive sets the maximum allowable size of the client request body, specified by the Content-Length line in the request header.

If the size is larger than the specified one, the client receives a "Request Entity Too Large" error (413). Increase this value when you are uploading files using the POST method;

`large_client_header_buffers 2 1k` - the directive assigns the maximum number and size of buffers for reading large headers from a client request.

By default, the size of one buffer is equal to the page size depending on the platform, it is either 4K or 8K.

If the connection goes into the keep-alive state at the end of the working request, then these buffers are released.

For example, 2x1k will accept a 2 KB data URI. It will also help fight bad bots and DoS attacks.

**12. IP Blacklist**

If you can identify the client IP addresses used for the attack, you can blacklist them with the deny directive so that NGINX and NGINX Plus won't accept their connections or requests:

```plain text
location / {
    deny 111.111.111.4;
    deny 111.111.111.0/24;
}

```

In this example, the first Deny blocks a specific IP address, and the second Deny blocks the entire range of IP addresses.

**13. IP Whitelist**

If your website or application is only allowed to be accessed from one or more specific sets or ranges of client IP addresses, you can use allow and deny directives together to allow only those addresses to access the site or application:

```plain text
location / {
    allow 111.111.111.4;
    deny all;
}

```

You can restrict access to only addresses on a specific local network. Here, the deny all directive blocks all client IP addresses that are not in the range specified by the allow directive.

**14. Block access to a file or location**

You can use Nginx to block access to a file or location completely. For example, if you notice that the register.php file is the target of an attack, you can completely block access to this file:

```plain text
location /register {
    deny all;
    return 444;
}

```

That will drop all requests that try to access this file. Code 444 closes the connection without a response.

**15. Enable sysctl based protection**

Enable sysctl-based protection. We can tweak the kernel and system variables on our server. Edit the `/etc/sysctl.conf` file and set these two lines to 1 as follows:

```plain text
net.ipv4.conf.all.rp_filter = 1
net.ipv4.tcp_syncookies = 1

```

The first setting enables IP address spoofing protection, and the second setting enables TCP SYN cookie protection.

**16. Nginx as a load balancer**

When Nginx is used as a load balancer, settings can be configured to limit the number of connections per server:

```plain text
upstream domain {
    server 111.111.111.4:80 max_conns=100;
    server 111.111.111.5:80 max_conns=100;
    queue 20 timeout=10s;
}

```

Here, the max_conns directive specifies the number of connections Nginx can open to the server. The queue directive limits the number of requests queued when all servers in that group have reached the connection limit.

Finally, the timeout directive specifies how long a request can be kept in the queue.

**17. Deny certain user agents**

You can easily block user agents i.e. crawlers, bots, and spammers that might be abusing your server:

```plain text
## Block download agents ##
if ($http_user_agent ~* LWP::Simple|BBBike|wget) {
    return 403;
}

```

```plain text
## Block some robots ##
if ($http_user_agent ~* msnbot|scrapbot|PxBroker) {
    return 403;
}

```

**18. Block Referral Spam**

Referral spam is dangerous. That can hurt your SEO ranking through weblogs (if published) since the referrer field links to their spam site. With these lines, you can block access to referrer spammers:

```plain text
if ( $http_referer ~* (babes|forsale|girl|jewelry|love|nudit|organic|poker|porn|sex|teen) )
{
    return 403;
}

```

**19. Stop image hotlinking**

An image hotlink or HTML means someone links your site to one of your images but displays it on their site.

You will end up paying for bandwidth and making content look like part of a site. This is usually done on forums and blogs. I highly recommend that you block and stop image hotlinking at the same level of your server:

```plain text
location /images/ {
  valid_referers none blocked www.domain.com domain.com;
   if ($invalid_referer) {
     return   403;
   }
}

```

**20. Limits the number of connections per IP address at the firewall level**

The web server must monitor connections and limit the number of connections per second. This serves 101. Both pf and iptables can block end-users before accessing your Nginx server.

Linux Iptables: Throttling Nginx connections per second.

In the following example, incoming connections will be dropped if the IP makes more than 15 connection attempts on port 80 within 60 seconds:

```plain text
/sbin/iptables -A INPUT -p tcp --dport 80 -i eth0 -m state --state NEW -m recent --set && \
/sbin/iptables -A INPUT -p tcp --dport 80 -i eth0 -m state --state NEW -m recent --update --seconds 60  --hitcount 15 -j DROP && \
service iptables save

```

**21. Disable content type sniffing in some browsers**

If a response specifies an incorrect content type, then browsers may process the response in unexpected ways.

If the content type is determined to be a renderable text-based format, then the browser will usually attempt to interpret the response as being in that format, regardless of the actual contents of the response.

Additionally, some other specified content types might sometimes be interpreted as HTML due to quirks in particular browsers. This might lead to otherwise “safe” content such as images being rendered as HTML, enabling cross-site scripting attacks in certain conditions. Add the following to your config:

`add_header X-Content-Type-Options nosniff;`

**22. Enable cross-site scripting (XSS) filter**

Add the following line to your configuration file to enable the X-XSS-Protection header on your Nginx web server. When you're done, save your changes and reload Nginx:

`add_header X-XSS-Protection "1; mode=block";`

**23. Configure HTTP/2 Support**

HTTP/2 is the successor to the HTTP 1.x network protocol. HTTP/2 is widely used to reduce latency, minimize protocol overhead, add support for request prioritization, and speed up web application loading.

Hence, it is vital to stay up to date with performance optimization techniques and strategies. The main focus of HTTP/2 is to reduce the overall load time of a web page, which optimizes performance.

It also focuses on the use of network and server resources, as well as increased security, since SSL/TLS encryption is mandatory when using HTTP/2.

As a prerequisite, ensure the Nginx version is 1.9.5 or higher as it is built, otherwise, you will have to add it manually, and the server must enable SSL/TLS. Your HTTPS server block should now look like this:

```plain text
server {
    listen 443 ssl http2;

    ssl_ciphers EECDH+CHACHA20:EECDH+AES128:RSA+AES128:EECDH+AES256:RSA+AES256:EECDH+3DES:RSA+3DES:!MD5;
}

```

**24. Add **_**Content-Security-Policy**_** (CSP)**

CSP is an additional layer of security that helps detect and mitigate certain types of attacks, including cross-site scripting (XSS) and data injection attacks. These attacks are used for data theft, website corruption, and malware distribution:

```plain text
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains; preload" always;

```

**25. Monitoring Nginx**

The state of your server must be constantly monitored. You can watch what the server returns by periodically sending requests to it, for example, through third-party paid services. Or you can fine-tune everything, and there are many ways to do it—more about it in the [Nginx blog](https://www.nginx.com/blog/monitoring-nginx/?ref=hackernoon.com).

**Instead conclusion**

It is essential to be flexible and develop a configuration for each case separately.
