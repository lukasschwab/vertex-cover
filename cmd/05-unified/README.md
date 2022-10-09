# 05-unified

Attempt at producing a lot of comparisons all at once, using learnings from the previous experiments.

## Interpretation

The `clever-degreePositive-tricky` graph was *very* unintuitive to me at first, but I think I've figured it out!

I expected this to be the *most* punishing graph for the clever algorithm... but the clever algorithm performs almost optimally, for all $n, k$! What gives?

First important realization: the theoretical lightest vertex cover with `DegreePositive` weights has weight $\|E\|$ by definition of the vertex cover! So all a "good" strategy can do in this context is *minimize redundancy.*

Vazirani is inherently redundant.

If I introduce a superlinear degree-correlated weight scheme, I can't imagine clever will perform well.

This is a great counterintuitive explanation.