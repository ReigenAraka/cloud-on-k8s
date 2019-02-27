// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package initcontainer

import (
	"testing"

	"github.com/elastic/k8s-operators/operators/pkg/controller/elasticsearch/volume"
	"github.com/stretchr/testify/assert"
)

func TestNewInitContainers(t *testing.T) {
	type args struct {
		elasticsearchImage string
		operatorImage      string
		linkedFiles        LinkedFilesArray
		SetVMMaxMapCount   bool
		nodeCertsVolume    volume.SecretVolume
	}
	tests := []struct {
		name                       string
		args                       args
		expectedNumberOfContainers int
	}{
		{
			name: "with SetVMMaxMapCount enabled",
			args: args{
				elasticsearchImage: "es-image",
				operatorImage:      "op-image",
				linkedFiles:        LinkedFilesArray{},
				SetVMMaxMapCount:   true,
				nodeCertsVolume:    volume.SecretVolume{},
			},
			expectedNumberOfContainers: 3,
		},
		{
			name: "with SetVMMaxMapCount disabled",
			args: args{
				elasticsearchImage: "es-image",
				operatorImage:      "op-image",
				linkedFiles:        LinkedFilesArray{},
				SetVMMaxMapCount:   false,
				nodeCertsVolume:    volume.SecretVolume{},
			},
			expectedNumberOfContainers: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			containers, err := NewInitContainers(tt.args.elasticsearchImage, tt.args.operatorImage, tt.args.linkedFiles, tt.args.SetVMMaxMapCount, tt.args.nodeCertsVolume)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedNumberOfContainers, len(containers))
		})
	}
}