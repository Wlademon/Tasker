package main

type Task interface {
	GetId() string
	GetArgs() []Argument
}

type Argument interface {
	GetName() string
	GetValue() interface{}
}
