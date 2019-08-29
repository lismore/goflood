[![Build Status](https://travis-ci.com/lismore/goflood.svg?branch=master)](https://travis-ci.com/lismore/goflood)

# GoFlood
A connection flood attack application written in Go
  - Enables security tester to perform a Denial of Service against a target of evaluation 
  
## New Features!

  - Initial Proof of Concept

### Tech

GoFlood is implemented in GO programming language

### Installation

### Step 1: Build

From the root directory open a terminal and type the following command to generate your own binary

``` shell
go build goflood.go YouExeName.exe 
```
or
```
go build goflood.go
```
This command will build a binary of the same name goflood.exe

### Step 2: Run
From command prompt/powershell/terminal execute the following binary to launch a connection flood attack against your target of evaluation
``` CMD
.\goflood.exe -targetIP="127.0.0.1" -targetPort=4444 -numOfRequests=10000
```
### Development

Want to contribute? Great!

Fork, submit a pull request

License
----

MIT
