# [Variables](https://en.wikipedia.org/wiki/Variable_(computer_science))

## What are they, and why do we need it ?

Variables allow us to work with values that we don't know during the writing the code,
but will be available when our program is running somewhere.

## Examples from real world with real world

### The Deliverer

In far fetched example, you can have a box that holds something,
You can ask someone to carry that box for you.
Therefore when a deliverer come to pick up your belongings,
all s/he need to know is that you have a couple of box that needs to be
delivered to an address, without the need to know the content.
To go further with this example, we can vaguley describe the deliverer job description
using the "package" and "location" keyword as variables in a sentence, without the need to know the actual
concrete examples for each upcoming delivery:
> In generally the deliverer deliver packages to a location


### Contact List

In a different usage scenario with variables,
you can have a friend in your contact address book,
you can call that person with your phone, without the need to know the
phone number of your friend. Therefore when that contact change they phone number,
you still call them in the same way, and don't have to memorize a new phone number.
This is especially good when you have a company phone address book,
and you search a someone by they role. Like HR department.
You don't have to know everyone to know who you need to contact,
in case you have a problem, instead you can use a symbolic group mailing address.

## Syntax / Written form

If types and other things don't tell you anything now,
don't worry, because we will have a detailed introduction for that in the future.
For now, it is enough to understand the examples above,
and just see a few example about how they could look in written form.

Most of the case, the single equal sign is a good way to spot variable declaration.

### In most static type language

In language where you have to work with types,
usually the following variations are the most common:

    dataType variableName = initialValue;
    var variableName dataType = initialValue;
    variableName := initialValue;

### In most dynamic type language

In most high level language, types usually not defined in the code.
So most of the time it is either a keyword like "var"
or even not that required to write down variable declaration.

    var variableName = initialValue;
    variableName = initialValue;
    SET variableName TO initialValue;

Just to mention it, types in most high level language are used,
but not mentioned in the code, to apply some programming principles with less boilerplate.
But more about that as well in the future!

