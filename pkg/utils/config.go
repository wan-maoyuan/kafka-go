package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		GrpcAddr string `mapstructure:"grpc_addr" yaml:"grpc_addr"`
	} `mapstructure:"server" yaml:"server"`
	Log struct {
		FilePath     string  `mapstructure:"file_path" yaml:"file_path"`
		MaxFileSize  float32 `mapstructure:"max_file_size" yaml:"max_file_size"`
		MaxFileCount uint32  `mapstructure:"max_file_count" yaml:"max_file_count"`
	} `mapstructure:"log" yaml:"log"`
}

var C Config

func InitializeConfig() error {
	v := viper.New()
	v.SetConfigName("service")
	v.AddConfigPath(".")
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("initialize config error: %v", err)
	}

	if err := v.Unmarshal(&C); err != nil {
		return fmt.Errorf("unmarshal config error: %v", err)
	}

	C.show()
	return nil
}

func (c Config) show() {
	fmt.Println(fmt.Sprintf(`
-----------------------------------------------------------------------------------------
%v
-----------------------------------------------------------------------------------------
	`, c))
}

func (c Config) String() string {
	if b, err := yaml.Marshal(c); err == nil {
		return string(b)
	}
	return ""
}
