[CPU]: https://en.wikipedia.org/wiki/Central_processing_unit
[MultiCore]: https://en.wikipedia.org/wiki/Multi-core_processor
[SimultaneousMultithreading]: https://en.wikipedia.org/wiki/Simultaneous_multithreading

# Understanding the [CPU][CPU]: Simplified

The Central Processing Unit (CPU) is like the brain of the computer.
It does basic math, answers simple questions, and manages data going in and out, except for some special data tasks.
When the computer is working on something, it's usually the CPU doing the job.

## Speeding it Up: Clock Speed

The CPU's speed is often talked about through its clock speed.
Just like a clock ticks every second, the CPU has its own ticks,
but it ticks way faster - billions of times in a second! 
A higher number of ticks (measured in gigahertz, or GHz) means it can do more work in the same amount of time.

## Doubling Up: Multi-Core CPUs

Having a multi-core CPU is like having multiple tiny brains working together inside one CPU.
If a program is designed to only do one thing at a time, having more cores won’t make that one thing happen faster.
However, if a program can split its work among the cores,
it can get a big speed boost from having more cores to work with.

## Sharing the Work: Simultaneous Multi Threading

Simultaneous Multi Threading (SMT) lets a single core handle multiple tasks at once, sharing its resources smartly.
This is like being given two coloring tasks at once and being able to work on them together,
instead of finishing one before starting the other.
For SMT to work well, the operating system needs to support it, 
directing the right tasks to the CPU to keep everything running smoothly.

Think about the following simplified example:
If you receive two task to do at the same time,
Like you have to colorize two fairy tale character shoes to the same color,
you can easily do both tasks without the need to ask any more information from the task giver.
If you would not be able to do multi-threading, you would have to ask for the task,
ask for the color, do the coloring, and repeat from the beginning.

## Example

2 core 4 thread 2 GHz
```txt
.
├── [core1 - 2GHZ]
│   ├── [thread1]
│   ├── [thread2]
│   ├── [thread3]
│   └── [thread4]
└── [core2 - 2GHZ]
    ├── [thread1]
    ├── [thread2]
    ├── [thread3]
    └── [thread4]
```
