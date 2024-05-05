package controllers

import (
	"goweb/logger"
	"goweb/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 索引页
func IndexGet(c *gin.Context) {
	articleList, err := models.QueryAllArticle()
	if err != nil {
		logger.Error("models.QueryCurrUserArticleWithPage failed", zap.Any("error", err))
	}
	logger.Debug("models.QueryCurrUserArticleWithPage", zap.Any("articleList", articleList))
	c.HTML(http.StatusOK, "index.html", gin.H{"articleList": articleList})
}
