# go-strings

go-strings prints out ASCII strings found in a binary file. This is similar to the `strings` command on \*NIX systems, but with fewer features and different flags.

## Why?

Mac strings crashes when given certain input files (e.g. Java `.class` files). Windows does not have a strings command at all.
This allows for a unified, cross-platform strings experience.

## Usage

```shell
strings [-length <length>] [-output-separator <separator>] [-t radix] filename
```
* Length is the minimum number of characters required for a string to be printed. The default is 4.
* Output separator is the string that will be printed in between matches.  
Regardless of output separator, a new line will be printed at the end.
* Radix controls how the beginning position of the string is printed.
  * `d` prints in decimal
  * `x` prints in hexadecimal
  * `o` prints in octal
  * If the `POSIXLY_CORRECT` or `POSIX_ME_HARDER` environment variables are set, then `0x` and `0o` will not be printed as prefixes. If neither of these are set, then these prefixes will be printed.  
  If this flag is not set, then no location will be printed.

