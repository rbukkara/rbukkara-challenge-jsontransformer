package jsontransformer

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

type TestIO struct {
	input  string
	output string
}

func TestJsonTransformer(t *testing.T) {
	tests := []TestIO{
		{
			input: `{
				"map_1": {
					"M": {
						"bool_1": {
							"BOOL": "truthy"
						},
						"null_1": {
							"NULL": "true"
						},
						"list_1": {
							"L": [
								{
									"S": ""
								},
								{
									"N": "011"
								},
								{
									"N": "5215s"
								},
								{
									"BOOL": "f"
								},
								{
									"NULL": "0"
								}
								]
							}
						}
					}
				}`,
			output: `{"map_1":{"list_1":[11,false],"null_1":null}}`,
		},
		{
			input: `{
				"map_1": {
					"M": {
						"bool_1": {
							"BOOL": "truthy"
						},
						"null_1": {
							"NULL": "true"
						},
						"list_1": {
							"L": [
								{
									"S": "test"
								},
								{
									"N": "011"
								},
								{
									"N": "5215s"
								},
								{
									"BOOL": "f"
								},
								{
									"NULL": "1"
								},
								{"list_1_inner": {
									"L": [
										{
											"S":"inner"
										}
									]
								}}
							]
						}
					}
				}}`,
			output: `{"map_1":{"list_1":["test",11,false],"null_1":null}}`,
		},
		{
			input: `{
				"number_1": {
					"N": "1.50"
				},
				" string_1": {
					"S": "784498 "
				},
				"string_2": {
					"S": " 2014-07-16T20:55:46Z"
				},
				"bool_1": {
					"BOOL ": " true"
				},
				"null_1": {
					"NULL": "true"
				},
				"list_1": {
					"L": [
						{
							"S": ""
						},
						{
							"N": "011"
						},
						{
							"N": "5215s"
						},
						{
							"BOOL": "f"
						},
						{
							"NULL": "0"
						}
					]
				},
				"list_2": {
					"L": "noop"
				},
				"list_3": {
					"L": [
					"noop"
					]
				},
				"": {
					"S": "noop"
				}
				}`,
			output: `{"bool_1":true,"list_1":[11,false],"number_1":1.5,"string_1":"784498","string_2":1405544146}`,
		},
		{
			input: `{
				"number_1": {
					" N": " 1.50"
				},
				"string_1": {
					"S": "784498 "
				},
				"string_2": {
					"S": "2014-07-16T20:55:46Z"
				},
				"map_1": {
					"M": {
					"bool_1": {
						"BOOL": "truthy"
					},
					"null_1": {
						"NULL ": "true"
					},
					"list_1": {
						"L": [
						{
							"S": ""
						},
						{
							"N": "011"
						},
						{
							"N": "5215s"
						},
						{
							"BOOL": "f"
						},
						{
							"NULL": "0"
						}
						]
					}
					}
				},
				"list_2": {
					"L": "noop"
				},
				"list_3": {
					"L": [
					"noop"
					]
				},
				"": {
					"S": "noop"
				}
				}`,
			output: `{"map_1":{"list_1":[11,false],"null_1":null},"number_1":1.5,"string_1":"784498","string_2":1405544146}`,
		},
	}
	for _, tc := range tests {
		reader := strings.NewReader(tc.input)
		got, _ := Transform(reader)

		buf := new(bytes.Buffer)
		buf.ReadFrom(got)
		outputString := buf.String()

		// return a string
		if !reflect.DeepEqual(tc.output, outputString) {
			t.Errorf("expected: %v, got: %v", tc.output, got)
		}
	}
}
