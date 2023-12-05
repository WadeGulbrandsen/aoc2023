![GitHub License](https://img.shields.io/github/license/WadeGulbrandsen/aoc2023?logo=github)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/WadeGulbrandsen/aoc2023?logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/WadeGulbrandsen/aoc2023)](https://goreportcard.com/report/github.com/WadeGulbrandsen/aoc2023)
![GitHub top language](https://img.shields.io/github/languages/top/WadeGulbrandsen/aoc2023?logo=github)
![GitHub repo size](https://img.shields.io/github/repo-size/WadeGulbrandsen/aoc2023?logo=github)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/WadeGulbrandsen/aoc2023/go.yml?logo=github&label=tests)
[![Go Coverage](https://github.com/WadeGulbrandsen/aoc2023/wiki/coverage.svg)](https://raw.githack.com/wiki/WadeGulbrandsen/aoc2023/coverage.html)

# Advent of Code 2023

I'm trying to do the [Advent of Code](https://adventofcode.com/) in [Go](https://go.dev/) this year.

I just started learning Go on Nov 28, 2023. After doing the first 2 days I feel like I'm starting to get the hang of Go.

## Day 5
Had some performance issues with my solution for the second problem.

First I added a global `backSteps` variable to be used by the `LocationToSeed` method. Which didn't make any noticable difference. I also removed some of the print statements which also made no effect on the speed. Replacing `for i := 0; i <= 157211394; i++` with `for i := 0; true; i++` for the condition in the loop that does most of the work did nothing either.

### No read locks
Since after creation there are no writes to the almanac struct it's safe not to use locks. This shaved 35 seconds from the time.

| Version | Sample 1 | Sample 2 | Input 1 | Input 2 |
| ------- | -------- | -------- | ------- | ------- |
| [Initial](https://github.com/WadeGulbrandsen/aoc2023/commit/6d7e10fc3ce737a352be12fcc445bcb1771afc80) | 432.21µs | 318.257µs | 1.476169ms | 1m25.011957731s |
| [No read locks](https://github.com/WadeGulbrandsen/aoc2023/commit/4c8d223ea4c53623d9c27d886fe2eaf41b22685e) | 471.923µs | 507.104µs | 1.281487ms | 50.185503074s |

## Test Coverage
![Go Coverage Chart](https://github.com/WadeGulbrandsen/aoc2023/wiki/coverage-chart.svg)