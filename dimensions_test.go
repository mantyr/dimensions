package dimensions

import (
    "testing"
)

type ParseTest struct {
    source        string
    value         string
    value_int     int
    value_int64   int64
    value_float64 float64
    value_type    string
}

var weightTests = []ParseTest{
    {"10 гр.",         "10",    10, 10, 10,    "GR"},
    {"10 gr.",         "10",    10, 10, 10,    "GR"},
    {"10",             "10",    10, 10, 10,    "DEFAULT_TYPE"},
    {"10 kg.",         "10",    10, 10, 10,    "KG"},
    {"10 кг.",         "10",    10, 10, 10,    "KG"},
    {"",               "0",      0,  0,  0,    "DEFAULT_TYPE"},
}

func TestWeight(t *testing.T) {
    p := NewDimension()
    p.SetDimensionWeight()
    p.SetDefaultType("DEFAULT_TYPE")

    for _, test := range weightTests {
        p.Parse(test.source)
        if p.Get() != test.value {
            t.Errorf("Error value, %q, %q", test.source, p.Get())
        }
        if p.GetInt() != test.value_int {
            t.Errorf("Error int, %q, %q", test.source, p.GetInt())
        }
        if p.GetInt64() != test.value_int64 {
            t.Errorf("Error int64, %q, %q", test.source, p.GetInt64())
        }
        if p.GetFloat64() != test.value_float64 {
            t.Errorf("Error float64, %q, %q", test.source, p.GetFloat64())
        }
        if p.GetType() != test.value_type {
            t.Errorf("Error type, %q, %q", test.source, p.GetType())
        }
    }
}

var widthTests = []ParseTest{
    {"10 километров",  "10",    10, 10, 10,    "KM"},
    {"10 км.",         "10",    10, 10, 10,    "KM"},
    {"10",             "10",    10, 10, 10,    "DEFAULT_TYPE"},
    {"10 метров",      "10",    10, 10, 10,    "M"},
    {"10 Метры",       "10",    10, 10, 10,    "M"},
    {"10 м",           "10",    10, 10, 10,    "M"},
    {"12.52 м",        "12.52", 13, 13, 12.52, "M"},
}

func TestWidth(t *testing.T) {
    p := NewDimension()
    p.SetDimensionWidth()
    p.SetDefaultType("DEFAULT_TYPE")

    for _, test := range widthTests {
        p.Parse(test.source)
        if p.Get() != test.value {
            t.Errorf("Error value, %q, %q, %q", test.source, p.Get(), test.value)
        }
        if p.GetInt() != test.value_int {
            t.Errorf("Error int, %q, %i, %i", test.source, p.GetInt(), test.value_int)
        }
        if p.GetInt64() != test.value_int64 {
            t.Errorf("Error int64, %q, %i, %i", test.source, p.GetInt64(), test.value_int64)
        }
        if p.GetFloat64() != test.value_float64 {
            t.Errorf("Error float64, %q, %f, %f", test.source, p.GetFloat64(), test.value_float64)
        }
        if p.GetType() != test.value_type {
            t.Errorf("Error type, %q, %q", test.source, p.GetType())
        }
    }
}

