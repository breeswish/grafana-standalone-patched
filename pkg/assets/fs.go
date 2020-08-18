package assets

import (
	"github.com/spf13/afero"
	"github.com/breeswish/go-bindata-afero"
)

var Fs afero.Fs = afero.NewMemMapFs()

func InitFs(basePath string) {
	Fs = afero.NewMemMapFs()
	err := bindataafero.WriteAssetsInDirectory(Fs, Asset, AssetInfo, AssetDir, basePath, "", "")
	if err != nil {
		panic(err)
	}
}
