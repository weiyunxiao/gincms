import service from '@/utils/request'

export const useCaptchaEnabledApi = () => {
	return service.get('admin_api/sys/auth_captcha_enabled')
}

export const useCaptchaApi = () => {
	return service.get('admin_api/sys/auth_captcha')
}

export const useSendCodeApi = (mobile: string) => {
	return service.post('/sys/auth/send/code?mobile=' + mobile)
}

export const useAccountLoginApi = (data: any) => {
	return service.post('admin_api/sys/auth_login', data)
}

export const useMobileLoginApi = (data: any) => {
	return service.post('/sys/auth/mobile', data)
}

export const useLogoutApi = () => {
	return service.post('admin_api/sys/auth_logout')
}
