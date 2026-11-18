---
title: "Base64 Encoding, Explained"
notion_id: 7575e9c6-1ba3-4f8f-8b4f-ddfd5a38fdbf
notion_url: https://app.notion.com/p/Base64-Encoding-Explained-7575e9c61ba34f8f8b4fddfd5a38fdbf
last_edited: 2026-09-21T16:58:00.000Z
source_url: https://writesoftwarewell.com/base64-encoding-explained
tags: ["English", "Programming", "Article", "Akshay's Blog"]
---
When you're programming, it's easy to get by with a superficial understanding of many things. You can easily fool yourself by thinking that you are programming when you are blindly copy + pasting code from Stack Overflow or some random article you stumbled upon.

Base64 encoding was one of these topics that was bugging me for a while. I often came across Base64 encoded images or URLs, and had no idea whatsoever it meant or why it was even used. Finally, I decided to do some research to fill that knowledge gap, and spent the Sunday reading [RFC 4648](https://datatracker.ietf.org/doc/html/rfc4648?ref=akshaykhot.com) (my idea of a fun weekend).

What follows is everything I learned about Base64 encoding.

## What is Base64 Encoding?

Base64 encoding takes binary data and converts it into text, specifically ASCII text. The resulting text contains only letters from `A-Z`, `a-z`, numbers from `0-9`, and the symbols `+` and `/`.

As there are 26 letters in the alphabet, we have `26 + 26 + 10 + 2` characters. Hence this encoding is named **`Base64`**. These 64 characters are considered "safe", that is, they cannot be misinterpreted by legacy computers and programs unlike characters such as `<`, `>`, `\n` and many others.

Here's what the text "Ruby on Rails" looks like when Base64 encoded: **UnVieSBvbiBSYWlscw==**.

**It's important to remember that we are not encrypting the text here.** Given Base64 encoded data, it's very easy to convert it back (decode) to the original text. We are only changing the representation of the data, i.e. encoding.

In its essence, Base64 encoding uses a specific, reduced set of characters to encode binary data, to prevent against data corruption.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The Base64 Alphabet

As there are only 64 characters available to encode into, we can represent them using only 6 bits, because `2^6 = 64`. Every Base64 digit represents 6 bits of data. There are 8 bits in a byte, and the closest common multiple of 8 and 6 is 24. So 24 bits, or 3 bytes, can be represented using four 6-bit Base64 digits.

_(If that last paragraph totally went over your head, __**don't worry**__. Hopefully it should be clear by the end of this post.)_

## Why Base64?

You must have included an image in your HTML document using the `<img src="nature.jpeg">` tag. **Did you know you can embed the image data directly into the HTML without linking to the external image file?** [Data URLs](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs?ref=akshaykhot.com) let you do this, and they use Base64 encoded text to embed files inline.

```plain text
<img src="data:image/gif;base64,xxxxbase64encodedtextxxxx">

data:[<mime type>][;charset=<charset>][;base64],<encoded data>
```

Historically it has been used to encode binary data in email messages where the email server might modify line-endings. A more modern example is the use of Base64 encoding to [embed image data directly in HTML source code](http://www.sweeting.org/mark/blog/2005/07/12/base64-encoded-images-embedded-in-html?ref=akshaykhot.com). Here it is necessary to encode the data to avoid characters like '<' and '>' being interpreted as tags.

From: [Why do we use Base64?](https://stackoverflow.com/questions/3538021/why-do-we-use-base64?ref=akshaykhot.com)

Another common use case is when we have to store or transmit some binary data over the network that's supposed to handle text, or US-ASCII data. This ensures data remains unchanged during transport. Base64 can also be used for passing data in URLs when that data includes non-URL friendly characters.

Base encoding is also used in many applications simply because it makes it possible to manipulate objects with text editors.

You can also transfer files as text, using Base64 encoding. First, get the file's bytes and encode them as Base64. Then transfer the Base64 encoded string, and then decode it back to the original file content on the receiving side.

Let's take a deeper look into this algorithm in the next section.

## Base64 Encoding Algorithm

Here's the simple algorithm that converts some text into Base64.

1. Convert the text to its binary representation.
2. Divide the bits into groups of 6 bits each.
3. Convert each group to a decimal number from 0-63. It cannot be greater than 64 as there are only 6 bits in each group.
4. Convert this decimal number to the equivalent Base64 character using the Base64 alphabet.

That's it. You have a Base64 encoded string. If there're insufficient bits in the final group, you can use `=` or `==` as padding.

Sounds confusing? Don't worry, the following example should make it pretty clear. Let's convert my name "Akshay" to its Base64 equivalent string.

- Convert the text "Akshay" to binary by first converting each character to its corresponding [ASCII number](https://www.ascii-code.com/?ref=akshaykhot.com) and then converting that decimal number to binary (or just use [this tool](https://www.rapidtables.com/convert/number/ascii-to-binary.html?ref=akshaykhot.com)):

```plain text
01000001 01101011 01110011 01101000 01100001 01111001

   A        k        s        h        a        y
```

- Divide the bits into groups of 6 bits:

```plain text
010000 010110 101101 110011 011010 000110 000101 111001
```

- Convert each group to a decimal number between 0 to 63:

```plain text
010000 010110 101101 110011 011010 000110 000101 111001

  16     22     45     51     26     6      5      57
```

- Now use the Base64 alphabet (see above image) to convert each decimal number to its Base64 representation:

```plain text
16  22  45  51  26  6  5  57

Q   W   t   z   a   G  F  5
```

And we're done. The name "Akshay" is represented in Base64 as `QWtzaGF5`.

At first glance, the benefit of Base64 encoding is not quite obvious. **What exactly did we achieve by converting "Akshay" to "QWtzaGF5"?**

Imagine, instead of my name, you had an image or a sensitive file (PDF, text, video, anything, really), and you wanted to store it as text. You could first convert it to binary, and then Base64 encode it to get corresponding ASCII text.

Now you could send or store that text anywhere and anyhow you like, without worrying whether some legacy device, protocol or software won't misinterpret the raw binary data to corrupt your file. Makes sense?

## How to Encode and Decode Base64

All programming languages have support for encoding and decoding data to and from the Base64 format.

Here is the Ruby code that takes some text as input and converts it into Base-64 encoded string.

```plain text
require "base64"

encoded = Base64.encode64("Ruby on Rails")  # "UnVieSBvbiBSYWlscw==\n"

decoded = Base64.decode64(encoded)  # "Ruby on Rails"
```

Here's the equivalent program in C#, my second-most favorite language:

```plain text
public static string ToBase64(string value)
{
    byte[] bytes = System.Text.Encoding.ASCII.GetBytes(value);

    string base64 = Convert.ToBase64String(bytes);

    return base64;
}

public static string FromBase64(string encoded)
{
    byte[] data = System.Convert.FromBase64String(encodedString);

    string decodedString = System.Text.Encoding.UTF8.GetString(data);
}
```

PHP makes it very simple with its `base64_encode` and `base64_decode` top-level functions.

```plain text
<?php
$str = "Ruby on Rails";
echo base64_encode($str);

$str = "UnVieSBvbiBSYWlscw==\n";
echo base64_decode($str);
?>
```

Similarly, in JavaScript, use the `btoa()` to encode and `atob()` functions to encode and decode the text.

```plain text
const text = "Ruby on Rails"
btoa(text) // "UnVieSBvbiBSYWlscw==\n"

const encoded_text = "UnVieSBvbiBSYWlscw==\n"
atob(encoded_text)  // Ruby on Rails
```

What's more, your terminal has built-in support for Base64 encoding. Try this in terminal:

```plain text
$ echo "akshay" | base64
YWtzaGF5Cg==

$ echo "YWtzaGF5Cg==" | base64 -d
akshay
```

That's a wrap. I hope you found this article helpful and you learned something new. If you are interested in learning more, I highly recommend you read [RFC 4648](https://datatracker.ietf.org/doc/html/rfc4648?ref=akshaykhot.com), which describes the Base64 encoding in detail.

As always, if you have any questions or feedback, didn't understand something, or found a mistake, please leave a comment below or [send me an email](mailto:akshay.khot@hey.com?ref=akshays-blog). I reply to all emails I get from developers, and I look forward to hearing from you.

If you'd like to receive future articles directly in your email, please [subscribe to my blog](https://www.akshaykhot.com/base64-encoding-explained/?utm_source=tldrwebdev#/portal/signup). If you're already a subscriber, thank you.
