package e2eclient

import (
	"fmt"
	"os"
)

func FailE2eTest(msg string)  {
	fmt.Println("× " + msg)
	os.Exit(1)
}

func SuccessE2eTest(msg string)  {
	fmt.Println("✔ " +msg)
}
