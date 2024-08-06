[domaindoc]: /topics/programming/ddd/domain.md

# Entity

When we model something in code, we're capturing its essential characteristics, behaviours, and relationships. This involves identifying the key aspects of the concept we're modelling and representing them using programming constructs like data structures, classes, or objects. 

An entity expresses data that has [domain][domaindoc] value, often called a [Domain][domaindoc] Entity. These terms are interchangeable.

A [domain][domaindoc] entity represents a data oriented concept from the problem [domain][domaindoc] we're solving. 
It's a model of something meaningful and valuable in our application.

Think of a [domain][domaindoc] entity as a container for data with a specific purpose or significance within our system. It’s not just a collection of random fields; it represents a cohesive unit of data with its own identity, behaviour, and constraints.

## Example: Representing a Note 

Let's take the example, where we want to represent a note in our code. 
A simple implementation might look like this:

```go
type Note struct {
    ID       string
    Title    string
    Content  string
    CreatedAt time.Time
}
```

In this case, `Note` is a [domain][domaindoc] entity that represents a single note with its associated data fields (ID, title, content, and creation date). By creating a separate type for the note, we've given it an identity and structure that reflects its meaning within our application.

A Note value refers to a specific instance of a [domain][domaindoc] entity.
Using the `Note` example again, a value would be an actual instance of the `Note` struct with its fields populated:

```go
note := Note{
    ID:       1,
    Title:    "My First Note",
    Content:  "This is my first note.",
    CreatedAt: time.Now(),
}
```

In this case, `note` is a [domain][domaindoc] value that represents a specific note with its own unique characteristics (ID, title, content, and creation date).


## Benefits of Modeling [Domain][domaindoc] Entities

By modeling [domain][domaindoc] entities in our code, we can:

- Improve understanding and communication among team members about the problem [domain][domaindoc].
- Improve the reusability of our code by having a common way to represent a [domain][domaindoc] value.
- Create more robust and maintainable software by separating concerns and encapsulating data and behavior.
- Better capture the essential characteristics and constraints of real-world concepts.
