package main

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

type ScopeItem struct {
	Enabled  bool   `json:"enabled"`
	File     string `json:"file"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}

type ScopeJsonFile struct {
	Target struct {
		Scope struct {
			AdvancedMode bool        `json:"advanced_mode"`
			Exclude      []ScopeItem `json:"exclude"`
			Include      []ScopeItem `json:"include"`
		} `json:"scope"`
	} `json:"target"`
}

func (f *ScopeJsonFile) IsWithinScopeProtocol(line string) bool {
	for _, v := range f.Target.Scope.Include {
		if !v.Enabled {
			continue
		}
		pattern := v.Protocol + "://" + v.Host[1:]
		if f.isInScope(line, pattern) {
			return true
		}
	}
	return false
}

func (f *ScopeJsonFile) isInScope(line, pattern string) bool {
	re, err := regexp.Compile(pattern)
	//TODO Handle this error better - should it halt? Or fail to a logfile?
	if err != nil {
		return false
	}
	if re.MatchString(line) {
		return true
	}
	return false
}

func (f *ScopeJsonFile) IsWithinScope(line string) bool {
	for _, v := range f.Target.Scope.Include {
		if !v.Enabled {
			continue
		}
		pattern := v.Host
		if f.isInScope(line, pattern) {
			return true
		}
	}
	return false
}

func (f *ScopeJsonFile) CheckScope(matchProtocol, invert bool, lines []string) (matched []string) {
	for _, line := range lines {
		inScope := false
		if matchProtocol {
			inScope = f.IsWithinScopeProtocol(string(line))
		} else {
			inScope = f.IsWithinScope(string(line))
		}

		if inScope || invert {
			matched = append(matched, line)
		}
	}
	return
}

func isValidFile(file string) (valid bool, err error) {
	_, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}
	valid = true
	return
}

func ParseFile(file string) (scope ScopeJsonFile, err error) {
	valid, err := isValidFile(file)
	if err != nil || !valid {
		return
	}
	data, err := ioutil.ReadFile(file)
	err = json.Unmarshal(data, &scope)
	if err != nil {
		return
	}
	return
}
