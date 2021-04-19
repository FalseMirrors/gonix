package apiJson

const urlHostAPI string = "https://jsonplaceholder.typicode.com"

var resoureces = map[string]string{
	"posts":		 "/posts",
	"comments":		 "/comments",
	"albums":		 "/albums",
	"photos":		 "/photos",
	"todos": 		 "/todos",
	"users":		 "/users",
}

func UrlHostAPI() string {
	return urlHostAPI
}

func Resources(idx string) string {
	if _, ok := resoureces[idx]; ok {
		return resoureces[idx]
	}
	return ""
}
