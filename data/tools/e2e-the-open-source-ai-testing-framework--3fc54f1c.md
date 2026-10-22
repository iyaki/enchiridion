---
title: "e2e: The open source AI testing framework"
notion_id: 3fc54f1c-7d23-81a9-b8e8-ef83be688551
notion_url: https://app.notion.com/p/e2e-The-open-source-AI-testing-framework-3fc54f1c7d2381a9b8e8ef83be688551
last_edited: 2026-10-17T02:57:00.000Z
source_url: https://tester.army/e2e
tags: ["Others", "Programming", "Testing", "TypeScript", "Automation", "Tool", "Medium"]
---
## Open Source AItesting framework

Customizable TypeScript testing framework for web, mobile apps and more.

import { test, expect } from 'e2e';

test('a member upgrades to Pro', async ({ app, agent, screen }) => {

await app.open('/settings/billing');

await agent.act('upgrade the workspace to the Pro plan');

await agent.assert('the invoice preview shows a prorated amount');

await expect(screen.getByRole('status')).toContainText('Pro');

});

[01] Why e2e

Most AI tests start from scratch.Every single time.

e2e remembers your agent's actionsand replays them when possible,without calling the model again.

Less repeated reasoning.Fewer tokens. Lower costs.

01/05

### Connect on your terms.

Use your preferred gateway and keep control
of how your AI connects.

import { gateway } from 'ai';

model: gateway('openai/gpt-6-luna-fast'),

02/05

### Your favorite AI? Bring it along.

Choose the model that works for you.
Your testing framework shouldn't decide for you.

import { gateway } from 'ai';

agents: {

default: {

model: gateway('openai/gpt-6-luna-fast'),

},

},

03/05

### Your agent. Your rules.

Give your agent a purpose. Define how it should approach
testing, what to look for, and what matters to you.

System prompt

agents: {

default: {

system: 'You are a QA testing agent.',

},

},

04/05

### Your app. Your environment.

Test on the web, go mobile, or bring your own engine.
One framework, built to adapt.

import { web } from '@e2e-dev/web';

export default {

targets: [{ engine: web(), app: { url: 'http://localhost:3000' } }],

agents: {

default: {

model: gateway('openai/gpt-6-luna-fast'),

},

},

} satisfies E2EConfig;

05/05

### Your stack. Ready to test.

Your model, your agent, your engine.
All connected through one testing framework.

import type { E2EConfig } from 'e2e';

import { web } from '@e2e-dev/web';

import { gateway } from 'ai';

export default {

targets: [{ engine: web(), app: { url: 'http://localhost:3000' } }],

agents: {

default: {

model: gateway('openai/gpt-6-luna-fast'),

system: 'You are a QA testing agent.',

},

},

} satisfies E2EConfig;

[03] Core features

## AI where it helps. Code where it counts.

### Don't make AI think twice.

Replay compatible cached actions without unnecessary model calls.

### Goals, not scripts.

Describe your testing goals in plain language. Let your agent handle the actions.

### Exact when it matters.

Combine AI-driven actions with precise locators and deterministic assertions.

### Your AI. Your choice.

Choose your preferred AI model and customize your testing agent.

### One API. Multiple environments.

Test web and mobile applications through one shared API.

### Bring your own engine.

Extend the framework with your own execution backend.

[04] Testimonials

## Loved by developers.

- [The future is verification-engineering.
Proofs, (e2e) tests, benchmarks, linters…
Some tests will be deterministic, some agentic. This looks great.(opens on X in a new tab)](https://x.com/rauchg/status/2105723481413550427)
- [The agentic mobile testing game just changed.(opens on X in a new tab)](https://x.com/expo/status/2105697228627247231)
- [We’ve all been building our version of this, but this is better. Well done!Jean-Michel Lemieux@jmwind(opens on X in a new tab)](https://x.com/jmwind/status/2105874398842605779)
- [e2e looks great!(opens on X in a new tab)](https://x.com/zeeg/status/2105798617403723885)
- [congrats on the launch! looks amazingEnzo Manuel Mangano@reactiive_(opens on X in a new tab)](https://x.com/reactiive_/status/2105704339725828103)
- [the future of end-to-end testing is partially agentic, partially deterministic. glad that libraries like e2e build in this area, supporting both mobile and web development from day one.
it was also a pleasure to collaborate with Oskar on making @agent_device, which is a driver behind mobile support, more reliable in deterministic flows.
let's go! write some tests!Michał Pierzchała@thymikee(opens on X in a new tab)](https://x.com/thymikee/status/2105753091890966561)
- [this is fire 🔥 congrats man(opens on X in a new tab)](https://x.com/vanshdevx/status/2106014512579252226)
- [Test your apps with X Premium or SuperGrok
@o_kwasniewski really cooked with this, you write tests in natural language, the agent uses tokens only once, it remembers what it did for all future runsMiłosz Jankiewicz@Jaaneek(opens on X in a new tab)](https://x.com/Jaaneek/status/2105682420607627445)
- [the mix of deterministic and agentic APIs sounds like a powerful approachBrjan | AI Builder@MrBrjan(opens on X in a new tab)](https://x.com/MrBrjan/status/2105679914913714653)
- [Esto tiene una pintaza espectacular!
He hecho 2-3 pruebas rapidas y wow... 👇🏼Translated from SpanishThis looks spectacular! I did 2-3 quick tests and wow...Carlos Beneyto@CarlosBeneyto(opens on X in a new tab)](https://x.com/CarlosBeneyto/status/2105740346374140145)
- [oh shit I was building this but this is better than the thing I was building(opens on X in a new tab)](https://x.com/cgenco/status/2105731949138743566)
- [yeah this is it. once agents write the code, the scarce skill is deciding what to prove vs what to agent-test. mixing both in one harness feels right(opens on X in a new tab)](https://x.com/idohodl/status/2105812419952161133)
- [verification loops are the future.
This massively enables that.(opens on X in a new tab)](https://x.com/TanayKmr/status/2105740428444324112)
- [it's fantastic. Just use it. trust me(opens on X in a new tab)](https://x.com/CalebPanza/status/2105727417029705780)
- [Can't wait to review this migration 🔥(opens on X in a new tab)](https://x.com/trozee_dev/status/2105676052471120226)
- [وجود حل يدمج بين العهد القديم والجديد في أختبار البرمجيات بشكل سلس كان تحدي خلال الشهور الماضية
ولكن الوضع اختلف مع e2e شي رهيب ان يقوم بعمل كاش للخطوات لتقليل التكلفه عند اعاده الاختبار مالم تتغير خطوات الاختبار
حل واعد جداً 💪Translated from ArabicSeamlessly combining the old and new approaches to software testing has been a challenge in recent months. That changed with e2e: it's great that it caches steps to cut the cost of re-running tests as long as the steps don't change. A very promising solution.Salah Alkhwlani@SalahAlkhwlani(opens on X in a new tab)](https://x.com/SalahAlkhwlani/status/2105787938563092506)
- [this looks interesting for reducing token usage when testing web and mobile apps.
it mixes deterministic and non deterministic testing, along with caching actions so you don’t have to call the model multiple timesElliott Callender@the__elliott(opens on X in a new tab)](https://x.com/the__elliott/status/2105913335594262589)
- [Needed this. Pure agentic e2e flakes hard in CI. Mixing deterministic assertions with agent steps, plus BYO model and infra, is the only setup that stays useful once agents actually drive real UIs.Edwin | AI Systems@EdwinAI_Systems(opens on X in a new tab)](https://x.com/EdwinAI_Systems/status/2105678137774555644)
- [npx e2e just saved me 0.60$ on this coffee order, thanks @TesterArmy 🙏
Works out of the box with @expo EAS Simulator 😎(opens on X in a new tab)](https://x.com/szdziedzic/status/2105748791999766942)
- [A probarlo! Me encanta este concepto mix de instrucciones deterministas y agenticas.
Lo mejor de todo es que tiene cache de los steps hechos por agentes.Translated from SpanishTime to try it! I love this mix of deterministic and agentic instructions. Best of all, it caches the steps the agents take.(opens on X in a new tab)](https://x.com/javicerodriguez/status/2105724821841576154)
- [congrats @TesterArmy team! you have really done some exceptional work with e2e, and i'm glad we got to be part of the launch 👏Danny Prevoznik@DanielPrevoznik(opens on X in a new tab)](https://x.com/DanielPrevoznik/status/2105676795999920286)
- [you have no idea how long I've been begging for this(opens on X in a new tab)](https://x.com/_kvnloo/status/2105863185509023808)
- [use e2e alongside @poteto's verification skill and you have a reliable workflow for mobile development.(opens on X in a new tab)](https://x.com/stringsaeed/status/2105734077085303106)
- [Love it!!! Testing it right now!Danny Suárez 🍄‍🟫 d/acc@dannysuarez(opens on X in a new tab)](https://x.com/dannysuarez/status/2105851280749162569)
- [Ya lo estoy probando pero toca tener cuidado con el consumo de tokens ando construyendo algo hibrido usando e2e y playwrightTranslated from SpanishAlready trying it, but you have to watch token usage. I'm building something hybrid with e2e and Playwright.Danny Suárez 🍄‍🟫 d/acc@dannysuarez(opens on X in a new tab)](https://x.com/dannysuarez/status/2106004279697817672)
- [Suena interesante, tengo que echarle un ojo.
Yo de momento resuelvo con Playwright generado por AI, a lo mejor es hora de cambiar...Translated from SpanishSounds interesting, I need to take a look. For now I get by with AI-generated Playwright, maybe it's time to switch...Carlos Castellanos@karloscodes(opens on X in a new tab)](https://x.com/karloscodes/status/2105910784903836042)
- [This is very interesting.
Most e2e tests either break on every UI tweak or hand the whole run to an agent that burns tokens redoing the same clicks.
e2e (TesterArmy) mixes both: stable locators where you know them, agent steps where you don’t - and caches the agent passes so the next run replays without calling the model again. Web, iOS, Android. BYO model.
Basically: Let AI figure it out once, and then replay without cost.(opens on X in a new tab)](https://x.com/hrobjartz/status/2105882275212575043)
- [Love this. I’ve been doing the same thing in a vacuum.(opens on X in a new tab)](https://x.com/dgaspard/status/2105782130051584380)
- [just started using it and i'm already in love w this. Demo coming soon🙂‍↕️Rami Maalouf@rami__maalouf(opens on X in a new tab)](https://x.com/rami__maalouf/status/2105681148722733423)
- [This is very cool! I was thinking about building something similar for @DispatchinoAI(opens on X in a new tab)](https://x.com/Amr0Khalid/status/2105759015003464011)
- [I like the action replay. An unchanged login screen shouldn’t need a fresh opinion every time CI runs.Ryan Gilpatric@RyanGilpatric(opens on X in a new tab)](https://x.com/RyanGilpatric/status/2105724758088393188)
- [This looks like magic 🪄 can’t wait to give it a goAntonie Huang@antoniehuang(opens on X in a new tab)](https://x.com/antoniehuang/status/2105718390841328028)
- [@TesterArmy E2E is amazing 🙌. It found a form submission case I haven't covered.
Basically a http 500 left the user with no error message during failure. Was easy to miss on my happy path tests.(opens on X in a new tab)](https://x.com/parzival1l/status/2106646701725364515)
- [wait the mix of deterministic + agentic apis is the bit tho. pure agent e2e always felt flaky in ci(opens on X in a new tab)](https://x.com/anytanreal/status/2105687026901168503)
- [playing around with e2e two days in the row.
it's amazing, the hybrid approach super useful!
kudos to @TesterArmy team(opens on X in a new tab)](https://x.com/dimasike_/status/2106776202979803604)

[05] FAQ

## Questions, answered.

The short version. When in doubt, the docs win.

[Read the docs](https://e2e.tester.army/docs)

Any AI SDK model works. Sign in with the ChatGPT, GitHub Copilot, or SuperGrok subscription you already pay for, bring a Vercel AI Gateway or OpenRouter key, or point the runner at a local server. The config selects the model; e2e adds no markup.

Secrets stay in environment variables: the config references named credentials and the runner fills them at run time, so values never live in your test code or transcripts.

Yes. @e2e-dev/mobile drives iOS simulators and Android emulators through agent-device, with the same tests, fixtures, and assertions you write for the web.

Yes. CI mode switches on automatically: one retry, a read-only trace cache, and agent steps replay verified actions without model calls when nothing changed. Pass your model key as a secret and run it like any other test job.

Yes: the docs include migration guides for Playwright, Cypress, Selenium, Detox, and Maestro. Locators and assertions map closely, and agent steps replace the most brittle parts first.

You get a report with the exact failing step, a screenshot, a trace, and logs under .e2e. Re-run with --headed or --debug to watch it live; every error code maps to a fix in the docs.
