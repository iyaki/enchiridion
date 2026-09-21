---
title: "Pandoc - universal document converter"
notion_id: fa076034-0480-42d5-ad9f-c604aa626c6e
notion_url: https://app.notion.com/p/Pandoc-universal-document-converter-fa076034048042d5ad9fc604aa626c6e
last_edited: 2023-01-20T13:28:00.000Z
source_url: https://pandoc.org/
tags: ["English", "Untried", "Office", "Tool"]
---
If you need to convert files from one markup format into another, pandoc is your swiss-army knife. Pandoc can convert between the following formats:

(← = conversion from; → = conversion to; ↔︎ = conversion from and to)

Pandoc understands a number of useful markdown syntax extensions, including document metadata (title, author, date); footnotes; tables; definition lists; superscript and subscript; strikeout; enhanced ordered lists (start number and numbering style are significant); running example lists; delimited code blocks with syntax highlighting; smart quotes, dashes, and ellipses; markdown inside HTML blocks; and inline LaTeX. If strict markdown compatibility is desired, all of these extensions can be turned off.

LaTeX math (and even macros) can be used in markdown documents. Several different methods of rendering math in HTML are provided, including MathJax and translation to MathML. LaTeX math is converted (as needed by the output format) to unicode, native Word equation objects, MathML, or roff eqn.

Pandoc includes a powerful system for automatic citations and bibliographies. This means that you can write a citation like

```plain text
[see @doe99, pp. 33-35; also @smith04, ch. 1]
```

and pandoc will convert it into a properly formatted citation using any of hundreds of [CSL](http://citationstyles.org/) styles (including footnote styles, numerical styles, and author-date styles), and add a properly formatted bibliography at the end of the document. The bibliographic data may be in [BibTeX](https://tug.org/bibtex/), [BibLaTeX](https://github.com/plk/biblatex), [CSL JSON](https://citeproc-js.readthedocs.io/en/latest/csl-json/markup.html), or CSL YAML format. Citations work in every output format.

There are many ways to customize pandoc to fit your needs, including a template system and a powerful system for writing filters.

Pandoc includes a Haskell library and a standalone command-line program. The library includes separate modules for each input and output format, so adding a new input or output format just requires adding a new module.

Pandoc is free software, released under the [GPL](https://www.gnu.org/copyleft/gpl.html). Copyright 2006–2022 [John MacFarlane](http://johnmacfarlane.net/).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
