# Version Control

Version control is a system that helps you manage changes to documents, code, or other digital content over time. 
It's an essential tool for any software development project, allowing multiple developers to collaborate on the same codebase and track changes made by each team member.

## Why Use Version Control?

Using version control offers several benefits:
* **Collaboration**: Multiple developers can work on the same codebase simultaneously without conflicts.
* **Backup and recovery**: All versions are stored securely, allowing for easy recovery in case of data loss or corruption.
* **Change tracking**: A clear record of changes made to each file or document helps identify who made changes and when.
  * especially useful for learning about what reasons contributed to changes of a given code
* **Rollback**: Easily revert to a previous version if changes cause issues.
  * experimenting is important, and version control enable safe points where we can always go back
    * recover mistakenly deleted code
    * finding a stable state to help identify changes which introduced a potential issue

## Types

1. **Centralized Version Control System (CVCS)**: All team members access the central repository, where all versions are stored. Examples include Subversion (SVN) and Perforce.
2. **Distributed Version Control System (DVCS)**: Each developer has a local copy of the entire project history, allowing them to work offline and synchronize changes with others when needed. Examples include Git and Mercurial.

## Key Concepts

* **Repository**: A central location where all versions are stored.
* **Branch**: A separate line of development within a repository, often used for feature development or bug fixes.
* **Commit**: The act of saving changes to a file or document.
* **Merge**: Combining changes from two branches into a single branch.

## Best Practices

To get the most out of version control:
* **Use meaningful commit messages** to describe changes made in each commit.
* **Create regular backups** of your repository to prevent data loss.
* **Communicate with team members** about changes and updates to avoid conflicts.
* **Test thoroughly** before merging changes into a production branch.

## Getting Started

- [Git](/topics/programming/foundation/version-control/git.md)

These companies give you remote repository hosting solutions using git:
    - [GitHub](https://github.com/)
    - [Bitbucket](https://bitbucket.org/)
    - [GitLab](https://gitlab.com/)
