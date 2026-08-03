# Chess Analysis

Analysing chess games from the [Lichess dataset](https://database.lichess.org/)

## Example outputs (using only the 8.3GB of games from 2013-2015)

### Win rate

```
Colour   Wins      Percent
White    22528070  50.02%
Black    20917851  46.45%
Draw     1591844   3.53%
Invalid  69        0.00%
Total: 45037834
```

### Game types

```
Key                         Value     Percent
Rated Blitz game            15385204  34.16%
Rated Classical game        11306052  25.10%
Rated Bullet game           10928701  24.27%
Rated Bullet tournament     3477032   7.72%
Rated Blitz tournament      3021226   6.71%
Rated Standard game         464972    1.03%
Rated Classical tournament  287658    0.64%
Rated Correspondence game   155308    0.34%
Rated Standard tournament   11681     0.03%
Total: 45037834
```

### Titles

```
Titles:
Key       Value     Percent
Untitled  89660540  99.54%
LM        154272    0.17%
NM        85319     0.09%
FM        63418     0.07%
IM        56434     0.06%
CM        40381     0.04%
GM        15035     0.02%
WFM       261       0.00%
WCM       8         0.00%
Total: 90075668
```

### Opening Families

```
Key                   Value   Percent
Sicilian Defense      341351  10.09%
French Defense        231010  6.83%
Queen's Pawn Game     169924  5.02%
King's Pawn Game      153550  4.54%
Scandinavian Defense  146299  4.33%
Italian Game          107998  3.19%
Caro-Kann Defense     99266   2.94%
English Opening       94809   2.80%
Bishop's Opening      91681   2.71%
Van't Kruijs Opening  86108   2.55%
Total: 3381744
```

### Openings with variations

```
Key                                            Value  Percent
Scandinavian Defense: Mieses-Kotroc Variation  64451  3.29%
French Defense: Knight Variation               52681  2.69%
Sicilian Defense: Bowdler Attack               50744  2.59%
Queen's Pawn Game: Chigorin Variation          37438  1.91%
King's Pawn Game: Wayward Queen Attack         33581  1.71%
King's Pawn Game: Leonardis Variation          30059  1.53%
French Defense: Normal Variation               27894  1.42%
Sicilian Defense: Old Sicilian                 27559  1.41%
Queen's Pawn Game: Zukertort Variation         26799  1.37%
Sicilian Defense: McDonnell Attack             22701  1.16%
Total: 1958472
```

### Time Controls

```
Key    Value   Percent
60+0   423745  12.53%
300+0  336788  9.96%
180+0  246507  7.29%
120+0  236293  6.99%
300+8  183002  5.41%
600+0  113885  3.37%
240+0  109456  3.24%
300+5  87161   2.58%
0+1    86368   2.55%
300+2  65745   1.94%
Total: 3381744
```

### Game termination

```
Key               Value     Percent
Normal            29714108  65.98%
Time forfeit      15078805  33.48%
Abandoned         242429    0.54%
Rules infraction  2427      0.01%
Unterminated      65        0.00%
Total: 45037834
```

