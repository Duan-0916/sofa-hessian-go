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
	"math"
)

func DecodeFloat64Hessian4V2(o *DecodeContext, reader *bufio.Reader) (float64, error) {
	var i float64
	err := DecodeFloat64ToHessian4V2(o, reader, &i)
	return i, err
}

func DecodeFloat64ToHessian4V2(o *DecodeContext, reader *bufio.Reader, i *float64) error {
	if o.tracer != nil {
		o.tracer.OnTraceStart("decodefloat64")
		defer o.tracer.OnTraceStop("decodefloat64")
	}

	c1, err := reader.ReadByte()
	if err != nil {
		return err
	}

	if c1 == 0x44 {
		u64, err := readUint64FromReader(reader)
		if err != nil {
			return err
		}
		*i = math.Float64frombits(u64)
		return nil
	}

	switch c1 {
	case 0x5b:
		*i = 0.0
		return nil

	case 0x5c:
		*i = 1.0
		return nil

	case 0x5d:
		c2, err := reader.ReadByte()
		if err != nil {
			return err
		}
		*i = float64(int8(c2))
		return nil

	case 0x5e:
		i16, err := readInt16FromReader(reader)
		if err != nil {
			return err
		}
		*i = float64(i16)
		return nil

	case 0x5f:
		i32, err := readInt32FromReader(reader)
		if err != nil {
			return err
		}
		*i = float64(i32) * 0.001
		return nil
	}

	return ErrDecodeMalformedDouble
}
