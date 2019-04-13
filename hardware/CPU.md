# [Central processing unit (CPU)](https://en.wikipedia.org/wiki/Central_processing_unit)

CPU performs the basic arithmetic, logic, controlling, and input/output (I/O) operations specified by the instructions.
This is when the brain try to use mathematics, answer a yes or no question or doing REM during night.
Almost anything the computer **do actively** is usually done by the CPU, except I/O operations.

## Clock Cycle / Clock Speed

One of its property is the clock speed which sometimes used as a measurement unit for its performance.
CPU clock speed, or clock rate, is measured in Hertz, generally in gigahertz, or GHz.
A CPU’s clock speed rate is a measure of how many clock cycles a CPU can perform per second.
For example, a CPU with a clock rate of 1.8 GHz can perform 1,800,000,000 clock cycles per second.
If we oversimplify and ignore all other bottleneck in the system, this can be interpreted as,
the more clock cycle it can do, the more [program instruction](https://en.wikipedia.org/wiki/Instruction_set_architecture) it will execute.

## [MultiCore](https://en.wikipedia.org/wiki/Multi-core_processor)

Multi-core means what you would assume, multiple CPU core is integrated together in one CPU component.
For applications written in a way that they [only execute one action a time](https://en.wikipedia.org/wiki/Synchronous_programming_language), this means,
you can run multiple application on a computer, but even if you have more core in the CPU, it will not increase the performance of one individual applications.
An oversimplified example is that for such programs it is no matter if your CPU has 4 core or Infinite.

If the application created in a way, that it can [distribute programming instructions between CPU cores](https://en.wikipedia.org/wiki/Asynchrony_(computer_programming)),
this also can easily mean that that kind of application can benefit from applications that the computer additional CPU cores.
An oversimplified example is that for such programs it can be a great performance boost if the system has 8 core instead of 4.

## [Simultaneous Multi Threading](https://en.wikipedia.org/wiki/Simultaneous_multithreading)

Multi threading from CPU perspective means one CPU core fetches multiple program instructions,
and some shared resources can be used more efficiently.

In practice, at high level from operation system perspective,
this looks like one CPU core represented by multiple one,
thus the operation system provides task to each CPU core it sees.
Now these programs the threads received are fed into the actually CPU core
to make the most of the unused assets in the chip,
and thus make possible that multiple processes can use the same resources.
For this to work with the Operation System,
it is required for simultaneous multithreading (SMT) to be supported by that OS.

Think about the following simplified example:
If you receive two task to do at the same time,
Like you have to colorize two fairy tale character shoe to the same color,
you can easily do both task without the need to ask any more information from the task giver.
If you would not be able to do multi threading, you would have to ask for the task,
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
