package context

import (
	"github.com/36625090/wechat/credential"
	"github.com/36625090/wechat/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
