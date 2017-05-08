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

This will create a structure and scoffold

```
├── cmd
│   └── server  # this is were main lives and where the binary will be built
├── install # provides resources for installing the app (wip)
│   ├── kubernetes
│   └── openshift
└── pkg # main packages dir
    ├── inmemdb # db interface framework/driver
    ├── shop # core domain and buisiness logic
    │   ├── dispatch
    │   ├── orders
    │   │   ├── README.md
    │   │   ├── history.go #a use case order history
    │   │   ├── history_test.go
    │   │   ├── interfaces.go
    │   │   ├── place.go # a use case (place order)
    │   │   └── place_test.go
    │   └── types.go # core domain model / entities
    └── web # web interface framework/driver

```    