package searchutil

import (
	"fmt"
	"testing"

	"github.com/shortedapp/shortedfunctions/pkg/awsutil"
	"github.com/stretchr/testify/assert"
)

type Searchutilclient struct {
	awsutil.AwsUtiler
}

func (s Searchutilclient) FetchDynamoDBLastModified(tableName string, keyName string) (string, error) {
	if tableName == "test" {
		return "", fmt.Errorf("error")
	}
	return "2006-02-09T15:04:05Z", nil
}

func TestGetSearchWindow(t *testing.T) {

	client := Searchutilclient{}
	testCases := []struct {
		table  string
		period SearchPeriod
		result int64
	}{
		{"test", Year, 1},
		{"test", Month, 1},
		{"test", Week, 7},
		{"test", Day, 1},
		{"test", Latest, 0},
		{"test2", Year, 1},
		{"test2", Month, 1},
		{"test2", Week, 7},
		{"test2", Day, 1},
		{"test2", Latest, 11},
	}

	for _, test := range testCases {
		res, res2 := GetSearchWindow(client, test.table, "", test.period)
		if test.period == Latest && test.table != "test" {
			assert.True(t, test.result <= (res2/10000-res/10000))
		} else if test.period == Latest {
			assert.True(t, test.result <= res2-res)
		} else if test.period == Year {
			assert.True(t, test.result == (res2/10000-res/10000))
		} else if test.period == Month {
			diff := (res2/100)%100 - (res/100)%100
			assert.True(t, (test.result == diff || diff == 11))
		} else if test.period == Week {
			diff := res2 - res
			assert.True(t, (test.result == diff || diff > 21))
		} else if test.period == Day {
			diff := res2%100 - res%100
			assert.True(t, (test.result == diff || diff > 27))
		}
	}
}

func TestStringToSearchPeriod(t *testing.T) {

	testCases := []struct {
		period string
		result SearchPeriod
	}{
		{"day", Day},
		{"week", Week},
		{"month", Month},
		{"year", Year},
		{"asd", Week},
	}

	for _, test := range testCases {
		assert.Equal(t, test.result, StringToSearchPeriod(test.period))
	}
}
