import service from '@/utils/request'

export const useAttachmentSubmitApi = (dataForm: any) => {
	return service.post('admin_api/sys/attachment', dataForm)
}
