package setting

//  go get gopkg.in/ini.v1
import "gopkg.in/ini.v1"

var Conf = new(AppConf)

type AppConf struct {
	Port       int  `ini:"port"`
	Release    bool `ini:"release"`
	*MySQLConf `ini:"mysql"`
}

type MySQLConf struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	DB       string `ini:"db"`
}

func InitConf(file string) error {
	return ini.MapTo(Conf, file)
}
