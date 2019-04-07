# [Environment Variables](https://en.wikipedia.org/wiki/Environment_variable)

In Shell, you can have special names that express values under the hood.
These env variables are binded to the shell where they have been created.

More about that later, but first of all, what problem variables can solve ?  

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

## Practical example for why variables in generally useful

In far fetched example, you can have a box that holds something,
You can ask someone to carry that box for you.
Therefore when a moving company come to pick up your belongings,
all they need to know is that you have a couple of box that needs to be
delivered to a new address, without the need to know the content.

In a different usage scenario with variables,
you can have a friend in your contact address book,
you can call that person with your phone, without the need to know the 
phone number of your friend. Therefore when that contact change they phone number,
you still call them in the same way, and don't have to memorize a new phone number.
This is especially good when you have a company phone address book,
and you search a someone by they role. Like HR department.
You don't have to know everyone to know who you need to contact,
in case you have a problem, instead you can use a symbolic group mailing address.  

In a more software engineering related way, this can be used for us to 
create a software application, that knows, that under ${HOME} environment variable,
you will find the user home directory without the need to know the user, or the 
s/he's system settings, to be able to work in the user's home directory.
Otherwise because every one of your user probably have a slightly different value in that variable,
you less likely want to know the concrete path during when you create your software.
Otherwise, you have to make sure you know every persons home folder path in this globe,
now and in from the future as well during when you create your software.

## The Scope/Visibility of such ENV Variable 

This means when one variable is created in one shell, 
is cannot be accessed from a different shell,
but any shell that is created in the shell where they have been created can see it.