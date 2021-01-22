package gotemplate

import (
	"testing"

	"github.com/georgemac/truce"
	"github.com/georgemac/truce/pkg/http"
	"gotest.tools/v3/assert"
)

func TestArgs(t *testing.T) {
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
			v := args(truce.Function{
				Arguments: tc.in,
			})
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
			out: "do(ctxt context.Context, v0 string) (err error)",
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
			out: "do(ctxt context.Context, v0 string, v1 int) (err error)",
		},
		{
			name: "return value",
			in: truce.Function{
				Name: "do",
				Arguments: []truce.Field{
					{Name: "a", Type: "string"},
				},
				Return: truce.Field{
					Name: "x",
					Type: "string",
				},
			},
			out: "do(ctxt context.Context, v0 string) (rtn string, err error)",
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
	v1 := method(&http.Function{Method: "GET"})
	assert.Equal(t, v1, "Get")

	v2 := method(&http.Function{Method: "get"})
	assert.Equal(t, v2, "Get")
}

func TestPath(t *testing.T) {
	testcases := []struct {
		name string
		in   http.Path
		out  string
	}{
		{
			name: "no variables",
			in: http.Path{
				{Type: "static", Value: "a"},
			},
			out: `"/a"`,
		},
		{
			name: "single variable",
			in: http.Path{
				{Var: "x", Type: "variable", Value: "a"},
			},
			out: `fmt.Sprintf("/%v", x)`,
		},
		{
			name: "multiple variables",
			in: http.Path{
				{Var: "x", Type: "variable", Value: "a"},
				{Var: "y", Type: "variable", Value: "b"},
				{Var: "z", Type: "variable", Value: "c"},
			},
			out: `fmt.Sprintf("/%v/%v/%v", x, y, z)`,
		},
		{
			name: "variables and static elements",
			in: http.Path{
				{Value: "api", Type: "static"},
				{Var: "x", Type: "variable", Value: "a"},
				{Var: "y", Type: "variable", Value: "b"},
				{Value: "private", Type: "static"},
				{Var: "z", Type: "variable", Value: "c"},
			},
			out: `fmt.Sprintf("/api/%v/%v/private/%v", x, y, z)`,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := path(&http.Function{Path: tc.in})
			assert.Equal(t, v, tc.out)
		})
	}
}