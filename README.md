
# gochess

## Demo

Demo picture:

<p align="center">
  <img src="./resources/example.png" width="300"/>
</p>

## Features

List of features

  - [ ] xxx

## Tasks

List of things to do

  - features.
    - [ ] flip the sides of the board option.
    - [ ] undo history.
  - style.
    - [ ] have a flag for **textual** mode: RNBQKBNR pppppppp
    - [ ] have a flag for **graphical** mode: RNBQKBNR pppppppp

Done

  - [X] implement CLI background color checkered board in black and white.
  - [X] implement CLI foreground colors for pieces (in black and white).

List of ideas

  - misc.
    - [ ] implement timer.
    - [ ] implement piece value display.
    - [ ] implement simple computer player (ie AI enemy).
    - [ ] option to customize background colors.
    - [ ] option to customize foreground colors.
    - [ ] allow new pieces, like the duck (moves everywhere, cannot be beat, blocks all attacks, is moved by both players, duck does not move, if move opens a new attack on the enemy player).
    - [ ] allow randomized starting positions.
    - [ ] allow randomized piece values (white and black adding up to the same total).
    - [ ] allow randomized piece values (white and black adding up to differing totals).
  - check for moves to be legal.
    - [ ] have a flag for **manual** mode: all moves allowed.
    - [ ] have a flag for **turns** mode: moves must follow turns.
    - [ ] have a flag for **checked** mode: all moves must be legal.
    - [ ] have a option for turning off **castling** mode: disallow castling at all.
    - [ ] have a option for turning on **attacked** mode: sallow castling through attacked squares.
    - [ ] have a option for turning off **en-passant** mode: disallow en-passants.
    - [ ] have a option for turning off **undos** mode: disallow undos.

## Installation

Install the program:

```bash
go install github.com/kraasch/gochess@latest
```

Get the package:

```bash
go get github.com/kraasch/gochess
```

## Usage

Use the program:

```bash
make build
./build/gochess --help
```

Use the package:

```go
import (
  "github.com/kraasch/gochess"
)

gochess.DoSomething("Hello")
```

## Feedback

I can be reached via [alex@kraasch.eu](mailto:alex@kraasch.eu).

## Contributing

Feel free to help me.

## Acknowledgments

Uses the following software:

  - see [go.mod](./go.mod) and [go.sum](./go.sum).

Made by the following people:

  - see Github info.

## License

View the [license file](./LICENSE).

