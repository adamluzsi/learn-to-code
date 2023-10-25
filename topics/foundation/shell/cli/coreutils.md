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
