package packages

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ChangeValue(v interface{}) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && val.Elem().CanSet() {
		switch val.Elem().Kind() {
		case reflect.Int:
			val.Elem().SetInt(100)
		case reflect.String:
			val.Elem().SetString("Updated via reflect")
		}
	}
}

type MyStruct struct{}

func (m MyStruct) Hello(name string) string {
	return "Hello, " + name
}

func InvokeMethod(obj interface{}, methodName string, args ...interface{}) []reflect.Value {
	val := reflect.ValueOf(obj)
	method := val.MethodByName(methodName)
	
	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}
	
	return method.Call(reflectArgs)
}

type Person struct {
	Name string
	Age  int
}

func GetAndSetField(obj interface{}, fieldName string, newValue interface{}) {
	val := reflect.ValueOf(obj).Elem()
	field := val.FieldByName(fieldName)
	
	if field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(newValue))
	}
}

func CreateAndSetFields(t reflect.Type, fieldValues map[string]interface{}) interface{} {
	newStruct := reflect.New(t).Elem()
	for fieldName, fieldValue := range fieldValues {
		field := newStruct.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(fieldValue))
		}
	}
	return newStruct.Interface()
}

func CompareStructs() {
	type Config struct {
		Host string
		Port int
	}
	
	cfg1 := Config{Host: "localhost", Port: 8080}
	cfg2 := Config{Host: "localhost", Port: 8080}
	cfg3 := Config{Host: "127.0.0.1", Port: 9090}
	
	fmt.Println(reflect.DeepEqual(cfg1, cfg2))
	fmt.Println(reflect.DeepEqual(cfg1, cfg3))
}

func InspectInterface(i interface{}) {
	val := reflect.ValueOf(i)
	typ := reflect.TypeOf(i)
	
	fmt.Printf("Тип интерфейса: %v\n", typ)
	fmt.Printf("Значение интерфейса: %v\n", val)
	
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		fmt.Printf("Значение по указателю: %v\n", val)
	}
}

func StartExample() {
	x := 10
	str := "Original"
	ChangeValue(&x)
	ChangeValue(&str)
	fmt.Println("Измененный int:", x)
	fmt.Println("Измененная строка:", str)
	
	m := MyStruct{}
	result := InvokeMethod(m, "Hello", "World")
	fmt.Println(result[0])
	
	p := &Person{Name: "Alice", Age: 25}
	GetAndSetField(p, "Name", "Bob")
	GetAndSetField(p, "Age", 30)
	fmt.Printf("Updated person: %+v\n", p)
	
	personType := reflect.TypeOf(Person{})
	newPerson := CreateAndSetFields(personType, map[string]interface{}{
		"Name": "Charlie",
		"Age":  22,
	}).(Person)
	fmt.Printf("Созданный человек: %+v\n", newPerson)
	
	CompareStructs()
	
	var i interface{} = "Some string"
	InspectInterface(i)
}

/*
 * Простенький пример функции которая читает JSON (маппер для кастомных тегов)
 */

type A struct {
	Name  string `custom:"name"`
	Age   int    `custom:"age"`
	Email string `custom:"email_address"`
}

func CustomMapper(data []byte, out interface{}) error {
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		return err
	}
	
	v := reflect.ValueOf(out).Elem()
	t := v.Type()
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		
		tag := field.Tag.Get("custom")
		if tag == "" {
			continue
		}
		
		if rawValue, ok := rawData[tag]; ok {
			fieldValue.Set(reflect.ValueOf(rawValue).Convert(field.Type))
		}
	}
	
	return nil
}
