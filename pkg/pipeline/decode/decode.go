/*
 * Copyright (C) 2021 IBM, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package decode

import (
	log "github.com/sirupsen/logrus"
	"github.ibm.com/MCNM/observability/flowlogs2metrics/pkg/config"
)

type Decoder interface {
	Decode(in []interface{}) []config.GenericMap
}

type decodeNone struct {
}

// Decode decodes input strings to a list of flow entries
func (c *decodeNone) Decode(in []interface{}) []config.GenericMap {
	log.Debugf("entering Decode none")
	log.Debugf("Decode none, in = %v", in)
	var f []config.GenericMap
	return f
}

// NewDecodeNone create a new decode
func NewDecodeNone() (Decoder, error) {
	return &decodeNone{}, nil
}
