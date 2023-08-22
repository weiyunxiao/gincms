import service from '@/utils/request'

export const usePostApi = (id: Number) => {
	return service.get('admin_api/sys/post?id=' + id)
}

export const usePostSubmitApi = (dataForm: any) => {
	if (dataForm.id) {
		return service.put('admin_api/sys/post', dataForm)
	} else {
		return service.post('admin_api/sys/post', dataForm)
	}
}

export const usePostListApi = () => {
	return service.get('admin_api/sys/post_list')
}
