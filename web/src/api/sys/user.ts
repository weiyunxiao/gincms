import service from '@/utils/request'
import cache from '@/utils/cache'
import constant from '@/utils/constant'

export const useUserInfoApi = () => {
	return service.get('admin_api/sys/user_info')
}

export const updatePasswordApi = (data: any) => {
	return service.put('admin_api/sys/user_info', data)
}

export const useUserApi = (id: number) => {
	return service.get('admin_api/sys/user?id=' + id)
}

export const useUserExportApi = () => {
	location.href = constant.apiUrl + 'admin_api/sys/user/export?access_token=' + cache.getToken()
}

export const useUserSubmitApi = (dataForm: any) => {
	if (dataForm.status !== undefined) {
		dataForm.status = Number(dataForm.status)
	}
	if (dataForm.id) {
		return service.put('admin_api/sys/user', dataForm)
	} else {
		dataForm.status = Number(dataForm.status)
		return service.post('admin_api/sys/user', dataForm)
	}
}
