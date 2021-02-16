## Work Sample for Product Role, Golang Variant

[What is this?](https://github.com/EQWorks/work-samples#what-is-this)

## Problems

You should receive problems that can be solved using this repo as a starting point when you [apply for the job](https://apply.workable.com/eqworks/).

## Notes on working through the problems

Try to leverage Golang's `channels` and/or `sync`.

## TODO
* [ ] Support counters by content selection and time, example counter Key `"sports:2020-01-08 22:01", Value {views: 100, clicks: 4}.`

* [ ] Implement a mock store for storing counters. It can be in-memory, filesystem-based, or satellite-based (satellite not provided).

* [ ] Create go routine to upload counters to the mock store every 5 seconds.
Global rate limit for stats handler.