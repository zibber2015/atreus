package main

var _headerTemplate = `
// Code generated by kratos tool mcgen. DO NOT EDIT.

NEWLINE
/* 
  Package {{.PkgName}} is a generated mc cache package.
  It is generated from:
  ARGS
*/
NEWLINE

package {{.PkgName}}

import (
	"context"
	"fmt"
	{{if .UseStrConv}}"strconv"{{end}}
	{{if .EnableBatch }}"sync"{{end}}
NEWLINE
	{{if .UseMemcached }}"github.com/mapgoo-lab/atreus/pkg/cache/memcache"{{end}}
	{{if .EnableBatch }}"github.com/mapgoo-lab/atreus/pkg/sync/errgroup"{{end}}
	"github.com/mapgoo-lab/atreus/pkg/log"
	{{.ImportPackage}}
)

var (
	_ _mc
)
`
