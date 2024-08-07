[user-story]: /topics/project-management/user-story.md
[technical-story]: /topics/project-management/technical-story.md
[user-epic]: /topics/project-management/epic.md
[backlog]: /topics/project-management/backlog.md

# NoteDown

The goal of the NoteDown project is to help you learn software engineering through a hands-on, practical approach.

Rather than focusing deeply on theory, we take a different route. 
Your tasks will primarily be structured as [user stories][user-story] and [technical stories][technical-story]. 
This method aims to provide you with value-oriented goals while you work and learn, closely mirroring a typical working environment. 
Moreover, these goals offer small milestones to celebrate.

Continuous learning can sometimes feel overwhelming. Without real-world context, you might feel anxious about the value of your knowledge. To avoid a situation where the student gets trapped in a feeling of being unprepared for actual work, we take a different approach, focusing on real-world tasks that mirror work expectations.

To support your work on these realistic tasks, we will provide guidance on the necessary theoretical knowledge for each task.
This will help you stay on track, learn how to learn alongside a task and to build confidence in yourself.

## [Backlog][backlog]

- A User desire to express a given note can be formulised in code
  - []
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
- Forgotten passwords can be recovered or reset
