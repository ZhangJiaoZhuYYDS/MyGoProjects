package service

import "github.com/google/wire"

// ProviderSet is service providers.  向 wire 中注入 Service 代码
var ProviderSet = wire.NewSet(NewGreeterService, NewStudentService)
