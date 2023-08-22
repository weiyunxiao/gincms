import service from '@/utils/request'

export const useOrgListApi = () => {
	return service.get('admin_api/sys/org_list')
}

export const useOrgApi = (id: Number) => {
	return service.get('admin_api/sys/org?id=' + id)
}

export const useOrgSubmitApi = (dataForm: any) => {
	if (dataForm.id) {
		return service.put('admin_api/sys/org', dataForm)
	} else {
		return service.post('admin_api/sys/org', dataForm)
	}
}
