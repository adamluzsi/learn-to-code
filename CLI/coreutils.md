# Unix/Unix-like system basic tools

There are a couple of basic tools that you can use on a Unix/Unix-like OS.
If you learn them, you will be at home easily on any OS variation that is at least Unix-like system.

## GNU coreutils as implementation

The GNU Core Utilities or coreutils is a package of GNU software containing reimplementations for many of the basic tools,
such as cat, ls, and rm, which are used on Unix-like operating systems.

The GNU core utilities support long options as parameters to the commands,
as well as the relaxed convention allowing options even after the regular arguments
(unless the POSIXLY_CORRECT environment variable is set).
Note that this environment variable enables a different functionality in BSD.

Alternative implementation packages are available in the FOSS ecosystem,
with a slightly different scope and focus, or license.
For example, GPLv2-licensed BusyBox and BSD-licensed Toybox are available for use in embedded devices.

## [Environment Variables](https://en.wikipedia.org/wiki/Environment_variable)

In Shell, you can have special names that express values under the hood.
In practice for example, you can have a box that holds something,
you can mention it as well that "use the content of that box",
without even have to know about what is actually in the box.
Or an another example is when you have a contact in your phonebook,
and you call this contact all the time,
but you don't remember they actually number out from your head.
Therefore when that contact change they phone number,
you still call them in the same way, and don't have to memorize a new phone number.

Basically this is the same reason why those named variables are being used.
For example you can create a software, that knows, that under "${HOME}",
your current user who run your software have a directory that belongs to that user.
And because every one of your user probably have a slightly different value in that variable,
you less likely want to know the conrete path during when you create your software.
Otherwise, you have to make sure you know every persons home folder path in this globe,
now and in from the future as well.

## Command Options and Arguments

The CLI works in the following way to interpret your requests.

> $ /path/to/executable/command [argument 1] [argument 2] [argument 3]

TODO

## [STDIN, STDOUT, STDERR](https://en.wikipedia.org/wiki/Standard_streams)

TODO

## Path and shell interpretation

	In the explanation, I made the assumption,
	that your shell is in your home directory,
	that you work on Unix/Unix-like environment,
	that your user's home directory is under /home/you.
	I intentionally don't use anything that would be able to resolve the need for those assumtions,
	because they would be noise in the examples.
	So please replace the path parts "/home/you" with the path you currently in.
	To find out where you are, scroll down for the `pwd` command.


In the CLI shell you can refer to files trough two way.
One is the absolute path, which starts with a slash,
and mention every directory until it reach the target.
> /home/you/Pictures/funny_cat.gif

The other solution is to use relative paths.
Relative paths always mean to be from your current working directory,
which is the directory you are currently with your shell/process.
This can be useful in case you don't know the full path,
or any reference that can be used to express a file location.
> Pictures/funny_cat.git

There is a special key which represent the "at where I am", and that is the dot.
You can use it that as well yo represent the current location
> ./Pictures/funny_cat.git
The dot itself represent the "at where I am" not binded to your shell, but to the path.
So in each part of the path which is separated by "/", you can say dot,
and that dot will represent the current directory til that path.
> ./Pictures/././././././ is equal to ./Pictures/. which is equal to ./Pictures which is equal to Pictures
There is one more special character, which is the ".." / double dot.
The double dot represent "one directory above from the point where I am" in the path.
> ../ is equal to /home/you


Some command line interpreter allows you wildcards,
which allows you to match multiple file in different ways.

For example, if you want to list every file in your Pictures directory:
> ./Pictures/*
Or if you want to list only the gif files:
> ./Pictures/*.gif
Or if you want list any file there which includes the name: cat
> ./Pictures/*cat*
Or if you want to list any file that name is cat, in one directory belove where you are:
> ./*/cat

In practical terms, this means that the command which receive the path in a way like this,
in reality will receive many path as separate arguments.
Like when you call the command `echo` (below more detail about echo)
with a path that includes wildcard, and it match multiple one,
the echo command will receive each path, and not the wildcard.
But only if the wildcard match actual files, else the wildcarded path will be received by the command.

	Most command line interpreter allows you to have even more advanced wildcarding,
	but for the sake of learning curve minimalization, I only wanted to share this one with you.
	This can be combined in so many way, that it would be hard to list all,
	but I hope you are on the same page as me now regarding paths.

## Commands

Each of the following commands can be executed from the CLI shell.
Most linux distribution have a system level shortcut for that with: CTRL + Shift + T

### man - an interface to the on-line reference manuals

The man command can provide you with documentation about commands that have `man` page.
This command is super helpful both in the early first encounters with the CLI interface,
and even when you are a experienced CLI wizard/witch.

You can use it simply by providing the name of the command you are interested about.
```sh
man andTheNameOfTheCommand
```
```sh
man ls
```

if you fail to find documentation for a given CLI command,
you can always try the following arguments that almost every CLI command understands.
```sh
NameOfTheCommand --help
NameOfTheCommand -h
```
```sh
ls --help
```

I personally often find myself use the --help argument
and the man command to figure out how to use a command which I don't know,
or don't remember anymore.

### ls - list directory contents

ls command allows you to list directory contents, and also to get more information
about the files located in a directory.

```sh
ls /absolute/path/to/the/directory
ls ./relative/path
```

```sh
ls Pictures
```

ls also able to tell you details about the files and directories it located under a directory path.
You can use the "-l" option.

```sh
ls -l Pictures
```

If you don't provide path to the ls command,
it will list out files and directories under your current working directory.

### echo - display a line of text

echo is a command / shell builtin that allows you to write to the standard out (STDOUT)
anything you give as an argument to the echo command.

In practice, you can use this command to provide input for other commands that expect something coming from standard in (STDIN).
