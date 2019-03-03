package lghLog

import (
	"testing"
	"fmt"
	"github.com/lghLog/log4go/l4g"
	"math/rand"
)

/**
    作者(Author): 林冠宏 / 指尖下的幽灵
    Created on : 2019/3/3
*/
func Test_SimpleInit(t *testing.T) {
	DefaultInit()
	defer l4g.Close()
	l4g.Info("name %s","simple"+fmt.Sprintf("%d",rand.Intn(2000)))
}

func Test_InitLogger(t *testing.T) {
	defer l4g.Close()
	Init("config","example.xml")
	l4g.Info("name %s","lgh"+fmt.Sprintf("%d",rand.Intn(2000)))
}

func Test_FindConfigFile(t *testing.T) {
	fmt.Println(findConfigFile("config","config.json"))
}

