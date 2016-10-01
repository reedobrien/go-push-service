// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.
// All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package resource

// SubscriptionSet is a resource which provides read and delete access for a
// collection of push message subscriptions.  A user agent receives push
// messages (Section 6.1) for all the push message subscriptions in the set.  A
// link relation of type "urn:ietf:params:push:set" identifies a push message
// subscription set.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-6.1
type SubscriptionSet struct{}
