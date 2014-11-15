// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// See https://code.google.com/p/go/source/browse/CONTRIBUTORS
// Licensed under the same terms as Go itself:
// https://code.google.com/p/go/source/browse/LICENSE

package http2

import "testing"

func TestErrCodeString(t *testing.T) {
	tests := []struct {
		err  ErrCode
		want string
	}{
		{ErrCodeProtocol, "PROTOCOL_ERROR"},
		{0xf, "unknown error code 0xf"},
	}
	for i, tt := range tests {
		got := tt.err.String()
		if got != tt.want {
			t.Errorf("%i. Error = %q; want %q", i, got, tt.want)
		}
	}
}
