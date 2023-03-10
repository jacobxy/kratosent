package main

import (
	"flag"
	olog "log"
	"os"
	"time"

	"kratosent/internal/conf"
	"kratosent/internal/midconf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mosn.io/holmes"

	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	polaris "github.com/go-kratos/kratos/contrib/polaris/v2"
	polarisV2 "github.com/polarismesh/polaris-go"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "kratosent"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server,
	data *conf.Data,
	register *conf.Register,
	midcf *conf.MidConfig,
) *kratos.App {
	// priviopolaris.NewProviderAPI()

	sdk, err := polarisV2.NewSDKContextByAddress(register.Polaris.Addrs...)
	if err != nil {
		log.Fatal(err)
	}
	p := polaris.New(sdk,
		polaris.WithNamespace(register.Polaris.Namespace),
	)
	midconf.NewMidConfig(midcf)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),

		kratos.Registrar(p.Registry(
			polaris.WithRegistryTTL(40),
		)),
	)
}

func main() {
	flag.Parse()

	// logger := log.With(log.NewStdLogger(os.Stdout),
	// 	"ts", log.DefaultTimestamp,
	// 	"caller", log.DefaultCaller,
	// 	"service.id", id,
	// 	"service.name", Name,
	// 	"service.version", Version,
	// 	"trace.id", tracing.TraceID(),
	// 	"span.id", tracing.SpanID(),
	// )
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	if bc.Log == nil {
		olog.Fatalln(bc.Log, flagconf)
		return
	}

	logger := NewZapLoger(bc.Log)

	// mHolmes := NewHolmes()
	// mHolmes.Start()
	// defer mHolmes.Stop()

	app, cleanup, err := wireApp(
		bc.Server,
		bc.Data,
		bc.Register,
		bc.MidConfig,
		bc.Middleware,
		logger,
	)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func NewZapLoger(logConf *conf.Log) log.Logger {
	lumberJackLogger := lumberjack.Logger{
		Filename:   logConf.File,
		MaxSize:    int(logConf.MaxSize),
		MaxBackups: int(logConf.MaxBackup),
		MaxAge:     int(logConf.MaxAge),
		Compress:   logConf.Compress,
	}
	writeSyncer := zapcore.AddSync(&lumberJackLogger)
	core := zapcore.NewCore(getEncoder(), writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return kzap.NewLogger(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func NewHolmes() *holmes.Holmes {
	h, _ := holmes.New(
		holmes.WithCollectInterval("5s"),
		holmes.WithDumpPath("./holmes"),
		holmes.WithTextDump(),

		holmes.WithCPUDump(10, 25, 80, time.Minute),
		holmes.WithMemDump(30, 25, 80, time.Minute),
		holmes.WithGCHeapDump(10, 20, 40, time.Minute),
		holmes.WithGoroutineDump(500, 25, 20000, 0, time.Minute),
		// holmes.WithCGroup(true),
	)

	h.EnableCPUDump().
		EnableGoroutineDump().
		EnableMemDump().
		EnableGCHeapDump()

	return h
}
