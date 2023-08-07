package router

import (
	"gincms/app/http/adminapi/controller"
	sysController "gincms/app/http/adminapi/controller/sys"
	"gincms/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func AdminApiRouter(r *gin.Engine) {
	route := r.Group("admin_api")
	routeNeedJwt := r.Group("admin_api", middleware.JWTCheck()) //需要jwt验证

	/************无需jwt验证***************/
	route.GET("sys/auth_captcha_enabled", controller.PublicCtl.LoginCaptchaEnabled) //是否开启登录需要验证码
	route.GET("sys/auth_captcha", controller.PublicCtl.Captcha)                     //验证码
	route.POST("sys/auth_login", controller.PublicCtl.Login)                        //登录
	/************无需jwt验证 end***************/

	/************后台系统进入前端需要调用的***************/
	routeNeedJwt.POST("sys/auth_logout", sysController.AuthCtl.Logout)       //用户退出
	routeNeedJwt.GET("sys/user_info", sysController.UserCtl.Info)            //登录进入后，获取用户信息
	routeNeedJwt.GET("sys/menu_authority", sysController.MenuCtl.Authority)  //获取用户的权限
	routeNeedJwt.GET("sys/menu_nav", sysController.MenuCtl.Nav)              //获取用户的菜单
	routeNeedJwt.GET("sys/dict_type_all", sysController.DictCtl.DictTypeAll) //获取系统的字典数据
	/************后台系统进入前端需要调用的 end***************/

	/************岗位管理***************/
	routeNeedJwt.POST("sys/post", sysController.PostCtl.AddPost)      //添加一个岗位
	routeNeedJwt.PUT("sys/post", sysController.PostCtl.UpdatePost)    //更新一个岗位
	routeNeedJwt.DELETE("sys/post", sysController.PostCtl.DelPost)    //删除岗位
	routeNeedJwt.GET("sys/post", sysController.PostCtl.GetPost)       //获取一个岗位信息
	routeNeedJwt.GET("sys/post_page", sysController.PostCtl.PostPage) //岗位分页列表
	routeNeedJwt.GET("sys/post_list", sysController.PostCtl.PostList) //岗位列表
	/************岗位管理 end***************/

	/************机构部门管理***************/
	routeNeedJwt.GET("sys/org_list", sysController.OrgCtl.OrgList) //机构部门列表
	routeNeedJwt.GET("sys/org", sysController.OrgCtl.GetOrg)       //得到一条机构记录
	routeNeedJwt.POST("sys/org", sysController.OrgCtl.AddOrg)      //添加一条机构记录
	routeNeedJwt.PUT("sys/org", sysController.OrgCtl.UpdateOrg)    //修改一条机构记录
	routeNeedJwt.DELETE("sys/org", sysController.OrgCtl.DelOrg)    //删除机构
	/************机构部门管理 end***************/

	/************角色管理***************/
	routeNeedJwt.POST("sys/role", sysController.RoleCtl.AddRole)      //添加一个角色
	routeNeedJwt.PUT("sys/role", sysController.RoleCtl.UpdateRole)    //更新
	routeNeedJwt.DELETE("sys/role", sysController.RoleCtl.DelRole)    //删除
	routeNeedJwt.GET("sys/role", sysController.RoleCtl.GetRole)       //获取一个角色
	routeNeedJwt.GET("sys/role_list", sysController.RoleCtl.RoleList) //角色列表
	routeNeedJwt.GET("sys/role_page", sysController.RoleCtl.RolePage) //角色带分页列表

	routeNeedJwt.GET("sys/role_menu", sysController.RoleCtl.RoleMenu)          //添加编辑角色时，显示的所有系统菜单
	routeNeedJwt.GET("sys/role_user_page", sysController.RoleCtl.RoleUserPage) //某个角色拥有的用户分页列表
	routeNeedJwt.POST("sys/role_user", sysController.RoleCtl.AddRoleUser)      //某个角色关联多个用户的操作
	routeNeedJwt.DELETE("sys/role_user", sysController.RoleCtl.DelRoleUser)    //移除某个角色下的用户

	/************角色管理 end***************/
	/************用户管理***************/
	routeNeedJwt.POST("sys/user", sysController.UserCtl.AddUser)      //添加一个用户
	routeNeedJwt.GET("sys/user", sysController.UserCtl.GetUser)       //获取一个用户
	routeNeedJwt.PUT("sys/user", sysController.UserCtl.UpdateUser)    //更新一个用户
	routeNeedJwt.DELETE("sys/user", sysController.UserCtl.DelUser)    //删除用户
	routeNeedJwt.GET("sys/user_page", sysController.UserCtl.UserPage) //用户带分页列表
	/************用户管理 end***************/

	{
		route.GET("/system_dir", controller.FileCtl.DirList)
	}
}
