package content

import "embed"

//go:embed * */*
var ContentFS embed.FS
