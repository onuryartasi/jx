// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	gits "github.com/jenkins-x/jx/pkg/gits"
)

func AnySliceOfGitsGitLabel() []gits.GitLabel {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]gits.GitLabel))(nil)).Elem()))
	var nullValue []gits.GitLabel
	return nullValue
}

func EqSliceOfGitsGitLabel(value []gits.GitLabel) []gits.GitLabel {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []gits.GitLabel
	return nullValue
}
