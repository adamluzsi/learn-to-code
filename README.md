[high-level-language]: https://en.wikipedia.org/wiki/High-level_programming_language "high-level programming language"

# Learn to Code

In this repository you will see some learning materials about programming, that probably help you begin your journey.

## Prologue

Change and learning go hand in hand. Understanding the mechanics of the brain and adopting effective learning habits
significantly boosted my performance. As a child, I struggled with studying, often feeling proficient in one area while
lacking in others. The turning point came when I realized the importance of context in learning. Whenever teachers
provided contextual information about a subject, it became easier to grasp and remember. In contrast, topics taught
without context felt harder to retain.

I've found that Software Engineering can be enjoyable and manageable when tackled incrementally. Starting with the core
concepts and building upon them makes the learning journey steady and less daunting. Without this step-by-step approach,
the field might seem overwhelming or exclusive to the "naturally gifted." However, if you find yourself struggling, it's
likely you've missed some fundamental concepts along the way.

In my view, persistence outweighs talent in this profession. Your interest in reading this affirms your determination.
If you need assistance, feel free to reach out; I'm here to help.

Each Section below here will be a reference to a separate file that focus on one topic

## Foundation

* [Cultivating Habits for a Growth-Oriented Mindset](topics/foundation/ThinkingHabits.md)
    - Embrace vulnerability; seek clarity by asking questions and accelerate your learning speed
    - Establish a solid knowledge base before progressing to more advanced concepts
    - Overcome your fears
        - Fear of the unknown is a natural response, but courage is honed when you challenge these fears.
        - Venture outside your comfort zone, even when failure seems probable.

* [Software? Hardware?](topics/foundation/SoftwareAndHardware.md)
    * [Hardware summary](topics/foundation/hardware/Summary.md)
* [Operation System Fundamentals](topics/foundation/OperationSystemFundamentals.md)
    * [GNU/Linux](topics/foundation/GNU-Linux.md)
* [What are the variables, and why they can be useful for us ?](topics/basics/Variables.md)
* [Shell / Command Line Interface](topics/foundation/shell/README.md)
    * [Graphical User Interface](topics/foundation/GraphicalUserInterface.md)
    * [Command Line Interface](topics/foundation/CommandLineInterface.md)
        * [Environment Variables](topics/foundation/shell/cli/EnvironmentVariables.md)
        * [How file/directory path(s) is interpreted](topics/foundation/shell/cli/Paths.md)
        * [How a command line request is interpreted by the shell](topics/foundation/shell/cli/CommandLineInterpretation.md)
        * [Unix/Unix-like system basic tools](topics/foundation/shell/cli/coreutils.md)

## Version control basics

- What is a version control system, and how does it help software developers?
- Choosing your first `version-control-system` (VCS)
    - the suggested `vcs` to start with is `git`
      [You can find Git basics here](/app/git/README.md)
- These companies give you remote repository hosting solutions using git.
    - [GitHub](https://github.com/)
    - [GitLab](https://gitlab.com/)
    - [Bitbucket](https://bitbucket.org/)

## Foundations of Programming

- Selecting Your First Programming Language
    - Opt for a general-purpose programming language with a concise specification.
      For instance, Go is a recommended choice due to its intentional simplicity,
      which masks complexity without concealing implementation details,
      thanks to its imperative language nature.

- Setting Up Your Development Environment
    - Install the chosen programming language on your machine.
    - Install a code editor:
        - [VSCode](https://code.visualstudio.com/)
        - [IntelliJ IDEA](https://www.jetbrains.com/idea/)

- Understanding Control Statements
    - Conditional Statements
        - if/else
        - switch/case
    - Loops and Iteration
        - for / loop / foreach
        - while / until
        - break / continue

- Introduction to the 'goto' Statement
    - It's not essential to use, but it's useful to understand the 'goto' statement for when you come across it.

- Delving into Package Management
    - Understanding Packages
    - The Rationale Behind Using Third-Party Packages
    - Discovering Packages
        - The `awesome` collections are usually a great starting point to explore new packages.
        - Community-driven package distribution platforms.

- Grasping Visibility
    - Levels of Visibility
        - Private/Protected (Unexported)
        - Public (Exported)

- Structuring
    - Classes/Structures
        - Methods
    - Singleton Functions

## Design

- Test Driven Development
    - What is TDD

- Design principle frameworks:
    - Clean Architecture
    - Domain Driven Design
    - Hexagonal Architecture

## Philosophy

* [The Unix Philosophy](topics/foundation/UnixPhilosophy.md)
* Understanding first the reason why a given programming language was invented helps to understand its design. 

