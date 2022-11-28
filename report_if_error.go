package schemacafe

import "fmt"

func ReportIfError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
