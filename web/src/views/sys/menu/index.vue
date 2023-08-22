<template>
	<el-card class="mod__menu">
		<el-form :inline="true">
			<el-form-item>
				<el-button v-auth="'post:/admin_api/sys/menu|'" type="primary" @click="addOrUpdateHandle()">新增</el-button>
			</el-form-item>
			<el-form-item>
				<el-button type="danger" @click="toggleExpandAll()">
					<el-icon style="width: 100%"><Sort />展开/折叠</el-icon>
				</el-button>
			</el-form-item>
		</el-form>
		<el-table
			v-if="refreshTable"
			v-loading="state.dataListLoading"
			:default-expand-all="isExpandAll"
			:data="state.dataList"
			row-key="id"
			border
			style="width: 100%"
		>
			<el-table-column prop="name" label="名称" header-align="center" width="180"></el-table-column>
			<el-table-column prop="icon" label="图标" header-align="center" align="center" width="80">
				<template #default="scope">
					<svg-icon :icon="scope.row.icon"></svg-icon>
				</template>
			</el-table-column>
			<el-table-column prop="type" label="类型" header-align="center" align="center" width="100">
				<template #default="scope">
					<el-tag v-if="scope.row.type === 0" type="info">菜单</el-tag>
					<el-tag v-if="scope.row.type === 1" type="success">按钮</el-tag>
					<el-tag v-if="scope.row.type === 2" type="warning">接口</el-tag>
				</template>
			</el-table-column>
			<el-table-column prop="openStyle" label="打开方式" header-align="center" align="center" width="100">
				<template #default="scope">
					<span v-if="scope.row.type !== 0"></span>
					<el-tag v-else-if="scope.row.openStyle === 0">内部打开</el-tag>
					<el-tag v-else type="info">外部打开</el-tag>
				</template>
			</el-table-column>
			<el-table-column prop="sort" label="排序" header-align="center" align="center" width="70"></el-table-column>
			<el-table-column prop="url" label="路由" header-align="center" align="center" width="150" :show-overflow-tooltip="true"></el-table-column>
			<el-table-column prop="authority" label="授权标识" header-align="center" align="left" :show-overflow-tooltip="true"></el-table-column>
			<el-table-column label="操作" fixed="right" header-align="center" align="center" width="150">
				<template #default="scope">
					<el-button v-auth="'put:/admin_api/sys/menu|'" type="primary" link @click="addOrUpdateHandle(scope.row.id)">修改</el-button>
					<el-button v-auth="'delete:/admin_api/sys/menu|'" type="primary" link @click="deleteHandle(scope.row.id)">删除</el-button>
				</template>
			</el-table-column>
		</el-table>
		<add-or-update ref="addOrUpdateRef" @refresh-data-list="getDataList"></add-or-update>
	</el-card>
</template>

<script setup lang="ts">
import { useCrud } from '@/hooks'
import { reactive, ref, nextTick } from 'vue'
import AddOrUpdate from './add-or-update.vue'
import { IHooksOptions } from '@/hooks/interface'
import { Sort } from '@element-plus/icons-vue'

const state: IHooksOptions = reactive({
	dataListUrl: 'admin_api/sys/menu_list',
	deleteUrl: 'admin_api/sys/menu',
	isPage: false
})

const addOrUpdateRef = ref()
const addOrUpdateHandle = (id?: number) => {
	addOrUpdateRef.value.init(id)
}

const { getDataList, deleteHandle } = useCrud(state)

// 是否展开，默认全部折叠
const isExpandAll = ref(false)
// 是否重新渲染表格状态
const refreshTable = ref(true)

/**
 * 切换 展开/折叠
 */
const toggleExpandAll = () => {
	refreshTable.value = false
	isExpandAll.value = !isExpandAll.value
	nextTick(() => {
		refreshTable.value = true
	})
}
</script>
