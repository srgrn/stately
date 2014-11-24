# Stately #
===
[![Build Status](https://travis-ci.org/srgrn/stately.svg?branch=master)](https://travis-ci.org/srgrn/stately)
[![Gobuild Download](http://gobuild.io/badge/github.com/srgrn/stately/downloads.svg)](http://gobuild.io/github.com/srgrn/stately)

Stately is a commandline multiplatform tool for handling projects that contains several repositories.

Rationale
------------
In Android development until recently in order to build a project that contained several android libs you had to bring all the code togather on the file system to allow the project.properties to contain the correct paths, in android studio (without gradle) you could even define them in the project iml/.idea files instead.
This may prove problmatic when you have each android lib as a different repo used by multiple applications.
Google have a wonderfull tool that supposed to help called google-repo however this tool is complex and works only on linux. 
stately was created to mitigate this issue and allowing collecting several git repositores at once.

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
both sources are from git however stately also supports svn urls

At the moment contains only two commands

#### Get
Run over a given configuration file and get the source into the target directory
stately at the moment knows two ways to get the configuration file:
* as a filepath
* as a url that contains the file as body (it doesn't strip html tags at the moment)

```
stately get .\example\projectdef.toml

stately get https://server.internal.com/projects/projectA.toml
```
(at the moment the stately code is ignoring certificates on ssl as the original work to support it were against a self signed certificate)
it can accept a target directory to put the source into instead of current directory by 
```
stately get -T Directory .\example\projectdef.toml
```

#### Freeze
Create configuration file from a directory.
usually you will already have project checked out atleast once. using the freeze command allows you to create a configuration file from a directory instead of writing it by hand.
it works on the current directory searching all top level directories for known source control systems (currently git and svn)
```
stately freeze projectdef.toml
```

License
---------------

This project is under the GNUv2 License. See the [LICENSE](https://github.com/srgrn/stately/blob/master/LICENSE) file for the full license text.

 