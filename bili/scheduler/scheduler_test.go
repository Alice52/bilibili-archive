package scheduler

import (
	"github.com/alice52/archive/bili/service"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/common/global"
	initialize "github.com/alice52/archive/common/init"
	"github.com/wordpress-plus/kit-logger/viperx"
	"github.com/wordpress-plus/kit-logger/zapx"
	"testing"
)

func init() {
	global.VIPER = viperx.Viper(&global.CONFIG, "../config-local.yaml")
	global.LOG = zapx.Zap(global.CONFIG.Zap)

	// init db and do migration
	global.DB = initialize.GormMysql()
	if global.DB.Error != nil {
		panic(global.DB.Error)
	} else {
		dal.SetDefault(global.DB)
	}
}

func TestUserUpperTagService(t *testing.T) {
	err := service.UserUpperTagService.SyncUpperTags()
	if err != nil {
		panic(err)
	}
}

func TestSyncUppersService(t *testing.T) {
	err := service.UserUpperService.SyncUppers()
	if err != nil {
		panic(err)
	}
}
