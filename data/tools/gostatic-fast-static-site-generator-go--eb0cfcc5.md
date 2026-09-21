---
title: "gostatic - Fast static site generator (go)"
notion_id: eb0cfcc5-e69e-438f-90c5-e4197776be9b
notion_url: https://app.notion.com/p/gostatic-Fast-static-site-generator-go-eb0cfcc5e69e438f90c5e4197776be9b
last_edited: 2022-12-20T23:37:00.000Z
source_url: https://github.com/piranha/gostatic/#
tags: ["English", "Frontend", "Blogging/Content Creation", "Untried", "Tool"]
---
Gostatic is a static site generator. It tracks file changes during compilation, which is why it works reasonably fast. Also it provides framework for configuration akin to Make, which makes it easy to understand and to write custom configurations.

Features include:

- No run-time dependencies, just a single binary - download it and run
- Dependency tracking and re-rendering only changed pages
- Markdown support - Commonmark via goldmark
- Simple config syntax
- Flexible filter system
- Support for pagination
- Plays well with external commands and scripts
- HTTP server and watcher (instant rendering on changes)
- Suitable for automation (ability to query state with gostatic --dump)

And all in all, it works nicely for me, so it may work for you!

## Installation

Just download a binary from releases page.

If you need to automate downloading latest release, I use this script (change 64-linux to the type you need):

```plain text
URL=$(curl -s https://api.github.com/repos/piranha/gostatic/releases | awk '/download_url.*64-linux/ { print $2; exit }') curl -Lso gostatic $(URL) chmod +x gostatic
```

When downloading in Macos, your file will be quarantined and to use it you'll have to assure Macos you know what you're doing:

```plain text
xattr -r -d com.apple.quarantine gostatic
```

Obviously, go get https://github.com/piranha/gostatic also works, if you want to compile from source.

## Quick Start

Run gostatic -i my-site to generate basic site in directory my-site. It will have a basic config file, which you should edit to put relevant variables at the top - it also contains description of how files in your src directory are treated.

src directory obviously contains sources of your site (name of this directory can be changed in config). You can follow general idea of this directory to create new blog posts or new pages. All files, which are not mentioned in config, are just copied over. Run gostatic -fv config to see how your src is processed.

site.tmpl is a file that defines templates your are able to use for your pages. You can see those templates mentioned in config.

And, finally, there is a Makefile, just for convenience. Run make to build your site once or make w to run watcher and server, to see your site changes in real time.

Also, you could look at my site for an example of advanced usage.

Good luck! And remember, your contributions either to gostatic or to documentation (even if it's just this README.md) are always very welcome!

## Documentation index:

- Approach
- Speed
- External Resources
- Configuration Constants
- Page Config
- Processors
- Template API Reference Global Functions Page interface Page list interface Site interface
- Extensibility

Also, see wiki - and feel free to add more information there!

## Approach

Each given file is processed through a pipeline of filters, which modify the file state and then rendered on disk. Single input file corresponds to a single output file, but filters can generate virtual input files (like tag files).

File is rendering in those cases:

- output file does not exists
- file source is newer than it's output
- one of those is the case for one of file's dependencies

All files are sorted by date. This date is taken in their config or, in case if date in config is absent or dates there are equal, by file modification time.

## Speed

Box Pages Full Noop Single post Macbook Air '15 (i7) 630 450ms 100ms 160ms (93 pages) Macbook Air '15 (i7) 250 180ms n/a n/a Macbook Air '20 (M1) 486 185ms 39ms 140ms (97 pages) Hetzner CX21 (2 vCPU/4GB) 486 502ms 89ms 414ms (97 pages)

This are results of forced full site rebuild, then checking there are no modification, and then re-rendering a single changed page (along with pages which depend on this page).

Also note that if you're using various external post-processors (like uglifyjs or sassc) they tend to slow down things a bit (for my specific use case both uglifyjs and sassc add another 0.25s when files they process change).

To reproduce numbers, download hyperfine, download gostatic, clone solovyov.net, comment out :google-closure-compiler in config and then run:

- hyperfine 'gostatic -f config'
- hyperfine 'gostatic config'
- hyperfine 'touch src/blog/2017/fuji-raw.md && gostatic config'

## External resources

- Jack Pearkes made a Heroku buildpack for gostatic.

## Configuration

Config syntax is Makefile-inspired with some simplifications, look at the example:

```plain text
TEMPLATES = site.tmpl templates-folder SOURCE = src OUTPUT = site # this is a comment *.md: config ext .html directorify tags tags/*.tag markdown template page # yeah, this is a comment as well index.md: blog/*.md config ext .html inner-template markdown template page *.tag: blog/*.md ext .html directorify template tag markdown template page
```

Here we have constants declaration (first three lines), a comment and then three rules. One for any markdown file, one specifically for index.md and one for generated tags.

Specific rules override generic matching rules, but logic is not exactly very smart, and there is no real precedence defined, so if you have several matches for a single file you could end up with any of them. Note that there is some order: exact path match, exact name match, glob path match, glob name match. NOTE: this may change in future.

Rules consist of path/match, list of dependencies (also paths and matches, the ones listed after colon) and commands.

Each command consists of a name of processor and (possibly) some arguments. Arguments are separated by spaces.

Note: if a file has no rules whatsoever, it will be copied to exactly same location at destination as it was in source without being read into memory. So heavy images etc shouldn't be a problem.

### Constants

There are three configuration constants:

- SOURCE - sources to read (relative to location of config)
- OUTPUT - directory for output (relative to location of config)
- TEMPLATES - list of files and/or directories (containing *.tmpl files), which will be parsed as Go templates. Each file can contain more than one template (see docs on that).

You can also use arbitrary names for constants to access later from templates - just use any other name (AUTHOR could be one).

All constants can also be accessed from the config itself, using $(CONSTANT_NAME) syntax, just like in Makefile.

## Page Config

Page config is only processed if you specify config processor for a page. It's format is name: value, for example:

```plain text
title: This is a page tags: test date: 2013-01-05
```

Parsed properties:

- title - page title.
- tags - list of tags, separated by ,.
- date - page date, could be used for blog. Accepts formats from bigger to smaller (from "2006-01-02 15:04:05 -07" to "2006-01-02")
- hide - false if not specified or is one of f, false, False, FALSE. True in other cases. Hides page from children and tag lists when true.

You can also define any other property you like, it's value will be treated as a string and it's key is capitalized and put on the .Other page property.

## Processors

You can always check list of available processors with gostatic --processors.

- config - reads config from content. Config should be in format "name: value" and separated by four dashes on empty line (----) from content.
- ignore - ignore file.
- rename <new-name> - rename a file to new-name. Note this does not change path to a file (you can use .., though, but be careful about platform differences). If new-name contains *, then it'll be replaced with content of * from path match. For example, with blog/*.md: rename ../blog-*.html this will rename blog/one.html to blog-one.html.
- ext <.ext> - change file extension to a given one (which should be prefixed with a dot).
- datefilename - rename a file from whatever/2021-02-08-name.html to whatever/name.html and set the page.Date to 2021-02-08.
- directorify - rename a file from whatever/name.html to whatever/name/index.html.
- markdown - process content as Markdown. markdown without any arguments will not do any code-block highlighting. markdown chroma=monokai will use the Chroma highlighter to highlight code blocks, using the Monokai style, with inline CSS styles. (No .css file needed). You can see the styles at the Chroma style previewer. The official list of styles is in the Chroma repo here.
- inner-template - process content as Go template.
- template <name> - pass page to a template named <name>.
- tags <path-pattern> - create a virtual page for all tags of a current page. This tag page has path formed by replacing * in <path-pattern> with a tag name and has a tag as its .Title (use {{ range .Site.Pages.WithTag .Title }}...{{end}} to get a list of tagged pages.
- relativize - change all urls archored at / to be relative (i.e. add appropriate amount of ../) so that generated content can be deployed in a subfolder of a site.
- external <command> <args...> - call external command with content of a page as stdin and using stdout as a new content of a page. Has a shortcut: :<command> <args...> (: is replaced with external ).
- paginate <n> <path-pattern> - create a virtual page for each n of pages (grouped by path-pattern, so you can paginate few groups of pages as a single one). path-pattern has * replaced by an index of this virtual page (1-based), and you can get a list of pages with {{ range paginator .}}...{{end}} (see paginator function). Using paginate with the same path-pattern on different types of pages will group them in same paginated list (request: please open an issue if you have an idea how to phrase this better).
- jekyllify - creating pages in jekyll style, for example, the page 2018-02-02-name.md will be converted to /2018/02/02/name.md.
- yaml - read the configuration for the page using yaml format (like jekyll).

## Template API Reference

Templating is provided using Go templates. See link for documentation on syntax.

Each template is executed in context of a page. This means it has certain properties and methods it can output or call to generate content, i.e. {{ .Content }} will output page content in place.

### Global functions

Go template system provides some convenient functions, and gostatic expands on that a bit:

- absurl <base> <url> - given an url and a base will join them to something sane: leave the <url> in place if it's absolute, or resolve it within <base> if it's not.
- changed <name> <value> - checks if <value> has changed since previous call with the same name. Storage used for checking is global over the whole run of gostatic, so choose unique names for different places.
- cut <begin> <end> <value> - cut partial content from <value>, delimited by regular expressions <begin> and <end>.
- hash <value> - return 32-bit hash of a given value.
- version <page> <path> - return relative URL to a page with resulting path <path> with ?v=<32-bit hash> appended (use to override cache settings for static files).
- truncate <length> <value> - truncate string to given length (if it's longer).
- strip_html <value> - remove all HTML tags from string.
- strip_newlines <value> - remove all line breaks and newlines from string.
- trim <value> - trim all whitespace (strings.TrimSpace).
- replace <old> <new> <value> - replace all occurrences of old with new in value.
- replacen <old> <new> <n> <value> - same as above, but only n times.
- replacere <pattern> <replacement> <value> - replace text in value according to regexp pattern and replacement.
- split <separator> <value> - split string by separator, generating an array (you can use range with result of this function).
- contains <needle> <value> - check if a string value contains needle.
- starts <needle> <value> - check if a string value starts with needle.
- ends <needle> <value> - check if a string value ends with needle.
- matches <pattern> <value> - check if a regexp pattern matches string value.
- refind <pattern> <value> - apply regexp pattern to a string value and return first submatch (the thing in parentheses), if any, or a whole matched string.
- markdown <value> - convert a string (value) from Markdown to HTML.
- paginator <page> - get a paginator object for current page (only works on pages created by paginate processor).
- exec <cmd> [<arg1> <arg2> ....] - exec a command with (optional) arguments.
- exectext <cmd> [<arg1> <arg2> ....] <text> - exec a command with (optional) arguments and last argument (presumably some text) bound to command's stdin. If you need to do something hard, use it like {{ exectext "sh" "-c" "pipe | line" .Content }}.
- excerpt <text> <maxWordCount> - Gets an excerpt from the given text, to a maximum of maxWordCount words. When the text is shortened, it will produce an [...] string, denoting there's more. For example, The quick brown fox with maxWordCount of 2 will result in The quick [...].
- even <n> - tests if n is divisible by 2.
- odd <n> - tests if n is not divisible by 2.
- count <text> - returns a number of words in text.
- reading_time <text> - returns reading time based on average reading speed being 200.
- some <x> <x> <x>.... - returns first non-nil value as a string
- dir <path> - returns all but the last element of a path (same as filepath.Dir)
- base <path> - returns the last element of a path (same as filepath.Base)
- abcsort <pages> - returns the pages sorted in alphabetical order of their .Name

### Page interface

- .Site - global site object.
- .Rule - rule object, matched by page.
- .Pattern - pattern, which matched this page.
- .Deps - list of pages, which are dependencies for this page.
- .Next - next page in a list of all site pages (use specific PageSlice's .Next method if you need more precise matching).
- .Prev - previous page in a list of all site pages (use specific PageSlice's .Prev method if you need more precise matching).
- .Source - relative path to page source.
- .FullPath - full path to page source.
- .Path - relative path to page destination.
- .OutputPath - full path to page destination.
- .ModTime - page last modification time.
- .Title - page title.
- .Tags - list of page tags.
- .Date - page date, as defined in page config.
- .Hide - boolean if page is going to be absent from {{ .Children }} or {{ .WithTag }} lists.
- .Other - map of all other properties (capitalized) from page config, like {{ .Other.Author }}.
- .Raw - page content after preprocessors (i.e. after config has stripped it part), that was originally read from the disk.
- .Content - page content.
- .Url - page url (i.e. .Path, but with index.html stripped from the end).
- .Name - page name (i.e. last part of .Url).
- .UrlTo <other-page> - relative url from current to some other page.
- .Rel <url> - relative url to given absolute (anchored at /) url.
- .Is <url> - checks if page is at passed url (or path) - use it for marking active elements in menu, for example.
- .UrlMatches <pattern> - checks if page url matches regular expression <pattern>.
- .Has <field> <value> - backend for .Where and .WhereNot, checks if field equals to value, or: "Url" - calls UrlMatches "Tag" - checks tag is present in .Tags "Source" - matches source path for value.

### Paginator interface

- .Number - number of paginator page, first is 1
- .PathPattern - whatever was passed as path-pattern to paginate (processor)[#processors]
- .Page - paginator's own page
- .Pages - list of pages
- .Prev - previous paginator object (if current is first, then nil)
- .Next - next paginator object (if current is last, then nil)

### Page list interface

- .Get <n> - page number <n>.
- .First - first page.
- .Last - last page.
- .Len - length of page list.
- .Prev <page> - return page with earlier date than given. Returns nil if no earlier pages exist or page is not in page list.
- .Next <page> - return page with later date than given. Returns nil if no later pages exist or page is not in page list.
- .Slice <from> <to> - return pages from from to to (i.e. from 0 to 10).
- .Children <prefix> - list of pages, nested under <prefix>.
- .WithTag <tag-name> - list of pages, tagged with <tag-name>.
- .Reverse - list of pages, sorted in reverse order.
- .BySource <path> - finds a page with source path <path>.
- .ByPath <path> - finds a page with resulting path <path>.
- .GlobSource <pattern> - list of pages, matching source path <pattern>.
- .Where <field> <value> - list of pages, which return true for .Has <field> <value>
- .WhereNot <field> <value> - list of pages, which return false for .Has <field> <value>

### Site interface

- .Pages - list of all pages.
- .Source - path to site source.
- .Output - path to site destination.
- .Templates - list of template files used for the site.
- .Other - any other properties (capitalized) defined in site config.

## Extensibility

Obviously, the easiest way to extend gostatic's functionality is to use external processor. It makes you able to process files in the way you want, but is more or less limited to that. There is no API right now to create pages on the fly (like tags processor does) using this method, for example.

But gostatic itself is a library, and you can write your own static site generator using this library. See gostatic.go for an example of one.
