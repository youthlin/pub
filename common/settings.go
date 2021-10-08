package common

import (
	"io"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/youthlin/z"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

var c = AppConfig{
	Logs: *z.DefaultConfig(),
	Web: WebConfig{
		Addr:  ":8080",
		Debug: true,
		AccessLog: []*z.Output{
			{Type: z.Console, File: lumberjack.Logger{Filename: z.Stdout}},
		},
		ErrorLog: []*z.Output{
			{Type: z.Console, File: lumberjack.Logger{Filename: z.Stderr}},
		},
	},
	LangPath: "conf/langs",
}

func initConfig() {
	// viper can not unmarshal yaml, so use yaml directly
	// > * cannot parse 'Logs.Level.Root' as int: strconv.ParseInt: parsing "warn": invalid syntax

	name := "conf/config.yaml"
	if n := os.Getenv("CONFIG_FILE"); n != "" {
		name = n
	}
	file, err := os.Open(name)
	if err != nil {
		panic(errors.Wrapf(err, "can not open config file: %v", name))
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(errors.Wrapf(err, "can not read config file: %v", name))
	}
	if err := yaml.Unmarshal(bytes, &c); err != nil {
		panic(errors.Wrapf(err, "failed to unmarshal config"))
	}
}

func Config() *AppConfig {
	return &c
}

type AppConfig struct {
	Logs     z.LogsConfig `json:"logs,omitempty" yaml:"logs"`
	Web      WebConfig    `json:"web,omitempty" yaml:"web"`
	LangPath string       `json:"lang_path" yaml:"lang_path"`
}

// WebConfig gin 配置
type WebConfig struct {
	Addr      string  `yaml:"addr" json:"addr,omitempty"`             // 监听地址
	Debug     bool    `yaml:"debug" json:"debug,omitempty"`           // 是否以 debug 模式运行
	AccessLog Outputs `yaml:"access_log" json:"access_log,omitempty"` // 访问日志文件
	ErrorLog  Outputs `yaml:"error_log" json:"error_log,omitempty"`   // 错误日志文件
}

type Outputs []*z.Output

func (os Outputs) MultiWriter() io.Writer {
	var out []io.Writer
	for _, o := range os {
		out = append(out, o)
	}
	return io.MultiWriter(out...)
}
