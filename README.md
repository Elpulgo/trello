![Go Linux](https://github.com/Elpulgo/trello/workflows/Go%20Linux/badge.svg?branch=master)
![release](https://github.com/Elpulgo/trello/workflows/release/badge.svg?branch=0.1.1&event=release)

# Trello CLI
Show boards, lists, cards, comments from your Trello board. Requires an ApiKey and a token.

## Features
 + List all boards for user, owned aswell as shared.
 + Persist Trello ApiKey and Token with a password.
 + Optionally store the password for a seamless use of the CLI. 
   Password is otherwise required for each command.
 + Boards
    + NumericShort
    + Name
    + Id
    + Link to board 
 
 + Single board
    + List
    + Cards
      + Id
      + Nr of comments
      + Name
 + Card
    + User @ Timestamp
    + Comments

## How to use
```
$ tre boards -b 2  
```
```
$ tre card -c 5efc3a17f7910117f0e3b88b
```


