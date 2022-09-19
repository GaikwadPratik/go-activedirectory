package goactivedirectory

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-ldap/ldap/v3"
)

type propMapStruct struct {
	TagValue  string
	FieldName string
	FieldType string
}

var propsByTag = make(map[string][]*propMapStruct)

// buildPropsByTagMap builds reflection map for easy lookup
// https://stackoverflow.com/a/55922473/2398353
// TODO:call using sync.once, use reflection
func buildPropsByTagMap(s interface{}) {

	if reflect.ValueOf(s).Kind() != reflect.Pointer {
		panic("Input is not pointer")
	}

	rt := reflect.ValueOf(s).Elem().Type()

	if rt.Kind() != reflect.Struct {
		panic("Input is not struct")
	}

	refName := rt.Name()

	if propsByTag[refName] == nil || len(propsByTag[refName]) == 0 {
		propsByTag[refName] = make([]*propMapStruct, 0)
	}

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get("activedirectory"), ",")[0] // use split to ignore tag "options"
		if v == "" || v == "-" {
			continue
		}
		propsByTag[refName] = append(propsByTag[refName], &propMapStruct{TagValue: v, FieldName: f.Name, FieldType: f.Type.String()})
	}
}

var propByTagsOnce sync.Once

func fillObjFromAttrMap(item interface{}, attributes []*ldap.EntryAttribute) error {
	if len(propsByTag) == 0 {
		propByTagsOnce.Do(func() {
			adu := &ActiveDirectoryUser{}
			buildPropsByTagMap(adu)
			adg := &ActiveDirectoryGroup{}
			buildPropsByTagMap(adg)
		})
	}

	v := reflect.ValueOf(item).Elem()

	if !v.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	fields, ok := propsByTag[v.Type().Name()]

	if !ok || len(fields) == 0 {
		return fmt.Errorf("prop map is not populated for: %s", v.Type().Name())
	}

	for _, attr := range attributes {
		var propInfo *propMapStruct
		for _, prop := range fields {
			if prop.TagValue == attr.Name {
				propInfo = prop
				break
			}
		}

		if propInfo == nil {
			return fmt.Errorf("prop not found by tag: %s", attr.Name)
		}

		prop := v.FieldByName(propInfo.FieldName)

		switch propInfo.FieldType {
		case "int":
			intVar, err := strconv.Atoi(attr.Values[0])
			if err != nil {
				return err
			}
			prop.SetInt(int64(intVar))
		case "string":
			prop.SetString(attr.Values[0])
		case "*time.Time":
			layout := time.RFC3339[:len(attr.Values[0])]
			dt, err := time.Parse(layout, attr.Values[0])
			if err != nil {
				fmt.Println(err)
				return err
			}
			prop.Set(reflect.ValueOf(&dt))
		case "[]string":
			prop.Set(reflect.ValueOf(attr.Values))
		default:
			return fmt.Errorf("wrong prop type found: %s", propInfo.FieldType)
		}
	}
	return nil
}
