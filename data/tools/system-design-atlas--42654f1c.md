---
title: "System Design Atlas"
notion_id: 42654f1c-7d23-81f8-9f43-cbd25e87b638
notion_url: https://app.notion.com/p/System-Design-Atlas-42654f1c7d2381f89f43cbd25e87b638
last_edited: 2026-11-28T06:15:00.000Z
source_url: https://systemdesignatlas.dev/
tags: ["English", "System Design / Software Architecture", "Databases", "APIs", "Performance", "Distributed Systems", "Algorithms", "Tool", "Guide", "Official Website"]
---
## Learn system design the way the interview tests it.

Concepts built from the ground up, the technologies real designs name, and worked designs that open with the hard part and end with the follow-up questions. Plus a coding track that draws every algorithm as it runs.

Most of the atlas is free to read, with no account.

L2 · batched sync
L2 · batched sync
L2 · batched sync
allow
deny
unreachable → FAIL OPEN
on local bucket only,
log and alert
86%DatabaseCacheExternal systemAPI gatewayLoad balancerFocusClick a node for details · ⌘/Ctrl + scroll to zoom

The live architecture from [Distributed Rate Limiter](https://systemdesignatlas.dev/docs/04-rate-limiter). Every design has one: click a box for its role and its trade-off.

- [11ConceptsThe ideas, in reading order](https://systemdesignatlas.dev/#concepts)
- [25DesignsWorked problems, hard part first](https://systemdesignatlas.dev/#designs)
- [16TechnologiesWhat each system is for](https://systemdesignatlas.dev/#technologies)
- [6OutagesReal incidents, retold](https://systemdesignatlas.dev/#outages)
- [13CodingAlgorithms, drawn as they run](https://systemdesignatlas.dev/#coding)
- [3Courses28 lessons, read in order](https://systemdesignatlas.dev/#courses)

### 22 more designs

- [Chat / Slack](https://systemdesignatlas.dev/docs/02-chat-slack)
- [URL Shortener](https://systemdesignatlas.dev/docs/03-url-shortener)
- [Distributed Rate Limiter](https://systemdesignatlas.dev/docs/04-rate-limiter)
- [Notification System](https://systemdesignatlas.dev/docs/05-notification-system)
- [Search and Typeahead](https://systemdesignatlas.dev/docs/06-search-typeahead)
- [Video Streaming](https://systemdesignatlas.dev/docs/08-video-streaming)
- [Web Crawler](https://systemdesignatlas.dev/docs/09-web-crawler)
- [Ticketmaster / Booking](https://systemdesignatlas.dev/docs/11-ticketmaster-booking)
- [Dropbox / File Sync](https://systemdesignatlas.dev/docs/12-dropbox-file-sync)
- [Ad Click Aggregation](https://systemdesignatlas.dev/docs/13-ad-click-aggregation)
- [Distributed Cache](https://systemdesignatlas.dev/docs/14-distributed-cache)
- [Google Docs](https://systemdesignatlas.dev/docs/15-google-docs-collaborative-editing)
- [AI Agent Platform / Lindy(premium)](https://systemdesignatlas.dev/docs/16-ai-agent-platform)
- [Researcher–Participant Chat(premium)](https://systemdesignatlas.dev/docs/17-researcher-participant-chat)
- [AI App Builder / Lovable(premium)](https://systemdesignatlas.dev/docs/18-ai-app-builder-lovable)
- [ChatGPT(premium)](https://systemdesignatlas.dev/docs/19-chatgpt)
- [RL Environment Platform(premium)](https://systemdesignatlas.dev/docs/20-rl-environment-platform)
- [ElevenLabs(premium)](https://systemdesignatlas.dev/docs/21-elevenlabs)
- [Figma(premium)](https://systemdesignatlas.dev/docs/22-figma)
- [Online Chess / Chess.com(premium)](https://systemdesignatlas.dev/docs/23-online-chess)
- [Collaborative Whiteboard / Miro(premium)](https://systemdesignatlas.dev/docs/24-collaborative-whiteboard)
- [Music Streaming / Spotify(premium)](https://systemdesignatlas.dev/docs/25-spotify)

## Real outages, retold as design lessons.

Each case study follows a public postmortem: what users saw, the timeline, the root cause, what made it spread, and the design lessons, each linked to the concept it tests. Then the interview answer it teaches.

[Read AWS us-east-1: the empty DNS record](https://systemdesignatlas.dev/docs/01-aws-dynamodb-dns)

- [AWS us-east-1: the empty DNS recordOctober 2025. A race in DynamoDB's DNS automation emptied its endpoint, and EC2, load balancers and Lambda fell in behind it for 14 hours.](https://systemdesignatlas.dev/docs/01-aws-dynamodb-dns)
- [Cloudflare: the file that doubledNovember 2025. A permissions change doubled a generated config file, and every Cloudflare proxy crashed loading it.](https://systemdesignatlas.dev/docs/02-cloudflare-feature-file)
- [CrowdStrike: field 21 of 20(premium)July 2024. A content update made every online Windows sensor read a field that didn't exist, and 8.5 million machines crashed on boot.](https://systemdesignatlas.dev/docs/03-crowdstrike-channel-file-291)
- [Facebook: the day the DNS servers hid(premium)October 2021. One command cut every Facebook data centre off, and its DNS servers, as designed, withdrew themselves from the internet.](https://systemdesignatlas.dev/docs/04-facebook-bgp-dns)
- [Roblox: 73 hours on one Consul cluster(premium)October 2021. Two bugs in the one Consul cluster everything depended on took Roblox offline for three days.](https://systemdesignatlas.dev/docs/05-roblox-consul)
- [GitHub: 43 seconds, 24 hours(premium)October 2018. A 43-second network cut moved GitHub's databases across the country, and GitHub chose a degraded day over losing writes.](https://systemdesignatlas.dev/docs/06-github-network-partition)

## Algorithms you can watch run.

The coding track opens every data structure with a step-through: the structure on one side, the code highlighting line by line on the other. Press play below.

target = 19find the two indices
nums[lo] + nums[hi]28 pairs to rule out
• lo
• hi
• answer
• dimmed = eliminatedpythonruns with the steps


01 / 9Sorted inputpairs 28
Eight sorted values, and a target of 19. Brute force would test all 28 pairs. The order in this array is what makes a cheaper answer possible — unsorted, none of what follows holds.
1. 
2. 
3. 
4. 
5. 
6. 
7. 
8. 
9.

- • lo
- • hi
- • answer
- • dimmed = eliminated

```plain text
1def two_sum_sorted(nums, target):
2    lo, hi = 0, len(nums) - 1
3    while lo < hi:
4        total = nums[lo] + nums[hi]
5        if total == target:
6            return [lo, hi]
7        if total > target:
8            hi -= 1
9        else:
10            lo += 1
11    return []
```

1. 1.
2. 2.
3. 3.
4. 4.
5. 5.
6. 6.
7. 7.
8. 8.
9. 9.

## Go deep on one subject.

Each course is a run of lessons read in order, from the overview to a finished system or a full mock round. Pick a course to see its lessons.

## Unlock every premium page for $19, once.

Lifetime Access opens the premium designs, the Interview Playbook, Building with LLMs, the courses, and Listen and Script on every page, plus every premium page added later. No subscription.

[See what Premium includes](https://systemdesignatlas.dev/premium)
