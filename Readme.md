# takuhai

is a CLI application to track packages in your terminal.

# Installation
```
$ git clone https://github.com/lapla-cogito/takuhai
$ make
$ echo PATH=$PATH:$(pwd)/takuhai/bin/takuhai >> ~/.bashrc
```

If you are using other unix shell like zsh, replace "~/.bashrc" and "bash" properly.

# Currently available carriers

- YAMATO TRANSPORT CO., LTD.

- SAGAWA EXPRESS CO.,LTD.

- Japan Post Co., Ltd.

# How to use

```
$ takuhai -h
A CLI application to track packages you registered.
currently, this application can track:
- SAGAWA TRANSPORTATION CO., LTD.
- YAMATO TRANSPORT CO., LTD.
- Nippon Express Co., LTD.

Usage:
  takuhai [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dereg       Deregister the specific package
  export      Export package informations
  help        Help about any command
  import      Import package informations from exported yaml file
  reg         Registers a package
  rename      Rename a package
  show        Shows the state of a specific package
  timeline    Shows the timeline of a specific package

Flags:
  -h, --help   help for takuhai

Use "takuhai [command] --help" for more information about a command.
```

# Usage examples

## Register a package delivered by Sagawa, its tracking number is 123456890, naming this as hoge

```
$ takuhai reg --sagawa -t 1234567890 -n hoge
```

## Rename hoge package to fuga package

```
$ takuhai rename -o hoge -n fuga
```

## Deregister the package specified above

```
$ takuhai dereg -t 1234567890
```

or

```
$ takuhai dereg -n hoge
```

## Get the states of all registered packages

```
$ takuhai show -a
```

## Get the state of a specific package

```
$ takuhai show -t 1234567890
```

or

```
$ takuhai show -n hoge
```

In the former example, the package information is specified by tracking number. If there are same numbers with different carriers, there should appear prompt to select which carrier you want to see the information.

The latter example, specifies package information by the name you selected before. Both examples show the latest status of the package.

## Get the timeline of a specific package

```
$ takuhai timeline -t 1234567890
```

or

```
$ takuhai timeline -n hoge
```

This shows the timeline of the package.

# Export package information and share them

This application can export some of package information you registered as yaml file, e.g. you can export packges named hoge, foo and bar to ./exp.yml by execute this:

```
$ takuhai export -n hoge -n foo -n bar -p exp.yml
```

This outputs the information of hoge, foo and bar packages to $(pwd)/exp.yml. Share this with people who are invovled in them like your coworkers.

# Import from yml file other person exported

If you receive exported yml file from other person, you can import it by running:

```
$ takuhai import -p exp.yml
```

By running this, the information of packages in exp.yml are imported to your environment.

# What's "takuhai"?

The word "takuhai" means delivery in Japanese.

# An article in Japanese is available at

[here](https://lapla.dev/posts/takuhai/)
