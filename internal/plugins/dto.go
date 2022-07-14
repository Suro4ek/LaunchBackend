package plugins

type CreatePluginDTO struct {
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
	SpigotID    string `form:"spigot_id" json:"spigot_id"`
}

type UpdatePluginDTO struct {
	ID          int32  `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
	SpigotID    string `form:"spigot_id" json:"spigot_id"`
}
