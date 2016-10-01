// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.
// All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package resource

// Message is a resource which identifies push messages that have been
// accepted for delivery (Section 5).  The push message resource is also
// deleted by the user agent to acknowledge receipt (Section 6.2) of a push
// message.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-5
type Message struct {
}
