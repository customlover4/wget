package utils

func FileExtension(contentType [2]string) string {
	mainType := contentType[0]
	subType := contentType[1]

	var result string
	switch mainType {
	case "text":
		switch subType {
		case "html":
			result = "html"
		case "css":
			result = "css"
		case "javascript":
			result = "js"
		case "js":
			result = "js"
		case "plain":
			result = "txt"
		case "xml":
			result = "xml"
		}
	case "application":
		switch subType {
		case "json":
			result = "json"
		case "xml":
			result = "xml"
		case "atom+xml":
			result = "xml"
		case "pdf":
			result = "pdf"
		case "zip":
			result = "zip"
		case "x-zip-compressed":
			result = "zip"
		case "gzip":
			result = "gz"
		case "x-gzip":
			result = "tar.gz"
		case "x-javascript":
			result = "js"
		case "javascript":
			result = "js"
		case "js":
			result = "js"
		case "octet-stream":
			result = "bin"
		}
	case "image":
		switch subType {
		case "jpeg":
			result = "jpeg"
		case "jpg":
			result = "jpg"
		case "png":
			result = "png"
		case "gif":
			result = "gif"
		case "webp":
			result = "webp"
		case "svg+xml":
			result = "svg"
		}
	case "audio":
		switch subType {
		case "mpeg":
			result = "mp3"
		case "wav":
			result = "wav"
		case "ogg":
			result = "ogg"
		}
	case "video":
		switch subType {
		case "mp4":
			result = "mp4"
		case "webm":
			result = "webm"
		}
	}

	if result == "" {
		return ""
	}

	return "." + result
}
