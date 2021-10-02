package drive

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
		ThumbWidth int `yaml:"thumbWidth"`
	}
}
