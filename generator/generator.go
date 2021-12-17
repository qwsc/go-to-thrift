package generator

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

var TypeMapping = map[string]string{
	"bool":    "bool",
	"int":     "i64",
	"int8":    "byte",
	"int16":   "i16",
	"int32":   "i32",
	"int64":   "i64",
	"uint":    "i64",
	"uint8":   "byte",
	"uint16":  "i16",
	"uint32":  "i32",
	"uint64":  "i64",
	"byte":    "byte",
	"float32": "double",
	"float64": "double",
	"string":  "string",
}

type Generator struct {
	file  *ast.File
	enums map[string][][2]string
}

func NewGenerator(file *ast.File) *Generator {
	return &Generator{
		file:  file,
		enums: map[string][][2]string{},
	}
}

func (g *Generator) generateIndent(indent *ast.Ident) string {
	if res, ok := TypeMapping[indent.Name]; ok {
		return res
	}
	return indent.Name
}

func (g *Generator) generateFieldType(fieldType ast.Expr, root bool) string {
	switch expr := fieldType.(type) {
	case *ast.Ident:
		return g.generateIndent(expr)
	case *ast.StarExpr:
		if root {
			return "optional " + g.generateFieldType(expr.X, false)
		} else {
			return g.generateFieldType(expr.X, false)
		}
	case *ast.ArrayType:
		return "list<" + g.generateFieldType(expr.Elt, false) + ">"
	case *ast.MapType:
		return "map<" + g.generateFieldType(expr.Key, false) + "," + g.generateFieldType(expr.Value, false) + ">"
	}
	return ""
}

func (g *Generator) generateField(field *ast.Field) []string {
	var typeStr string
	typeStr = g.generateFieldType(field.Type, true)

	var goTag string
	if field.Tag != nil {
		goTag = strings.Trim(field.Tag.Value, "`")
	}

	var comment string
	if field.Comment != nil {
		comment = strings.TrimSpace(field.Comment.Text())
	}

	var result []string
	for _, name := range field.Names {
		builder := strings.Builder{}
		builder.WriteString(typeStr + " ")
		builder.WriteString(name.Name + " ")
		if len(goTag) > 0 {
			builder.WriteString("(go.tag=" + strconv.Quote(goTag) + ") ")
		}
		if len(comment) > 0 {
			builder.WriteString("// " + comment)
		}
		result = append(result, builder.String())
	}
	return result
}

func (g *Generator) generateStructType(structType *ast.StructType) string {
	builder := strings.Builder{}
	builder.WriteString("{\n")
	var fieldIndex = 1
	for _, field := range structType.Fields.List {
		for _, fieldStr := range g.generateField(field) {
			builder.WriteString("    " + strconv.Itoa(fieldIndex) + ": " + fieldStr + "\n")
			fieldIndex++
		}
	}
	builder.WriteString("}\n\n")
	return builder.String()
}

func (g *Generator) generateTypeSpec(typeSpec *ast.TypeSpec) string {
	builder := strings.Builder{}
	switch expr := typeSpec.Type.(type) {
	case *ast.StructType:
		builder.WriteString("struct " + typeSpec.Name.Name + " " + g.generateStructType(expr))
	case *ast.Ident:
		// TODO
	}
	return builder.String()
}

func (g *Generator) generateValueSpec(valueSpec *ast.ValueSpec) {
	valueType, ok := valueSpec.Type.(*ast.Ident)
	if !ok {
		return
	}
	if _, ok := TypeMapping[valueType.Name]; ok {
		return
	}
	for i, name := range valueSpec.Names {
		value := valueSpec.Values[i]
		if value, ok := value.(*ast.BasicLit); ok {
			if value, err := strconv.Atoi(value.Value); err == nil {
				g.enums[valueType.Name] = append(g.enums[valueType.Name], [2]string{name.Name, strconv.Itoa(value)})
			}
		}
	}
}

func (g *Generator) generateGenDecl(genDecl *ast.GenDecl) string {
	builder := strings.Builder{}
	switch genDecl.Tok {
	case token.TYPE:
		for _, spec := range genDecl.Specs {
			switch spec := spec.(type) {
			case *ast.TypeSpec:
				builder.WriteString(g.generateTypeSpec(spec))
			}
		}
		return builder.String()
	case token.CONST:
		for _, spec := range genDecl.Specs {
			switch spec := spec.(type) {
			case *ast.ValueSpec:
				g.generateValueSpec(spec)
			}
		}
	}
	return ""
}

func (g *Generator) generateDecl(decl ast.Decl) string {
	switch decl := decl.(type) {
	case *ast.GenDecl:
		return g.generateGenDecl(decl)
	}
	return ""
}

func (g *Generator) Generate() string {
	structBuilder := strings.Builder{}
	for _, decl := range g.file.Decls {
		structBuilder.WriteString(g.generateDecl(decl))
	}
	enumBuilder := strings.Builder{}
	for enumName, kvs := range g.enums {
		enumBuilder.WriteString("enum " + enumName + " {\n")
		for _, kv := range kvs {
			enumBuilder.WriteString("    " + kv[0] + " = " + kv[1] + "\n")
		}
		enumBuilder.WriteString("}\n\n")
	}
	return enumBuilder.String() + structBuilder.String()
}
