//  Copyright (c) 2012 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package auth

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseJWT(t *testing.T) {
	header := base64.RawURLEncoding.EncodeToString([]byte("header"))
	payload := base64.RawURLEncoding.EncodeToString([]byte("payload"))
	signature := base64.RawURLEncoding.EncodeToString([]byte("signature"))
	jwt := header + "." + payload + "." + signature
	payloadBytes, err := parseJWT(jwt)
	require.NoError(t, err, "malformed JSON Web Token")
	assert.Equal(t, "payload", string(payloadBytes))
}
