---
title: "johnkerl/miller: Miller is like awk, sed, cut, join, and sort for name-indexed data such as CSV, TSV, and tabular JSON"
notion_id: 2d654f1c-7d23-8164-9c84-d52ab4807273
notion_url: https://app.notion.com/p/johnkerl-miller-Miller-is-like-awk-sed-cut-join-and-sort-for-name-indexed-data-such-as-CSV-TSV-2d654f1c7d2381649c84d52ab4807273
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://github.com/johnkerl/miller
tags: ["English", "Command Line", "Data Manipulation", "CSV", "Tools", "Tool", "Article", "GitHub"]
---
# What is Miller?

**Miller is like awk, sed, cut, join, and sort for data formats such as CSV, TSV, JSON, JSON Lines, and positionally-indexed.**

# What can Miller do for me?

With Miller, you get to use named fields without needing to count positional indices, using familiar formats such as CSV, TSV, JSON, JSON Lines, and positionally-indexed. Then, on the fly, you can add new fields which are functions of existing fields, drop fields, sort, aggregate statistically, pretty-print, and more.

![image](https://github.com/johnkerl/miller/raw/main/docs/src/coverart/cover-combined.png)

- 
- 

In the above image you can see how Miller embraces the common themes of key-value-pair data in a variety of data formats.

# Getting started

- [Miller in 10 minutes](https://miller.readthedocs.io/en/latest/10min)
- [A Guide To Command-Line Data Manipulation](https://www.smashingmagazine.com/2022/12/guide-command-line-data-manipulation-cli-miller)
- [A quick tutorial on Miller](https://www.ict4g.net/adolfo/notes/data-analysis/miller-quick-tutorial.html)
- [Miller Exercises](https://github.com/GuilloteauQ/miller-exercises)
- [Tools to manipulate CSV files from the Command Line](https://www.ict4g.net/adolfo/notes/data-analysis/tools-to-manipulate-csv.html)
- [www.togaware.com/linux/survivor/CSV_Files.html](https://www.togaware.com/linux/survivor/CSV_Files.html)
- [MLR for CSV manipulation](https://guillim.github.io/terminal/2018/06/19/MLR-for-CSV-manipulation.html)
- [Linux Magazine: Process structured text files with Miller](https://www.linux-magazine.com/Issues/2016/187/Miller)
- [Miller: Command Line CSV File Processing](https://onepointzero.app/posts/miller-command-line-csv-file-processing/)
- [Miller - A Swiss Army Chainsaw for CSV Data, Data Science and Data Munging](https://fuzzyblog.io/blog/data_science/2022/05/13/miller-a-swiss-army-chainsaw-for-csv-data-data-science-and-data-munging.html)
- [Pandas Killer: mlr, the Scientist](https://xvzftube.xyz/posts/pandas_killers/#mlr%3A-the-scientist)

# More documentation links

- [**Full documentation**](https://miller.readthedocs.io/)
- [Miller's license is two-clause BSD](https://github.com/johnkerl/miller/blob/main/LICENSE.txt)
- [Notes about issue-labeling in the Github repo](https://github.com/johnkerl/miller/wiki/Issue-labeling)
- [Active issues](https://github.com/johnkerl/miller/issues?q=is%3Aissue%20is%3Aopen%20sort%3Aupdated-desc)

# Installing

There's a good chance you can get Miller pre-built for your system:

![image](https://camo.githubusercontent.com/8828d837fa43cc62ca7bc1675f22d555e3ee28c5980ff3c0d6a7105e91b81057/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d7562756e74752d6462343932332e737667)

![image](https://camo.githubusercontent.com/d0d1b55dfddc2189fad4e8ca78ae76fd60e1e1a8d24b581eedac31f65388e59b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d7562756e7475313630346c74732d6462343932332e737667)

![image](https://camo.githubusercontent.com/3f842e0e1613527d055ad6c96c2438151057785eb13ec7a23bf0c6d9fd1de6a2/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d6665646f72612d3137336237302e737667)

![image](https://camo.githubusercontent.com/25f34b5d39178489662c05c3dc87c09e3596be8fd056d7c6abcd553dc1221e05/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d64656269616e2d6337303033362e737667)

![image](https://camo.githubusercontent.com/0a77d607254c6d408f949e437af3526c4e1b73665600a7794f43fd6f3baa0529/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d67656e746f6f2d3465343337312e737667)

![image](https://camo.githubusercontent.com/3ed8532663fcddc64720d7ff2ca6c37c9c2ef31ca4aee1bc675b9604b6929441/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d70726f6c696e75782d3361363739642e737667)

![image](https://camo.githubusercontent.com/ca3803e460b7c1a7888b4049e26b8c77c8b487cb3fe1b09f40e7204ddbd3c1a7/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d617263686c696e75782d3137393264302e737667)

![image](https://camo.githubusercontent.com/581b044785a0092e41bb04c18fda62ab8f2f140c1b9d514673a39c5f871dffc6/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d6e65746273642d6632363731312e737667)

![image](https://camo.githubusercontent.com/51e6ebc9379f7aeadc582a653325b932586cecd7cde9a3321e667fde0ecf7a1b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d667265656273642d3863303730372e737667)

![image](https://camo.githubusercontent.com/8dd43d97d7f29c1d4c8aa284b372b92110e75f1cc002ed62f5bcd46e161c3716/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d616e61636f6e64612d3633616434312e737667)

![image](https://camo.githubusercontent.com/eac585ffc1dd6339af0f2904f03a50b75835ad007e7bb789c70b8e583094ba33/68747470733a2f2f736e617063726166742e696f2f)

![image](https://camo.githubusercontent.com/7ff22e696c76053037a710bb9df850a60c12739e01edc68ae8db3f5dbb2b48dd/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d686f6d65627265772d6261383332622e737667)

![image](https://camo.githubusercontent.com/13fc8174ee9cd2ebff64fad0d3f26a8b5ec32e0fd81f04bb7d3fe5b8a8cf367f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d6d6163706f7274732d3133373665632e737667)

![image](https://camo.githubusercontent.com/138541e339c4b95e6ff39f4de585e2c694b070db627cbcc56023e3a577bc988e/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d63686f636f6c617465792d7265642e737667)

![image](https://camo.githubusercontent.com/1cfae973a2b4a98a6f716bd7b000b79321ed5eff0dccca5e560f322a6b51ce27/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f64697374726f732d77696e6765742d3339326635352e737667)

| OS | Installation command |
| --- | --- |
| Linux | `yum install miller` `apt-get install miller` `snap install miller` |
| Mac | `brew install millerport install miller` |
| Windows | `choco install millerwinget install Miller.Millerscoop install main/miller` |

See also [README-versions.md](https://github.com/johnkerl/miller/blob/main/README-versions.md) for a full list of package versions. Note that long-term-support (LtS) releases will likely be on older versions.

See also [building from source](https://miller.readthedocs.io/en/latest/build.html).

# Community

![image](https://camo.githubusercontent.com/8c2840e7df5d04fe6cb53b456eb2b0d6b922f59f73a57ab0a111730a0891f182/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f73746172732f6a6f686e6b65726c2f6d696c6c65722e7376673f6c6162656c3d4769744875622532307374617273)

![image](https://camo.githubusercontent.com/3f9bba492fcb7f35ca269ff8cc907fd679e611c5b25e18cff4c30e2c684317da/68747470733a2f2f6261646765732e77656172656f70656e736f757263652e6d652f686f6d65627265772f696e7374616c6c732f64792f6d696c6c65723f6c6162656c3d486f6d6562726577253230646f776e6c6f616473)

![image](https://camo.githubusercontent.com/f30b147950dd0a23389c4b1f04d951724f3574388c16ed95c9ac97787314acc3/68747470733a2f2f616e61636f6e64612e6f72672f636f6e64612d666f7267652f6d696c6c65722f6261646765732f646f776e6c6f6164732e7376673f6c6162656c3d436f6e6461253230646f776e6c6f616473)

![image](https://camo.githubusercontent.com/a64d66aa2c63fba77fe0d7e5d29ed5cac1afcf39dc7464b7908705f615abef26/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f616c6c5f636f6e7472696275746f72732d34312d6f72616e67652e7376673f7374796c653d666c61742d737175617265)

- Discussion forum: [https://github.com/johnkerl/miller/discussions](https://github.com/johnkerl/miller/discussions)
- Feature requests / bug reports: [https://github.com/johnkerl/miller/issues](https://github.com/johnkerl/miller/issues)
- How to contribute: [https://miller.readthedocs.io/en/latest/contributing/](https://miller.readthedocs.io/en/latest/contributing/)

# Build status

![image](https://github.com/johnkerl/miller/actions/workflows/go.yml/badge.svg)

![image](https://github.com/johnkerl/miller/actions/workflows/codeql-analysis.yml/badge.svg)

![image](https://github.com/johnkerl/miller/actions/workflows/codespell.yml/badge.svg)

# Building from source

- First: 
- With `make`: 
- Without `make`: 
- See also the doc page on [building from source](https://miller.readthedocs.io/en/latest/build).
- For more developer information please see [README-dev.md](https://github.com/johnkerl/miller/blob/main/README-dev.md).

# For developers

- [README-dev.md](https://github.com/johnkerl/miller/blob/main/README-dev.md)
- [How to contribute](https://miller.readthedocs.io/en/latest/contributing/)

# License

[License: BSD2](https://github.com/johnkerl/miller/blob/main/LICENSE.txt)

# Features

- 
- 
- 
- 
- 
- 
- 
- 
- 
- 
- 
- 

# What people are saying about Miller

> Today I discovered Miller—it's like jq but for CSV: https://t.co/pn5Ni241KM

> Underappreciated swiss-army command-line chainsaw.

> Miller looks like a great command line tool for working with CSV data. Sed, awk, cut, join all rolled into one: http://t.co/9BBb6VCZ6Y

> Miller is like sed, awk, cut, join, and sort for name-indexed data such as CSV: http://t.co/1zPbfg6B2W - handy tool!

> Btw, I think Miller is the best CLI tool to deal with CSV. I used to use this when I need to preprocess too big CSVs to load into R (now we have vroom, so such cases might be rare, though...)https://t.co/kUjrSSGJoT

> Miller: a *format-aware* data munging tool By @__jo_ker__ to overcome limitations with *line-aware* workshorses like awk, sed et al https://t.co/LCyPkhYvt9

> Holy holly data swiss army knife batman! How did no one suggest Miller https://t.co/JGQpmRAZLv for solving database cleaning / ETL issues to me before

> 🤯@__jo_ker__'s Miller easily reads, transforms, + writes all sorts of tabular data. It's standalone, fast, and built for streaming data (operating on one line at a time, so you can work on files larger than memory).

## Contributors ✨

Thanks to all the fine people who help make Miller better ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

![image](https://camo.githubusercontent.com/3dd511316e039f6c199497d117d68ff9b5c43e8583e202ea1687f012a5e682ab/68747470733a2f2f636f6e7472696275746f72732d696d672e7765622e6170702f696d6167653f7265706f3d6a6f686e6b65726c2f6d696c6c6572)

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind are welcome!
