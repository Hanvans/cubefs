// Copyright 2023 The CubeFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package compressor_test

import (
	"crypto/rand"
	"testing"

	"github.com/cubefs/cubefs/util/compressor"
	"github.com/stretchr/testify/require"
)

func TestCompressor_New(t *testing.T) {
	for _, encoding := range []string{"", "none", "balaa"} {
		buf := make([]byte, 1024)
		rand.Read(buf)

		c := compressor.New(encoding)
		require.NotNil(t, c)
		cbuf, err := c.Compress(buf)
		require.NoError(t, err)
		require.Equal(t, buf, cbuf)
		pbuf, err := c.Decompress(cbuf)
		require.NoError(t, err)
		require.Equal(t, buf, pbuf)
	}
}
