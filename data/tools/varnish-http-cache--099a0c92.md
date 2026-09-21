---
title: "Varnish HTTP Cache"
notion_id: 099a0c92-5e6d-4add-8cbb-8da51806f6fc
notion_url: https://app.notion.com/p/Varnish-HTTP-Cache-099a0c925e6d4add8cbb8da51806f6fc
last_edited: 2022-12-19T18:34:00.000Z
source_url: https://varnish-cache.org/
tags: ["Web Development", "SysAdmin", "Untried", "Tool", "English"]
---
## What is Varnish Cache?

Varnish Cache is a distribution of the open-source Vinyl Cache project, made for the modern cloud: ready for Kubernetes, with built-in TLS support, and a long-term-support release cadence with extra modules and tooling layered on top.

Deploying behind HTTPS or inside a container platform? This is the distribution you want.

## How does Varnish work?

Varnish sits between clients and your origin servers. Requests come in over HTTP or HTTPS; Varnish answers from its in-memory cache when it can, and proxies through to the origin when it can't. This is illustrated in the diagram below.

A diagram that illustrates how Varnish works.

Cached objects honor the standard HTTP cache headers and are evicted when they expire.

Caching policy is expressed in VCL, the Varnish Configuration Language, which compiles to native code and runs inside the worker. This is what lets Varnish handle policy decisions at line rate.

## Key features

Varnish has a rich set of features and capabilities, which you can read about in the documentation.

## Performance

Intel and Varnish Software measured Varnish Enterprise at 1.5 Tbps per server on a single CDN node. The same high-performance TLS code now powers Varnish Cache.

On the cache-hit path, Varnish Cache is structurally lean. Per cached response, on a stock build:

- 3 system calls. read the request, poll the socket, writev the response. That's the entire kernel interaction.
- No IO logging. Varnish logs to shared memory, no disk IO needed.
- 0 kernel memory operations. No mmap, no brk, no munmap, Varnish never asks the kernel for a new page during cache-hit traffic. Workspaces and mempools are pre-allocated.

These are properties of the code, not of any particular benchmark setup.

## Getting started with Varnish

Varnish Cache runs on Linux and ships in container images, Helm charts, and native packages for the major DEB- and RPM-based distributions.

Pull in the Docker image:

docker pull varnish

Run the Docker container with the standard configuration:

docker run --rm -p 80:80 --name varnish varnish

Run the Docker container with a custom VCL file:

docker run --rm -p 80:80 --name varnish -v $(pwd)/default.vcl:/etc/varnish/default.vcl:ro varnish

Mounts default.vcl into the Docker container to customize the caching policies of Varnish. Makes a Varnish available on port 80 of your local machine.

Customize the VCL:

vcl 4.1; backend default { .host = "origin.example.com"; .port = "80"; }

This VCL change configures the backend server for Varnish to forward requests to origin.example.com on port 80. Use the -v $(pwd)/default.vcl:/etc/varnish/default.vcl:ro option in Docker to mount the default.vcl file and customize the caching policies of Varnish.

Learn more about deploying Varnish with Docker by reading the install guide.

Install guide

Install the Helm chart:

helm install varnish -f values.yaml oci://docker.io/varnish/varnish-cache

Customize the values.yaml file to configure the Varnish Helm chart.

Run kubectl get svc varnish-varnish-cache to get the cluster IP address and node port of the Varnish Cache service. Other service types and Ingress are also supported.

Learn more about deploying Varnish on Kubernetes by reading the install guide.

Install guide

Update the package list:

sudo apt-get update

Configure the package registry to install Varnish on Debian and Ubuntu Linux servers:

curl -Ls https://packages.varnish-software.com/varnish/bootstrap-deb.sh | sh

Install Varnish:

sudo apt-get install -y varnish

Edit the varnish systemd service:

sudo systemctl edit --full varnish

Change the port number and cache size in the varnish.service systemd parameters:

ExecStart=/usr/sbin/varnishd \ -a :80 \ -a localhost:8443,PROXY \ -f /etc/varnish/default.vcl \ -P %t/%N/varnishd.pid \ -p feature=+http2 \ -s malloc,2g

This systemd configuration change sets the Varnish service to listen on port 80, and allocates 2GB of memory for caching.

Open the VCL file:

sudo vim /etc/varnish/default.vcl

Customize the VCL:

vcl 4.1; backend default { .host = "origin.example.com"; .port = "80"; }

This VCL change configures the backend server for Varnish to forward requests to origin.example.com on port 80.

Restart Varnish to activate the changes:

sudo systemctl restart varnish

Learn more about installing Varnish on Debian or Ubuntu by reading the install guide.

Install guide

Configure the package registry to install Varnish on RPM-based Linux servers:

curl -s https://packages.varnish-software.com/varnish/bootstrap-rpm.sh | sh

Install Varnish:

sudo yum install -y varnish

Edit the varnish systemd service:

sudo systemctl edit --full varnish

Change the port number and cache size in the varnish.service systemd parameters:

ExecStart=/usr/sbin/varnishd \ -a :80 \ -a localhost:8443,PROXY \ -f /etc/varnish/default.vcl \ -P %t/%N/varnishd.pid \ -p feature=+http2 \ -s malloc,2g

This systemd configuration change sets the Varnish service to listen on port 80, and allocates 2GB of memory for caching.

Open the VCL file:

sudo vim /etc/varnish/default.vcl

Customize the VCL:

vcl 4.1; backend default { .host = "origin.example.com"; .port = "80"; }

This VCL change configures the backend server for Varnish to forward requests to origin.example.com on port 80.

Restart Varnish to activate the changes:

sudo systemctl restart varnish

Learn more about installing Varnish on Red Hat Linux and other compatible distributions by reading the install guide.

Install guide

## Getting help

If you need help with Varnish and the docs give you no clue on what to do there are several options available to you.

### Discord Channel

While newer and less populated than the IRC channel, the Varnish discord server is also a good place to ask questions about configuration, HTTP, vmods or anything that’s even remotely related to Varnish.

### StackOverflow

You can ask questions and find answers about Varnish and VCL on StackOverflow using the varnish or the varnish-vcl tag.
