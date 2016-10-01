// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.
// All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.
package resource_test

import (
	"testing"

	"gitlab.com/fffd/wps/assert"
	"gitlab.com/fffd/wps/resource"
)

func TestSubscriptionSetLinkReLinkRell(t *testing.T) {
	ss := resource.NewSubscriptionSet()
	assert.Equals(t, ss.LinkRel(), "urn:ietf:params:push:set")
}
