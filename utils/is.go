package utils

func IsLink(k string) bool {
	switch k {
	case "href", "src", "srcset", "data", "action", "formaction", "background", "poster":
		return true
	}

	return false
}
