# 01-clever-random

Compares the `clever` strategy against `vazirani` on

+ `random` graphs
+ `uniform` weights

And works out some kings in diagramming.

## Interpretation

`clever` seems better for low-$p$, high-$n$ graphs --- lots of vertices, but pretty sparse.

I might need to normalize for n here (the total possible weight), or express `clever`'s performance as a *multiple* of `vazirani` instead of a constant delta. There's more "weight" on offer for large graphs than for small ones.

Actually, normalizing suggests a different behavior: they're really very comparable, but `clever` is better for low-$n$, low-$p$. Neat; we've taken out the proportional effect of high $n$.