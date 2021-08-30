package main

import (
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

type OuterTask interface {
	GetName() string
	GetArgs() []Argument
}

type oTask struct {
	id   string
	Name string
	args []Argument
}

func NewOuterTask(name string, args string) OuterTask {
	return oTask{id: uuid.NewV4().String(), Name: name, args: createArgs(args)}
}

func (O oTask) GetId() string {
	return O.id
}

func (O oTask) GetName() string {
	return O.Name
}

func (O oTask) GetArgs() []Argument {
	return O.args
}

func createArgs(argsLine string) []Argument {
	argsLine = strings.ReplaceAll(argsLine, "  ", " ")
	argsLine = strings.TrimSpace(argsLine)
	args := strings.Split(argsLine, " ")
	var strArgs []Argument
	for i, arg := range args {
		strArgs = append(strArgs, NewOuterArg(i, arg))
	}

	return strArgs
}

type oArg struct {
	index int
	value string
}

func NewOuterArg(index int, value string) *oArg {
	return &oArg{index: index, value: value}
}

func (o oArg) GetName() string {
	return strconv.Itoa(o.index)
}

func (o oArg) GetValue() interface{} {
	return o.value
}
