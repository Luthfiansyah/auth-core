//
//  document.go
//  Auth Core
//
//  Copyright © 2018 Auth Core. All rights reserved.
//
package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/auth-core/logging"
)

var log = logging.MustGetLogger("auth-core")

type Document struct {
	Name         string
	Schema       string
	Hash         string
	RefreshAfter int
	Data         *[]byte
}

// used to represent the document section of the document
type DocumentHeader struct {
	HeaderData struct {
		Schema       string `json:"schema"`
		Name         string `json:"name"`
		Revision     int    `json:"revision"`
		CreationDate string `json:"creation_date"`
		RefreshAfter int    `json:"refresh_after"`
	} `json:"document"`
}

type DocumentService struct {
	Directory string
	documents map[string](*Document)
}

func NewDocumentService(directory string) *DocumentService {
	return &DocumentService{
		Directory: directory,
		documents: make(map[string](*Document)),
	}
}

func (ds *DocumentService) ReadAll() error {
	files, err := ioutil.ReadDir(ds.Directory)
	if err != nil {
		return err
	}
	for _, file := range files {
		// does not follow directories recusively
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".json") {
			log.Infof(logging.INTERNAL, "skipping file %s (no .json suffix)", file.Name())
			continue
		}
		log.Debugf(logging.INTERNAL, "found file %s", filepath.Join(ds.Directory, file.Name()))
		doc, err := readFile(filepath.Join(ds.Directory, file.Name()))
		if err != nil {
			return err
		}
		key := (*doc).Name + "|" + (*doc).Schema
		log.Infof(logging.INTERNAL, "adding document %+v with key %s", *doc, key)
		ds.documents[key] = doc
	}
	return nil
}

func (ds *DocumentService) GetDocument(name string, schema string) (*Document, bool) {
	key := name + "|" + schema
	if doc, ok := (*ds).documents[key]; ok {
		return doc, true
	}
	return nil, false
}

func readFile(fileName string) (*Document, error) {
	// read file
	log.Debugf(logging.INTERNAL, "reading file %s", fileName)
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Debugf(logging.INTERNAL, "error reading file %s", err.Error())
		return nil, err
	}
	// parse header (document tag)
	header, err := parseHeader(bytes)
	if err != nil {
		log.Debugf(logging.INTERNAL, "error parsing header %s", err.Error())
		return nil, err
	}
	log.Debugf(logging.INTERNAL, "parsed header %+v", *header)
	// compute hash
	hashBytes := sha256.Sum256(bytes)
	hexString := hex.EncodeToString(hashBytes[:])
	log.Debugf(logging.INTERNAL, "hash: %s", hexString)
	d := Document{
		Name:         (*header).HeaderData.Name,
		Schema:       (*header).HeaderData.Schema,
		Hash:         hexString,
		RefreshAfter: (*header).HeaderData.RefreshAfter,
		Data:         &bytes,
	}
	return &d, nil
}

func parseHeader(bytes []byte) (*DocumentHeader, error) {
	header := DocumentHeader{}
	if err := json.Unmarshal(bytes, &header); err != nil {
		return nil, err
	}
	return &header, nil
}
