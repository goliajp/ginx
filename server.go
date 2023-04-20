package ginx

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	DefaultAddr              = "0.0.0.0" // DefaultAddr 默认监听地址
	DefaultPort              = 60002     // DefaultPort 默认监听端口
	DefaultReadTimeout       = "120s"    // DefaultReadTimeout 默认读取超时
	DefaultReadHeaderTimeout = "120s"    // DefaultReadHeaderTimeout 默认读取header超时
	DefaultWriteTimeout      = "120s"    // DefaultWriteTimeout 默认写入超时
	DefaultIdleTimeout       = "120s"    // DefaultIdleTimeout 默认处理超时
)

type HttpConfig struct {
	Addr              string        `json:"addr" mapstructure:"addr"`
	Port              int           `json:"port" mapstructure:"port"`
	ReadTimeout       time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	ReadHeaderTimeout time.Duration `json:"read_header_timeout" mapstructure:"read_header_timeout"`
	WriteTimeout      time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
	IdleTimeout       time.Duration `json:"idle_timeout" mapstructure:"idle_timeout"`
}

func NewServer(handler http.Handler, config *HttpConfig, configFilePaths ...string) *http.Server {
	var cfg HttpConfig
	if config != nil {
		cfg = *config
	} else {
		v := viper.New()
		v.SetDefault("addr", DefaultAddr)
		v.SetDefault("port", DefaultPort)
		v.SetDefault("read_timeout", DefaultReadTimeout)
		v.SetDefault("read_header_timeout", DefaultReadHeaderTimeout)
		v.SetDefault("write_timeout", DefaultWriteTimeout)
		v.SetDefault("idle_timeout", DefaultIdleTimeout)
		v.SetEnvPrefix("HTTP")
		v.AutomaticEnv()
		v.SetConfigName("http.config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		for _, path := range configFilePaths {
			v.AddConfigPath(path)
		}
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&cfg)
	}
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		Handler:           handler,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}
}

func GracefulServe(listenSignal chan os.Signal, server *http.Server) error {
	go func() {
		log.Printf("http served on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorln(err)
		}
	}()
	signal.Notify(listenSignal, os.Interrupt)
	<-listenSignal
	log.Info("正在关闭服务器")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
