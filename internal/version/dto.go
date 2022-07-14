package version

type CreateVersionDTO struct {
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Version     string  `form:"version" json:"version"`
	Url         *string `form:"url" json:"url"`
	JavaVersion string  `form:"java_version" json:"java_version"`
}

type UpdateVersionDTO struct {
	ID          string  `form:"id" json:"id"`
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Version     string  `form:"version" json:"version"`
	Url         *string `form:"url" json:"url"`
	JavaVersion string  `form:"java_version" json:"java_version"`
}
