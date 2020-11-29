package gcasbin

import (
  "github.com/casbin/casbin/v2"
  "github.com/casbin/casbin/v2/model"
  gormadapter "github.com/casbin/gorm-adapter/v3"
  config "go-element-admin-api/configs"
  "strconv"
)

var Enforcer *casbin.SyncedEnforcer

var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func init()  {
  conn := config.DatabaseConfig.Username + ":" + config.DatabaseConfig.Password + "@tcp(" + config.DatabaseConfig.Host + ":" + strconv.Itoa(config.DatabaseConfig.Port) + ")/" + config.DatabaseConfig.Database
  Apter, err := gormadapter.NewAdapter(config.DatabaseConfig.Dbtype, conn, true)
  if err != nil {
    panic(err)
  }

  m, err := model.NewModelFromString(text)
  if err != nil {
    panic(err)
  }

  Enforcer, err = casbin.NewSyncedEnforcer(m, Apter)
  if err != nil {
    panic(err)
  }

  // 开启权限认证日志
  Enforcer.EnableLog(true)

  // 加载策略
  err = Enforcer.LoadPolicy()
  if err != nil {
    panic(err)
  }
}
