# Curriculum

## Thinking Basics

- [How to build habits that lead to a growth mindset](ThinkingHabits.md)
    - [ ] Don't fear showing weakness, ask questions to learn faster
    - [ ] build the foundation of your knowledge before you start building on top of it
    - [ ] conquer your fears
        - the fear from the unknown is in our instincts, bravery is some you learn while going against your fear.
        - give it a try even if you fear a likely failure

## Know your Development Environment
- [ ] [Software? Hardware?](SoftwareAndHardware.md)
- [ ] [Hardware summary](hardware/Summary.md)
- [ ] [Operation System Fundamentals](OperationSystemFundamentals.md)
  - [ ] [GNU/Linux](GNU-Linux.md)
- [ ] [What are the variables, and why they can be useful for us?](Variables.md)
- [Shell / Command Line Interface](Shell.md)
    - [ ] [Graphical User Interface](GraphicalUserInterface.md)
    - [ ] [Command Line Interface](CommandLineInterface.md)
        - [ ] [Environment Variables](cli/EnvironmentVariables.md)
        - [ ] [How file/directory path(s) is interpreted](cli/Paths.md)
        - [ ] [How a command line request is interpreted by the shell](cli/CommandLineInterpretation.md)
        - [ ] [Unix/Unix-like system basic tools](cli/coreutils.md)

### Practical exercise with a Command Language Interpreter
- [ ] Choose your login Shell (Bash/Zsh/Dash)
    - `bash` 
      - `bash` is the default login shell for most unix based systems
      - experience you gain with `bash` is reusable when you will work on servers

- shell utilities to know
    - [ ] cd
        - allows you to change the current working directory.        
    - [ ] ls
        - list the files and directories
    - [ ] mkdir + rmdir
        - makes/removes directories
    - [ ] rm
        - removes file (optionally recursively directories as well)
    - [ ] cat
        - print the content of a file into the STDOUT of your shell session
    - [ ] grep
        - search matching lines in files or in the STDIN of the grep command 
    - [ ] ps
        - list running processes
    - [ ] export
        - export a shell variable, so a subshell can inherit it
    - [ ] echo
        - print out a text from its arguments into the shell session's STDOUT
    - [ ] . 
        - or `source`
        - source the content of a shell script into current shell session
    - [ ] vim
        - it's a text editor you can use from the shell.
        - just the basics is enough
            - press `i` for insert mode that allow you to type
            - press `esc` for cancelling insert mode
            - press `:` and type `w` then hit `enter` to write changes, aka save changes to the file (`:w`)
            - press `:` and type `q` then hit `enter` to exit vim
            - you can combine commands like `:wq` to write changes and then exit vim.
            - press `dd` to cut a line and then press `p` to paste the previously cut line to the cursors position.

## Version control basics

- [ ] What is a version control system, and how does it help software developers?
- [ ] Choosing your first `version-control-system` (VCS)
    - the suggested `vcs` to start with is `git`
      [You can find Git basics here](/app/git/README.md)
- [ ] repository hosting companies
    - [GitHub](https://github.com/)
    - [GitLab](https://gitlab.com/)
    - [Bitbucket](https://bitbucket.org/)

## Programming basics

- [ ] Choose your language first language
    - Try to use a general-purpose programming language that has a small language specification. For example, due to its intentional simplicity, I suggest that Go hide complexity while not hiding implementation details due to its imperative language nature. 
- [ ] Installing your language of choice on your machine
- [ ] install a code editor
    - [VSCode](https://code.visualstudio.com/)
    - [IntelliJ IDEA](https://www.jetbrains.com/idea/)

- control statements
    - conditional statements
        - if/else
        - switch/case
- loop / iteration / enumeration
    - [ ] for / loop / foreach
    - [ ] while / until
    - [ ] break / continue

- [ ] goto statement
    - not necessary for using, but to know about the statement when you encounter with it

- [ ] Package management
    - what is a package
    - why do we work with third-party packages
    - where can you find packages
        - the `awesome` collections usually a good starter to find new things to play with
        - what are the community driven package distribution solutions 

- visibility 
    - types of visibility
        - [ ] private,protected / unexported
        - [ ] public / exported
    - Who's visibility can be managed
        - singleton functions
        - classes/structures
        - methods

- [ ] 

## Team conventions to speed up mental model building time
- [ ] common code formatting

## Design

- [ ] Test Driven Development
    - What is TDD

- Design principle frameworks:
    - [ ] Clean Architecture
    - [ ] Domain Driven Design
    - [ ] Hexagonal Architecture

## Philosophy

- [ ] [The Unix Philosophy](UnixPhilosophy.md)
- [ ] Understanding first the reason why a given programming language was invented helps to understand its design. 

