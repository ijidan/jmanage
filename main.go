package main

import (
	"github.com/ijidan/jmanage/router"
	"golang.org/x/sync/errgroup"
	"log"
)

//入口函数
func main() {
	var g errgroup.Group
	//前台服务
	//g.Go(func() error {
	//	server:=router.WWWServer()
	//	return server.ListenAndServe()
	//})
	//后台服务
	g.Go(func() error {
		server:=router.AdminServer()
		return server.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}