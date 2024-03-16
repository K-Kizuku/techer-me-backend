package schema

type GenerateImageInput struct {
	ObjectName string `json:"object_name"`
	Type       string `json:"type"`
}

type GenerateImageOutput struct {
	URL string `json:"url"`
}
