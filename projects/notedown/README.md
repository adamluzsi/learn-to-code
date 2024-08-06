[user-story]: /topics/project-management/user-story.md
[technical-story]: /topics/project-management/technical-story.md
[user-epic]: /topics/project-management/epic.md
[backlog]: /topics/project-management/backlog.md
[entity]: /topics/programming/entity.md
[domain]: /topics/programming/ddd/domain.md
[modelling]: /topics//programming/modelling.md
[stakeholder]: /topics/project-management/stackeholder.md

# **NoteDown: A Simple Note-Taking Web Application**

NoteDown is a lightweight note-taking web application designed to help users create, edit, and organize their notes.
NoteDown will allow users to register and log in to access their private notes, format text using Markdown syntax, upload images, and search for specific notes.

## Planned Key Features

* User registration and login functionality
* Private note storage with Markdown formatting
* Image uploading and display
* Search functionality
* Simple, intuitive interface

## [Backlog][backlog]

- Present to [Stakeholders][stakeholder] what hardware and devices will be used as part of the project
  - oh now, the stakeholders are not familiar with some IT fundamentals, let's describe it to them (verbally enough).
    * [what are computer hardware, and why we need it for this project.](/topics/programming/foundation/hardware/Summary.md)
    * [What is software, what is hardware, and why we need both for this project.](/topics/programming/foundation/SoftwareAndHardware.md)
    * [what is an Operation System, and why it is benefitial for using one for the project.](/topics/programming/foundation/OperationSystemFundamentals.md)
- Enable yourself to be able to work on the project
  - [getting familiar with the terminal](/topics/programming/foundation/CommandLineInterface.md)
  - [learning about shell](/topics/programming/foundation/shell/README.md)
  - [Version control basics](/topics/programming/foundation/version-control/README.md)
    - [getting familiar with Git](/topics/programming/foundation/version-control/)
  - installing a code editor
  - creating a project directory
- Enable the system to express a given note as code
  - [Domain][domain] [Entity][entity] is a good topic to check first
  - [learning about the modelling process can be valuable for this step](/topics/programming/modelling.md)

<!-- 
- User can create new notes
- User can view all saved notes
- User can edit existing notes
- User can delete notes
- Notes are persisted across application restarts:
  - Store notes in a file system
  - Implement caching for faster access (optional)
  - Store notes in a database (future implementation)
- Multiple users can register with unique usernames and passwords:
    - Create user registration form
    - Validate username and password input
    - Hash and store passwords securely
    - Implement login functionality
- Registered users can log in to access their own notes
- Users can organize notes by category or tag (optional)
- Users can search for specific notes
- Users can format notes using Markdown syntax:
    - Accept Markdown text input from user
    - Convert Markdown text to HTML on the server-side
    - Display formatted HTML note content to user
- Users can upload images to include in notes
- Users can view uploaded images within notes
- Application displays formatted notes with HTML rendering
- Notes are stored securely and privately for each user
- User accounts can be deleted or deactivated
- Forgotten passwords can be recovered or reset -->
