// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoc-gen-go binary is a protoc plugin to generate a Go protocol
// buffer package.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stderr, "%v %v\n", filepath.Base(os.Args[0]), protoimpl.VersionString())
		os.Exit(1)
	}

	var (
		flags        flag.FlagSet
		plugins      = flags.String("plugins", "", "deprecated option")
		importPrefix = flags.String("import_prefix", "", "deprecated option")
		opts         = &protogen.Options{
			ParamFunc: flags.Set,
		}
	)
	protogen.Run(opts, func(gen *protogen.Plugin) error {
		if *plugins != "" {
			return errors.New("protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC")
		}
		if *importPrefix != "" {
			return errors.New("protoc-gen-go: import_prefix is not supported")
		}
		for _, f := range gen.Files {
			if f.Generate {
				gengo.GenerateFile(gen, f)
			}
		}
		return nil
	})
}
