<template>
	<el-row :gutter="10">
		<el-col :span="14">
			<el-card shadow="hover" header="欢迎">
				<div class="welcome">
					<div class="logo">
						<img src="@/assets/logo.png" />
						<h2>欢迎体验(GO语言开发的开源免费可商用系统)</h2>
					</div>
					<div class="tips">
						<div class="tips-item">
							<div class="tips-item-icon"><svg-icon icon="icon-database-fill"></svg-icon></div>
							<div class="tips-item-message">这里是项目控制台，你可以点击右上方的“...”按钮来定义你喜欢的界面风格。</div>
						</div>
						<div class="tips-item">
							<div class="tips-item-icon"><svg-icon icon="icon-flag"></svg-icon></div>
							<div class="tips-item-message">在时间允许下，会持续更新精磨。</div>
						</div>
						<div class="tips-item">
							<div class="tips-item-icon"><svg-icon icon="icon-send"></svg-icon></div>
							<div class="tips-item-message">项目目的：让go程序员开发项目应用更容易。</div>
						</div>
					</div>
					<div class="actions">
						<el-button type="primary" size="large" @click="godoc"><svg-icon
								icon="icon-file-text"></svg-icon>&nbsp;开发文档</el-button>
					</div>
				</div>
			</el-card>
		</el-col>
		<el-col :span="10">
			<el-card shadow="hover" header="时钟" class="item-background_time">
				<div class="time">
					<h2>{{ time }}</h2>
					<p>{{ day }}</p>
				</div>
			</el-card>
			<br>

			<el-card shadow="hover" header="关于项目" class="item-background">
				<p>gincms是使用go语言gin框架开发一套快速开发系统</p>
				<P>前端后台管理系统是基于maku-admin</P>
				<p>包含:</p>
				<p>go语言后台管理系统(基础rbac)</p>
				<p>go语言cms系统(开发中)</p>
				<p>go语言博客系统(开发中)</p>
				<p>使用门槛极低，且采用MIT开源协议，完全免费开源，可免费用于商业项目等场景</p>
				<br>
				<el-button type="danger" size="small" @click="gostar"><svg-icon icon="icon-heart-fill"></svg-icon>&nbsp;去帮作者star</el-button>
			</el-card>
		</el-col>
	</el-row>
</template>
<script lang="ts" setup>
import xeUtils from 'xe-utils'
import { ref, onMounted } from 'vue'
const time = ref('')
const day = ref('')
const godoc = () => {
	window.open('http://www.gincms.com')
}
const gostar = () => {
	window.open('https://github.com/weiyunxiao/gincms')
}

const showTime = () => {
	time.value = xeUtils.toDateString(new Date(), 'HH:mm:ss')
	day.value = xeUtils.toDateString(new Date(), 'yyyy-dd-MM')
}
onMounted(() => {
	showTime()
	setInterval(() => {
		showTime()
	}, 1000)
})
</script>
<style scoped>
/* .welcome {} */

.welcome .logo {
	text-align: center;
}

.welcome .logo img {
	vertical-align: bottom;
	width: 300px;
	height: 100px;
	margin-bottom: 20px;
}

.welcome .logo h2 {
	font-size: 30px;
	font-weight: normal;
	display: flex;
	align-items: center;
	justify-content: center;
}

.tips {
	margin-top: 20px;
	padding: 0 40px;
}

.tips-item {
	display: flex;
	align-items: center;
	justify-content: center;
	padding: 7.5px 0;
}

.tips-item-icon {
	width: 40px;
	height: 40px;
	display: flex;
	align-items: center;
	justify-content: center;
	border-radius: 50%;
	font-size: 18px;
	margin-right: 20px;
	color: var(--el-color-primary);
	background: rgba(180, 180, 180, 0.1);
}

.tips-item-message {
	flex: 1;
	font-size: 14px;
}

.actions {
	text-align: center;
	margin: 40px 0 20px 0;
}

.item-background p {
	color: #999;
	margin-top: 10px;
	line-height: 1.8;
}

.item-background_time {
	background: linear-gradient(to right, #8E54E9, #4776E6);
	color: #fff;
}

.time h2 {
	font-size: 40px;
}

.time p {
	font-size: 14px;
	margin-top: 13px;
	opacity: 0.7;
}
</style>
