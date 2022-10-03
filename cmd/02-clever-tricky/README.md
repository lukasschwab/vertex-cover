# 01-clever-random

Compares the `clever` and `lavrov` strategies against `vazirani` on tricky graphs.

+ `tricky` graphs: $a$ and $k$ calculated with a variable $\lambda \in [0.1, 1] \rightarrow k = \lambda a$.
+ `uniform` weights

## Interpretation

As expected, `clever` falls apart for high $a$, $k$.

`lavrov` *does* solve the problem, and performs basically as well as `vazirani` across the board. Maybe this is to be expected; both are 2-approximations on unweighted graphs.

Come to think of it, because of the way $t(v)$ is updated in Vazirani, the algorithms are effectively equivalent: the vertices all have the same weight, so every iteration of Vazirani takes both vertices! The variation between the two must just be a matter of their pseudorandom selection of edges.

Cool. The question is whether this'll hold when the edge weights vary.