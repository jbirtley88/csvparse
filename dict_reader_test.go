package csvparse_test

import (
	"encoding/csv"
	"strings"
	"testing"

	csvreader "github.com/Dan6erbond/csvparse"
	"github.com/stretchr/testify/assert"
)

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
