package gotemplate

import (
	"fmt"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/TruceRPC/truce"
)

var tmplFuncs = template.FuncMap{
	"backtick":            backtick,
	"callUnexported":      callUnexported,
	"errorFmt":            errorFmt,
	"goType":              goType,
	"method":              method,
	"name":                name,
	"pathJoin":            pathJoin,
	"posArgs":             posArgs,
	"signature":           signature,
	"sortQueryParams":     sortQueryParams,
	"stringHasPrefix":     strings.HasPrefix,
	"tags":                tags,
	"unexportedSignature": unexportedSignature,
}

func unexported(v string) string {
	if v == "" {
		return v
	}
	return string(unicode.ToLower(rune(v[0]))) + v[1:]
}

func name(f truce.Field) (v string) {
	parts := strings.Split(f.Name, "_")
	for _, p := range parts {
		v += strings.Title(p)
	}
	return
}

func goType(typ string) string {
	switch typ {
	case "int":
		return "int64"
	case "object":
		return "map[string]interface{}"
	case "timestamp", "*timestamp":
		if typ[0] == '*' {
			return "*time.Time"
		}

		return "time.Time"
	default:
		return typ
	}
}

func tags(f truce.Field) string {
	return fmt.Sprintf("`json:%q`", f.Name)
}

func signature(f truce.Function) string {
	builder := &strings.Builder{}
	fmt.Fprintf(builder, "%s(ctxt context.Context, ", f.Name)
	for i, arg := range f.Arguments {
		if i > 0 {
			builder.WriteString(", ")
		}
		fmt.Fprintf(builder, "%s %s", arg.Name, goType(arg.Type))
	}
	builder.WriteString(") ")
	builder.WriteString(signatureReturn(f))
	return builder.String()
}

func unexportedSignature(f truce.Function) string {
	builder := &strings.Builder{}
	fmt.Fprintf(builder, "%s(ctxt context.Context, ", unexported(f.Name))
	for i, arg := range f.Arguments {
		if i > 0 {
			builder.WriteString(", ")
		}
		fmt.Fprintf(builder, "v%d %s", i, goType(arg.Type))
	}
	builder.WriteString(") ")
	builder.WriteString(signatureReturn(f))
	return builder.String()
}

func signatureReturn(f truce.Function) string {
	builder := &strings.Builder{}
	builder.WriteString("(")
	if rtn := f.Return; rtn.Present && rtn.Name != "" {
		fmt.Fprintf(builder, "%s, ", goType(rtn.Type))
	}
	builder.WriteString("error)")
	return builder.String()
}

func callUnexported(f truce.Function) string {
	builder := &strings.Builder{}
	fmt.Fprintf(builder, "%s(ctxt, ", unexported(f.Name))
	for i, arg := range f.Arguments {
		if i > 0 {
			builder.WriteString(", ")
		}
		fmt.Fprintf(builder, "%s", arg.Name)
	}
	builder.WriteString(")")
	return builder.String()
}

func pathJoin(path Path) string {
	if len(path) == 0 {
		return fmt.Sprintf("%q", path)
	}

	var hasVariables bool
	for _, elem := range path {
		if elem.Type == "variable" {
			hasVariables = true
		}
	}

	if hasVariables {
		return fmt.Sprintf(`fmt.Sprintf(%q, %s)`, path.FmtString(), path.ArgString())
	}
	return fmt.Sprintf(`%q`, path.FmtString())
}

func posArgs(fields []truce.Field) (v string) {
	for i := range fields {
		if i > 0 {
			v += ", "
		}
		v += fmt.Sprintf("v%d", i)
	}
	return
}

func method(b *Function) string {
	return strings.Title(strings.ToLower(b.Method))
}

func errorFmt(t truce.Type) string {
	var (
		i              int
		fmtStr, argStr string
	)
	for _, field := range sortedFields(t.Fields) {
		if i > 0 {
			fmtStr += " "
			argStr += ", "
		}

		fmtStr += fmt.Sprintf("%s=%%q", field.Name)
		argStr += fmt.Sprintf("e.%s", name(field))
		i++
	}

	return fmt.Sprintf(`"error: %s", %s`, fmtStr, argStr)
}

func backtick(v string) string {
	return "`" + v + "`"
}

func sortedFields(m map[string]truce.Field) []truce.Field {
	var l []truce.Field
	for _, f := range m {
		l = append(l, f)
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].Name < l[j].Name
	})
	return l
}

func sortQueryParams(m map[string]QueryParam) []QueryParam {
	var l []QueryParam
	for _, f := range m {
		l = append(l, f)
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].Pos < l[j].Pos
	})
	return l
}
