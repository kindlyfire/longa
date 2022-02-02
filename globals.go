package main

import (
	"embed"

	"github.com/kindlyfire/longa/ent"
)

// Not great to have use globals but... meh
var client *ent.Client
var config Config

//go:embed frontend/dist
var frontendContent embed.FS

type H = map[string]interface{}
