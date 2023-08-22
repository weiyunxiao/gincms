import service from '@/utils/request'

export const useRoleMenuApi = () => {
	return service.get('admin_api/sys/role_menu')
}

export const useRoleApi = (id: number) => {
	return service.get('admin_api/sys/role?id=' + id)
}

export const useRoleListApi = () => {
	return service.get('admin_api/sys/role_list')
}

export const useRoleSubmitApi = (dataForm: any) => {
	if (dataForm.id) {
		return service.put('admin_api/sys/role', dataForm)
	} else {
		return service.post('admin_api/sys/role', dataForm)
	}
}

export const useRoleDataScopeSubmitApi = (dataForm: any) => {
	return service.put('admin_api/sys/role_data-scope', dataForm)
}

export const useRoleUserSubmitApi = (roleId: number, dataForm: any) => {
	return service.post('admin_api/sys/role_user?roleId=' + roleId, { id_arr: dataForm })
}
