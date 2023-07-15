
# password-manager

## Introduction

Pass Manager is a simple and secure password manager written in Go. It allows you to store and manage your passwords in a secure and convenient way, so you don’t have to remember them all or write them down on paper.
__But make sure to always remember your secret !!__

## Features

- __Secure storage:__ Pass Manager uses strong encryption to protect your passwords and keep them safe from prying eyes. Your passwords are encrypted and decrypted on the fly, so they are never stored in plain text.
- __Easy to use:__ Pass Manager has an intuitive command-line interface that makes it easy to use, even if you’re not familiar with the command line. You can quickly add, retrieve, update, and delete passwords with just a few simple commands.
- __Cross-platform:__ Pass Manager is written in Go, which means it can run on any platform that supports Go, including Windows, macOS, and Linux.
- __Password generation__: Supports random password generation of various types eg. numerical, alpha-numerical, etc..

## Installation

If you have go installed, you can run:

``` shell
go install github.com/sum28it/pass-manager@latest
```

This will put the binary inside __$GOPATH/bin__ directory.

Else you can download the binary for your os and architecture from bin folder and add it to your 'environment Path variable' to access it from anywhere.

## Usage

``` text
Usage:
  pass-manager [flags]
  pass-manager [command]

Available Commands:
  add         Used for adding a new user
  completion  Generate the autocompletion script for the specified shell
  delete      Deletes one or more user data
  generate    Generates a random password
  get         Retrieve user data
  help        Help about any command
  info        Prints the location of data files
  init        Initializes the application
  reset       Resets the application
  resetSecret Command to reset your secret

Flags:
  -h, --help     help for pass-manager

Use "pass-manager [command] --help" for more information about a command.
```

## Examples

``` shell
> pass-manager init 
Your secret:
Data is stored at:  C:\Users\User\.password-manager-data

> pass-manager add -a Leetcode -e prasad28sumit@gmail.com -u sum28it 
Enter password:
Your secret:
User Added!
{
        App: Leetcode
        UserId: sum28it
        Email: prasad28sumit@gmail.com
        Password: secret
        Description:
        ModifiedAt: 2023-07-04 05:01:22
}
> pass-manager get -a Leetcode 
Your secret: 
App: Leetcode   Password: Something

> pass-manager delete -a Leetcode 
Your secret:
Deleted!

> pass-manager resetSecret 
Your secret:
New secret:
Re-enter sceret:
Your secret has been changed successfully!

> pass-manager reset 
Your secret:
This will remove all your data including any password that might be saved.
Are you sure you want to do this? (Yes/No)
yes
Your application has been successsfully reset.
Use init command again before adding users.   
```
