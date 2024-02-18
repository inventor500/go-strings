# strings

`strings` prints out ASCII strings found in a binary file. This is a *mostly* POSIX-compatible implementation.

The `strings` command is extraordinarily useful for finding information in binary files. Many file formats include metadata in plain text, even when the rest of the file is binary. Possible use-cases include:
* Finding PDF and other document metadata
* Finding image, audio, and video metadata
* Finding which libraries an executable is linked with
* Finding filenames in a ZIP archive (without any installed ZIP decompression utility)

If you want to find out what interesting info a file may have, then `strings <filename> | less` is a good starting place.

## Why?

There are already several versions of the `strings` command out there, since it is part of the POSIX specification.
* GNU `strings` is great, but it is not always easy to put on non-GNU platforms.
* Mac `strings` has various issues, including crashing on some files (e.g. Java `.class` files). When these files are piped in, Mac `strings` prints out binary data.
* Go is trivial to port to other platforms.

## Usage

```shell
strings [-n <length>] [--output-separator <separator>] [-t <radix>] [-a] [-w] filename
```
* `-n` is the minimum number of characters required for a string to be printed. The default is 4.
* `--output-separator` is the string that will be printed in between matches.  
Regardless of output separator, a new line will be printed at the end.
* `-t` controls how the beginning position of the string is printed.
  * `d` prints in decimal
  * `x` prints in hexadecimal
  * `o` prints in octal
  * If this flag is not set, then no location will be printed.
  * This program prints `0x` before hexadecimal numbers and `0o` before octal numbers. While this makes these numbers easier to identify, this is not POSIX-standard behavior.
    * The POSIX behavior can be turned on by setting the `POSIXLY_CORRECT` OR `POSIX_ME_HARDER` environment variables. These names were chosen because they have been previously used by GNU to enable POSIX-compliant, user-unfriendly behavior.
* `-a` will always be ignored. It is included only for POSIX compatibility.
* `-w` enables counting new lines (whitespace) as valid characters.
  * Keep in mind that printing a carriage return to `stdout` may have the unintended consequence of overwriting some text if the immediate next character is not a line feed.
  * This flag is best used in conjunction with `--output-separator` to differentiate between new lines and separators.
* If the input file is "`-`", then strings will read from standard input instead of from a file. If you really want to read from a file named "`-`", use "`./-`" instead.

### Example Usage

```shell
$ strings image.png
# ...
# <?xpacket begin="
# Lots of XML metadata...
# <?xpacket end="w"?>

# All PDFs start with the version. Much of the metadata, if provided, is in plain text
$ strings document.pdf
# %PDF-1.4
# 1 0 obj
# /Creator ...
# /CreationDate ...
# ...
# <?xpacket begin="
# " id=...
# More XML metadata...
# <?xpacket end="w"?>
# ...

$ strings song.opus
# OggS
# OpusHead
# OggS
# zOpusTags
# <encoder was here>
# language=eng
# album=<album name was here>
# title=<title was here>
# encoder=<encoder was here>
# ...

# MP3 metadata tables are mostly plain text
$ strings -n 8 song.mp3
# <encoder was here>
# <title>TPE1
# <artist>TALB
# <album>TDRC
# <date>TRCK
# ...
# <encoder was here>
# ...

# Videos put much of their metadata in plain text
$ strings -n 10 video.webm
# <encoder>
# <encoder>
# MAJOR_BRANDD
# MINOR_VERSIOND
# <encoder>
# HANDLER_NAMED
# <comment about producer and date>
# VENDOR_IDD
# ...

# Encrypted zips don't encrypt the file names!
# (Don't use these if data integrity is important - attackers can also just remove or replace the file in the zip)
$ strings encrypted-zip.zip
# <filename 1>UT
# ...
# <filename 2>UT
# ...
# <filename 3>UT
# ...
# <filename 1>UT
# <filename 2>UT
# <filename 3>UT


# Compiled code has much useful information still in plain text...
# This just lists some of the copyright info, for binaries that still have it
$ for file in /bin/*
$ do
$ 	$(echo "$file: $(strings $file | grep -i copyright)"
$ done
# ...
# /bin/[: Copyright %s %d Free Software Foundation, Inc.
# ...
# /bin/bzip2:    Copyright (C) 1996-2019 by Julian Seward.
# ...

$ strings <compiled-rust-and-c++-program>
# <linker>
# ...
# curl_easy_perform
# curl_easy_init
# curl_easy_strerror
# curl_easy_cleanup
# curl_easy_setopt
# <other interesting stuff>
# ...
# <too much info to put here>
# <GCC version>
# <rust version>
# ...
```

## Compilation

Compiling for your current OS/architecture is easy:

```shell
$ cd cmd/strings
$ go build
```

If you want to compile for a different OS/CPU architecture, just set the `GOOS` and `GOARCH` environment variables.
