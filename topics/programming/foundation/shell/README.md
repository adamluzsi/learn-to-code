# Shell

In computing, a shell serves as the user interface for accessing an operating system's services.
Operating system shells typically use either a command-line interface (CLI) or a graphical user interface (GUI),
depending on the computer's role and specific tasks.
The term "shell" comes from its function as the outer layer that interacts with the operating system kernel.
CLI shells necessitate familiarity with specific commands and their syntax, 
as well as an understanding of the shell's own scripting language, like bash script for example.
These shells are also more compatible with refreshable braille displays and offer some benefits for screen readers.

> TODO:
- [ ] Choose your login Shell (Bash/Zsh/Dash)
    - `bash`
        - `bash` is the default login shell for most unix based systems
        - experience you gain with `bash` is reusable when you will work on servers

## Useful shell commands to learn

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
    - example: `export MY_ENV_VAR_NAME="the value"`
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
