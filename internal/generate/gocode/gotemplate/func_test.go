package gotemplate

import (
	"testing"

	"github.com/TruceRPC/truce"
	"gotest.tools/v3/assert"
)

func TestPosArgs(t *testing.T) {
	testcases := []struct {
		name string
		in   []truce.Field
		out  string
	}{
		{
			name: "single field",
			in: []truce.Field{
				{Name: "a", Type: "string"},
			},
			out: "v0",
		},
		{
			name: "multiple fields",
			in: []truce.Field{
				{Name: "a", Type: "string"},
				{Name: "b", Type: "string"},
			},
			out: "v0, v1",
		},
		{
			name: "no fields",
			in:   []truce.Field{},
			out:  "",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := posArgs(tc.in)
			assert.Equal(t, v, tc.out)
		})
	}
}

func TestName(t *testing.T) {
	testcases := []struct {
		name string
		in   truce.Field
		out  string
	}{
		{
			name: "simple",
			in:   truce.Field{Name: "a"},
			out:  "A",
		},
		{
			name: "multi-segment",
			in:   truce.Field{Name: "a_thing_that_is_endless"},
			out:  "AThingThatIsEndless",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := name(tc.in)
			assert.Equal(t, v, tc.out)
		})
	}
}

func TestGoType(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{
			in:  "string",
			out: "string",
		},
		{
			in:  "int",
			out: "int64",
		},
		{
			in:  "object",
			out: "map[string]interface{}",
		},
		{
			in:  "timestamp",
			out: "time.Time",
		},
		{
			in:  "*timestamp",
			out: "*time.Time",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, goType(tc.in), tc.out)
		})
	}
}

func TestSignature(t *testing.T) {
	testcases := []struct {
		name string
		in   truce.Function
		out  string
	}{
		{
			name: "single argument",
			in: truce.Function{
				Name: "do",
				Arguments: []truce.Field{
					{Name: "a", Type: "string"},
				},
			},
			out: "do(ctxt context.Context, a string) (error)",
		},
		{
			name: "multiple arguments",
			in: truce.Function{
				Name: "do",
				Arguments: []truce.Field{
					{Name: "a", Type: "string"},
					{Name: "b", Type: "int"},
				},
			},
			out: "do(ctxt context.Context, a string, b int64) (error)",
		},
		{
			name: "return value",
			in: truce.Function{
				Name: "do",
				Arguments: []truce.Field{
					{Name: "a", Type: "string"},
				},
				Return: truce.OptionalField{
					Present: true,
					Field: truce.Field{
						Name: "x",
						Type: "string",
					},
				},
			},
			out: "do(ctxt context.Context, a string) (string, error)",
		},
		{
			name: "no return value",
			in: truce.Function{
				Name: "do",
				Arguments: []truce.Field{
					{Name: "a", Type: "string"},
				},
			},
			out: "do(ctxt context.Context, a string) (error)",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := signature(tc.in)
			assert.Equal(t, v, tc.out)
		})
	}
}

func TestFlags(t *testing.T) {
	v := tags(truce.Field{Name: "a_field"})
	assert.Equal(t, v, "`json:\"a_field\"`")
}

func TestMethod(t *testing.T) {
	v1 := method(&Function{Method: "GET"})
	assert.Equal(t, v1, "Get")

	v2 := method(&Function{Method: "get"})
	assert.Equal(t, v2, "Get")
}

func TestPath(t *testing.T) {
	testcases := []struct {
		name string
		in   Path
		out  string
	}{
		{
			name: "no variables",
			in: Path{
				{Type: "static", Value: "a"},
			},
			out: `"/a"`,
		},
		{
			name: "single variable",
			in: Path{
				{VarPos: 0, Var: "x", Type: "variable", Value: "a"},
			},
			out: `fmt.Sprintf("/%v", v0)`,
		},
		{
			name: "multiple variables",
			in: Path{
				{VarPos: 0, Var: "x", Type: "variable", Value: "a"},
				{VarPos: 1, Var: "y", Type: "variable", Value: "b"},
				{VarPos: 2, Var: "z", Type: "variable", Value: "c"},
			},
			out: `fmt.Sprintf("/%v/%v/%v", v0, v1, v2)`,
		},
		{
			name: "variables and static elements",
			in: Path{
				{Value: "api", Type: "static"},
				{VarPos: 0, Var: "x", Type: "variable", Value: "a"},
				{VarPos: 1, Var: "y", Type: "variable", Value: "b"},
				{Value: "private", Type: "static"},
				{VarPos: 2, Var: "z", Type: "variable", Value: "c"},
			},
			out: `fmt.Sprintf("/api/%v/%v/private/%v", v0, v1, v2)`,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := pathJoin(tc.in)
			assert.Equal(t, v, tc.out)
		})
	}
}

func TestErrorFmt(t *testing.T) {
	testcases := []struct {
		name string
		in   truce.Type
		out  string
	}{
		{
			name: "single field",
			in: truce.Type{
				Fields: map[string]truce.Field{
					"aoeu": {Name: "x", Type: "string"},
				},
			},
			out: `"error: x=%q", e.X`,
		},
		{
			name: "multiple fields",
			in: truce.Type{
				Fields: map[string]truce.Field{
					"x": {Name: "x", Type: "string"},
					"y": {Name: "y", Type: "string"},
				},
			},
			out: `"error: x=%q y=%q", e.X, e.Y`,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := errorFmt(tc.in)
			assert.Equal(t, v, tc.out)
		})
	}
}

func TestBacktick(t *testing.T) {
	assert.Equal(t, backtick("hello"), "`hello`")
}
