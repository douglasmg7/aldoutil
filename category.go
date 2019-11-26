package aldoutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

// Aldo product.
type Category struct {
	Name        string `db:"name"`
	Text        string `db:"text"`
	ProductsQty int    `db:"productsQty"`
	Selected    bool   `db:"selected"`
}

//  ReadCategoryList read list, lowercase, remove spaces and create a list of lines.
func ReadCategoryList(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// s := strings.Replace(string(b), " ", "", -1)
	// s = strings.ToLower(s)
	return strings.Split(string(b), "\n")
}

//  WriteCategoryList write a list to a file.
func WriteCategoryList(m *map[string]int, fileName string) {
	b := bytes.Buffer{}
	ss := []string{}
	// Sort.
	for k, v := range *m {
		ss = append(ss, fmt.Sprintf("%s (%d)\n", strings.ToLower(k), v))
	}
	sort.Strings(ss)
	// To buffer.
	for _, s := range ss {
		b.WriteString(s)
	}
	// Write to file.
	err := ioutil.WriteFile(fileName, bytes.TrimRight(b.Bytes(), "\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//  WriteCategoryList write a list from textcto a file.
func WriteCategoryListFromString(str, fileName string) error {
	b := bytes.Buffer{}
	ss := strings.Split(str, "\n")
	sort.Strings(ss)
	for _, s := range ss {
		reg := regexp.MustCompile(`\s+`)
		s = reg.ReplaceAllString(s, " ")
		b.WriteString(strings.TrimSpace(s) + "\n")
	}
	// Write to file.
	return ioutil.WriteFile(fileName, bytes.TrimRight(b.Bytes(), "\n"), 0644)
}
