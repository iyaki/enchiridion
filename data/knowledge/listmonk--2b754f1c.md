---
title: "Listmonk"
notion_id: 2b754f1c-7d23-8174-bb1e-c8a87750f4fc
notion_url: https://app.notion.com/p/Listmonk-2b754f1c7d238174bb1ec8a87750f4fc
last_edited: 2025-11-26T19:06:00.000Z
source_url: https://www.ssp.sh/brain/listmonk/
tags: ["Tool", "Article", "FreeCodeCamp", "English", "DevOps", "Automation", "Email", "Open Source", "Self-hosted"]
---
Self-hosted newsletter and mailing list manager. Performance and features packed into a single binary. Free and open source.

My new [Newsletter](https://www.ssp.sh/brain/newsletter) is published with Listmonk on [list.ssp.sh](https://list.ssp.sh/subscription/form). So far I used Mailchimp. I also liked Buttondown, unfortunately, it costs 9$ for above 100 subscribers.

By using open source Listmonk, I have the sending emails and collecting and unsubscribing from a newsletter, which I needed. But I had to deploy it myself. I used Railway.app.

## [# Settings](https://www.ssp.sh/brain/listmonk/#settings)

[**# Automatic Confirm Subscription Email by Clicking in Email**](https://www.ssp.sh/brain/listmonk/#automatic-confirm-subscription-email-by-clicking-in-email)

Default behaviour, if you click on the confirm button in your email:

![image](https://www.ssp.sh/brain/img_Listmonk_1744707958247.webp)

It wil infact not confirm yet, you need another click once again `confirm` on the linked page. I believe this is why many have not confirmed my email subscribtion, tough many are subcribing itself.

There is a [workaround](https://github.com/knadh/listmonk/issues/556?ref=ssp.sh#issuecomment-1830369917) for this by pasting the following piece of JavaScript in `Settings -> Appearance -> Public -> Custom Javascript`:

| <br><br>`1` | <br><br>`document.querySelector(".optin-form").submit();` |
| --- | --- |

After that, the above confirm will open a browser window and automatically click/submit the button after opening the browser. I think this should be default as it won’t work if JavaScript it blocked in your browser.

## [# Automation](https://www.ssp.sh/brain/listmonk/#automation)

[**# Sending Emails of Newsletter with Amazon SES**](https://www.ssp.sh/brain/listmonk/#sending-emails-of-newsletter-with-amazon-ses)

Sending it with Amazon SES.

Make sure you set up and verify the domain, and set some other settings like SPF, DKIM, and DMARC records. Check Bulk or mass email service for Newsletter and Email Automation.

I think I used these checklists:

- Tutorial: [Setting up DMARC, SPF, DKIM with Amazon SES - EmailOctopus knowledge base](https://help.emailoctopus.com/article/68-setting-up-dmarc-spf-dkim?ref=ssp.sh) 

### [# RSS](https://www.ssp.sh/brain/listmonk/#rss)

Based on RSS new email - Provide an RSS integration for listmonk, a self-hosted newsletter and mailing list manager. [GitHub - listmonk-rss:](https://github.com/ping13/listmonk-rss?ref=ssp.sh) by Stephan Heuel.

I just enabled [automatic email send](https://github.com/sspaeti/listmonk-rss/actions?ref=ssp.sh) when I create a new article based on my [RSS-Feed](https://www.ssp.sh/index.xml). I used Above repo by Stephan and [forked it](https://github.com/sspaeti/listmonk-rss?ref=ssp.sh). Locally I can run `make dry_run` or live:

| <br><br>`1<br>2` | <br><br>`cd ~/Documents/git/sspaeti.com/listmonk-rss<br>make create_campaign` |
| --- | --- |

It works like a charm. This will schedule a campaign with the latest articles since `LAST_UPDATE` with a delay of 180 minutes.

### [# This is how it looks](https://www.ssp.sh/brain/listmonk/#this-is-how-it-looks)

Make command:

![image](https://www.ssp.sh/brain/img_Listmonk_1744101659613.webp)

New created campain on Listmonk:

![image](https://www.ssp.sh/brain/img_Listmonk_1744101663722.webp)

![image](https://www.ssp.sh/brain/img_Listmonk_1744101710902.webp)

And updated `LAST_UPDATE` on GitHub:

![image](https://www.ssp.sh/brain/img_Listmonk_1744101654460.webp)

And the Email:

![image](https://www.ssp.sh/brain/img_Listmonk_1744101979218.webp)

Wdyt? :)

### [# Welcome Emails](https://www.ssp.sh/brain/listmonk/#welcome-emails)

Listmonk does not support automatically sending welcome emails or other messages. It is better to use Email Automation, something like Mautic, as mentioned in the [GitHub issue](https://github.com/knadh/listmonk/issues/1391?ref=ssp.sh).

A hacky solution with Listmonk on [Auto reply on subscription (Issue #206)](https://github.com/knadh/listmonk/issues/206?ref=ssp.sh#issuecomment-2423559646).

The alternative I use for now is PayHip, see on [My Services](https://ssp.sh/services/).

### [# Opt-In Campain to none-confirmed subscribers](https://www.ssp.sh/brain/listmonk/#opt-in-campain-to-none-confirmed-subscribers)

![image](https://www.ssp.sh/brain/img_Listmonk_1747224341392.webp)

## [# Tech Setup and Maintenance](https://www.ssp.sh/brain/listmonk/#tech-setup-and-maintenance)

[**# Upgrade Listmonk on Railapp**](https://www.ssp.sh/brain/listmonk/#upgrade-listmonk-on-railapp)

- Head to your dashboard, and select your Listmonk project.
- Select the GitHub deployment service.
- In the Deployment tab, head to the latest deployment, click on the three vertical dots to the right, and select “Redeploy”.

![image](https://user-images.githubusercontent.com/55474996/226517149-6dc512d5-f862-46f7-a57d-5e55b781ff53.png)

Railway Redeploy option

see more on [Upgrade - listmonk / Documentation](https://listmonk.app/docs/upgrade/?ref=ssp.sh)

### [# Backup Postgres DB](https://www.ssp.sh/brain/listmonk/#backup-postgres-db)

Our data should is now being automatically backed up to our S3 bucket. If we ever want to restore this data to a database, we can do so very easily using `pg_restore`.

First, unzip the backup which is stored as a `tar.gz` file and then run:

In the above snippet, replace the `DATABASE_URL` with the connection URL of the database you’d like to restore the data to and the `BACKUP_FOLDER` with the name of the folder to which you extracted your backup zip file and voila, your backup has been restored!

Origin: [Newsletter](https://www.ssp.sh/brain/newsletter)

References: [listmonk - Free and open source self-hosted newsletter, mailing list manager, and transactional mails](https://listmonk.app/?ref=ssp.sh), Send Checklist for Reaching Audience

Created 2023-04-09
