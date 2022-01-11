package app

import "fmt"

const (
	VersionName = "ZLab Drive"

	VersionNumber = "0.2.0"

	Website = "https://zlab.dev"

	// http://patorjk.com/software/taag/#p=display&h=0&f=Small%20Slant&t=RPC
	banner = `
   ___          _
  / _ \  ____  (_) _  __ ___
 / // / / __/ / / | |/ // -_)
/____/ /_/   /_/  |___/ \__/  %s %s
High performance, App framework
Support by %s
%s
____________________________________O/_______
                                    O\
`
)

func Banner(message string) {
	fmt.Printf(banner, VersionName, VersionNumber, Website, message)
}
