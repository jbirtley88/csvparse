package csvparse_test

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"

	csvreader "github.com/Dan6erbond/csvparse"
	"github.com/stretchr/testify/assert"
)

func TestDictReader_Headers(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	dr := csvreader.NewDictReader(r)

	headers, err := dr.Headers()

	assert.NoError(t, err)

	assert.Equal(t, "first_name", headers[0])
	assert.Equal(t, "last_name", headers[1])
	assert.Equal(t, "username", headers[2])
}

func ExampleDictReader_Headers() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	dr := csvreader.NewDictReader(r)

	headers, _ := dr.Headers()

	fmt.Println(headers)

	// Output: [first_name last_name username]
}

func TestDictReader_Read(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	dr := csvreader.NewDictReader(r)

	row, err := dr.Read()

	assert.NoError(t, err)

	assert.Equal(t, "Rob", row["first_name"])
}

func ExampleDictReader_Read() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	dr := csvreader.NewDictReader(r)

	row, _ := dr.Read()

	fmt.Println(row["first_name"])

	// Output: Rob
}

func TestDictReader_ReadAll(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	dr := csvreader.NewDictReader(r)

	records, err := dr.ReadAll()

	assert.NoError(t, err)

	assert.Equal(t, "Robert", records[2]["first_name"])
}

func ExampleDictReader_ReadAll() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	dr := csvreader.NewDictReader(r)

	records, _ := dr.ReadAll()

	fmt.Println(records[2]["first_name"])

	// Output: Robert
}
