---
title: "PHPat - PHP Architecture Tester"
notion_id: db7ea7e7-d8af-4352-a960-b7f25ec7fe50
notion_url: https://app.notion.com/p/PHPat-PHP-Architecture-Tester-db7ea7e7d8af4352a960b7f25ec7fe50
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://github.com/carlosas/phpat#introduction-
tags: ["Tool", "English", "System Design / Software Architecture", "PHP", "Untried"]
---
[https://github.com/carlosas/phpat#introduction-](https://github.com/carlosas/phpat#introduction-)

## Easy to use architecture testing tool for PHP

### Introduction 📜

PHP Architecture Tester is a PHPStan extension (Static Analysis tool) designed to verify architectural requirements.

It provides a natural language abstraction that enables you to define your own architectural rules and and assess their compliance in your code.

### Getting started 🚀

Require PHPat with Composer:

```plain text
composer require --dev phpat/phpat
```

Activate the extension using one of the following methods:

Automatic activation

```plain text
composer require --dev phpstan/extension-installer
```

Manual activation

```plain text
# phpstan.neon includes: - vendor/phpat/phpat/extension.neon
```

For further information, check out the documentation at phpat.dev

You can visit the Examples section to get some ideas of typical use cases.

PHP Architecture Tester is open source, contributions are welcome!

> Warning The launch of early-stage releases (0.x.x) could break the API according to Semantic Versioning 2.0. We are using minor for breaking changes until the release of the stable 1.0.0 version.

### Sponsors 💙
