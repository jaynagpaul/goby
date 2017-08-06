package vm

import "testing"

func TestPGConnectionPing(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`
			require "db"

			db = DB.open("postgres", "user=postgres dbname=postgres sslmode=disable")
			db.driver.ping
			`,
			true},
		{`
			require "db"

			db = DB.open("postgres", "user=test dbname=postgres sslmode=disable")
			db.driver.ping
			`,
			false},
	}

	for i, tt := range tests {
		v := initTestVM()
		evaluated := v.testEval(t, tt.input)
		checkExpected(t, i, evaluated, tt.expected)
		v.checkCFP(t, i, 0)
		v.checkSP(t, i, 1)
	}
}