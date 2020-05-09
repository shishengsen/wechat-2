/*
 * 纸喵软件
 * Copyright (c) 2017~2020 http://zhimiao.org All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2020/3/3 下午4:26
 */

package api

import (
	"gitee.com/zhimiao/wechat-sdk/miniprogram"
	"gitee.com/zhimiao/wechat/common/utils"
	"gitee.com/zhimiao/wechat/models"
	"gitee.com/zhimiao/wechat/req"
	"gitee.com/zhimiao/wechat/resp"
	ws "gitee.com/zhimiao/wechat/wechat_service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type miniprogramApi struct{}

var Miniprogram = &miniprogramApi{}

// @Summary 小程序列表
// @Tags 小程序信息获取
// @Produce  json
// @Accept  json
// @Param body query req.ListMiniAppParam true "入参集合"
// @Success 200 {array} resp.MiniAppListsVO ""
// @Router /miniprogram/Lists [get]
func (m *miniprogramApi) Lists(c *gin.Context) {
	param := &req.ListMiniAppParam{}
	if err := c.ShouldBind(param); err != nil {
		resp.NewApiResult(-4, utils.Validator(err)).Json(c)
		return
	}
	s := &models.Miniprogram{}
	lists, rows := s.List(param)
	vo := make([]resp.MiniAppListsVO, len(lists))
	for k, v := range lists {
		utils.SuperConvert(&v, &vo[k])
	}
	resp.NewApiResult(1, "读取成功", resp.PageInfo{
		Page:      param.Page,
		PageSize:  param.PageSize,
		TotalSize: rows,
		Rows:      vo,
	}).Json(c)
}

// @Summary 小程序配置
// @Description 小程序配置
// @Tags 小程序控制
// @Produce json
// @Param body body req.MiniAppConfigParam true "body参数"
// @Success 200 {object} resp.ApiResult "{"code": 1,"msg": "操作成功","data": null}"
// @Router /miniprogram/Config [post]
func (m *miniprogramApi) Config(c *gin.Context) {
	param := &req.MiniAppConfigParam{}
	if err := c.ShouldBind(param); err != nil {
		resp.NewApiResult(-4, utils.Validator(err)).Json(c)
		return
	}
	s := &models.Miniprogram{}
	utils.SuperConvert(param, s)
	if ok := s.UpdateByAppID(); ok {
		resp.NewApiResult(1).Json(c)
		return
	}
	resp.NewApiResult(-5, "更新失败请重新尝试").Json(c)
}

// @Summary 获取线上小程序码
// @Tags 小程序信息获取
// @Param MiniProgramID query string true "小程序APPID"
// @Param Page query string false "页面"
// @Param Scene query string false "场景，不得超过32位，不得含有特殊符号"
// @Router /miniprogram/GetWXACodeUnlimit [get]
func (m *miniprogramApi) GetWXACodeUnlimit(c *gin.Context) {
	param := &req.WxACodeParam{}
	if err := c.ShouldBind(param); err != nil {
		resp.NewApiResult(-4, utils.Validator(err)).Json(c)
		return
	}
	maService := ws.GetMiniApp(param.MiniProgramID)
	if maService == nil {
		resp.NewApiResult(-5, "小程序配置无效").Json(c)
		return
	}
	rq := miniprogram.QRCoder{Page: "", Scene: "0"}
	if param.Page != "" {
		rq.Page = param.Page
	}
	if param.Scene != "" {
		rq.Scene = param.Scene
	}
	ret, err := maService.GetWXACodeUnlimit(rq)
	if err != nil {
		logrus.Warn("小程序体验码获取失败,%s", err.Error())
		c.Status(500)
		return
	}
	c.Data(http.StatusOK, "image/jpeg", ret)
}