# Border arrays

We define the border array `ba` such that `ba[i]` is the longest border of `x[1:i+1]`.

Implement a function, `border_array(x) -> ba`, that computes this border array.

Now, define `bax` such that `bax[i]` is the longest border of `x[1:i+1]` where `x[bax[i]]` differs from `x[i+1]`, that is, the next character after the border and the position in `x` must be different. We call this the *strict border array*.

Can you derive an algorithm for computing `bax`? What is the running time?

Implement a function, `strict_border_array(x) -> bax` that computes the strict border array.

A hint here is that either `bax[i]=ba[i]`, when the additional condition is satisfied, or else it must be the longest border of `x[1:ba[i]]` where the condition is satisfied (can you show why?). This is a value that can be found in `bax`.
