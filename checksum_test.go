// MIT License

// Copyright (c) [year] [fullname]

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package checksum

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	t.Run("can checksum lowercase address", func(t *testing.T) {
		var result string
		lowcase := "0xfb6916095ca1df60bb79ce92ce3ea74c37c5d359"
		exp := "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359"
		result = checksum(lowcase)
		if result != exp {
			t.Errorf("expected \n %v, got \n %v \n", exp, result)
		}
	})
	t.Run("can checksum uppercase address", func(t *testing.T) {
		var result string
		uppcase := "0xFB6916095CA1DF60BB79CE92CE3EA74C37C5D359"
		exp := "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359"
		result = checksum(uppcase)
		if result != exp {
			t.Errorf("expected \n %v, got \n %v \n", exp, result)
		}
	})

}

func TestIsChecksum(t *testing.T) {
	t.Run("should return false if checksum is wrong", func(t *testing.T) {
		wrongcheck := "0xc1912fEE45d61C87Cc5EA59DaE31190FFFFf232D"
		if isChecksum(wrongcheck) {
			t.Errorf("expected \n %v, got \n %v \n", true, isChecksum(wrongcheck))
		}
	})
	t.Run("should return true if checksum is correct", func(t *testing.T) {
		rightcheck := "0xc1912fEE45d61C87Cc5EA59DaE31190FFFFf232d"
		if !isChecksum(rightcheck) {
			t.Errorf("expected \n %v, got \n %v \n", true, isChecksum(rightcheck))
		}
	})
}
