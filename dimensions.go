package dimensions

import (
    "strconv"
    "strings"
    "regexp"
    "math/big"
)

var regexp_numbers = regexp.MustCompile("[.,0-9]+")
func init() {
    regexp_numbers = regexp.MustCompile("[,.0-9]+")
}

func NewDimension() (d *Dimension) {
    d = new(Dimension)
    d.Value_rat = new(big.Rat)
    return d
}

func (d *Dimension) SetDimensionType(name string) *Dimension {
    d.dimension_type = name
    return d
}

func (d *Dimension) SetDimensionWeight() *Dimension {
    d.SetDimensionType("weight")
    return d
}

func (d *Dimension) SetDimensionWidth() *Dimension {
    d.SetDimensionType("width")
    return d
}

func (d *Dimension) SetDefaultType(name string) *Dimension {
    d.default_type = name
    return d
}

func (d *Dimension) Parse(value string) *Dimension {
    d.Value_source = value
    d.Value        = d.parse_value()
    d.Value_type   = d.parse_type()

    d.Value_rat.SetString(d.Value)
    return d
}


func (d *Dimension) parse_value() string {
    v := strings.Join(regexp_numbers.FindAllString(d.Value_source, -1), "")
    v  = strings.Replace(v, ",", ".", -1)

    c := strings.IndexRune(v, '.')
    if c > -1 {
        c2 := strings.IndexRune(v[c+1:], '.')
        if c2 > -1 {
            v = v[:c+c2+2]
        }
    }

    if v[len(v)-1:] == "." {
        v = v[:len(v)-1]
    }
    return v
}

func (d *Dimension) parse_type() string {
    var values []SearchType

    switch d.dimension_type {
        case "weight":
            values = WeightTypes
        case "width":
            values = WidthTypes
        default:
            return d.default_type
    }

    for _, value := range values {
        if strings.Index(strings.ToLower(d.Value_source), value.Search) > -1 {
            return value.Result
        }
    }
    return d.default_type
}

func (d *Dimension) String() string {
    return d.Get()+" "+d.Value_type
}

func (d *Dimension) Get() string {
    value := d.Value_rat.FloatString(2)
    if value[len(value)-2:] == "00" {
        return value[:len(value)-3]
    }
    return value
}

func (d *Dimension) GetFloatString(prec int) string {
    return d.Value_rat.FloatString(prec)
}

func (d *Dimension) GetInt() int {
    value := d.GetFloatString(0)

    v, err := strconv.Atoi(value)
    if err != nil {
        return 0
    }
    return v
}

func (d *Dimension) GetInt64() int64 {
    value := d.GetFloatString(0)

    v, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
        return 0
    }
    return v
}

func (d *Dimension) GetFloat64() float64 {
    v, _ := d.Value_rat.Float64()
    return v
}

func (d *Dimension) GetTypeDimension() string {
    return d.dimension_type
}

func (d *Dimension) GetType() string {
    return d.Value_type
}
