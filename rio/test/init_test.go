package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"

	"github.com/qiniu/stack-go/components/cookieauth"
	"github.com/qiniu/stack-go/components/log"
	"github.com/qiniu/stack-go/rio"
)

var (
	testConfigPath string

	testStack    *rio.Stack
	testZoneID   = "nova"
	testCIDR     = "064c33f6-8736-4c67-970f-365740d57aaa"
	testEIPID    = "10db31d9-db84-4f0e-b4b6-fd552b3650c3"
	testDiskID   = "a3954481-6f55-4928-a4b3-077ec56d0211"
	testServerID = "6bcfcb3b-25fa-4b6e-aac4-cd02f68ba65d"
	l            log.Logger
)

func TestMain(m *testing.M) {
	srcPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	l = log.NewSimpleLog()
	testConfigPath = strings.Split(srcPath, "rio")[0] + "rio/config.test.yaml"
	cfg := setup(testConfigPath)

	testStack = rio.New(log.NewSimpleLog(), cfg)

	code := m.Run()

	os.Exit(code)
}

func setup(filePath string) *cookieauth.Config {
	viper.SetConfigType("yaml")

	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("无法读取配置文件 %v: %v", filePath, err))
	}

	defer f.Close()

	if err = viper.ReadConfig(f); err != nil {
		panic(fmt.Sprintf("读取配置文件错误 %v: %v", filePath, err))
	}

	var cfg *cookieauth.Config
	if err = viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Sprintf("解析配置文件失败 %v: %v", filePath, err))
	}

	return cfg
}
