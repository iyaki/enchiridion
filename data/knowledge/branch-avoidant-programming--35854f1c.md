---
title: "Branch‑Avoidant Programming"
notion_id: 35854f1c-7d23-81ef-bbe2-e18374bd5d5e
notion_url: https://app.notion.com/p/Branch-Avoidant-Programming-35854f1c7d2381efbbe2e18374bd5d5e
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://easylang.online/blog/branchless
tags: ["Medium", "English", "Programming", "Performance", "System Design / Software Architecture", "Article"]
---
With modern CPUs, **avoiding branch mispredictions** is a key method for speeding up programs. One of the most effective ways to reduce mispredictions is to simply avoid branches.

Let’s consider a simple task: Iterate through an array and copy all numbers less than 500 into a new array. A straightforward implementation in C looks like this:

```plain text
for (int i = 0, j = 0; i < 1000; i++) {
    if (numbers[i] < 500) {
        small_numbers[j] = numbers[i];
        j += 1;
    }
}
```

If the numbers are randomly distributed, the result of the ‘if’ condition becomes unpredictable for the CPU’s branch predictor. This results in a high misprediction rate, significantly impairing performance because the processor must repeatedly empty its pipeline and restart execution.

We can avoid this problem by using a branchless approach:

```plain text
for (int i = 0, j = 0; i < 1000; i++) {
    small_numbers[j] = numbers[i];
    j += (numbers[i] < 500);
}
```

While this version introduces an unconditional store (writing to memory even if the condition isn’t met), its cost is typically far lower than the heavy penalty of a branch misprediction.

**Performance Comparison**

To get meaningful results, we run both routines for 1000 iterations using 100,000 random numbers.

```plain text
// SPDX-License-Identifier: MIT
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>

#define SIZE 100000
int numbers[SIZE];
int small_numbers[SIZE];

double ts(void) {
    static double t0;
    struct timeval tv;
    gettimeofday(&tv, NULL);
    double h = t0;
    t0 = tv.tv_sec + tv.tv_usec / 1000000.0;
    return t0 - h;
}
void init(void) {
    srand(1);
    for (int i = 0; i < SIZE; i++) numbers[i] = rand() % 1000;
    ts();
}
int small_if(void) {
    int i, j;
    for (i = 0, j = 0; i < SIZE; i++) {
        if (numbers[i] < 500) {
            small_numbers[j] = numbers[i];
            j += 1;
        }
    }
    return j;
}
int small_bl(void) {
    int i, j;
    for (i = 0, j = 0; i < SIZE; i++) {
        small_numbers[j] = numbers[i];
        j += (numbers[i] < 500);
    }
    return j;
}

int main(void) {
    init();
    for (int i = 0; i < 1000; i++) small_if();
    printf("Time if: %.3f\n", ts());
    init();
    for (int i = 0; i < 1000; i++) small_bl();
    printf("Time bl: %.3f\n", ts());
}
```

Benchmark results were obtained on an Apple M1 with `-O1` optimization. On this benchmark, `-O3` did not yield additional speedup, only larger generated code.

```plain text
Time if: 0.329
Time bl: 0.036
```

**Assembly Analysis**

Using `objdump -d --no-show-raw-insn`, we can see the difference in how the CPU processes these loops.

With the _if_ version, the compiler generates a conditional branch `b.gt`.

```plain text
ldr     w12, [x9], #0x4            ; Load numbers[i]
cmp     w12, #0x1f3                ; Compare with 499 (0x1f3)
b.gt    0x100000624                ; Branch if greater (The bottleneck)
str     w12, [x10, w8, sxtw #2]    ; Store value in small_numbers[j]
add     w8, w8, #0x1               ; j += 1
```

With the Branchless version the compiler uses the `cinc` (Conditional Increment) instruction.

```plain text
ldr     w12, [x9], #0x4            ; Load numbers[i]
str     w12, [x10, w8, uxtw #2]    ; Unconditional store to small_numbers[j]
cmp     w12, #0x1f4                ; Compare with 500
cinc    w8, w8, lt                 ; Conditional increment (No branch)
```

In this version, there is no jump based on the data. The control flow is perfectly linear, allowing the CPU to use its full pipeline width.

**Why doesn’t the compiler optimize the ‘if’ case automatically?**

The compiler doesn’t know the nature of your data. If only a tiny fraction of numbers were less than 500, the branch version might actually be faster because the predictor would almost always be right.

The branchless version performs a write operation (`str`) every single time. The compiler cannot safely assume that it is permitted to write to memory if a write operation in the original source code was intended to occur only under a specific condition - for example, to avoid writing past the end of an allocated buffer.

**On x86_64**

Intel Xeon (E5-2680)

```plain text
Time if: 0.594
Time bl: 0.110
```

With the _if_ version, the compiler generates a conditional branch `jg`.

```plain text
mov    edx, DWORD PTR [rax]     ; Load numbers[i]
cmp    edx, 0x1f3               ; Compare value with 499 (0x1f3)
jg     1240                     ; Branch if greater
movsxd rdi, ecx                 ; 'if' body
mov    DWORD PTR [r8+rdi*4],edx ; Store value in small_numbers[j]
add    ecx, 0x1                 ; j += 1
```

With the Branchless version the compiler uses the `setle` instruction.

```plain text
mov    ecx, DWORD PTR [rax]     ; Load numbers[i]
movsxd rsi, edx                 ; Prepare current index j
mov    DWORD PTR [rdi+rsi*4],ecx; Unconditional store in small_numbers[j]
cmp    ecx, 0x1f3               ; Compare value with 499
setle  cl                       ; cl = 1 if condition met, else 0 (No branch)
movzx  ecx, cl                  ; Zero-extend cl to ecx
add    edx, ecx                 ; j += (0 or 1)
```

**The Branch Predictor’s Memory**

Reducing the size to 10,000 elements or fewer narrows the performance gap because the branch predictor’s history tables are large enough to capture the pattern. By repeatedly processing the same array, the CPU learns the outcome of every ‘if’ condition, achieving near-perfect prediction accuracy. Increasing the size to 100,000 overwhelms this internal memory.

**Quicksort**

Quicksort is very well suited for branchless programming, since the algorithm’s main task is to partition the data by moving all elements smaller than a pivot to one side.

[A Fast Quicksort with Branch‑Avoidant Programming](https://easylang.online/blog/qsort)

christof.kaser@gmail.com
