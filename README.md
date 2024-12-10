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
- I got reminded that when I have problem I have to debug it on smaller, simpler input and CAREFULLY and SLOWLY go through the steps, ONE BY ONE. With patience and focus.

## Day 5

This one took me quite long. I think I was just messing up on small things but then I go nervous and demotivated and it hust became worse. Finally made it (mainly the 2dn part, the first one was pretty easy). The solution is most likely quite shit. Appearently this was classic dependency (topological graph) problem which I did bot realize. I could redo it but I am intentionally leaving all of the code here messy.

## Day 6

Part 1 was manageable. It was loads of fun and I know what I want to do from the start. Programming it took me longer that I expected and the initial version was so ugly and cumbersome that I had to refactor it so that it was easier to debug... However, I did the refactor with the AI and it did really nice job. It created the structs and objects, which was something new to me and I think I will use it quite a lot from now on. The logic was exactly the same though.

The second part was a lot trickier for me. I did couple of attempts to find possible loops with the idea that they should comply to certain geometrical criteria (top right is one row below top left...) however I could not program that. At least not that fast. Then I was starting to be desperate because I needed to get this done so I checked the subreddit and found out how possible solution can work. When I saw that I thought like a complete idiot - because this is so simple. Put wall to all previously visited positions and check for loops. Again it sounded so simple, but when I tried to code it, I was strugling. I could not get it done, even when I would swear that the logic was right, there were some errors that made it not work. I really wanted to have it done today so that I can move on so I used cursor to help me fix the mistakes

### What I have learned?

- Structs, Maps
- Shallow vs Deep copy
- Colouring and replacing terminal output (I think)

## Day 7

I didn't really know how to approach the problem, other than brute force so I checked out how this could be solved. AI recomended backtrack, which was basically brute force and some other options like dynamic programming. I had no idea what it even is and the wikipedia page looked too complicated. I also checked out binary expression trees, but that didn't quite work out and I didn't want to spend the whole day with it. I opted for the backtracking option and it's surprisingly fast. Although I have found out that the tree solution was probably not working because I had made mistakes converting the strings into numbers :(.

### What I have learned?

- Be careful with the conversions string -> int and vice versa
- Backtracking, Binary trees (kinda, just a little intro, because I didn't go for it further)
