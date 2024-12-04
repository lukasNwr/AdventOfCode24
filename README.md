# AdventOfCode24

> This year I am trying to learn GO. Haven't tried GO before so the solutions might be not the most elegant ever :D

## Day 1

Haha, this took me much longer than I anticipated. I have forgotten a lot of the CS basics and had to research how to do the quick sort and how to do binary search (even though I didn't even use it in the end)

### What have a learned?

- I have learned how to read file in GO
- I have practiced some array/slices, conversions, maps and other basics
- I have revisited and implemented QuickSort and LinearSearch

## Day 2

Today it was pretty easy. I got stuck for a while because I didn't notice one condition for unsafe report (the equal levels). Then I was unsure about the for loop skipping in go but that was easy to figure out. The solution could most likely be cleaner, I went with pretty verbose if "early returns" but I like that it's more specific even if it's less concise.

### What have a learned?

- Revisited the file reading, string -> int conversions, abs values...

## Day 3

First part took me a while. The logic was clear to me from the start - I knew what I wanted to do. However, the implementation took me longer. I didn't think of some small quirks and the implementaion that I have could be probably optimized even with my current knowledge of GO. BUT, it's working :D. The second part was super easy and took me only a few minutes.

### What have a learned?

- I have played around with slices some more (I still don't feel super confident with them)
- File read as bytes rather than line by line

## Day 4

Today, the task was in my opinion conceptually very easy. However, the implementation was dependent on 2d array checks, which I have always hated. It's based on playing (or not playing but knowing :D) with indices and I always get annoyed by this. This time, I got it right, but it took me some time to fix other mistake. The mistake was that I didn't think of the case where diagonals could contain more than one occurence of XMAS. After I realized that, it was easy to fix.

### What have a learned?

- More work with arrays/slices (which is greate because I feel I am getting more confortable with them)
- More work with files
- I got introduced into the runes - way to work with individual characters in strings (since strings in go are immutable and you access only idx byte, not necessarily character)
- I have learned that when I have problem I have to debug on smaller, simpler input and CAREFULLY and SLOWLY go through the steps, ONE BY ONE.
