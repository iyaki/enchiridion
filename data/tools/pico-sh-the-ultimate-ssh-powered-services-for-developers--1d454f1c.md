---
title: "pico.sh - The ultimate ssh powered services for developers"
notion_id: 1d454f1c-7d23-8172-927b-e5fbed24e217
notion_url: https://app.notion.com/p/pico-sh-The-ultimate-ssh-powered-services-for-developers-1d454f1c7d238172927be5fbed24e217
last_edited: 2025-06-22T02:03:00.000Z
source_url: https://pico.sh/
tags: ["English", "Hosting", "Producer (Individual Contributor)", "Shell/Bash", "Service"]
---
The ultimate `ssh` powered services for developers

[GET STARTED](https://pico.sh/getting-started)

Our mission is to enable developers with services that help them rapidly prototype on the web. We want to make it easier than ever for developers to share their projects with the world.

Our services allow users to publish content without needing to install anything. We accomplish this with the SSH tools (`rsync`, `sftp`, `sshfs`) you already have installed on your system.

Use our platform entirely using SSH and our TUI.

Read about what motivates us: [RFC-001 radical experimentation](https://blog.pico.sh/rfc-001-radical-experimentation)

### [pages](https://pico.sh/pgs)

Host static sites on our global platform using SSH.

### [tuns](https://pico.sh/tuns)

Host public web services on localhost using SSH.

### [pipe](https://pico.sh/pipe)

Stream data between computers using our authenticated *nix pipes using SSH.

### [prose](https://pico.sh/prose)

Serve your blog using SSH.

### [rss-to-email](https://pico.sh/feeds)

Receive email digests for your RSS feeds using SSH.

### [pastes](https://pico.sh/pastes)

Upload code snippets using `rsync`, `scp`, and `sftp`.

### Deploy a site with a single command

Upload your static site to us:

```plain text
rsync --delete -rv ./public/ pgs.sh:/mysite/
```

Now your site is available with TLS handled for you: **https://{user}-mysite.pgs.sh**

We also automatically handle TLS for your custom domains!

### Access localhost using https

If you have a local webserver on `localhost:8000`, activate an SSH tunnel to us:

```plain text
ssh -R dev:80:localhost:8000 tuns.sh
```

Now your local dev server is available on the web: **https://{user}-dev.tuns.sh**

### Authenticated *nix pipes over SSH

Have one terminal listen for an event and another terminal send the event:

```plain text
ssh pipe.pico.sh sub mytopic
```

```plain text
echo "Hello world!" | ssh pipe.pico.sh pub mytopic
```

The `sub` will receive "Hello world!"

### Publish blog articles with a single command

Create your first post, (e.g. `hello-world.md`):

```plain text
# hello world!

This is my first blog post.

Cya!
```

Upload the post to us:

```plain text
scp hello-world.md prose.sh:/
```

Congrats! You just published a blog article, accessible here: **https://{user}.prose.sh/hello-world**

### Easily share code snippets

Pipe some stdout to us:

```plain text
git diff | ssh pastes.sh changes.patch
```

And instantly share your code snippets: **https://{user}.pastes.sh/changes.patch**

### Receive email notifications for your favorite rss feeds

Create a `blogs.txt` file:

```plain text
=: email rss@myemail.com
=: digest_interval 1day
=> https://pico.prose.sh/rss
=> https://erock.prose.sh/rss
```

Then upload it to us:

```plain text
scp blogs.txt feeds.pico.sh:/
```

After the daily interval has been reached, you will receive an email with your feeds!

### Ready to join?

[GET STARTED](https://pico.sh/getting-started)

Built by **pico.sh LLC**

206 E Huron St, Ann Arbor MI 48104
