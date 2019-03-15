// Copyright 2017 <614422099@qq.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// ini file resolve.
//
// https://github.com/dingdayu/golangtools
package ini

import (
	"github.com/go-gcfg/gcfg"
)

// resolve ini
func Unmarshal(f []byte, result interface{}) error {
	return gcfg.ReadStringInto(result, string(f))
}
