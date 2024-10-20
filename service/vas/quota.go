package vas

import (
	model "github.com/TW527E/MineCloudreve/v3/models"
	"github.com/TW527E/MineCloudreve/v3/pkg/serializer"
	"github.com/gin-gonic/gin"
)

// GeneralVASService 通用增值服务
type GeneralVASService struct {
}

// Quota 获取容量配额信息
func (service *GeneralVASService) Quota(c *gin.Context, user *model.User) serializer.Response {
	packs := user.GetAvailableStoragePacks()
	return serializer.BuildUserQuotaResponse(user, packs)
}
