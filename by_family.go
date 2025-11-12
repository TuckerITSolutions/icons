package icons

import (
	"strings"

	"github.com/a-h/templ"
)

func BoxIcon(name string, props ...IconProps) templ.Component {
	if !strings.Contains(name, "bx") {
		name = "bx-" + name
	}

	if len(props) > 0 {
		props[0].Family = BOX
		return Icon(name, props[0])
	}
	return Icon(name, IconProps{Family: BOX})
}

func BootstrapIcon(name string, props ...IconProps) templ.Component {
	if len(props) > 0 {
		props[0].Family = BOOTSTRAP
		return Icon(name, props[0])
	}

	return Icon(name, IconProps{Family: BOOTSTRAP})
}

func IonicIcon(name string, props ...IconProps) templ.Component {
	if len(props) > 0 {
		props[0].Family = IONIC
		return Icon(name, props[0])
	}
	return Icon(name, IconProps{Family: IONIC})
}

func LineIcon(name string, props ...IconProps) templ.Component {
	if len(props) > 0 {
		props[0].Family = LINE
		return Icon(name, props[0])
	}
	return Icon(name, IconProps{Family: LINE})
}

func MaterialDesignIcon(name string, props ...IconProps) templ.Component {
	if len(props) > 0 {
		props[0].Family = MDI
		return Icon(name, props[0])
	}
	return Icon(name, IconProps{Family: MDI})
}

func HeroIcon(name string, props ...IconProps) templ.Component {
	if len(props) > 0 {
		props[0].Family = HERO
		return Icon(name, props[0])
	}
	return Icon(name, IconProps{Family: HERO})
}

func FontAwesomeIcon(name string, props ...IconProps) templ.Component {
	if len(props) > 0 {
		props[0].Family = FA
		return Icon(name, props[0])
	}
	return Icon(name, IconProps{Family: FA})

}
func LucideIcon(name string, props ...IconProps) templ.Component {
	if len(props) > 0 {
		props[0].Family = LUCIDE
		return Icon(name, props[0])
	}
	return Icon(name, IconProps{Family: LUCIDE})

}
