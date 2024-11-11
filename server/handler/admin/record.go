package admin

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/record"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取提交记录信息（提交信息+评测结果）
func AdminRecordInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 获取提交信息
	sid := uint64(id)
	records, err := record.SelectBySubmissionId(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}

// 获取提交记录列表
func AdminRecordList(c *gin.Context) {
	records, err := record.SelectAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}

// 删除提交记录（提交信息+评测结果）
func AdminRecordRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	sid := uint64(id)
	err = record.DeleteBySubmissionId(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
