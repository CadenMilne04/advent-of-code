# Star 1:

## Calculating area:
1. Try every square 
    1. Try all neighbors
    2. If neighbor is the same as the current string letter, then add 1 to the area

## Calculating perimeter:
1. Try every square.

AAAA
XAXX

# Star 2:

measureing sides...         

how could it be a corner? well if dir and adjacent are both off limits
1. up == oob right == oob
2. right == oob down == oob
3. down == oob left == oob
4. left == oob up == oob


aaa
axa <- diff than a or oob, and surroundings are same as a
aaa

a has 8 sides


xx <- diff than a or oob, and surroundings are diff than a or oob
ax

a has 4 sides.

if a corner is a different type or out of bounds.
