package input

type Gamma int32

const (
	GammaOne  Gamma = 1
	GammaTwo  Gamma = 2
	GamaThree Gamma = 3
)

type Beta struct {
	Bool    bool    `json:"bool_field,omitempty"`   // bool field
	Int32   int32   `json:"int_32_field,omitempty"` // int32 field
	Float64 float64 `json:"float_64,omitempty"`     // float64 field
	String  string  `json:"string,omitempty"`       // string field
}

type Alpha struct {
	Bool    bool    `json:"bool_field,omitempty"`
	Int32   int32   `json:"int_32_field,omitempty"`
	Float64 float64 `json:"float_64,omitempty"`
	String  string  `json:"string,omitempty"`

	BoolPointer    *bool    `json:"bool_pointer,omitempty"`
	Int32Pointer   *int32   `json:"int_32_pointer,omitempty"`
	Float64Pointer *float64 `json:"float_64_pointer,omitempty"`
	StringPointer  *string  `json:"string_pointer,omitempty"`

	BoolSlice    []bool    `json:"bool_slice,omitempty"`
	Int32Slice   []int32   `json:"int_32_slice,omitempty"`
	Float64Slice []float64 `json:"float_64_slice,omitempty"`
	StringSlice  []string  `json:"string_slice,omitempty"`

	Beta         Beta             `json:"beta,omitempty"`
	BetaPointer  *Beta            `json:"beta_pointer,omitempty"`
	BetaSlice    []*Beta          `json:"beta_slice,omitempty"`
	StringToBeta map[string]*Beta `json:"string_to_beta,omitempty"`

	Gamma Gamma `json:"gamma,omitempty"`
}
