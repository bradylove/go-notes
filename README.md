# src-notes
========

Rails rake notes command for all source code using Go

## Installation

With [Go](http://golang.org) 1.2 or higher installed (may work on older versions,
but I have not tried) run the following commands:

    $ go get github.com/bradylove/src-notes
    $ go install github.com/bradylove/src-notes

Once I get a few more features built in I will start releasing binaries for Mac
and Linux.

## Usage

src-notes currently supports comments with "todo" and "note. Your notes should
have the languages single line comment, then a space, then `Note:` or `Todo:`
then a space then your comment. See example below.

    // Todo: This is a valid comment

To run the app and get a list of notes and todo's in a directory run the
`src-notes` command to search in the current directory, or you can optionally
provide the path to a directory. See examples below.

    // Search for notes and todo's in current directory
    $ src-notes

    // Search for notes and todo's in app/controllers directory
    $ src-notes app/controllers

## Note

Things are pretty basic right now, and there arent any options except for the
 directory to search in. I have a few things in mind to add and will work on
 that over the next little while, however if you wish to add a feature for
 yourself feel free to send a pull request.

To see a list of supported files check out the `mappings.go` file, and if
you don't see your favorite language supported please let me know or send a
pull request.

## License

The MIT License (MIT)

Copyright (c) 2014 Brady Love

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
