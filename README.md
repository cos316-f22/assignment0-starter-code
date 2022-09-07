# COS316, Assignment 0: Setting up the Development Environment

In this "warmup" assignment, we you will get your development environment set up
with the necessary language compilers, libraries, and tools to complete future
assignments and participate in precepts. In the end, you will run a command that
exercises all of these tools and write a computed value to a file:

``` sh
$ make
[Answer printed here and also saved to the file `answer`]
```

The value will be a long string of seemingly random hexademical characters (a
SHA256 hash of expected values from each necessary tool).

You will be "graded" on the correctness of this value, verifying that your
tool-chain is setup.

## Development Environment for Assignments and Precepts

The assignments and precepts in this course assume that you have a few tools and
libraries available in your development environment. Any way you get them is
fine, and they should generally be relatively easy to find and install on most
Linux distributions, BSDs, OS X, and Windows using Windows Subsystem for
Linux (WSL)[^1].

[^1]: You _may_ be able to complete assignments with these tools installed on
    Windows _without_ WSL, but your mileage may vary and course staff may not be
    able to offer help with such a setup. We strongly recommend enabling and
    using WSL if you are on Windows.

  * [Git](https://git-scm.org)
  * [Go](https://go.dev) version 1.18
  * [SQLite](https://www.sqlite.org/index.html) version 3
  * [curl](https://curl.se)
  * [gcc/g++](https://gcc.gnu.org/) or [clang/clang++](https://clang.llvm.org/) compilers
    for C & C++ (other C/C++ compiler may or may not work).
  * [make](https://www.gnu.org/software/make/)

In additional you should have either Firefox or Chrom[e|ium] installed for one of the precepts where we will inspect HTTP traffic using those browsers[^2].

[^2]: Recent versions of the Edge browser which are based on Chromium and
include Chromium's developer tools will likely also work.

## Verifying you have the right tools installed

Once you've installed the tools (instructions in the next sections), running the
following commands in a shell should result in similar output (slight variations
are expected on different operating systems).

### Git

``` sh
$ git version
git version 2.34.1
```

### Go

``` sh
$ go version
go version go1.18.5 linux/amd64
```

### SQLite 3

``` sh
$ echo "select (316 + 1021)" | sqlite3
1337
```

### curl

``` sh
$ curl -I https://cos316.princeton.edu
HTTP/2 200
server: nginx
date: Wed, 26 Jan 2022 16:41:20 GMT
content-type: text/html
content-length: 3859
etag: "xsmg9fnyjsw7c3qidhrrb62g071vfnls"
accept-ranges: bytes
```

### C++ compiler

``` sh
$ g++ --version
g++ (GCC) 10.3.0
Copyright (C) 2020 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```

or

``` sh
$ clang --version
clang version 7.1.0 (tags/RELEASE_710/final)
Target: x86_64-unknown-linux-gnu
Thread model: posix
InstalledDir: /nix/store/ass1sf1bx07qvlrg02nymxnyzp1cpxz7-clang-7.1.0/bin
```

### Make

``` sh
$ make --version
GNU Make 4.3
Built for x86_64-pc-linux-gnu
Copyright (C) 1988-2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
```

## Installing on Linux

Most recent Linux distributions should have the necessary tools available from
their respective package managers. Some Long-Term-Support (LTS) versions may
have an older version of Go, in which case you can use the instructions on the
Go website to install Go 1.17.

Below are instructions for a couple popular distributions. You will typically
need root permissions (i.e. prefix the commands with `sudo` or login as root) to
install packages:

### Debian (Backports/Testing/Unstable) & Ubuntu (Rolling)

``` sh
$ apt-get install git curl build-essential golang sqlite3
```

### Arch Linux

``` sh
$ pacman -S git go sqlite3 curl base-devel
```

### Alpine Linux

``` sh
$ apk add git go build-base sqlite curl
```

## Installing on OS X

First, you must have the [XCode Command Line Tools](https://developer.apple.com/downloads/index.action) from Apple installed. These already include a C/C++ compiler and `make`.

You can obtain the remaining tools from a number of package managers available for OS X (e.g. Homebrew, MacPorts), or you can also download and install universal packages from each tool's respective website.

A simple way using [MacPorts](https://www.macports.org/) is with the following command:

``` sh
$ sudo port install git go curl sqlite3
```

## Installing on Windows under WSL

To work in a Windows environment, we recommend using Windows Subsystem for
Linux, which provides a Linux environment alongside Windows. By default, WSL
provides a Ubuntu distribution of Linux. Follow the installation instructions
for WSL from the [official
documentation](https://docs.microsoft.com/en-us/windows/wsl/install), then use
the instructions above for Linux to install the tools.

Most likely, you will be using the default Ubuntu flavor of WSL, in which case
the following command will install all the necessary tools:

``` sh
$ apt-get install git curl build-essential golang sqlite3
```

## Submission and grading

To the generated `answer`, commit the `answer` file and push it to your GitHub repository.

``` sh
$ git add answer
$ git commit
[Write a commit message]
$ git push
```

See the [Instructions on the course
website](https://cos316.princeton.edu/assignments) for more detailed
instructions.

We will grade your submission by validating the answer generated by the build
script matches the expected result. This indicates that you have all the tools
installed and are using the correct version of the Go compiler.

You may submit in this way and receive feedback as many times as you like,
whenever you like, but a lateness penalty will be applied to submissions
received after the deadline.
