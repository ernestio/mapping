# Mapping

master:  [![CircleCI](https://circleci.com/gh/ernestio/mapping/tree/master.svg?style=shield)](https://circleci.com/gh/ernestio/mapping/tree/master)  

Mapping is the internal representation of an ernest environment.
This library is used to interact with environments inside ernest, it has shortcuts to access many of the env's related endpoints. Some examples are:

You can create a new mapping by only providing environment name:
```go
env := mapping.New(c, "my_env")
```

And you can interact with it through:
```go
def := definition.New()
// Applies the given definition to current environment
env.Apply(def)
// Deletes the current environment
env.Delete()
// Imports to current environment the result for the specified env
env.Import(filters)
// Diff : gets a mapping for a diff between two environment builds
env.Diff("my_env_1", "my_env_0")
```

## Querying system

Mapping package introduces a querying system described [here](query)

## Installation

```
make deps
make install
```

## Running Tests

```
make deps
make test
```

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
