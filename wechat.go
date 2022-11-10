package wechat

import (
	"github.com/36625090/wechat/cache"
	"github.com/36625090/wechat/miniprogram"
	miniConfig "github.com/36625090/wechat/miniprogram/config"
	"github.com/36625090/wechat/officialaccount"
	offConfig "github.com/36625090/wechat/officialaccount/config"
	"github.com/36625090/wechat/openplatform"
	openConfig "github.com/36625090/wechat/openplatform/config"
	"github.com/36625090/wechat/work"
	workConfig "github.com/36625090/wechat/work/config"
)

func init() {
}

// Wechat WeChat struct
type Wechat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

// SetCache 设置cache
func (wc *Wechat) SetCache(cache cache.Cache) {
	wc.cache = cache
}

// GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(cfg)
}

// GetWork 获取企业微信的实例
func (wc *Wechat) GetWork(cfg *workConfig.Config) *work.Work {
	return work.NewWork(cfg)
}
