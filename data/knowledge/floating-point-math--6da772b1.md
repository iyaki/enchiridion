---
title: "Floating Point Math"
notion_id: 6da772b1-b786-441d-a5bc-4e38562d6b8f
notion_url: https://app.notion.com/p/Floating-Point-Math-6da772b1b786441da5bc4e38562d6b8f
last_edited: 2023-01-25T18:31:00.000Z
source_url: https://0.30000000000000004.com/
tags: ["English", "Programming", "Website", "Article"]
---


## 









| Language | Code | Result |
| --- | --- | --- |
|  |  |  |
| PowerShell by default uses double type, but because it runs on .NET it has the same types as C# does. Thanks to that the [Decimal type](https://docs.microsoft.com/en-us/dotnet/api/system.decimal?view=net-5.0) can be used - directly by providing the type name `[decimal]` or via [suffix ](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_numeric_literals?view=powershell-7.1#real-literals)[`d`](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_numeric_literals?view=powershell-7.1#real-literals). <br>More about that in the [C# section](https://0.30000000000000004.com/#csharp). |  |  |
| **  ABAP ** | <br><br>`WRITE / CONV f( '.1' + '.2' ).` and <br><br>`WRITE / CONV decfloat16( '.1' + '.2' ).` | <br><br>`0.30000000000000004` and <br><br>`0.3` |
| **  APL ** | <br><br>`0.1 + 0.2` and <br><br>`⎕PP ← 17<br>0.1 + 0.2` and <br><br>`0.3 = 0.1 + 0.2` and <br><br>`⎕CT←0<br>0.3 = 0.1 + 0.2` and <br><br>`⎕FR ← 1287<br>⎕PP ← 34<br>0.1 + 0.2` and <br><br>`⎕FR ← 1287<br>⎕DCT ← 0<br>0.3 = 0.1 + 0.2` | <br><br>`0.3` and <br><br>`0.30000000000000004` and <br><br>`1` and <br><br>`0` and <br><br>`0.3` and <br><br>`1` |
| [APL](https://aplwiki.com/) has a default [printing precision of 10 significant digits](http://microapl.com/apl_help/ch_020_070_640.htm). Setting `⎕PP` to 17 shows the error, however `0.3 = 0.1 + 0.2` is still true (`1`) because there’s a default [comparison tolerance of about 10](https://help.dyalog.com/latest/#Language/System%20Functions/ct.htm)[-14](https://help.dyalog.com/latest/#Language/System%20Functions/ct.htm). Setting `⎕CT` to 0 shows the inequality. [Dyalog APL](https://aplwiki.com/wiki/Dyalog_APL) also supports 128-bit decimal numbers (activated by setting the float representation, `⎕FR`, to 1287, i.e. 128-bit decimal), where even setting the decimal comparison tolerance (`⎕DCT`) to zero still makes the equation hold true. [Try it online!](https://tio.run/##SyzI0U2pTMzJT/8PBAZ6hgraCgZ6RlyP@qYGBCg8apugYGjOhRA20DNWsFVAUeYcAlRlgE3GLQhigJGFOZJ5xiZc@FS5OIeA@RgmAgA) Multi-precision floats, unlimited precision rationals, and ball arithmetic are available in [NARS2000](https://aplwiki.com/wiki/NARS2000). |  |  |
| **  Ada ** | <br><br>`with Ada.Text_IO; use Ada.Text_IO;<br>procedure Sum is<br>  A : Float := 0.1;<br>  B : Float := 0.2;<br>  C : Float := A + B;<br>begin<br>  Put_Line(Float'Image(C));<br>  Put_Line(Float'Image(0.1 + 0.2));<br>end Sum;` | <br><br>`3.00000E-01  <br>3.00000E-01` |
| **  AutoHotkey ** | <br><br>`MsgBox, % 0.1 + 0.2` | <br><br>`0.3` |
| **  C ** | <br><br>`#include <stdio.h><br><br>int main(int argc, char** argv) {<br>  printf("%.17f\n", .1 + .2);<br>  return 0;<br>}` | <br><br>`0.30000000000000004` |
| **  C# ** | <br><br>`Console.WriteLine("{0:R}", .1 + .2);` and <br><br>`Console.WriteLine("{0:R}", .1f + .2f);` and <br><br>`Console.WriteLine("{0:R}", .1m + .2m);` | <br><br>`0.30000000000000004` and <br><br>`0.3` and <br><br>`0.3` |
| [C# has support for 128-bit decimal numbers](https://msdn.microsoft.com/en-us/library/364x0z75.aspx), with 28-29 significant digits of precision. Their range, however, is smaller than that of both the single and double precision floating point types. Decimal literals are denoted with the `m` suffix. |  |  |
| **  C++ ** | <br><br>`#include <iomanip><br>#include <iostream><br><br>int main() {<br>  std::cout << std::setprecision(17) << 0.1 + 0.2;<br>}` | <br><br>`0.30000000000000004` |
| **  Clojure ** | <br><br>`(+ 0.1 0.2)` | <br><br>`0.30000000000000004` |
| Clojure supports arbitrary precision and ratios. `(+ 0.1M 0.2M)` returns `0.3M`, while `(+ 1/10 2/10)` returns `3/10`. |  |  |
| **  ColdFusion ** | <br><br>`<cfset foo = .1 + .2><br><cfoutput>#foo#</cfoutput>` | <br><br>`0.3` |
| **  Common Lisp ** | <br><br>`(+ .1 .2)` and <br><br>`(+ 1/10 2/10)` and <br><br>`(+ 0.1d0 0.2d0)` and <br><br>`(- 1.2 1.0)` | <br><br>`0.3` and <br><br>`3/10` and <br><br>`0.30000000000000004d0` and <br><br>`0.20000005` |
| CL’s spec doesn’t actually even require radix-2 floats (let alone specifically 32-bit singles and 64-bit doubles), but the high-performance implementations all seem to use IEEE floats with the usual sizes. This was tested on SBCL and ECL in particular. |  |  |
| **  Crystal ** | <br><br>`puts 0.1 + 0.2` and <br><br>`puts 0.1_f32 + 0.2_f32` | <br><br>`0.30000000000000004` and <br><br>`0.3` |
| **  D ** | <br><br>`import std.stdio;<br><br>void main(string[] args) {<br>  writefln("%.17f", .1+.2);<br>  writefln("%.17f", .1f+.2f);<br>  writefln("%.17f", .1L+.2L);<br>}` | <br><br>`0.29999999999999999  <br>0.30000001192092896  <br>0.30000000000000000` |
| **  Dart ** | <br><br>`print(.1 + .2);` | <br><br>`0.30000000000000004` |
| **  Delphi XE5 ** | <br><br>`writeln(0.1 + 0.2);` | <br><br>`0.3` |
| **  Elixir ** | <br><br>`IO.puts(0.1 + 0.2)` | <br><br>`0.30000000000000004` |
| **  Elm ** | <br><br>`0.1 + 0.2` | <br><br>`0.30000000000000004` |
| **  Elvish ** | <br><br>`+ .1 .2` | <br><br>`0.30000000000000004` |
| Elvish uses Go’s `double` for numerical operations. |  |  |
| **  Emacs Lisp ** | <br><br>`(+ .1 .2)` | <br><br>`0.30000000000000004` |
| **  Erlang ** | <br><br>`io:format("~w~n", [0.1 + 0.2]).<br>io:format("~f~n", [0.1 + 0.2]).<br>io:format("~e~n", [0.1 + 0.2]).<br>io_lib:format("~.1f~n", [0.1 + 0.2]).<br>io_lib:format("~.2f~n", [0.1 + 0.2]).` | <br><br>`0.30000000000000004<br>0.300000<br>3.00000e-1<br>"0.3\n"<br>"0.30\n"` |
| **  FORTRAN ** | <br><br>`program FLOATMATHTEST<br>  real(kind=4) :: x4, y4<br>  real(kind=8) :: x8, y8<br>  real(kind=16) :: x16, y16<br>  ! REAL literals are single precision, use _8 or _16<br>  ! if the literal should be wider.<br>  x4 = .1; x8 = .1_8; x16 = .1_16<br>  y4 = .2; y8 = .2_8; y16 = .2_16<br>  write (*,*) x4 + y4, x8 + y8, x16 + y16<br>end` | <br><br>`0.300000012  <br>0.30000000000000004  <br>0.300000000000000000000000000000000039` |
| **  Fish ** | <br><br>`math .1 + .2` | <br><br>`0.3` |
| **  GHC (Haskell) ** | <br><br>`0.1 + 0.2 :: Double` and <br><br>`0.1 + 0.2 :: Float` and <br><br>`0.1 + 0.2 :: Rational` | <br><br>`0.30000000000000004` and <br><br>`0.3` and <br><br>`3 % 10` |
| If you need real numbers, packages like [exact-real](https://hackage.haskell.org/package/exact-real) give you the correct answer. |  |  |
| **  GNU Octave ** | <br><br>`0.1 + 0.2` and <br><br>`single(0.1)+single(0.2)` and <br><br>`double(0.1)+double(0.2)` and <br><br>`0.1+single(0.2)` and <br><br>`0.1+double(0.2)` and <br><br>`sprintf('%.17f',0.1+0.2)` | <br><br>`0.3` and <br><br>`0.3` and <br><br>`0.3` and <br><br>`0.3` and <br><br>`0.3` and <br><br>`0.30000000000000004` |
| **  Gforth ** | <br><br>`0.1e 0.2e f+ f.` and <br><br>`0.1e 0.2e f+ 0.3e f= .` and <br><br>`0.3e 0.3e f= .` | <br><br>`0.3` and <br><br>`0` and <br><br>`-1` |
| In Gforth `0` means `false` and `-1` means `true`. First example print `0.3` but it’s not equal to actuall `0.3`. |  |  |
| **  Go ** | <br><br>`package main<br>import "fmt"<br><br>func main() {<br>  fmt.Println(.1 + .2)<br>  var a float64 = .1<br>  var b float64 = .2<br>  fmt.Println(a + b)<br>  fmt.Printf("%.54f\n", .1 + .2)<br>}` | <br><br>`0.3  <br>0.30000000000000004  <br>0.299999999999999988897769753748434595763683319091796875` |
| Go numeric constants [have arbitrary precision](http://blog.golang.org/constants#TOC_8.). |  |  |
| **  Groovy ** | <br><br>`println 0.1 + 0.2` | <br><br>`0.3` |
| Literal decimal values in Groovy are instances of [java.math.BigDecimal](https://docs.oracle.com/javase/8/docs/api/java/math/BigDecimal.html). |  |  |
| **  Guile ** | <br><br>`(+ 0.1 0.2)` and <br><br>`(+ 1/10 2/10)` | <br><br>`0.30000000000000004` and <br><br>`3/10` |
| **  Hugs (Haskell) ** | <br><br>`0.1 + 0.2` | <br><br>`0.3` |
| **  Io ** | <br><br>`(0.1 + 0.2) print` | <br><br>`0.3` |
| **  Java ** | <br><br>`System.out.println(.1 + .2);` and <br><br>`System.out.println(.1F + .2F);` | <br><br>`0.30000000000000004` and <br><br>`0.3` |
| Java has built-in support for arbitrary-precision numbers using the [BigDecimal](http://docs.oracle.com/javase/8/docs/api/java/math/BigDecimal.html) class. |  |  |
| **  JavaScript ** | <br><br>`console.log(.1 + .2);` | <br><br>`0.30000000000000004` |
| The [decimal.js](http://mikemcl.github.io/decimal.js/) library provides an arbitrary-precision Decimal type for JavaScript. |  |  |
| **  Julia ** | <br><br>`.1 + .2` | <br><br>`0.30000000000000004` |
| Julia has built-in [rational numbers support](https://docs.julialang.org/en/v1/manual/complex-and-rational-numbers/#Rational-Numbers-1) and also a built-in [arbitrary-precision BigFloat](https://docs.julialang.org/en/v1/manual/integers-and-floating-point-numbers/#Arbitrary-Precision-Arithmetic-1) data type. To get the math right, `1//10 +<br>2//10` returns `3//10`. |  |  |
| **  K (Kona) ** | <br><br>`0.1 + 0.2` | <br><br>`0.3` |
| **  Kotlin ** | <br><br>`println(.1 + .2)` and <br><br>`println(.1F + .2F)` | <br><br>`0.30000000000000004` and <br><br>`0.3` |
| See [Reference documentation](https://kotlinlang.org/docs/reference/basic-types.html). |  |  |
| **  Lua ** | <br><br>`print(.1 + .2)` and <br><br>`print(string.format("%0.17f", 0.1 + 0.2))` | <br><br>`0.3` and <br><br>`0.30000000000000004` |
| **  MATLAB ** | <br><br>`0.1 + 0.2` and <br><br>`sprintf('%.17f', 0.1 + 0.2)` | <br><br>`0.3` and <br><br>`0.30000000000000004` |
| **  MIT/GNU Scheme ** | <br><br>`(+ 0.1 0.2)` and <br><br>`(+ \#e0.1 \#e0.2)` | <br><br>`0.30000000000000004` and <br><br>`3/10` |
| The scheme specification has a concept [exactness](https://people.csail.mit.edu/jaffer/r3rs_8.html#SEC48). |  |  |
| **  Mathematica ** | <br><br>`0.1 + 0.2` | <br><br>`0.3` |
| Mathematica has a fairly thorough internal mechanism for dealing with [numerical precision](https://reference.wolfram.com/language/tutorial/Numbers.html#21155) and supports arbitrary precision. <br>[By default](https://reference.wolfram.com/language/tutorial/MachinePrecisionNumbers.html), the inputs `0.1` and `0.2` in the example are taken to have [MachinePrecision](https://reference.wolfram.com/language/ref/FullForm.html). At a common `MachinePrecision` of `15.9546` digits, `0.1 + 0.2` actually has a [FullForm][4] of `0.30000000000000004`, but is printed as `0.3`. <br>Mathematica supports rational numbers: `1/10 + 2/10` is `3/10` (which has a `FullForm` of `Rational[3, 10]`). |  |  |
| **  MySQL ** | <br><br>`SELECT .1 + .2;` | <br><br>`0.3` |
| **  Nim ** | <br><br>`echo(0.1 + 0.2)` | <br><br>`0.3` |
| **  OCaml ** | <br><br>`0.1 +. 0.2;;` | <br><br>`float = 0.300000000000000044` |
| **  Objective-C ** | <br><br>`#import <Foundation/Foundation.h><br><br>int main(int argc, const char * argv[]) {<br>  @autoreleasepool {<br>    NSLog(@"%.17f\n", .1+.2);<br>  }<br>  return 0;<br>}` | <br><br>`0.30000000000000004` |
| **  PHP ** | <br><br>`echo .1 + .2;` and <br><br>`var_dump(.1 + .2);` and <br><br>`var_dump(bcadd(.1, .2, 1));` | <br><br>`0.3` and <br><br>`float(0.30000000000000004441)` and <br><br>`string(3) "0.3"` |
| PHP `echo` converts `0.30000000000000004441` to a string and shortens it to “0.3”. To achieve the desired floating-point result, adjust the precision setting: `ini_set("precision", 17)`. |  |  |
| **  Perl ** | <br><br>`perl -E 'say 0.1+0.2'` and <br><br>`perl -e 'printf q{%.17f}, 0.1+0.2'` and <br><br>`perl -MMath::BigFloat -E 'say Math::BigFloat->new(q{0.1}) + Math::BigFloat->new(q{0.2})'` | <br><br>`0.3` and <br><br>`0.30000000000000004` and <br><br>`0.3` |
| The addition of float primitives only appears to print correctly because [not all of the 17 digits are printed by default](https://github.com/perl/perl5/issues/15119). The core [Math::BigFloat](https://metacpan.org/pod/Math::BigFloat) allows true arbitrary precision floating point operations by never using numeric primitives. |  |  |
| **  PicoLisp ** | <br><br>`[load "frac.min.l"]<br>[println (+ (/ 1 10) (/ 2 10))]` | <br><br>`(/ 3 10)` |
| You must [load file “frac.min.l”](https://gist.github.com/DKordic/6016d743c4c124a1c04fc12accf7ef17/raw/cde44c880c51c79ec3a93ea17b9fec93db8e149f). |  |  |
| **  PostgreSQL ** | <br><br>`SELECT 0.1::float + 0.2::float;` and <br><br>`SELECT 0.1 + 0.2;` | <br><br>`0.30000000000000004` and <br><br>`0.3` |
| PostgreSQL treats decimal literals as [arbitrary precision numbers with fixed point](https://www.postgresql.org/docs/12/datatype-numeric.html#DATATYPE-NUMERIC-DECIMAL). Explicit type casts are required to get floating-point numbers. <br>PostgreSQL 11 and earlier outputs `0.3` as a result for query `SELECT 0.1::float + 0.2::float;`, but the result is rounded only for display, and under the hood it is still good old `0.30000000000000004`. <br>In PostgreSQL 12 default behavior for textual output of floats was changed from more human-readable rounded format to shortest-precise format. Format can be customized by the [`extra_float_digits`](https://www.postgresql.org/docs/12/runtime-config-client.html#GUC-EXTRA-FLOAT-DIGITS) configuration parameter. |  |  |
| **  Prolog (SWI-Prolog) ** | <br><br>`?- X is 0.1 + 0.2.` | <br><br>`X = 0.30000000000000004.` |
| **  Pyret ** | <br><br>`0.1 + 0.2` and <br><br>`~0.1 + ~0.2` | <br><br>`0.3` and <br><br>`~0.30000000000000004` |
| Pyret has built-in support for both rational numbers and floating points. Numbers written normally are assumed to be exact. In contrast, RoughNums are represented by floating points, and are written prefixed with a `~`, indicating that they are not precise answers – the `~` is meant to visually evoke hand-waving. A user who sees a computation produce `~0.30000000000000004` knows to treat the value with skepticism. RoughNums cannot be compared directly for equality; they can only be compared up to a given tolerance. |  |  |
| **  Python 2 ** | <br><br>`print .1 + .2` and <br><br>`.1 + .2` and <br><br>`float(decimal.Decimal(".1") + decimal.Decimal(".2"))` and <br><br>`float(fractions.Fraction('0.1') + fractions.Fraction('0.2'))` | <br><br>`0.3` and <br><br>`0.30000000000000004` and <br><br>`0.3` and <br><br>`0.3` |
| Python 2’s `print` statement converts `0.30000000000000004` to a string and shortens it to “0.3”. To achieve the desired floating point result, use `print repr(.1 + .2)`. This was fixed in Python 3 (see below). |  |  |
| **  Python 3 ** | <br><br>`print(.1 + .2)` and <br><br>`.1 + .2` and <br><br>`float(decimal.Decimal('.1') + decimal.Decimal('.2'))` and <br><br>`float(fractions.Fraction('0.1') + fractions.Fraction('0.2'))` | <br><br>`0.30000000000000004` and <br><br>`0.30000000000000004` and <br><br>`0.3` and <br><br>`0.3` |
| Python (both 2 and 3) supports decimal arithmetic with the [decimal](https://docs.python.org/3/library/decimal.html) module, and true rational numbers with the [fractions](https://docs.python.org/3.7/library/fractions.html) module. |  |  |
| **  R ** | <br><br>`print(.1 + .2)` and <br><br>`print(.1 + .2, digits=18)` | <br><br>`0.3` and <br><br>`0.30000000000000004` |
| **  Racket (PLT Scheme) ** | <br><br>`(+ .1 .2)` and <br><br>`(+ 1/10 2/10)` | <br><br>`0.30000000000000004` and <br><br>`3/10` |
| **  Raku ** | <br><br>`raku -e 'say 0.1 + 0.2'` and <br><br>`raku -e 'say (0.1 + 0.2).fmt(\"%.17f\")'` and <br><br>`raku -e 'say 1/10 + 2/10'` and <br><br>`raku -e 'say 0.1e0 + 0.2e0'` | <br><br>`0.3` and <br><br>`0.30000000000000000` and <br><br>`0.3` and <br><br>`0.30000000000000004` |
| [Raku uses rationals by default](https://docs.raku.org/type/Rational), so `.1` is stored something like `{<br>numerator => 1, denominator => 10 }`. To actually trigger the behavior, you must force the numbers to be of type Num (double in C terms) and use the base function instead of the `sprintf` or `fmt` functions (since those functions have a bug that limits the precision of the output). |  |  |
| **  Regina REXX ** | <br><br>`say '.1+.2'` | <br><br>`0.3` |
| **  Ruby ** | <br><br>`puts 0.1 + 0.2` and <br><br>`puts 1/10r + 2/10r` | <br><br>`0.30000000000000004` and <br><br>`3/10` |
| Ruby supports rational numbers in syntax with version 2.1 and newer directly. For older versions use [Rational](http://ruby-doc.org/core/classes/Rational.html). Ruby also has a library specifically for decimals: [BigDecimal](http://ruby-doc.org/stdlib/libdoc/bigdecimal/rdoc/index.html). |  |  |
| **  Rust ** | <br><br>`extern crate num;<br>use num::rational::Ratio;<br><br>fn main() {<br>  println!("{}", 0.1 + 0.2);<br>  println!("{}", 0.1_f32 + 0.2_f32);<br>  println!("1/10 + 2/10 = {}", Ratio::new(1, 10) + Ratio::new(2, 10));<br>}` | <br><br>`0.30000000000000004<br>0.3<br>1/10 + 2/10 = 3/10` |
| Rust has [rational number support](https://rust-num.github.io/num/num_rational/struct.Ratio.html) from the [num crate](https://crates.io/crates/num). |  |  |
| **  SageMath ** | <br><br>`.1 + .2` and <br><br>`RDF(.1) + RDF(.2)` and <br><br>`RBF('.1') + RBF('.2')` and <br><br>`QQ('1/10') + QQ('2/10')` | <br><br>`0.3` and <br><br>`0.30000000000000004` and <br><br>`["0.300000000000000 +/- 1.64e-16"]` and <br><br>`3/10` |
| [SageMath](https://www.sagemath.org/) supports [various fields](http://doc.sagemath.org/html/en/reference/rings_numerical/index.html) for arithmetic: [Arbitrary Precision Real Numbers](http://doc.sagemath.org/html/en/reference/rings_numerical/sage/rings/real_mpfr.html), [RealDoubleField](http://doc.sagemath.org/html/en/reference/rings_numerical/sage/rings/real_double.html), [Ball Arithmetic](http://doc.sagemath.org/html/en/reference/rings_numerical/sage/rings/real_arb.html), [Rational Numbers](http://doc.sagemath.org/html/en/reference/rings_standard/sage/rings/rational_field.html), etc. |  |  |
| **  Scala ** | <br><br>`scala -e 'println(0.1 + 0.2)'` and <br><br>`scala -e 'println(0.1F + 0.2F)'` and <br><br>`scala -e 'println(BigDecimal(\"0.1\") + BigDecimal(\"0.2\"))'` | <br><br>`0.30000000000000004` and <br><br>`0.3` and <br><br>`0.3` |
| **  Smalltalk ** | <br><br>`(1/10) + (2/10).` and <br><br>`0.1 + 0.2.` and <br><br>`0.1s17 + 0.2s17.` | <br><br>`(3/10)` and <br><br>`0.30000000000000004` and <br><br>`0.30000000000000000s17` |
| Smalltalk uses fractions by default in most operations; in fact, standard devision results in fractions, not floating point numbers. Squeak and similar Smalltalks provide “scaled decimals” that allow fixed-point real numbers (`s`-suffix indicating precision places). |  |  |
| **  Swift ** | <br><br>`0.1 + 0.2` and <br><br>`Decimal(0.1) + Decimal(0.2)` | <br><br>`0.30000000000000004` and <br><br>`0.3` |
| Swift supports decimal arithmetic with the [Foundation](https://developer.apple.com/documentation/foundation/decimal) module. |  |  |
| **  TCL ** | <br><br>`puts [expr .1 + .2]` | <br><br>`0.30000000000000004` |
| **  Turbo Pascal 7.0 ** | <br><br>`writeln(0.1 + 0.2);` | <br><br>`0.3` |
| **  Vala ** | <br><br>`static int main(string[] args) {<br>  stdout.printf("%.17f\n", 0.1 + 0.2);<br>  return 0;<br>}` | <br><br>`0.30000000000000004` |
| **  Visual Basic 6 ** | <br><br>`a# = 0.1 + 0.2: b# = 0.3<br>Debug.Print Format(a - b, "0." & String(16, "0"))<br>Debug.Print a = b` | <br><br>`0.0000000000000001  <br>False` |
| Appending the identifier type character `#` to any identifier forces it to Double. |  |  |
| **  WebAssembly (WAST) ** | <br><br>`(func $add_f32 (result f32)<br>  f32.const 0.1<br>  f32.const 0.2<br>  f32.add)<br>(export "add_f32" (func $add_f32))` and <br><br>`(func $add_f64 (result f64)<br>  f64.const 0.1<br>  f64.const 0.2<br>  f64.add)<br>(export "add_f64" (func $add_f64))` | <br><br>`0.30000001192092896` and <br><br>`0.30000000000000004` |
| See [demo](https://webassembly.studio/?f=r739k6d6q4t). |  |  |
| **  awk ** | <br><br>`awk 'BEGIN { print 0.1 + 0.2 }'` | <br><br>`0.3` |
| **  bc ** | <br><br>`0.1 + 0.2` | <br><br>`0.3` |
| **  dc ** | <br><br>`0.1 0.2 + p` | <br><br>`0.3` |
| **  ivy ** | <br><br>`0.1 + 0.2` and <br><br>`0.1 + sqrt(0.04)` | <br><br>`3/10` and <br><br>`0.3` |
| [Ivy](https://pkg.go.dev/robpike.io/ivy) is an interpreter for an APL-like language. It uses exact rational arithmetic so it can handle arbitrary precision. When ivy evaluates an irrational function, the result is stored in a high-precision floating-point number (default 256 bits of mantissa). |  |  |
| **  zsh ** | <br><br>`echo "$((.1 + .2))"` | <br><br>`0.30000000000000004` |
