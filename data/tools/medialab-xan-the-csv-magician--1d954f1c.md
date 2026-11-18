---
title: "medialab/xan: The CSV magician"
notion_id: 1d954f1c-7d23-81c6-8a1c-c0d2c2afedd9
notion_url: https://app.notion.com/p/medialab-xan-The-CSV-magician-1d954f1c7d2381c68a1cc0d2c2afedd9
last_edited: 2025-06-22T02:06:00.000Z
source_url: https://github.com/medialab/xan
tags: ["English", "Office", "Tool"]
---
![image](https://github.com/medialab/xan/workflows/Tests/badge.svg)

# `xan`, the CSV magician

`xan` is a command line tool that can be used to process CSV files directly from the shell.

It has been written in Rust to be as fast as possible, use as little memory as possible, and can easily handle very large CSV files (Gigabytes). It is also able to leverage parallelism (through multithreading) to make some tasks complete as fast as your computer can allow.

It can easily preview, filter, slice, aggregate, sort, join CSV files, and exposes a large collection of composable commands that can be chained together to perform a wide variety of typical tasks.

`xan` also leverages its own expression language so you can perform complex tasks that cannot be done by relying on the simplest commands. This minimalistic language has been tailored for CSV data and is faster than evaluating typical dynamically-typed languages such as Python, Lua, JavaScript etc.

Note that this tool is originally a fork of [BurntSushi](https://github.com/BurntSushi)'s [`xsv`](https://github.com/BurntSushi/xsv), but has been nearly entirely rewritten at that point, to fit [SciencesPo's médialab](https://github.com/medialab) use-cases, rooted in web data collection and analysis geared towards social sciences (you might think CSV is outdated by now, but read our [love letter](https://github.com/medialab/xan/blob/master/docs/LOVE_LETTER.md) to the format before judging too quickly). `xan` therefore goes beyond typical data manipulation and expose utilities related to lexicometry, graph theory and even scraping.

Finally, `xan` can be used to display CSV files in the terminal, for easy exploration, and can even be used to draw basic data visualisations:

<!-- unsupported block: child_database -->

## Summary

- [How to install](https://github.com/medialab/xan#how-to-install) 
- [Cargo](https://github.com/medialab/xan#cargo)
- [Scoop (Windows)](https://github.com/medialab/xan#scoop-windows)
- [Homebrew (macOS)](https://github.com/medialab/xan#homebrew-macos)
- [Arch Linux](https://github.com/medialab/xan#arch-linux)
- [Nix](https://github.com/medialab/xan#nix)
- [Pre-built binaries](https://github.com/medialab/xan#pre-built-binaries)
- [Installing completions](https://github.com/medialab/xan#installing-completions)
- [Quick tour](https://github.com/medialab/xan#quick-tour)
- [Available commands](https://github.com/medialab/xan#available-commands)
- [General flags and IO model](https://github.com/medialab/xan#general-flags-and-io-model)
- [Expression language reference](https://github.com/medialab/xan#expression-language-reference)
- [Cookbook](https://github.com/medialab/xan#cookbook)
- [News](https://github.com/medialab/xan#news)
- [Frequently Asked Questions](https://github.com/medialab/xan#frequently-asked-questions)

## How to install

### Cargo

`xan` can be installed using cargo (it usually comes with [Rust](https://www.rust-lang.org/tools/install)):

```plain text
cargo install xan
```

You can also tweak the build flags to make sure the Rust compiler is able to leverage all your CPU's features:

```plain text
CARGO_BUILD_RUSTFLAGS='-C target-cpu=native' cargo install xan
```

You can also install the latest dev version thusly:

```plain text
cargo install --git https://github.com/medialab/xan
```

### Scoop (Windows)

`xan` can be installed using [Scoop](https://scoop.sh/) on Windows:

```plain text
scoop bucket add extras
scoop install xan
```

### Homebrew (macOS)

`xan` can be installed with [Homebrew](https://brew.sh/) on macOS thusly:

```plain text
brew install xan
```

### Arch Linux

You can install `xan` from the [extra repository](https://archlinux.org/packages/extra/x86_64/xan/) using `pacman`:

```plain text
sudo pacman -S xan
```

### Nix

`xan` is packaged for Nix, and is available in Nixpkgs as of 25.05 release. To install it, you may add it to your `environment.systemPackages` as `pkgs.xan` or use `nix-shell` to enter an ephemeral shell.

```plain text
nix-shell -p xan
```

### Pre-built binaries

Pre-built binaries can be found attached to every GitHub [releases](https://github.com/medialab/xan/releases/latest).

Currently supported targets include:

- `x86_64-unknown-linux-musl`
- `x86_64-pc-windows-gnu`

Feel free to open a PR to improve the CI by adding relevant targets.

### Installing completions

Note that `xan` also exposes handy automatic completions for command and header/column names that you can install through the `xan completions` command.

Run the following command to understand how to install those completions:

```plain text
xan completions -h
```

## Quick tour

Let's learn about the most commonly used `xan` commands by exploring a corpus of French medias:

### Downloading the corpus

```plain text
curl -LO https://github.com/medialab/corpora/raw/master/polarisation/medias.csv
```

### Displaying the file's headers

```plain text
xan headers medias.csv
```

```plain text
0   webentity_id
1   name
2   prefixes
3   home_page
4   start_pages
5   indegree
6   hyphe_creation_timestamp
7   hyphe_last_modification_timestamp
8   outreach
9   foundation_year
10  batch
11  edito
12  parody
13  origin
14  digital_native
15  mediacloud_ids
16  wheel_category
17  wheel_subcategory
18  has_paywall
19  inactive

```

### Counting the number of rows

```plain text
xan count medias.csv
```

```plain text
478

```

### Previewing the file in the terminal

```plain text
xan view medias.csv
```

```plain text
Displaying 5/20 cols from 10 first rows of medias.csv
┌───┬───────────────┬───────────────┬────────────┬───┬─────────────┬──────────┐
│ - │ name          │ prefixes      │ home_page  │ … │ has_paywall │ inactive │
├───┼───────────────┼───────────────┼────────────┼───┼─────────────┼──────────┤
│ 0 │ Acrimed.org   │ http://acrim… │ http://ww… │ … │ false       │ <empty>  │
│ 1 │ 24matins.fr   │ http://24mat… │ https://w… │ … │ false       │ <empty>  │
│ 2 │ Actumag.info  │ http://actum… │ https://a… │ … │ false       │ <empty>  │
│ 3 │ 2012un-Nouve… │ http://2012u… │ http://ww… │ … │ false       │ <empty>  │
│ 4 │ 24heuresactu… │ http://24heu… │ http://24… │ … │ false       │ <empty>  │
│ 5 │ AgoraVox      │ http://agora… │ http://ww… │ … │ false       │ <empty>  │
│ 6 │ Al-Kanz.org   │ http://al-ka… │ https://w… │ … │ false       │ <empty>  │
│ 7 │ Alalumieredu… │ http://alalu… │ http://al… │ … │ false       │ <empty>  │
│ 8 │ Allodocteurs… │ http://allod… │ https://w… │ … │ false       │ <empty>  │
│ 9 │ Alterinfo.net │ http://alter… │ http://ww… │ … │ <empty>     │ true     │
│ … │ …             │ …             │ …          │ … │ …           │ …        │
└───┴───────────────┴───────────────┴────────────┴───┴─────────────┴──────────┘

```

On unix, don't hesitate to use the `-p` flag to automagically forward the full output to an appropriate pager and skim through all the columns.

### Reading a flattened representation of the first row

```plain text
# NOTE: drop -c to avoid truncating the values
xan flatten -c medias.csv
```

```plain text
Row n°0
───────────────────────────────────────────────────────────────────────────────
webentity_id                      1
name                              Acrimed.org
prefixes                          http://acrimed.org|http://acrimed69.blogspot…
home_page                         http://www.acrimed.org
start_pages                       http://acrimed.org|http://acrimed69.blogspot…
indegree                          61
hyphe_creation_timestamp          1560347020330
hyphe_last_modification_timestamp 1560526005389
outreach                          nationale
foundation_year                   2002
batch                             1
edito                             media
parody                            false
origin                            france
digital_native                    true
mediacloud_ids                    258269
wheel_category                    Opinion Journalism
wheel_subcategory                 Left Wing
has_paywall                       false
inactive                          <empty>

Row n°1
───────────────────────────────────────────────────────────────────────────────
webentity_id                      2
...

```

### Searching for rows

```plain text
xan search -s outreach internationale medias.csv | xan view
```

```plain text
Displaying 4/20 cols from 10 first rows of <stdin>
┌───┬──────────────┬────────────────────┬───┬─────────────┬──────────┐
│ - │ webentity_id │ name               │ … │ has_paywall │ inactive │
├───┼──────────────┼────────────────────┼───┼─────────────┼──────────┤
│ 0 │ 25           │ Businessinsider.fr │ … │ false       │ <empty>  │
│ 1 │ 59           │ Europe-Israel.org  │ … │ false       │ <empty>  │
│ 2 │ 66           │ France 24          │ … │ false       │ <empty>  │
│ 3 │ 220          │ RFI                │ … │ false       │ <empty>  │
│ 4 │ 231          │ fr.Sott.net        │ … │ false       │ <empty>  │
│ 5 │ 246          │ Voltairenet.org    │ … │ true        │ <empty>  │
│ 6 │ 254          │ Afp.com /fr        │ … │ false       │ <empty>  │
│ 7 │ 265          │ Euronews FR        │ … │ false       │ <empty>  │
│ 8 │ 333          │ Arte.tv            │ … │ false       │ <empty>  │
│ 9 │ 341          │ I24News.tv         │ … │ false       │ <empty>  │
│ … │ …            │ …                  │ … │ …           │ …        │
└───┴──────────────┴────────────────────┴───┴─────────────┴──────────┘

```

### Selecting some columns

```plain text
xan select foundation_year,name medias.csv | xan view
```

```plain text
Displaying 2 cols from 10 first rows of <stdin>
┌───┬─────────────────┬───────────────────────────────────────┐
│ - │ foundation_year │ name                                  │
├───┼─────────────────┼───────────────────────────────────────┤
│ 0 │ 2002            │ Acrimed.org                           │
│ 1 │ 2006            │ 24matins.fr                           │
│ 2 │ 2013            │ Actumag.info                          │
│ 3 │ 2012            │ 2012un-Nouveau-Paradigme.com          │
│ 4 │ 2010            │ 24heuresactu.com                      │
│ 5 │ 2005            │ AgoraVox                              │
│ 6 │ 2008            │ Al-Kanz.org                           │
│ 7 │ 2012            │ Alalumieredunouveaumonde.blogspot.com │
│ 8 │ 2005            │ Allodocteurs.fr                       │
│ 9 │ 2005            │ Alterinfo.net                         │
│ … │ …               │ …                                     │
└───┴─────────────────┴───────────────────────────────────────┘

```

### Sorting the file

```plain text
xan sort -s foundation_year medias.csv | xan view -s name,foundation_year
```

```plain text
Displaying 2 cols from 10 first rows of <stdin>
┌───┬────────────────────────────────────┬─────────────────┐
│ - │ name                               │ foundation_year │
├───┼────────────────────────────────────┼─────────────────┤
│ 0 │ Le Monde Numérique (Ouest France)  │ <empty>         │
│ 1 │ Le Figaro                          │ 1826            │
│ 2 │ Le journal de Saône-et-Loire       │ 1826            │
│ 3 │ L'Indépendant                      │ 1846            │
│ 4 │ Le Progrès                         │ 1859            │
│ 5 │ La Dépêche du Midi                 │ 1870            │
│ 6 │ Le Pélerin                         │ 1873            │
│ 7 │ Dernières Nouvelles d'Alsace (DNA) │ 1877            │
│ 8 │ La Croix                           │ 1883            │
│ 9 │ Le Chasseur Francais               │ 1885            │
│ … │ …                                  │ …               │
└───┴────────────────────────────────────┴─────────────────┘

```

### Deduplicating the file on some column

```plain text
# Some medias of our corpus have the same ids on mediacloud.org
xan dedup -s mediacloud_ids medias.csv | xan count && xan count medias.csv
```

```plain text
457
478

```

Deduplicating can also be done while sorting:

```plain text
xan sort -s mediacloud_ids -u medias.csv
```

### Computing frequency tables

```plain text
xan frequency -s edito medias.csv | xan view
```

```plain text
Displaying 3 cols from 5 rows of <stdin>
┌───┬───────┬────────────┬───────┐
│ - │ field │ value      │ count │
├───┼───────┼────────────┼───────┤
│ 0 │ edito │ media      │ 423   │
│ 1 │ edito │ individu   │ 30    │
│ 2 │ edito │ plateforme │ 14    │
│ 3 │ edito │ agrégateur │ 10    │
│ 4 │ edito │ agence     │ 1     │
└───┴───────┴────────────┴───────┘

```

### Printing a histogram

```plain text
xan frequency -s edito medias.csv | xan hist
```

```plain text
Histogram for edito (bars: 5, sum: 478, max: 423):

media      |423  88.49%|━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━|
individu   | 30   6.28%|━━━╸                                                  |
plateforme | 14   2.93%|━╸                                                    |
agrégateur | 10   2.09%|━╸                                                    |
agence     |  1   0.21%|╸                                                     |

```

### Computing descriptive statistics

```plain text
xan stats -s indegree,edito medias.csv | xan transpose | xan view -I
```

```plain text
Displaying 2 cols from 14 rows of <stdin>
┌─────────────┬───────────────────┬────────────┐
│ field       │ indegree          │ edito      │
├─────────────┼───────────────────┼────────────┤
│ count       │ 463               │ 478        │
│ count_empty │ 15                │ 0          │
│ type        │ int               │ string     │
│ types       │ int|empty         │ string     │
│ sum         │ 25987             │ <empty>    │
│ mean        │ 56.12742980561554 │ <empty>    │
│ variance    │ 4234.530197929737 │ <empty>    │
│ stddev      │ 65.07326792108829 │ <empty>    │
│ min         │ 0                 │ <empty>    │
│ max         │ 424               │ <empty>    │
│ lex_first   │ 0                 │ agence     │
│ lex_last    │ 99                │ plateforme │
│ min_length  │ 0                 │ 5          │
│ max_length  │ 3                 │ 11         │
└─────────────┴───────────────────┴────────────┘

```

### Evaluating an expression to filter a file

```plain text
xan filter 'batch > 1' medias.csv | xan count
```

```plain text
130

```

To access the expression language's [cheatsheet](https://github.com/medialab/xan/blob/master/docs/moonblade/cheatsheet.md), run `xan help cheatsheet`. To display the full list of available [functions](https://github.com/medialab/xan/blob/master/docs/moonblade/functions.md), run `xan help functions`.

### Evaluating an expression to create a new column based on other ones

```plain text
xan map 'fmt("{} ({})", name, foundation_year)' key medias.csv | xan select key | xan slice -l 10
```

```plain text
key
Acrimed.org (2002)
24matins.fr (2006)
Actumag.info (2013)
2012un-Nouveau-Paradigme.com (2012)
24heuresactu.com (2010)
AgoraVox (2005)
Al-Kanz.org (2008)
Alalumieredunouveaumonde.blogspot.com (2012)
Allodocteurs.fr (2005)
Alterinfo.net (2005)

```

To access the expression language's [cheatsheet](https://github.com/medialab/xan/blob/master/docs/moonblade/cheatsheet.md), run `xan help cheatsheet`. To display the full list of available [functions](https://github.com/medialab/xan/blob/master/docs/moonblade/functions.md), run `xan help functions`.

### Transform a column by evaluating an expression

```plain text
xan transform name 'split(name, ".") | first | upper' medias.csv | xan select name | xan slice -l 10
```

```plain text
name
ACRIMED
24MATINS
ACTUMAG
2012UN-NOUVEAU-PARADIGME
24HEURESACTU
AGORAVOX
AL-KANZ
ALALUMIEREDUNOUVEAUMONDE
ALLODOCTEURS
ALTERINFO

```

To access the expression language's [cheatsheet](https://github.com/medialab/xan/blob/master/docs/moonblade/cheatsheet.md), run `xan help cheatsheet`. To display the full list of available [functions](https://github.com/medialab/xan/blob/master/docs/moonblade/functions.md), run `xan help functions`.

### Performing custom aggregation

```plain text
xan agg 'sum(indegree) as total_indegree, mean(indegree) as mean_indegree' medias.csv | xan view -I
```

```plain text
Displaying 1 col from 1 rows of <stdin>
┌────────────────┬───────────────────┐
│ total_indegree │ mean_indegree     │
├────────────────┼───────────────────┤
│ 25987          │ 56.12742980561554 │
└────────────────┴───────────────────┘

```

To access the expression language's [cheatsheet](https://github.com/medialab/xan/blob/master/docs/moonblade/cheatsheet.md), run `xan help cheatsheet`. To display the full list of available [functions](https://github.com/medialab/xan/blob/master/docs/moonblade/functions.md), run `xan help functions`. Finally, to display the list of available [aggregation functions](https://github.com/medialab/xan/blob/master/docs/moonblade/aggs.md), run `xan help aggs`.

### Grouping rows and performing per-group aggregation

```plain text
xan groupby edito 'sum(indegree) as indegree' medias.csv | xan view -I
```

```plain text
Displaying 1 col from 5 rows of <stdin>
┌────────────┬──────────┐
│ edito      │ indegree │
├────────────┼──────────┤
│ agence     │ 50       │
│ agrégateur │ 459      │
│ plateforme │ 658      │
│ media      │ 24161    │
│ individu   │ 659      │
└────────────┴──────────┘

```

To access the expression language's [cheatsheet](https://github.com/medialab/xan/blob/master/docs/moonblade/cheatsheet.md), run `xan help cheatsheet`. To display the full list of available [functions](https://github.com/medialab/xan/blob/master/docs/moonblade/functions.md), run `xan help functions`. Finally, to display the list of available [aggregation functions](https://github.com/medialab/xan/blob/master/docs/moonblade/aggs.md), run `xan help aggs`.

## Available commands

- [**help**](https://github.com/medialab/xan/blob/master/docs/cmd/help.md): Get help regarding the expression language

_Explore & visualize_

- [**count (c)**](https://github.com/medialab/xan/blob/master/docs/cmd/count.md): Count rows in file
- [**headers (h)**](https://github.com/medialab/xan/blob/master/docs/cmd/headers.md): Show header names
- [**view (v)**](https://github.com/medialab/xan/blob/master/docs/cmd/view.md): Preview a CSV file in a human-friendly way
- [**flatten**](https://github.com/medialab/xan/blob/master/docs/cmd/flatten.md): Display a flattened version of each row of a file
- [**hist**](https://github.com/medialab/xan/blob/master/docs/cmd/hist.md): Print a histogram with rows of CSV file as bars
- [**plot**](https://github.com/medialab/xan/blob/master/docs/cmd/plot.md): Draw a scatter plot or line chart
- [**heatmap**](https://github.com/medialab/xan/blob/master/docs/cmd/heatmap.md): Draw a heatmap of a CSV matrix
- [**progress**](https://github.com/medialab/xan/blob/master/docs/cmd/progress.md): Display a progress bar while reading CSV data

_Search & filter_

- [**search**](https://github.com/medialab/xan/blob/master/docs/cmd/search.md): Search for patterns in CSV data
- [**filter**](https://github.com/medialab/xan/blob/master/docs/cmd/filter.md): Only keep some CSV rows based on an evaluated expression
- [**slice**](https://github.com/medialab/xan/blob/master/docs/cmd/slice.md): Slice rows of CSV file
- [**top**](https://github.com/medialab/xan/blob/master/docs/cmd/top.md): Find top rows of a CSV file according to some column
- [**sample**](https://github.com/medialab/xan/blob/master/docs/cmd/sample.md): Randomly sample CSV data

_Sort & deduplicate_

- [**sort**](https://github.com/medialab/xan/blob/master/docs/cmd/sort.md): Sort CSV data
- [**dedup**](https://github.com/medialab/xan/blob/master/docs/cmd/dedup.md): Deduplicate a CSV file
- [**shuffle**](https://github.com/medialab/xan/blob/master/docs/cmd/shuffle.md): Shuffle CSV data

_Aggregate_

- [**frequency (freq)**](https://github.com/medialab/xan/blob/master/docs/cmd/frequency.md): Show frequency tables
- [**groupby**](https://github.com/medialab/xan/blob/master/docs/cmd/groupby.md): Aggregate data by groups of a CSV file
- [**stats**](https://github.com/medialab/xan/blob/master/docs/cmd/stats.md): Compute basic statistics
- [**agg**](https://github.com/medialab/xan/blob/master/docs/cmd/agg.md): Aggregate data from CSV file
- [**bins**](https://github.com/medialab/xan/blob/master/docs/cmd/bins.md): Dispatch numeric columns into bins

_Combine multiple CSV files_

- [**cat**](https://github.com/medialab/xan/blob/master/docs/cmd/cat.md): Concatenate by row or column
- [**join**](https://github.com/medialab/xan/blob/master/docs/cmd/join.md): Join CSV files
- [**regex-join**](https://github.com/medialab/xan/blob/master/docs/cmd/regex-join.md): Fuzzy join CSV files using regex patterns
- [**url-join**](https://github.com/medialab/xan/blob/master/docs/cmd/url-join.md): Join CSV files on url prefixes
- [**merge**](https://github.com/medialab/xan/blob/master/docs/cmd/merge.md): Merge multiple similar already sorted CSV files

_Add, transform, drop and move columns_

- [**select**](https://github.com/medialab/xan/blob/master/docs/cmd/select.md): Select columns from a CSV file
- [**drop**](https://github.com/medialab/xan/blob/master/docs/cmd/drop.md): Drop columns from a CSV file
- [**map**](https://github.com/medialab/xan/blob/master/docs/cmd/map.md): Create a new column by evaluating an expression on each CSV row
- [**transform**](https://github.com/medialab/xan/blob/master/docs/cmd/transform.md): Transform a column by evaluating an expression on each CSV row
- [**enum**](https://github.com/medialab/xan/blob/master/docs/cmd/enum.md): Enumerate CSV file by preprending an index column
- [**flatmap**](https://github.com/medialab/xan/blob/master/docs/cmd/flatmap.md): Emit one row per value yielded by an expression evaluated for each CSV row
- [**fill**](https://github.com/medialab/xan/blob/master/docs/cmd/fill.md): Fill empty cells
- [**blank**](https://github.com/medialab/xan/blob/master/docs/cmd/blank.md): Blank down contiguous identical cell values

_Format, convert & recombobulate_

- [**behead**](https://github.com/medialab/xan/blob/master/docs/cmd/behead.md): Drop header from CSV file
- [**rename**](https://github.com/medialab/xan/blob/master/docs/cmd/rename.md): Rename columns of a CSV file
- [**input**](https://github.com/medialab/xan/blob/master/docs/cmd/input.md): Read unusually formatted CSV data
- [**fixlengths**](https://github.com/medialab/xan/blob/master/docs/cmd/fixlengths.md): Makes all rows have same length
- [**fmt**](https://github.com/medialab/xan/blob/master/docs/cmd/fmt.md): Format CSV output (change field delimiter)
- [**explode**](https://github.com/medialab/xan/blob/master/docs/cmd/explode.md): Explode rows based on some column separator
- [**implode**](https://github.com/medialab/xan/blob/master/docs/cmd/implode.md): Collapse consecutive identical rows based on a diverging column
- [**from**](https://github.com/medialab/xan/blob/master/docs/cmd/from.md): Convert a variety of formats to CSV
- [**to**](https://github.com/medialab/xan/blob/master/docs/cmd/to.md): Convert a CSV file to a variety of data formats
- [**scrape**](https://github.com/medialab/xan/blob/master/docs/cmd/scrape.md): Scrape HTML into CSV data
- [**reverse**](https://github.com/medialab/xan/blob/master/docs/cmd/reverse.md): Reverse rows of CSV data
- [**transpose (t)**](https://github.com/medialab/xan/blob/master/docs/cmd/transpose.md): Transpose CSV file

_Split a CSV file into multiple_

- [**split**](https://github.com/medialab/xan/blob/master/docs/cmd/split.md): Split CSV data into chunks
- [**partition**](https://github.com/medialab/xan/blob/master/docs/cmd/partition.md): Partition CSV data based on a column value

_Parallel operation over multiple CSV files_

- [**parallel (p)**](https://github.com/medialab/xan/blob/master/docs/cmd/parallel.md): Map-reduce-like parallel computation over multiple CSV files

_Generate CSV files_

- [**range**](https://github.com/medialab/xan/blob/master/docs/cmd/range.md): Create a CSV file from a numerical range

_Perform side-effects_

- [**eval**](https://github.com/medialab/xan/blob/master/docs/cmd/eval.md): Evaluate/debug a single expression
- [**foreach**](https://github.com/medialab/xan/blob/master/docs/cmd/foreach.md): Loop over a CSV file to perform side effects

_Lexicometry & fuzzy matching_

- [**tokenize**](https://github.com/medialab/xan/blob/master/docs/cmd/tokenize.md): Tokenize a text column
- [**vocab**](https://github.com/medialab/xan/blob/master/docs/cmd/vocab.md): Build a vocabulary over tokenized documents
- [**cluster**](https://github.com/medialab/xan/blob/master/docs/cmd/cluster.md): Cluster CSV data to find near-duplicates

_Matrix & network-related commands_

- [**matrix**](https://github.com/medialab/xan/blob/master/docs/cmd/matrix.md): Convert CSV data to matrix data
- [**network**](https://github.com/medialab/xan/blob/master/docs/cmd/network.md): Convert CSV data to network data

## General flags and IO model

### Getting help

If you ever feel lost, each command has a `-h/--help` flag that will print the related documentation.

If you need help about the expression language, check out the `help` command itself:

```plain text
# Help about help ;)
xan help --help
```

### Regarding input & output formats

All `xan` commands expect a "standard" CSV file, e.g. comma-delimited, with proper double-quote escaping. This said, `xan` is also perfectly able to infer the delimiter from typical file extensions such as `.tsv` or `.tab`.

If you need to process a file with a custom delimiter, you can either use the `xan input` command or use the `-d/--delimiter` flag available with all commands.

If you need to output a custom CSV dialect (e.g. using `;` delimiters), feel free to use the `xan fmt` command.

Finally, even if most `xan` commands won't even need to decode the file's bytes, some might still need to. In this case, `xan` will expect correctly formatted UTF-8 text. Please use `iconv` or other utils if you need to process other encodings such as `latin1` ahead of `xan`.

### Working with headless CSV file

Even if this is good practice to name your columns, some CSV file simply don't have headers. Most commands are able to deal with those file if you give the `-n/--no-headers` flag.

Note that this flag always relates to the input, not the output. If for some reason you want to drop a CSV output's header row, use the `xan behead` command.

### Regarding stdin

By default, all commands will try to read from stdin when the file path is not specified. This makes piping easy and comfortable as it respects typical unix standards. Some commands may have multiple inputs (`xan join`, for instance), in which case stdin is usually specifiable using the `-` character:

```plain text
# First file given to join will be read from stdin
cat file1.csv | xan join col1 - col2 file2.csv
```

Note that the command will also warn you when stdin cannot be read, in case you forgot to indicate the file's path.

### Regarding stdout

By default, all commands will print their output to stdout (note that this output is usually buffered for performance reasons).

In addition, all commands expose a `-o/--output` flag that can be use to specify where to write the output. This can be useful if you do not want to or cannot use `>` (typically in some Windows shells). In which case, `-` as a output path will mean forwarding to stdout also. This can be useful when scripting sometimes.

### Gzipped files

`xan` is able to read gzipped files (having a `.gz` extension) out of the box.

## Expression language reference

- [Cheatsheet](https://github.com/medialab/xan/blob/master/docs/moonblade/cheatsheet.md)
- [Comprehensive list of functions & operators](https://github.com/medialab/xan/blob/master/docs/moonblade/functions.md)
- [Comprehensive list of aggregation functions](https://github.com/medialab/xan/blob/master/docs/moonblade/aggs.md)
- [Scraping DSL](https://github.com/medialab/xan/blob/master/docs/moonblade/scraping.md)

## Cookbook

- [Merging frequency tables, three ways](https://github.com/medialab/xan/blob/master/docs/cookbook/frequency_tables.md)
- [Parsing and visualizing dates with xan](https://github.com/medialab/xan/blob/master/docs/cookbook/dates.md)
- [Joining files by URL prefixes](https://github.com/medialab/xan/blob/master/docs/cookbook/urls.md)
- [Miscellaneous](https://github.com/medialab/xan/blob/master/docs/cookbook/misc.md)

## News

For news about the tool's evolutions feel free to read:

1. the [changelog](https://github.com/medialab/xan/blob/master/CHANGELOG.md)
2. the [xan zines](https://github.com/medialab/xan/blob/master/docs/XANZINE.md)

## Frequently Asked Questions

### How to display a vertical bar chart?

Rotate your screen ;)
