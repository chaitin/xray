package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/cel-go/common"
	"github.com/google/cel-go/parser"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

type Rule struct {
	Expression string `yaml:"expression"`
}

type YamlPOC struct {
	Set   map[string]string `yaml:"set"`
	Rules []*Rule           `yaml:"rules"`
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		os.Exit(1)
	}
}

func fmtExpr(expStr string) (string, error) {
	expr, errs := parser.Parse(common.NewTextSource(expStr))
	if len(errs.GetErrors()) > 0 {
		return "", fmt.Errorf("parse expr %q failed, error: %v", expStr, errs.GetErrors())
	}
	newExpStr, err := parser.Unparse(expr.GetExpr(), nil)
	if err != nil {
		return "", fmt.Errorf("unparse expr %q failed, error: %v", expStr, err)
	}
	return newExpStr, nil
}

type CheckDetail struct {
	Origin    string
	Formatted string
}
type CheckResult struct {
	File    string
	Details []CheckDetail
}

func main() {
	args := os.Args
	if len(args) < 2 {
		handleErr(errors.New("usage: ./cellint FILE"))
	}
	results := make([]CheckResult, 0, len(args)-1)
	editFile := os.Getenv("cellint_edit_file") == "true"
	for _, file := range args[1:] {
		content, err := ioutil.ReadFile(file)
		handleErr(err)
		poc := YamlPOC{}
		err = yaml.Unmarshal(content, &poc)
		handleErr(err)
		result := CheckResult{File: file, Details: make([]CheckDetail, 0, 2)}
		for _, item := range poc.Set {
			expr := strings.TrimSpace(item)
			newExpr, err := fmtExpr(expr)
			handleErr(err)
			if expr != newExpr {
				result.Details = append(result.Details, CheckDetail{Origin: expr, Formatted: newExpr})
			}
		}
		for _, item := range poc.Rules {
			expr := strings.TrimSpace(item.Expression)
			newExpr, err := fmtExpr(expr)
			handleErr(err)
			if expr != newExpr {
				result.Details = append(result.Details, CheckDetail{Origin: expr, Formatted: newExpr})
			}
		}
		results = append(results, result)
		if editFile {
			for _, item := range result.Details {
				content = bytes.Replace(content, []byte(item.Origin), []byte(item.Formatted), 1)
			}
			fileInfo, err := os.Stat(file)
			handleErr(err)
			err = ioutil.WriteFile(file, content, fileInfo.Mode())
			handleErr(err)
		}
	}
	if os.Getenv("cellint_json_output") == "true" {
		jsonResult, err := json.MarshalIndent(results, "", "    ")
		handleErr(err)
		fmt.Println(string(jsonResult))
	} else {
		hasError := false
		for _, result := range results {
			for idx, detail := range result.Details {
				if idx == 0 {
					fmt.Println("File: ", result.File)
				}
				fmt.Println("Origin: ", detail.Origin, "\nNew   : ", detail.Formatted)
				hasError = true
				if idx == len(result.Details)-1 {
					fmt.Printf("\n\n")
				}
			}
		}
		if hasError {
			os.Exit(2)
		}
	}
}