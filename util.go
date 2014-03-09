package reveltang

import (
	"fmt"

	"github.com/itang/gotang"
	"github.com/revel/revel"
)

// 获取应用配置， 如果未配置，则panic
func ForceGetConfig(key string) string {
	value, exists := revel.Config.String(key)
	gotang.Assert(exists, fmt.Sprintf("Missing revel app config for key: %s", key))

	return value
}
