// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package arrow

import (
	"reflect"
	"unsafe"

	"github.com/apache/arrow/go/v11/arrow/decimal256"
	"github.com/apache/arrow/go/v11/arrow/endian"
)

// Decimal256 traits
var Decimal256Traits decimal256Traits

const (
	Decimal256SizeBytes = int(unsafe.Sizeof(decimal256.Num{}))
)

type decimal256Traits struct{}

func (decimal256Traits) BytesRequired(n int) int { return Decimal256SizeBytes * n }

func (decimal256Traits) PutValue(b []byte, v decimal256.Num) {
	for i, a := range v.Array() {
		start := i * 8
		endian.Native.PutUint64(b[start:], a)
	}
}

// CastFromBytes reinterprets the slice b to a slice of decimal256
func (decimal256Traits) CastFromBytes(b []byte) []decimal256.Num {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	var res []decimal256.Num
	s := (*reflect.SliceHeader)(unsafe.Pointer(&res))
	s.Data = h.Data
	s.Len = h.Len / Decimal256SizeBytes
	s.Cap = h.Cap / Decimal256SizeBytes

	return res
}

func (decimal256Traits) CastToBytes(b []decimal256.Num) []byte {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	var res []byte
	s := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	s.Data = h.Data
	s.Len = h.Len * Decimal256SizeBytes
	s.Cap = h.Cap * Decimal256SizeBytes

	return res
}

func (decimal256Traits) Copy(dst, src []decimal256.Num) { copy(dst, src) }
