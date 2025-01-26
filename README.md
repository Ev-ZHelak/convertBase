# Number Base Converter in Go

This project is a converter for numbers between different number systems.

## Description

Main features:
- Displays intermediate steps when converting to the hexadecimal system.
- Simple user interface with prompts.

## How to Use

The program asks the user to input a number and the base system, for example:

```
Enter a number and the base system (e.g., 29 16):
```

After input, the program performs the conversion and displays the result in the specified base.

Example program output:

```
Enter a number and the base system (e.g., 29 16): 29 2
Performing conversion:
29 / 2 = 14 remainder 1
14 / 2 = 7 remainder 0
7 / 2 = 3 remainder 1
3 / 2 = 1 remainder 1
1 / 2 = 0 remainder 1
Rewriting remainders in reverse order:
Result = 11101
```

## Features

- The program supports conversion to any base between 2 and 16.
- For hexadecimal conversion, the program provides additional explanations of how numbers 10-15 are represented as letters A-F.

## Educational Purpose

This project is created with the goal of teaching the basics of programming in Go, including handling user input, working with loops and conditionals, and using packages for output formatting.
