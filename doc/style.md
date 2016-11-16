# Coding Style

## Git

- add `[type | action]` prefix to commit message ie: `[doc] update db list`, `[fix] wrong configuration for cassandra` 
- use an editor that support [editorconfig](http://editorconfig.org/)
- use `git pull --rebase` to avoid unnecessary merge commit when update local repository
- DO NOT use `git push -f` unless you have a good reason, ie: on another branch that has conflict with master
- stick to master unless you have some features that diverge from current design, `fix/xx`, `feature/xxx` is overkill
- find a good image from http://lgtm.in/ when you close an issue, merge a PR, publish a new release.

## Test

- If you do have time, ie: Algo test tommorrow, then don't write tests!
- Coverage is important but 100% coverage always ruin a good day, so have a nice day!

## Language Specific

### Bash

- write the `!#/usr/bin/bash` line
- DO NOT use zsh

### Java

- use Java 8

### Golang

- use Golang 1.7
- use `vendor` with [glide](https://github.com/Masterminds/glide) for dependency management
- use specific version for you dependency


#### Interface 

- each package put interface in its package name file, ie: `generator/generator.go`
- check implementation match interface in package name test file, ie : `generator/generator_test.go`, 
using `var _ InterafaceName = (*ImplementationName)(nil)`

### JavaScript

- do what ever you want

### PHP

- it's the best language
