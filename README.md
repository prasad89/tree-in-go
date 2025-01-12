# Tree CLI Application

## Overview

This is a simple Go program that replicates the functionality of the `tree` command, displaying the directory structure of a specified root folder.

## Features

- Displays directory structure in a tree format.
- Skips hidden files and directories (those starting with `.`).
- Handles errors gracefully if the specified directory is inaccessible.

## Setup and Usage

### Build the Application

To build the application, use the provided `Makefile`. Run:

```sh
make build
```

The binary will be generated in the `bin/` directory with the name `tree`.

### Run the Application

To execute the application, use:

```sh
make run
```

By default, the program displays the structure of the current directory. You can specify a different directory:

```sh
./bin/tree /path/to/directory
```

### Clean Up

To remove the built binary, run:

```sh
make clean
```

## Example Output

For the following directory structure:

```
project/
├── file1.txt
├── file2.txt
└── folder
    └── file3.txt
```

The program will output:

```
project
├── file1.txt
├── file2.txt
└── folder
    └── file3.txt
```
