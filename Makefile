# Makefile for Snake and Ladder Game

BINARY_NAME=snake-ladder
CONFIG_FILE=config/gameConfigOne.json

build:
	go build -o $(BINARY_NAME) -v

run: build
	./$(BINARY_NAME)

.PHONY: build run