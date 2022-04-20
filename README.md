# un will [un]structure your structured json logs

Parse json logs and make them pretty in the console

## Intro

Un will help you view your structured json logs in a terminal

## Install

    go get github.com/yavosh/un

## Examples

    cat file.json | un
    stern container -o raw --tail 0 | un
