package slices

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tcase struct {
	test     string
	input    any
	expected any
}

var s = []struct {
	Name string
}{
	{
		Name: "hello earth",
	},
	{
		Name: "hello mars",
	},
	{
		Name: "hello mars",
	},
}

func TestFind(t *testing.T) {

	testCases := []tcase{
		{
			test: "it should get hello mars",
			input: func(item struct {
				Name string
			}, index int) bool {

				return item.Name == "hello mars"
			},
			expected: &struct{ Name string }{Name: "hello mars"},
		},

		{
			test: "it should get nil pointer",
			input: func(item struct {
				Name string
			}, index int) bool {

				return false
			},
			expected: (*struct{ Name string })(nil),
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.test, func(t *testing.T) {
			asserts := assert.New(t)

			ex := Find[struct{ Name string }](&s, tcase.input.(func(item struct{ Name string }, index int) bool))

			asserts.Equal(tcase.expected, ex)
		})
	}

}

func TestFilter(t *testing.T) {

	testCases := []tcase{
		{
			test: "it should get hello mars",
			input: func(item struct {
				Name string
			}, index int) bool {
				return strings.Contains(item.Name, "hello")
			},
			expected: []struct{ Name string }{
				{
					Name: "hello earth",
				},
				{
					Name: "hello mars",
				},
				{
					Name: "hello mars",
				},
			},
		},

		{
			test: "it should get nil pointer",
			input: func(item struct {
				Name string
			}, index int) bool {
				return strings.Contains(item.Name, "rello")
			},
			expected: []struct{ Name string }{},
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.test, func(t *testing.T) {
			asserts := assert.New(t)

			ex := Filter[struct{ Name string }](&s, tcase.input.(func(item struct{ Name string }, index int) bool))

			asserts.Equal(tcase.expected, ex)
		})
	}

}

func TestMap(t *testing.T) {

	type Mapped struct {
		Hello string
	}

	testCases := []tcase{
		{
			test: "it should get slice of mapped for each element",
			input: func(item struct {
				Name string
			}, index int) Mapped {
				return Mapped{
					Hello: item.Name,
				}
			},
			expected: []Mapped{
				{
					Hello: "hello earth",
				},
				{
					Hello: "hello mars",
				},
				{
					Hello: "hello mars",
				},
			},
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.test, func(t *testing.T) {
			asserts := assert.New(t)

			ex := Map(&s, tcase.input.(func(item struct{ Name string }, index int) Mapped))

			asserts.Equal(tcase.expected, ex)
		})
	}

}

func TestForeach(t *testing.T) {

	testCases := []tcase{
		{
			test:     "it should return nil",
			input:    nil,
			expected: nil,
		},
		{
			test:     "it should return error",
			input:    errors.New("error"),
			expected: errors.New("error"),
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.test, func(t *testing.T) {
			err := Foreach(&s, func(el struct{ Name string }, index int) error {
				if tcase.input != nil {
					assert.Equal(t, 0, index)
					return tcase.input.(error)
				}
				assert.GreaterOrEqual(t, 4, index)
				return nil
			})

			assert.Equal(t, tcase.expected, err)
		})
	}

}
