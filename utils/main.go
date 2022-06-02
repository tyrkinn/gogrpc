package utils

import "log"

type CheckErrArgs struct {
	Message string
}

func CheckErr(err error, args *CheckErrArgs) {
	if args != nil {
		log.Fatal(args.Message)
	}
	args = &CheckErrArgs{Message: "Error occurred: %v"}
	if err != nil {
		log.Fatalf(args.Message, err)
	}
}
