casbin使用redis作为缓存，使用例：
```go
		// ...
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(3600)
		syncedCachedEnforcer.EnableCache(true)
		syncedCachedEnforcer.SetCache(persist.NewCasbinCache())  // 使用redis自定义缓存
```