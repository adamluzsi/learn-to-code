# Create a Temporary Data Structure for Storing Notes

Create an in-memory repository for storing notes, with methods for creating, reading, updating, and deleting notes.

This implementation will be written in Go.

## Implementation Scaffold

To help get you started, here is a basic scaffold for the in-memory notes repository:

```go

type InMemoryNotesRepository struct{
    notes map[NoteID]Note
}

func NewInMemoryNotesRepository() NotesRepository {
    return &InMemoryNotesRepository{/* TODO: implement me! */}
}

// Implement the methods here:

func (r *InMemoryNotesRepository) Create(ctx context.Context, ptr *Note) error {
    // TODO: implement me!
}

func (r *InMemoryNotesRepository) FindAllNow(ctx context.Context, id NoteID) ([]Note, error) {
    // TODO: implement me!
}

func (r *InMemoryNotesRepository) FindByID(ctx context.Context, id NoteID) (_ Note, found bool, _ error) {
    // TODO: implement me!
}

func (r *InMemoryNotesRepository) Update(ctx context.Context, ptr *Note) error {
    // TODO: implement me!
}

func (r *InMemoryNotesRepository) DeleteByID(ctx context.Context, id NoteID) error {
    // TODO: implement me!
}

```

And here is an example test that you can use to verify your implementation

```go
func TestNotesRepository(t *testing.T) {
    repo := notedown.NewInMemoryNotesRepository()

    t.Run("Create", func(t *testing.T) {
        note := notedown.Note{Title: "Test Note"}
        err := repo.Create(context.Background(), &note)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }
        if note.ID == "" {
            t.Errorf("expected non-empty ID, got empty")
        }
    })

    t.Run("FindByID", func(t *testing.T) {
        note := notedown.Note{Title: "Test Note"}
        err := repo.Create(context.Background(), &note)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        foundNote, found, err := repo.FindByID(context.Background(), note.ID)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }
        if !found {
            t.Errorf("expected note to be found")
        }
        if foundNote.Title != "Test Note" {
            t.Errorf("expected title to match, got %q instead of %q", foundNote.Title, "Test Note")
        }
    })

    t.Run("Update", func(t *testing.T) {
        note := notedown.Note{Title: "Test Note"}
        err := repo.Create(context.Background(), &note)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        updatedNote := notedown.Note{ID: note.ID, Title: "Updated Test Note"}
        err = repo.Update(context.Background(), &updatedNote)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        foundNote, _, err := repo.FindByID(context.Background(), note.ID)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }
        if foundNote.Title != "Updated Test Note" {
            t.Errorf("expected title to match, got %q instead of %q", foundNote.Title, "Updated Test Note")
        }
    })

    t.Run("DeleteByID", func(t *testing.T) {
        note := notedown.Note{Title: "Test Note"}
        err := repo.Create(context.Background(), &note)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        err = repo.DeleteByID(context.Background(), note.ID)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        _, found, _ := repo.FindByID(context.Background(), note.ID)
        if found {
            t.Errorf("expected note to be deleted")
        }
    })
}
```

## Glossary

**Repository Pattern**

The Repository Pattern is a design pattern that abstracts the data access layer of an application. 
It provides a centralized interface for accessing and manipulating data, decoupling the business logic from the data storage.

In essence, a repository acts as an intermediary between the business logic and the data storage, 
providing methods for CRUD (Create, Read, Update, Delete) operations on the data.

**CRUD**

CRUD is an acronym that stands for:
* **C**: Create - adding new data to the system
* **R**: Read - retrieving existing data from the system
* **U**: Update - modifying existing data in the system
* **D**: Delete - removing data from the system

These four operations are the fundamental actions that can be performed on data,
and they form the basis of most data access patterns.

In the context of the Repository Pattern, CRUD operations are typically implemented as methods on the repository interface. 
For example:

* `Create`: adds a new entity to the data storage
* `Read` (or `FindByID`): retrieves an existing entity from the data storage by its ID
* `Update`: modifies an existing entity in the data storage
* `Delete` (or `DeleteByID`): removes an entity from the data storage by its ID

By using a repository to abstract the CRUD operations, you can decouple your business logic from the underlying data storage technology and make it easier to switch between different databases or data storage systems.
