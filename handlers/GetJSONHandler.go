package handlers

import (
	"GoReadNote/logger"
	"GoReadNote/sprider"
	"github.com/gin-gonic/gin"
)

//返回Json的一个模板  Code在不同情况下有不同作用
type JsonRet struct {
	Code int         `json:"code"`
	List interface{} `json:"list"`
}

//下面是个基本的例子
func GetJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetJsonHandler")
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	holder := JsonHolder{Id: 1, Name: "123"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(200, holder)

	return
}

func GetSearchNoteJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchNoteJsonHandler")

	h := gin.H{}

	noteName, exist := c.GetQuery("notename")
	if !exist {
		c.JSON(500, h)
		return
	}
	noteListMap, find := sprider.SearchNoteByName(noteName)
	if !find {
		//没有找到 再试试直接Get
		chptMap, ok := sprider.GetNoteChapterListByNoteName(noteName)
		if !ok {
			c.JSON(500, h)
			return
		}
		var cpInfo []sprider.ChapterInfo

		for i := 1; i <= len(chptMap); i++ {
			cpInfo = append(cpInfo, chptMap[i])
		}
		//code = 0 为一个结果  code = 1为小说列表
		retJson := JsonRet{Code: 0, List: cpInfo}
		c.JSON(200, retJson)
		return
	}
	var noteInfo []sprider.SearchNote

	for i := 1; i <= len(noteListMap); i++ {
		noteInfo = append(noteInfo, noteListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, List: noteInfo}
	c.JSON(200, retJson)
	return
}

func GetBookContentJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchNoteJsonHandler")
	h := gin.H{}

	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
		return
	}
	url = sprider.URL + url
	logger.ALogger().Debug("url = ", url)
	chp := sprider.GetNoteContent(url)
	if chp == nil {
		c.JSON(500, h)
		return
	}

	c.JSON(200, chp)
	return

}

func GetTopNoteListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTopNoteListJsonHandler")
	h := gin.H{}
	type JsonRet struct {
		Code int         `json:"code"` //code = 0 为一个结果  code = 1为小说列表
		List interface{} `json:"list"`
	}
	noteListMap, ok := sprider.GetTopNoteList()
	if !ok {
		c.JSON(500, h)
		return
	}
	var noteInfo []sprider.TopNote

	for i := 1; i <= len(noteListMap); i++ {
		noteInfo = append(noteInfo, noteListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, List: noteInfo}
	c.JSON(200, retJson)
	return

}
