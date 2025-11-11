package icons

func Alert() IconFunc {
	return Icon("alert", MDI)
}

func Information(variant ...string) IconFunc {
	name := "information"
	if len(variant) > 0 {
		for _, v := range variant {
			name += "-" + v
		}
	}
	return Icon(name, MDI)
}

func Success() IconFunc {
	return Icon("check", MDI)
}

func Close(variant ...string) IconFunc {
	name := "close"
	if len(variant) > 0 {
		for _, v := range variant {
			name += "-" + v
		}
	}
	return Icon(name, MDI)
}

func Chevron(variant ...string) IconFunc {
	name := "chevron"

	if len(variant) > 0 {
		for _, d := range variant {
			name += "-" + d
		}
	} else {
		name = "chevron-down"
	}

	return Icon(name, MDI)
}
