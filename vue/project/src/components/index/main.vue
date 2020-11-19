<template>
	<div id="area">
		<div id="content">
			<el-row id="first">欢迎使用课程信息管理系统!</el-row>
			<el-row id="second">本系统具有以下几种功能</el-row>
			<el-row class="feature">1.个人信息，查看或编辑你的个人信息</el-row>
			<el-row class="feature">2.课程信息，查看当前已选择的课程信息，还有课程表形式展示课程方便查看</el-row>
			<el-row class="feature">3.签到管理，查询当前已选择课程相关的签到信息，学生可以此页面查看或完成签到，教师可以选择开启或关闭签到</el-row>
		</div>
	</div>
</template>

<script>
	export default {
		data() {
			return {
				id: this.$route.params.id
			}
		},
		beforeMount() {
			var s = sessionStorage.getItem('data')
			if (s != null && s != "{}") {
				var o = JSON.parse(s)
				this.id = o.id
			}

		},
		mounted() {
			var _self = this
			window.addEventListener("beforeunload",()=>{
				sessionStorage.setItem('data',JSON.stringify({id : _self.id, index : "main"}))
			})
		}
	}
</script>

<style scoped>
	#area {
	margin-top: 0px;
	margin-left: 0px;
	position: absolute;
	width: 100%;
	height: 100%;
	text-align: center;
}

#content{
	width: 900px;
	height: wrap-content;
	padding-top: 20px;
	padding-left: 20px;
	padding-right: 20px;
}

#first {
	margin-top: 40px;
	font-size: 60px;
}

#second {
	margin-top: 80px;
	text-align: left;
	font-size: 25px;
}

.feature {
	margin-top: 20px;
	text-align: left;
	font-size: 25px;
}
</style>