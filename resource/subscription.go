// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.
// All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package resource

// Subscription is a resource which provides read and delete access for a
// message subscription.  A user agent receives push messages (Section 6) using
// a push message subscription.  Every push message subscription has exactly
// one push resource associated with it.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-6
type Subscription struct{}
