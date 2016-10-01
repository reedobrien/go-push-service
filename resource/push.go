// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.  All rights
// reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package resource

// Push is the mechanism by which an application server requests the delivery
// (Section 5) of a push message using a push resource.  A link relation of
// type "urn:ietf:params:push" identifies a push resource.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-5
type Push struct {
}
