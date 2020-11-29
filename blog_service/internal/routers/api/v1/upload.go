package v1

import (
	"GoProgrammingJourney/blog_service/global"
	"GoProgrammingJourney/blog_service/internal/service"
	"GoProgrammingJourney/blog_service/pkg/app"
	"GoProgrammingJourney/blog_service/pkg/convert"
	"GoProgrammingJourney/blog_service/pkg/errcode"
	"GoProgrammingJourney/blog_service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (up Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})

	return
}
