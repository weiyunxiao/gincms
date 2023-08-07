package comdata

import "gincms/app/model"

// TreeMenu 生成菜单树
func TreeMenu(result []model.SysMenu, id uint) (tree []model.SysMenu) {
	tree = make([]model.SysMenu, 0)
	// 遍历结果集
	for _, item := range result {
		if item.Pid == id {
			node := model.SysMenu{}
			node.ID = item.ID
			node.Pid = item.Pid
			node.Name = item.Name
			node.URL = item.URL
			node.Type = item.Type
			node.OpenStyle = item.OpenStyle
			node.Icon = item.Icon
			node.Authority = item.Authority
			node.Sort = item.Sort
			node.CreateTime = item.CreateTime
			node.Children = TreeMenu(result, item.ID)
			tree = append(tree, node)
		}
	}
	return tree
}

// TreeOrg 组织架构菜单树
func TreeOrg(result []model.SysOrg, pid uint) (tree []model.SysOrg) {
	tree = make([]model.SysOrg, 0)
	for _, item := range result {
		if item.Pid == pid {
			node := model.SysOrg{}
			node = item
			node.Children = TreeOrg(result, item.ID)
			if item.Pid == 0 {
				node.ParentName = "一级机构"
			} else {
				node.ParentName = node.ParenOrg.Name
			}
			tree = append(tree, node)
		}
	}
	return
}
