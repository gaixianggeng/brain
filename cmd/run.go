package main

import (
	"doraemon/api"
	"doraemon/conf"
	"doraemon/internal/engine"
	"doraemon/internal/index"
	"flag"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	action int64
)

const (
	START_SERVER = 1
	START_INDEX  = 2
)

func run() {
	// TODO: 命令行启动参数
	confPath := "./conf/conf.toml"
	c, err := conf.ReadConf(confPath)
	if err != nil {
		log.Fatal(err)
	}

	meta, err := engine.ParseMeta(c)
	if err != nil {
		panic(err)
	}
	if meta == nil {
		panic("meta is nil")
	}

	// 定时同步meta数据
	ticker := time.NewTicker(time.Second * 1)
	go meta.SyncByTicker(ticker)

	if c.Version != meta.Version {
		panic(fmt.Sprintf("version not match, %s != %s", c.Version, meta.Version))
	}

	start(c, meta)

	// close
	func() {
		// 最后同步元数据至文件
		log.Info("close")
		meta.SyncMeta()
		log.Info("close")
		ticker.Stop()
		log.Info("close")
	}()
}

func start(c *conf.Config, meta *engine.Meta) {
	if action == START_SERVER {
		log.Debugf("start server...")
		api.Start(meta, c)
	} else if action == START_INDEX {
		log.Debugf("start")
		index.Run(meta, c)
		log.Debug("end")
	}
}

// 获取flag参数
func flagInit() {
	flag.Int64Var(&action, "flag", 1, "start flag 1:serv 2:index")
	flag.Parse()
}

func init() {
	flagInit()
}
