# Stately #
===
[![Build Status](https://travis-ci.org/srgrn/stately.svg?branch=master)](https://travis-ci.org/srgrn/stately)
[![Gobuild Download](http://gobuild.io/badge/github.com/srgrn/stately/downloads.svg)](http://gobuild.io/github.com/srgrn/stately)

Stately is a commandline multiplatform tool for handling projects that contains several repositories.

Installation
------------

You may build ah from sources or just download proper binary from
[gobuild](http://gobuild.io/github.com/srgrn/stately).

Or

To install it from sources, just do following:

```
go get github.com/srgrn/stately
```

Usage
-----------
Stately uses [toml](https://github.com/toml-lang/toml) for its configuration files for example take the following configuration file

```
name = "test"
[[sources]]
url = "git@github.com:srgrn/google-drive-trash-cleaner.git"
target = "trash"
[[sources]]
url = "git@github.com:srgrn/stately.git"
target = "stately"
branch = "master"
```

The configuration file above defines a project called "test",
with two sources on in a directory called trash the other in a directory called stately.


At the moment contains only two commands

#### Get
Run over a given configuration file and get the source into the target directory

#### Freeze
Create configuration file from a directory


 