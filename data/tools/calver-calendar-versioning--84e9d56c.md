---
title: "CalVer - Calendar Versioning"
notion_id: 84e9d56c-5b1c-44ff-9ae5-80fc72b5f93d
notion_url: https://app.notion.com/p/CalVer-Calendar-Versioning-84e9d56c5b1c44ff9ae580fc72b5f93d
last_edited: 2023-01-18T13:57:00.000Z
source_url: https://calver.org/
tags: ["English", "Programming", "System Design / Software Architecture", "Project Management", "Website"]
---
_CalVer is a versioning convention based on your project's release calendar, instead of arbitrary numbers._

**Versioning gets better with time.**

For maintainers, versioning allows us to specify precise dependencies within an ever-expanding ecosystem. For sellers and promoters, a project's version is a dynamic part of a brand. For all of us, versioning lets us reference the past while upgrading to the future.

Different projects use different systems for versioning, but common practices have emerged. For instance, point-separated numbers (e.g., _3.1.4_) are all but given. Another common versioning pattern incorporates a time-based element, usually part of the release date.

This date-based approach has come to be called Calendar Versioning, or **CalVer** for short.

Contents

- [Scheme](https://calver.org/#scheme)
- [Case studies](https://calver.org/#case-studies)
- [Ubuntu](https://calver.org/#ubuntu)
- [Twisted](https://calver.org/#twisted)
- [pytz](https://calver.org/#pytz)
- [Teradata](https://calver.org/#teradata)

There are multiple calendar versioning schemes, long used by projects big and small. Rather than declaring a single scheme to be CalVer, it's important to recognize the practicality of each and [design the scheme](http://sedimental.org/designing_a_version.html) to fit the project. First, the parts of the version:

- **Major** - The first number in the version. 2 and 3 are Python's famous major versions. The major segment is the most common calendar-based component.
- **Minor** - The second number in the version. 7 is the most popular minor version of Python.
- **Micro** - The third and usually final number in the version. Sometimes referred to as the "patch" segment.
- **Modifier** - An optional text tag, such as "dev", "alpha", "beta", "rc1", and so on.

The vast majority of modern version identifiers are composed of two or three numeric segments, plus the optional modifier. Convention suggests that four-numeric-segment versions are discouraged.

As seen in the [case studies](https://calver.org/#case_studies) below, projects have found more than one useful way to leverage dates in their versions. Rather than choose a single scheme, CalVer introduces standard terminology for developers, in addition to the "semantic" versions:

- **`YYYY`** - Full year - 2006, 2016, 2106
- **`YY`** - Short year - 6, 16, 106
- **`0Y`** - Zero-padded year - 06, 16, 106
- **`MM`** - Short month - 1, 2 ... 11, 12
- **`0M`** - Zero-padded month - 01, 02 ... 11, 12
- **`WW`** - Short week (since start of year) - 1, 2, 33, 52
- **`0W`** - Zero-padded week - 01, 02, 33, 52
- **`DD`** - Short day - 1, 2 ... 30, 31
- **`0D`** - Zero-padded day - 01, 02 ... 30, 31

Note that traditional, incremented version numbers are 0-based, whereas date segments are 1-based, and the short and zero-padded years are relative to the year 2000. Also note that usage of weeks is usually mutually exclusive with months/days.

The [Gregorian calendar](https://en.wikipedia.org/wiki/Gregorian_calendar) is assumed, as is the convention of [UTC](https://en.wikipedia.org/wiki/Coordinated_Universal_Time). Technically any calendar can be used, provided projects state which one.

CalVer has quite a few users. These are projects selected for their notability and variety of use cases.

[**Ubuntu**](http://www.ubuntu.com/), one of the most prominent Linux-based operating systems available, uses a three-segment CalVer scheme, with a short year and zero-padded month. It has done so [from the very start](https://en.wikipedia.org/wiki/List_of_Ubuntu_releases), in October 2004, making 4.10 the first general release of Ubuntu.

Even a simple operating system involves many, many parts, making it difficult to communicate much meaning with an arbitrary number. By dating the project release, the calendar-based version is much more than an arbitrary number, communicating useful information that is rooted in simple fact.

Ubuntu derives additional benefit from its CalVer scheme, by integrating it with their support schedule. Ubuntu currently has five-year support periods for their long-term support (LTS) releases, and only 9 months for non-LTS releases. Thanks to CalVer and elementary arithmetic, any user can easily determine whether their version is still supported. The current LTS release at the time of writing, 16.04, will be supported until April 2021.

[**Twisted**](https://twistedmatrix.com/), the venerated Python networking and asynchronous execution framework, uses a three-segment CalVer scheme, with a short year in the major version slot, short month in the minor version slot, and micro/patch version in the third and final slot.

First released in 2002 and still actively developed today, Twisted is a [mature](https://en.wikipedia.org/wiki/Twisted_%28software%29) library that has grown to match its large scope. It features everything from an IRC client to an HTTP server to a slew of utilities for concurrent programming. Like an operating system, Twisted has a lot of parts, making SemVer a poor fit due to the individual parts deprecating and breaking compatibility individually.

The non-deprecated parts of Twisted are backwards-compatible between each successive version, and breaking changes are done on a time basis, where one year must pass and two releases issued between the release deprecating the functionality and the removal of the functionality.

Its versioning scheme has spread to related projects, including [Klein](https://github.com/twisted/klein), [Treq](https://github.com/twisted/treq), and even one of Twisted's dependencies, [PyOpenSSL](https://github.com/pyca/pyopenssl).

[**youtube_dl**](https://youtube-dl.org/), the understated ally of Internet media archivists everywhere, uses a three-segment CalVer scheme, including full year, zero-padded month, and zero-padded day. The version is almost completely calendar-driven, except for a micro segment that is added in some technical contexts.

Despite the name, youtube_dl's scope is expansive. It supports extracting audio and video from a long, ever-expanding list of sites. Consider the rapid release cycle of supported services, and it becomes clear why the project has adopted CalVer to such a great degree.

[**pytz**](https://pypi.python.org/pypi/pytz) is the Python translation of the [IANA/Olson timezone database](https://www.iana.org/time-zones), the database behind accurate times for all of computerdom. pytz uses a two-segment CalVer scheme, including full year and minor version.

While Python has a history of "batteries-included" architecture, and the datetime module frequently mentions timezones, the core Python runtime does not include timezone information. This is because timezone updates do not follow a fixed schedule, and are subject to politics and legislative whim. Calendar versioning offers a date-stamped snapshot of an otherwise chaotic system.

The [**Teradata UDA client**](https://pypi.python.org/pypi/teradata) provides [next-generation access](https://developer.teradata.com/tools/reference/teradata-python-module) to [Teradata](http://www.teradata.com/)'s data warehousing technologies.

Teradata's usage is notable not for the prominence of the technology or company, but because there have been multiple releases in 2016 which were versioned as `15.10`. This may seem breaking at first, but the meaning and utility is clear.

The library maintainers have crafted a resourceful hybrid of [semantic versioning](http://semver.org/) and calendar versioning. The **YY.MM** part of the version are used as a combined SemVer major version. That is, for new releases, the API of the library remains the same as it did in October 2015. Dependent code written since then is safe to upgrade. We will see the year and month segments update next time there is a breaking API change.

- [Unity](https://unity3d.com/unity/whats-new/) - **`YYYY.MINOR.MICRO`** - Cross-platform game engine.
- [pip](https://pip.pypa.io/en/stable/news/) - **`YY.MINOR.MICRO`** - Official package manager for Python.
- [PyCharm](https://www.jetbrains.com/pycharm/download/) - **`YYYY.MINOR.MICRO`** - A leading Python IDE.
- [OpenSCAD](http://www.openscad.org/) - **`YYYY.0M`** - The premiere open-source offering for solid 3D CAD modelling.
- [fusefs-ntfs](http://www.freshports.org/sysutils/fusefs-ntfs) - **`YYYY.MM.DD_MICRO`** - One of the earliest and most cross-compatible NTFS access layers for Unix systems.
- [certifi](https://pypi.python.org/pypi/certifi) - **`YYYY.MM.DD`** - certifi is a wrapper around Mozilla's certificate authority bundle, used for secure Internet communication. Similar to [pytz](https://calver.org/#pytz), certificate updates do not follow a fixed schedule, but timely, dateable updates are critical to security.
- [boltons](http://boltons.readthedocs.io/en/latest/) - **`YY.MINOR.MICRO`** - A broad library of utilities supplementing the Python standard library.

See the [Users page](https://calver.org/users.html) for a growing list of CalVer users.

If both you and people you don't know use your project seriously, then use a serious version. Luckily, the decision on whether to use CalVer for that version is easier than ever:

- Does your project feature a large or constantly-changing scope?
- Is your project time-sensitive in any way? Do other external changes drive new project releases?

If you answered yes to any of these questions, CalVer's semantics make it a strong choice for your project.
