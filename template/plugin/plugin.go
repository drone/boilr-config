// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/config"
)

// TODO replace or remove
const defaultPipeline = `
kind: pipeline
name: default

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test -v
`

// New returns a new config plugin.
func New(param1, param2 string) config.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		param1: param1,
		param2: param2,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	param1 string
	param2 string
}

func (p *plugin) Find(ctx context.Context, req *config.Request) (*drone.Config, error) {
	// TODO replace or remove
	// this demonstrates how we can override
	// and return a custom configuration file
	if req.Repo.Namespace == "some-organization" {
		return &drone.Config{
			Data: defaultPipeline,
		}, nil
	}

	// return nil and Drone will fallback to
	// the standard behavior for getting the
	// configuration file.
	return nil, nil
}
