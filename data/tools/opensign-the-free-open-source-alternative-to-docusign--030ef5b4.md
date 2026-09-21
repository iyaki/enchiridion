---
title: "OpenSign - The free & Open Source Alternative to DocuSign"
notion_id: 030ef5b4-85a2-4acc-980c-03dda2c54c0c
notion_url: https://app.notion.com/p/OpenSign-The-free-Open-Source-Alternative-to-DocuSign-030ef5b485a24acc980c03dda2c54c0c
last_edited: 2023-10-31T13:32:00.000Z
source_url: https://github.com/OpenSignLabs/OpenSign
tags: ["Tool", "Service", "English", "Office", "Untried"]
---
### Table of Contents

1. Introduction
2. Features
3. Installation
4. Usage
5. Contribution Guidelines
6. License
7. Acknowledgments

Please star ⭐ the repo to support us! 😀

### Introduction

Welcome to OpenSign, the premier open source docusign alternative - document e-signing solution designed to provide a secure, reliable and free alternative to commercial esign platforms like DocuSign, PandaDoc, SignNow, Adobe Sign, Smartwaiver, SignRequest, HelloSign & Zoho sign. Our mission is to democratize the document signing process, making it accessible and straightforward for everyone.

### Features

- Secure PDF E-Signing: With the help of robust encryption algorithms, OpenSign™ ensures maximum security, privacy & compatibility. Now sign unlimited documents even on the cloud hosted free version of OpenSign.
- Annotate Documents: OpenSign™ allows you to annotate PDF documents with an advanced signing pad that allows hand drawn signatures, uploaded images, typed signatures & saved signatures for the simplest open source document signing experience ever.
- User-Friendly Interface: OpenSign™ was built while keeping Intuitive design in mind for ease of use. Features like "Sign yourself", "Templates", "One click signatures" and "OpenSign™ Drive" makes it stand out of the crowd and even makes it better than a lot of so-called industry leaders. OpenSign intends to provide the best document signing experience in the open source ecosystem.
- Multi-signer Support: OpenSign's ability to invite multiple signers for signing along with the ability to invite by sharing signing links & being able to enforce signing in a sequence makes it the only open source solution that is fully loaded and allows it to compete head-to-head with established players in e-signature space.
- Email Unique Code(OTP) verification support for guest signers: With OpenSign™, your documents are fully secure even when being signed by guest users. Guest signers can only sign the document after entering a unique code sent to their email address.
- "Expiring Docs" & "Rejection": You can set documents to expire after certain number of days after which nobody will be able to sign. Not just this, OpenSign™ also allows signers to reject signing a document with a reason that will be promptly shared with the sender.
- Beautiful email templates: All document signing invitations, completion notifications & reminders are formatted using great looking email templates. Not just this, you are even allowed to customise the email templates making your free document signing invitations look the way you always wanted them to be.
- PDF Template Creation: OpenSign™ allows you to create and store PDF document templates for repeated use thereby saving you a lot of time & collect e-signatures seamlessly.
- OpenSign™ Drive: It is a centralised secure vault for your digital documents that makes storing, signing, organizing, sharing & archieving your docs a breeze.
- Audit Trails & completion certificate: Being a security focused solution, OpenSign™ makes it a top priority to save detailed logs for tracking document activities along with time-stamps, IP addresses, email IDs & phone numbers. A completion certificate is generated as soon as document is completed which contains all the document related logs for added safety.
- API Support: OpenSign™ API allows seamless integration into existing systems and software. You can generate an API key from the app and refer the official API docs to start integrating it in your existing applications.
- Integrations: The open source document signing experience becomes even more seamless because of integrations with various Cloud storage systems, CRMs & enterprise platforms. We also have a Zapier integration that allows you to integrate it with virtually any application.

### Deploy

Note: The default MongoDB instance used in deployment is not persistant and will be cleared on every restart. To retain your data, configure and supply your own MongoDB connection URL.

### DigitalOcean

### Docker

The simplest way to install OpenSign on your own server is using official docker images by running the following command -

Command for linux/MacOS

```plain text
export HOST_URL=https://opensign.yourdomain.com && curl --remote-name-all https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/docker-compose.yml https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/Caddyfile https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/.env.local_dev && mv .env.local_dev .env.prod && docker compose up --force-recreate
```

Command for Windows (Powershell)

```plain text
$env:HOST_URL="https://opensign.yourdomain.com"; Invoke-WebRequest -Uri https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/docker-compose.yml -OutFile docker-compose.yml; Invoke-WebRequest -Uri https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/Caddyfile -OutFile Caddyfile; Invoke-WebRequest -Uri https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/.env.local_dev -OutFile .env.local_dev; Rename-Item -Path .env.local_dev -NewName .env.prod; docker compose up --force-recreate
```

Command for Windows (CMD/Terminal)

```plain text
set HOST_URL=https://opensign.yourdomain.com && curl -O https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/docker-compose.yml && curl -O https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/Caddyfile && curl -O https://raw.githubusercontent.com/OpenSignLabs/OpenSign/main/.env.local_dev && rename .env.local_dev .env.prod && docker compose up --force-recreate
```

Make sure that you have Docker and git installed before you run this command -

Please refer to the Installation Guide for detailed instructions on how to install OpenSign on your system.

### Usage

For comprehensive guidelines on how to use OpenSign™, please consult our User Manual.

### Contribution Guidelines

We welcome contributions from the open-source community. For more information on how to contribute, please read our Contribution Guidelines.

### License

OpenSign is licensed under the AGPL-3 License. For more details, see the LICENSE file.

### Acknowledgments

We would like to thank all our contributors and users for their support and feedback. Special thanks to OpenSignLabs for spearheading this initiative.

## Contributors

Aleksandar Jakovljevic💻 Priyanshu Dwivedi💻 Akriti Sengar💻 Parth Chawande💻 Rishabh Dewangan💻 Nitin Mishra💻 Jobin Selvanose📖 Hans Fraiponts📖 Monil Prajapati💻 Edogbanya Emmanuel🐛 pranav514💻 Aria💻 Soumyadipto Pal💻 Andrey Didenko💻 VishakhaSainani💻 Andrew💻 Rishab💻 Maurizio Pillitu🐛 Luis Parra️️️️♿️ Govinda Kocharekar💻 Bilal Ahmad Bhat💻 Vikram💻 ugoconsonni💻 Daniel Mutwiri💻 Zathiel💻 1024mb🌍 Eugene Howe💻

This project is tested with BrowserStack.
