
# password-manager

## Introduction

A CLI application to securely manage all your passwords at one place. All the passwords are stored on your local machine and are protected by a secret that is set by the user.
You can also backup the generated files anywhere, so that you can port them on any other machine.
__But make sure to always remember your secret !!__

## Installation

If you have go installed, you can run:

``` shell
go install github.com/sum28it/pass-manager@latest
```

This will put the binary inside $GOPATH/bin directory.

Else you can download the binary for your os and architecture from bin folder and add it to your 'environment Path variable' to access it from anywhere.

## Usage

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

``` shell
> pass-manager init secret 
Data is stored at:  C:\Users\User\.password-manager-data

> pass-manager add -a Leetcode -e prasad28sumit@gmail.com -u sum28it -p Something secret
User Added!

> pass-manager get -a Leetcode secret
App: Leetcode   Password: Something

> pass-manager delete -a Leetcode secret
Deleted!

> pass-manager reset secret
This will remove all your data including any password that might be saved.
Are you sure you want to do this? (Yes/No)
yes
Your application has been successsfully reset.
Use init command again before adding users.   
```
