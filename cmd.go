// Copyright 2022 bonzai-example Authors
// SPDX-License-Identifier: Apache-2.0

package emb

import (
	"embed"
	_ "embed"
	"os"
	"strings"
	"text/template"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fn/each"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{

	Name:      `emb`,
	Summary:   `access embedded files`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Source:    `git@github.com:rwxrob/emb.git`,
	Issues:    `github.com/rwxrob/emb/issues`,

	Init: func(x *Z.Cmd, args ...string) error {
		var zero embed.FS
		switch {
		case FS == zero:
			return MissingFS{}
		case Top == "":
			return MissingTop{}
		}
		return nil
	},

	Commands: []*Z.Cmd{help.Cmd, catCmd, listCmd},

	Dynamic: template.FuncMap{
		`list`: func() string { return strings.Join(RelPaths(), "\n") },
	},

	Description: `
		The **{{.Name}}** command provides access to files embedded within
		this binary:

		  {{ list }}

		`,
}

var catCmd = &Z.Cmd{
	Name:     `cat`,
	Aliases:  []string{`print`},
	MinArgs:  1,
	Summary:  `concatenate content of embedded files`,
	Usage:    `GLOB ...`,
	Commands: []*Z.Cmd{help.Cmd},
	Comp:     new(comp),

	Description: `
		The {{ cmd .Name }} command will concatenate the contents of each
		embedded files that matches one or more Go globs. To first see the
		list of files use {{ cmd "list" }}.

		Note that no extraneous line returns are added meaning that binary
		and other files can be included as well with other text embedded
		files.

	`,

	Call: func(x *Z.Cmd, args ...string) error {
		byt, err := Cat(args...)
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(byt)
		if err != nil {
			return err
		}
		return nil
	},
}

var listCmd = &Z.Cmd{
	Name:     `list`,
	Aliases:  []string{`files`, `ls`},
	Summary:  `list all embedded files and directories`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		each.Println(RelPaths())
		return nil
	},
}
