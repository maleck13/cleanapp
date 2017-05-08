// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"os"

	"github.com/maleck13/cleanapp/cmd/golang"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init will scaffold a new web based application",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		// TODO: Work your own magic here
		t, err := cmd.Flags().GetString("template")
		if err != nil || t == "" {
			log.Fatal("no template location specified")
		}
		a, err := cmd.Flags().GetString("name")
		if err != nil || a == "" {
			log.Fatal("no app name specified")
		}
		r, err := cmd.Flags().GetString("runtime")
		if err != nil || r == "" {
			log.Fatal("no runtime specified")
		}
		o, err := cmd.Flags().GetString("out")
		if err != nil {
			log.Fatal("no outputdir specified")
		}
		if o == "" {
			o, err = os.Getwd()
			if err != nil {
				log.Fatal("failed to get current dir ", err.Error())
			}
		}
		switch r {
		case "golang":
			err = golang.Template(o, t, a)
		}
		if err != nil {
			log.Fatal("failed to init app ", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("template", "t", "", "location of template you want to use")
	initCmd.Flags().StringP("name", "n", "app", "the name of your app")
	initCmd.Flags().StringP("runtime", "r", "golang", "the runtime of your app")
	initCmd.Flags().StringP("out", "o", "", "the output path")

}
