# 03-clever-tricky-random

Trying out random weights. I expect the same results for large graphs/high reps, since an unweighted algo on a graph with random weights will have an average vertex weight matching the distribution average.

## Interpretation

I was wrong! And I should've seen this.

Vazirani is relatively better here because its performance *isnt'* weight-agnostic. It might take strategies with more vertices but lower total weight. Nice!
