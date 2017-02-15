# jsoncmp
A simple command line tool for comparing two JSON files

## Install
There are two options for installing ```jsoncmp```:
### 1) Use ```go get``` to retrieve the project into your GOPATH and manually build/run it
```
go get github.com/opreaalex/jsoncmp

cd $GOPATH/src/github.com/opreaalex/jsoncmp

go build

./jsoncmp <your_first_json_file> <your_second_json_file>
```
### 2) Download the binary, save it to any directory and add it to your PATH (Linux x64)
```
wget -P $HOME/.local/bin https://github.com/opreaalex/jsoncmp/releases/download/0.1.0/jsoncmp

export PATH=$PATH:/$HOME/.local/bin/

chmod u+x $HOME/.local/bin/jsoncmp

jsoncmp <your_first_json_file> <your_second_json_file>
```
