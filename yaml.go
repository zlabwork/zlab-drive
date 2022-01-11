package app

var Cfg = Yaml{}

type Yaml struct {
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
