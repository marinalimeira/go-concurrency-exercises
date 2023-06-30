# Notes

## 29th June, 2023

concurrency () vs parallelism (things at the same tiem)
how computers with 1 cpu worked with multiple processes?
unix -- > 100 things (using files/signals to communicate between programms)
time slicing is the simplest. each one has x time to process.
fnc 1 -> infinite for that prints hello.
fnc 2-> infinit for that prints goodbye.

fnc 1 -> infinite for that prints hello.
fnc 2-> infinit for that prints goodbye.
each func has a time slice and swap to the other func.
run 1 for 1 sec. stop. run 2 for 1 sec.
saving the state before stopping and resuming it.
the fnc needs to know it was interrupted. e.g. it prints "hel"
so far as the fnc 1 knows, it has the _whole_ machine to itself.
"am I awake now"? -> always yes.
what it fnc 2 sleeps for 24h before it prints? then by time slice it wastes time
"go" - let this task run until it does something.
gos scheduler. has a number of tasks and a number of resources.
-> run as many tasks as quickly as possible. keep the core as busy as possible.
no tasks waits too long. no cpu is idle.
main fnc. garbage collector (its own go routine - it's always running).
by having a "main" - you have one routine.
"go" runs a function.
scheduler: do we  have any free resource? "wake up" when a resource is free.
e.g. go executes a blocking call (time.sleep, i/o) -> scheduler "parks" it. whats the next task?
---> its ready -> goes to ready queue.
---> three states: running, ready, parked.
"go" statement -> "ready queue". It does not start it. "please run me when its convenient".
