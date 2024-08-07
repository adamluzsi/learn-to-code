# [Environment Variables](https://en.wikipedia.org/wiki/Environment_variable)

In Shell, you can have special names that express values under the hood.
These env variables are set to a specific shell runtime where they have been created.

If you don't know what are variables, [please read this](../../../basics/Variables.md).
More about that later, but first of all, what problem variables can solve ?

## Why and how we can use them ?

When we create a software application, that knows, that under ${HOME} environment variable,
you will find the user home directory without the need to know the user, or the
s/he's system settings, to be able to work in the user's home directory.
Otherwise because every one of your user probably have a slightly different value in that variable,
you less likely want to know the concrete path during when you create your software.
Otherwise, you have to make sure you know every persons home folder path in this globe,
now and in from the future as well during when you create your software.

## how they are usually represented ?

In most shell they are represented as words prefixed with dollar sign.
> $AN_EXAMPLE_ENVIRONMENT_VARIABLE

They are often enclosed with curly brackets as well.
> ${AN_EXAMPLE_ENVIRONMENT_VARIABLE}

    I personally love to use curly brackets to explicitly
    tell what is the variable name I'm interested about.
    This is especially important when those environment
    variables are used in a way, where might be part of
    something and it's not easy to tell where the
    variable name ends.

Most of the time, they receive value trough assigment:
> AN_EXAMPLE_ENVIRONMENT_VARIABLE='hello, world!'

When we assign value to a variable in shell,
we don't use the dollar sign prefix.
That is only used for retrieve the value.
This is because when you assign value to a variable,
you cannot do anything else in that shell command.
But using the env variables can happen in many place,
and trough many form, so in order to make the shell
recognize that we need a dynamic value here,
we prefix the word that mean to be interpreted
as variable with a dollar sign.

## The Scope/Visibility of such ENV Variable

When you create a shell variable in a shell,
that variable will be accessible by that shell session.

If the variable is marked to be available for child processes with `export`,
than any new process and shell session that is created from our current shell session will see the marked variables

But even by using `export` marking on the variable,
the variable will not become available in processes / shell sessions which created out current shell session.
Because of this, if you create a variable, and export it, no other shell sessions that are independent from your current session will see it.

    In a made up but still real world example,
    think about a scenario, where you have created a homework assignment.
    And you hand it over to your friend who by chance forgot to do it (shell spawning a child process or new shell session)
    Your friend is in possession of your work, and do some modification on it, before handing in for the teacher.
    Your friend had the knowledge about your homework assignment content,
    while you have no knowledge what that friend changed in it, afterwards.

### Example

```
.
├── Shell1 created
│   ├── Variable1 created and readable from Shell1 and will be readable even from any shell created in Shell1
│   └── Shell2 created by Shell1
│       ├── Variable1 readable here
│       ├── Variable1 cannot be changed in Shell1 from here but you can assign new value here that will remain for Shell2
│       └── Variable2 created and readable but no one else see it. Shell1 for example have no knowledge about it.
└── Shell3 created
    ├── Variable3 created and readable but no one else can see it, like Shell1 and Shell2 have no knowledge about it.
    └── Variable3 created and readable but no one else can see it, like Shell1 and Shell2 have no knowledge about it.
```