import service from '@/utils/request'

export const useDictTypeAllApi = () => {
	return service.get('admin_api/sys/dict_type_all')
}

export const useDictTypeApi = (id: Number) => {
	return service.get('admin_api/sys/dict_type?id=' + id)
}

export const useDictTypeSubmitApi = (dataForm: any) => {
	if (dataForm.id) {
		return service.put('admin_api/sys/dict_type', dataForm)
	} else {
		return service.post('admin_api/sys/dict_type', dataForm)
	}
}

export const useDictDataApi = (id: Number) => {
	return service.get('admin_api/sys/dict_data?id=' + id)
}

export const useDictDataSubmitApi = (dataForm: any) => {
	if (dataForm.id) {
		return service.put('admin_api/sys/dict_data', dataForm)
	} else {
		return service.post('admin_api/sys/dict_data', dataForm)
	}
}
