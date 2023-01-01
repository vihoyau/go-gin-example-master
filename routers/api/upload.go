package api

import (
	"context"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/gcos"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage2(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}

func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}

	file, Header, err := c.Request.FormFile("file")
	filename := Header.Filename
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	client := gcos.Setup()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Upload the file to COS
	v, err := client.Object.Put(context.Background(), filename, file, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		panic(err)
	}
	gcos.Log_status(err)
	fmt.Printf("Case2 done, %v\n", v)
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      v.Request.URL.Path,
		"image_save_url": v.Request.URL.Host + v.Request.URL.Path,
	})
}
