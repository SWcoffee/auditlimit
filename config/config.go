package config

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	PORT           = 8080
	PlusModels     = garray.NewStrArrayFrom([]string{"gpt-4", "gpt-4o", "gpt-4-browsing", "gpt-4-plugins", "gpt-4-mobile", "gpt-4-code-interpreter", "gpt-4-dalle", "gpt-4-gizmo", "gpt-4-magic-create"})
	ForbiddenWords = []string{}    // 禁止词
	LIMIT          = 40            // 限制次数
	PER            = time.Hour * 3 // 限制时间
	OAIKEY         = ""            // OAIKEY
	OAIKEYLOG      = ""            // OAIKEYLOG 隐藏
	// MODERATION     = "https://api.openai.com/v1/moderations" // OPENAI Moderation 检测
	MODERATION = "https://gateway.ai.cloudflare.com/v1/a8cace244ffbc233655fefeaca37d515/xyhelper/openai/moderations"
)

func init() {
	ctx := gctx.GetInitCtx()

	// 日志配置
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		m := map[string]interface{}{
			"stdout": false,
			"path":  "app/logs", // 此处必须重新设置，才可以实现db的log写入文件
		}
		in.Logger.SetConfigWithMap(m)
		in.Next(ctx)
	})


	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)
	limit := g.Cfg().MustGetWithEnv(ctx, "LIMIT").Int()
	if limit > 0 {
		LIMIT = limit
	}
	g.Log().Info(ctx, "LIMIT:", LIMIT)
	per := g.Cfg().MustGetWithEnv(ctx, "PER").Duration()
	if per > 0 {
		PER = per
	}
	g.Log().Info(ctx, "PER:", PER)
	oaikey := g.Cfg().MustGetWithEnv(ctx, "OAIKEY").String()
	// oaikey 不为空
	if oaikey != "" {
		OAIKEY = oaikey
		// 日志隐藏 oaikey，有 * 代表有值
		OAIKEYLOG = "******"
	}
	g.Log().Info(ctx, "OAIKEY:", OAIKEYLOG)
	moderation := g.Cfg().MustGetWithEnv(ctx, "MODERATION").String()
	if moderation != "" {
		MODERATION = moderation
	}
	g.Log().Info(ctx, "MODERATION:", MODERATION)
}
