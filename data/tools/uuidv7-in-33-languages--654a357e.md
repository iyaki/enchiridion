---
title: "UUIDv7 in 33 languages"
notion_id: 654a357e-e4dc-4fdf-a37c-61cb8f4a7417
notion_url: https://app.notion.com/p/UUIDv7-in-33-languages-654a357ee4dc4fdfa37c61cb8f4a7417
last_edited: 2024-07-15T23:43:00.000Z
source_url: https://antonz.org/uuidv7/
tags: ["English", "Programming", "System Design / Software Architecture", "Website", "Framework/Library"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

UUIDv7 is a 128-bit unique identifier like it's older siblings, such as the widely used UUIDv4. But unlike v4, UUIDv7 is time-sortable with 1 ms precision. By combining the timestamp and the random parts, UUIDv7 becomes an excellent choice for record identifiers in databases, including distributed ones.

Let's briefly explore the UUIDv7 structure and move on to the zero-dependency implementations in 33 languages (as ranked by the Stack Overflow survey).

> These implementations may not be the fastest or most idiomatic, but they are concise and easy to understand. Many of the code examples are interactive (but the results are cached, so you won't see different UUIDs very often).

As you can imagine, I am by no means fluent in all these languages, so if you spot a bug — please submit a [pull request](https://github.com/nalgeon/uuidv7). PRs for other languages are also welcome!

[JavaScript](https://antonz.org/uuidv7/#javascript) • [Python](https://antonz.org/uuidv7/#python) • [SQL](https://antonz.org/uuidv7/#sql) • [Shell](https://antonz.org/uuidv7/#shell) • [Java](https://antonz.org/uuidv7/#java) • [C#](https://antonz.org/uuidv7/#csharp) • [C++](https://antonz.org/uuidv7/#cpp) • [C](https://antonz.org/uuidv7/#c) • [PHP](https://antonz.org/uuidv7/#php) • [PowerShell](https://antonz.org/uuidv7/#powershell) • [Go](https://antonz.org/uuidv7/#go) • [Rust](https://antonz.org/uuidv7/#rust) • [Kotlin](https://antonz.org/uuidv7/#kotlin) • [Ruby](https://antonz.org/uuidv7/#ruby) • [Lua](https://antonz.org/uuidv7/#lua) • [Dart](https://antonz.org/uuidv7/#dart) • [Swift](https://antonz.org/uuidv7/#swift) • [R](https://antonz.org/uuidv7/#r) • [Pascal](https://antonz.org/uuidv7/#pascal) • [Perl](https://antonz.org/uuidv7/#perl) • [Elixir](https://antonz.org/uuidv7/#elixir) • [Clojure](https://antonz.org/uuidv7/#clojure) • [Julia](https://antonz.org/uuidv7/#julia) • [Erlang](https://antonz.org/uuidv7/#erlang) • [Zig](https://antonz.org/uuidv7/#zig) • [Crystal](https://antonz.org/uuidv7/#crystal) • [Nim](https://antonz.org/uuidv7/#nim) • [Gleam](https://antonz.org/uuidv7/#gleam) • [Tcl](https://antonz.org/uuidv7/#tcl) • [V](https://antonz.org/uuidv7/#v) • [Emacs Lisp](https://antonz.org/uuidv7/#emacs-lisp) • [Vimscript](https://antonz.org/uuidv7/#vimscript) • [Nushell](https://antonz.org/uuidv7/#nushell)

## Structure

UUIDv7 looks like this when represented as a string:

```plain text
0190163d-8694-739b-aea5-966c26f8ad91
└─timestamp─┘ │└─┤ │└───rand_b─────┘
             ver │var
              rand_a

```

The 128-bit value consists of several parts:

- `timestamp` (48 bits) is a Unix timestamp in milliseconds.
- `ver` (4 bits) is a UUID version (`7`).
- `rand_a` (12 bits) is randomly generated.
- `var`(2 bits) is equal to `10`.
- `rand_b` (62 bits) is randomly generated.
- In string representation, each symbol encodes 4 bits as a hex number, so the `a` in the example is `1010`, where the first two bits are the fixed variant (`10`) and the next two are random. So the resulting hex number can be either `8` (`1000`), `9` (`1001`), `a` (`1010`) or `b` (`1011`).

See [RFC 9652](https://www.rfc-editor.org/rfc/rfc9562#name-uuid-version-7) for details.

## JavaScript

Initialize a random array with `crypto.getRandomValues()`, get the current timestamp with `Date.now()`, fill the array from the timestamp, set version and variant.

```plain text
function uuidv7() {
    // random bytes
    const value = new Uint8Array(16);
    crypto.getRandomValues(value);
    // current timestamp in ms
    const timestamp = BigInt(Date.now());
    // timestamp
    value[0] = Number((timestamp >> 40n) & 0xffn);
    value[1] = Number((timestamp >> 32n) & 0xffn);
    value[2] = Number((timestamp >> 24n) & 0xffn);
    value[3] = Number((timestamp >> 16n) & 0xffn);
    value[4] = Number((timestamp >> 8n) & 0xffn);
    value[5] = Number(timestamp & 0xffn);
    // version and variant
    value[6] = (value[6] & 0x0f) | 0x70;
    value[8] = (value[8] & 0x3f) | 0x80;
    return value;
}
const uuidVal = uuidv7();
const uuidStr = Array.from(uuidVal)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
console.log(uuidStr);

```

[Edit](https://antonz.org/uuidv7/#edit)

The TypeScript version is identical, the only difference is the function signature:

```plain text
function uuidv7(): Uint8Array {
    // ...
}

```

## Python

Initialize a random array with `os.urandom()`, get the current timestamp with `time.time()`, fill the array from the timestamp, set version and variant.

```plain text
import os
import time
def uuidv7():
    # random bytes
    value = bytearray(os.urandom(16))
    # current timestamp in ms
    timestamp = int(time.time() * 1000)
    # timestamp
    value[0] = (timestamp >> 40) & 0xFF
    value[1] = (timestamp >> 32) & 0xFF
    value[2] = (timestamp >> 24) & 0xFF
    value[3] = (timestamp >> 16) & 0xFF
    value[4] = (timestamp >> 8) & 0xFF
    value[5] = timestamp & 0xFF
    # version and variant
    value[6] = (value[6] & 0x0F) | 0x70
    value[8] = (value[8] & 0x3F) | 0x80
    return value
if __name__ == "__main__":
    uuid_val = uuidv7()
    print(''.join(f'{byte:02x}' for byte in uuid_val))

```

[Edit](https://antonz.org/uuidv7/#edit)

## SQL

Get the current timestamp parts with `strftime()` (SQLite) or `now()` (PostgreSQL), get the random parts with `random()`, concatenate everything into a UUID string.

SQLite (by [Fabio Lima](https://gist.github.com/fabiolimace/e3c3d354d1afe0b3175f65be2d962523)):

```plain text
select
  -- timestamp
  format('%08x', ((strftime('%s') * 1000) >> 16)) || '-' ||
  format('%04x', ((strftime('%s') * 1000)
    + ((strftime('%f') * 1000) % 1000)) & 0xffff) || '-' ||
  -- version / rand_a
  format('%04x', 0x7000 + abs(random()) % 0x0fff) || '-' ||
  -- variant / rand_b
  format('%04x', 0x8000 + abs(random()) % 0x3fff) || '-' ||
  -- rand_b
  format('%012x', abs(random()) >> 16) as value;

```

[Edit](https://antonz.org/uuidv7/#edit)

PostgreSQL:

```plain text
select
  -- timestamp
  lpad(to_hex(((extract(epoch from now()) * 1000)::bigint >> 16)), 8, '0') || '-' ||
  lpad(to_hex(((extract(epoch from now()) * 1000
    + (date_part('milliseconds', now())::bigint % 1000))::bigint & 0xffff)), 4, '0') || '-' ||
  -- version / rand_a
  lpad(to_hex((0x7000 + (random() * 0x0fff)::int)), 4, '0') || '-' ||
  -- variant / rand_b
  lpad(to_hex((0x8000 + (random() * 0x3fff)::int)), 4, '0') || '-' ||
  -- rand_b
  lpad(to_hex((floor(random() * (2^48))::bigint >> 16)), 12, '0') AS value;

```

[Edit](https://antonz.org/uuidv7/#edit)

SQL Server (by [Onur Keskin](https://github.com/keskinonur)):

```plain text
create procedure GenerateUUIDv7
as
begin
    -- Declare variables to hold different parts of the UUID
    declare @timestamp_hex    char(8)
    declare @milliseconds_hex char(4)
    declare @version_rand_a   char(4)
    declare @variant_rand_b   char(4)
    declare @rand_b           char(12)
    -- Calculate the current time in milliseconds since the Unix epoch
    declare @currentTime bigint = datediff_big(millisecond, '1970-01-01', getutcdate())
    -- Convert the timestamp to hexadecimal, divided into two parts
    set @timestamp_hex = right('00000000' + convert(varchar(8), convert(varbinary(4), @currentTime / 65536), 2), 8)
    set @milliseconds_hex = right('0000' + convert(varchar(4), convert(varbinary(2), @currentTime % 65536), 2), 4)
    -- Generate a random value for the version/rand_a part, combine it with the version identifier (0x7000), and convert to hexadecimal
    declare @rand_a int = cast(floor(rand(checksum(newid())) * 4096) as int) -- 0x0FFF = 4095
    set @version_rand_a = right('0000' + convert(varchar(4), convert(varbinary(2), 0x7000 + @rand_a), 2), 4)
    -- Generate a random value for the variant/rand_b part, combine it with the variant identifier (0x8000), and convert to hexadecimal
    declare @rand_b_part int = cast(floor(rand(checksum(newid())) * 16384) as int) -- 0x3FFF = 16383
    set @variant_rand_b = right('0000' + convert(varchar(4), convert(varbinary(2), 0x8000 + @rand_b_part), 2), 4)
    -- Generate a large random value for the rand_b part and convert to hexadecimal
    declare @rand_b_bigint bigint = floor(rand(checksum(newid())) * power(cast(2 as bigint), 48))
    set @rand_b = right('000000000000' + convert(varchar(12), convert(varbinary(6), @rand_b_bigint / 65536), 2), 12)
    -- Combine all parts to form the final UUID v7-like value
    declare @UUIDv7 char(36)
    set @UUIDv7 = @timestamp_hex + '-' + @milliseconds_hex + '-' + @version_rand_a + '-' + @variant_rand_b + '-' + @rand_b
    -- Return the generated UUID v7-like value
    select @UUIDv7 as UUIDv7
end

```

## Shell

Initialize a random byte string with `/dev/urandom`, get the current timestamp with `date`, fill the array from the timestamp and the byte string, set version and variant.

```plain text
#!/bin/sh
uuidv7() {
    # current timestamp in ms. POSIX date does not support %3N.
    timestamp=$(date +%s)000
    timestamp_hi=$(( timestamp >> 16 ))
    timestamp_lo=$(( timestamp & 0xFFFF ))
    # 16 random bits (12 will be used)
    rand_a=0x$(LC_ALL=C tr -dc '0-9a-f' < /dev/urandom|head -c4)
    # ver is 0x7
    ver_rand_a=$(( 0x7000 | ( 0xFFF & rand_a ) ))
    # 16 random bits (14 will be used)
    rand_b_hi=0x$(LC_ALL=C tr -dc '0-9a-f' < /dev/urandom|head -c4)
    # var is 0b10
    var_rand_hi=$(( 0x8000 | ( 0x3FFF & rand_b_hi ) ))
    # remaining 48 bits of rand b
    rand_b_lo=$(LC_ALL=C tr -dc '0-9a-f' < /dev/urandom|head -c12)
    printf "%08x-%04x-%04x-%4x-%s" "$timestamp_hi" "$timestamp_lo" "$ver_rand_a" "$var_rand_hi" "$rand_b_lo"
}
echo $(uuidv7)

```

[Edit](https://antonz.org/uuidv7/#edit)

by [Brian Ewins](https://github.com/bewinsnw)

## Java

Initialize a random array with `SecureRandom.nextBytes()`, get the current timestamp with `System.currentTimeMillis()`, fill the array from the timestamp, set version and variant.

```plain text
import java.nio.ByteBuffer;
import java.security.SecureRandom;
import java.util.UUID;
public class UUIDv7 {
    private static final SecureRandom random = new SecureRandom();
    public static UUID randomUUID() {
        byte[] value = randomBytes();
        ByteBuffer buf = ByteBuffer.wrap(value);
        long high = buf.getLong();
        long low = buf.getLong();
        return new UUID(high, low);
    }
    public static byte[] randomBytes() {
        // random bytes
        byte[] value = new byte[16];
        random.nextBytes(value);
        // current timestamp in ms
        ByteBuffer timestamp = ByteBuffer.allocate(Long.BYTES);
        timestamp.putLong(System.currentTimeMillis());
        // timestamp
        System.arraycopy(timestamp.array(), 2, value, 0, 6);
        // version and variant
        value[6] = (byte) ((value[6] & 0x0F) | 0x70);
        value[8] = (byte) ((value[8] & 0x3F) | 0x80);
        return value;
    }
    public static void main(String[] args) {
        var uuid = UUIDv7.randomUUID();
        System.out.println(uuid);
    }
}

```

[Edit](https://antonz.org/uuidv7/#edit)

by [David Ankin](https://github.com/alexanderankin)

## C#

Initialize a random array with `RandomNumberGenerator.GetBytes()`, get the current timestamp with `DateTimeOffset.UtcNow`, fill the array from the timestamp, set version and variant.

```plain text
using System;
using System.Security.Cryptography;
public class UUIDv7 {
    private static readonly RandomNumberGenerator random =
        RandomNumberGenerator.Create();
    public static byte[] Generate() {
        // random bytes
        byte[] value = new byte[16];
        random.GetBytes(value);
        // current timestamp in ms
        long timestamp = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds();
        // timestamp
        value[0] = (byte)((timestamp >> 40) & 0xFF);
        value[1] = (byte)((timestamp >> 32) & 0xFF);
        value[2] = (byte)((timestamp >> 24) & 0xFF);
        value[3] = (byte)((timestamp >> 16) & 0xFF);
        value[4] = (byte)((timestamp >> 8) & 0xFF);
        value[5] = (byte)(timestamp & 0xFF);
        // version and variant
        value[6] = (byte)((value[6] & 0x0F) | 0x70);
        value[8] = (byte)((value[8] & 0x3F) | 0x80);
        return value;
    }
    public static void Main(string[] args) {
        byte[] uuidVal = Generate();
        foreach (byte b in uuidVal) {
            Console.Write("{0:x2}", b);
        }
        Console.WriteLine();
    }
}

```

[Edit](https://antonz.org/uuidv7/#edit)

## C++

Initialize a random array with `random_device`, get the current timestamp with `system_clock.time_since_epoch()`, fill the array from the timestamp, set version and variant.

```plain text
#include <array>#include <chrono>#include <cstdint>#include <cstdio>#include <random>std::array<uint8_t, 16> uuidv7() {
    // random bytes
    std::random_device rd;
    std::array<uint8_t, 16> random_bytes;
    std::generate(random_bytes.begin(), random_bytes.end(), std::ref(rd));
    std::array<uint8_t, 16> value;
    std::copy(random_bytes.begin(), random_bytes.end(), value.begin());
    // current timestamp in ms
    auto now = std::chrono::system_clock::now();
    auto millis = std::chrono::duration_cast<std::chrono::milliseconds>(
        now.time_since_epoch()
    ).count();
    // timestamp
    value[0] = (millis >> 40) & 0xFF;
    value[1] = (millis >> 32) & 0xFF;
    value[2] = (millis >> 24) & 0xFF;
    value[3] = (millis >> 16) & 0xFF;
    value[4] = (millis >> 8) & 0xFF;
    value[5] = millis & 0xFF;
    // version and variant
    value[6] = (value[6] & 0x0F) | 0x70;
    value[8] = (value[8] & 0x3F) | 0x80;
    return value;
}
int main() {
    auto uuid_val = uuidv7();
    for (const auto& byte : uuid_val) {
        printf("%02x", byte);
    }
    printf("\n");
    return 0;
}

```

[Edit](https://antonz.org/uuidv7/#edit)

## C

Initialize a random array with `getentropy()`, get the current timestamp with `timespec_get()`, fill the array from the timestamp, set version and variant.

```plain text
#include <stdio.h>#include <stdlib.h>#include <stdint.h>#include <time.h>#include <unistd.h>int uuidv7(uint8_t* value) {
    // random bytes
    int err = getentropy(value, 16);
    if (err != EXIT_SUCCESS) {
        return EXIT_FAILURE;
    }
    // current timestamp in ms
    struct timespec ts;
    int ok = timespec_get(&ts, TIME_UTC);
    if (ok == 0) {
        return EXIT_FAILURE;
    }
    uint64_t timestamp = (uint64_t)ts.tv_sec * 1000 + ts.tv_nsec / 1000000;
    // timestamp
    value[0] = (timestamp >> 40) & 0xFF;
    value[1] = (timestamp >> 32) & 0xFF;
    value[2] = (timestamp >> 24) & 0xFF;
    value[3] = (timestamp >> 16) & 0xFF;
    value[4] = (timestamp >> 8) & 0xFF;
    value[5] = timestamp & 0xFF;
    // version and variant
    value[6] = (value[6] & 0x0F) | 0x70;
    value[8] = (value[8] & 0x3F) | 0x80;
    return EXIT_SUCCESS;
}
int main() {
    uint8_t uuid_val[16];
    uuidv7(uuid_val);
    for (size_t i = 0; i < 16; i++) {
        printf("%02x", uuid_val[i]);
    }
    printf("\n");
}

```

[Edit](https://antonz.org/uuidv7/#edit)

## PHP

Initialize a random string with `random_bytes()`, get the current timestamp with `microtime()`, fill the characters from the timestamp, set version and variant.

```plain text
<?php
function uuidv7() {
    // random bytes
    $value = random_bytes(16);
    // current timestamp in ms
    $timestamp = intval(microtime(true) * 1000);
    // timestamp
    $value[0] = chr(($timestamp >> 40) & 0xFF);
    $value[1] = chr(($timestamp >> 32) & 0xFF);
    $value[2] = chr(($timestamp >> 24) & 0xFF);
    $value[3] = chr(($timestamp >> 16) & 0xFF);
    $value[4] = chr(($timestamp >> 8) & 0xFF);
    $value[5] = chr($timestamp & 0xFF);
    // version and variant
    $value[6] = chr((ord($value[6]) & 0x0F) | 0x70);
    $value[8] = chr((ord($value[8]) & 0x3F) | 0x80);
    return $value;
}
$uuid_val = uuidv7();
echo bin2hex($uuid_val);

```

[Edit](https://antonz.org/uuidv7/#edit)

## PowerShell

Initialize a random array with `Random.GetBytes()`, get the current timestamp with `DateTimeOffset.UtcNow`, fill the array from the timestamp, set version and variant.

```plain text
function New-Uuidv7
{
  # random bytes
  $value = [byte[]]::new(16)
  [System.Random]::new().NextBytes($value)
  # current timestamp
  $timestamp = [DateTimeOffset]::UtcNow.ToUnixTimeMilliseconds()
  [System.BitConverter]::GetBytes($timestamp)[5..0].Copyto($value, 0)
  # version and variant
  $value[6] = ($value[6] -band 0x0F) -bor 0x70
  $value[8] = ($value[8] -band 0x0F) -bor 0x80
  $value
}
(New-Uuidv7 | ForEach-Object ToString x2) -join ''

```

by [lost](https://github.com/lost22git)

## Go

Initialize a random array with `rand.Read()`, get the current timestamp with `time.Now()`, fill the array from the timestamp, set version and variant.

```plain text
package main
import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)
func uuidv7() ([16]byte, error) {
	// random bytes
	var value [16]byte
	_, err := rand.Read(value[:])
	if err != nil {
		return value, err
	}
	// current timestamp in ms
	timestamp := big.NewInt(time.Now().UnixMilli())
	// timestamp
	timestamp.FillBytes(value[0:6])
	// version and variant
	value[6] = (value[6] & 0x0F) | 0x70
	value[8] = (value[8] & 0x3F) | 0x80
	return value, nil
}
func main() {
	uuidVal, _ := uuidv7()
	fmt.Printf("%x\n", uuidVal)
}

```

[Edit](https://antonz.org/uuidv7/#edit)

## Rust

Initialize a random array with `getrandom::getrandom()`, get the current timestamp with `SystemTime::now()`, fill the array from the timestamp, set version and variant.

```plain text
use std::error::Error;
use std::time::{SystemTime, UNIX_EPOCH};
fn uuidv7() -> Result<[u8; 16], Box<dyn Error>> {
    // random bytes
    let mut value = [0u8; 16];
    getrandom::getrandom(&mut value)?;
    // current timestamp in ms
    let timestamp = match SystemTime::now().duration_since(UNIX_EPOCH) {
        Ok(duration) => duration.as_millis() as u64,
        Err(_) => return Err(Box::from("Failed to get system time")),
    };
    // timestamp
    value[0] = (timestamp >> 40) as u8;
    value[1] = (timestamp >> 32) as u8;
    value[2] = (timestamp >> 24) as u8;
    value[3] = (timestamp >> 16) as u8;
    value[4] = (timestamp >> 8) as u8;
    value[5] = timestamp as u8;
    // version and variant
    value[6] = (value[6] & 0x0F) | 0x70;
    value[8] = (value[8] & 0x3F) | 0x80;
    Ok(value)
}
fn main() {
    match uuidv7() {
        Ok(uuid_val) => {
            for byte in &uuid_val {
                print!("{:02x}", byte);
            }
            println!();
        }
        Err(e) => eprintln!("Error: {}", e),
    }
}

```

by [Rodolphe Bréard](https://github.com/breard-r)

## Kotlin

Initialize a random array with `SecureRandom.nextBytes()`, get the current timestamp with `Instant.now()`, fill the array from the timestamp, set version and variant.

```plain text
import java.security.SecureRandom
import java.time.Instant
object UUIDv7 {
    private val random = SecureRandom()
    fun generate(): ByteArray {
        // random bytes
        val value = ByteArray(16)
        random.nextBytes(value)
        // current timestamp in ms
        val timestamp = Instant.now().toEpochMilli()
        // timestamp
        value[0] = ((timestamp shr 40) and 0xFF).toByte()
        value[1] = ((timestamp shr 32) and 0xFF).toByte()
        value[2] = ((timestamp shr 24) and 0xFF).toByte()
        value[3] = ((timestamp shr 16) and 0xFF).toByte()
        value[4] = ((timestamp shr 8) and 0xFF).toByte()
        value[5] = (timestamp and 0xFF).toByte()
        // version and variant
        value[6] = (value[6].toInt() and 0x0F or 0x70).toByte()
        value[8] = (value[8].toInt() and 0x3F or 0x80).toByte()
        return value
    }
    @JvmStatic
    fun main(args: Array<String>) {
        val uuidVal = generate()
        uuidVal.forEach { b -> print("%02x".format(b)) }
        println()
    }
}

```

## Ruby

Initialize a random array with `SecureRandom.random_bytes()`, get the current timestamp with `Time.now`, fill the array from the timestamp, set version and variant.

```plain text
require 'securerandom'
require 'time'
def uuidv7
  # random bytes
  value = SecureRandom.random_bytes(16).bytes
  # current timestamp in ms
  timestamp = (Time.now.to_f * 1000).to_i
  # timestamp
  value[0] = (timestamp >> 40) & 0xFF
  value[1] = (timestamp >> 32) & 0xFF
  value[2] = (timestamp >> 24) & 0xFF
  value[3] = (timestamp >> 16) & 0xFF
  value[4] = (timestamp >> 8) & 0xFF
  value[5] = timestamp & 0xFF
  # version and variant
  value[6] = (value[6] & 0x0F) | 0x70
  value[8] = (value[8] & 0x3F) | 0x80
  value
end
if __FILE__ == $0
  uuid_val = uuidv7
  puts uuid_val.pack('C*').unpack1('H*')
end

```

[Edit](https://antonz.org/uuidv7/#edit)

## Lua

Initialize a random table with `math.random()`, get the current timestamp with `os.time()`, fill the list from the timestamp, set version and variant.

```plain text
local function uuidv7()
    -- random bytes
    local value = {}
    for i = 1, 16 do
        value[i] = math.random(0, 255)
    end
    -- current timestamp in ms
    local timestamp = os.time() * 1000
    -- timestamp
    value[1] = (timestamp >> 40) & 0xFF
    value[2] = (timestamp >> 32) & 0xFF
    value[3] = (timestamp >> 24) & 0xFF
    value[4] = (timestamp >> 16) & 0xFF
    value[5] = (timestamp >> 8) & 0xFF
    value[6] = timestamp & 0xFF
    -- version and variant
    value[7] = (value[7] & 0x0F) | 0x70
    value[9] = (value[9] & 0x3F) | 0x80
    return value
end
local uuid_val = uuidv7()
for i = 1, #uuid_val do
    io.write(string.format('%02x', uuid_val[i]))
end
print()

```

[Edit](https://antonz.org/uuidv7/#edit)

## Dart

Initialize a random list with `Random.nextInt()`, get the current timestamp with `DateTime.now()`, fill the list from the timestamp, set version and variant.

```plain text
import 'dart:math';
import 'dart:typed_data';
Uint8List uuidv7() {
  // random bytes
  final rng = Random.secure();
  final value = Uint8List(16);
  for (int i = 0; i < 16; i++) {
    value[i] = rng.nextInt(256);
  }
  // current timestamp in ms
  final timestamp = DateTime.now().millisecondsSinceEpoch;
  // timestamp
  value[0] = (timestamp ~/ pow(2, 40)) & 0xFF;
  value[1] = (timestamp ~/ pow(2, 32)) & 0xFF;
  value[2] = (timestamp ~/ pow(2, 24)) & 0xFF;
  value[3] = (timestamp ~/ pow(2, 16)) & 0xFF;
  value[4] = (timestamp ~/ pow(2, 8)) & 0xFF;
  value[5] = timestamp & 0xFF;
  // version and variant
  value[6] = (value[6] & 0x0F) | 0x70;
  value[8] = (value[8] & 0x3F) | 0x80;
  return value;
}
void main() {
  final uuidVal = uuidv7();
  print(uuidVal.map((byte) => byte.toRadixString(16).padLeft(2, '0')).join());
}

```

## Swift

Extend the existing `UUID` type with a `v7` function. Initialize a random tuple with `UInt8.random()`, get the current timestamp with `Date().timeIntervalSince1970`, fill the tuple from the timestamp, set version and variant.

```plain text
import Foundation
extension UUID {
    static func v7() -> Self {
        // random bytes
        var value = (
            UInt8(0),
            UInt8(0),
            UInt8(0),
            UInt8(0),
            UInt8(0),
            UInt8(0),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255),
            UInt8.random(in: 0...255)
        )
        // current timestamp in ms
        let timestamp: Int = .init(Date().timeIntervalSince1970 * 1000)
        // timestamp
        value.0 = .init((timestamp >> 40) & 0xFF)
        value.1 = .init((timestamp >> 32) & 0xFF)
        value.2 = .init((timestamp >> 24) & 0xFF)
        value.3 = .init((timestamp >> 16) & 0xFF)
        value.4 = .init((timestamp >> 8) & 0xFF)
        value.5 = .init(timestamp & 0xFF)
        // version and variant
        value.6 = (value.6 & 0x0F) | 0x70
        value.8 = (value.8 & 0x3F) | 0x80
        return UUID(uuid: value)
    }
}
let uuidVal: UUID = .v7()
print(uuidVal)

```

by [Prathamesh Kowarkar](https://github.com/p16r), [Sam Dean](https://github.com/deanWombourne)

## R

Initialize a random vector with `sample()`, get the current timestamp with `Sys.time()`, fill the vector from the timestamp, set version and variant.

```plain text
uuidv7 <- function() {
  # Initialise vector with current timestamp & random numbers
  value = numeric(16)
  value[1:6] <- as.numeric(Sys.time()) * 1000
  value[7:16] <- sample(0:255, 10L, replace = TRUE)
  # timestamp
  value[1:6] <- value[1:6] %/% 2^c(40, 32, 24, 16, 8, 1) %% 256L
  # version and variant
  value[7] <- bitwOr(bitwAnd(value[7], 0x0F), 0x70)
  value[9] <- bitwOr(bitwAnd(value[9], 0x3F), 0x80)
  as.raw(value)
}
uuid_val <- uuidv7()
cat(paste(sprintf('%02x', as.integer(uuid_val)), collapse = ''))

```

[Edit](https://antonz.org/uuidv7/#edit)

by [Colin Gillespie](https://github.com/csgillespie)

## Pascal

Initialize a random array with `Random()`, get the current timestamp with `DateTimeToUnix()`, fill the array from the timestamp, set version and variant.

```plain text
// Use as a regular unit from Delphi, or run as a console app from FreePascal
unit uuidv7;

interface

uses
  SysUtils, DateUtils;

function GenerateUUIDv7: TGUID;

implementation

function GenerateUUIDv7: TGUID;
var
  timestamp: Int64;
  randomBytes: array[0..9] of Byte;
  uuid: TGUID;
  i: Integer;
begin
  FillChar(uuid, SizeOf(uuid), 0);
  {$IFDEF FPC}
  timestamp := DateTimeToUnix(Now) * 1000; // seconds accuracy
  {$ELSE}
  timestamp := DateTimeToMilliseconds(Now) - Int64(UnixDateDelta + DateDelta) * MSecsPerDay; // millisecond accuracy
  {$ENDIF}

  // Generate 10 random bytes
  for i := 0 to 9 do
    randomBytes[i] := Random($FF);

  // Populate the TGUID fields
  uuid.D1 := (timestamp shr 16) and $FFFFFFFF;       // Top 32 bits of the 48-bit timestamp
  uuid.D2 := ((timestamp shr 4) and $0FFF) or $7000; // Next 12 bits of the timestamp and version 7
  uuid.D3 := ((timestamp and $0000000F) shl 12) or   // the last 4 bits of timestamp
              (randomBytes[0] and $F0);              // the top 4 bits of randomBytes[0]
  uuid.D4[0] := (randomBytes[0] and $0F) or $80;     // Set the variant to 10xx
  Move(randomBytes[1], uuid.D4[1], 7);               // Remaining 7 bytes

  Result := uuid;
end;

// Optionally remove this to make a regular unit for FPC too
{$IFDEF FPC}
var i: Integer;
begin
  Randomize;
  for i := 0 to 30 do
    writeln(GUIDToString(GenerateUUIDv7).ToLower);
  readln;
{$ELSE}
initialization
  Randomize;
{$ENDIF}
end.

```

by [Jim McKeeth](https://github.com/jimmckeeth)

## Perl

Initialize a random array with `rand()`, get the current timestamp with `Time::HiRes::time()`, fill the array from the timestamp, set version and variant.

```plain text
#!/usr/bin/env perl
use v5.16;
use Time::HiRes;
sub uuidv7 {
	my $type = shift() || "";
	# 16 random bytes (4 * 4)
	my $uuid = "";
	for (my $i = 0; $i < 4; $i++) {
		$uuid .= pack('I', int(rand(2 ** 32)));
	}
	# current timestamp in ms
	my $timestamp = int(Time::HiRes::time() * 1000);
	# timestamp
	substr($uuid, 0, 1, chr(($timestamp >> 40) & 0xFF));
	substr($uuid, 1, 1, chr(($timestamp >> 32) & 0xFF));
	substr($uuid, 2, 1, chr(($timestamp >> 24) & 0xFF));
	substr($uuid, 3, 1, chr(($timestamp >> 16) & 0xFF));
	substr($uuid, 4, 1, chr(($timestamp >>  8) & 0xFF));
	substr($uuid, 5, 1, chr($timestamp         & 0xFF));
	# version and variant
	substr($uuid, 6, 1, chr((ord(substr($uuid, 6, 1)) & 0x0F) | 0x70));
	substr($uuid, 8, 1, chr((ord(substr($uuid, 8, 1)) & 0x3F) | 0x80));
	return $uuid;
}
my $uuid_val = uuidv7('hex');
printf(unpack("H*", $uuid_val));

```

by [Scott Baker](https://github.com/scottchiefbaker)

## Elixir

Initialize a random list with `crypto.strong_rand_bytes()`, get the current timestamp with `os.system_time()`, fill the list from the timestamp, set version and variant.

```plain text
use Bitwise
defmodule UUIDv7 do
  def generate do
    # random bytes
    value = :crypto.strong_rand_bytes(16) |> :binary.bin_to_list()
    # current timestamp in ms
    timestamp = :os.system_time(:millisecond)
    # timestamp
    value = List.replace_at(value, 0, (timestamp >>> 40) &&& 0xFF)
    value = List.replace_at(value, 1, (timestamp >>> 32) &&& 0xFF)
    value = List.replace_at(value, 2, (timestamp >>> 24) &&& 0xFF)
    value = List.replace_at(value, 3, (timestamp >>> 16) &&& 0xFF)
    value = List.replace_at(value, 4, (timestamp >>> 8) &&& 0xFF)
    value = List.replace_at(value, 5, timestamp &&& 0xFF)
    # timestamp
    value = List.replace_at(value, 6, (Enum.at(value, 6) &&& 0x0F) ||| 0x70)
    value = List.replace_at(value, 8, (Enum.at(value, 8) &&& 0x3F) ||| 0x80)
    value
  end
end
uuid_val = UUIDv7.generate()
Enum.map(uuid_val, &Integer.to_string(&1, 16))
|> Enum.map(&String.pad_leading(&1, 2, "0"))
|> Enum.join()
|> IO.puts()

```

## Clojure

Initialize a random array with `SecureRandom.nextBytes()`, get the current timestamp with `System.currentTimeMillis()`, fill the array from the timestamp, set version and variant.

```plain text
(ns uuidv7
  (:require [clojure.string :as str])
  (:import (java.security SecureRandom)))
(defn gen-uuid-v7 []
  (let [rand-array (byte-array 10)]
    (.nextBytes (SecureRandom.) rand-array)
    (concat
      ;; timestamp
      (map byte (.toByteArray (biginteger (System/currentTimeMillis))))
      ;; version
      [(bit-or (bit-and (first rand-array) 0x0F) 0x70)]
      [(nth rand-array 1)]
      ;; variant
      [(bit-or (bit-and (nth rand-array 2) 0x3F) 0x80)]
      (drop 3 rand-array))))
(defn uuid-to-string [uuid-bytes]
  (apply str (map #(format "%02x" %) uuid-bytes)))
(def uuid-bytes (gen-uuid-v7))
(println (uuid-to-string uuid-bytes))

```

by [Saidone](https://github.com/saidone75)

## Julia

Initialize a random array with `rand()`, get the current timestamp with `time()`, fill the array from the timestamp, set version and variant.

```plain text
function uuidv7()
  # random bytes
  value = rand(UInt8, 16)
  # current timestamp
  timestamp = trunc(UInt64, time() * 1000)
  digits!(UInt8[0, 0, 0, 0, 0, 0, 0, 0], hton(timestamp), base=256) |> x -> copyto!(value, 1, x, 3, 6)
  # version and variant
  value[7] = value[7] & 0x0F | 0x70
  value[9] = value[9] & 0x3F | 0x80
  value
end
uuidv7() |> bytes2hex |> println

```

by [lost](https://github.com/lost22git)

## Erlang

Generate random bytes with `crypto:strong_rand_bytes()`, get the current timestamp with `os:system_time()`, set version and variant, combine everything together.

```plain text
-module(uuidv7).
-export([generate/0, main/1]).
-spec generate() -> binary().
generate() ->
    <<RandA:12, RandB:62, _:6>> = crypto:strong_rand_bytes(10),
    UnixTsMs = os:system_time(millisecond),
    Ver = 2#0111,
    Var = 2#10,
    <<UnixTsMs:48, Ver:4, RandA:12, Var:2, RandB:62>>.
main(_) ->
    UUIDv7 = generate(),
    %% note: if you use an erlang release newer than OTP23,
    %%       there is binary:encode_hex/1,2
    io:format("~s~n", [[io_lib:format("~2.16.0b",[X]) || <<X:8>> <= UUIDv7]]).

```

by [Stefan Grundmann](https://github.com/sg2342)

## Zig

Initialize a random array with `std.crypto.random.bytes()`, get the current timestamp with `std.time.milliTimestamp()`, fill the array from the timestamp, set version and variant.

```plain text
const std = @import("std");
fn uuidv7() [16]u8 {
    // random bytes
    var value: [16]u8 = undefined;
    std.crypto.random.bytes(value[6..]);
    // current timestamp in ms
    const timestamp: u48 = @intCast(std.time.milliTimestamp());
    // timestamp
    std.mem.writeInt(u48, value[0..6], timestamp, .big);
    // version and variant
    value[6] = (value[6] & 0x0F) | 0x70;
    value[8] = (value[8] & 0x3F) | 0x80;
    return value;
}
pub fn main() void {
    const uuid_val = uuidv7();
    std.debug.print("{s}\n", .{std.fmt.bytesToHex(uuid_val, .upper)});
}

```

[Edit](https://antonz.org/uuidv7/#edit)

by [Frank Denis](https://github.com/jedisct1)

## Crystal

Initialize a random slice with `rand.random_bytes`, get the current timestamp with `Time.utc`, fill the slice from the timestamp, set version and variant.

```plain text
require "uuid"
class Uuidv7
  @@rand = Random.new
  @uuid : UUID
  forward_missing_to @uuid
  def initialize
    # random bytes
    value = @@rand.random_bytes(16)
    # current timestamp in ms
    timestamp = Time.utc.to_unix_ms
    # timestamp
    timestamp_bytes = StaticArray(UInt8, 8).new(0).to_slice
    IO::ByteFormat::BigEndian.encode(timestamp, timestamp_bytes)
    timestamp_bytes[2..].copy_to(value)
    # version and variant
    value[6] = (value[6] & 0x0F) | 0x70
    value[8] = (value[8] & 0x0F) | 0x80
    @uuid = UUID.new(value)
  end
end
puts Uuidv7.new.hexstring

```

by [lost](https://github.com/lost22git)

## Nim

Initialize a random sequence with `random.rand()`, get the current timestamp with `times.epochTime()`, fill the sequence from the timestamp, set version and variant.

```plain text
import std/[times, strutils, sequtils, sysrand]
proc uuidv7(): seq[byte] =
  # random bytes
  result = urandom(16)
  # current timestamp in ms
  let timestamp = epochTime().uint64 * 1000
  # timestamp
  result[0] = (timestamp shr 40).byte and 0xFF
  result[1] = (timestamp shr 32).byte and 0xFF
  result[2] = (timestamp shr 24).byte and 0xFF
  result[3] = (timestamp shr 16).byte and 0xFF
  result[4] = (timestamp shr 8).byte and 0xFF
  result[5] = timestamp.byte and 0xFF
  # version and variant
  result[6] = (result[6] and 0x0F) or 0x70
  result[8] = (result[8] and 0x3F) or 0x80
var uuidVal = uuidv7()
echo uuidVal.mapIt(it.toHex(2)).join()

```

by [Sultan Al Isaiee](https://github.com/foxoman)

## Gleam

Initialize a random bit array with `crypto.strong_rand_bytes()`, get the current timestamp with `os.system_time()`, write the timestamp, set version and variant into the array.

```plain text
import gleam/int
import gleam/io

@external(erlang, "crypto", "strong_rand_bytes")
pub fn strong_random_bytes(a: Int) -> BitArray

@external(erlang, "os", "system_time")
pub fn system_time(time_unit: Int) -> Int

pub fn uuiv7() -> BitArray {
  let assert <<a:size(12), b:size(62), _:size(6)>> = strong_random_bytes(10)

  let timestamp = system_time(1000)
  let version = 7
  let var = 10

  <<timestamp:48, version:4, a:12, var:2, b:62>>
}

pub fn to_string(ints: BitArray) -> String {
  to_base16(ints, 0, "")
}

fn to_base16(ints: BitArray, position: Int, acc: String) -> String {
  case position {
    8 | 13 | 18 | 23 -> to_base16(ints, position + 1, acc)
    _ ->
      case ints {
        <<i:size(4), rest:bits>> -> {
          to_base16(rest, position + 1, acc <> int.to_base16(i))
        }
        _ -> acc
      }
  }
}

pub fn main() {
  let uuid = uuiv7()
  io.debug(to_string(uuid))
}

```

by [Pavel](https://github.com/brennschlus)

## Tcl

Initialize a random array with `rand()`, get the current timestamp with `clock milliseconds`, fill the array from the timestamp, set version and variant.

```plain text
package require Tcl 8.6
namespace eval uuidv7 {
    namespace export uuidv7
}
proc ::uuidv7::generate { } {
    # random bytes
    set randomBytes {}
    for {set i 0} {$i < 16} {incr i} {
        lappend randomBytes [expr {int(rand() * 256)}]
    }
    # current timestamp in ms
    set timestamp_ms [expr {[clock milliseconds]}]
    # timestamp
    set timestamp_bytes {}
    for {set i 5} {$i >= 0} {incr i -1} {
        lappend timestamp_bytes [expr {($timestamp_ms >> ($i * 8)) & 0xFF}]
    }
    # version and variant
    set bytes [lreplace $randomBytes 0 5 {*}$timestamp_bytes]
    lset bytes 6 [expr {([lindex $bytes 6] & 0x0F) | 0x70}]
    lset bytes 8 [expr {([lindex $bytes 8] & 0x3F) | 0x80}]
    return [binary format c* $bytes]
}
proc ::uuidv7::tostring { uuid } {
    binary scan $uuid H* s
    return [string tolower $s]
}
puts [uuidv7::tostring [uuidv7::generate]]

```

by [Ștefan Alecu](https://github.com/overanalytcl)

## V

Initialize a random array with `rand.bytes()`, get the current timestamp with `time.now()`, fill the array from the timestamp, set version and variant.

```plain text
import rand
import time
fn uuidv7() ![]u8 {
	mut value := rand.bytes(16)!
	// current timestamp in ms
	timestamp := u64(time.now().unix_milli())
	// timestamp
	value[0] = u8((timestamp >> 40) & 0xFF)
	value[1] = u8((timestamp >> 32) & 0xFF)
	value[2] = u8((timestamp >> 24) & 0xFF)
	value[3] = u8((timestamp >> 16) & 0xFF)
	value[4] = u8((timestamp >> 8) & 0xFF)
	value[5] = u8(timestamp & 0xFF)
	// version and variant
	value[6] = (value[6] & 0x0F) | 0x70
	value[8] = (value[8] & 0x3F) | 0x80
	return value
}
fn main() {
	uuid_val := uuidv7()!
	for _, val in uuid_val {
		print('${val:02x}')
	}
	println('')
}

```

by [Laurent Cheylus](https://github.com/lcheylus)

## Emacs Lisp

Initialize a random array with `random`, get the current timestamp with `current-time`, fill the array from the timestamp, set version and variant.

```plain text
(require 'cl-lib)
(defun uuidv7 ()
  "Generates an array representing the bytes of an UUIDv7 label."
  (let* ((timestamp (car (time-convert (current-time) 1000)))
         (timestamp-bytes
          (cl-loop for i from 5 downto 0
                   collect (logand (ash timestamp (* i -8)) #xFF)))
         (uuid (make-vector 16 0)))
    (cl-loop for i below 16 do
	     (aset uuid i
		   (if (< i 6)
		       (nth i timestamp-bytes)
		   (random 256))))
    (aset uuid 6 (logior (logand (elt uuid 6) #x0F) #x70))
    (aset uuid 8 (logior (logand (elt uuid 8) #x3F) #x80))
    uuid))
(defun bytes-to-hexstring (bytes)
  "Converts a vector of bytes into a hexadecimal string."
  (cl-loop for byte across bytes
           concat (format "%02x" byte)))
(let ((uuid-bytes (uuidv7)))
  (message "%s" (bytes-to-hexstring uuid-bytes)))

```

by [Ștefan Alecu](https://github.com/overanalytcl)

## Vimscript

Initialize a random array with `rand`, get the current timestamp with `localtime`, fill the array from the timestamp, set version and variant.

```plain text
function! s:uuidv7() abort	let timestamp = localtime() * 1000	let uuid = []	for i in range(0, 15)		if i < 6			call add(uuid, and(timestamp >> (40 - 8*i), 255))		else			call add(uuid, rand() % 256)		endif	endfor	let uuid[6] = or(and(uuid[6], 15), 112)	let uuid[8] = or(and(uuid[8], 63), 128)	return uuidendfunctionfunction! s:bytes_to_hexstring(bytes) abort	let hexstring = ''	for byte in a:bytes		let hex = printf("%02x", byte)		let hexstring .= hex	endfor	return hexstringendfunctionlet uuid_bytes = s:uuidv7()let hex_uuid = s:bytes_to_hexstring(uuid_bytes)echo hex_uuid
```

by [Ștefan Alecu](https://github.com/overanalytcl)

## Nushell

Generate random bytes with `random int`, get the current timestamp with `date now`, set version and variant, combine everything together.

```plain text
def "random uuid v7" [] {
  # timestamp in ms
  let timestamp_ms = (date now | into int) // 1_000_000
  let timestamp = $timestamp_ms | into binary | bytes at 0..=5 | bytes reverse

  # random bytes
  let rand =  1..=10 | each { random int 0..=255 | into binary | bytes at 0..=0  } | bytes collect

  # version and variant
  let version = $rand | bytes at 0..=0 | bits and 0x[0F] | bits or 0x[70]
  let variant = $rand | bytes at 2..=2 | bits and 0x[3F] | bits or 0x[80]

  [ $timestamp $version ($rand | bytes at 1..=1) $variant ($rand | bytes at 3..) ] | bytes collect
}

random uuid v7 | encode hex

```

by [lost](https://github.com/lost22git)

## Final thoughts

The previous version of the UUID specification ([RFC 4122](https://datatracker.ietf.org/doc/html/rfc4122)) was published in 2005. A K-sorted, time-ordered UUID is a much-needed standard refresher that will hopefully serve us well for years to come.

──

P.S. Interactive examples in this post are powered by [codapi](https://codapi.org/) — an open source tool I'm building. Use it to embed live code snippets into your product docs, online course or blog.

[★ Subscribe](https://antonz.org/subscribe/) to keep up with new posts.
