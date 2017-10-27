# sol2proto sample

1. Install [proto](https://github.com/golang/protobuf) and set up other tools

```
$ make setup
```

2. Build the contract code

```
$ make run name=NameService pkg=name_service
```

3. Build server binary

```
make server
```

4. Build client binary

```
make client
```