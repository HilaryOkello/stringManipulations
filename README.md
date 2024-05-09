
go-reloaded is a simple text completion, editing, and auto-correction tool written in GO. It provides various functionalities to manipulate text according to specified rules.

Functionality:

    Replace hexadecimal representation with decimal.
    Replace binary representation with decimal.
    Convert words to uppercase.
    Convert words to lowercase.
    Capitalize words.
    Modify case for a specified number of words.
    Correct punctuation spacing.
    Properly place apostrophes.
    Adjust articles "a" to "an" based on the following word.

Cloning the Project:

    $ git clone https://learn.zone01kisumu.ke/git/hilaokello/go-reloaded.git

Testing:
To test the program, execute:

    $ go test

For verbose output:

    $ go test -v

Project Structure:

    go-reloaded/
    ├── go.mod
    ├── main.go
    ├── main_test.go
    ├── manipulations/
    │   ├── manipulator.go
    │   ├── manipulator_test.go
    │   ├── punctuation.go
    │   └── punctuation_test.go
    └── README.md

Package manipulations:

    manipulator.go: Contains functions to manipulate words based on specified rules.
    manipulator_test.go: Tests the functionality of manipulator.go.
    punctuation.go: Functions to handle punctuation formatting.
    punctuation_test.go: Tests the punctuation handling functions.

Main:

    main.go: Reads from sample.txt, applies manipulations, and writes the output to result.txt.
    main_test.go: Tests that our main.go runs as required and passes the following tests:

    $ cat sample.txt
    it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
    $ go run . sample.txt result.txt
    $ cat result.txt
    It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
    $ cat sample.txt
    Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
    $ go run . sample.txt result.txt
    $ cat result.txt
    Simply add 66 and 2 and you will see the result is 68.
    $ cat sample.txt
    There is no greater agony than bearing a untold story inside you.
    $ go run . sample.txt result.txt
    $ cat result.txt
    There is no greater agony than bearing an untold story inside you.
    $ cat sample.txt
    Punctuation tests are ... kinda boring ,don't you think !?
    $ go run . sample.txt result.txt
    $ cat result.txt
    Punctuation tests are... kinda boring, don't you think!?


