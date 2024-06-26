// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package pwa

import (
	json "encoding/json"
	target "github.com/chromedp/cdproto/target"
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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa(in *jlexer.Lexer, out *UninstallParams) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestId":
			out.ManifestID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa(out *jwriter.Writer, in UninstallParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UninstallParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UninstallParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UninstallParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UninstallParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa1(in *jlexer.Lexer, out *OpenCurrentPageInAppParams) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestId":
			out.ManifestID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa1(out *jwriter.Writer, in OpenCurrentPageInAppParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OpenCurrentPageInAppParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OpenCurrentPageInAppParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OpenCurrentPageInAppParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OpenCurrentPageInAppParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa2(in *jlexer.Lexer, out *LaunchReturns) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "targetId":
			out.TargetID = target.ID(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa2(out *jwriter.Writer, in LaunchReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.TargetID != "" {
		const prefix string = ",\"targetId\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.TargetID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LaunchReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LaunchReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LaunchReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LaunchReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa3(in *jlexer.Lexer, out *LaunchParams) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestId":
			out.ManifestID = string(in.String())
		case "url":
			out.URL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa3(out *jwriter.Writer, in LaunchParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LaunchParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LaunchParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LaunchParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LaunchParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa4(in *jlexer.Lexer, out *LaunchFilesInAppReturns) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "targetIds":
			if in.IsNull() {
				in.Skip()
				out.TargetIDs = nil
			} else {
				in.Delim('[')
				if out.TargetIDs == nil {
					if !in.IsDelim(']') {
						out.TargetIDs = make([]target.ID, 0, 4)
					} else {
						out.TargetIDs = []target.ID{}
					}
				} else {
					out.TargetIDs = (out.TargetIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 target.ID
					v1 = target.ID(in.String())
					out.TargetIDs = append(out.TargetIDs, v1)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa4(out *jwriter.Writer, in LaunchFilesInAppReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.TargetIDs) != 0 {
		const prefix string = ",\"targetIds\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.TargetIDs {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LaunchFilesInAppReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LaunchFilesInAppReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LaunchFilesInAppReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LaunchFilesInAppReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa5(in *jlexer.Lexer, out *LaunchFilesInAppParams) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestId":
			out.ManifestID = string(in.String())
		case "files":
			if in.IsNull() {
				in.Skip()
				out.Files = nil
			} else {
				in.Delim('[')
				if out.Files == nil {
					if !in.IsDelim(']') {
						out.Files = make([]string, 0, 4)
					} else {
						out.Files = []string{}
					}
				} else {
					out.Files = (out.Files)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Files = append(out.Files, v4)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa5(out *jwriter.Writer, in LaunchFilesInAppParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	{
		const prefix string = ",\"files\":"
		out.RawString(prefix)
		if in.Files == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Files {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LaunchFilesInAppParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LaunchFilesInAppParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LaunchFilesInAppParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LaunchFilesInAppParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa6(in *jlexer.Lexer, out *InstallParams) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestId":
			out.ManifestID = string(in.String())
		case "installUrlOrBundleUrl":
			out.InstallURLOrBundleURL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa6(out *jwriter.Writer, in InstallParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	if in.InstallURLOrBundleURL != "" {
		const prefix string = ",\"installUrlOrBundleUrl\":"
		out.RawString(prefix)
		out.String(string(in.InstallURLOrBundleURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InstallParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InstallParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InstallParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InstallParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa7(in *jlexer.Lexer, out *GetOsAppStateReturns) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "badgeCount":
			out.BadgeCount = int64(in.Int64())
		case "fileHandlers":
			if in.IsNull() {
				in.Skip()
				out.FileHandlers = nil
			} else {
				in.Delim('[')
				if out.FileHandlers == nil {
					if !in.IsDelim(']') {
						out.FileHandlers = make([]*FileHandler, 0, 8)
					} else {
						out.FileHandlers = []*FileHandler{}
					}
				} else {
					out.FileHandlers = (out.FileHandlers)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *FileHandler
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(FileHandler)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.FileHandlers = append(out.FileHandlers, v7)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa7(out *jwriter.Writer, in GetOsAppStateReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.BadgeCount != 0 {
		const prefix string = ",\"badgeCount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.BadgeCount))
	}
	if len(in.FileHandlers) != 0 {
		const prefix string = ",\"fileHandlers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.FileHandlers {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOsAppStateReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOsAppStateReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOsAppStateReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOsAppStateReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa8(in *jlexer.Lexer, out *GetOsAppStateParams) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestId":
			out.ManifestID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa8(out *jwriter.Writer, in GetOsAppStateParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOsAppStateParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOsAppStateParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOsAppStateParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOsAppStateParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa8(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa9(in *jlexer.Lexer, out *FileHandlerAccept) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "mediaType":
			out.MediaType = string(in.String())
		case "fileExtensions":
			if in.IsNull() {
				in.Skip()
				out.FileExtensions = nil
			} else {
				in.Delim('[')
				if out.FileExtensions == nil {
					if !in.IsDelim(']') {
						out.FileExtensions = make([]string, 0, 4)
					} else {
						out.FileExtensions = []string{}
					}
				} else {
					out.FileExtensions = (out.FileExtensions)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.FileExtensions = append(out.FileExtensions, v10)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa9(out *jwriter.Writer, in FileHandlerAccept) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"mediaType\":"
		out.RawString(prefix[1:])
		out.String(string(in.MediaType))
	}
	{
		const prefix string = ",\"fileExtensions\":"
		out.RawString(prefix)
		if in.FileExtensions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.FileExtensions {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileHandlerAccept) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileHandlerAccept) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileHandlerAccept) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileHandlerAccept) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa9(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa10(in *jlexer.Lexer, out *FileHandler) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "action":
			out.Action = string(in.String())
		case "accepts":
			if in.IsNull() {
				in.Skip()
				out.Accepts = nil
			} else {
				in.Delim('[')
				if out.Accepts == nil {
					if !in.IsDelim(']') {
						out.Accepts = make([]*FileHandlerAccept, 0, 8)
					} else {
						out.Accepts = []*FileHandlerAccept{}
					}
				} else {
					out.Accepts = (out.Accepts)[:0]
				}
				for !in.IsDelim(']') {
					var v13 *FileHandlerAccept
					if in.IsNull() {
						in.Skip()
						v13 = nil
					} else {
						if v13 == nil {
							v13 = new(FileHandlerAccept)
						}
						(*v13).UnmarshalEasyJSON(in)
					}
					out.Accepts = append(out.Accepts, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "displayName":
			out.DisplayName = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa10(out *jwriter.Writer, in FileHandler) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"action\":"
		out.RawString(prefix[1:])
		out.String(string(in.Action))
	}
	{
		const prefix string = ",\"accepts\":"
		out.RawString(prefix)
		if in.Accepts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Accepts {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"displayName\":"
		out.RawString(prefix)
		out.String(string(in.DisplayName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileHandler) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileHandler) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileHandler) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileHandler) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa10(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa11(in *jlexer.Lexer, out *ChangeAppUserSettingsParams) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestId":
			out.ManifestID = string(in.String())
		case "linkCapturing":
			out.LinkCapturing = bool(in.Bool())
		case "displayMode":
			(out.DisplayMode).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa11(out *jwriter.Writer, in ChangeAppUserSettingsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	if in.LinkCapturing {
		const prefix string = ",\"linkCapturing\":"
		out.RawString(prefix)
		out.Bool(bool(in.LinkCapturing))
	}
	if in.DisplayMode != "" {
		const prefix string = ",\"displayMode\":"
		out.RawString(prefix)
		(in.DisplayMode).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChangeAppUserSettingsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChangeAppUserSettingsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChangeAppUserSettingsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChangeAppUserSettingsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa11(l, v)
}
