package sdk

import "regexp"

type parserVariantSpecification struct {
	mapping []*mappingRule
	regexp  *regexp.Regexp
}

//func (vs *parserVariantSpecification) compile(settings ParserVariantSettings) (*parserVariantSpecification, error) {
//	vs.regexp = regexp.MustCompile(settings.Regexp)
//
//	for _, expr := range settings.Mapping {
//		mr, err := new(mappingRule).compile(expr, true)
//
//		if err != nil {
//			return nil, err
//		}
//
//		vs.mapping = append(vs.mapping, mr)
//	}
//
//	return vs, nil
//}
