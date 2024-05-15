package helper

import (
	"Libeery/database"
	model "Libeery/model"

	"github.com/gin-gonic/gin"
)

func GetAllAcara(c *gin.Context) ([]*model.MsAcara, error) {
	db := database.GlobalDB
	var acara []*model.MsAcara
	if err := db.Find(&acara).Error; err != nil {
		return nil, err
	}

	return acara, nil
}

func GetAcaraPagination(c *gin.Context, page, take int) ([]*model.MsAcara, error) {
	db := database.GlobalDB
	var acara []*model.MsAcara
	offset := (page - 1) * take

	if err := db.Offset(offset).Limit(take).Find(&acara).Error; err != nil {
		return nil, err
	}
	return acara, nil
}

func GetAcaraDetails(c *gin.Context, id int) (*model.MsAcara, error) {
	db := database.GlobalDB
	var acara model.MsAcara
	if err := db.Where("acara_id = ?", id).First(&acara).Error; err != nil {
		return nil, err
	}
	return &acara, nil
}

func CreateAcara(c *gin.Context, reqBody model.MsAcaraRequestDTO) error {
	db := database.GlobalDB
	acara := model.MsAcara{
		AcaraName:      reqBody.AcaraName,
		AcaraStartTime: reqBody.AcaraStartTime,
		AcaraEndTime:   reqBody.AcaraEndTime,
		AcaraDate:      reqBody.AcaraDate,
		AcaraLocation:  reqBody.AcaraLocation,
		AcaraDetails:   reqBody.AcaraDetails,
		SpeakerName:    reqBody.SpeakerName,
		RegisterLink:   reqBody.RegisterLink,
		AcaraImage:     reqBody.AcaraImage,
		Stsrc:          "A",
	}

	if err := db.Create(&acara).Error; err != nil {
		return err
	}
	return nil
}
