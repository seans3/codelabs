/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package foo

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

const defaultFilename = "default.yaml"

var (
	fooLong = templates.LongDesc(i18n.T(`
This is the foo command long description.
`))

	fooExample = templates.Examples(i18n.T(`
		# Foo command example
		kubectl foo --count 3 --filename foo-resource.yaml
`))
)

// FooOptions are the knobs available for the "foo" command.
type FooOptions struct {
	Count    int
	Filename string
}

// NewCmdFoo a new Cobra command encasulating the "foo" command.
func NewCmdFoo(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	o := &FooOptions{}

	cmd := &cobra.Command{
		Use: "foo [--count=COUNT] --filename=FILENAME",
		DisableFlagsInUseLine: true,
		Short:   i18n.T("Foo short description"),
		Long:    fooLong,
		Example: fooExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(args))
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.RunFoo())
		},
	}

	cmd.Flags().IntVarP(&o.Count, "count", "c", o.Count, "Usage for count flag.")
	cmd.Flags().StringVarP(&o.Filename, "filename", "f", o.Filename, i18n.T("Usage for filename flag."))

	return cmd
}

// Complete fills in all the FooOptions fields, including defaults.
func (o *FooOptions) Complete(args []string) error {
	if len(o.Filename) == 0 {
		o.Filename = defaultFilename
	}

	return nil
}

// Validate ensures all FooOptions fields are valid.
func (o *FooOptions) Validate() error {
	if o.Count < 0 {
		return fmt.Errorf("Count is negative")
	}

	if len(o.Filename) == 0 {
		return fmt.Errorf("Filename is empty")
	}

	return nil
}

// RunFoo executes the foo command.
func (o *FooOptions) RunFoo() error {

	fmt.Printf("Count: %d\n", o.Count)
	fmt.Printf("Filename: %s\n", o.Filename)

	return nil
}
