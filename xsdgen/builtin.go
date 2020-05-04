package xsdgen

import (
	"go/ast"

	"github.com/CognitoIQ/go-xml/xsd"
)

//func builtinExpr(b xsd.Builtin) ast.Expr {
//	if int(b) > len(builtinTbl) || b < 0 {
//		return nil
//	}
//	return builtinTbl[b]
//}

func builtinExpr(b xsd.Builtin, pointer bool) ast.Expr {
	goLangTypeInfo, exists := goLangTypeMappings[int(b)]
	if !exists {
		return nil
	}

	var golangType string
	if pointer {
		golangType = "*" + goLangTypeInfo.goLangType
	} else {
		golangType = goLangTypeInfo.goLangType
	}

	if goLangTypeInfo.isarray {
		return &ast.ArrayType{Elt: &ast.Ident{Name: golangType}}
	}
	return &ast.Ident{Name: golangType}

}

// Returns true if t is an xsd.Builtin that is not trivially mapped to a
// builtin Go type; it requires additional marshal/unmarshal methods.
func nonTrivialBuiltin(t xsd.Type) bool {
	b, ok := t.(xsd.Builtin)
	if !ok {
		return false
	}
	switch b {
	case xsd.Base64Binary, xsd.HexBinary,
		xsd.Date, xsd.Time, xsd.DateTime,
		xsd.GDay, xsd.GMonth, xsd.GMonthDay, xsd.GYear, xsd.GYearMonth:
		return true
	}
	return false
}

type goLangTypes struct {
	goLangType string
	isarray    bool
}

//the 45 built in types
var goLangTypeMappings = map[int]goLangTypes{
	int(xsd.AnyType):      {goLangType: "string"},
	int(xsd.ENTITIES):     {goLangType: "string", isarray: true},
	int(xsd.ENTITY):       {goLangType: "string"},
	int(xsd.ID):           {goLangType: "string"},
	int(xsd.IDREF):        {goLangType: "string"},
	int(xsd.IDREFS):       {goLangType: "string", isarray: true},
	int(xsd.NCName):       {goLangType: "string"},
	int(xsd.NMTOKEN):      {goLangType: "string"},
	int(xsd.NMTOKENS):     {goLangType: "string", isarray: true},
	int(xsd.NOTATION):     {goLangType: "string", isarray: true},
	int(xsd.Name):         {goLangType: "string"},
	int(xsd.QName):        {goLangType: "xml.Name"},
	int(xsd.AnyURI):       {goLangType: "string"},
	int(xsd.Base64Binary): {goLangType: "byte", isarray: true},
	int(xsd.Boolean):      {goLangType: "bool"},
	int(xsd.Byte):         {goLangType: "byte"},
	int(xsd.Date):         {goLangType: "time.Time"},

	//int(xsd.DateTime) : {goLangType : "time.Time"},
	int(xsd.DateTime): {goLangType: "time.Time"},

	int(xsd.Decimal): {goLangType: "float64"},
	int(xsd.Double):  {goLangType: "float64"},
	// the "duration" built-in is especially broken, so we
	// don't parse it at all.
	int(xsd.Duration):           {goLangType: "string"},
	int(xsd.Float):              {goLangType: "float32"},
	int(xsd.GDay):               {goLangType: "time.Time"},
	int(xsd.GMonth):             {goLangType: "time.Time"},
	int(xsd.GMonthDay):          {goLangType: "time.Time"},
	int(xsd.GYear):              {goLangType: "time.Time"},
	int(xsd.GYearMonth):         {goLangType: "time.Time"},
	int(xsd.HexBinary):          {goLangType: "byte", isarray: true},
	int(xsd.Int):                {goLangType: "int"},
	int(xsd.Integer):            {goLangType: "int"},
	int(xsd.Language):           {goLangType: "string"},
	int(xsd.Long):               {goLangType: "int64"},
	int(xsd.NegativeInteger):    {goLangType: "int"},
	int(xsd.NonNegativeInteger): {goLangType: "int"},
	int(xsd.NormalizedString):   {goLangType: "string"},
	int(xsd.NonPositiveInteger): {goLangType: "int"},
	int(xsd.PositiveInteger):    {goLangType: "int"},
	int(xsd.Short):              {goLangType: "int"},
	int(xsd.String):             {goLangType: "string"},
	int(xsd.Time):               {goLangType: "time.Time"},
	int(xsd.Token):              {goLangType: "string"},
	int(xsd.UnsignedByte):       {goLangType: "byte"},
	int(xsd.UnsignedInt):        {goLangType: "uint"},
	int(xsd.UnsignedLong):       {goLangType: "uint64"},
	int(xsd.UnsignedShort):      {goLangType: "uint"},
}

// The 45 built-in types of the XSD schema
//var builtinTbl = []ast.Expr{
//	xsd.AnyType:      &ast.Ident{Name: "string"},
//	xsd.ENTITIES:     &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
//	xsd.ENTITY:       &ast.Ident{Name: "string"},
//	xsd.ID:           &ast.Ident{Name: "string"},
//	xsd.IDREF:        &ast.Ident{Name: "string"},
//	xsd.IDREFS:       &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
//	xsd.NCName:       &ast.Ident{Name: "string"},
//	xsd.NMTOKEN:      &ast.Ident{Name: "string"},
//	xsd.NMTOKENS:     &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
//	xsd.NOTATION:     &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
//	xsd.Name:         &ast.Ident{Name: "string"},
//	xsd.QName:        &ast.Ident{Name: "xml.Name"},
//	xsd.AnyURI:       &ast.Ident{Name: "string"},
//	xsd.Base64Binary: &ast.ArrayType{Elt: &ast.Ident{Name: "byte"}},
//	xsd.Boolean:      &ast.Ident{Name: "*bool"},
//	xsd.Byte:         &ast.Ident{Name: "byte"},
//	xsd.Date:         &ast.Ident{Name: "time.Time"},
//
//	//xsd.DateTime: &ast.Ident{Name: "time.Time"},
//	xsd.DateTime: &ast.Ident{Name: "time.Time"},
//
//	xsd.Decimal: &ast.Ident{Name: "*float64"},
//	xsd.Double:  &ast.Ident{Name: "*float64"},
//	// the "duration" built-in is especially broken, so we
//	// don't parse it at all.
//	xsd.Duration:           &ast.Ident{Name: "*string"},
//	xsd.Float:              &ast.Ident{Name: "*float32"},
//	xsd.GDay:               &ast.Ident{Name: "time.Time"},
//	xsd.GMonth:             &ast.Ident{Name: "time.Time"},
//	xsd.GMonthDay:          &ast.Ident{Name: "time.Time"},
//	xsd.GYear:              &ast.Ident{Name: "time.Time"},
//	xsd.GYearMonth:         &ast.Ident{Name: "time.Time"},
//	xsd.HexBinary:          &ast.ArrayType{Elt: &ast.Ident{Name: "byte"}},
//	xsd.Int:                &ast.Ident{Name: "*int"},
//	xsd.Integer:            &ast.Ident{Name: "*int"},
//	xsd.Language:           &ast.Ident{Name: "*string"},
//	xsd.Long:               &ast.Ident{Name: "*int64"},
//	xsd.NegativeInteger:    &ast.Ident{Name: "*int"},
//	xsd.NonNegativeInteger: &ast.Ident{Name: "*int"},
//	xsd.NormalizedString:   &ast.Ident{Name: "*string"},
//	xsd.NonPositiveInteger: &ast.Ident{Name: "*int"},
//	xsd.PositiveInteger:    &ast.Ident{Name: "*int"},
//	xsd.Short:              &ast.Ident{Name: "*int"},
//	xsd.String:             &ast.Ident{Name: "*string"},
//	xsd.Time:               &ast.Ident{Name: "time.Time"},
//	xsd.Token:              &ast.Ident{Name: "string"},
//	xsd.UnsignedByte:       &ast.Ident{Name: "*byte"},
//	xsd.UnsignedInt:        &ast.Ident{Name: "*uint"},
//	xsd.UnsignedLong:       &ast.Ident{Name: "*uint64"},
//	xsd.UnsignedShort:      &ast.Ident{Name: "*uint"},
//}
