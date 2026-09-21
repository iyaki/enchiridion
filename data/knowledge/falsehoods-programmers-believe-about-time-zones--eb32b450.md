---
title: "Falsehoods programmers believe about time zones"
notion_id: eb32b450-9301-4acc-8ba6-7b502ffa2b4d
notion_url: https://app.notion.com/p/Falsehoods-programmers-believe-about-time-zones-eb32b45093014acc8ba67b502ffa2b4d
last_edited: 2024-03-27T19:55:00.000Z
source_url: https://www.zainrizvi.io/blog/falsehoods-programmers-believe-about-time-zones/
tags: ["Article", "ZAIN RIZVI Blog", "English", "Programming", "System Design / Software Architecture"]
---
My aunt has a problem

She loves joining Zoom meetings, but they're all hosted in different time zones. It's hard to remember if she should add 4 hours, subtract 3, or what. She's not the most technical person, so google isn't an option. She has to ask for help.

Every. Single. Time.

And, for the less technically minded, it's also error-prone.

It got me thinking:

What if event organizers could share a link that would do the work for you? If someone clicked on [mytime.at/5pm/EST](https://www.zainrizvi.io/blog/falsehoods-programmers-believe-about-time-zones/mytime.io/5pm/EST), they would see their local version of that time. It sounded simple enough.

I began coding.

![image](https://www.zainrizvi.io/content/images/2020/10/image-5.png)

I later discovered mytime.io had already implemented a very similar thing, and run into the same pitfalls

I knew trying to manage time [is a fool's errand](https://www.youtube.com/watch?v=-5wpm-gesOY&ref=zainrizvi.io), but that's what datetime libraries are for. I would merely build an extra time zone conversion layer on top.

Surely that couldn't be complicated

...Right?

I soon discovered just how wrong I was. One after another, I kept learning the falsehood of yet another "fact" that had seemed obviously true. Eventually my original vision became literally impossible to pull off without making serious compromises (more about that in a future blog post).

Hopefully this list will help you avoid the landmines I stepped on. All the falsehoods below are ones I'd considered true at some point in my adult life.

Most of them I believed just one month ago.

### **Misconception #1: UTC offsets go from -12 to +12**

Turns out, UTC offsets span from -12 to +14. Yeah, +14. That's gives you 27 hours UTC can be offset by (don't forget the zero offset)

How does it work? UTC-12 has the same time as UTC+12, but is one day behind. Same goes for UTC-11 and UTC+13, etc.

Why that crazy range? That was a result of pacific islanders decided they wanted to be on a specific side of the international date line.

It makes for a very jagged international date line

![image](https://www.zainrizvi.io/content/images/2020/10/pasted-image-0--1-.png)

### Misconception #2: Every UTC offset corresponds to exactly one time zone

Here are 10 distinct time zones which are all at UTC+5:

- Aqtobe Time
- Mawson Time
- Maldives Time
- Oral Time
- Pakistan Standard Time
- French Southern and Antarctic Time
- Tajikistan Time
- Turkmenistan Time
- Uzbekistan Time
- Yekaterinburg Time

You might be wondering: if they’re all at the same UTC offset, why couldn’t all those countries just use the same time zone? Perhaps Pakistanis weren’t keen about being on “Yekaterinburg Time”

### Misconception #3: **There are more countries in the world than time zones**

How could this one possibly be wrong? Well...

1. Many countries want their very own time zone (how many do you think run on Myanmar Time?)
2. Some countries split themselves up into multiple time zones (e.g. eastern and western times)
3. Military time alone uses 25 time zones, one for each hour from UTC-12 to UTC+12
4. DST. More on this one below

All together, there are[ 244 time zones](https://www.timeanddate.com/time/zones/?ref=zainrizvi.io) used by the 195 countries in the world.

### Misconception #4: **Every time zone has exactly one agreed upon name**

Ever notice how every time zone consists only of English words? Awfully kind of Spanish and French speaking countries to graciously use our language, right?

Hah, Yeah right.

Eastern Standard Time, Tiempo del Este, and Heure Normale de l'Est are all different names for the exact same time zone.

Have fun coding that into your library.

### Misconception #5: **Time zones are always offset from UTC by an integer number of hours**

India standard time is five and a half hours off of UTC. There are many more examples

### Misconception #6: **Fine, time zones are always offset from UTC by an integer number of half-hours**

Nepal likes to be at the 45 minute UTC offset.

Why does that extra 15 minutes matter so much to them? Because[ they really want](https://www.bbc.com/news/world-asia-33815153?ref=zainrizvi.io#:~:text=Nepal%20is%205%20hours%20and,a%20mountain%20east%20of%20Kathmandu.&text=It%20gets%20trickier%20in%20the,t%20officially%20have%20time%20zones.) their mountain to have the sun right above it at noon.

But it makes you wonder: what would happen if the mountain ever shifted?

![image](https://www.zainrizvi.io/content/images/2020/10/mountain-moved-3.jpeg)

### Misconception #7: **A country stays at the same UTC offset all year long**

Don't forget about Daylight Saving Time! Or as the Europeans call it "Summer Time."

Countries practicing DST change their UTC offset twice every year.

### Misconception #8: **There is a standard format for declaring time zones**

Hah, I wish. Here are some standards I discovered, there may be more:

### Common name

These are the traditional time zone names we’re used to. Example: Pacific Standard Time.

But I don't know if there's an official term for these names, they just that unstandardized.

### IANA zone keys

This is as close to the official standard as you can get. It's not at all official, but it's something the developer community has rallied around.

It's a [painstakingly maintained database](https://www.iana.org/time-zones?ref=zainrizvi.io) which contains all known time zone data representing the entire history of local time for places around the globe. It doesn't give any zone a name though, preferring to use the name of the most prominent city in there, which leads to:

### Prominent city based

This one is "[basically bad UI that derives from the IANA zone keys](https://twitter.com/pganssle/status/1319794747876253697?ref=zainrizvi.io)"

Full time zone names come with naming complications, which we discussed above. If that wasn't enough fun, there's also the political implications of recognizing certain time zones such as Israel Standard Time.

Some developers took the safer route and identified time zones only by the name of a prominent city in it, not bothering to map it to a common name. That's why the Ubuntu time zone picker makes you select "New York'' instead of Eastern Standard Time.

![image](https://www.zainrizvi.io/content/images/2020/10/image-4.png)

### Forget time zones, use the raw UTC offset

W3's international standard gave up on the notion of time zones and declared that engineers should only store a timestamp's[ raw UTC offset](https://www.w3.org/TR/NOTE-datetime?ref=zainrizvi.io).

### GPS Coordinates

Fun fact: Many APIs for getting a region's UTC offset only want a UTC time and latitude/longitude coordinates. This lets them define any moment unambiguously and not have to worry about Daylight Saving Time.

If you squint your eyes a bit, you could consider this a fourth standard.

### Misconception #9: **Daylight Saving Time starts at the same time every year**

Did you think this would be the one thing world powers agree on? Each country choose when to start it's own DST

### Misconception #10: **A country's time zone never changes**

Almost every year some country will pass a law to edit their time zone.

In a particularly memorable example, a few years ago the Samoan islands wanted to be on the other side of the international date line to get the same weekends as their Australian trading partners. So on midnight Dec 29th, they changed their UTC offset from -11 to +13 UTC, skipping Dec 30th and going straight to Dec 31st.

Samoan citizens had one less day to celebrate the holidays that year.

On the plus side, just 40 miles away the American Samoa Islands stayed on the other side of the international date line. Now Samoans can celebrate new years on the Western Island, and then boat over to American Samoa for a second new year’s party the next night.

![image](https://www.zainrizvi.io/content/images/2020/10/pasted-image-0--2-.png)

### Misconception #11: **A country stays in the same time zone during Daylight Saving Time**

Funny thing about DST, it doesn't actually change the time zone's UTC offset. Instead, Daylight Saving Time countries switch to a different time zone, with a different name.

For example:

Texas goes from Central Standard Time to Central Daylight Time.

Chile goes from Chile Standard Time to Chile Summer Time

![image](https://www.zainrizvi.io/content/images/2020/10/dst-shift.jpeg)

### Misconception #12: **Daylight Saving Time starts around March and ends around October**

The Southern hemisphere has their summer in the other half of the year. The pattern flips.

### Misconception #13: **Every time zone has it’s own name**

Which country should get to claim "Eastern Standard Time"?

North America claimed dibs by virtue of inventing the name, but do you think no one objected? Australia thought it sounded like a fine name to use, and so even though the rest of the world refer to their time zone as Australian Eastern Standard Time, [some](http://disq.us/p/2cqg1bc?ref=zainrizvi.io) of it's own citizens just call it "Eastern Standard Time" (not all of them [call it that though](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9qv6cc/?ref=zainrizvi.io)).

### Misconception #14: **Every time zone has its own abbreviation**

Which of these was meant when someone says CST?

- Central Standard Time
- China Standard Time
- Cuba Standard Time

And remember how the time zone name changes during Daylight Saving Time? Many people don’t know that and keep using the wrong abbreviations during DST months. CST might be used for Central Daylight Time.

If there's no standard name for time zones, can you really expect one for the abbreviations?

![image](https://www.zainrizvi.io/content/images/2020/10/fake-franks.jpeg)

### Misconception #15: **There is always an unambiguous conversion from one time zone to another**

If I say I want to convert 5pm Eastern Standard Time to Pakistan Standard Time, am I talking about the American or Australian Eastern Standard Time?

And is Daylight Saving Time in effect or not?

Okay, it’s tricky. But surely if we include the date and the exact city, then we'd be able to do the conversion reliably, right?

What if the date and time are 1:30 am on Nov 1st, 2020, right when US DST ends and the clock moves backwards?

1:30am occurs twice that morning, how do you know which instance was intended?

![image](https://www.zainrizvi.io/content/images/2020/10/deja-vu.jpeg)

### Misconception #16: **Your time zone library can recognize any time zone (you are using a library for this, right?)**

Remember all those different potential time zone names and formats? Most libraries will only support one.

And they might be limited by the time zones installed on your local machine.

Yeah, really.

Remember, if time zones can change based on the whims of a local government, then the library will need some external dataset to base its calculations off of. That external dataset just might be the time zones installed on your PC.

### Misconception #17: **The entire country always shifts during Daylight Saving Time**

In the US, Arizona doesn't practice Daylight Saving Time

![image](https://www.zainrizvi.io/content/images/2020/10/arizona-dst.jpeg)

### Misconception #18: **The entire state always shifts during Daylight Saving Time**

Within Arizona, the Navajo Nation happily follows Daylight Saving Time

![image](https://www.zainrizvi.io/content/images/2020/10/navajo-nation.jpeg)

### Misconception #19: **Other than DST, every city within a state follows the same time zone**

In Indiana, USA, most cities follow Eastern Standard Time but a few decided to follow Central Standard Time

### Misconception #20: **Every city sits within exactly one time zone**

A few times in history, state line or time zone lies got drawn without paying attention to who actually lived on that border, cutting a city in half. [There are a surprisingly large number of examples](https://www.quora.com/Are-there-any-major-cities-divided-by-two-time-zones?ref=zainrizvi.io).

This enables some really unusual sleep schedules

![image](https://www.zainrizvi.io/content/images/2020/10/Wake-up.jpeg)

It is also why GPS coordinates are more reliable than city names to determine the time zone

### Misconception #21: **There’s a designated time zone for every location in the world**

The north and south poles have no official time zone. Researchers there just follow their own country's time.

There's no way that could get confusing.

### Misconception #22: **This is a comprehensive list of misconceptions**

These are the misconceptions I've uncovered so far, but I'm sure there are many more waiting to be discovered. Heck, I didn't even realize UTC offsets went all the way up to +14 until just 10 hours before I published this list!

### ~~Misconception #23: Redditors will agree these are all misconceptions~~

Who am I kidding? Reddit never agrees on anything. Seriously though, this post sparked [a conversation](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/?ref=zainrizvi.io) which highlighted a few more misconceptions I'd held. Thanks Reddit!

Some of the highlights:

- **#24: Daylight Saving Time starts and stops exactly once each year:** When the month of Ramadan starts, some Muslim countries will exit DST time, and then re-enter DST once Ramadan ends. It makes sunset (the time to end your fast) arrive faster (via [matthieum](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9wf8o3/?context=3&ref=zainrizvi.io))
- **#25: DST offsets are always exactly one hour**: The Lord Howe island uses a 30 minute DST offset (via [paulrpg](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9qr08v/?ref=zainrizvi.io)), and England once had a 2 hour DST offset (via [bandwidthcrisis](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9rsvzq/?ref=zainrizvi.io))
- **#26: Standard Time is the same as Time Zone**: They're [different concepts](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9ql7cl/?ref=zainrizvi.io) (and apparently I've been using them wrong this whole article :P) (via [lpsmith](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9ql7cl/?ref=zainrizvi.io))
- **#27: Timezones are always offset from UTC by an integer number of quarter-hours:** Amsterdam was once at the UTC + 19 minute, 32.13 second offset. Most of the world simplified it to [UTC+0:20](https://en.wikipedia.org/wiki/UTC%2B00:20?ref=zainrizvi.io) (via [DJDavio](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9qwqp6/?ref=zainrizvi.io))
- **#28: Everyone follows their official time zone:** Some western parts of China have their own unofficial time zone (via [jl2352](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9r1wxz/?context=3&ref=zainrizvi.io)), and some industries independently decide to ignore DST to mitigate the timezone madness (via [bitchkat](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9relwp/?ref=zainrizvi.io))
- **#29: You can solve your problems by saving the time as UTC:** Saving future timestamps in UTC can [still lead to confusion](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9r90ni/?context=3&ref=zainrizvi.io) (via [AryA_ch](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9r90ni/?context=3&ref=zainrizvi.io))
- **#30: Birth dates tell you who is older**: [Not necessarily](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9raph0/?context=3&ref=zainrizvi.io) (via [oshkarr](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9raph0/?context=3&ref=zainrizvi.io))
- **#31: There are exactly 195 countries in the world:** Not exactly time zones but another misconception (via [kankyo](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/g9qztxm/?context=3&ref=zainrizvi.io))
- **#32: If you have a UTC timestamp and the GPS coordinates, you can always determine the local time**: Palestine and Israel have different time zones. So in the West Bank, the time zone [depends on if you're Palestinian or an Israeli settler](https://www.972mag.com/the-worlds-only-ethnic-time-zone/?ref=zainrizvi.io). If you don't know which person you're computing the time zone for, the local time is ambiguous (via [haxney](https://www.reddit.com/r/programming/comments/jggx3l/falsehoods_programmers_believe_about_time_zones/gam4vvs/?ref=zainrizvi.io))
- **#33: Historical UTC offsets for a region never change**: The 1927 time zone shift for Shanghai [has been adjusted at least twice](https://stackoverflow.com/questions/6841333/why-is-subtracting-these-two-times-in-1927-giving-a-strange-result?ref=zainrizvi.io) since July 2011 (via [JoeIngeno](https://twitter.com/JoeIngeno/status/1323429822211756034?ref=zainrizvi.io))
- **#34: Daylight Saving Time is the only timezone adjustment: **Sometimes countries instead move their clocks backwards in the Winter and call it [Winter Time](https://en.wikipedia.org/wiki/Winter_time_(clock_lag)?ref=zainrizvi.io) (via [Radek Zajic](https://twitter.com/zajdee/status/1325039075108265985?s=21&ref=zainrizvi.io))

Want to hear about how I designed the time zone project, worked around the obstacles, balanced trade offs, what I finally ended up building?

Sign up below to get that essay once it's ready.
