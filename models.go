package petstore

// Category represents a category that a Pet belongs to
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Tag represents a set of free form tags that a Pet can be associated with
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Pet is is a single instance of a Pet in the store which holds its name
// photo, category, tags and its status.
type Pet struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	Photos []string `json:"photoUrls"`

	Category Category `json:"category"`
	Tags     []Tag    `json:"tags"`

	Status string `json:"status"`
}

// APIResponse is a structure returned by all endpoints to incidate a response
type APIResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
