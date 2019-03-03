package lghLog

import (
	"path/filepath"
	"os"
	"strings"
)

/**
    作者(Author): 林冠宏 / 指尖下的幽灵
    Created on : 2019/3/3
*/

const dirLevel = 4

func findConfigFile(configDir,fileName string) string {
	targetArr := []string{}
	targetArr = append(targetArr,configDir+"/"+fileName)
	targetArr = append(targetArr,"./"+configDir+"/" + fileName)
	for i:=1; i<=dirLevel; i++ {
		prefix := strings.Repeat("../",i)
		targetArr = append(targetArr,prefix+configDir+"/" + fileName)
	}
	found := false
	for _,item := range targetArr {
		if _, err := os.Stat(item); err == nil {
			fileName, _ = filepath.Abs(item)
			found = true
			break
		}
	}
	if !found {
		fileName, _ = filepath.Abs(fileName)
	}
	return fileName
}