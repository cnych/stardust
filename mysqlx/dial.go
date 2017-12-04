package mysqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	cacher *xorm.LRUCacher
)

type Address struct {
	Addr     string `json:"addr" toml:"addr"`
	Database string `json:"db" toml:"db"`
	//Timeout  int    `json:"timeout" toml:"timeout"`
	Username string `json:"username" toml:"username"`
	Password string `json:"password" toml:"password"`
	Charset  string `json:"charset" toml:"charset"`
	ShowSQL  bool   `json:"show_sql" toml:"show_sql"`
	Cache bool `json:"cache" toml:"cache"`
}

func (ad Address) String() string {
	//dsn = fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, dbname)
	netAddr := fmt.Sprintf("%s(%s)", "tcp", ad.Addr)
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s", ad.Username, ad.Password, netAddr, ad.Database, ad.Charset)
}

const driveName = "mysql"

func Dial(addr Address) (*xorm.Engine, error) {
	orm, err := xorm.NewEngine(driveName, addr.String())
	if err != nil {
		return nil, err
	}
	// 内存缓存
	if addr.Cache {
		ccStore := xorm.NewMemoryStore()
		cacher = xorm.NewLRUCacher(ccStore, 1000)
		if cacher != nil {
			orm.SetDefaultCacher(cacher)
		}
	}
	orm.ShowSQL(addr.ShowSQL)
	//config := map[string]string{
	//	"conn": "127.0.0.1:6379",
	//}
	//ccStore := cachestore.NewRedisCache(config)
	//fmt.Printf("ccStore=%v\n", ccStore)
	//cacher := xorm.NewLRUCacher(ccStore, 99999999)
	//orm.SetDefaultCacher(cacher)
	return orm, nil
}
