package bootstrap

import (
	// "bytes"
	// "crypto/aes"
	// "crypto/cipher"
	// "encoding/gob"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "os"
	// "strconv"

	"github.com/TW527E/MineCloudreve/v3/pkg/conf"
	"github.com/TW527E/MineCloudreve/v3/pkg/vol"

	// "github.com/TW527E/MineCloudreve/v3/pkg/request"
	// "github.com/TW527E/MineCloudreve/v3/pkg/util"
	"github.com/gin-gonic/gin"
)

var matrix []byte
var APPID string

// InitApplication 初始化应用常量
func InitApplication() {
	fmt.Print(`
   ___ _                 _                    
  / __\ | ___  _   _  __| |_ __ _____   _____ 
 / /  | |/ _ \| | | |/ _  | '__/ _ \ \ / / _ \	
/ /___| | (_) | |_| | (_| | | |  __/\ V /  __/
\____/|_|\___/ \__,_|\__,_|_|  \___| \_/ \___|

   V` + conf.BackendVersion + `-` + conf.PlusVersion + `-` + conf.Tw527eVersion + `  Commit #` + conf.LastCommit + `  Plus=` + conf.IsPlus + `  TW527E-Edition=` + conf.IsTw527e + `
================================================

`)
	// data, err := ioutil.ReadFile(util.RelativePath(string([]byte{107, 101, 121, 46, 98, 105, 110})))
	// if err != nil {
	// 	util.Log().Panic("%s", err)
	// }

	//table := deSign(data)
	//constant.HashIDTable = table["table"].([]int)
	//APPID = table["id"].(string)
	//matrix = table["pic"].([]byte)
	APPID = `2018110303550773058`
	matrix = []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 2, 10, 0, 0, 1, 16, 8, 6, 0, 0, 0, 209, 100, 92, 53, 0, 0, 6, 112, 73, 68, 65, 84, 120, 156, 237, 214, 209, 142, 212, 54, 0, 64, 81, 187, 223, 4, 69, 162, 255, 255, 93, 230, 97, 50, 187, 147, 196, 119, 89, 84, 169, 165, 234, 57, 66, 66, 12, 153, 141, 99, 59, 222, 59, 191, 126, 251, 107, 141, 15, 173, 177, 214, 28, 99, 172, 49, 231, 124, 255, 116, 141, 49, 142, 127, 206, 211, 135, 99, 172, 57, 199, 124, 253, 254, 241, 175, 247, 207, 158, 151, 174, 49, 94, 175, 93, 107, 60, 110, 53, 199, 203, 173, 198, 26, 235, 254, 217, 26, 99, 30, 223, 63, 127, 56, 199, 120, 25, 219, 219, 253, 47, 151, 62, 199, 186, 249, 161, 99, 141, 251, 248, 231, 186, 223, 107, 141, 57, 142, 63, 207, 75, 199, 115, 50, 207, 227, 31, 99, 158, 198, 244, 28, 234, 58, 223, 107, 55, 166, 199, 211, 31, 227, 127, 157, 255, 199, 55, 215, 188, 204, 255, 28, 227, 50, 211, 99, 173, 121, 251, 248, 121, 233, 245, 251, 235, 184, 240, 190, 126, 231, 89, 121, 60, 235, 121, 79, 188, 173, 255, 120, 93, 171, 143, 238, 191, 110, 123, 229, 88, 213, 235, 19, 92, 198, 52, 114, 82, 31, 207, 243, 153, 231, 127, 95, 207, 211, 253, 215, 126, 255, 94, 215, 127, 191, 84, 253, 252, 99, 94, 118, 213, 118, 169, 30, 243, 113, 29, 211, 110, 254, 31, 235, 127, 25, 192, 118, 81, 251, 253, 217, 191, 127, 113, 255, 117, 93, 235, 113, 159, 255, 231, 11, 112, 154, 148, 15, 246, 255, 230, 249, 55, 155, 242, 56, 127, 110, 175, 197, 126, 255, 205, 57, 206, 75, 117, 188, 63, 215, 93, 241, 193, 249, 51, 175, 247, 223, 158, 95, 235, 184, 199, 79, 95, 170, 237, 248, 247, 75, 85, 251, 239, 57, 212, 203, 250, 199, 253, 215, 237, 252, 252, 236, 82, 237, 222, 191, 62, 191, 119, 231, 199, 120, 94, 121, 157, 255, 203, 248, 173, 223, 63, 181, 126, 239, 191, 107, 62, 241, 2, 30, 99, 157, 183, 115, 109, 126, 249, 246, 125, 237, 254, 227, 250, 178, 239, 39, 225, 163, 75, 47, 155, 224, 237, 192, 189, 140, 247, 111, 77, 194, 175, 109, 162, 253, 38, 156, 247, 159, 154, 191, 4, 206, 7, 240, 207, 54, 193, 53, 2, 238, 191, 196, 226, 16, 61, 14, 246, 243, 237, 118, 27, 246, 119, 142, 184, 205, 33, 178, 221, 152, 247, 251, 111, 131, 173, 94, 226, 205, 47, 241, 109, 196, 125, 244, 18, 237, 38, 96, 179, 169, 183, 251, 247, 250, 11, 244, 229, 251, 187, 3, 119, 31, 140, 215, 91, 109, 158, 255, 109, 175, 238, 46, 189, 63, 255, 24, 243, 118, 136, 61, 246, 218, 252, 220, 253, 71, 29, 120, 191, 16, 236, 231, 31, 26, 247, 223, 205, 223, 216, 30, 10, 255, 233, 253, 191, 157, 191, 62, 127, 106, 78, 119, 231, 215, 245, 210, 218, 255, 189, 127, 199, 245, 195, 199, 223, 187, 243, 123, 156, 173, 99, 173, 230, 237, 251, 251, 253, 119, 63, 171, 227, 253, 25, 251, 181, 218, 142, 127, 140, 177, 198, 206, 203, 92, 175, 245, 54, 203, 251, 107, 199, 237, 217, 120, 55, 191, 254, 249, 189, 230, 13, 0, 248, 159, 251, 227, 223, 30, 0, 0, 240, 251, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 18, 10, 0, 64, 250, 1, 122, 83, 2, 40, 214, 38, 234, 140, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
	vol.ClientSecret = APPID
}

// InitCustomRoute 初始化自定义路由
func InitCustomRoute(group *gin.RouterGroup) {
	group.GET(string([]byte{98, 103}), func(c *gin.Context) {
		c.Header("content-type", "image/png")
		c.Writer.Write(matrix)
	})
	group.GET("id", func(c *gin.Context) {
		c.String(200, APPID)
	})
}

// func deSign(data []byte) map[string]interface{} {
// 	res := decode(data, seed())
// 	dec := gob.NewDecoder(bytes.NewReader(res))
// 	obj := map[string]interface{}{}
// 	err := dec.Decode(&obj)
// 	if err != nil {
// 		util.Log().Panic("您仍在使用旧版的授权文件，请前往 https://pro.cloudreve.org/ 登录下载最新的授权文件")
// 		os.Exit(-1)
// 	}
// 	return checkKeyUpdate(obj)
// }

// func checkKeyUpdate(table map[string]interface{}) map[string]interface{} {
// 	if table["version"].(string) != conf.KeyVersion {
// 		util.Log().Info("正在自动更新授权文件...")
// 		reqBody := map[string]string{
// 			"secret": table["secret"].(string),
// 			"id":     table["id"].(string),
// 		}
// 		reqBodyString, _ := json.Marshal(reqBody)
// 		client := request.NewClient()
// 		resp := client.Request("POST", "https://pro.cloudreve.org/Api/UpdateKey",
// 			bytes.NewReader(reqBodyString)).CheckHTTPResponse(200)
// 		if resp.Err != nil {
// 			util.Log().Panic("授权文件更新失败, %s", resp.Err)
// 		}
// 		keyContent, _ := ioutil.ReadAll(resp.Response.Body)
// 		ioutil.WriteFile(util.RelativePath(string([]byte{107, 101, 121, 46, 98, 105, 110})), keyContent, os.ModePerm)

// 		return deSign(keyContent)
// 	}

// 	return table
// }

// func seed() []byte {
// 	res := []int{8}
// 	s := "20210323"
// 	m := 1 << 20
// 	a := 9
// 	b := 7
// 	for i := 1; i < 23; i++ {
// 		res = append(res, (a*res[i-1]+b)%m)
// 		s += strconv.Itoa(res[i])
// 	}
// 	return []byte(s)
// }

// func decode(cryted []byte, key []byte) []byte {
// 	block, _ := aes.NewCipher(key[:32])
// 	blockSize := block.BlockSize()
// 	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
// 	orig := make([]byte, len(cryted))
// 	blockMode.CryptBlocks(orig, cryted)
// 	orig = pKCS7UnPadding(orig)
// 	return orig
// }

// func pKCS7UnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	unpadding := int(origData[length-1])
// 	return origData[:(length - unpadding)]
// }
