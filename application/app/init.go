package app

import (
	"btc-pool-appserver/application/config"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/lang"
	"btc-pool-appserver/application/library/log"
	"os"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	// 下载不下来啊
	// rotateLogs "github.com/lestrrat-go/file-rotatelogs"
)

var (
	isInit       int32
	defaultEnv   string
	dockerEnvMap map[string]bool
)

// Init app
func Init(typeStr string, workDir string) {
	if !atomic.CompareAndSwapInt32(&isInit, 0, 1) {
		log.Info("app has been init")
		return
	}

	//读取环境变量
	env := "dev" //os.Getenv("RUNTIME_ENVIRONMENT")
	if env == "" {
		env = defaultEnv
	}

	if v := dockerEnvMap[env]; v {
		log.InitLogger(os.Stdout)
	} else {
		// var (
		// 	baseLogPath   string
		// 	logFileMaxAge time.Duration
		// 	logRotateTime time.Duration
		// )

		// switch typeStr {
		// case "web":
		// 	baseLogPath, logFileMaxAge, logRotateTime = path.Join(workDir+"storage/logs", "web.log"), 7*time.Hour*24, time.Second*20
		// default:
		// 	panic("unknown type")
		// }

		// logWriter, err = rotateLogs.New(
		// 	baseLogPath+".%Y%m%d",
		// 	rotateLogs.WithMaxAge(logFileMaxAge),       //文件最大保存时间
		// 	ratateLogs.WithRotationTime(logRotateTime), //日志切割时间间隔
		// )

		// if err != nil {
		// 	log.Panicf("config local file system logger error. %+v", errors.WithStack(err))
		// }
		// log.InitLogger(logWriter)
	}

	log.Info("app run in mode " + env)

	// 加载配置，区分运行环境
	config.Load(workDir + "conf/env." + env + "/")

	// 初始化语言包
	lang.Load(workDir + "conf/lang/")

	// 初始化常用错误， err_msg依赖lang
	errs.Init()
}

// Exit app, 收尾工作
func Exit() {

}

// Start app
func Start() error {
	engine := gin.New()
	var err error

	// TODO:
	// config gin

	// 健康检查

	// 性能监控

	// 记录请求，过滤敏感信息

	// 重设cookie，向下请求接口用

	// 检验sign

	// 初始化路由

	if err = InitRouter(engine); err != nil {
		return err
	}

	// 初始化完成， 开始运行
	if err = engine.Run(); err != nil {
		return err
	}

	return nil
}

func setBtcpool(engine *gin.Engine) error {

	// TODO:
	return nil
}
