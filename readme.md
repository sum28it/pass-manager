
# password-manager

A CLI tool to securely manage all your passwords at one place. All the passwords are stored on your local machine and are protected by a secret that is set by the user.

## Installation

If you have go installed, you can run:

``` bash
go install github.com/sum28it/pass-manager@latest
```

This will put the binary inside $GOPATH/bin directory.

Else you can download the built binary from bin folder and add it to your environment Path variable.

``` text
Usage:
  pass-manager [command]

Available Commands:
  add         Used for adding a new user
  completion  Generate the autocompletion script for the specified shell
  delete      Deletes one or more user data
  get         Retrieve user data
  help        Help about any command
  info        Prints the location of data files
  init        Initializes the application
  reset       Resets the application

Flags:
  -h, --help     help for pass-manager
  -t, --toggle   Help message for toggle

Use "pass-manager [command] --help" for more information about a command.
```

## Examples

