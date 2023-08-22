<template>
	<el-card>
		<div class="sc-file-select">
			<div class="sc-file-select__side" v-loading="menuLoading">
				<div class="sc-file-select__side-menu">
					<el-scrollbar ref="scrollbar">
						<el-tree ref="group" class="menu" :data="menu" :node-key="'id'" :props="treeProps"
							:current-node-key="menu.length > 0 ? menu[0]['id'] : ''" highlight-current
							@node-click="groupClick">
							<template #default="{ node }">
								<svg-icon icon="icon-folder"></svg-icon>
								&nbsp;
								<span class="el-tree-node__label">
									{{ node.label }}
								</span>
							</template>
						</el-tree>
					</el-scrollbar>
				</div>
			</div>
			<div class="sc-file-select__files" v-loading="listLoading">
				<div class="sc-file-select__list">
					<el-scrollbar ref="scrollbar">
						<el-empty v-if="data.length == 0" description="无数据" :image-size="80"></el-empty>
						<div v-for="item in data" :key="item['id']" class="sc-file-select__item"
							:class="{ active: value.includes(item['url']) }" @click="select(item)">
							<div class="sc-file-select__item__file">
								<div class="sc-file-select__item__select"></div>
								<div class="sc-file-select__item__box"></div>
								<el-image v-if="_isImg(item['url'])" :src="item['url']" fit="contain" lazy></el-image>
								<div v-else class="item-file item-file-doc" style="background-color: #a6cff9; color: #fff">
									<svg-icon icon="icon-file-unknown"></svg-icon>
									{{ _getExt(item['url']) }}文件
								</div>
							</div>
							<p :title="item['fileName']">{{ item['fileName'] }}</p>
						</div>
					</el-scrollbar>
				</div>
			</div>
		</div>
	</el-card>
</template>
<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import service from '@/utils/request'

let menuLoading = ref(false)
let listLoading = ref(false)
let treeProps = reactive({
	key: 'id',
	label: 'label',
	children: 'children'
})
let fileProps = reactive({
	key: 'id',
	fileName: 'fileName',
	url: 'url'
})
let menuId = ref('')
let value = ref('')
const menu = ref([])
const data = ref([])
let getMenu = () => {
	menuLoading.value = true
	service.get('admin_api/sys/fileManage_dir_list').then((res) => {
		menu.value = res.data
		menuId.value = res.data[0].id
		getData()
	})
	menuLoading.value = false
}
let getData = () => {
	let dirName = menuId.value
	listLoading.value = true
	service.get('admin_api/sys/fileManage_dirFile_list?dirName=' + dirName).then((res) => {
		data.value = res.data
	})
	listLoading.value = false
}
onMounted(() => {
	getMenu()
})

let groupClick = (data: any) => {
	menuId.value = data.id
	getData()
}


let select = (item: any) => {
	const itemUrl = item['url']
	if (value.value.includes(itemUrl)) {
		value.value = ''
	} else {
		value.value = itemUrl
	}
	console.log(value.value);

}
let _isImg = (fileUrl: String) => {
	const imgExt = ['.jpg', '.jpeg', '.png', '.gif', '.bmp']
	const fileExt = fileUrl.substring(fileUrl.lastIndexOf("."))
	return imgExt.indexOf(fileExt) != -1
}
let _getExt = (fileUrl: String) => {
	return fileUrl.substring(fileUrl.lastIndexOf(".") + 1)
}

</script>

<style scoped>
.sc-file-select {
	display: flex;
}

.sc-file-select__files {
	flex: 1;
}

.sc-file-select__list {
	height: 700px;
}

.sc-file-select__item {
	display: inline-block;
	float: left;
	margin: 0 15px 25px 0;
	width: 110px;
	cursor: pointer;
}

.sc-file-select__item__file {
	width: 110px;
	height: 110px;
	position: relative;
}

.sc-file-select__item__file .el-image {
	width: 110px;
	height: 110px;
}

.sc-file-select__item__box {
	position: absolute;
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	border: 2px solid var(--el-color-success);
	z-index: 1;
	display: none;
}

.sc-file-select__item__box::before {
	content: '';
	position: absolute;
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	background: var(--el-color-success);
	opacity: 0.2;
	display: none;
}

.sc-file-select__item:hover .sc-file-select__item__box {
	display: block;
}

.sc-file-select__item.active .sc-file-select__item__box {
	display: block;
}

.sc-file-select__item.active .sc-file-select__item__box::before {
	display: block;
}

.sc-file-select__item p {
	margin-top: 10px;
	white-space: nowrap;
	text-overflow: ellipsis;
	overflow: hidden;
	-webkit-text-overflow: ellipsis;
	text-align: center;
}

.sc-file-select__item__checkbox {
	position: absolute;
	width: 20px;
	height: 20px;
	top: 7px;
	right: 7px;
	z-index: 2;
	background: rgba(0, 0, 0, 0.2);
	border: 1px solid #fff;
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
}

.sc-file-select__item__checkbox i {
	font-size: 14px;
	color: #fff;
	font-weight: bold;
	display: none;
}

.sc-file-select__item__select {
	position: absolute;
	width: 20px;
	height: 20px;
	top: 0px;
	right: 0px;
	z-index: 2;
	background: var(--el-color-success);
	display: none;
	flex-direction: column;
	align-items: center;
	justify-content: center;
}

.sc-file-select__item__select i {
	font-size: 14px;
	color: #fff;
	font-weight: bold;
}

.sc-file-select__item.active .sc-file-select__item__checkbox {
	background: var(--el-color-success);
}

.sc-file-select__item.active .sc-file-select__item__checkbox i {
	display: block;
}

.sc-file-select__item.active .sc-file-select__item__select {
	display: flex;
}

.sc-file-select__item__file .item-file {
	width: 110px;
	height: 110px;
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
}

.sc-file-select__item__file .item-file i {
	font-size: 40px;
}

.sc-file-select__item__file .item-file.item-file-doc {
	color: #409eff;
}

.sc-file-select__item__upload {
	position: absolute;
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	z-index: 1;
	background: rgba(255, 255, 255, 0.7);
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
}

.sc-file-select__side {
	width: 200px;
	margin-right: 15px;
	border-right: 1px solid rgba(128, 128, 128, 0.2);
	display: flex;
	flex-flow: column;
}

.sc-file-select__side-menu {
	flex: 1;
	max-height: 700px;
}

.sc-file-select__side-msg {
	height: 32px;
	line-height: 32px;
}

.sc-file-select__top {
	margin-bottom: 15px;
	display: flex;
	justify-content: space-between;
}

.sc-file-select__upload {
	display: inline-block;
}

.sc-file-select__top .tips {
	font-size: 12px;
	margin-left: 10px;
	color: #999;
}

.sc-file-select__top .tips i {
	font-size: 14px;
	margin-right: 5px;
	position: relative;
	bottom: -0.125em;
}

.sc-file-select__pagination {
	margin: 15px 0;
}

.sc-file-select__do {
	text-align: right;
}
</style>
