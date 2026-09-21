---
title: "Testcafe - e2e testing"
notion_id: 084b904c-b69f-4d0d-b657-f8ed5ec46cb3
notion_url: https://app.notion.com/p/Testcafe-e2e-testing-084b904cb69f4d0db657f8ed5ec46cb3
last_edited: 2026-09-21T17:08:00.000Z
source_url: https://testcafe.io/
tags: ["Web Development", "Testing", "Tool", "English"]
---
## From zero to testing in minutes

- 1% npm i -g testcafeJust one npm package.
- 2% testcafe chrome test.jsWorks with common browsers out of the box.
- 3% docker pull testcafe/testcafeReady for your CI/CD pipeline.

Get Started

## Write or Record Tests

- Get started with our free and open source framework in minutes.
- Create easy-to-read JavaScript and TypeScript tests.
- Simulate complex page interactions and multi-window scenarios.
- Code your tests by hand or record in your browser.

## Run and Analyze

- Run your tests in any modern browser — local or remote.
- Run your tests concurrently to speed up the testing process.
- Easily integrate TestCafe with your CI solution of choice.
- Store test reports in many convenient formats.

## Write tests with ease

The intuitive syntax of TestCafe makes teams more productive from day one.Check the test below and see for yourself.

Select a tab

TestCafe 😊

```plain text
fixture('Pizza Palace') .page('https://testcafe-demo-page.glitch.me/'); test('Submit a form', async t => { await t // automatically dismiss dialog boxes .setNativeDialogHandler(() => true) // drag the pizza size slider .drag('.noUi-handle', 100, 0) // select the toppings .click('.next-step') .click('label[for="pepperoni"]') .click('#step2 .next-step') // fill the address form .click('.confirm-address') .typeText('#phone-input', '+1-541-754-3001') .click('#step3 .next-step') // zoom into the iframe map .switchToIframe('.restaurant-location iframe') .click('button[title="Zoom in"]') // submit the order .switchToMainWindow() .click('.complete-order'); });
```

Selenium (JavaScript) 😕

```plain text
const {Builder, By, Key, until} = require('selenium-webdriver'); (async function pizzaPalace() { const driver = await new Builder().forBrowser('firefox').build(); try { await driver.get('https://testcafe-demo-page.glitch.me/'); // drag the pizza size slider const sourceEle = driver.findElement(By.className("noUi-handle")); const actions = driver.actions({async: true}); await actions.dragAndDrop(sourceEle, {x:100, y:0}).perform(); // select the toppings await driver.findElement(By.className("next-step")).click(); await driver.findElement(By.css('label[for="pepperoni"]')).click(); await driver.findElement(By.css('#step2 .next-step')).click(); // fill the address form await driver.wait(until.elementLocated(By.css('.google-map')),10000); await driver.findElement(By.className("confirm-address")).click(); await driver.findElement(By.id('phone-input')).sendKeys('+1-541-754-3001'); await driver.findElement(By.css('#step3 .next-step')).click(); // zoom into the iframe map await driver.wait(until.elementLocated(By.css('.restaurant-location iframe')),10000); const restaurantLocationFrame = await driver.findElement(By.css('.restaurant-location iframe')); await driver.switchTo().frame(restaurantLocationFrame); await driver.wait(until.elementLocated(By.css('button[title="Zoom in"]')),10000); await driver.findElement(By.css('button[title="Zoom in"]')).click(); // submit the order await driver.switchTo().defaultContent(); await driver.findElement(By.className("complete-order")).click(); // dismiss the confirmation dialog await driver.wait(until.alertIsPresent()); const confirmationMessage = driver.switchTo().alert(); await confirmationMessage.accept(); } finally { await driver.quit(); } })();
```

## Deploy without fear

- CI/CD-readyTestCafe integrates with all popular CI/CD solutions.
- Concurrent Test RunsSave time and resources: run your tests in multiple browsers at once.
- If something goes wrong...Use the built-in Debug Mode to pin-point the source of your frustration.

Get Started
