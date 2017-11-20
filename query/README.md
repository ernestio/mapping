# Querying data-stores

Query object is intended to be an easy to use interface to retrieve data-stores information.

## Query structure

A basic example will retrieve all builds on the system

```go
var nc akira.Connector
var builds []builds.Builds
  err := query.New(nc, "builds.find").
    Run(builds)
```

As you can see first of all we will be creating a new query object, it will receive 2 arguments:

- An [akira.Connector](https://github.com/r3labs/akira/blob/master/nats.go#L15) in order to interact with the messaging system
- A subject "builds.find", which will be used as subject to create the request

 Finally, once the query is fully configured we can run it with `Run` method:
 ```go
 Run(builds)
```
in this case we will need to pass a collection builds, so it gets filled with the query results.


## Query filters

A query object has some [public methods](https://github.com/ernestio/mapping/blob/master/query/params.go) to modify the query parameters, this is helpful if we want to add one, or many filters to the serach:

The next example will be retrieving all builds matching the environment_id

```go
var nc akira.Connector
var builds []builds.Builds
  err := query.New(nc, "mapping.get.create").
    Run(builds)
```

Filters is the basic method to add filters to a query, however you'll find some shortcuts like:

### ID
filters the query results by id

```go
ID("my_env_id").
```

### Name

filters the query results by name field

```go
Name("my_env_id").
```
