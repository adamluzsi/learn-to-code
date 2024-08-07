# Getting started with Git

**Step 1: Install Git**

Download and install the latest version of Git from the official website:

[https://git-scm.com/downloads](https://git-scm.com/downloads)

Follow the installation instructions for your operating system.


**Step 2: Set up your Git account**

Open a terminal or command prompt and run the following commands to set up your Git account:
```
$ git config --global user.name "Your Name"
$ git config --global user.email "your@email.com"
```
Replace `"Your Name"` and `"your@email.com"` with your actual name and email address.


**Step 3: Create a new repository**

Create a new directory for your project and navigate to it in the terminal or command prompt:
```bash
$ mkdir myproject
$ cd myproject
```
 Initialize a new Git repository:
```
$ git init
```
This will create a new `.git` directory in your project folder.


**Step 4: Add files to the repository**

Create a new file called `hello.txt` and add some content to it:
```bash
$ echo "Hello, World!" > hello.txt
```
Add the file to the Git repository:
```
$ git add hello.txt
```
This will stage the file for the next commit.


**Step 5: Commit changes**

Commit the changes with a meaningful message:
```
$ git commit -m "Initial commit"
```
This will create a new snapshot of your project and save it to the Git history.


**Step 6: Create a branch**

Create a new branch called `feature/new-feature`:
```bash
$ git branch feature/new-feature
```
Switch to the new branch:
```bash
$ git checkout feature/new-feature
```
Make some changes to the `hello.txt` file and commit them:
```bash
$ echo "This is a new feature!" >> hello.txt
$ git add hello.txt
$ git commit -m "Added new feature"
```


**Step 7: Merge branches**

Switch back to the master branch:
```
$ git checkout master
```
Merge the changes from the `feature/new-feature` branch into the master branch:
```bash
$ git merge feature/new-feature
```
This will update the master branch with the new changes.


**Step 8: Push to a remote repository**

Create a new repository on GitHub or another Git hosting service.

Add the remote repository to your local Git configuration:
```
$ git remote add origin https://github.com/your-username/your-repo-name.git
```
Push the changes to the remote repository:
```bash
$ git push -u origin master
```
This will upload your project to the remote repository.

That's it! You've now completed the basic Git workflow.
