package meta

import "os"

type Meta_Command_Result int8

const (
	META_COMMAND_SUCCESS      Meta_Command_Result = 0
	META_COMMAND_UNRECOGNIZED Meta_Command_Result = 1
)

func DoMetaCommamd(command string) Meta_Command_Result {
	if command == ".exit" {
		os.Exit(int(META_COMMAND_SUCCESS))
	}
	return META_COMMAND_UNRECOGNIZED

}
