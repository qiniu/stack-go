package main

import (
	"fmt"

	"github.com/qbox/qvm-base/components/logger"
	"github.com/qiniu/stack-go"
	"github.com/qiniu/stack-go/components/auth"
	"github.com/qiniu/stack-go/ecs"
)

func main() {
	s := stack.New(logger.StdLog, "https://api-qvm.qiniu.com", auth.NewCredential("ak", "sk"))

	name := "海外服务代理1"
	resp, err := s.ECS().ModifyInstanceAttribute(&ecs.ModifyInstanceAttribute{
		RegionId:     "cn-hongkong",
		InstanceName: &name,
		InstanceId:   "i-j6ci4wx78lgqewyxihx1",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.RequestID)
}
