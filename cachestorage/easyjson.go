// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package cachestorage

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage(in *jlexer.Lexer, out *RequestEntriesReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheDataEntries":
			if in.IsNull() {
				in.Skip()
				out.CacheDataEntries = nil
			} else {
				in.Delim('[')
				if out.CacheDataEntries == nil {
					if !in.IsDelim(']') {
						out.CacheDataEntries = make([]*DataEntry, 0, 8)
					} else {
						out.CacheDataEntries = []*DataEntry{}
					}
				} else {
					out.CacheDataEntries = (out.CacheDataEntries)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DataEntry
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DataEntry)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.CacheDataEntries = append(out.CacheDataEntries, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "hasMore":
			out.HasMore = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage(out *jwriter.Writer, in RequestEntriesReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CacheDataEntries) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cacheDataEntries\":")
		if in.CacheDataEntries == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.CacheDataEntries {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.HasMore {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"hasMore\":")
		out.Bool(bool(in.HasMore))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestEntriesReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestEntriesReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestEntriesReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestEntriesReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage1(in *jlexer.Lexer, out *RequestEntriesParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "skipCount":
			out.SkipCount = int64(in.Int64())
		case "pageSize":
			out.PageSize = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage1(out *jwriter.Writer, in RequestEntriesParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cacheId\":")
	out.String(string(in.CacheID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"skipCount\":")
	out.Int64(int64(in.SkipCount))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"pageSize\":")
	out.Int64(int64(in.PageSize))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestEntriesParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestEntriesParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestEntriesParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestEntriesParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage2(in *jlexer.Lexer, out *RequestCachedResponseReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "response":
			if in.IsNull() {
				in.Skip()
				out.Response = nil
			} else {
				if out.Response == nil {
					out.Response = new(CachedResponse)
				}
				(*out.Response).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage2(out *jwriter.Writer, in RequestCachedResponseReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Response != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"response\":")
		if in.Response == nil {
			out.RawString("null")
		} else {
			(*in.Response).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCachedResponseReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCachedResponseReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCachedResponseReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCachedResponseReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage3(in *jlexer.Lexer, out *RequestCachedResponseParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "requestURL":
			out.RequestURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage3(out *jwriter.Writer, in RequestCachedResponseParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cacheId\":")
	out.String(string(in.CacheID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"requestURL\":")
	out.String(string(in.RequestURL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCachedResponseParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCachedResponseParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCachedResponseParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCachedResponseParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage4(in *jlexer.Lexer, out *RequestCacheNamesReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "caches":
			if in.IsNull() {
				in.Skip()
				out.Caches = nil
			} else {
				in.Delim('[')
				if out.Caches == nil {
					if !in.IsDelim(']') {
						out.Caches = make([]*Cache, 0, 8)
					} else {
						out.Caches = []*Cache{}
					}
				} else {
					out.Caches = (out.Caches)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Cache
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Cache)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Caches = append(out.Caches, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage4(out *jwriter.Writer, in RequestCacheNamesReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Caches) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"caches\":")
		if in.Caches == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Caches {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCacheNamesReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCacheNamesReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCacheNamesReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCacheNamesReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage5(in *jlexer.Lexer, out *RequestCacheNamesParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "securityOrigin":
			out.SecurityOrigin = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage5(out *jwriter.Writer, in RequestCacheNamesParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"securityOrigin\":")
	out.String(string(in.SecurityOrigin))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCacheNamesParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCacheNamesParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCacheNamesParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCacheNamesParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage6(in *jlexer.Lexer, out *DeleteEntryParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "request":
			out.Request = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage6(out *jwriter.Writer, in DeleteEntryParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cacheId\":")
	out.String(string(in.CacheID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	out.String(string(in.Request))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteEntryParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteEntryParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteEntryParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteEntryParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage6(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage7(in *jlexer.Lexer, out *DeleteCacheParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage7(out *jwriter.Writer, in DeleteCacheParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cacheId\":")
	out.String(string(in.CacheID))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteCacheParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteCacheParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteCacheParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteCacheParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage7(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage8(in *jlexer.Lexer, out *DataEntry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			out.Request = string(in.String())
		case "response":
			out.Response = string(in.String())
		case "responseTime":
			out.ResponseTime = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage8(out *jwriter.Writer, in DataEntry) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	out.String(string(in.Request))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"response\":")
	out.String(string(in.Response))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"responseTime\":")
	out.Float64(float64(in.ResponseTime))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DataEntry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DataEntry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DataEntry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DataEntry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage8(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage9(in *jlexer.Lexer, out *CachedResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "headers":
			(out.Headers).UnmarshalEasyJSON(in)
		case "body":
			out.Body = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage9(out *jwriter.Writer, in CachedResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"headers\":")
	(in.Headers).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"body\":")
	out.String(string(in.Body))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CachedResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CachedResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CachedResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CachedResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage9(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage10(in *jlexer.Lexer, out *Cache) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "securityOrigin":
			out.SecurityOrigin = string(in.String())
		case "cacheName":
			out.CacheName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage10(out *jwriter.Writer, in Cache) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cacheId\":")
	out.String(string(in.CacheID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"securityOrigin\":")
	out.String(string(in.SecurityOrigin))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cacheName\":")
	out.String(string(in.CacheName))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Cache) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Cache) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpCachestorage10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Cache) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Cache) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpCachestorage10(l, v)
}
