# geordi

A key-value database written in Go.

## Available Commands

A `<value>` can be a string, integer, or float. These data types are explained
in the protocol section.

* **SET:** `SET "key" <value>`
* **GET:** `GET "key"`
* **INCR:** `INCR "key"` or by a specific value `INCR "key" 0.5`
* **HSET:** `HSET "hash-key" "key" <value>`
* **HGET:** `HGET "hash-key" "key"`

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
SET "name" "Dane"\r
+ "OK"\r
GET "an-error"\r
- "record not found"\r
GET "name"\r
+ "Dane"\r
```

### Client/Server Interaction

```
client> SET "name" "Dane"
server> + "OK"

client> GET "an-error"
server> - "record not found"

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
