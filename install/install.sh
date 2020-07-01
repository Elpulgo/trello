#!/usr/bin/env bash

install() {
    export FILE="tre"
    curl "https://raw.githubusercontent.com/Elpulgo/trello/master/dist/tre.tar.gz" | tar xvz
    chmod +x tre
    sudo mv tre /usr/local/bin/tre
}

install