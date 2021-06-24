package http

var ToPropagate = []string{
	"authorization",
	"x-b3-traceid",
	"x-b3-spanid",
	"signid"}

func Propagate(rawHeaders map[string]string, headersKey []string) map[string]string {

	result := make(map[string]string)
	for _, key := range headersKey {
		if val, ok := rawHeaders[key]; ok {
			result[key] = val
		}
	}

	return result
}
