package docker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wemanity-belgium/hyperclair/docker/httpclient"
	"github.com/wemanity-belgium/hyperclair/xerrors"
)

//Pull Image from Registry or Hub depending on image name
func Pull(imageName string) (Image, error) {
	image, err := Parse(imageName)
	if err != nil {
		return Image{}, err
	}

	log.Println("pulling image: ", image)

	mURI := fmt.Sprintf("%v/%v/manifests/%v", image.Registry, image.Name, image.Tag)

	client := httpclient.Get()
	request, err := http.NewRequest("GET", mURI, nil)
	response, err := client.Do(request)
	if err != nil {
		return Image{}, fmt.Errorf("retrieving manifest: %v", err)
	}

	if response.StatusCode == http.StatusUnauthorized {
		log.Println("Pull is Unauthorized")
		err := Authenticate(response, request)

		if err != nil {
			return Image{}, fmt.Errorf("authenticating: %v", err)
		}

		response, err = client.Do(request)
		if err != nil {
			return Image{}, fmt.Errorf("retrieving manifest: %v", err)
		}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Image{}, fmt.Errorf("reading manifest body: %v", err)
	}

	if response.StatusCode != 200 {
		switch response.StatusCode {
		case http.StatusUnauthorized:
			return Image{}, xerrors.Unauthorized
		case http.StatusNotFound:
			return Image{}, xerrors.NotFound
		default:
			return Image{}, fmt.Errorf("%d - %s", response.StatusCode, string(body))
		}
	}
	if err := image.parseManifest(body); err != nil {
		return Image{}, fmt.Errorf("parsing manifest: %v", err)
	}

	return image, nil
}

func (image *Image) parseManifest(body []byte) error {

	err := json.Unmarshal(body, &image)

	if err != nil {
		return fmt.Errorf("unmarshalling manifest body: %v", err)
	}

	image.uniqueLayers()
	return nil
}

func (image *Image) uniqueLayers() {
	encountered := map[Layer]bool{}
	result := []Layer{}

	for index := range image.FsLayers {
		if encountered[image.FsLayers[index]] != true {
			encountered[image.FsLayers[index]] = true
			result = append(result, image.FsLayers[index])
		}
	}
	image.FsLayers = result
}