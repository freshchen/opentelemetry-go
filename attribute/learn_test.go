// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attribute_test

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/attribute"
)

func TestKeyValue(t *testing.T) {
	sAttr := attribute.String("gen_ai.input", "hello world")

	fmt.Printf("sAttr Value Emit: %v\n", sAttr.Value.Emit())

	sSliceAttr := attribute.StringSlice("gen_ai.input", []string{"hello world", "你好", "Hi"})
	fmt.Printf("sSliceAttr Value Emit: %v\n", sSliceAttr.Value.Emit())

	attrs := attribute.NewSet(sAttr, sSliceAttr)
	fmt.Printf("%#v\n", attrs)
	attrsVByKey, _ := attrs.Value("gen_ai.input")
	fmt.Printf("%#v\n", attrsVByKey)

	fmt.Printf("%#v\n", attrs.Encoded(attribute.DefaultEncoder()))
}
