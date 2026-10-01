---
title: "Visual guide to SSH tunnels"
notion_id: f305d4bb-409e-49df-be5a-43f066708705
notion_url: https://app.notion.com/p/Visual-guide-to-SSH-tunnels-f305d4bb409e49dfbe5a43f066708705
last_edited: 2022-12-21T18:06:00.000Z
source_url: https://robotmoon.com/ssh-tunnels/
tags: ["English", "Network", "SysAdmin", "Article"]
---
This page explains use cases and examples of SSH tunnels while visually presenting the traffic flows. For example, here's a reverse tunnel that allows only users from IP address 1.2.3.4 access to port 80 on the SSH client through an SSH server.

SSH tunnels are encrypted TCP connections between SSH clients and servers that allows traffic entering one side of the tunnel to transparently exit through the other. While the term originally referred to tunnels using TUN/TAP virtual network interfaces, it's commonly used to refer to SSH port forwarding nowadays. Use cases include:

- Providing encrypted channels for protocols that use plaintext
- Opening backdoors into private networks
- Bypassing firewalls

This page is a work-in-progress. Feel free to [suggest examples or corrections on Github](https://github.com/linrock/ssh-tunnels)

## Port forwarding

### Forwards a port from one system (local or remote) to another

### Local port forwarding

Local port forwarding allows you to forward traffic on the SSH client to some destination through an SSH server. This lets you access remote services over an encrypted connection as if they were local services. Example use cases:

- Accessing a remote service (redis, memcached, etc.) listening on internal IPs
- Locally accessing resources available on a private network
- Transparently proxying a request to a remote service

### Remote port forwarding

Remote port forwarding allows you to forward traffic on an SSH server to a destination through either the SSH client or another remote host. This gives users on public networks access to resources on private networks. Example use cases:

- Making a local development server available over a public network
- Granting IP-restricted access to a remote resource on a private network

### Forwarding from privileged ports

If you want to open a privileged port (ports 1-1023) to forward traffic, you'll need to run SSH with superuser privileges on the system that opens the port.

`sudo ssh -L 80:example.com:80 ssh-server`

### SSH command-line flags

These are some useful SSH command-line flags when establishing tunnels

- `f # forks the ssh process into the background``n # prevents reading from STDIN``N # do not run remote commands. Used when only forwarding ports``T # disables TTY allocation`

Here's an example of a command you would run to create an SSH tunnel in the background that forwards a local port through the ssh server.

`ssh -fnNT -L 127.0.0.1:8080:example.org:80 ssh-server`

For brevity, the examples below will leave out these command-line flags.

## Local port forwarding

### Forwards connections from a port on a local system to a port on a remote host

`ssh -L 127.0.0.1:8080:example.org:80 ssh-server`

Forwards connections to 127.0.0.1:8080 on your local system to port 80 on example.org through ssh-server. The traffic between your local system and ssh-server is wrapped in an SSH tunnel, but the traffic between ssh-server and example.org is not. From the perspective of example.org the traffic originates from ssh-server.

`ssh -L 8080:example.org:80 ssh-server``ssh -L *:8080:example.org:80 ssh-server`

Forwards connections to port 8080 on all interfaces on your local system to example.org:80 through a tunnel to ssh-server.

### SSH configurations

Make sure that TCP forwarding is enabled on the SSH server. By default it should be.

`AllowTcpForwarding yes`

If you're forwarding ports on interfaces other than 127.0.0.1 then you'll need to enable GatewayPorts on your local system, either within ssh_config or as a command-line option

`GatewayPorts yes`

### Use cases

If you want to use a secure connection to access a remote service that communicates over plaintext. For example, redis and memcached all use plaintext protocols. If you securely access one of these services on a remote server over public networks, you can tunnel a connection from your local system to the remote server instead of having it listen over the public internet.

## Remote port forwarding

### Forwards a port on a remote system to another system

`ssh -R 8080:localhost:80 ssh-server`

Forwards traffic to all interfaces on port 8080 on ssh-server to localhost port 80 on your local computer. If one of these interfaces is available to the public internet, traffic connecting to port 8080 will be forwarded to your local system.

`ssh -R 8080:example.org:80 ssh-server`

Forwards traffic to all interfaces on ssh-server:8080 to localhost:80 on your local system. From your local system, traffic is then forwarded to example.org:80. From the perspective of example.org the traffic is originating from your local system.

### SSH server configuration

By default, forwarded ports are not accessible to the public internet. You'll need to add this to your sshd_config on your remote server to forward public internet traffic to your local computer.

`GatewayPorts yes`

Or if you'd like to specify which clients are allowed access, you can use the following in your sshd_config instead

`GatewayPorts clientspecified`

## Dynamic port forwarding

### Forward traffic from a range of ports to a remote server

`ssh -D 3000 ssh-server`

Opens a SOCKS proxy on port 3000 of all interfaces on your local system. This allows you to forward traffic sent through the proxy to the ssh-server on any port or destination host. By default, SSH will use the SOCKS5 protocol, which forwards TCP and UDP traffic.

When you have a SOCKS proxy running, you can configure your web browser to use the proxy to access resources as if connections were originating from ssh-server. For example, if ssh-server had access to other servers within a private network, by using using a SOCKS proxy you could access those other servers locally as if you were on the network, without needing to set up a VPN.

When you have a SOCKS proxy running, you can test it like so

`curl -x socks5://127.0.0.1:12345 https://example.org`

### SSH client configuration

If you want the SOCKS proxy to be available to more interfaces than just localhost, make sure to enable GatewayPorts on your local system.

`GatewayPorts yes`

Since GatewayPorts is being configured on the SSH client here, you can also configure it with a command-line option instead of ssh_config.

`ssh -o GatewayPorts=yes -D 3000 ssh-server`

## Jump hosts and proxy commands

### Transparently connecting to a remote host through intermediate hosts

`ssh -J user1@jump-host user2@remote-host``ssh -o "ProxyJump user1@jump-host" user2@remote-host`

Establishes an SSH connection with jump-host and forwards TCP traffic to remote-host. Connecting to remote-host through an intermediate jump-host. The above command should work out of the box if jump-host already has SSH access to remote-host. If it does not, you can use agent forwarding to forward the SSH identity of your local computer to remote-host.

`ssh -o ProxyCommand="nc -X 5 -x localhost:3000 %h %p" user@remote-host`

Connecting to a remote server through a SOCKS5 proxy using netcat. From the perspective of the server, the originating IP is from proxy-host. However, the SSH connection itself is end-to-end encrypted so proxy-host only sees an encrypted stream of traffic between the local system and remote-host.

### SSH client configuration

To enable agent forwarding, you can use ssh-add to add your local SSH identity to your local ssh agent.

`ssh-add`

## Reliable SSH Tunnels

### How to keep SSH tunnels open through network failures

The commands listed above work on an ad-hoc basis, but if you want to maintain SSH tunnels through network outages or unreliable connections, you'll have to do some additional setup.

By default, the TCP connection used to establish an SSH tunnel may time out after a period of inactivity. To prevent timeouts, you can configure the server to send heartbeat messages.

`ClientAliveInterval 15``ClientAliveCountMax 4`

You can also configure the client to send heartbeat messages.

`ServerAliveInterval 15``ServerAliveCountMax 4`

### Using AutoSSH

While the above options may prevent a connection from dropping due to inactivity, they will not re-establish dropped connections. To ensure that an SSH tunnel will be re-established, you can use autossh, which builds an SSH tunnel and monitors its health.

AutoSSH accepts the same arguments for port forwarding as SSH.

`autossh -R 2222:localhost:22 ssh-server`

This establishes a reverse tunnel that comes back after network failures. By default, AutoSSH will open extra ports on the SSH client and server for health checks. If traffic appears to no longer pass between the health check ports, AutoSSH will restart the SSH tunnel.

`autossh -R 2222:localhost:22 \``  -M 0 \``  -o "ServerAliveInterval 10" -o "ServerAliveCountMax 3" \``  remote-host`

Using the -M 0 flag disables the health check ports and allows the SSH client to handle the health checks. In this example, the SSH client expects the server to send a heartbeat every 10 seconds. If 3 heartbeats fail in a row, the SSH client exits, and AutoSSH will re-establish a new connection.

## More examples and use cases

### Transparent access to remote resource on a private network

Let's say there's a git repository on a private network that's only accessible through a private server on the network. This server is not accessible to the public internet. You have direct access to the server, but don't have VPN access to the private network.

For convenience, you'd like to access this private git repository as if you were connecting to it directly from your local system. If you have SSH access to another server that's accessible from both your local system and the private server, you can accomplish this by establishing an SSH tunnel and using a couple of ProxyCommand directives.

`ssh -L 127.0.0.1:22:127.0.0.1:2222 intermediate-host`

This forwards port 2222 on intermediate-host to port 22 on the private server. Now, if you SSH to port 2222 from intermediate-host, you're connecting to the SSH server on the private server despite the private server not being accessible by the public internet.

`ssh -p 2222 user@localhost`

If you'd like to make the backdoor even more convenient, you can add some directives to your local ~/.ssh/config

`Host git.private.network``  HostName git.private.network``  ForwardAgent yes``  ProxyCommand ssh private nc %h %p``Host private``  HostName localhost``  Port 2222``  User private-user``  ForwardAgent yes``  ProxyCommand ssh tunnel@intermediate-host nc %h %p`

Now you have access to the private git repository as if you were on the private network.

## Resources
