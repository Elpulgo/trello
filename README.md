![Go Linux](https://github.com/Elpulgo/trello/workflows/Go%20Linux/badge.svg?branch=master)
![release](https://github.com/Elpulgo/trello/workflows/release/badge.svg?branch=0.1.1&event=release)

# Trello CLI
Show boards, lists, cards, comments from your Trello board. Requires an ApiKey and a token.

## Features
 + List all boards for user, owned aswell as shared.
 + Add, move and comment on single card
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
+ `boards`
    + `-b / --board`
      + Show cards on a specific board, specified with either # or id.
    + `-l / --listname`
      +  Pass listname, for a specific list. Must be combined with -b/-n
    + `-n / --name`
      + Show cards on a specific board, specified with a name.

+ `card`
    + `-c / --id`
      + (*) Card Id, required.
    + `add`
      + `-d / --description`
        + Description for the card.
      + `-l / --listid`
        + (*) List id. The list the card should belong to.
      + `-t / --title`
        + (*) Title of the card.
    + `comment`
      + `-i / --cardid`
        + (*) Card id. The card the comment should belong to.
      + `-c / --comment`
        + Comment.
    + `move`
      + `-c / --cardid`
        + (*) Id of the card, required.
      + `-l / --listid`
        + (*) List id. The list the card should be moved to.
    
+ `credentials`
    + `-p / --passphrase`
      + (*) Passphrase for API credentials.
    + `-s / --store`
      + Should store passphrase in 'pass.dat' (y/n)
    + `-k / --trello-key`
      + (*) Trello API key.
    + `-t / --trello-token`
      + (*) Trello API token.

+ `help / -h`
  + Help command, can be executed with any command.

## Installation
+ Linux `$ curl "https://raw.githubusercontent.com/Elpulgo/trello/master/install/install.sh" | bash`
+ Windows <a href="https://github.com/Elpulgo/trello/releases/download/v0.1.3/tre-windows-amd64.tar.gz">.tar.gz for Windows</a>

## Screen

## Future functionality
  + Trello link to card/member
  + Attachments link on cards
  + Badges, showing different colors
  + Checklist on card


