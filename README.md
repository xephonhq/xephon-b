# xephon-b

<h1 align="center">
	<br>
	<img width="400" src="https://raw.githubusercontent.com/at15/artwork/master/logo/xephonhq/xephon-b.png" alt="xephon-b">
	<br>
	<br>
	<br>
</h1>

[![Build Status](https://travis-ci.org/xephonhq/xephon-b.svg?branch=feature%2Fdata-generation)](https://travis-ci.org/xephonhq/xephon-b)
[![Join the chat at https://gitter.im/xephonhq/xephon-b](https://badges.gitter.im/xephonhq/xephon-b.svg)](https://gitter.im/xephonhq/xephon-b?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/xephonhq/xephon-b?status.svg)](https://godoc.org/github.com/xephonhq/xephon-b)

A time series database benchmark tool and benchmark result sharing platform

- [Documentation](doc)
- [Slide: Introduce Xephon-B](http://www.slideshare.net/ssuser7e134a/intoduce-xephonb)

![system design](doc/system-design.png)

NOTE: this repository is outdated, the development is now in a monolithic repository [Xephon-K](https://github.com/xephonhq/xephon-k)

## FAQ

Q: How can I get involved?

A: Thanks for your interest, but this project is not actively maintained, most of its code has been migrated
to [Xephon-K](https://github.com/xephonhq/xephon-k) as its load tester.

Q: Is this benchmark suite for time series database only?

A: Yes, its load is quite different from both RDMS and NoSQL store. That's why we create a new tool.
However its components are quite flexible, you can fork it and make it a more general database benchmark tool.

Q: Why call it Xephon-B?

A: B is for benchmark and Xephon comes from the animation [RahXephon](https://en.wikipedia.org/wiki/RahXephon).

## Related Projects

- [awesome-time-series-database](https://github.com/xephonhq/awesome-time-series-database)
- [Xephon-K A time series database using Cassandra as backend, modeled after KairosDB](https://github.com/xephonhq/xephon-k)

## License

MIT

## Authors

- [Pinglei Guo](https://at15.github.io) [@at15](https://github.com/at15), [linkedin](https://www.linkedin.com/in/at1510086), [twitter](https://twitter.com/at1510086)
- [Zheyuan Chen](http://czheo.github.io/) [@czheo](https://github.com/czheo), [linkedin](https://www.linkedin.com/in/zheyuan-chen), [twitter](https://twitter.com/czheo)
