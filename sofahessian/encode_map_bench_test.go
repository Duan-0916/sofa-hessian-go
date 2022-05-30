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
	"testing"
)

func BenchmarkEncodeMap(b *testing.B) {
	m := make(map[int32]string, 256)
	m[1] = "fee"
	m[16] = "fie"
	m[256] = "foe"
	ectx := NewEncodeContext()

	var dst []byte
	var err error

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dst, err = EncodeToHessian4V2(ectx, dst[:0], m)
		b.SetBytes(int64(len(dst)))
		if err != nil {
			b.Fatal(err)
		}
	}
}