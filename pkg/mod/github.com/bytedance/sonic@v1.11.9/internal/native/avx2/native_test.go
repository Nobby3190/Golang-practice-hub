// +build amd64 


// Code generated by Makefile, DO NOT EDIT.

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package avx2

import (
    `encoding/hex`
    `fmt`
    `math`
    `strings`
    `testing`
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

func TestNative_Value(t *testing.T) {
    var v types.JsonState
    s := `   -12345`
    p := (*rt.GoString)(unsafe.Pointer(&s))
    x := value(p.Ptr, p.Len, 0, &v, 0)
    assert.Equal(t, 9, x)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    assert.Equal(t, int64(-12345), v.Iv)
    assert.Equal(t, 3, v.Ep)
}

func TestNative_Value_OutOfBound(t *testing.T) {
    var v types.JsonState
    mem := []byte{'"', '"'}
    s := rt.Mem2Str(mem[:1])
    p := (*rt.GoString)(unsafe.Pointer(&s))
    x := value(p.Ptr, p.Len, 0, &v, 0)
    assert.Equal(t, 1, x)
    assert.Equal(t, -int(types.ERR_EOF), int(v.Vt))
}

func TestNative_Quote(t *testing.T) {
    s := "hello\b\f\n\r\t\\\"\u666fworld"
    d := make([]byte, 256)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := quote(sp.Ptr, sp.Len, dp.Ptr, &dp.Len, 0)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    assert.Equal(t, len(s), rv)
    assert.Equal(t, 35, len(d))
    assert.Equal(t, `hello\u0008\u000c\n\r\t\\\"景world`, string(d))
}

func TestNative_QuoteNoMem(t *testing.T) {
    s := "hello\b\f\n\r\t\\\"\u666fworld"
    d := make([]byte, 10)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := quote(sp.Ptr, sp.Len, dp.Ptr, &dp.Len, 0)
    assert.Equal(t, -6, rv)
    assert.Equal(t, 5, len(d))
    assert.Equal(t, `hello`, string(d))
}

func TestNative_DoubleQuote(t *testing.T) {
    s := "hello\b\f\n\r\t\\\"\u666fworld"
    d := make([]byte, 256)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := quote(sp.Ptr, sp.Len, dp.Ptr, &dp.Len, types.F_DOUBLE_UNQUOTE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    assert.Equal(t, len(s), rv)
    assert.Equal(t, 44, len(d))
    assert.Equal(t, `hello\\u0008\\u000c\\n\\r\\t\\\\\\\"景world`, string(d))
}

func TestNative_Unquote(t *testing.T) {
    s := `hello\b\f\n\r\t\\\"\u2333world`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\b\f\n\r\t\\\"\u2333world", string(d))
}

func TestNative_UnquoteError(t *testing.T) {
    s := `asdf\`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_EOF), rv)
    assert.Equal(t, 5, ep)
    s = `asdf\gqwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_ESCAPE), rv)
    assert.Equal(t, 5, ep)
    s = `asdf\u1gggqwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_CHAR), rv)
    assert.Equal(t, 7, ep)
    s = `asdf\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 6, ep)
    s = `asdf\\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_DOUBLE_UNQUOTE)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 7, ep)
    s = `asdf\ud800\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 12, ep)
    s = `asdf\\ud800\\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_DOUBLE_UNQUOTE)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 14, ep)
}

func TestNative_DoubleUnquote(t *testing.T) {
    s := `hello\\b\\f\\n\\r\\t\\\\\\\"\\u2333world`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_DOUBLE_UNQUOTE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\b\f\n\r\t\\\"\u2333world", string(d))
}

func TestNative_UnquoteUnicodeReplacement(t *testing.T) {
    s := `hello\ud800world`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_UNICODE_REPLACE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\ufffdworld", string(d))
    s = `hello\ud800\ud800world`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_UNICODE_REPLACE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\ufffd\ufffdworld", string(d))
}

func TestNative_HTMLEscape(t *testing.T) {
    s := "hello\u2029\u2028<&>world"
    d := make([]byte, 256)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := html_escape(sp.Ptr, sp.Len, dp.Ptr, &dp.Len)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    assert.Equal(t, len(s), rv)
    assert.Equal(t, 40, len(d))
    assert.Equal(t, `hello\u2029\u2028\u003c\u0026\u003eworld`, string(d))
}

func TestNative_HTMLEscapeNoMem(t *testing.T) {
    s := "hello\u2029\u2028<&>world"
    d := make([]byte, 10)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := html_escape(sp.Ptr, sp.Len, dp.Ptr, &dp.Len)
    assert.Equal(t, -6, rv)
    assert.Equal(t, 5, len(d))
    assert.Equal(t, `hello`, string(d))
}

func TestNative_Vstring(t *testing.T) {
    var v types.JsonState
    i := 0
    s := `test"test\n2"`
    vstring(&s, &i, &v, 0)
    assert.Equal(t, 5, i)
    assert.Equal(t, -1, v.Ep)
    assert.Equal(t, int64(0), v.Iv)
    vstring(&s, &i, &v, 0)
    assert.Equal(t, 13, i)
    assert.Equal(t, 9, v.Ep)
    assert.Equal(t, int64(5), v.Iv)
}

func TestNative_Vstring_ValidUnescapedChars(t *testing.T) {
    var v types.JsonState
    valid := uint64(types.F_VALIDATE_STRING)
    i := 0
    s := "test\x1f\""
    vstring(&s, &i, &v, valid)
    assert.Equal(t, -int(types.ERR_INVALID_CHAR), int(v.Vt))
}

func TestNative_VstringEscapeEOF(t *testing.T) {
    var v types.JsonState
    i := 0
    s := `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\"xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"x`
    vstring(&s, &i, &v, 0)
    assert.Equal(t, 95, i)
    assert.Equal(t, 63, v.Ep)
    assert.Equal(t, int64(0), v.Iv)
}

func TestNative_VstringHangUpOnRandomData(t *testing.T) {
    v, e := hex.DecodeString(
        "228dc61efd54ef80a908fb6026b7f2d5f92a257ba8b347c995f259eb8685376a" +
        "8c4500262d9c308b3f3ec2577689cf345d9f86f9b5d18d3e463bec5c22df2d2e" +
        "4506010eba1dae7278",
    )
    assert.Nil(t, e)
    p := 1
    s := rt.Mem2Str(v)
    var js types.JsonState
    vstring(&s, &p, &js, 0)
    fmt.Printf("js: %s\n", spew.Sdump(js))
}

func TestNative_Vnumber(t *testing.T) {
    var v types.JsonState
    i := 0
    s := "1234"
    vnumber(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "1.234"
    vnumber(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 1.234, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "1.234e5"
    vnumber(&s, &i, &v)
    assert.Equal(t, 7, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 1.234e5, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "0.0125"
    vnumber(&s, &i, &v)
    assert.Equal(t, 6, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 0.0125, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "100000000000000000000"
    vnumber(&s, &i, &v)
    assert.Equal(t, 21, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 100000000000000000000.0, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "999999999999999900000"
    vnumber(&s, &i, &v)
    assert.Equal(t, 21, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 999999999999999900000.0, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "-1.234"
    vnumber(&s, &i, &v)
    assert.Equal(t, 6, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, -1.234, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
}

func TestNative_Vsigned(t *testing.T) {
    var v types.JsonState
    i := 0
    s := "1234"
    vsigned(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "-1234"
    vsigned(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(-1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "9223372036854775807"
    vsigned(&s, &i, &v)
    assert.Equal(t, 19, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(math.MaxInt64), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "-9223372036854775808"
    vsigned(&s, &i, &v)
    assert.Equal(t, 20, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(math.MinInt64), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "9223372036854775808"
    vsigned(&s, &i, &v)
    assert.Equal(t, 18, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INTEGER_OVERFLOW)), v.Vt)
    i = 0
    s = "-9223372036854775809"
    vsigned(&s, &i, &v)
    assert.Equal(t, 19, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INTEGER_OVERFLOW)), v.Vt)
    i = 0
    s = "1.234"
    vsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "0.0125"
    vsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1234e5"
    vsigned(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1234e-5"
    vsigned(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
}

func TestNative_Vunsigned(t *testing.T) {
    var v types.JsonState
    i := 0
    s := "1234"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "18446744073709551615"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 20, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, ^int64(0), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "18446744073709551616"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 19, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INTEGER_OVERFLOW)), v.Vt)
    i = 0
    s = "-1234"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "1.234"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "0.0125"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "1234e5"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1234e5"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1.234e5"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1.234e-5"
    vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
}

func TestNative_SkipOne(t *testing.T) {
    p := 0
    s := ` {"asdf": [null, true, false, 1, 2.0, -3]}, 1234.5`
    q := skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 42, p)
    assert.Equal(t, 1, q)
    p = 0
    s = `1 2.5 -3 "asdf\nqwer" true false null {} []`
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 1, p)
    assert.Equal(t, 0, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 5, p)
    assert.Equal(t, 2, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 8, p)
    assert.Equal(t, 6, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 21, p)
    assert.Equal(t, 9, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 26, p)
    assert.Equal(t, 22, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 32, p)
    assert.Equal(t, 27, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 37, p)
    assert.Equal(t, 33, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 40, p)
    assert.Equal(t, 38, q)
    q = skip_one(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, 43, p)
    assert.Equal(t, 41, q)
}

func TestNative_SkipOne_Error(t *testing.T) {
    for _, s := range([]string{
        "-", "+", "0.", "0. ", "+1", "0.0e ", "9e+", "0e-",
        "tru", "fals", "nul", "trux", "fals ", 
        `"asdf`, `"\\\"`,
    }) {
        p := 0
        q := skip_one(&s, &p, &types.StateMachine{}, uint64(0))
        assert.True(t, q < 0)
    }
}

func TestNative_SkipArray(t *testing.T) {
    p := 0
    s := `null, true, false, 1, 2.0, -3, {"asdf": "wqer"}],`
    skip_array(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, p, 48)
}

func TestNative_SkipObject(t *testing.T) {
    p := 0
    s := `"asdf": "wqer"},`
    skip_object(&s, &p, &types.StateMachine{}, uint64(0))
    assert.Equal(t, p, 15)
}

func TestNative_SkipNumber(t *testing.T) {
    p := 0
    s := `-1.23e+12`
    q := skip_number(&s, &p)
    assert.Equal(t, 9, p)
    assert.Equal(t, 0, q)
}

func TestNative_SkipNumberInJson(t *testing.T) {
    p := 0x13
    s := "{\"h\":\"1.00000\",\"i\":true,\"pass3\":1}"
    q := skip_number(&s, &p)
    assert.Equal(t, 0x13, p)
    assert.Equal(t, -2, q)
}

func TestNative_SkipOneFast(t *testing.T) {
    p := 0
    s := ` {"asdf": [null, true, false, 1, 2.0, -3]}, 1234.5`
    q := skip_one_fast(&s, &p)
    assert.Equal(t, 42, p)
    assert.Equal(t, 1, q)
    p = 0
    s = `1, 2.5, -3, "asdf\nqwer", true, false, null, {}, [],`
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 1, p)
    assert.Equal(t, 0, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 6, p)
    assert.Equal(t, 3, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 10, p)
    assert.Equal(t, 8, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 24, p)
    assert.Equal(t, 12, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 30, p)
    assert.Equal(t, 26, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 37, p)
    assert.Equal(t, 32, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 43, p)
    assert.Equal(t, 39, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 47, p)
    assert.Equal(t, 45, q)
    p += 1
    q = skip_one_fast(&s, &p)
    assert.Equal(t, 51, p)
    assert.Equal(t, 49, q)
}

func TestNative_SkipOneFast_Error(t *testing.T) {
    for _, s := range([]string{
        "{{", "[{",  "{{}",
        `"asdf`, `"\\\"`,
    }) {
        p := 0
        q := skip_one_fast(&s, &p)
        assert.True(t, q < 0)
    }
}

func TestNative_GetByPath(t *testing.T) {
    s := `{"asdf": [null, true, false, 1, 2.0, -3]}, 1234.5`
    p := 0
    path := []interface{}{"asdf", 4}
    ret := get_by_path(&s, &p, &path, types.NewStateMachine())
    assert.Equal(t, strings.Index(s, "2.0"), ret)
}

func BenchmarkNative_SkipOneFast(b *testing.B) {
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        s := `{"asdf": [null, true, false, 1, 2.0, -3]}, 1234.5`
        p := 0
        _ = skip_one_fast(&s, &p)
    }
}

func BenchmarkNative_GetByPath(b *testing.B) {
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        s := `{"asdf": [null, true, false, 1, 2.0, -3]}, 1234.5`
        p := 0
        path := []interface{}{"asdf", 3}
        sm := types.NewStateMachine()
        _ = get_by_path(&s, &p, &path, sm)
        types.FreeStateMachine(sm)
    }
}