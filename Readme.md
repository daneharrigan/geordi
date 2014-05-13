# geordi

A key-value database written in Go.

## Protocol

Geordi's wire protocol is meant to be simple and easy read and parse.

* New-lines, tabs and space characters are considered whitespace and are ignored
* Command names are uppercased
* Strings are always double quoted
* Integers are never quoted
* Floats are not required to begin with 0 when the value is less than 1
* Operations end with a `\r` character
* Successful responses begin with a `+` character
* Error responses begin with a `-` character

```
client> SET "name" "Dane"
server> + "OK"

client> GET "name"
server> + "Dane"

client> SET "age" 28
server> + "OK"

client> GET "age"
server> + 28
```

## Developing

```console
$ go get github.com/daneharrigan/geordi
$ cd $GOPATH/src/github.com/daneharrigan/geordi
$ ./build.sh
```

The `build.sh` script will compile the `geordi-server`, `geordi-cli` and create
your PEM and KEY files.

Run the server with:

```console
$ bin/geordi-server -pem service.pem -key service.key
```

Connect to the server with the CLI:

```console
$ bin/geordi-cli
```
