# vertex-cover

Testing vertex cover algorithms discussed in [Graphs at Work](http://lukasschwab.me/blog/gen/graphs-at-work.html).

## Runtime

I originally reached for Go because of the `testing` benchmark tooling, which I haven't ended up using. I'm not prepared to actually optimize runtime in Go. Where would one start?

+ Currently there's lots of copying. Make the graph objects publicly-mutable (weaken the abstraction). In particular, removing a vertex from a graph should be fast.

+ Consider converting deeply-recursive searches to iteration.

So maybe I would've been happier in Python.

## Variables

+ Graph topology
    + Pseudorandom
        + n vertices: would be more interesting looking at runtime. In general, prefer large n here (less artifacting).
        + p probability of a possible edge in g existing.
    + Antagonistic: see the family described in Lavrov, implemented in `NewTricky`. Have $a*Hk$ nodes for $a$, $k<a$, where $H_k$ is the $k$th harmonic number.
        + Because these vertices are always selected (by degree) and they have arbitrary weight, there's no bound on the performance vs optimum.
        + Also I guess you're always taking $(\floor Hk\floor - 1)a$ vertices.
+ Vertex weights
    + Uniform
    + Random
    + Degree-correlated (positive, negative). Negative correlation makes for delicious greedy cases.
+ Strategy
    + Vazirani
    + Clever
    + Lavrov: like Clever, but picks the edge with highest incident degree and takes *both* vertices. Outperforms clever on tricky graphs with high k.
    + Exhaustive: basically useless, even for moderate n. 2^n grows fast.

## Project structure

Can put experiment scaffolding in some pkg, then different experiments in different cmd. Each README can describe the experiment.

## To do

+ Docstrings and style.