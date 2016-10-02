// Package resource is part of the wps -- an implementation of Generic Event
// Delivery Using HTTP Push.
// Copyright © 2016 Reed O'Brien <reed+oss@reedobrien.com>.  All rights
// reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package resource

import "fmt"

// urgency MAY be set by an application server or a user agent.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-5.3
// We make it an int so we can sort messages in a subscription set by urgency.
// NB: The urgency decreases inverst to the value.
type urgency int

// TODO(ro) 2016-10-01 fixme.
const (
	// High urgency
	High = iota
	// Normal urgency
	Normal
	// Low urgency
	Low
	// VeryLow urgency
	VeryLow
	// None urgency
	None
)

// Message is a resource which identifies push messages that have been
// accepted for delivery (Section 5).  The push message resource is also
// deleted by the user agent to acknowledge receipt (Section 6.2) of a push
// message.
// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-5
type Message struct {

	// An application server can include the Prefer header field [RFC7240] with
	// the "respond-async" preference to request confirmation from the push
	// service when a push message is delivered and then acknowledged by the
	// user agent. The push service MUST support delivery confirmations.
	Prefer string // RFC7420 https://tools.ietf.org/html/rfc7240

	// Only push messages that have been assigned a topic can be replaced.
	// A push message with a topic replaces any outstanding push message
	// with an identical topic.

	// A push message topic is a string carried in a Topic header field.  A
	// topic is used to correlate push messages sent to the same
	// subscription and does not convey any other semantics.

	// The grammar for the Topic header field uses the "token" rule defined
	// in [RFC7230] <https://tools.ietf.org/html/rfc7230>.
	Topic string // Max 32 https://tools.ietf.org/html/rfc4648 characters

	// Time to live in seconds.
	// https://tools.ietf.org/html/draft-ietf-webpush-protocol-07#section-5.2
	TTL int // Required from application server, or MUST return 400.

	// An application server MAY include an Urgency header field in its request
	// for push message delivery.  This header field indicates the message
	// urgency.  The push service MUST not forward the Urgency header field to
	// the user agent.  A push message without the Urgency header field
	// defaults to a value of "normal".

	// A user agent MAY include the Urgency header field when monitoring for
	// push messages to indicate the lowest urgency of push messages that it is
	// willing to receive.  A push service MUST NOT deliver push messages with
	// lower urgency than the value indicated by the user agent in its
	// monitoring request.  Push messages of any urgency are delivered to a
	// user agent that does not include an Urgency header field when monitoring
	// for messages.

	// Urgency = 1#(urgency-option)
	// urgency-option = ("very-low" / "low" / "normal" / "high")
	// Multiple values for the Urgency header field MUST NOT be included in
	// requests; otherwise, the push service MUST return a 400 (Bad Request)
	// status code.
	Urgency urgency
}

// NewMessage returns a new Message unless it receives an invalid value of
// urgency
func NewMessage(u int) (Message, error) {
	if u > -1 && u < 5 {
		return Message{Urgency: urgency(u)}, nil
	}
	return Message{}, fmt.Errorf("invalid urgency, '%d'", u)
}

// ByUrgency implements sort.Sort for []Message based on the Urgency field.
type ByUrgency []Message

func (u ByUrgency) Len() int           { return len(u) }
func (u ByUrgency) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u ByUrgency) Less(i, j int) bool { return u[i].Urgency < u[j].Urgency }
