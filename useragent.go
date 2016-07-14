// Copyright (c) 2016, Reed O'Brien
// All rights reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.
package wps

// UserAgen is mispelled to test linting. This is done so gitlab CI can be
// configured.
type UserAgent struct { 
ID string
}

// NewUserAgent is a function we can test. See above.
func NewUserAgent(id string) UserAgent {
    return UserAgent{ID:    id}
}
