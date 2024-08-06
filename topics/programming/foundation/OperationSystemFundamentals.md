# Operation System Fundamentals

By Wiki Definition:
An operating system (OS) is system software that manages computer hardware and software resources and provides common services for computer programs.
Time-sharing operating systems schedule tasks for efficient use of the system and may also include accounting software for cost allocation of processor time, mass storage, printing, and other resources.
For hardware functions such as input and output and memory allocation, the operating system acts as an intermediary between programs and the computer hardware, although the application code is usually executed directly by the hardware and frequently makes system calls to an OS function or is interrupted by it.
Operating systems are found on many devices that contain a computer – from cellular phones and video game consoles to web servers and supercomputers.

The dominant desktop operating system is:
* Microsoft Windows
* MacOS by Apple Inc
* GNU/Linux variations

In other words, it is a system software created with the responsibility to abstract away hardware of a system from the user softwares.
If you try out a new operation system, you probably encounter some challenges along the learning progress.
One common notice is the realization that because user software applications are not depending on hardware, but on the operation system (kernel),
an application that able to communicate with the windows kernel, will probably have problems talking with the XNU Kernel (Mac OS).

But what is this kernel I talking about ?

As we learned before, system software main purpose is to provide platform for user software applications to run,
without the need to know about the actually hardware.
The core, the foundation of every Operation System is this core component called [kernel](https://en.wikipedia.org/wiki/Kernel_(operating_system)).
The kernel make it possible for user software applications to use the system resources like [CPU](https://en.wikipedia.org/wiki/Central_processing_unit), [RAM](https://en.wikipedia.org/wiki/Random-access_memory) and other hardware devices on the system,
in a way, that is a stable and reliable even if the actually hardware components changed under the hood.
An oversimplified example is that your future user software application will ask the kernel to save the file down to the [disk](https://en.wikipedia.org/wiki/Disk_storage),
and you don't have to know that it is a [HDD](https://en.wikipedia.org/wiki/Hard_disk_drive) or an [SSD](https://en.wikipedia.org/wiki/Solid-state_drive) or on a [floppy](https://en.wikipedia.org/wiki/Floppy_disk).
All you care is that the data is saved and can be retrieved later on. And kernel make this possible for you.

So this kernel has its own "language", that it can speak, which is called system calls.
What one kernel can understand, it is a totally unknown for the other type of kernel.
And this is the reason why the Windows NT kernel system calls is unknown by Linux Kernel.
So when your application try to talk with one type of system calls to a different type of kernel, it will not be able to function.
And on really high level, this result in that your windows software not running on Linux environment and vice versa.
Of course, there are projects that aims to provide compatibility layer, which is kind a like a translator for the given Kernel.
some translations succeed, while some could fail. It really depends on the quality of the translator.
One example is the [WINE](https://www.winehq.org/) project on Linux that aims to allow user software applications which talks windows NT system calls,
to be able to make themself understood with the Linux Kernel.

To go back to our main topic, on top of this kernel, an operation system usually have more user software application, that aims to provide all kind of different solutions for common problems.
One example to this is the [windows manager](https://wiki.archlinux.org/index.php/window_manager) which controls the placement and appearance of windows within a windowing system in a graphical user interface.
Other example could be the [Windows File Explorer](https://en.wikipedia.org/wiki/File_Explorer), which allows you to do file operations in an easy to understood graphical user interface.

You as a Software Engineer will most likely meet with different Operation Systems.
Most language abstract away system calls from you, so you rarely will feel the urge to learn more about kernels.
This was a really high level introduction to the Operation Systems, which should be enough for your future lessons.
