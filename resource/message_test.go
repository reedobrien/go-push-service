// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.
// All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.
package resource_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"gitlab.com/fffd/wps/assert"
	"gitlab.com/fffd/wps/resource"
)

func TestMessageByUrgency(t *testing.T) {
	// TODO(ro) 2016-10-01 use structs and more info so we can check errors)
	tests := [][]int{
		[]int{1, 2, 3, 4},
		[]int{4, 4, 4, 3, 2, 3, 4, 3, 2, 1, 0},
		[]int{1, 2, 2, 4, 0, 3},
		[]int{0, 1, 2, 2, 2, 2, 2, 3, 4},
		[]int{2, 3, 1, 4, 3, 0, 7}, // Fail on idx 4
	}
	var msgs resource.ByUrgency
	for _, test := range tests {
		var msg resource.Message
		var err error
		for i, u := range test {
			fmt.Println("run:", i)
			msg, err = resource.NewMessage(u)
			if err != nil {
				assert.Equals(t, strings.HasPrefix(err.Error(), "invalid urgency,"), true)
			} else {
				assert.OK(t, err)
			}
			msgs = append(msgs, msg)

			sort.Sort(resource.ByUrgency(msgs))
			var last int
			for _, m := range msgs {
				assert.Equals(t, true, last <= int(m.Urgency))
				last = int(m.Urgency)

			}
		}
	}
}
