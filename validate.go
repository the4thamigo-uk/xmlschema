package xmlschema

//go:generate go run gen_schema_data.go

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/parser"
	"github.com/lestrrat-go/libxml2/xsd"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
)

var (
	cacheSchemas map[Schema]*xsd.Schema
)

// Schema represents the name of a schema used to validate XML
type Schema string

const (
	catalogFile = "catalog"
	libxml2Env  = "XML_CATALOG_FILES"

	// Protocol is the name of the SAML protocol schema file
	Protocol Schema = "saml-schema-protocol-2.0.xsd"

	// Metadata is the name of the SAML metadata schema file
	Metadata Schema = "saml-schema-metadata-2.0.xsd"

	// XHTML is the name of the XHTML schema file
	XHTML Schema = "xhtml1-strict.xsd"
)

func inflate(data []byte) ([]byte, error) {
	return ioutil.ReadAll(flate.NewReader(bytes.NewReader(data)))
}

func unzip(dir string) error {
	b, err := base64.StdEncoding.DecodeString(schemaData)
	if err != nil {
		return err
	}
	b, err = inflate(b)
	if err != nil {
		return err
	}
	var fds map[string][]byte
	err = json.Unmarshal(b, &fds)
	if err != nil {
		return err
	}
	for fn, fd := range fds {
		err = ioutil.WriteFile(path.Join(dir, fn), fd, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func unzipTemp() (string, error) {
	tmp, err := ioutil.TempDir("", "samlxml")
	if err != nil {
		return "", err
	}
	err = unzip(tmp)
	if err != nil {
		return "", err
	}
	return tmp, nil
}

func parseSchemas(dir string) (map[Schema]*xsd.Schema, error) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	m := map[Schema]*xsd.Schema{}
	for _, fi := range fis {
		if path.Ext(fi.Name()) == ".xsd" {
			s, err := xsd.ParseFromFile(path.Join(dir, fi.Name()))
			if err != nil {
				return nil, err
			}
			m[Schema(fi.Name())] = s
		}
	}
	return m, nil
}

// loadSchemas extracts the schema data from schemaData, stores it in a temporary folder.
// The XML_CATALOG_FILES environment variable is set to point to this folder which enables
// the schemas to be loaded into memory. The folder is removed once the schemas are loaded.
func loadSchemas() (map[Schema]*xsd.Schema, error) {
	if cacheSchemas == nil {
		schemaDir, err := unzipTemp()
		if err != nil {
			return nil, err
		}
		defer func() { _ = os.RemoveAll(schemaDir) }() //#nosec

		oldEnv, oldEnvOk := os.LookupEnv(libxml2Env)
		err = os.Setenv(libxml2Env, path.Join(schemaDir, catalogFile))
		if err != nil {
			return nil, err
		}
		defer func() {
			if oldEnvOk {
				err = os.Setenv(libxml2Env, oldEnv)
			} else {
				_ = os.Unsetenv(libxml2Env) //#nosec
			}
		}()

		schemas, err := parseSchemas(schemaDir)
		if err != nil {
			return nil, err
		}
		cacheSchemas = schemas
	}
	return cacheSchemas, nil

}

// Validate performs validation of the given xml data against the named schema
func Validate(xml []byte, schemaName Schema) error {
	schemas, err := loadSchemas()
	if err != nil {
		return err
	}
	schema := schemas[schemaName]
	doc, err := libxml2.Parse(xml, parser.XMLParseNoError|parser.XMLParseNoWarning)
	if err != nil {
		return err
	}
	err = schema.Validate(doc)
	if err != nil {
		allErr := errors.New("Failed to validate XML")
		for _, e := range err.(xsd.SchemaValidationError).Errors() {
			allErr = errors.Wrap(e, err.Error())
		}
		return allErr
	}
	return nil
}
