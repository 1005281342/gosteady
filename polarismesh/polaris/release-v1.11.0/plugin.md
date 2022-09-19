## 接口

在plugin/plugin.go中定义

```go
// Plugin 通用插件接口
type Plugin interface {
	Name() string
	Initialize(c *ConfigEntry) error
	Destroy() error
}
```



## 插件配置

plugin/plugin.go

```go
// Config 插件配置
type Config struct {
	CMDB                 ConfigEntry `yaml:"cmdb"`
	RateLimit            ConfigEntry `yaml:"ratelimit"`
	History              ConfigEntry `yaml:"history"`
	Statis               ConfigEntry `yaml:"statis"`
	DiscoverStatis       ConfigEntry `yaml:"discoverStatis"`
	ParsePassword        ConfigEntry `yaml:"parsePassword"`
	Auth                 ConfigEntry `yaml:"auth"`
	MeshResourceValidate ConfigEntry `yaml:"meshResourceValidate"`
	DiscoverEvent        ConfigEntry `yaml:"discoverEvent"`
}
```

新增插件需要在这里注册配置才能够正确解析到yaml配置



## 插件实现

plugin/auth/platform/platform.go

以 Auth 鉴权插件为例 

```go
// Auth 鉴权插件
type Auth struct {
	dbType       string
	dbSourceName string
	interval     time.Duration
	whiteList    string
	firstUpdate  bool
	lastMtime    time.Time
	ids          *sync.Map
}
```

实现插件定义的三个接口

```go
// Name 返回插件名字
func (a *Auth) Name() string {
	return PluginName
}

// Initialize 初始化鉴权插件
func (a *Auth) Initialize(conf *plugin.ConfigEntry) error {
	dbType, _ := conf.Option["dbType"].(string)
	dbAddr, _ := conf.Option["dbAddr"].(string)
	dbName, _ := conf.Option["dbName"].(string)
	interval, _ := conf.Option["interval"].(int)
	whiteList, _ := conf.Option["white-list"].(string)

	if dbType == "" || dbAddr == "" || dbName == "" {
		return fmt.Errorf("config Plugin %s missing database params", PluginName)
	}

	if interval == 0 {
		return fmt.Errorf("config Plugin %s has error interval: %d", PluginName, interval)
	}

	a.dbType = dbType
	a.dbSourceName = dbAddr + "/" + dbName
	a.interval = time.Duration(interval) * time.Second
	a.whiteList = whiteList
	a.ids = new(sync.Map)
	a.lastMtime = time.Unix(0, 0)
	a.firstUpdate = true

	if err := a.update(); err != nil {
		log.Errorf("[Plugin][%s] update err: %s", PluginName, err.Error())
		return err
	}
	a.firstUpdate = false

	go a.run()
	return nil
}

// Destroy 销毁插件
func (a *Auth) Destroy() error {
	return nil
}
```



## 插件注册

plugin/auth/platform/platform.go

```go
// init 初始化注册函数
func init() {
	plugin.RegisterPlugin(PluginName, &Auth{})
}
```



## 插件获取

plugin/auth.go

```go
// GetAuth 获取Auth插件
func GetAuth() Auth {
	c := &config.Auth

	plugin, exist := pluginSet[c.Name]
	if !exist {
		log.Error("[Plugin][Auth] not found", zap.String("name", c.Name))
		return nil
	}

	authOnce.Do(func() {
		if err := plugin.Initialize(c); err != nil {
			log.Errorf("plugin init err: %s", err.Error())
			os.Exit(-1)
		}
	})

	return plugin.(Auth)
}
```

