package v1

import (
	"example/models"
	"example/pkg/e"
	"example/pkg/settings"
	"example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetAnArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams

	var data interface{}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "参数：" + err.Key + "错误， " + err.Message,
				"data": make(map[string]string),
			})
			return
		}
	}

	if models.ExistArticleByID(id) {
		data = models.GetAnArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ErrorNotExistArticle
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	code := e.InvalidParams
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "参数：" + err.Key + "错误， " + err.Message,
				"data": make(map[string]string),
			})
			return
		}
	}

	code = e.SUCCESS
	data["lists"] = models.GetArticles(util.GetPage(c), settings.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddAnArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.Query("state")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("必须大于0")
	valid.Required(title, "title").Message("不能为空")
	valid.Required(desc, "desc").Message("不能为空")
	valid.Required(content, "content").Message("不能为空")
	valid.Required(createdBy, "createdBy").Message("不能为空")
	valid.Required(title, "title").Message("不能为空")
	valid.Range(state, 0, 1, "state").Message("只能为0或1")

	code := e.InvalidParams
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "参数：" + err.Key + "错误， " + err.Message,
				"data": make(map[string]string),
			})
			return
		}
	}

	if models.ExistArticleByTitle(title) {
		code = e.ErrorExistArticle

	} else {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddAnArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ErrorNotExistTag
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func EditAnArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(title, "title").Message("不能为空")
	valid.Required(desc, "desc").Message("不能为空")
	valid.Required(content, "content").Message("不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.InvalidParams

	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagId) {
				data := make(map[string]interface{})

				data["tag_id"] = tagId
				data["title"] = title
				data["desc"] = desc
				data["content"] = content
				data["modified_by"] = modifiedBy

				models.EditAnArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ErrorNotExistTag
			}
		} else {
			code = e.ErrorNotExistArticle
		}
	} else {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "参数:" + err.Key + "错误，" + err.Message,
				"data": make(map[string]string),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteAnArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.InvalidParams

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistArticleByID(id) {
			models.DeleteAnArticle(id)
		} else {
			code = e.ErrorNotExistArticle
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
