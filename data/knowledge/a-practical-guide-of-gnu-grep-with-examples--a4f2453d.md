---
title: "A Practical Guide of GNU grep With Examples"
notion_id: a4f2453d-0d05-4f87-aa3c-107055cf1184
notion_url: https://app.notion.com/p/A-Practical-Guide-of-GNU-grep-With-Examples-a4f2453d0d054f87aa3c107055cf1184
last_edited: 2023-09-04T18:33:00.000Z
source_url: https://thevaluable.dev/grep-cli-guide-examples/
tags: ["English", "Shell/Bash", "Article", "Guide", "The Valuable Dev"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->















- 
- 
- 
- 
- 
- 
- 
- 







## 



### 



```

```



1. 
2. 
3. 





```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



## 



### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



```

```





```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 



```

```









| Capability | Description | Default |
| --- | --- | --- |
| `sl` | The line where the pattern is matched (the selected line), without the match itself. | `sl=` (empty) |
| `cx` | The lines where the pattern is not matched (the context lines). | `cx=` (empty) |
| `ms` | The pattern matched (match in the selected line). | `ms=01;31` |
| `fn` | The eventual filename prefixing the selected line. | `fn=35` |
| `ln` | The eventual line numbers prefixing the selected line. | `ln=32` |
| `se` | Any separator displayed by grep. | `se=36` |





| Integer | Description |
| --- | --- |
| `1` | Bold. |
| `3` | Italic. |
| `4` | Underline. |
| `5` | Blink (to impress your coworkers). |
| `7` | Inverse the foreground and background color. |
| `39` | Default foreground color of you terminal. |
| `30` to `37` | Foreground colors set up for your terminal. |
| `38;5;0` to `38;5;255` | Foreground color (256 colors ANSI). |
| `49` | Default background color. |
| `40` to `47` | Background colors set up for your terminal. |
| `48;5;0` to `48;5;255` | Background color (256 colors ANSI). **TODO to test** |



```

```





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



```

```







```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



```

```



### 





```

```



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 







## 







### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 



```

```





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



```

```



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



| Option | Description | Example |
| --- | --- | --- |
| `-v` | Invert grep’s output: output the lines _not_ matching the pattern. | `grep -v 'div' styles.css` |
| `-o` | Output only the matches. | `grep -o 'div' styles.css` |
| `-n` | Add the line numbers to the output. | `grep -n 'div' styles.css` |
| `-c` | Only output the count of matches. | `grep -c 'div' styles.css` |

## 



### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 





```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 



```

```





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



| Option | Description | Example(s) |
| --- | --- | --- |
| `-h` | Always hide the filenames in the output. | `grep -h 'div' styles.css Makefile` |
| `-H` | Always display the filenames in the output. | `grep -H 'div' styles.css Makefile` |
| `-l` | Only output the filenames where the pattern is matched. | `grep -l 'div' styles.css README.md` |
| `-L` | Only output the filenames where the pattern _doesn’t_ match. | `grep -L 'div' styles.css README.md` |

## 







### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



```

```

### 



| Option | Description | Example(s) |
| --- | --- | --- |
| `-A <value>` | Output `<value>` lines after the lines matching the pattern. | `grep -A 3 'div' styles.css` |
| `-B <value>` | Output `<value>` lines before the lines matching the pattern. | `grep -B 3 'div' styles.css` |
| `-C <value>` | Output `<value>` lines before and after the lines matching the pattern for a full context. | `grep -C 3 'div' styles.css` |



## 



### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





### 



```

```



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 



| Option | Description |
| --- | --- |
| `--exclude <value>` | Exclude the files `<value>`. Globs can be used here. |
| `--include <value>` | Include the files `<value>`. Globs can be used here. |

## 







```

```



```

```





```

```

## 







- 
- 
- 
- 
- 





```

```

```

```

```

```

```

```

```

```

```

```

```

```



## 



| CLI | Description |
| --- | --- |
| [ripgrep-all](https://github.com/phiresky/ripgrep-all) | Similar to ripgrep, except that you can parse many types of files: PDF, ebooks, office documents… |
| [ugrep](https://github.com/Genivia/ugrep) | Very fast grep-like CLI, apparently even faster than ripgrep. It also offers a TUI to search in your files. |



## 





- 
- 
- 
- 
- 
- 
- 





- 
- 

### 






