package icons

import "github.com/a-h/templ"

func Alert() templ.Component {
	return Icon("alert", IconProps{Family: MDI})
}

func Home(variant ...string) templ.Component {
	name := "home"

	if len(variant) != 0 {
		for _, v := range variant {
			name += "-" + v
		}
	}
	return Icon(name, IconProps{Family: MDI})
}

func Information(variant ...string) templ.Component {
	name := "information"
	if len(variant) > 0 {
		for _, v := range variant {
			name += "-" + v
		}
	}
	return Icon(name, IconProps{Family: MDI})
}

func Success() templ.Component {
	return Icon("check", IconProps{Family: MDI})
}

func Close(variant ...string) templ.Component {
	name := "close"
	if len(variant) > 0 {
		for _, v := range variant {
			name += "-" + v
		}
	}
	return Icon(name, IconProps{Family: MDI})
}

func Chevron(variant ...string) templ.Component {
	name := "chevron"

	if len(variant) > 0 {
		for _, v := range variant {
			name += "-" + v
		}
	} else {
		name = "chevron-down"
	}

	return Icon(name, IconProps{Family: MDI})
}

func Settings(variant ...string) templ.Component {
	name := "cog"

	for _, v := range variant {

		name += "-" + v
	}

	return Icon(name, IconProps{Family: MDI})
}
