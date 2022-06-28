package utils

import "log"

type CheckErrArgs struct {
	Message string
}

func CheckErr(err error, args *CheckErrArgs) {
	if err != nil {
		if args != nil {
			log.Fatal(args.Message)
		}
		args = &CheckErrArgs{Message: "Error occurred: %v"}
		log.Fatalf(args.Message, err)
	}
}
