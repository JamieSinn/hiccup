package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	file, err := ParseFile("testfiles/valid.json")
	if err != nil {
		t.Fail()
	}
	fmt.Println(file)
}

func TestScopeJsonFile_CheckScope(t *testing.T) {
	inScope := []string {"github.com", "test.githubapp.com", "test.github.net"}
	outOfScope := []string {"shop.github.com", "agithubapp.com", "enterprise.github.com"}
	file, err := ParseFile("testfiles/valid.json")
	if err != nil {
		t.Fail()
		return
	}
	inScopeCheck := file.CheckScope(false, false, inScope)
	outOfScopeCheck := file.CheckScope(false, false, outOfScope)

	if !reflect.DeepEqual(inScope, inScopeCheck) {
		fmt.Println("Failed check for in-scope.")
		fmt.Println("Expected")
		fmt.Println(inScope)
		fmt.Println("Real")
		fmt.Println(inScopeCheck)
		t.Fail()
	}
	if len(outOfScopeCheck) != 0 {
		fmt.Println("Failed check for out-of-scope.")
		fmt.Println("Expected")
		fmt.Println([]string{})
		fmt.Println("Real")
		fmt.Println(outOfScopeCheck)
		t.Fail()
	}
}

func TestScopeJsonFile_IsWithinScope(t *testing.T) {
	file, err := ParseFile("testfiles/valid.json")
	if err != nil {
		t.Fail()
		return
	}

	if !file.IsWithinScope("github.com") {
		t.Fail()
	}
}

func TestScopeJsonFile_IsWithinScopeProtocol(t *testing.T) {
	file, err := ParseFile("testfiles/valid.json")
	if err != nil {
		t.Fail()
		return
	}

	if file.IsWithinScope("shop.github.com") {
		t.Fail()
	}
}