package config

import (
	"aladinfun.com/ZonePet/ZonePetServer/common/cfgmgr"
	frw "aladinfun.com/ZonePet/ZonePetServer/framework"
	"errors"
	"github.com/vaughan0/go-ini"
)

type ServerConfig struct {
	SkeCfg   *frw.SkeletonConfig
	LogPath  string
	LogLevel string
	Platform string
	HttpPort string
	CacheDir string
}

var (
	SrvCfg ServerConfig // 基础配置
)

func LoadConfig(filename string) (cfg *ServerConfig, err error) {
	f, err := ini.LoadFile(filename)
	if err != nil {
		return
	}
	// 加载服务配置
	if err = loadSrvConfig(f); err != nil {
		return
	}
	// 加载框架配置
	SrvCfg.SkeCfg, err = loadFrwConfig(f)
	if err != nil {
		return
	}
	cfg = &SrvCfg
	return
}

func Close() {
	cfgmgr.Close()
}

func loadSrvConfig(f ini.File) (err error) {
	ok := true
	section := "Server"
	SrvCfg.LogLevel, ok = f.Get(section, "log_level")
	if !ok {
		err = errors.New("get load log_level config failed")
		return
	}
	SrvCfg.LogPath, ok = f.Get(section, "log_path")
	if !ok {
		err = errors.New("get load log_path config failed")
		return
	}
	SrvCfg.Platform, ok = f.Get(section, "platform")
	if !ok {
		err = errors.New("get load platform config failed")
		return
	}
	SrvCfg.HttpPort, ok = f.Get(section, "http_port")
	if !ok {
		err = errors.New("get load http_port config failed")
		return
	}
	SrvCfg.CacheDir, ok = f.Get(section, "cache_dir")
	if !ok {
		err = errors.New("get load cache_dir config failed")
		return
	}
	return
}

func loadFrwConfig(f ini.File) (cfg *frw.SkeletonConfig, err error) {
	cfg = &frw.SkeletonConfig{}
	ok := true

	section := "Framework"
	cfg.NatsAddrs, ok = f.Get(section, "nats_addrs")
	if !ok {
		err = errors.New("get load nats_addrs config failed")
		return
	}
	cfg.NatsTimeout, ok = f.GetInt(section, "nats_timeout_msec")
	if !ok {
		err = errors.New("get load nats_timeout_msec config failed")
		return
	}
	cfg.EtcdAddrs, ok = f.Get(section, "etcd_addrs")
	if !ok {
		err = errors.New("get load etcd_addrs config failed")
		return
	}
	cfg.ListenPort, ok = f.GetInt(section, "listen_port")
	if !ok {
		err = errors.New("get load listen_port config failed")
		return
	}
	cfg.MaxRoutineNum, ok = f.GetInt(section, "max_routine_num")
	if !ok {
		err = errors.New("get load max_routine_num config failed")
		return
	}
	return
}
