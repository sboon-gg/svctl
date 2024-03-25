package prbf2update

import (
	"archive/zip"
	"encoding/json"
	"io"
	"path/filepath"
	"strings"
)

type patchDataInfo struct {
	Version         string `json:"Version"`
	RequiresVersion string `json:"RequiresVersion"`
	PatchData       []struct {
		Method      string `json:"Method"`
		Source      string `json:"Source,omitempty"`
		Destination string `json:"Destination"`
		Entry       string `json:"Entry,omitempty"`
		BeforeHash  string `json:"BeforeHash,omitempty"`
		AfterHash   string `json:"AfterHash,omitempty"`
	} `json:"PatchData"`
}

func patchChangedFiles(archive string) ([]string, error) {
	info, err := readPatchDataInfo(archive)
	if err != nil {
		return nil, err
	}

	var changedFiles []string

	for _, patch := range info.PatchData {
		if patch.Method == "FileAdd" {
			parts := strings.Split(patch.Destination, "\\")
			changedFiles = append(changedFiles, filepath.Join(parts[1:]...))
		}
	}

	return changedFiles, nil
}

func readPatchDataInfo(archive string) (*patchDataInfo, error) {
	data, err := readFileFromZIP(archive, "info.json")
	if err != nil {
		return nil, err
	}

	var info patchDataInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func readFileFromZIP(archive, path string) ([]byte, error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.Name == path {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			return io.ReadAll(rc)
		}
	}

	return nil, io.EOF
}
