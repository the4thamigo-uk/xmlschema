<p align="center"><img src="gophers.png" alt="gophers" style="width: 50%; height: 50%"></p>

# XML XSD schema validation library 

_Library providing XML schema validation using libxml2_

[![Build Status](https://secure.travis-ci.org/miracl/xmlschema.png?branch=master)](https://travis-ci.org/miracl/xmlschema?branch=master)
[![Coverage Status](https://coveralls.io/repos/miracl/xmlschema/badge.svg?branch=master&service=github)](https://coveralls.io/github/miracl/xmlschema?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/miracl/xmlschema)](https://goreportcard.com/report/github.com/miracl/xmlschema)

## Description

This library builds on the libxml2 wrapper library (github.com/lestrrat-go/libxml2) and provides a way to validate XML against a schema using the 
XML_CATALOG_FILES mechanism described in http://xmlsoft.org/catalog.html. This mechanism of libxml2, is intended to cater for validation against 
complex schemas that are defined in terms of multiple sub-schemas. Normally this approach requires the schema files to be present on disk at time of validation
(e.g. run time), therefore requiring the files to be deployed to the target machine, as separate files, along with your application (which you may not want to do).
This library works-around this limitation by storing the schemas in your application code and when executed deploys them in a transitory fashion to a 
temporary folder, purely for the purposes of loading them using libxml2.
