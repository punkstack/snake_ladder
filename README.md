# Snake And Ladder Game

## Overview

The objective is to design and implement a snake and ladder game, which should support the following set of operations/functionality and be extensible for the future:

1. Create a board of size 0-100.
2. Make 2 types of dice - Normal (1-6)/ Crooked (only odd numbers)
3. Users should be able to create a game
4. Users should be able to join a game.
5. Whoever reaches the final position first is the winner.

## Packages

- `models`: Contains the data models for the game.
- `services`: Contains the service layer for game logic.
- `main.go`: Entry point of the application.

## MAKE Commands
Use Make commands to build application.

```bash
make build
```

```bash
make run
```

## Assumptions

* PLayer will have only one dice to roll

## Next Steps

* Improvement of I/O validation
* More validations logic for the game
* Need more test cases to improve the test coverage
* CI pipeline for tests