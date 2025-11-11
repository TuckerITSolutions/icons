package icons

import "strings"

func BoxIcon(name string) IconFunc {
	if !strings.Contains(name, "bx") {
		name = "bx-" + name
	}

	return Icon(name, BOX)
}

func BootstrapIcon(name string) IconFunc {
	return Icon(name, BOOTSTRAP)
}

func IonicIcon(name string) IconFunc {
	return Icon(name, IONIC)
}

func LineIcon(name string) IconFunc {
	return Icon(name, LINE)
}

func MaterialDesignIcon(name string) IconFunc {
	return Icon(name, MDI)
}

func HeroIcon(name string) IconFunc {
	return Icon(name, HERO)
}
