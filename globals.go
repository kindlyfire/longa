package main

import "github.com/kindlyfire/longa/ent"

// Not great to have use globals but... meh
var client *ent.Client
var config Config

type H = map[string]interface{}
