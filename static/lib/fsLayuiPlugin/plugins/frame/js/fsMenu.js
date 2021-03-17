/**
 * @Description: 菜单配置
 * @Copyright: 2017 www.fallsea.com Inc. All rights reserved.
 * @author: fallsea
 * @version 1.6.4
 * @License：MIT
 */
layui.define(['element',"fsConfig","fsCommon"], function(exports){
	
	var menuConfig = {
			dataType : "server" , //获取数据方式，local本地获取，server 服务端获取
			loadUrl : "/menu/getMenu", //加载数据地址
			rootMenuId : "0", //根目录菜单id
			defaultSelectTopMenuId : "1", //默认选中头部菜单id
			defaultSelectLeftMenuId : "28", //默认选中左边菜单id
			menuIdField : "menuId", //菜单id
			menuNameField : "menuName", //菜单名称
			menuIconField : "menuIcon" , //菜单图标，图标必须用css
			menuHrefField : "menuHref" , //菜单链接
			parentMenuIdField : "parentMenuId" ,//父菜单id
			data : [
				// {"menuId":"1","menuName":"海房项目库","menuIcon":"fa-cog","menuHref":"","parentMenuId":"0"},
				// {"menuId":"2","menuName":"移民项目库","menuIcon":"","menuHref":"","parentMenuId":"0"},
				// {"menuId":"3","menuName":"系统配置","menuIcon":"","menuHref":"","parentMenuId":"0"},
				// {"menuId":"4","menuName":"功能配置","menuIcon":"","menuHref":"","parentMenuId":"0"},
				// {"menuId":"8","menuName":"后台首页","menuIcon":"","menuHref":"/home/index/index","parentMenuId":"1"},
				// {"menuId":"11","menuName":"海房项目","menuIcon":"","menuHref":"","parentMenuId":"1"},
				// {"menuId":"12","menuName":"海房配置","menuIcon":"","menuHref":"","parentMenuId":"1"},
                // {"menuId":"13","menuName":"爬虫库配置","menuIcon":"","menuHref":"","parentMenuId":"1"},
				// {"menuId":"21","menuName":"移民项目","menuIcon":"","menuHref":"","parentMenuId":"2"},
				// {"menuId":"22","menuName":"移民配置","menuIcon":"","menuHref":"","parentMenuId":"2"},
				// {"menuId":"31","menuName":"工具配置","menuIcon":"","menuHref":"","parentMenuId":"3"},
				// {"menuId":"32","menuName":"用户管理","menuIcon":"","menuHref":"","parentMenuId":"3"},
				// {"menuId":"111","menuName":"海外房产","menuIcon":"&#xe68e;","menuHref":"/estate/home/index","parentMenuId":"11"},
				// {"menuId":"112","menuName":"海外考察","menuIcon":"&#xe68e;","menuHref":"/estate/investigate/index","parentMenuId":"11"},
                // {"menuId":"113","menuName":"二手房产","menuIcon":"&#xe68e;","menuHref":"/estate/secondHouse/index?type=2","parentMenuId":"11"},
				// {"menuId":"121","menuName":"国家配置","menuIcon":"<i class=\"layui-icon\">&#xe614;</i>","menuHref":"/config/country/index","parentMenuId":"12"},
				// {"menuId":"122","menuName":"热门配置","menuIcon":"<i class=\"layui-icon\">&#xe756;</i>","menuHref":"/estate/recommend/index","parentMenuId":"12"},
				// {"menuId":"123","menuName":"搜索词配置","menuIcon":"<i class=\"layui-icon\">&#xe615;</i>","menuHref":"/estate/search_recommend/index","parentMenuId":"12"},
                // {"menuId":"131","menuName":"爬虫二手房","menuIcon":"<i class=\"layui-icon\">&#xe614;</i>","menuHref":"/estate/secondHouse/index?type=1","parentMenuId":"13"},
				// {"menuId":"211","menuName":"海外服务","menuIcon":"&#xe68e;","menuHref":"/project/Oversea_Service/index","parentMenuId":"21"},
				// /*{"menuId":"212","menuName":"海外考察","menuIcon":"&#xe68e;","menuHref":"views/home/index.html","parentMenuId":"21"},
				// {"menuId":"221","menuName":"国家配置","menuIcon":"<i class=\"layui-icon\">&#xe652;</i>","menuHref":"404.html","parentMenuId":"22"},
				// {"menuId":"222","menuName":"热门配置","menuIcon":"<i class=\"layui-icon\">&#xe756;</i>","menuHref":"404.html","parentMenuId":"22"},*/
				// {"menuId":"331","menuName":"日志查询","menuIcon":"<i class=\"layui-icon\">&#xe614;</i>","menuHref":"/system/log/index","parentMenuId":"31"},
				// {"menuId":"332","menuName":"SQL查询","menuIcon":"<i class=\"layui-icon\">&#xe614;</i>","menuHref":"/system/sql/index","parentMenuId":"31"},
				// {"menuId":"333","menuName":"发布代码","menuIcon":"<i class=\"layui-icon\">&#xe609;</i>","menuHref":"/system/git/index","parentMenuId":"31"},
				// {"menuId":"334","menuName":"用户管理","menuIcon":"<i class=\"layui-icon\">&#xe770;</i>","menuHref":"/home/account/index","parentMenuId":"32"},
				// {"menuId":"335","menuName":"角色管理","menuIcon":"<i class=\"layui-icon\">&#xe613;</i>","menuHref":"/admin/admin/admincate","parentMenuId":"32"},
				// {"menuId":"336","menuName":"权限管理","menuIcon":"<i class=\"layui-icon\">&#xe672;</i>","menuHref":"/admin/menu/index","parentMenuId":"32"},
				// {"menuId":"41","menuName":"APP配置","menuIcon":"","menuHref":"","parentMenuId":"4"},
				// {"menuId":"332","menuName":"APP首页icon","menuIcon":"<i class=\"layui-icon\">&#xe614;</i>","menuHref":"/config/appicon/index","parentMenuId":"41"},
		] //本地数据
	};
	
	var element = layui.element,
	fsCommon = layui.fsCommon,
	fsConfig = layui.fsConfig,
	statusName = $.result(fsConfig,"global.result.statusName","errorNo"),
  msgName = $.result(fsConfig,"global.result.msgName","errorInfo"),
  dataName = $.result(fsConfig,"global.result.dataName","results.data"),
	FsMenu = function (){
		
	};
	
	
	FsMenu.prototype.render = function(){
		
		this.loadData();
		
		this.showMenu();
	};
	
	/**
	 * 加载数据
	 */
	FsMenu.prototype.loadData = function(){
		
		if(menuConfig.dataType == "server"){//服务端拉取数据
			
			var url = menuConfig.loadUrl;
			if($.isEmpty(url)){
				fsCommon.errorMsg("未配置请求地址！");
				return;
			}
			
			fsCommon.invoke(url,{},function(data){
  			if(data[statusName] == "0")
  			{
  				menuConfig.data = $.result(data,dataName);
  			}
  			else
  			{
  				//提示错误消息
  				fsCommon.errorMsg(data[msgName]);
  			}
  		},false);
			
		}
		
	}
	
	
	/**
	 * 获取图标
	 */
	FsMenu.prototype.getIcon = function(menuIcon){
		
		if(!$.isEmpty(menuIcon)){
			
			if(menuIcon.indexOf("<i") == 0){
				return menuIcon;
			}else if (menuIcon.indexOf("&#") == 0){
				return '<i class="layui-icon">'+menuIcon+'</i>';
			}else if (menuIcon.indexOf("fa-") == 0){
				return '<i class="fa '+menuIcon+'"></i>';
			}else {
				return '<i class="'+menuIcon+'"></i>';
			}
		}
		return "";
	};
	
	/**
	 * 显示菜单
	 */
	FsMenu.prototype.showMenu = function(){
		var thisMenu = this;
		var data = menuConfig.data;
		if(!$.isEmpty(data)){
			//显示顶部一级菜单
			var fsTopMenuElem = $("#fsTopMenu");
			var fsLeftMenu = $("#fsLeftMenu");
			$.each(data,function(i,v){
				if(menuConfig.rootMenuId === v[menuConfig.parentMenuIdField]){
					var topStr = '<li class="layui-nav-item';
					if(!$.isEmpty(menuConfig.defaultSelectTopMenuId) && menuConfig.defaultSelectTopMenuId == v[menuConfig.menuIdField]){//默认选中处理
						topStr += ' layui-this';
					}
					topStr += '" dataPid="'+v[menuConfig.menuIdField]+'"><a href="javascript:;">'+thisMenu.getIcon(v[menuConfig.menuIconField])+' <cite>'+v[menuConfig.menuNameField]+'</cite></a></li>';
					fsTopMenuElem.append(topStr);
					
					//显示二级菜单，循环判断是否有子栏目
					$.each(data,function(i2,v2){
						if(v[menuConfig.menuIdField] === v2[menuConfig.parentMenuIdField]){
							
							var menuRow = '<li class="layui-nav-item';
							if(!$.isEmpty(menuConfig.defaultSelectLeftMenuId) && menuConfig.defaultSelectLeftMenuId == v2[menuConfig.menuIdField]){//默认选中处理
								menuRow += ' layui-this';
							}
							//显示三级菜单，循环判断是否有子栏目
							var menuRow3 = "";
							$.each(data,function(i3,v3){
								if(v2[menuConfig.menuIdField] === v3[menuConfig.parentMenuIdField]){
									if($.isEmpty(menuRow3)){
										menuRow3 = '<dl class="layui-nav-child">';
									}
									menuRow3 += '<dd';
									if(!$.isEmpty(menuConfig.defaultSelectLeftMenuId) && menuConfig.defaultSelectLeftMenuId == v3[menuConfig.menuIdField]){//默认选中处理
										menuRow3 += ' class="layui-this"';
										menuRow += ' layui-nav-itemed';//默认展开二级菜单
									}
									
									menuRow3 += ' lay-id="'+v3[menuConfig.menuIdField]+'"><a href="javascript:;" menuId="'+v3[menuConfig.menuIdField]+'" dataUrl="'+v3[menuConfig.menuHrefField]+'">'+thisMenu.getIcon(v3[menuConfig.menuIconField])+' <cite>'+v3[menuConfig.menuNameField]+'</cite></a></dd>';
									
								}
								
							});
							
							menuRow += '" lay-id="'+v2[menuConfig.menuIdField]+'" dataPid="'+v2[menuConfig.parentMenuIdField]+'" style="display: none;"><a href="javascript:;" menuId="'+v2[menuConfig.menuIdField]+'" dataUrl="'+v2[menuConfig.menuHrefField]+'">'+thisMenu.getIcon(v2[menuConfig.menuIconField])+' <cite>'+v2[menuConfig.menuNameField]+'</cite></a>';
							
							
							if(!$.isEmpty(menuRow3)){
								menuRow3 += '</dl>';
								
								menuRow += menuRow3;
							}
							
							menuRow += '</li>';
							
							fsLeftMenu.append(menuRow);
						}
						
					});
					
				}
			});
		}
		element.render("nav");
	};
	
	var fsMenu = new FsMenu();
	exports("fsMenu",fsMenu);
});

