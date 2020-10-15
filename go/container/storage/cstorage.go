package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/containers/storage"
	"github.com/spf13/viper"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
)

// Layer is an enum representing the fileystem layer on which a file
// exists. Can be a mount, container layer, image, or unknown if not found.
type Layer uint8

const (
	LMOUNT     Layer = iota
	LCONT      Layer = iota
	LIMAGE     Layer = iota
	LHOST      Layer = iota
	LUNKNOWN   Layer = iota
	MOUNT_ROOT       = "mountRoot"
	ROOT_DIR         = "rootDir"
)

// CStorage is an object representing the layered filesystem of OCI containers.
type CStorage struct {
	store storage.Store
	//contCache map[string]*ContFileSystem
	config    *viper.Viper
	mountRoot bool
	rootDir   string
}

// ContFileSystem describes a container's filesystem based on the layers and mounts it uses.
type ContFileSystem struct {
	UserDir             string
	ContTopLayer        string
	ContTopLayerMerged  string
	ImageTopLayer       string
	ImageTopLayerMerged string
	ContainerConfig     map[string]interface{}
	Mounts              []*Mount
}

// Mount represents a filesystem mount.
type Mount struct {
	Type string
	Src  string
	Dst  string
}

// NewContainerStore creates a CStorage object given a configuration.
func NewContainerStore(conf *viper.Viper) (*CStorage, error) {
	options, err := storage.DefaultStoreOptions(false, 0)
	if err != nil {
		return nil, err
	}
	store, err := storage.GetStore(options)
	if err != nil {
		return nil, err
	}
	var mr bool
	var rd string
	if conf.IsSet(MOUNT_ROOT) && conf.IsSet(ROOT_DIR) {
		mr = conf.GetBool(MOUNT_ROOT)
		rd = conf.GetString(ROOT_DIR)
	} else {
		return nil, errors.New("mountRoot and rootDir variables must be set in the config")
	}
	//return &CStorage{store, make(map[string]*ContFileSystem), conf, mr, rd}, nil
	return &CStorage{store, conf, mr, rd}, nil
}

// GetStatus gives the current status of the filesystem.
func (c *CStorage) GetStatus() ([][2]string, error) {
	return c.store.Status()
}

// GetContainer returns the storage object representing a container based on its container ID.
func (c *CStorage) GetContainer(id string) (*storage.Container, error) {
	cont, err := c.store.Container(id)
	if err != nil {
		return nil, err
	}
	return cont, nil
}

// GetContTopLayer returns the path to the writable layer of the container given a container ID.
func (c *CStorage) GetContTopLayer(id string) (string, error) {
	cont, err := c.GetContainer(id)
	if err != nil {
		return "", err
	}

	path := c.store.GraphRoot() + "/overlay/" + cont.LayerID
	return path, nil
}

// GetContUserDir returns the path to the directory that holsd container metadata.
func (c *CStorage) GetContUserDir(id string) (string, error) {
	path, err := c.store.ContainerDirectory(id)
	if err != nil {
		return "", err
	}
	return path, nil
}

// GetImage returns the storage object representing a container image  based on its container ID.
func (c *CStorage) GetImage(ContainerID string) (*storage.Image, error) {
	cont, err := c.GetContainer(ContainerID)
	if err != nil {
		return nil, err
	}
	image, err := c.store.Image(cont.ImageID)
	if err != nil {
		return nil, err
	}
	return image, nil
}

// GetImageTopLayer returns the path to the top layer of the container image  given a container ID.
func (c *CStorage) GetImageTopLayer(ContainerID string) (string, error) {
	image, err := c.GetImage(ContainerID)
	if err != nil {
		return "", err
	}
	path := c.store.GraphRoot() + "/overlay/" + image.TopLayer
	return path, nil

}

// GetContainerConfig returns the configuration information about a container given its container Id.
func (c *CStorage) GetContainerConfig(ContainerID string) (map[string]interface{}, error) {
	path, err := c.GetContUserDir(ContainerID)
	f, err := os.Open(path + "/config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)
	var v map[string]interface{}
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// getValue returns a value casted as a string given a string key and a map.
func getValue(m map[string]interface{}, k string) string {
	if v, ok := m[k]; ok {
		return v.(string)
	}
	return ""
}

// createMount creates a mount object given a map of attributes.
func createMount(m map[string]interface{}) *Mount {
	return &Mount{Type: getValue(m, "type"), Src: getValue(m, "source"), Dst: getValue(m, "destination")}
}

// GetMounts returns the set of mounts for a given contaienr given a container ID.
func (c *CStorage) GetMounts(ContainerID string) ([]*Mount, error) {
	var mountMap []*Mount
	config, err := c.GetContainerConfig(ContainerID)
	if err != nil {
		return nil, err
	}
	if mts, ok := config["mounts"]; ok {
		mounts := mts.([]interface{})
		for _, k := range mounts {
			mount := k.(map[string]interface{})
			m := createMount(mount)
			mountMap = append(mountMap, m)
		}

	}
	return mountMap, nil

}

// GetCFileSystem returns filesystem information for a container given its ID.
func (c *CStorage) GetCFileSystem(ContainerID string) (*ContFileSystem, error) {
	mounts, err := c.GetMounts(ContainerID)
	if err != nil {
		return nil, err
	}
	userDir, err := c.GetContUserDir(ContainerID)
	if err != nil {
		return nil, err
	}
	config, err := c.GetContainerConfig(ContainerID)
	if err != nil {
		return nil, err
	}
	topLayer, err := c.GetContTopLayer(ContainerID)
	if err != nil {
		return nil, err
	}
	imageTopLayer, err := c.GetImageTopLayer(ContainerID)
	if err != nil {
		return nil, err
	}

	return &ContFileSystem{UserDir: userDir, ContTopLayer: topLayer + "/diff/",
		ContTopLayerMerged: topLayer + "/merged/", ImageTopLayer: imageTopLayer + "/diff/",
		ImageTopLayerMerged: topLayer + "/merged/", ContainerConfig: config, Mounts: mounts}, nil

}

/*
// GetCFileSystem caches and returns container filesystem information given a container ID
func (c *CStorage) GetCFileSystem(ContainerID string) (*ContFileSystem, error) {
	if cfs, ok := c.contCache[ContainerID]; ok {
		return cfs, nil
	}
	cfs, err := c.getCFileSystem(ContainerID)
	if err != nil {
		return nil, err
	}
	c.contCache[ContainerID] = cfs
	return cfs, nil
}*/

/*
// RemoveCFileSystem removes container filesystem information from cache given a container ID
func (c *CStorage) RemoveCFileSystem(ContainerID string) {
	delete(c.contCache, ContainerID)
}*/

// GetContainerFilePath finds the true host filesystem path of a file given its container path and container ID.
// Function also returns the filesystem layer on which the path was found  (mount, container layer, or image layer).
func (c *CStorage) GetContainerFilePath(ContainerID string, path string) (string, Layer, error) {
	cont, err := c.GetCFileSystem(ContainerID)
	if err != nil {
		return "", LUNKNOWN, err
	}
	for _, mnt := range cont.Mounts {
		if mnt.Type == "bind" && strings.HasPrefix(path, mnt.Dst) {
			suffix := strings.TrimPrefix(path, mnt.Dst)
			trueLoc := mnt.Src + suffix
			if c.mountRoot {
				trueLoc = c.rootDir + "/" + trueLoc
			}
			if exists, _ := ioutils.FileExists(trueLoc); exists {
				return trueLoc, LMOUNT, nil
			}
		}

	}

	trueLoc := cont.ContTopLayer + path
	if exists, _ := ioutils.FileExists(trueLoc); exists {
		return trueLoc, LCONT, nil
	}
	trueLoc = cont.ImageTopLayerMerged + path
	if exists, _ := ioutils.FileExists(trueLoc); exists {
		return trueLoc, LIMAGE, nil
	}
	return "", LUNKNOWN, nil

}

// GetHostFilePath finds the true host filesystem path of a file if there is a root mount.
// Function also returns the filesystem layer on which the path was found which will always be HOST  (mount, container layer, or image layer).
func (c *CStorage) GetHostFilePath(path string) (string, Layer, error) {
	trueLoc := path
	if c.mountRoot {
		trueLoc = c.rootDir + "/" + trueLoc
	}
	if exists, _ := ioutils.FileExists(trueLoc); exists {
		return trueLoc, LHOST, nil
	}
	return "", LUNKNOWN, nil
}

// GetFilePath finds the true container or  filesystem path of a file. If the container id is "", the api assumes the file is host
// Function also returns the filesystem layer on which the path was found which will always be HOST for a host file, but the following for a container:  (mount, container layer, or image layer).
func (c *CStorage) GetFilePath(contId string, path string) (string, Layer, error) {
	if contId == "" {
		return c.GetHostFilePath(path)
	}
	return c.GetContainerFilePath(contId, path)
}
