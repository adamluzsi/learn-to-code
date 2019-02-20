# Software 'n' Hardware

So we going to be Software Engineers. But essentially, what is software in the first place ?
And what exactly "soft" mean in this case ? Let's begin!

Back in the past, computers were made for specific tasks in mind.
The early computers each specialized to a specific task usually.
To achieve that task, they physical parts were made in a way to allow them to do so.
In other words most computer in the past were created trough [hardware](https://en.wikipedia.org/wiki/Computer_hardware) solutions.

Solving problems with hardware can often result in (by the current time's technology standards) top-notch performing solutions.
An example for this with today's technology standards is [active noise cancellation](https://en.wikipedia.org/wiki/Active_noise_control).
We can build processing unit that can process audio signals within the necessary time,
and then send instructions to a speaker to generate sound that will cancel the original sound waves.
Achieving this would be rather difficult today, without a hardware solution that aims to solve exactly this.
And the reason for that is that a software solution would be simply slower,
and for places where time and distance matter a lot, like a headphone, it is not acceptable.
This means in practice, that hardware engineers create an actual physical chip that main purpose is to calculate inverse audio wave signals.
But this hardware solution is specialized to solve just that.

For most high level business needs, hardware solutions are [bottleneck](https://en.wikipedia.org/wiki/Bottleneck_(engineering)), and it is much harder to keep evolve it on a system that being used actively.
For example if business goal is to create an accounting system that is specific to they company rules,
a hardware solution would be hard to keep up to date and evolve with the company since on each major business logic change which can happen often,
a new hardware would be necessary. Of course this is oversimplification, and only want to serve as an example.

The name itself suggest that **hard**ware is hard to change.
And this problem inspirited the need for **soft**ware which should be easy to do so compared to hardware.
This pushed things into a direction, where hardware manufacturers started to aim creating hardware that can perform simple tasks in an acceptable efficiently rate,
which then can be used by softwares to combine those simple task into a complex solution, which solve a specific [problem domain](https://en.wikipedia.org/wiki/Problem_domain).
With software, it became possible to solve multiple domain's problem on the same hardware,
in an (in case of well disciplined [architecture](https://en.wikipedia.org/wiki/Software_architecture)) efficient speed as well.

Then domain specific software application solutions and hardware began to separate even more by introducing system softwares,
which goal was to provide platform for the user software applications, so user software applications can focus mainly on the needs instead of the given hardware.
You will read more about this in the Operation System section later on.

To wrap this up, software is a collection of data or computer instructions that tell the computer how to work.
User software's goal is to achieve a task like browsing the internet, or solve a specific domain's problem.
System software's goal is to provide a stable platform for the user softwares, and make user software's independent from the hardware.

You as a software engineer will probably create user applications for domain problems,
and system software for infrastructure and platform needs.
