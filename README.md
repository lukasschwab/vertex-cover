# vertex-cover

Benchmarks for vertex vertex cover algorithms.

## Runtime

I originally reached for Go because of the `testing` benchmark tooling, which I haven't ended up using. I'm not prepared to actually optimize runtime in Go. Where would one start?

+ Currently there's lots of copying. Make the graph objects publicly-mutable (weaken the abstraction). In particular, removing a vertex from a graph should be fast.

+ Consider converting deeply-recursive searches to iteration.

So maybe I would've been happier in Python.