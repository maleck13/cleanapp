## Clean App

*Note this is a very early version that still needs lots of work*

tool for scaffolding a new clean app in go, the idea is that is will create a structure that follows or encourages using onion or clean architecture.


### Usage example

```
cd $GOPATH/src/github.com
mkdir maleck13 && cd maleck13
git clone git@github.com:maleck13/template_server.git
cd $GOPATH/src/github.com && mkdir myapp
cd myapp
cleanapp -n=myapp -t=$GOPATH/src/github.com/maleck13/template_server
````
