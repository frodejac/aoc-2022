# Advent of Code 2022

My go (get it?) at [Advent of Code 2022](https://adventofcode.com/2022).

## How to run

Makefile should do all the heavy lifting. If you want to build and run everything, just do `make all`.

Run a specific puzzle solution:
```bash
make day([0-1][0-9])|(2[0-4])
```

## Use your own input
AoC input depends on the logged in user. My input can be found in the [`assets`](assets) directory. It's organized by one subdirectory for each day, each with a `input.txt` input file, and `part1.txt` and `part2.txt` solution files (if I've managed to solve the puzzle). Just replace the appropriate input files if you want to run it with your own input (and also replace the solution files if you want to run the tests to verify).