package main

import (
	"reserve/conf"
	"reserve/server"
)

// @title go项目模板
// @version 1.0
func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	s := server.NewServer()
	if err := s.Init(conf.Conf); err != nil {
		panic(err)
	}
	err := s.Run()
	if err != nil {
		panic(err)
	}
}
