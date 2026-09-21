---
title: "InfluxDB - Time series data platform"
notion_id: 37bbf8b4-9f64-4d10-a15b-f2a9ebd509d0
notion_url: https://app.notion.com/p/InfluxDB-Time-series-data-platform-37bbf8b49f644d10a15bf2a9ebd509d0
last_edited: 2026-09-21T17:27:00.000Z
source_url: https://www.influxdata.com/
tags: ["English", "Español", "Others", "Databases", "Untried", "Tool", "Framework/Library", "Service"]
---
[https://www.influxdata.com/](https://www.influxdata.com/)

The database for real-time systems powering physical AI

## Developers choose InfluxDB

Built for systems that can’t afford to lag. Capture, analyze, and act on high-resolution operational data in real-time.

Try InfluxDB

1B+

Downloads of InfluxDB

1M+

Live open source instances

5B+

InfluxData Telegraf downloads

2,800+

Open source contributors

#1

Time series database Source: DB Engines

## Most databases weren’t built for time series

Telemetry, edge devices, and physical AI generate data that most databases can’t handle. InfluxDB handles continuous, high-resolution data with precision and speed.

## Engineered for performance at scale

Handle continuous, high-resolution data without breaking cost, latency, or reliability

See ways to get started

### Real-time performance without cost runaway

Ingest at millions of data points per second while keeping storage and compute predictable

### Edge-to-cloud continuity

Capture data where it’s generated and analyze it anywhere without redesigning your pipeline

### Purpose-built architecture for continuous data

Optimized for fast ingest on high-velocity, high-resolution data with efficient compression and storage

### Integrates with your stack

Deploy fast, scale seamlessly, and easily integrate without requiring specialized infrastructure

## Physical AI runs on time series

### Close the loop between telemetry and action.

AI systems need a stream of high-precision telemetry to improve. InfluxDB captures granular sensor data enabling AI to detect, respond and predict in real-time.

Learn More

### Aerospace

Aircraft and spacecraft require pinpoint precision for mission critical navigation. Capture high-rate sensor streams to power real-time diagnostics and adaptive control.

### Manufacturing

Production lines generate constant signals. Feed real-time telemetry into AI models to predict equipment failures, optimize throughput, and maintain product quality.

### Energy & Utilities

Manage grid volatility across distributed assets. Real-time sensor streams feed predictive models to prevent costly edge failures.

## Deploy anywhere

On-prem, edge, or cloud. A single engine for every environment.

## Open and extensible

Integrate time series across AI/ML workloads & visualize instantly

Explore integrations

### Code in the languages you love

### Build and ship faster with client libraries

Read more in docs

Python Javascript Go C# Java

ReadWrite

```plain text
from influxdb_client_3 import InfluxDBClient3 import pandas import os database = os.getenv('INFLUX_DATABASE') token = os.getenv('INFLUX_TOKEN') host="https://us-east-1-1.aws.cloud2.influxdata.com" def querySQL(): client = InfluxDBClient3(host, database=database, token=token) table = client.query( '''SELECT room, DATE_BIN(INTERVAL '1 day', time) AS _time, AVG(temp) AS temp, AVG(hum) AS hum, AVG(co) AS co FROM home WHERE time >= now() - INTERVAL '90 days' GROUP BY room, _time ORDER BY _time''' ) print(table.to_pandas().to_markdown()) client.close() querySQL()
```

```plain text
from influxdb_client_3 import InfluxDBClient3 import os database = os.getenv('INFLUX_DATABASE') token = os.getenv('INFLUX_TOKEN') host="https://us-east-1-1.aws.cloud2.influxdata.com" def write_line_protocol(): client = InfluxDBClient3(host, database=database, token=token) record = "home,room=Living\\ Room temp=22.2,hum=36.4,co=17i" print("Writing record:", record ) client.write(record) client.close() write_line_protocol()
```

```plain text
import {InfluxDBClient} from '@influxdata/influxdb3-client' import {tableFromArrays} from 'apache-arrow'; const database = process.env.INFLUX_DATABASE; const token = process.env.INFLUX_TOKEN; const host = "https://us-east-1-1.aws.cloud2.influxdata.com"; async function main() { const client = new InfluxDBClient({host, token}) const query = ` SELECT room, DATE_BIN(INTERVAL '1 day', time) AS _time, AVG(temp) AS temp, AVG(hum) AS hum, AVG(co) AS co FROM home WHERE time >= now() - INTERVAL '90 days' GROUP BY room, _time ORDER BY _time ` const result = await client.query(query, database) const data = {room: [], day: [], temp: []} for await (const row of result) { data.day.push(new Date(row._time).toISOString()) data.room.push(row.room) data.temp.push(row.temp) } console.table([...tableFromArrays(data)]) client.close() } main()
```

```plain text
import {InfluxDBClient} from '@influxdata/influxdb3-client' const database = process.env.INFLUX_DATABASE; const token = process.env.INFLUX_TOKEN; const host = "https://us-east-1-1.aws.cloud2.influxdata.com"; async function main() { const client = new InfluxDBClient({host, token}) const record = "home,room=Living\\ Room temp=22.2,hum=36.4,co=17i" await client.write(record, database) client.close() } main()
```

```plain text
package influxdbv3 import ( "context" "fmt" "io" "os" "text/tabwriter" "github.com/apache/arrow/go/v12/arrow" "github.com/InfluxCommunity/influxdb3-go/influx" ) func QuerySQL() error { url := "https://us-east-1-1.aws.cloud2.influxdata.com" token := os.Getenv("INFLUX_TOKEN") database := os.Getenv("INFLUX_DATABASE") client, err := influx.New(influx.Configs{ HostURL: url, AuthToken: token, }) defer func (client *influx.Client) { err := client.Close() if err != nil { panic(err) } }(client) query := ` SELECT room, DATE_BIN(INTERVAL '1 day', time) AS _time, AVG(temp) AS temp, AVG(hum) AS hum, AVG(co) AS co FROM home WHERE time >= now() - INTERVAL '90 days' GROUP BY room, _time ORDER BY _time ` iterator, err := client.Query(context.Background(), database, query) if err != nil { panic(err) } w := tabwriter.NewWriter(io.Discard, 4, 4, 1, ' ', 0) w.Init(os.Stdout, 0, 8, 0, '\t', 0) fmt.Fprintln(w, "day\troom\ttemp") for iterator.Next() { row := iterator.Value() day := (row["_time"].(arrow.Timestamp)).ToTime(arrow.TimeUnit(arrow.Nanosecond)) fmt.Fprintf(w, "%s\t%s\t%.2f\n", day, row["room"], row["temp"]) } w.Flush() return nil }
```

```plain text
package influxdbv3 import ( "context" "os" "fmt" "github.com/InfluxCommunity/influxdb3-go/influx" ) func WriteLineProtocol() error { url := "https://us-east-1-1.aws.cloud2.influxdata.com" token := os.Getenv("INFLUX_TOKEN") database := os.Getenv("INFLUX_DATABASE") client, err := influx.New(influx.Configs{ HostURL: url, AuthToken: token, }) defer func (client *influx.Client) { err := client.Close() if err != nil { panic(err) } }(client) record := "home,room=Living\\ Room temp=22.2,hum=36.4,co=17i" fmt.Println("Writing record: ", record) err = client.Write(context.Background(), database, []byte(record)) if err != nil { panic(err) } return nil }
```

```plain text
using System; using System.Threading.Tasks; using InfluxDB3.Client; using InfluxDB3.Client.Query; namespace InfluxDBv3; public class Query { static async Task QuerySQL() { const string hostUrl = "https://us-east-1-1.aws.cloud2.influxdata.com"; string? database = System.Environment.GetEnvironmentVariable("INFLUX_DATABASE"); string? authToken = System.Environment.GetEnvironmentVariable("INFLUX_TOKEN"); using var client = new InfluxDBClient(hostUrl, authToken: authToken, database: database); const string sql = @" SELECT room, DATE_BIN(INTERVAL '1 day', time) AS _time, AVG(temp) AS temp, AVG(hum) AS hum, AVG(co) AS co FROM home WHERE time >= now() - INTERVAL '90 days' GROUP BY room, _time ORDER BY _time "; Console.WriteLine("{0,-30}{1,-15}{2,-15}", "day", "room", "temp"); await foreach (var row in client.Query(query: sql)) { Console.WriteLine("{0,-30}{1,-15}{2,-15}", row[1], row[0], row[2]); } Console.WriteLine(); } }
```

```plain text
using System; using System.Threading.Tasks; using InfluxDB3.Client; using InfluxDB3.Client.Query; namespace InfluxDBv3; public class Write { public static async Task WriteLineProtocol() { const string hostUrl = "https://us-east-1-1.aws.cloud2.influxdata.com"; string? database = System.Environment.GetEnvironmentVariable("INFLUX_DATABASE"); string? authToken = System.Environment.GetEnvironmentVariable("INFLUX_TOKEN"); using var client = new InfluxDBClient(hostUrl, authToken: authToken, database: database); const string record = "home,room=Living\\ Room temp=22.2,hum=36.4,co=17i"; Console.WriteLine("Write record: {0,-30}", record); await client.WriteRecordAsync(record: record); } }
```

```plain text
package com.influxdb.v3; import java.time.Instant; import java.util.stream.Stream; import com.influxdb.v3.client.InfluxDBClient; import com.influxdb.v3.client.query.QueryOptions; import com.influxdb.v3.client.Point; public class IOxExample { public static void main(String[] args) throws Exception { String host = "https://us-east-1-1.aws.cloud2.influxdata.com"; char[] token = "my-token".toCharArray(); String database = "database"; try (InfluxDBClient client = InfluxDBClient.getInstance(host, token, database)) { // // Query by SQL // System.out.printf("--------------------------------------------------------%n"); System.out.printf("| %-8s | %-8s | %-30s |%n", "location", "value", "time"); System.out.printf("--------------------------------------------------------%n"); String sql = "select time,location,value from temperature order by time desc limit 10"; try (Stream stream = client.query(sql)) { stream.forEach(row -> System.out.printf("| %-8s | %-8s | %-30s |%n", row[1], row[2], row[0])); } System.out.printf("--------------------------------------------------------%n%n"); } } }
```

```plain text
package com.influxdb.v3; import java.time.Instant; import java.util.stream.Stream; import com.influxdb.v3.client.InfluxDBClient; import com.influxdb.v3.client.query.QueryOptions; import com.influxdb.v3.client.Point; public class IOxExample { public static void main(String[] args) throws Exception { String host = "https://us-east-1-1.aws.cloud2.influxdata.com"; char[] token = "my-token".toCharArray(); String database = "database"; try (InfluxDBClient client = InfluxDBClient.getInstance(host, token, database)) { // // Write by Point // Point point = Point.measurement("temperature") .setTag("location", "west") .setField("value", 55.15) .setTimestamp(Instant.now().minusSeconds(-10)); client.writePoint(point); // // Write by LineProtocol // String record = "temperature,location=north value=60.0"; client.writeRecord(record); } } }
```

### 400+ Telegraf plugins

Seamless integration with your favorite tools with Telegraf, our popular open source connector with 5B+ downloads.

### Client libraries

InfluxDB client libraries make it easy to integrate time series data into your applications using your favorite programming languages.

### Community & ecosystem

InfluxDB includes a massive community of cloud and open source developers to help you work the way you want.
