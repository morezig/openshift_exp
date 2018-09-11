package common

import (
	"net/http"
	"strings"
)

type BFS struct {
	FileSystem http.FileSystem
}

func (bfs *BFS) Open(name string) (http.File, error) {
	return bfs.FileSystem.Open(name)
}

func (bfs *BFS) Exists(prefix string, filepath string) bool {
	var err error
	var url string
	url = strings.TrimPrefix(filepath, prefix)
	if len(url) < len(filepath) {
		_, err = bfs.FileSystem.Open(url)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

// func GetBFS(root string) *BFS {
// 	var afs *assetfs.AssetFS
// 	afs = &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}

// 	return &BFS{afs}
// }
