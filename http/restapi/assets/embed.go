package assets

import (
	"embed"
)

//go:embed swagger-ui-dist/*
var SwaggerUIDist embed.FS
