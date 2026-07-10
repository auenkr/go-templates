// Package pkgfx contains the fx.Module for the pkg package.
package pkgfx

import (
	"github.com/auenkr/go-templates/connect-server/pkg/config"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"pkg",
	config.Module,
)
