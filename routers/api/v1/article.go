package v1

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get a single article
// @Tags Article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {
}

// @Summary Search multiple articles
// @Tags Article
// @Produce  json
// @Param key_words body int true "Search key words"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [get]
func SearchArticles(c *gin.Context) {
}

type AddArticleForm struct {
	TagID         int    `form:"tag_id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required;MaxSize(255)"`
	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
	Author        string `form:"author" valid:"Required;MaxSize(100)"`
	CoverImageUrl string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add article
// @Tags Article
// @Produce  json
// @Param tag_id body int true "TagID"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param author body string true "Author"
// @Param state body int true "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
}

type EditArticleForm struct {
	ID            int    `form:"id" valid:"Required;Min(1)"`
	TagID         int    `form:"tag_id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required;MaxSize(255)"`
	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
	CoverImageUrl string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update article
// @Tags Article
// @Produce  json
// @Param id path int true "ID"
// @Param tag_id body string false "TagID"
// @Param title body string false "Title"
// @Param desc body string false "Desc"
// @Param content body string false "Content"
// @Param modified_by body string true "ModifiedBy"
// @Param state body int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context) {
}

// @Summary Delete article
// @Tags Article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
}

//const (
//	QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
//)
//
//func GenerateArticlePoster(c *gin.Context) {
//	appG := app.Gin{C: c}
//	article := &article_service.Article{}
//	qr := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
//	posterName := article_service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
//	articlePoster := article_service.NewArticlePoster(posterName, article, qr)
//	articlePosterBgService := article_service.NewArticlePosterBg(
//		"bg.jpg",
//		articlePoster,
//		&article_service.Rect{
//			X0: 0,
//			Y0: 0,
//			X1: 550,
//			Y1: 700,
//		},
//		&article_service.Pt{
//			X: 125,
//			Y: 298,
//		},
//	)
//
//	_, filePath, err := articlePosterBgService.Generate()
//	if err != nil {
//		appG.Response(http.StatusInternalServerError, e.ERROR_GEN_ARTICLE_POSTER_FAIL, nil)
//		return
//	}
//
//	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
//		"poster_url":      qrcode.GetQrCodeFullUrl(posterName),
//		"poster_save_url": filePath + posterName,
//	})
//}
