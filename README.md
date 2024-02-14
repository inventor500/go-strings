# go-strings

go-strings prints out ASCII strings found in a binary file. This is similar to the `strings` command on \*NIX systems, but with fewer features and different flags.

## Why?

Mac strings crashes when given certain input files (e.g. Java `.class` files). Windows does not have a strings command at all.
This allows for a unified, cross-platform strings experience.

## Usage

```shell
strings [-length <length>] [-output-separator <separator>] filename
```
* Length is the minimum number of characters required for a string to be printed.
* Output separator is the string that will be printed in between matches.  
Regardless of output separator, a new line will be printed at the end.

