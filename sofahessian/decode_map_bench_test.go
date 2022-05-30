// nolint
// Copyright 20xx The Alipay Authors.
//
// @authors[0]: bingwu.ybw(bingwu.ybw@antfin.com|detailyang@gmail.com)
// @authors[1]: robotx(robotx@antfin.com)
//
// *Legal Disclaimer*
// Within this source code, the comments in Chinese shall be the original, governing version. Any comment in other languages are for reference only. In the event of any conflict between the Chinese language version comments and other language version comments, the Chinese language version shall prevail.
// *法律免责声明*
// 关于代码注释部分，中文注释为官方版本，其它语言注释仅做参考。中文注释可能与其它语言注释存在不一致，当中文注释与其它语言注释存在不一致时，请以中文注释为准。
//
//

package sofahessian

import (
	"bufio"
	"bytes"
	"testing"
)

func BenchmarkDecodeMap(b *testing.B) {
	buf := readFile(b, "testdata/map/car.bin")
	r := bytes.NewReader(buf)
	br := bufio.NewReader(r)

	o := NewDecodeContext()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.SetBytes(int64(len(buf)))
		_, err := DecodeHessian4V2(o, br)
		if err != nil {
			b.Fatal(err)
		}
		r.Reset(buf)
		br.Reset(r)
	}
}
