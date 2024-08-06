# NoteDown

The NoteDown project will help you learn web development using the Go language.

## modules

### Module 1.1: Create a New Project

- [ ] Set up a new Go project using your preferred IDE or text editor
- [ ] Initialize a Git repository and make an initial commit

### Module 1.2: Hello world!

- [ ] create `cmd/myapp` directories in your project
- [ ] Have a `main.go` in it.
  - [ ] when running this app using `go run` we should have some startup message greeting us


### Module 1.2: Define the Note App Requirements

In this module, we'll define the initial requirements for our note taking app. 
We'll consider the following aspects:

- Health check endpoint:
  - [ ] the app should be able to retrun back an HTTP status code 200 when the `GET /health` HTTP endpoint is called
    - for this you need to learn about go's http package
- Landing Page:
  - [ ] The app should display a landing page that welcomes users and invites them to start creating notes.
    - index.html could be found in the `assets` directory, however you nead to learn how to read this file and serve it over HTTP
  	- [ ] Update the landing page to should include a clear call-to-action (CTA) to create a new note.
- Note Creation:
  - [ ] Users should be able to create new notes with a content.
  - [ ] Notes should be saved automatically as the user press a button 
- Note Listing:
  - [ ] The app should display a list of all created notes on the landing page or a separate page.
  - [ ] Notes should be displayed in a reverse chronological order (newest first).
  - [ ] Each note should have a clear title, content preview, and tags (if any).
- Note Editing:
  - [ ] Users should be able to edit existing notes by clicking on them from the list view.
  - [ ] The app should allow users to update the note's title, content, and tags.

