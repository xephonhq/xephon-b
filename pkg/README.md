# Packages

This is the folder keep all the golang packages under the `github.com/xephonhq/xephon-b` namespace.
The structure is similar to https://github.com/grafana/grafana/tree/master/pkg
(And it's personal falvor I don't like too much folder in project root)

If certain package is stable and can be shared among other projects,
they may be moved to `github.com/xephon-contrib` namespace and adopt semantic version.
Before that, all the packages in `github.com/xephonhq/xephon-b` have no guarantee of API stability

External packages are managed by [glide](https://github.com/Masterminds/glide) in `vendor` folder

The outer `cmd` folder is used for better binary name when using `go get`
