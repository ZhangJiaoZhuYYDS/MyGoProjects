//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func Init() BroadCast {
	wire.Build(NewBroadCast, NewChannel, NewMessage)
	return BroadCast{}
}
