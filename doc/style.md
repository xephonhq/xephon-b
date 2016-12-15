# Coding Style

## Git

- add `[type | action]` prefix to commit message ie: `[doc] update db list`, `[fix] wrong configuration for cassandra` 
- use an editor that support [editorconfig](http://editorconfig.org/)
- use `git pull --rebase` to avoid unnecessary merge commit when update local repository
- DO NOT use `git push -f` unless you have a good reason, ie: on another branch that has conflict with master
- stick to master unless you have some features that diverge from current design, `fix/xx`, `feature/xxx` is overkill
- find a good image from http://lgtm.in/ when you close an issue, merge a PR, publish a new release.

## Test

- No need to test if everything is unstable
- If you don't have time, ie: Algo test tommorrow, then don't write tests!
- Coverage is important but 100% coverage always ruin a good day, so have a nice day!

## Language Specific

### Bash

- write the `#!/usr/bin/env bash` line
- DO NOT use zsh

### Java

- use Java 8

### Golang

- use Golang 1.7
- use `vendor` with [glide](https://github.com/Masterminds/glide) for dependency management
- use specific version for you dependency, do NOT use latest

#### Interface 

- TODO: each package put interface in its package name file, ie: `generator/generator.go`
- check implementation match interface in package name test file, ie : `generator/generator_test.go`, 
using `var _ InterafaceName = (*ImplementationName)(nil)`

### JavaScript

- TBD

## Golang

### Error handling

- deal with error first, avoid nesting
- panic when the developer made some stupid mistake
- wrap the error using `pkg/errors`
- use type conversion to extract information from specific error
- don't parse/check error string unless in test, `Error()` is for human

Bad

````go
if err := doA(); err == nil {
      if err = doB(); err == nil {
          fmt.Println("a and b success")
      }else{
          return err
      }
}else{
    return err
}
````

Good 

````go
if err := doA(); err != nil {
    return err
}
if err = doB(); err != nil {
    return err
}
fmt.Println("a and b success")
````

### Package organization

- use as namespace, nest when there is need
- split one package into server files by their functionality

### Test

- use `t.Parallel()` whenever you can
- test race problem

### Log

- separate logger for each package

### Doc 

- NO need to add doc if it is WIP and/or very unstable
- remove outdated regularly