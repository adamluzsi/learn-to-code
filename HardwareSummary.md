# Hardware Summary

There are a few components in the commonly available personal computers, that I would like to mention.
This will be a high level, oversimplified summary.

For the sake of representation, I will make references about
similarities between the purpose of a hardware components and the behavior of the brain for such tasks.
Again, it will be oversimplified, and the examples will be not 100% correct from Neuroscience point of view.
While 1:1 references are inaccurate, it will be still good for representation.

## [Central processing unit (CPU)](https://en.wikipedia.org/wiki/Central_processing_unit)

CPU performs the basic arithmetic, logic, controlling, and input/output (I/O) operations specified by the instructions.
This is when the brain try to use mathematics, answer a yes or no question or doing REM during night.
Almost anything the computer **do actively** is usually done by the CPU, except I/O operations.

### Clock Cycle / Clock Speed

One of its property is the clock speed which sometimes used as a measurement unit for its performance.
CPU clock speed, or clock rate, is measured in Hertz, generally in gigahertz, or GHz.
A CPU’s clock speed rate is a measure of how many clock cycles a CPU can perform per second.
For example, a CPU with a clock rate of 1.8 GHz can perform 1,800,000,000 clock cycles per second.
If we oversimplify and ignore all other bottleneck in the system, this can be interpreted as,
the more clock cycle it can do, the more [program instruction](https://en.wikipedia.org/wiki/Instruction_set_architecture) it will execute.

### [MultiCore](https://en.wikipedia.org/wiki/Multi-core_processor)

Multi-core means what you would assume, multiple CPU core is integrated together in one CPU component.
For applications written in a way that they [only execute one action a time](https://en.wikipedia.org/wiki/Synchronous_programming_language), this means,
you can run multiple application on a computer, but even if you have more core in the CPU, it will not increase the performance of one individual applications.
An oversimplified example is that for such programs it is no matter if your CPU has 4 core or Infinite.

If the application created in a way, that it can [distribute programming instructions between CPU cores](https://en.wikipedia.org/wiki/Asynchrony_(computer_programming)),
this also can easily mean that that kind of application can benefit from applications that the computer additional CPU cores.
An oversimplified example is that for such programs it can be a great performance boost if the system has 8 core instead of 4.

### [Simultaneous Multi Threading](https://en.wikipedia.org/wiki/Simultaneous_multithreading)

In multi-threading as general, the CPU core poses to the Software as N separate cores
and fetches N programs in parallel, just like N separate cores would.
Now these N programs are fed into the core to make the most of the unused assets in the chip,
and thus make possible that multiple processes can use the same resources.
In addition, it is required for simultaneous multithreading (SMT) that the operating system support it.

### Example

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

## [Random-access memory (RAM)](https://en.wikipedia.org/wiki/Random-access_memory)

RAM is the place where the short - term memorization happen in case of the brain.
This is like when you hear something new during a day, that you cannot match with something you already know,
or it is not that easy, so you remember about it for a certain period of time.

The "random" part here means that when you need something that you memorized for a given "address",
you can receive it back directly without the need for searching it in the memory.
Basically when you memorize the restaurant WC password from the receipt, at the door,
you don't have to think trough everything what you remember until you find the code,
but only think about the "wc password from the receipt", and you got it.
I fear such scenario would be troublesome otherwise.

Using short-term memory is much faster and easier for both the brain and the computer,
than the long-term memory, but more about that later.

Modern RAMs can only keep they supplied with electricity.
So if your power off / restart that device, the RAM will lose its content.

### [Clock Speed meaning for ram (MHz)](https://en.wikipedia.org/wiki/DDR_SDRAM)

MHz of RAM is the maximum number of clock cycles per second that the RAM operates on.
With Double Data Rate (DDR) RAM, it actually communicates twice per cycle.
This is why chips are now named for their bandwidth, not their frequency alone.
It's still necessary to know the clock rate, to ensure that the motherboard/CPU can operate at that clock.

Example with a 200 MHz DDR (PC-3200):
> 200 MHz clock rate × 2 (for DDR, 1 for SDR) × 8 Bytes = 3,200 MB/s bandwidth

# [Input/Output (I/O or IO)](https://en.wikipedia.org/wiki/Input/output)

Inputs are the signals or data received by the system and outputs are the signals or data sent from it.
The term can also be used as part of an action; to "perform I/O" is to perform an input or output operation.
Oversimplifying it, this is communication between anything that is not CPU or RAM.
This includes network communication (e.g. internet),
or reading content from your storage device (place where your files live on your PC).

## [Local Storage](https://en.wikipedia.org/wiki/Local_storage)

This device sole purpose is to provide the computer with long term memory.
These devices usually much more slower than RAM.
This is usually not a big deal, because most modern [operation systems](OperationSystemFundamentals.md) try to keep things in RAM,
and do as many as reasonable I/O operation together on the local storage.

It is like when you think hard to remember something you don't often use,
when you finally remember it, your brain put a working-copy of that information into your short-term memory,
so if you need it with in a given time, you still can easily remember about it.
Also when you learn things during the day, you will most likely memorize it in your long-term memory during your sleep's [REM](https://en.wikipedia.org/wiki/Rapid_eye_movement_sleep).

## [Keyboard](https://en.wikipedia.org/wiki/Computer_keyboard) (copied from wiki)

In computing, a computer keyboard is a typewriter-style device which uses an arrangement of buttons or keys to act as mechanical levers or electronic switches.
Following the decline of punch cards and paper tape, interaction via teleprinter-style keyboards became the main input method for computers.
Keyboard keys (buttons) typically have characters engraved or printed on them,
and each press of a key typically corresponds to a single written symbol.
However, producing some symbols may require pressing and holding several keys simultaneously or in sequence.
While most keyboard keys produce letters, numbers or signs (characters),
other keys or simultaneous key presses can produce actions or execute computer commands.
In normal usage, the keyboard is used as a text entry interface for typing text and numbers into a word processor, text editor or any other program.
In a modern computer, the interpretation of key presses is generally left to the software.
A computer keyboard distinguishes each physical key from every other key and reports all key presses to the controlling software.

## [Mouse](https://en.wikipedia.org/wiki/Computer_mouse) (copied from wiki)

A computer mouse is a hand-held pointing device that detects two-dimensional motion relative to a surface.
This motion is typically translated into the motion of a pointer on a display,
which allows a smooth control of the graphical user interface.
The first public demonstration of a mouse controlling a computer system was in 1968.
Originally wired to a computer, many modern mice are cordless, relying on short-range radio communication with the connected system.
Mice originally used a ball rolling on a surface to detect motion,
but modern mice often have optical sensors that have no moving parts.
In addition to moving a cursor, computer mice have one or more buttons to allow operations such as selection of a menu item on a display.
Mice often also feature other elements, such as touch surfaces and "wheels",
which enable additional control and dimensional input.

## [Touchpad](https://en.wikipedia.org/wiki/Touchpad) (copied from wiki)

A touchpad or trackpad is a pointing device featuring a tactile sensor,
a specialized surface that can translate the motion and position of a user's fingers to a relative position on the operating system that is made output to the screen.
Touchpads are a common feature of laptop computers, and are also used as a substitute for a mouse where desk space is scarce.
Because they vary in size, they can also be found on personal digital assistants (PDAs) and some portable media players.
Wireless touchpads are also available as detached accessories.

## [Monitor](https://en.wikipedia.org/wiki/Computer_monitor) (copied from wiki)

A computer monitor is an output device that displays information in pictorial form.
A monitor usually comprises the display device, circuitry, casing, and power supply.
The display device in modern monitors is typically a thin film transistor liquid crystal display (TFT-LCD) with LED backlighting having replaced cold-cathode fluorescent lamp (CCFL) backlighting.
Older monitors used a cathode ray tube (CRT). Monitors are connected to the computer via VGA, Digital Visual Interface (DVI), HDMI, DisplayPort, Thunderbolt, low-voltage differential signaling (LVDS) or other proprietary connectors and signals.

Originally, computer monitors were used for data processing while television receivers were used for entertainment.
From the 1980s onwards, computers (and their monitors) have been used for both data processing and entertainment,
while televisions have implemented some computer functionality.
The common aspect ratio of televisions, and computer monitors, has changed from 4:3 to 16:10, to 16:9.

Modern computer monitors are easily interchangeable with conventional television sets.
However, as computer monitors do not necessarily include components such as a television tuner and speakers,
it may not be possible to use a computer monitor as a television without external components.

## [Touchscreen](https://en.wikipedia.org/wiki/Touchscreen) (copied from wiki)

A touchscreen, or touch screen, is an input device and normally layered on the top of an electronic visual display of an information processing system.
A user can give input or control the information processing system through simple or multi-touch gestures by touching the screen with a special stylus or one or more fingers.
Some touchscreens use ordinary or specially coated gloves to work while others may only work using a special stylus or pen.
The user can use the touchscreen to react to what is displayed and, if the software allows, to control how it is displayed;
for example, zooming to increase the text size.

The touchscreen enables the user to interact directly with what is displayed, rather than using a mouse, touchpad,
or other such devices (other than a stylus, which is optional for most modern touchscreens).

Touchscreens are common in devices such as Nintendo game consoles, personal computers, electronic voting machines, and point-of-sale (POS) systems.
They can also be attached to computers or, as terminals, to networks.
They play a prominent role in the design of digital appliances such as personal digital assistants (PDAs) and some e-readers.

The popularity of smartphones, tablets, and many types of information appliances is driving the demand and acceptance of common touchscreens for portable and functional electronics.
Touchscreens are found in the medical field, heavy industry, automated teller machines (ATMs),
and kiosks such as museum displays or room automation, where keyboard and mouse systems do not allow a suitably intuitive,
rapid, or accurate interaction by the user with the display's content.

Historically, the touchscreen sensor and its accompanying controller-based firmware have been made available by a wide array of
after-market system integrators, and not by display, chip, or motherboard manufacturers.
Display manufacturers and chip manufacturers have acknowledged the trend toward acceptance of touchscreens
as a user interface component and have begun to integrate touchscreens into the fundamental design of their products.

## [Network+NetworkCards](https://en.wikipedia.org/wiki/Computer_network) (copied from wiki)

A computer network is a digital telecommunications network which allows nodes to share resources.
In computer networks, computing devices exchange data with each other using connections (data links) between nodes.
These data links are established over cable media such as wires or optic cables, or wireless media such as Wi-Fi.

Network computer devices that originate, route and terminate the data are called network nodes.
Nodes are generally identified by network addresses, and can include hosts such as
personal computers, phones, and servers, as well as networking hardware such as routers and switches.
Two such devices can be said to be networked together when one device is able to exchange information with the other device,
whether or not they have a direct connection to each other. In most cases,
application-specific communications protocols are layered (i.e. carried as payload) over other more general communications protocols.
This formidable collection of information technology requires skilled network management to keep it all running reliably.

Computer networks support an enormous number of applications and services such as access to the World Wide Web,
digital video, digital audio, shared use of application and storage servers, printers, and fax machines,
and use of email and instant messaging applications as well as many others.
Computer networks differ in the transmission medium used to carry their signals,
communications protocols to organize network traffic, the network's size, topology,
traffic control mechanism and organizational intent. The best-known computer network is the Internet.
