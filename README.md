# cyoa
Choose Your Own Adventure Series Web Implementation

## Motivation
This project aims to recreate __Choose Your Own Adventure__ series experience via a web application
where each page will be a portion of a story, and at the end of every page the user will be given a   
series of options  to choose from (or be told that they have reached the end of that particular story arc).

## Workings
Stories will be provided via a JSON file with the following format:
```json
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
```   
For example:
```json
{
    "intro": {
        "title": "The Little Blue Gopher",
        "story": [
            "Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
            "One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
            "On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
        ],
        "options": [
            {
                "text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
                "arc": "new-york"
            },
            {
                "text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
                "arc": "denver"
            }
        ]
    }
}
```
A few things worth noting:   
  - Stories could be cyclical if a user chooses options that keep leading to the same place. 
  - All stories will have a story arc named "intro" that is where the story starts. 
  
## Usage
 - Clone/Download the repo
 - from the root, run  
  ```go run ./cmd/web -addr :8000 -json sample.json```   
  where sample.json is a json file containing stories in the same format as above. 
  
