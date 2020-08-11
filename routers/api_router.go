package routers

import (
	"aries/handlers/api"
	"aries/middlewares"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (a *ApiRouter) InitApiRouter(rootPath string, router *gin.Engine) {
	articleHandler := api.ArticleHandler{}
	authHandler := api.AuthHandler{}
	categoryHandler := api.CategoryHandler{}
	commentHandler := api.CommentHandler{}
	linkHandler := api.LinkHandler{}
	navHandler := api.NavHandler{}
	sysSettingHandler := api.SysSettingHandler{}
	tagHandler := api.TagHandler{}
	userHandler := api.UserHandler{}

	ArticleApiRouter := router.Group(rootPath, middlewares.JWTAuthMiddleWare())
	{
		ArticleApiRouter.GET("/all_articles", articleHandler.GetAllArticles)
		ArticleApiRouter.GET("/articles/:id", articleHandler.GetArticleById)
		ArticleApiRouter.GET("/articles", articleHandler.GetArticlesByPage)
		ArticleApiRouter.POST("/articles", articleHandler.AddArticle)
		ArticleApiRouter.PUT("/articles", articleHandler.UpdateArticle)
		ArticleApiRouter.DELETE("/articles/:id", articleHandler.DeleteArticle)
		ArticleApiRouter.DELETE("/articles", articleHandler.MultiDelArticles)
		ArticleApiRouter.POST("/articles/files", articleHandler.ImportArticlesFromFiles)
		ArticleApiRouter.POST("/articles/up", articleHandler.MoveArticleUp)
		ArticleApiRouter.POST("/articles/down", articleHandler.MoveArticleDown)
	}

	authApiRouter := router.Group(rootPath)
	{
		authApiRouter.POST("/auth/login", authHandler.Login)
		authApiRouter.POST("/auth/register", authHandler.Register)
		authApiRouter.GET("/auth/captcha", authHandler.CreateCaptcha)
		authApiRouter.POST("/auth/pwd/forget", authHandler.ForgetPwd)
		authApiRouter.POST("/auth/pwd/reset", authHandler.ResetPwd)
	}

	categoryApiRouter := router.Group(rootPath, middlewares.JWTAuthMiddleWare())
	{
		categoryApiRouter.GET("/all_categories", categoryHandler.GetAllCategories)
		categoryApiRouter.GET("/parent_categories", categoryHandler.GetAllParentCategories)
		categoryApiRouter.GET("/categories", categoryHandler.GetCategoriesByPage)
		categoryApiRouter.POST("/categories/article", categoryHandler.AddArticleCategory)
		categoryApiRouter.PUT("/categories/article", categoryHandler.UpdateArticleCategory)
		categoryApiRouter.POST("/categories/link", categoryHandler.AddLinkCategory)
		categoryApiRouter.PUT("/categories/link", categoryHandler.UpdateLinkCategory)
		categoryApiRouter.DELETE("/categories/:id", categoryHandler.DeleteCategory)
		categoryApiRouter.DELETE("/categories", categoryHandler.MultiDelCategories)
	}

	commentRouter := router.Group(rootPath, middlewares.JWTAuthMiddleWare())
	{
		commentRouter.GET("/all_comments", commentHandler.GetAllComments)
		commentRouter.GET("/comments", commentHandler.GetCommentsByPage)
		commentRouter.POST("/comments", commentHandler.AddComment)
		commentRouter.PUT("/comments", commentHandler.UpdateComment)
		commentRouter.DELETE("/comments/:id", commentHandler.DeleteComment)
		commentRouter.DELETE("/comments", commentHandler.MultiDelComments)
	}

	linkApiRouter := router.Group(rootPath, middlewares.JWTAuthMiddleWare())
	{
		linkApiRouter.GET("/all_links", linkHandler.GetAllLinks)
		linkApiRouter.GET("/links", linkHandler.GetLinksByPage)
		linkApiRouter.POST("/links", linkHandler.CreateLink)
		linkApiRouter.PUT("/links", linkHandler.UpdateLink)
		linkApiRouter.DELETE("/links/:id", linkHandler.DeleteLink)
		linkApiRouter.DELETE("/links", linkHandler.MultiDelLinks)
	}

	navApiRouter := router.Group(rootPath)
	{
		navApiRouter.GET("/navs", navHandler.GetAllNavs)
		navApiRouter.POST("/navs", navHandler.CreateNav)
		navApiRouter.PUT("/navs", navHandler.UpdateNav)
		navApiRouter.DELETE("/navs/:id", navHandler.DeleteNav)
		navApiRouter.DELETE("/navs", navHandler.MultiDelNavs)
	}

	sysSettingApiRouter := router.Group(rootPath, middlewares.JWTAuthMiddleWare())
	{
		sysSettingApiRouter.GET("/sys_setting/items", sysSettingHandler.GetSysSettingItem)
		sysSettingApiRouter.POST("/sys_setting/site", sysSettingHandler.SaveSiteSetting)
		sysSettingApiRouter.POST("/sys_setting/smtp", sysSettingHandler.SaveSMTPSetting)
		sysSettingApiRouter.POST("/sys_setting/pic_bed", sysSettingHandler.SavePicBedSetting)
		sysSettingApiRouter.POST("/sys_setting/email/test", sysSettingHandler.SendTestEmail)
		sysSettingApiRouter.GET("/sys_setting/index_info", sysSettingHandler.GetAdminIndexData)
	}

	tagApiRouter := router.Group(rootPath, middlewares.JWTAuthMiddleWare())
	{
		tagApiRouter.GET("/all_tags", tagHandler.GetAllTags)
		tagApiRouter.GET("/tags", tagHandler.GetTagsByPage)
		tagApiRouter.GET("/tags/:id", tagHandler.GetTagById)
		tagApiRouter.POST("/tags", tagHandler.AddTag)
		tagApiRouter.PUT("/tags", tagHandler.UpdateTag)
		tagApiRouter.DELETE("/tags/:id", tagHandler.DeleteTag)
		tagApiRouter.DELETE("/tags", tagHandler.MultiDelTags)
	}

	userApiRouter := router.Group(rootPath)
	{
		userApiRouter.GET("/all_users", userHandler.GetAllUsers)
		userApiRouter.PUT("/users", userHandler.UpdateUser)
		userApiRouter.PUT("/users/pwd", userHandler.UpdateUserPwd)
	}
}
