// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package exporters

import (
	"fmt"
	"testing"

	"github.com/dapr/components-contrib/exporters"
	daprt "github.com/dapr/dapr/pkg/testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateFullName(t *testing.T) {
	t.Run("create zipkin exporter key name", func(t *testing.T) {
		assert.Equal(t, "exporters.zipkin", createFullName("zipkin"))
	})

	t.Run("create string key name", func(t *testing.T) {
		assert.Equal(t, "exporters.string", createFullName("string"))
	})
}

func TestNewExporterRegistry(t *testing.T) {
	registry0 := NewRegistry()
	registry1 := NewRegistry()

	assert.Equal(t, registry0, registry1, "should be the same object")
}

func TestCreateExporter(t *testing.T) {
	testRegistry := NewRegistry()

	t.Run("exporter is registered", func(t *testing.T) {
		const ExporterName = "mockExporter"
		// Initiate mock object
		mockExporter := new(daprt.MockExporter)

		// act
		RegisterExporter(ExporterName, func() exporters.Exporter {
			return mockExporter
		})
		p, e := testRegistry.CreateExporter(createFullName(ExporterName))

		// assert
		assert.Equal(t, mockExporter, p)
		assert.Nil(t, e)
	})

	t.Run("exporter is not registered", func(t *testing.T) {
		const ExporterName = "fakeExporter"

		// act
		p, e := testRegistry.CreateExporter(createFullName(ExporterName))

		// assert
		assert.Nil(t, p)
		assert.Equal(t, fmt.Errorf("couldn't find exporter %s", createFullName(ExporterName)), e)
	})
}
