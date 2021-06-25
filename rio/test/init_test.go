package test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/qiniu/stack-go/components/auth"
	"github.com/qiniu/stack-go/components/log"
	"github.com/qiniu/stack-go/rio"
)

var (
	testConfigPath string

	testStack  *rio.Stack
	testZoneID = "nova"

	testDiskID = "a3954481-6f55-4928-a4b3-077ec56d0211"

	l log.Logger
)

func TestMain(m *testing.M) {
	srcPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	l = log.NewSimpleLog()
	testConfigPath = strings.Split(srcPath, "rio")[0] + "rio/config.test.json"
	endpoint, credential := setup(testConfigPath)

	testStack = rio.New(log.NewSimpleLog(), endpoint, credential)

	code := m.Run()

	os.Exit(code)
}

func setup(filePath string) (string, *auth.Credential) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("无法读取配置文件 %v: %v", filePath, err))
	}

	defer f.Close()

	var (
		conf struct {
			Endpoint  string `json:"endpoint"`
			AccessKey string `json:"access_key"`
			SecretKey string `json:"secret_key"`
		}
	)

	err = json.NewDecoder(f).Decode(&conf)

	return conf.Endpoint, auth.NewCredential(conf.AccessKey, conf.SecretKey)
}
