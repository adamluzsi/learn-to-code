[user-story]: /topics/project-management/user-story.md
[technical-story]: /topics/project-management/technical-story.md
[user-epic]: /topics/project-management/epic.md
[backlog]: /topics/project-management/backlog.md
[entity]: /topics/programming/entity.md
[domain]: /topics/programming/ddd/domain.md
[modelling]: /topics//programming/modelling.md
[stakeholder]: /topics/project-management/stackeholder.md
[hardware]: /topics/programming/foundation/hardware/Summary.md
[software-hardware]: /topics/programming/foundation/SoftwareAndHardware.md
[operation-system]: /topics/programming/foundation/OperationSystemFundamentals.md
[command-line-interface]: /topics/programming/foundation/CommandLineInterface.md
[shell]: /topics/programming/foundation/shell/README.md
[version-control]: /topics/programming/version-control/README.md
[git-getting-started]: /topics/programming/version-control/git/getting-started.md
[code-editor]: /topics/editor.md
[client]: /topics/project-management/client.md

# Learn to Code

In this repository you will see some learning materials about programming, that probably help you begin your journey.

The goal of this project is to help you learn software engineering through a hands-on, practical approach.

Rather than focusing deeply on theory, we take a different route. 
Your tasks will primarily be structured as [user stories][user-story] and [technical stories][technical-story]. 
This method aims to provide you with value-oriented goals while you work and learn, 
closely mirroring a typical working environment. 
Moreover, these goals offer small milestones to celebrate.

Continuous learning can sometimes feel overwhelming. Without real-world context, you might feel anxious about the value of your knowledge. To avoid a situation where the student gets trapped in a feeling of being unprepared for actual work, we take a different approach, focusing on real-world tasks that mirror work expectations.

To support your work on these realistic tasks, we will provide guidance on the necessary theoretical knowledge for each task.
This will help you stay on track, learn how to learn alongside a task and to build confidence in yourself.

Things marked with `#LE` are meant to practiced with a mantor, as a live exercise.

## **NoteDown: A Minimalist Note-Taking Web Application**

NoteDown is a lightweight note-taking web application designed to help users create, edit, and organize their notes.
NoteDown will allow users to register and log in to access their private notes, format text using Markdown syntax, upload images, and search for specific notes.

### [Backlog][backlog]

- Investigate what is the vision of the [stakeholders][stakeholder] about the project #LE
  - quick introduction on how to gather requirements from a [client][client]/[stakeholder][stakeholder]
- Present to [Stakeholders][stakeholder] what hardware and devices will be used as part of the project #LE
  - oh now, the stakeholders are not familiar with some IT fundamentals, let's describe it to them (verbally enough).
    * [what is computer hardware, and why do we need it for this project?][hardware]
    * [What is software, what is hardware, and why do we need both for this project?][software-hardware]
    * [what is an Operation System, and why it is beneficial to use one for the project?][operation-system]
- Enable yourself to be able to work on the project
  - [getting familiar with the terminal][command-line-interface]
  - [learning about shell][shell]
  - [Version control basics][version-control]
    - [getting familiar with Git][git-getting-started]
  - [choose and install a code editor][code-editor]
  - creating a project directory
    - [bonus points if you create the directory using the shell][shell]
- Enable the system to express a given note as code
  - [Domain Entity][entity] is a good topic to check first
    - if you are unsure what is ["domain" is, then you can read more about it here][domain]
  - [learning about the modelling process can be valuable for this step][modelling]

<!-- 
- Users can create new notes
- Users can view all saved notes
- Users can edit existing notes
- Users can delete notes
- Notes are persisted across application restarts:
  - Store notes in a file system
  - Implement caching for faster access (optional)
  - Store notes in a database (future implementation)
- Multiple users can register with unique usernames and passwords:
    - Create a user registration form
    - Validate username and password input
    - Hash and store passwords securely
    - Implement login functionality
- Registered users can log in to access their notes
- Users can organize notes by category or tag (optional)
- Users can search for specific notes
- Users can format notes using Markdown syntax:
    - Accept Markdown text input from a user
    - Convert Markdown text to HTML on the server side
    - Display formatted HTML note content to the user
- Users can upload images to include in notes
- Users can view uploaded images within notes
- The application displays formatted notes with HTML rendering
- Notes are stored securely and privately for each user
- User accounts can be deleted or deactivated
- Forgotten passwords can be recovered or reset 

-->