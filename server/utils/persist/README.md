casbinʹ��redis��Ϊ���棬ʹ������
```go
		// ...
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(3600)
		syncedCachedEnforcer.EnableCache(true)
		syncedCachedEnforcer.SetCache(persist.NewCasbinCache())  // ʹ��redis�Զ��建��
```