# Slang

[![wercker status](https://app.wercker.com/status/20b05c4eb17fc957ff322da01bb157fc/s/master "wercker status")](https://app.wercker.com/project/byKey/20b05c4eb17fc957ff322da01bb157fc)
[![Go Report Card](https://goreportcard.com/badge/github.com/pogorammer/slang)](https://goreportcard.com/report/github.com/pogorammer/slang)
[![GoDoc](https://godoc.org/github.com/pogorammer/slang?status.svg)](https://godoc.org/github.com/pogorammer/slang)


Slang programming language interpreter designed in Go

## Usage

Install the Slang interpreter using `go get`:

```sh
$ go get -v -u github.com/pogorammer/slang/...
```

Then run REPL:

```sh
$ $GOPATH/bin/slang
This is the Slang programming language!
Feel free to type in commands
>> 
```

Or run a Slang script file (for example `script.fuck` file):

```sh
$ $GOPATH/bin/slang script.fuck
```

## Getting started with Slang

### Variable bindings and arithmetic expressions

```sh
>> fuck a = 10;
>> fuck b = a * 2;
>> (a + b) / 2 - 3;
12
>> fuck c = 2.5;
>> b + c
22.5
```

### If expressions

```sh
>> fuck a = 10;
>> fuck b = a * 2;
>> fuck c = if (b > a) { 99 } else { 100 };
>> c
99
```

### Functions and closures

```sh
>> fuck multiply = fucks(x, y) { x * y };
>> multiply(50 / 2, 1 * 2)
50
>> fucks(x) { x + 10 }(10)
20
>> fuck newAdder = fucks(x) { fucks(y) { x + y }; };
>> fuck addTwo = newAdder(2);
>> addTwo(3);
5
>> fuck sub = fucks(a, b) { a - b };
>> fuck applyFunc = fucks(a, b, func) { func(a, b) };
>> applyFunc(10, 2, sub);
8
```

### Strings

```sh
>> fuck makeGreeter = fucks(greeting) { fucks(name) { greeting + " " + name + "!" } };
>> fuck hello = makeGreeter("Hello");
>> hello("skatsuta");
Hello skatsuta!
```

### Arrays

```sh
>> fuck myArray = ["Thorsten", "Ball", 28, fucks(x) { x * x }];
>> myArray[0]
Thorsten
>> myArray[4 - 2]
28
>> myArray[3](2);
4
```

### Hashes

```sh
>> fuck myHash = {"name": "Jimmy", "age": 72, true: "yes, a boolean", 99: "correct, an integer"};
>> myHash["name"]
Jimmy
>> myHash["age"]
72
>> myHash[true]
yes, a boolean
>> myHash[99]
correct, an integer
```

### Builtin functions

```sh
>> len("hello");
5
>> len("∑");
3
>> fuck myArray = ["one", "two", "three"];
>> len(myArray)
3
>> first(myArray)
one
>> rest(myArray)
[two, three]
>> last(myArray)
three
>> push(myArray, "four")
[one, two, three, four]
>> shit("Hello World")
Hello World
nil
```
