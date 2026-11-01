---
title: "Haskell for all: NixOS in production"
notion_id: 3f454f1c-7d23-815b-bc34-c19ba251cf07
notion_url: https://app.notion.com/p/Haskell-for-all-NixOS-in-production-3f454f1c7d23815bbc34c19ba251cf07
last_edited: 2026-10-09T04:31:00.000Z
source_url: https://haskellforall.com/2018/08/nixos-in-production
tags: ["English", "DevOps", "NixOS", "Linux", "Deployment", "Automation", "Terraform", "Article", "Tutorial", "Haskell for all"]
---
This is a short post summarizing what I wish I had known when I first started using NixOS in production. Hopefully other people will find this helpful to avoid the pitfalls that I ran into.

The main issue with NixOS is that [the manual](https://nixos.org/nixos/manual/) recommends a workflow that is not suitable for deployment to production. Specifically, the manual encourages you to:

- keep Nix source code on the destination machine
- build on the destination machine
- use Nix’s channel system to obtain `nixpkgs` code

This post will describe how you can instead:

- build a source-free binary closure on a build machine
- transfer and deploy that binary closure to a separate destination machine

This guide overlaps substantially with what the NixOps tool does and you can think of this guide as a way to transplant a limited subset of what NixOps does to work with other provisioning tools (such as Terraform).

You might also find this guide useful even when using NixOS as a desktop operating system for handling more exotic scenarios not covered by the `nixos-rebuild` command-line tool.

## Building the closure

We’ll build up to the final solution by slowly changing the workflow recommended in the NixOS manual.

Suppose that you already have an `/etc/nixos/configuration.nix` file and you use `nixos-rebuild switch` to deploy your system. You can wean yourself off of `nixos-rebuild` by building the binary closure for the system yourself. In other words, you can reimplement the `nixos-rebuild build` command from scratch.

To do so, create the following file anywhere on your system:

```plain text
# ./nixos.nix

import <nixpkgs/nixos> {
system = "x86_64-linux";

configuration = {
imports = [
/etc/nixos/configuration.nix
];
};
}
```

Then run:

```plain text
$ nix-build --attr system ./nixos.nix
```

Congratulations, you’ve just built a binary closure for your system!

## Deploying the closure

`nix-build` deposits a `result` symlink in the current directory pointing to the built binary closure. To deploy the current system to match the built closure, run:

```plain text
$ sudo ./result/bin/switch-to-configuration switch
```

Congratulations, you’ve just done the equivalent of `nixos-rebuild switch`!

As the above command suggests, the closure contains a `./bin/switch-to-configuration` which understands a subset of the commands that the `nixos-rebuild` command does. In particular, the `switch-to-configuration` script accepts these commands:

```plain text
$ ./result/bin/switch-to-configuration --help
Usage: ../../result/bin/switch-to-configuration [switch|boot|test]
switch:       make the configuration the boot default and activate now
boot:         make the configuration the boot default
test:         activate the configuration, but don't make it the boot default
dry-activate: show what would be done if this configuration were activated
```

## Adding a profile

The `nixos-rebuild` command actually does one more thing in addition to buiding the binary closure and deploying the system. The `nixos-rebuild` command also creates a symlink pointing to the current system configuration so that you can roll back to that configuration later. The symlink also acts like a garbage collection root, preventing the system from being garbage collected until you remove the symlink (either directly using `rm` or a higher-level utility such as `nix-collect-garbage`)

You can record the system configuration in the same way as `nixos-rebuild` using the `nix-env` command:

```plain text
$ nix-env --profile /nix/var/nix/profiles/system --set ./result
```

## Querying system options

You can use the same `nixos.nix` file to query what options you’ve set for your system, just like the `nixos-option` utility. For example, if you want to compute the final value of the `networking.firewall.allowedTCPPorts` option then you run this command:

```plain text
$ nix-instantiate --eval --strict --attr config.networking.firewall.allowedTCPPorts nixos.nix
[ 22 ]
```

## Pinning `nixpkgs`

Now that you’ve taken control of the build you can do fancier things like pin `nixpkgs` to a specific revision of `nixpkgs` so that you don’t need to use a channels or the `NIX_PATH` any longer:

```plain text
let
nixpkgs = builtins.fetchTarball {
url = "https://github.com/NixOS/nixpkgs/archive/b74b1cdb2fecc31ff7a127c5bc89771f887c93bb.tar.gz";
sha256 = "0ncr4g29220amqm4riaa1xf4jz55v2nmh9fi16f1gzhww1gplk8h";
};

in
import "${nixpkgs}/nixos" {
configuration = {
imports = [
/etc/nixos/configuration.nix
];
};

system = "x86_64-linux";
}
```

In fact, this makes your build completely insensitive to the `NIX_PATH`, eliminating a potential source of non-determinism from the build.

## Building remotely

Now that you’ve removed `nixos-rebuild` from the equation you can build the binary closure on a separate machine from the one that you deploy to. You can check your `nixos.nix`, `configuration.nix` and `hardware-configuration.nix` files into version control and `nix-build` the system on any machine that can check out your version controlled Nix configuration. All you have to do is change the import path to be a relative path to the `configuration.nix` file within the same repository instead of an absolute path:

```plain text
# ./nixos.nix

let
nixpkgs = builtins.fetchTarball {
url = "https://github.com/NixOS/nixpkgs/archive/b74b1cdb2fecc31ff7a127c5bc89771f887c93bb.tar.gz";
sha256 = "0ncr4g29220amqm4riaa1xf4jz55v2nmh9fi16f1gzhww1gplk8h";
};

in
import "${nixpkgs}/nixos" {
configuration = {
imports = [
./relative/path/to/configuration.nix
];
};

system = "x86_64-linux";
}
```

Then all your build machine has to do is:

```plain text
$ git clone "git@github.com:${OWNER}/${REPOSITORY}.git"
$ nix-build --attr system "./${REPOSITORY}/path/to/nixos.nix"
```

To deploy the built binary closure to another machine, use `nix copy`. If you have SSH access to the destination machine then you can `nix copy` directly to that machine:

```plain text
$ nix copy --to "${USER}@${MACHINE}" ./result
```

If you do not have SSH access to the destination machine you can instead use `nix-copy` to create a binary archive file containing the binary closure of your system:

```plain text
$ # On the build machine
$ nix copy --to 'file:///tmp/system' ./result
```

… and upload the binary archive located at `/tmp/system` to the destination machine using your upload method of choice. Then import the binary archive into the `/nix/store` on the destination machine using `nix copy`:

```plain text
$ # On the destination machine
$ nix copy --from 'file:///tmp/system' ./result
```

Once the binary closure is on the machine, you install the closure the same way as before:

```plain text
$ sudo /nix/store/.../bin/switch-to-configuration switch
```

… replacing `/nix/store/...` with the `/nix/store` path of your closure (since there is no `result` symlink on the destination machine).

## Conclusion

That’s it! Now you should be able to store your NixOS configuration in version control, build a binary closure as part of continuous integration, and deploy that binary closure to a separate destination machine. You can also now pin your build to a specific revision of Nixpkgs so that your build is more deterministic.

I wanted to credit my teammate [Parnell Springmeyer](https://twitter.com/digitalmentat) who taught me the `./result/bin/switch-to-configuration` trick for deploying a NixOS system and who codified the trick into the [`nix-deploy`](https://github.com/awakesecurity/nix-deploy) command-line tool. I also wanted to credit [Remy Goldschmidt](https://twitter.com/taktoa1) who interned on our team over the previous summer and taught me how to reliably pin Nixpkgs.
