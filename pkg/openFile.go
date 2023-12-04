package pkg

import "os"

func OpenFile(path string) *os.File {
	dataInput, err := os.Open(os.Getenv("ABSOLUTE_PATH") + path)
	if err != nil {
		panic(err)
	}

	return dataInput
}