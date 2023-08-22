import service from '@/utils/request'

export const useMenuNavApi = () => {
	return service.get('admin_api/sys/menu_nav')
}

export const useAuthorityListApi = () => {
	return service.get('admin_api/sys/menu_authority')
}

export const useMenuListApi = (type: Number) => {
	// 菜单类型 0：菜单  1：按钮  2：接口
	const menuType = type === 2 ? 2 : 0

	return service.get('admin_api/sys/menu_list?type=' + menuType)
}

export const useMenuApi = (id: Number) => {
	return service.get('admin_api/sys/menu?id=' + id)
}

export const useMenuSubmitApi = (dataForm: any) => {
	if (dataForm.id) {
		return service.put('admin_api/sys/menu', dataForm)
	} else {
		return service.post('admin_api/sys/menu', dataForm)
	}
}
