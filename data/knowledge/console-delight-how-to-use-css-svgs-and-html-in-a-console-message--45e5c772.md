---
title: "console.delight, how to use CSS, SVGs, and HTML in a console message"
notion_id: 45e5c772-1ce1-4539-be22-738711a78f92
notion_url: https://app.notion.com/p/console-delight-how-to-use-CSS-SVGs-and-HTML-in-a-console-message-45e5c7721ce14539be22738711a78f92
last_edited: 2024-03-11T14:18:00.000Z
source_url: https://frontendmasters.com/blog/console-delight/
tags: ["Article", "Frontend Masters Blog", "English", "Javascript", "Web Development"]
---
![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/console-delight-thumb.jpg?w=1000&ssl=1)





## 











```

```



![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/console-info-chrome_iazb3v.webp?w=782&ssl=1)









```

```

![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/photo_zuevwp.webp?w=692&ssl=1)





## 

### 



| Capability | Supported? |
| --- | --- |
| Any “regular” SVG element (e.g. `<circle>`, `<path>`, etc) | ✅ |
| Gradients | ✅ |
| Clip paths | ✅ |
| Masks | ✅ |
| Filters | ✅ |
| Transforms | ✅[1](https://frontendmasters.com/blog/console-delight/?utm_source=tldrwebdev#d3e35373-b7fa-40a0-abe1-de55b75a23c2) |
| `<a>` tag | ❌ |
| External images | ❌ |
| SMIL animation | ✅ |
| `foreignObject` | ✅[2](https://frontendmasters.com/blog/console-delight/?utm_source=tldrwebdev#b456d489-1a38-4839-b543-b60424b8b4c6) |
| Patterns | ✅ |

### 



```

```



```

```



```

```

### 

| Capability | Supported? |
| --- | --- |
| CSS animations | ✅ |
| `:hover` | ❌ |
| CSS variables | ✅[3](https://frontendmasters.com/blog/console-delight/?utm_source=tldrwebdev#6d1c8864-8210-4c56-b14d-1efb64fc6987) |
| `@media` queries | ✅[4](https://frontendmasters.com/blog/console-delight/?utm_source=tldrwebdev#3392d0fc-fce0-4f23-a949-f61bf3fcf3d4) |
| Viewport units | ✅[5](https://frontendmasters.com/blog/console-delight/?utm_source=tldrwebdev#d10d6186-dd40-452c-a382-0821d7722689) |
| `background-image: linear-gradient()` | ✅ |
| `background-image: url()` | ❌ |
| `@import` | ❌ |
| `background-clip` | ❌ |

### 

| Capability | Supported? |
| --- | --- |
| Console commands | ❌ |
| DOM references within the SVG | ✅ |
| DOM reference to the SVG itself | ✅ |
| DOM reference outside of the SVG | ❌ |
| `.appendChild` | ❌ |
| Changing the SVG’s DOM (attributes, text, etc.) | ✅ |
| `setTimeout` | ❌[6](https://frontendmasters.com/blog/console-delight/?utm_source=tldrwebdev#621356bc-ae4e-4d36-8dba-c6293f976027) |
| `requestAnimationFrame` | ❌[6](https://frontendmasters.com/blog/console-delight/?utm_source=tldrwebdev#621356bc-ae4e-4d36-8dba-c6293f976027) |
| WAAPI | ❌ |
| “Global” variables via multiple `<script>` tags | ✅ |
| Setting CSS variables | ✅ |
| Importing external JS | ❌ |
| `prompt` | ❌ |
| `.addEventListener` | ❌ |

## 









```

```

## 







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

![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/datetime_zzmiav.webp?w=926&ssl=1)





```

```



```

```

## 





```

```



![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/link-chrome_drdw37.webp?w=1580&ssl=1)







## 







## 









## 









## 





### 



### 



![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/console-info-firefox_hgoffa.webp?w=768&ssl=1)





```

```



![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/console-info-firefox-in-chrome_wgjdme.webp?w=792&ssl=1)









### 







## 

- 
- 
- 
- 
- 

## 





- 
- 
- 
- 



```

```

## 



```

```

![image](https://i0.wp.com/frontendmasters.com/blog/wp-content/uploads/2024/01/gradient-text_hkv18l.webp?w=1106&ssl=1)



## 



```

```





## 



## 



1. 
2. 
3. 
4. 
5. 
6. 
