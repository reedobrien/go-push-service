// Package resource is part of the wps -- an implementation of Generic Event
// Delivery Using HTTP Push.
// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.
// All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.
package resource

const receiptRelLink = "urn:ietf:params:push:receipt"

// Receipt is the resource via which an application server receives delivery
// confirmations (Section 5.1) for push messages using a receipt subscription.
// A link relation of type "urn:ietf:params:push:receipt" identifies a receipt
// subscription.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-5.1
type Receipt struct {
	linkrel string
}

// NewReceipt returns a new Reciept rsource.
func NewReceipt() Receipt {
	return Receipt{linkrel: receiptRelLink}
}

// LinkRel returns a Receipt resource's relation link type
func (r *Receipt) LinkRel() string {
	return r.linkrel
}
