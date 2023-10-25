# Path and shell interpretation

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
