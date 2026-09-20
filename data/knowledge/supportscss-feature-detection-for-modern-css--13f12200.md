---
title: "SupportsCSS - Feature Detection
 for Modern CSS"
notion_id: 13f12200-5812-4364-a1c4-8921243631eb
notion_url: https://app.notion.com/p/SupportsCSS-Feature-Detection-for-Modern-CSS-13f1220058124364a1c48921243631eb
last_edited: 2023-06-15T23:49:00.000Z
source_url: https://supportscss.dev/
tags: ["Frontend", "CSS", "Untried", "Framework/Library", "English"]
---




- 
- 
- 
- 

## 







### 











## 



- 
- 







| Feature Class | Global Name | Test Condition |
| --- | --- | --- |
| Supportedat-container | AtContainer | `window.CSSContainerRule` |
| Unsupportedat-container-style-properties | AtContainerStyleProperties | * [See explanation](https://supportscss.dev/#atcontainerstyleproperties-test) |
| Supportedat-counter-style | AtCounterStyle | `window.CSSCounterStyleRule` |
| Supportedat-layer | AtLayer | `window.CSSLayerBlockRule` |
| Supportedat-property | AtProperty | `window.CSSPropertyRule` |
| Unsupportedat-scope | AtScope | `window.CSSScopeRule` |
| Unsupportedanchor | Anchor | `CSS.supports('left: anchor(center)')` |
| Unsupportedcolor-function | ColorFunction | `CSS.supports('color: color(srgb 0 0 1)')` |
| Unsupportedcolor-mix | ColorMix | `CSS.supports('color: color-mix(in lch, white, black)')` |
| Supportedcontainer-units | ContainerUnits | `CSS.supports('width: 1cqi')` |
| Unsupporteddynamic-viewport-units | DynamicViewportUnits | `CSS.supports('width: 1dvi')` |
| Supportedhas | Has | `CSS.supports('selector(:has(+ *))')` (_Possible false positive in Firefox 112_) |
| Supportedhoudini-paint-api | HoudiniPaintApi | `window.CSS.paintWorklet` |
| Supportedindividual-transforms | IndividualTransforms | `CSS.supports('transform: scale(1)')` |
| Supportedlogical-properties | LogicalProperties | `CSS.supports('border-start-start-radius: 1px')` |
| Supportedmedia-range-syntax | MediaRangeSyntax | `window.matchMedia('(width >= 1px)')` |
| Unsupportednesting | Nesting | `CSS.supports('selector(& a)')` |
| Unsupportednth-of-s | NthOfS | `CSS.supports('selector(:nth-child(1 of .a))')` |
| Supportedoverscroll-behavior | OverscrollBehavior | `CSS.supports('overscroll-behavior: none')` |
| Unsupportedrelative-color-syntax | RelativeColorSyntax | `CSS.supports('color: rgb(from red r g b / 1%)')` |
| Unsupportedscroll-timeline | ScrollTimeline | `CSS.supports('scroll-timeline-name: a')` |
| Unsupportedsubgrid | Subgrid | `CSS.supports('grid-template-rows: subgrid')` |
| Unsupportedtext-box-trim | TextBoxTrim | `CSS.supports('(leading-trim: both) or (text-box-trim: both)')` |
| Unsupportedtrigonometry | Trigonometry | `CSS.supports('width: calc(1px * cos(1deg))')` |
| Unsupportedview-timeline | ViewTimeline | `window.ViewTimeline` |
| Unsupportedview-transitions | ViewTransitions | `window.ViewTransition` |

### 



- 
- 
- 
- 
- 

## 





### 



```

```



```

```

### 

```

```

### 



```

```



```

```



### 



- 
- 

## 



- 
- 
- 



```

```

## 



### 



### 



### 





```

```



### 







- 
- 
- 
- 



```

```


