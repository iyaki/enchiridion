---
title: "TinyBase - The reactive data store for local-first apps"
notion_id: 01d24f88-e550-497c-9f7b-34969747e3c0
notion_url: https://app.notion.com/p/TinyBase-The-reactive-data-store-for-local-first-apps-01d24f88e550497c9f7b34969747e3c0
last_edited: 2024-08-16T19:42:00.000Z
source_url: https://tinybase.org/
tags: ["English", "Web Development", "Databases", "Frontend", "Untried", "Framework/Library"]
---
![image](https://tinybase.org/favicon.svg)

## The _reactive_ data store for local-first apps.

[_NEW!_](https://tinybase.org/guides/releases/#v5-0)[ v5.0 release](https://tinybase.org/guides/releases/#v5-0) "The One You Can Sync"

[Get started](https://tinybase.org/guides/the-basics/getting-started/)

[Try the demos](https://tinybase.org/demos/)

[Read the docs](https://tinybase.org/api/store/interfaces/store/store/)

## It's _Reactive_

TinyBase lets you [listen to changes](https://tinybase.org/#register-granular-listeners) made to any part of your data. This means your app will be fast, since you only spend rendering cycles on things that change. The optional [bindings to React](https://tinybase.org/#call-hooks-to-bind-to-data) and [pre-built components](https://tinybase.org/#pre-built-reactive-components) let you easily build fully reactive UIs on top of TinyBase. You even get a built-in [undo stack](https://tinybase.org/#set-checkpoints-for-an-undo-stack), and [developer tools](https://tinybase.org/#an-inspector-for-your-data)!

## It's _Database-Like_

[`Store`](https://tinybase.org/api/store/interfaces/store/store/) [key-value data](https://tinybase.org/#start-with-a-simple-key-value-store) and [tabular data](https://tinybase.org/#level-up-to-use-tabular-data) with optional typed [schematization](https://tinybase.org/#apply-schemas-to-tables-values) to model your app's data structures. TinyBase provides built-in [indexing](https://tinybase.org/#create-indexes-for-fast-lookups), [metric aggregation](https://tinybase.org/#define-metrics-and-aggregations), and tabular [relationships](https://tinybase.org/#model-table-relationships) APIs - or the powerful [query engine](https://tinybase.org/#build-complex-queries-with-tinyql) to select, join, filter, and group data (reactively!) without SQL.

## It _Synchronizes_

TinyBase is an in-memory data store, but you can easily [persist](https://tinybase.org/#persist-to-storage-sqlite-more) your data to [browser storage](https://tinybase.org/api/persister-browser), [IndexedDB](https://tinybase.org/api/persister-indexed-db), many flavors of [database](https://tinybase.org/guides/persistence/database-persistence/), and [more](https://tinybase.org/guides/persistence/third-party-crdt-persistence/). TinyBase now has [native CRDT](https://tinybase.org/#synchronize-between-devices) support, meaning that you can [natively synchronize](https://tinybase.org/guides/synchronization/) and merge data across multiple sources and clients.

## It's Built For A _Local-First_ World

TinyBase works anywhere that JavaScript does, but it's especially great for local-first apps: where data is stored locally on the user's device and that can be run offline. It's tiny by name, tiny by nature: just [5.3kB - 12.7kB](https://tinybase.org/#did-we-say-tiny) and with no dependencies - yet [100% tested](https://tinybase.org/#well-tested-and-documented), [fully documented](https://tinybase.org/guides/the-basics/getting-started/), and of course, [open source](https://github.com/tinyplex/tinybase)!

## TinyBase works great on its own, but also plays well with friends!

[React](https://tinybase.org/guides/building-uis/getting-started-with-ui-react)

[PartyKit](https://tinybase.org/api/persister-partykit-client)

[Expo SQLite](https://tinybase.org/guides/schemas-and-persistence/database-persistence)

[ElectricSQL](https://tinybase.org/guides/schemas-and-persistence/database-persistence)

[SQLite](https://tinybase.org/guides/schemas-and-persistence/database-persistence)

[Turso](https://tinybase.org/guides/schemas-and-persistence/database-persistence)

[PowerSync](https://tinybase.org/guides/schemas-and-persistence/database-persistence)

[IndexedDB](https://tinybase.org/api/persister-indexed-db/functions/creation/createindexeddbpersister)

[YJS](https://tinybase.org/api/persister-yjs/functions/creation/createyjspersister)

[CR-SQLite](https://tinybase.org/api/persister-cr-sqlite-wasm)

[Automerge](https://tinybase.org/api/persister-automerge)

## Start with a simple key-value store.

Creating a [`Store`](https://tinybase.org/api/store/interfaces/store/store/) requires just a simple call to the [`createStore`](https://tinybase.org/api/store/functions/creation/createstore/) function. Once you have one, you can easily set [`Values`](https://tinybase.org/api/store/type-aliases/store/values/) in it by unique [`Id`](https://tinybase.org/api/common/type-aliases/identity/id/). And of course you can easily get them back out again.

Read more about using keyed value data in [The Basics](https://tinybase.org/guides/the-basics/) guide.

```javascript
import {createStore} from 'tinybase';

const store =createStore()
  .setValues({employees: 3})
  .setValue('open', true);

console.log(store.getValues());
// -> {employees: 3, open: true}
```

## Level up to use tabular data.

For other types of data applications, a tabular data structure is more useful. TinyBase lets you set and get nested [`Table`](https://tinybase.org/api/store/type-aliases/store/table/), [`Row`](https://tinybase.org/api/store/type-aliases/store/row/), or [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/) data, by unique [`Id`](https://tinybase.org/api/common/type-aliases/identity/id/) - and in the same [`Store`](https://tinybase.org/api/store/interfaces/store/store/) as the keyed values!

Read more about setting and changing data in [The Basics](https://tinybase.org/guides/the-basics/) guide.

```javascript
store
  .setTable('pets', {fido: {species: 'dog'}})
  .setCell('pets', 'fido', 'color', 'brown');

console.log(store.getRow('pets', 'fido'));
// -> {species: 'dog', color: 'brown'}
```

## Register granular listeners.

The magic starts to happen when you register listeners on a [`Value`](https://tinybase.org/api/store/type-aliases/store/value/), [`Table`](https://tinybase.org/api/store/type-aliases/store/table/), [`Row`](https://tinybase.org/api/store/type-aliases/store/row/), or [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/). They get called when any part of that object changes. You can also use wildcards - useful when you don't know the [`Id`](https://tinybase.org/api/common/type-aliases/identity/id/) of the objects that might change.

Read more about listeners in the [Listening To Stores](https://tinybase.org/guides/the-basics/listening-to-stores/) guide.

```javascript
const listenerId = store.addTableListener('pets', () =>
  console.log('changed'),
);

store.setCell('pets', 'fido', 'sold', false);
// -> 'changed'

store.delListener(listenerId);
```

## Call hooks to bind to data.

If you're using React in your application, the optional [`ui-react`](https://tinybase.org/api/ui-react/) module provides hooks to bind to the data in a [`Store`](https://tinybase.org/api/store/interfaces/store/store/).

More magic! The [`useCell`](https://tinybase.org/api/ui-react/functions/store-hooks/usecell/) hook in this example fetches the dog's color. But it also registers a listener on that cell that will fire and re-render the component whenever the value changes.

Basically you simply describe what data you want in your user interface and TinyBase will take care of the whole lifecycle of updating it for you.

Read more about the using hooks in the [Using React Hooks](https://tinybase.org/guides/building-uis/using-react-hooks/) guide.

```javascript
import React from 'react';
import {createRoot} from 'react-dom/client';
import {useCell} from 'tinybase/ui-react';

const App1 = () => {
  const color =useCell('pets', 'fido', 'color', store);
  return <>Color: {color}</>;
};

const app = document.createElement('div');
const root = createRoot(app);
root.render(<App1 />);
console.log(app.innerHTML);
// -> 'Color: brown'

store.setCell('pets', 'fido', 'color', 'walnut');
console.log(app.innerHTML);
// -> 'Color: walnut'
```

## Pre-built reactive components.

The [`ui-react`](https://tinybase.org/api/ui-react/) module provides bare React components that let you build up a fully reactive user interface based on a [`Store`](https://tinybase.org/api/store/interfaces/store/store/).

For web applications in particular, the new [`ui-react-dom`](https://tinybase.org/api/ui-react-dom/) module provides pre-built components for tabular display of your data, with lots of customization and interactivity options.

Try them out in the [UI Components](https://tinybase.org/demos/ui-components/) demos, and read more about the underlying [`ui-react`](https://tinybase.org/api/ui-react/) module in the [Building UIs](https://tinybase.org/guides/building-uis/) guides.

## An inspector for your data.

If you are building a web application, the new [`Inspector`](https://tinybase.org/api/ui-react-inspector/functions/development-components/inspector/) component lets you overlay a view of the data in your [`Store`](https://tinybase.org/api/store/interfaces/store/store/), [`Indexes`](https://tinybase.org/api/indexes/interfaces/indexes/indexes/), [`Relationships`](https://tinybase.org/api/relationships/interfaces/relationships/relationships/), and so on. You can even edit the data in place and see it update in your app immediately.

Read more about this powerful new tool in the [Inspecting Data](https://tinybase.org/guides/developer-tools/inspecting-data/) guide.

## Apply schemas to tables & values.

By default, a [`Store`](https://tinybase.org/api/store/interfaces/store/store/) can contain any arbitrary [`Value`](https://tinybase.org/api/store/type-aliases/store/value/), and a [`Row`](https://tinybase.org/api/store/type-aliases/store/row/) can contain any arbitrary [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/). But you can add a [`ValuesSchema`](https://tinybase.org/api/store/type-aliases/schema/valuesschema/) or a [`TablesSchema`](https://tinybase.org/api/store/type-aliases/schema/tablesschema/) to a [`Store`](https://tinybase.org/api/store/interfaces/store/store/) to ensure that the values are always what you expect: constraining their types, and providing defaults.

In this example, we set a new [`Row`](https://tinybase.org/api/store/type-aliases/store/row/) without the `sold` [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/) in it. The schema ensures it's present with default of `false`.

Read more about schemas in the [Schemas](https://tinybase.org/guides/schemas/) guide.

```javascript
store.setTablesSchema({
  pets: {
    species: {type: 'string'},
    color: {type: 'string'},
    sold: {type: 'boolean', default: false},
  },
});

store.setRow('pets', 'polly', {species: 'parrot'});
console.log(store.getRow('pets', 'polly'));
// -> {species: 'parrot', sold: false}

store.delTablesSchema();
```

## Synchronize between devices.

The [`MergeableStore`](https://tinybase.org/api/mergeable-store/interfaces/mergeable/mergeablestore/) type acts as a native CRDT, letting you merge data and synchronize it between clients and systems. The synchronization protocol can run over WebSockets, the browser BroadcastChannel, or your own custom synchronization medium.

Read more about these techniques in the [Synchronization](https://tinybase.org/guides/synchronization/) guides.

```javascript
import {WebSocketServer, WebSocket} from 'ws';
import {createMergeableStore} from 'tinybase';
import {createWsServer} from 'tinybase/synchronizers/synchronizer-ws-server';
import {createWsSynchronizer} from 'tinybase/synchronizers/synchronizer-ws-client';

// On a server machine:
const server =createWsServer(
  new WebSocketServer({port: 8040}),
);

// On a client machine:
const store1 =createMergeableStore();
const synchronizer1 = awaitcreateWsSynchronizer(
  store1,
  new WebSocket('ws://localhost:8040'),
);
await synchronizer1.startSync();

// ...

synchronizer1.destroy();
server.destroy();
```

## Persist to storage, SQLite, & more.

You can easily persist a [`Store`](https://tinybase.org/api/store/interfaces/store/store/) between browser page reloads or sessions. You can also synchronize it with a web endpoint, or (if you're using TinyBase in an appropriate environment), load and save it to a file. You can bind TinyBase to various flavors of [SQLite](https://tinybase.org/guides/schemas-and-persistence/database-persistence/), or to [Yjs](https://yjs.dev/) and [Automerge](https://automerge.org/) CRDT documents.

Read more about persisters in the [Persistence](https://tinybase.org/guides/persistence/) guides.

```javascript
import {createSessionPersister} from 'tinybase/persisters/persister-browser';

const persister =createSessionPersister(store, 'demo');
await persister.save();

console.log(sessionStorage.getItem('demo'));
// ->
`
[
  {
    "pets":{
      "fido":{"species":"dog","color":"walnut","sold":false},
      "polly":{"species":"parrot","sold":false}
    }
  },
  {"employees":3,"open":true}
]
`;

persister.destroy();
sessionStorage.clear();
```

## Build complex queries with [TinyQL](https://tinybase.org/guides/using-queries/tinyql/).

The [`Queries`](https://tinybase.org/api/queries/interfaces/queries/queries/) object lets you query data across tables, with filtering and aggregation - using a SQL-adjacent syntax called [TinyQL](https://tinybase.org/guides/using-queries/tinyql/).

Accessors and listeners let you sort and paginate the results efficiently, making building rich tabular interfaces easier than ever.

In this example, we have two tables: of pets and their owners. They are joined together by the pet's ownerId [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/). We select the pet's species, and the owner's state, and then aggregate the prices for the combinations.

We access the results by descending price, essentially answering the question: "which is the highest-priced species, and in which state?"

Needless to say, the results are reactive too! You can add listeners to queries just as easily as you do to raw tables.

Read more about [`Queries`](https://tinybase.org/api/queries/interfaces/queries/queries/) in the [v2.0 Release Notes](https://tinybase.org/guides/releases/#v2-0), the [Using Queries](https://tinybase.org/guides/using-queries/) guide, and the [Car Analysis](https://tinybase.org/demos/car-analysis/) demo and [Movie Database](https://tinybase.org/demos/movie-database/) demo.

```javascript
import {createQueries} from 'tinybase';

store
  .setTable('pets', {
    fido: {species: 'dog', ownerId: '1', price: 5},
    rex: {species: 'dog', ownerId: '2', price: 4},
    felix: {species: 'cat', ownerId: '2', price: 3},
    cujo: {species: 'dog', ownerId: '3', price: 4},
  })
  .setTable('owners', {
    1: {name: 'Alice', state: 'CA'},
    2: {name: 'Bob', state: 'CA'},
    3: {name: 'Carol', state: 'WA'},
  });

const queries =createQueries(store);
queries.setQueryDefinition(
  'prices',
  'pets',
  ({select, join, group}) => {
    select('species');
    select('owners', 'state');
    select('price');
    join('owners', 'ownerId');
    group('price', 'avg').as('avgPrice');
  },
);

queries
  .getResultSortedRowIds('prices', 'avgPrice', true)
  .forEach((rowId) => {
    console.log(queries.getResultRow('prices', rowId));
  });
// -> {species: 'dog', state: 'CA', avgPrice: 4.5}
// -> {species: 'dog', state: 'WA', avgPrice: 4}
// -> {species: 'cat', state: 'CA', avgPrice: 3}

queries.destroy();
```

## Define metrics and aggregations.

A [`Metrics`](https://tinybase.org/api/metrics/interfaces/metrics/metrics/) object makes it easy to keep a running aggregation of [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/) values in each [`Row`](https://tinybase.org/api/store/type-aliases/store/row/) of a [`Table`](https://tinybase.org/api/store/type-aliases/store/table/). This is useful for counting rows, but also supports averages, ranges of values, or arbitrary aggregations.

In this example, we create a new table of the pet species, and keep a track of which is most expensive. When we add horses to our pet store, the listener detects that the highest price has changed.

Read more about [`Metrics`](https://tinybase.org/api/metrics/interfaces/metrics/metrics/) in the [Using Metrics](https://tinybase.org/guides/using-metrics/) guide.

```javascript
import {createMetrics} from 'tinybase';

store.setTable('species', {
  dog: {price: 5},
  cat: {price: 4},
  worm: {price: 1},
});

const metrics =createMetrics(store);
metrics.setMetricDefinition(
  'highestPrice', // metricId
  'species', //      tableId to aggregate
  'max', //          aggregation
  'price', //        cellId to aggregate
);

console.log(metrics.getMetric('highestPrice'));
// -> 5

metrics.addMetricListener('highestPrice', () =>
  console.log(metrics.getMetric('highestPrice')),
);
store.setCell('species', 'horse', 'price', 20);
// -> 20

metrics.destroy();
```

## Create indexes for fast lookups.

An [`Indexes`](https://tinybase.org/api/indexes/interfaces/indexes/indexes/) object makes it easy to look up all the [`Row`](https://tinybase.org/api/store/type-aliases/store/row/) objects that have a certain value in a [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/).

In this example, we create an index on the `species` [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/) values. We can then get the the list of distinct [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/) value present for that index (known as 'slices'), and the set of [`Row`](https://tinybase.org/api/store/type-aliases/store/row/) objects that match each value.

[`Indexes`](https://tinybase.org/api/indexes/interfaces/indexes/indexes/) objects are reactive too. So you can set listeners on them just as you do for the data in the underlying [`Store`](https://tinybase.org/api/store/interfaces/store/store/).

Read more about [`Indexes`](https://tinybase.org/api/indexes/interfaces/indexes/indexes/) in the [Using Indexes](https://tinybase.org/guides/using-indexes/) guide.

```javascript
import {createIndexes} from 'tinybase';

const indexes =createIndexes(store);
indexes.setIndexDefinition(
  'bySpecies', // indexId
  'pets', //      tableId to index
  'species', //   cellId to index
);

console.log(indexes.getSliceIds('bySpecies'));
// -> ['dog', 'cat']
console.log(indexes.getSliceRowIds('bySpecies', 'dog'));
// -> ['fido', 'rex', 'cujo']

indexes.addSliceIdsListener('bySpecies', () =>
  console.log(indexes.getSliceIds('bySpecies')),
);
store.setRow('pets', 'lowly', {species: 'worm'});
// -> ['dog', 'cat', 'worm']

indexes.destroy();
```

## Model table relationships.

A [`Relationships`](https://tinybase.org/api/relationships/interfaces/relationships/relationships/) object lets you associate a [`Row`](https://tinybase.org/api/store/type-aliases/store/row/) in a local [`Table`](https://tinybase.org/api/store/type-aliases/store/table/) with the [`Id`](https://tinybase.org/api/common/type-aliases/identity/id/) of a [`Row`](https://tinybase.org/api/store/type-aliases/store/row/) in a remote [`Table`](https://tinybase.org/api/store/type-aliases/store/table/). You can also reference a table to itself to create linked lists.

In this example, the `species` [`Cell`](https://tinybase.org/api/store/type-aliases/store/cell/) of the `pets` [`Table`](https://tinybase.org/api/store/type-aliases/store/table/) is used to create a relationship to the `species` [`Table`](https://tinybase.org/api/store/type-aliases/store/table/), so that we can access the price of a given pet.

Like everything else, you can set listeners on [`Relationships`](https://tinybase.org/api/relationships/interfaces/relationships/relationships/) too.

Read more about [`Relationships`](https://tinybase.org/api/relationships/interfaces/relationships/relationships/) in the [Using Relationships](https://tinybase.org/guides/using-relationships/) guide.

```javascript
import {createRelationships} from 'tinybase';

const relationships =createRelationships(store);
relationships.setRelationshipDefinition(
  'petSpecies', // relationshipId
  'pets', //       local tableId to link from
  'species', //    remote tableId to link to
  'species', //    cellId containing remote key
);

console.log(
  store.getCell(
    relationships.getRemoteTableId('petSpecies'),
    relationships.getRemoteRowId('petSpecies', 'fido'),
    'price',
  ),
);
// -> 5

relationships.destroy();
```

## Set checkpoints for an undo stack.

A [`Checkpoints`](https://tinybase.org/api/checkpoints/interfaces/checkpoints/checkpoints/) object lets you set checkpoints on a [`Store`](https://tinybase.org/api/store/interfaces/store/store/). Move forward and backward through them to create undo and redo functions.

In this example, we set a checkpoint, then sell one of the pets. Later, the pet is brought back to the shop, and we go back to that checkpoint to revert the store to its previous state.

Read more about [`Checkpoints`](https://tinybase.org/api/checkpoints/interfaces/checkpoints/checkpoints/) in the [Using Checkpoints](https://tinybase.org/guides/using-checkpoints/) guide.

```javascript
import {createCheckpoints} from 'tinybase';

const checkpoints =createCheckpoints(store);

store.setCell('pets', 'felix', 'sold', false);
checkpoints.addCheckpoint('pre-sale');

store.setCell('pets', 'felix', 'sold', true);
console.log(store.getCell('pets', 'felix', 'sold'));
// -> true

checkpoints.goBackward();
console.log(store.getCell('pets', 'felix', 'sold'));
// -> false
```

## Type definitions & ORM-like APIs

TinyBase has comprehensive type definitions, and even offers definitions that infer API types from the data schemas you apply.

Furthermore, you can easily create TypeScript `.d.ts` definitions that model your data and encourage type-safety when reading and writing data - as well as `.ts` implementations that provide ORM-like methods for your named tables.

Read more about type support in the [TinyBase And TypeScript](https://tinybase.org/guides/the-basics/tinybase-and-typescript/) guide.

```typescript
const tools =createTools(store);
const [dTs, ts] = tools.getStoreApi('shop');

// -- shop.d.ts --
/* Represents the 'pets' Table. */
export type PetsTable = {[rowId: Id]: PetsRow};
/* Represents a Row when getting the content of the 'pets' Table. */
export type PetsRow = {species: string /* ... */};
//...

// -- shop.ts --
export const createShop: typeof createShopDecl = () => {
  //...
};
```

## Did we say tiny?

If you use the basic [`store`](https://tinybase.org/api/store/) module alone, you'll only add a gzipped _5.3kB_ to your app. Incrementally add the other modules as you need more functionality, or get it all for _12.7kB_.

The optional [`ui-react`](https://tinybase.org/api/ui-react/) module is just _4.6kB_, the ui-react-dom components are another _2.5kB_, and everything is super fast. Life's easy when you have zero dependencies!

Read more about how TinyBase is structured and packaged in the [Architecture](https://tinybase.org/guides/how-tinybase-is-built/architecture/) guide.

|  | Minified .js.gz | Source .js |
| --- | --- | --- |
| [tinybase/store](https://tinybase.org/api/store/) (minimal) | 5.3kB | 52.3kB |
| tinybase (complete) | 12.7kB | 129.8kB |
| [ui-react](https://tinybase.org/api/ui-react/) | 4.6kB | 48.3kB |
| [ui-react-dom](https://tinybase.org/api/ui-react-dom/) | 2.5kB | 20.4kB |

## Well tested and documented.

TinyBase has _100.0%_ test coverage, including the code throughout the documentation - even on this page! The guides, demos, and API examples are designed to make it as easy as possible for you to get your TinyBase-powered app up and running.

Read more about how TinyBase is tested in the Unit [Testing](https://tinybase.org/guides/how-tinybase-is-built/testing/) guide.

|  | Total | Tested | Coverage |
| --- | --- | --- | --- |
| Lines | 2,373 | 2,373 | 100.0% |
| Statements | 2,561 | 2,561 | 100.0% |
| Functions | 1,004 | 1,004 | 100.0% |
| Branches | 947 | 947 | 100.0% |
| Tests | 6,821 |  |  |
| Assertions | 32,043 |  |  |

## Proud to be sponsored by:

## Excited to be used by:

[Get started](https://tinybase.org/guides/the-basics/getting-started/)

[Try the demos](https://tinybase.org/demos/)

[Read the docs](https://tinybase.org/api/store/interfaces/store/store/)

## About

Modern apps deserve better. Why trade reactive user experiences to be able to use relational data? Or sacrifice features for bundle size? And why does the cloud do all the work [anyway](https://localfirstweb.dev/)?

Building TinyBase was originally an interesting exercise for [me](https://tripleodeon.com/) in API design, minification, and documentation. But now it has taken on a life of its own, and has grown beyond my wildest expectations.

It could not have been built without these great [projects](https://tinybase.org/guides/how-tinybase-is-built/credits/#giants) and [friends](https://tinybase.org/guides/how-tinybase-is-built/credits/#and-friends), and I hope you enjoy using it as much as I do building it!
