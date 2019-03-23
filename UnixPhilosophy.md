# [The Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy)

My personal experience and opinion is that these philosophy directives actually so generally applicable,
that they are technology and age independent.
You will see later during our studies, that they are reinvented by different people, under different names, in different technologies in every 2-3 years.

The biggest challenge with them usually not to learn them, or to understand it,
but to teach them to others as a habit instead of giving them a framework that force the ideology on them.
Most of the time, those frameworks suffers because the false illusion about safety they provide,
because the user of the given framework lack the knowledge about the original reasons and believe that anything that can be achieved in the given framework, must be then good.
But that is again is my opinion, handle with a grain of salt please.

## The UNIX philosophy is documented by Doug McIlroy in the Bell System Technical Journal from 1978:
* Make each program do one thing well. To do a new job, build afresh rather than complicate old programs by adding new "features".
* Expect the output of every program to become the input to another, as yet unknown, program.
  * Don't clutter output with extraneous information.
  * Avoid stringently columnar or binary input formats.
  * Don't insist on interactive input.
* Design and build software, even operating systems, to be tried early, ideally within weeks.
  * Don't hesitate to throw away the clumsy parts and rebuild them.
* Use tools in preference to unskilled help to lighten a programming task
  * even if you have to detour to build the tools and expect to throw some of them out after you've finished using them.

## It was later summarized by Peter H. Salus in A Quarter-Century of Unix (1994):
* Write programs that do one thing and do it well.
* Write programs to work together.
* Write programs to handle text streams, because that is a universal interface.

## Examples And Similarities

The [GNU coreutils](https://www.gnu.org/software/coreutils/manual/html_node/index.html) are a perfect example for how it look when applications are built with Unix philosophy.
But you can find similar practices in software development techniques
that is originated from the same root case as the unix philosophy,
therefore share a lot of directives.
One example to this is the micro service architecture where you build small software applications
where each of them usually have small definable main purpose/goal.
We will learn more about that later on, when it will be needed.

To put it into a easy to grasp example, while [Swiss Army knife](https://en.wikipedia.org/wiki/Swiss_Army_knife) is a well famed tool,
it is still more preferable to use a proper knife for most preparing dishes,
or cutting wood with a Wood saw that serve only that purpose.
