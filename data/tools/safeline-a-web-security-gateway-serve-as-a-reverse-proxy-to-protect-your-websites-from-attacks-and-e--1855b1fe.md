---
title: "SafeLine - A web security gateway, serve as a reverse proxy to protect your websites from attacks and exploits"
notion_id: 1855b1fe-806c-4a93-9289-1070e55cfc5a
notion_url: https://app.notion.com/p/SafeLine-A-web-security-gateway-serve-as-a-reverse-proxy-to-protect-your-websites-from-attacks-an-1855b1fe806c4a9392891070e55cfc5a
last_edited: 2024-07-15T23:41:00.000Z
source_url: https://github.com/chaitin/SafeLine
tags: ["English", "Information Security", "Network", "Infrastructure", "REST API", "Untried", "Tool"]
---
# SafeLine, the best free WAF for webmaster

![image](https://github.com/chaitin/SafeLine/raw/main/documents/static/images/403.svg)

SafeLine is a web security gateway to protect your websites from attacks and exploits.

It defenses for all of web attacks, such as sql injection, code injection, os command injection, CRLF injection, ldap injection, xpath injection, rce, xss, xxe, ssrf, path traversal, backdoor, bruteforce, http-flood, bot abused and so on.

[🏠Home](https://waf.chaitin.com/) | [📖Documentation](https://docs.waf.chaitin.com/) | [🔍Live Demo](https://demo.waf.chaitin.com:9443/dashboard) | [中文版](https://waf-ce.chaitin.cn/)

# Screenshots

![image](https://github.com/chaitin/SafeLine/raw/main/images/safeline_en.png)

# How It Works

![image](https://github.com/chaitin/SafeLine/raw/main/images/safeline-as-proxy.png)

SafeLine is developed based on nginx, it serves as a reverse proxy middleware to detect and cleans web attacks, its core capabilities include:

- Defenses for web attacks
- Proactive bot abused defense
- HTML & JS code encryption
- IP-based rate limiting
- Web Access Control List

# Installation

**中国大陆用户安装国际版可能会导致无法连接云服务，请查看** [中文版安装文档](https://docs.waf-ce.chaitin.cn/zh/%E4%B8%8A%E6%89%8B%E6%8C%87%E5%8D%97/%E5%AE%89%E8%A3%85%E9%9B%B7%E6%B1%A0)

## Automatic Deploy

> 

👍Recommended

Use the following command to start the automated installation of SafeLine. (This process requires root privileges)

```plain text
bash -c "$(curl -fsSLk https://waf.chaitin.com/release/latest/setup.sh)"
```

After the command is executed, it means the installation is successfully. Please go to "Use Web UI" directly.

## Mannually Deploy

to see [Documentation](https://docs.waf.chaitin.com/en/tutorials/install)

# Usage

## Login

Open the web console page `https://<safeline-ip>:9443/` in the browser, then you will see below.

Execute the following command to get administrator account

![image](https://github.com/chaitin/SafeLine/raw/main/images/login.png)

```plain text
docker exec safeline-mgt /app/mgt-cli reset-admin --once
```

After the command is successfully executed, you will see the following content

> 

Please must remember this content

```plain text
[SafeLine] Initial username：admin
[SafeLine] Initial password：**********
[SafeLine] Done

```

Enter the password in the previous step and you will successfully logged into SafeLine.

## Protecting a website

Log into the SafeLine Web Admin Console, go to the "Site" -> "Website" page and click the "Add Site" button in the upper right corner.

In the next dialog box, enter the information to the original website.

![image](https://github.com/chaitin/SafeLine/raw/main/images/add-site-1.png)

- **Domain**: domain name of your original website, or hostname, or ip address, for example: `www.chaitin.com`
- **Port**: port that SafeLine will listen, such as 80 or 443. (for `https` websites, please check the `SSL` option)
- **Upstream**: real address of your original website, through which SafeLine will forward traffic to it

After completing the above settings, please resolve the domain name you just entered to the IP address of the server where SafeLine is located.

![image](https://github.com/chaitin/SafeLine/raw/main/images/add-site-2.png)

Then you can access the website protected by the SafeLine through the domain name like this.

![image](https://github.com/chaitin/SafeLine/raw/main/images/safeline-as-proxy-2.png)

## Try to attack your website

Now, your website is protected by SafeLine, let’s try tp attack it and see what happens.

If [https://chaitin.com](https://chaitin.com/) is a website protected by SafeLine, here are some test cases for common attacks:

- SQL Injection: `https://chaitin.com/?id=1+and+1=2+union+select+1`
- XSS: `https://chaitin.com/?id=<img+src=x+onerror=alert()>`
- Path Traversal: `https://chaitin.com/?id=../../../../etc/passwd`
- Code Injection: `https://chaitin.com/?id=phpinfo();system('id')`
- XXE: `https://chaitin.com/?id=<?xml+version="1.0"?><!DOCTYPE+foo+SYSTEM+"">`

Replace `chaitin.com` in the above cases with your website domain name and try to access it.

![image](https://github.com/chaitin/SafeLine/raw/main/images/blocked.png)

Check the web console of SafeLine to see the attack list

To view the specific details of the attack, click "detail"

![image](https://github.com/chaitin/SafeLine/raw/main/images/log-list.png)

![image](https://github.com/chaitin/SafeLine/raw/main/images/log-detail.png)

## Star History

![image](https://camo.githubusercontent.com/eec2037d3539d6e2fb0a46d30a81825f28450b1d0dfe1e191a8541732b46cfaf/68747470733a2f2f6170692e737461722d686973746f72792e636f6d2f7376673f7265706f733d6368616974696e2f736166656c696e6526747970653d44617465)

## Related Repo

[Automaton Generator](https://github.com/chaitin/yanshi) | [Lua Plugin](https://github.com/chaitin/safeline-open-platform) | [T1K Protocol](https://github.com/chaitin/lua-resty-t1k) | [WAF Test Tool](https://github.com/chaitin/blazehttp)
