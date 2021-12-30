package app

var Cfg = Yaml{}

type Yaml struct {
	Db struct {
		Mysql struct {
			Host string
			Port string
			User string
			Pass string
			Name string
		}
	}
	Image struct {
		Thumb struct {
			Width  int `yaml:"width"`
			Height int `yaml:"height"`
		}
		Large struct {
			Width  int `yaml:"width"`
			Height int `yaml:"height"`
		}
	}
}
