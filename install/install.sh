#!/usr/bin/env bash

install() {
    export FILE="tre"
    curl "https://github.com/Elpulgo/trello/releases/download/v0.1.3/tre-linux-amd64.tar.gz" | tar xvz
    chmod +x tre
    sudo mv tre /usr/local/bin/tre
}

install