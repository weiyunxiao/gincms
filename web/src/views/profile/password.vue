<template>
	<el-card shadow="never" header="个人信息">
		<el-form ref="dataFormRef" :model="dataForm" :rules="dataRules" label-width="200px" @keyup.enter="handleDataForm()">
			<el-form-item label="账号">
				<el-input v-model="username" disabled></el-input>
				<div class="el-form-item-msg">账号信息用于登录，系统不允许修改</div>
			</el-form-item>

			<el-form-item label="用户头像">
				<el-upload ref="uploadRef" v-model:file-list="fileList" :action="constant.uploadUrl"
					list-type="picture-card" :on-preview="handlePictureCardPreview" :on-remove="handleRemove" :limit="1"
					:on-success="handleSuccess">
					<el-icon>
						<Plus />
					</el-icon>
				</el-upload>
			</el-form-item>
			<el-dialog v-model="dialogVisible">
				<img w-full :src="dialogImageUrl" alt="Preview Image" />
			</el-dialog>
			<el-form-item prop="password" label="原密码">
				<el-input v-model="dataForm.password" type="password"></el-input>
			</el-form-item>
			<el-form-item prop="newPassword" label="新密码">
				<el-input v-model="dataForm.newPassword" type="password"></el-input>
			</el-form-item>
			<el-form-item prop="confirmPassword" label="确认密码">
				<el-input v-model="dataForm.confirmPassword" type="password"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="handleDataForm">确定</el-button>
			</el-form-item>
		</el-form>
	</el-card>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { useRouter, useRoute } from 'vue-router'
import { validatePassword } from '@/utils/validate'
import { updatePasswordApi } from '@/api/sys/user'
import { ElMessage } from 'element-plus'
import store from '@/store'
import { closeTab } from '@/utils/tabs'
import type { UploadProps, UploadUserFile } from 'element-plus'
import type { UploadInstance } from 'element-plus'
import constant from '@/utils/constant'
const username = store.userStore.user.username
const fileList = ref<UploadUserFile[]>([
	{
		name: username,
		url: store.userStore.user.avatar,
	}
])
const uploadRef = ref<UploadInstance>()

const dialogImageUrl = ref('')
const dialogVisible = ref(false)

const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
	//console.log(uploadFile, uploadFiles)
}

const logout = () => {
	store.userStore.logoutAction().then(() => {
		// router.push({ path: '/login' })

		// 刷新页面
		location.reload()
	})
}

const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
	dialogImageUrl.value = uploadFile.url!
	dialogVisible.value = true
}

const router = useRouter()
const route = useRoute()
const dataFormRef: any = ref(null)
const dataForm = reactive({
	avatar: store.userStore.user.avatar,
	password: '',
	newPassword: '',
	confirmPassword: ''
})

const dataRules = ref({
	password: [{ required: true, message: '必填', trigger: 'blur' }],
	newPassword: [{ required: true, validator: validatePassword, trigger: 'blur' }],
	confirmPassword: [{ required: true, message: '必填', trigger: 'blur' }]
})

const handleSuccess: UploadProps['onSuccess'] = (res, file) => {
	if (res.code !== 0) {
		ElMessage.error('图片上传失败：' + res.msg)
		return false
	}
	dataForm.avatar = res.data.filePath
}

const handleDataForm = () => {
	dataFormRef.value.validate((valid: boolean) => {
		if (!valid) {
			return false
		}
		if (dataForm.newPassword !== dataForm.confirmPassword) {
			ElMessage.error('确认密码必须与新密码一致')
			return
		}
		// 修改密码
		updatePasswordApi(dataForm).then(() => {
			ElMessage.success('修改成功')
			// 关闭当前tab
			logout()
		})
	})
}
</script>
<style scoped>
.el-form-item-msg {
	font-size: 12px;
	color: #999;
	clear: both;
	width: 100%
}
</style>