package statement

import "fmt"

type PrepareResult int8

const (
	PREPARE_SUCCESS               PrepareResult = 0
	PREPARE_UNRECONIZED_STATEMENT PrepareResult = 1
)

type StatementType int8

const (
	STATEMENT_INSERT StatementType = 0
	STATEMENT_SELECT StatementType = 1
)

type Statement struct {
	st_type StatementType
}

func PrepareStatement(command string, statement *Statement) PrepareResult {
	if command[:5] == "insert" {
		statement.st_type = STATEMENT_INSERT
		return PREPARE_SUCCESS
	}

	if command[:6] == "select" {
		statement.st_type = STATEMENT_SELECT
		return PREPARE_SUCCESS
	}

	return PREPARE_UNRECONIZED_STATEMENT
}

func ExecuteStatement(statement *Statement) {
	switch statement.st_type {
	case STATEMENT_INSERT:
		fmt.Println("This is where we would do an insert.")
	case STATEMENT_SELECT:
		fmt.Println("This is where we would do an select")
	}
}
