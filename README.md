# go-strings

go-strings prints out ASCII strings found in a binary file. This is *mostly* POSIX-compatible.

The `strings` command is extraordinarily useful for finding information in binary files. Many file formats include metadata in plain text, even when the rest of the file is binary. Possible use-cases include:
* Finding PDF and other document metadata
* Finding image, audio, and video metadata
* Finding which libraries an executable is linked with
* Finding filenames in a ZIP archive (without any installed ZIP decompression utility)

## Why?

Mac `strings` crashes when given certain input files (e.g. Java `.class` files). Windows does not have a `strings` command at all, since it is not a POSIX system.
This allows for a unified, cross-platform `strings` experience.

## Usage

```shell
strings [-n <length>] [--output-separator <separator>] [-t radix] [-a] [-w] filename
```
* `-n` is the minimum number of characters required for a string to be printed. The default is 4.
* `--output-separator` is the string that will be printed in between matches.  
Regardless of output separator, a new line will be printed at the end.
* `-t` controls how the beginning position of the string is printed.
  * `d` prints in decimal
  * `x` prints in hexadecimal
  * `o` prints in octal
  * If this flag is not set, then no location will be printed.
  * If the `POSIXLY_CORRECT` or `POSIX_ME_HARDER` environment variables are set, then `0x` and `0o` will not be printed as prefixes. If neither of these are set, then these prefixes will be printed.
* `-a` will always be ignored. It is included only for POSIX compatibility.
* `-w` enables counting new lines (whitespace) as valid characters.
  * Keep in mind that printing a carriage return to `stdout` may have the unintended consequence of overwriting some text if the immediate next character is not a line feed.
  * This flag is best used in conjunction with `--output-separator` to differentiate between new lines and separators.
