package router

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	HEAD    HTTPMethod = "HEAD"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	CONNECT HTTPMethod = "CONNECT"
	OPTIONS HTTPMethod = "OPTIONS"
	TRACE   HTTPMethod = "TRACE"
	PATCH   HTTPMethod = "PATCH"
)

type HttpStatus struct {
	Code        int
	Description string
	Category    string
}

var httpStatusCodes = map[int]HttpStatus{
	100: {100, "Continue", "Informational"},
	101: {101, "Switching Protocols", "Informational"},
	102: {102, "Processing", "Informational"},
	103: {103, "Early Hints", "Informational"},
	200: {200, "OK", "Success"},
	201: {201, "Created", "Success"},
	202: {202, "Accepted", "Success"},
	203: {203, "Non-Authoritative Information", "Success"},
	204: {204, "No Content", "Success"},
	205: {205, "Reset Content", "Success"},
	206: {206, "Partial Content", "Success"},
	207: {207, "Multi-Status", "Success"},
	208: {208, "Already Reported", "Success"},
	226: {226, "IM Used", "Success"},
	300: {300, "Multiple Choices", "Redirection"},
	301: {301, "Moved Permanently", "Redirection"},
	302: {302, "Found", "Redirection"},
	303: {303, "See Other", "Redirection"},
	304: {304, "Not Modified", "Redirection"},
	305: {305, "Use Proxy", "Redirection"},
	306: {306, "(Unused)", "Redirection"},
	307: {307, "Temporary Redirect", "Redirection"},
	308: {308, "Permanent Redirect", "Redirection"},
	400: {400, "Bad Request", "Client Error"},
	401: {401, "Unauthorized", "Client Error"},
	402: {402, "Payment Required", "Client Error"},
	403: {403, "Forbidden", "Client Error"},
	404: {404, "Not Found", "Client Error"},
	405: {405, "Method Not Allowed", "Client Error"},
	406: {406, "Not Acceptable", "Client Error"},
	407: {407, "Proxy Authentication Required", "Client Error"},
	408: {408, "Request Timeout", "Client Error"},
	409: {409, "Conflict", "Client Error"},
	410: {410, "Gone", "Client Error"},
	411: {411, "Length Required", "Client Error"},
	412: {412, "Precondition Failed", "Client Error"},
	413: {413, "Payload Too Large", "Client Error"},
	414: {414, "URI Too Long", "Client Error"},
	415: {415, "Unsupported Media Type", "Client Error"},
	416: {416, "Range Not Satisfiable", "Client Error"},
	417: {417, "Expectation Failed", "Client Error"},
	418: {418, "I'm a teapot", "Client Error"},
	421: {421, "Misdirected Request", "Client Error"},
	422: {422, "Unprocessable Entity", "Client Error"},
	423: {423, "Locked", "Client Error"},
	424: {424, "Failed Dependency", "Client Error"},
	425: {425, "Too Early", "Client Error"},
	426: {426, "Upgrade Required", "Client Error"},
	427: {427, "Precondition Required", "Client Error"},
	428: {428, "Too Many Requests", "Client Error"},
	429: {429, "Request Header Fields Too Large", "Client Error"},
	431: {431, "Unavailable For Legal Reasons", "Client Error"},
	451: {451, "Unavailable For Legal Reasons", "Client Error"},
	500: {500, "Internal Server Error", "Server Error"},
	501: {501, "Not Implemented", "Server Error"},
	502: {502, "Bad Gateway", "Server Error"},
	503: {503, "Service Unavailable", "Server Error"},
	504: {504, "Gateway Timeout", "Server Error"},
	505: {505, "HTTP Version Not Supported", "Server Error"},
	506: {506, "Variant Also Negotiates", "Server Error"},
	507: {507, "Insufficient Storage", "Server Error"},
	508: {508, "Loop Detected", "Server Error"},
	510: {510, "Not Extended", "Server Error"},
	511: {511, "Network Authentication Required", "Server Error"},
}

const (
	JSON            = "application/json"
	XML             = "application/xml"
	HTML            = "text/html"
	PlainText       = "text/plain"
	JPEG            = "image/jpeg"
	PNG             = "image/png"
	GIF             = "image/gif"
	CSS             = "text/css"
	JavaScript      = "application/javascript"
	PDF             = "application/pdf"
	CSV             = "text/csv"
	MP4             = "video/mp4"
	MP3             = "audio/mpeg"
	OGG             = "audio/ogg"
	ZIP             = "application/zip"
	RAR             = "application/x-rar-compressed"
	Tar             = "application/x-tar"
	GZIP            = "application/gzip"
	Markdown        = "text/markdown"
	SVG             = "image/svg+xml"
	WebP            = "image/webp"
	WOFF            = "application/font-woff"
	WOFF2           = "application/font-woff2"
	ODT             = "application/vnd.oasis.opendocument.text"
	RTF             = "application/vnd.rtf"
	APK             = "application/vnd.android.package-archive"
	DockerImage     = "application/vnd.docker.container.image.v1+json"
	ShellScript     = "application/x-sh"
	JavaScriptPatch = "application/json-patch+json"
	BZIP            = "application/x-bzip"
	BZIP2           = "application/x-bzip2"
	TTF             = "application/x-font-ttf"
)
