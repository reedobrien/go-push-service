// Package resource is part of the wps -- an implementation of Generic Event
// Delivery Using HTTP Push.
// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.  All rights
// reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package resource

const pushLinkRel = "urn:ietf:params:push"

// Push is the mechanism by which an application server requests the delivery
// (Section 5) of a push message using a push resource.  A link relation of
// type "urn:ietf:params:push" identifies a push resource.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-5
type Push struct {
	linkrel string
}

// NewPush returns a new Push resource.
func NewPush() Push {
	return Push{linkrel: pushLinkRel}
}

// LinkRel returns the Push resource's link relation type.
func (r *Push) LinkRel() string {
	return r.linkrel
}
