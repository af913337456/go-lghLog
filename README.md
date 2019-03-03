#### DES

```golang

func Test_SimpleInit(t *testing.T) {
	DefaultInit()
	defer l4g.Close()
	l4g.Info("name %s","simple"+fmt.Sprintf("%d",rand.Intn(2000)))
}

```

`DefaultInit` will help you auto create ConfigFile dir、config file、log files dir

`DefaultInit` 将会默认帮助创建日志的配置文件和日志的存放目录。做到一键处理。

`Init(configDir,fileName string)` create what you define.

`Init(configDir,fileName string)` 自定义创建
