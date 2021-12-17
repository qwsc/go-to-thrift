enum Gamma {
    GammaOne = 1
    GammaTwo = 2
    GamaThree = 3
}

struct Beta {
    1: bool Bool (go.tag="json:\"bool_field,omitempty\"") // bool field
    2: i32 Int32 (go.tag="json:\"int_32_field,omitempty\"") // int32 field
    3: double Float64 (go.tag="json:\"float_64,omitempty\"") // float64 field
    4: string String (go.tag="json:\"string,omitempty\"") // string field
}

struct Alpha {
    1: bool Bool (go.tag="json:\"bool_field,omitempty\"") 
    2: i32 Int32 (go.tag="json:\"int_32_field,omitempty\"") 
    3: double Float64 (go.tag="json:\"float_64,omitempty\"") 
    4: string String (go.tag="json:\"string,omitempty\"") 
    5: optional bool BoolPointer (go.tag="json:\"bool_pointer,omitempty\"") 
    6: optional i32 Int32Pointer (go.tag="json:\"int_32_pointer,omitempty\"") 
    7: optional double Float64Pointer (go.tag="json:\"float_64_pointer,omitempty\"") 
    8: optional string StringPointer (go.tag="json:\"string_pointer,omitempty\"") 
    9: list<bool> BoolSlice (go.tag="json:\"bool_slice,omitempty\"") 
    10: list<i32> Int32Slice (go.tag="json:\"int_32_slice,omitempty\"") 
    11: list<double> Float64Slice (go.tag="json:\"float_64_slice,omitempty\"") 
    12: list<string> StringSlice (go.tag="json:\"string_slice,omitempty\"") 
    13: Beta Beta (go.tag="json:\"beta,omitempty\"") 
    14: optional Beta BetaPointer (go.tag="json:\"beta_pointer,omitempty\"") 
    15: list<Beta> BetaSlice (go.tag="json:\"beta_slice,omitempty\"") 
    16: map<string,Beta> StringToBeta (go.tag="json:\"string_to_beta,omitempty\"") 
    17: Gamma Gamma (go.tag="json:\"gamma,omitempty\"") 
}

