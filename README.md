# Prisoner-Problem Visualization

## Premise

In a particular prison, there are a 100 prisoners labelled 1-100 on their uniforms. To avoid execution, they must go into a room containing 100 boxes (1-100) and find their corresponding number on their uniforms in one of the boxes. The catch is that the boxes <b>need not contain the same number labelled on the outside of the box</b> and they are allowed to check a <b>maximum of 50 boxes</b>

For example <b>Box 1</b> can contain the number <b>73</b>

They cannot communicate in any way with each other once the first prisoner heads inside the room. However the prisoners are allowed to strategize before going in.

What is the best strategy that optimizes their chances of surviving?

Veritasium Explanation of the solution - https://www.youtube.com/watch?v=iSNsgj1OCLA

## What is the scope of the project?

1. To simulate the prisoner problem experiment
2. To verify the statistics in a number of runs (1000 as a sample size)
3. To gather some useful data
4. To compare the real world results with the mathematical solution

## Results

Out of 10 iterations, The prisoner won 3 times and lost 7 times
Out of 100 iterations, The prisoner won 30 times and lost 70 times
Out of 1000 iterations, The prisoner won 297 times and lost 703 times
Out of 10000 iterations, The prisoner won 3111 times and lost 6889 times
Out of 100000 iterations, The prisoner won 31145 times and lost 68855 times

The results are very close to the mathematical solution of 31.1%
