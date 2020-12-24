# go-touch

The go interpetation of UNIX `touch` command. The motivation behind this was
I couldn't find anything like `touch` command in windows, so I decided to create
my own.

It does not completely copies the command, if file exists it changes its
timestamp, if it doesn't, creates a new file.

## Installation

Go to the [releases](1) page, install `touch.exe`, and add directory, where `touch.exe` lives,
to the `PATH` variable.

In Windows

```bat
set PATH=%PATH%;C:\path\to\dir\of\touch
```

### Installation from source

#### Requirements for installing from  sourcre

* golang compiler >= 1.15

I didn't test it with old versions of go, so I wrote `>=1.15`, but I don't think
it could fail with older versions
<hr>

Clone the repository, and `cd` into it

```bat
git clone https://github.com/spitfire-hash/go-touch
cd go-touch
```

Install the package

```bat
go install spitfire-hash/touch
```

Don't forget to add your `~/go/bin` directory to your `PATH` variable.
First find the `bin` path:

```bat
go list -f '{{.Target}}'
```

Then the output should be something like this:

```bat
'C:\Users\<User>\go\bin\touch.exe'
```

You should add the `C:\Users\<User>\go\bin` to your `PATH`, you can do it by

```bat
set PATH=%PATH%;C:\path\to\your\install\directory
```

## Usage

```sh
touch FILE
```

## Tests

For running test just do

```sh
go test
```

In the repository directory.
