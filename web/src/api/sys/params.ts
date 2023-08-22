import service from '@/utils/request'

export const useParamsApi = (id: number) => {
	return service.get('admin_api/sys/params?id=' + id)
}

export const useParamsSubmitApi = (dataForm: any) => {
	if (dataForm.id) {
		return service.put('admin_api/sys/params', dataForm)
	} else {
		return service.post('admin_api/sys/params', dataForm)
	}
}
