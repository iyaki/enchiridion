---
title: "tradingview/lightweight-charts - Performant financial charts built with HTML5 canvas"
notion_id: fdfccde7-e186-4d17-8d0e-9c1142312b8f
notion_url: https://app.notion.com/p/tradingview-lightweight-charts-Performant-financial-charts-built-with-HTML5-canvas-fdfccde7e1864d178d0e9c1142312b8f
last_edited: 2024-06-05T18:35:00.000Z
source_url: https://github.com/tradingview/lightweight-charts
tags: ["Frontend", "HTML", "Untried", "Framework/Library", "English"]
---
![image](https://github.com/tradingview/lightweight-charts/raw/master/.github/logo.svg?sanitize=true)

# Lightweight Charts™

TradingView Lightweight Charts™ are one of the smallest and fastest financial HTML5 charts.

The Lightweight Charts™ library is the best choice for you if you want to display financial data as an interactive chart on your web page without affecting your web page loading speed and performance.

It is the best choice for you if you want to replace static image charts with interactive ones. The size of the library is close to static images but if you have dozens of image charts on a web page then using this library can make the size of your web page smaller.

## Installing

### es6 via npm

```shell
npm install lightweight-charts
```

```javascript
import { createChart } from 'lightweight-charts';

const chart = createChart(document.body, { width: 400, height: 300 });
const lineSeries = chart.addLineSeries();
lineSeries.setData([
    { time: '2019-04-11', value: 80.01 },
    { time: '2019-04-12', value: 96.63 },
    { time: '2019-04-13', value: 76.64 },
    { time: '2019-04-14', value: 81.89 },
    { time: '2019-04-15', value: 74.43 },
    { time: '2019-04-16', value: 80.01 },
    { time: '2019-04-17', value: 96.63 },
    { time: '2019-04-18', value: 76.64 },
    { time: '2019-04-19', value: 81.89 },
    { time: '2019-04-20', value: 74.43 },
]);
```

### CDN

You can use [unpkg](https://unpkg.com/):

[https://unpkg.com/lightweight-charts/dist/lightweight-charts.standalone.production.js](https://unpkg.com/lightweight-charts/dist/lightweight-charts.standalone.production.js)

The standalone version creates `window.LightweightCharts` object with all exports from `esm` version:

```javascript
const chart = LightweightCharts.createChart(document.body, { width: 400, height: 300 });
const lineSeries = chart.addLineSeries();
lineSeries.setData([
    { time: '2019-04-11', value: 80.01 },
    { time: '2019-04-12', value: 96.63 },
    { time: '2019-04-13', value: 76.64 },
    { time: '2019-04-14', value: 81.89 },
    { time: '2019-04-15', value: 74.43 },
    { time: '2019-04-16', value: 80.01 },
    { time: '2019-04-17', value: 96.63 },
    { time: '2019-04-18', value: 76.64 },
    { time: '2019-04-19', value: 81.89 },
    { time: '2019-04-20', value: 74.43 },
]);
```

### Build Variants

| Dependencies included | Mode | ES module | CommonJS ⚠️ | IIFE (`window.LightweightCharts`) |
| --- | --- | --- | --- | --- |
| No | PROD | `lightweight-charts.production.mjs` | `lightweight-charts.production.cjs` | N/A |
| No | DEV | `lightweight-charts.development.mjs` | `lightweight-charts.development.cjs` | N/A |
| Yes (standalone) | PROD | `lightweight-charts.standalone.production.mjs` | - | `lightweight-charts.standalone.production.js` |
| Yes (standalone) | DEV | `lightweight-charts.standalone.development.mjs` | - | `lightweight-charts.standalone.development.js` |

⚠️ **Deprecation note:** CommonJS support will be removed from the library at the start of 2024.

## Development

See [BUILDING.md](https://github.com/tradingview/lightweight-charts/blob/master/BUILDING.md) for instructions on how to build `lightweight-charts` from source.

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this software except in compliance with the License. You may obtain a copy of the License at LICENSE file. Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

This software incorporates several parts of tslib ([https://github.com/Microsoft/tslib](https://github.com/Microsoft/tslib), (c) Microsoft Corporation) that are covered by BSD Zero Clause License.

This license requires specifying TradingView as the product creator. You shall add the "attribution notice" from the NOTICE file and a link to [https://www.tradingview.com/](https://www.tradingview.com/) to the page of your website or mobile application that is available to your users. As thanks for creating this product, we'd be grateful if you add it in a prominent place.
