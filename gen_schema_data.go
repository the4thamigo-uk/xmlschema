// +build ignore

package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func failIfError(err error) {
	if err != nil {
		fmt.Println("FAILED")
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	// get the directory containing this source file
	_, callerFile, _, _ := runtime.Caller(0)
	callerDir := path.Dir(callerFile)
	xsdDir := path.Join(callerDir, "xsd")

	fmt.Printf("Importing XSD schema from %v...", xsdDir)

	data, err := readFiles("./xsd")
	if err != nil {
		failIfError(err)
	}

	gocode := fmt.Sprintf(`package xmlschema

const schemaData = "%v"
`, data)

	fi, err := os.Stat(path.Join(callerDir, "./gen_schema_data.go"))
	failIfError(err)

	err = ioutil.WriteFile(path.Join(callerDir, "./schema_data.go"), []byte(gocode), fi.Mode())
	failIfError(err)

	fmt.Println("OK")
}

func readFiles(dir string) (string, error) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	fds := map[string][]byte{}
	for _, fi := range fis {
		fd, err := ioutil.ReadFile(path.Join(dir, fi.Name()))
		if err != nil {
			return "", err
		}
		fds[fi.Name()] = fd
	}
	b, err := json.Marshal(fds)
	if err != nil {
		return "", err
	}
	b, err = deflate(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// Deflate is a helper function for compress/flate
func deflate(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	fw, err := flate.NewWriter(b, flate.BestCompression)
	if err != nil {
		return nil, err
	}
	_, err = fw.Write(data)
	if err != nil {
		return nil, err
	}
	err = fw.Close()
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
