# Command Line Interpretation

Your CLI/TerminalEmulator has a strict way of interpreting your instructions.
Let's see this trough an example.

> $ /path/to/executable/command [argument 1] [argument 2] [argument 3]

Here the dollar sign is just symbolic, and I do representation of the terminal input line.
I only use it in the examples, because it is an often used symbol default representation in most environments.

## Command

> $ **/path/to/executable/command** [argument 1] [argument 2] [argument 3]

The first value is a path to a file that can be executed.
In other word, is a software that can be executed on your system.

One commonly used software is the `pwd` command.
it is located under the following path: `/usr/bin/pwd`
if you type this path to your terminal,
then the software will be executed.
The `pwd` software will tell you that what is the directory,
where your shell session currently in.

> $ /usr/bin/pwd

    /home/you

As we learned before, there are some special environment variables,
that can change some behavior for certain applications.
The shell/terminal it self also aware for some.
One such env variable is the `PATH`  variable.
The `PATH` variable holds directory paths,
where we hold executable files.
The directory paths are seperated by ":"
on unix/unix-like environments.
Why this is useful for us ?
If we type full paths all the time for our commands,
it can be rather repetitive and inefficient.
And here comes in picture why `PATH` become useful for us.
If we want to execute a software that located in one of the directory,
that the `PATH` environment variable holds,
then we don't have to type the full path, only the command name.
In our case with the above mentioned example,
we can simplify our command to this:

> $ pwd

    /home/you

This is possible, because `/usr/bin/` is mentioned in the `PATH` env variable.
We can even add our own paths as well to the `PATH` env variable,
if we use often an executable from directories that are not mentioned yet in `PATH`.

#### What happen if multiple executable located in multiple directory  ?

The `PATH` will always look for the executable file by going from left to right in the list.
As soon it find a matching executable, it will stop looking for more, and execute it.

To make it more easier to imagine, let's assume we have 3 folder,
which each holds an executable file called `x`.
```
.
├── a
│   └── x
├── b
│   └── x
└── c
    └── x

```
Also let's assume that our `PATH` current value is the following: "a:b:c"
Then when we execute the `x` command in the terminal,
the `a/x` executable will be executed, because that was the first in our list,
that is separated by `:`.

Why do I mention this ?
If you happen to start modify the `PATH` environment variable,
with your own, don't forget that if by mistake you create an application
that is already exist in the default `PATH`,
and you by chance put your directory path in the front of the `PATH` env var value,
then your directory will be the first to be located.
In this case if you by chance create an executable like `ls`,
all the applications that previously used the GNU `ls` will now start use your `ls` command.
and if the two not works in the same way, you may see some strange behavior in your shell session.

## Command Arguments

> $ /path/to/executable/command **[argument 1] [argument 2] [argument 3]**

The values you pass after the command are called `arguments`.
Often commands expect some `argument` to be able to fullfil they purpose.

    FunFact:
    The term was adopted by computer scientists when they applied mathematical reasoning to programming in the mid 20th century.
    The word argument has the general sense of something from which another thing may be deduced.
    It comes from the Latin. arguere word which means: make clear, make known, prove, declare, demonstrate.
    Its use in English to mean a “mathematical quantity from which another … quantity may be deduced, or on which its calculation depends” is attested as early as 1386:

### Why they are needed ?

When we talked about [variables](../Variables.md) in generals,
it was mentioned, that when we create software,
we work with values in  a way that they are become concrete only during runtime.
A way to request concrete values from a user is to use the CLI `arguments`

For example, if our software can only write out to the console the value it receive as an argument:
> $ echo my-free-text-argument

#### Application specific interpretation as special keywords

Some application use custom ways to interpret your arguments.
One really wide spread custom interpretation technic is to handle `arguments` separately
if they prefixed with a minus symbol (`-`).
They are often mentioned as `options`.
In practice `options` are totally application dependent thing,
how they interpret your arguments,
and has nothing to do with how shell interpret your command,
but it is in widespread use for most cli software application,
that it is wort to mention here as well.

For example:
> $ echo **-n** hello world

Most `option` interpreter solutions for applications use two type of option evaluation.
One is when you write each option as a separate argument prefixed with the minus symbol (`-`).
And the other is when you can use multiple option in one `argument` which is also prefixed by a minus symbol.
x and y here represent example option characters.
* -x -y
* -xy or -yx

## Order of interpretation

By default, the shell will first try to interpret and resolve into values each argument including the command.
Then it try to resolve the path to the command, and if there is anything that needs to be done,
to have a concrete value for the arguments, the shell will execute that first.
and then last the command will be called with the concrete arguments.

More about that later, when we do some programming. :)
