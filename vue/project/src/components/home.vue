<template>
	<div id="area">
		<div id="content">
			<div id="background"></div>
			<el-row id="main">
				<el-col :span="6" id="tac">
					<el-menu :default-active="index" class="menu" @open="handleOpen" @close="handleClose" @select="handleSelect" unique-opened>
						<el-menu-item index="main">
							<i class="el-icon-menu"></i>
							<span slot="title">主页</span>
						</el-menu-item>
						<el-menu-item index="profile">
							<i class="el-icon-menu"></i>
							<span slot="title">个人信息</span>
						</el-menu-item>
						<el-menu-item index="classinfo">
							<i class="el-icon-menu"></i>
							<span slot="title">课程信息</span>
						</el-menu-item>
						<el-menu-item index="checkin">
							<i class="el-icon-menu"></i>
							<span slot="title">签到管理</span>
						</el-menu-item>
					</el-menu>
				</el-col>

				<el-col :span="18" id="show">
					<div id="head"></div>
					<div>
						<transition name="fade" mode="out-in">
							<router-view></router-view>
						</transition>
					</div>
				</el-col>
			</el-row>
		</div>

		<div id="toast">
			<el-popover popper-class="pop" width="160" v-model="visible">
				<p style="text-align: center;">{{message}}</p>
			</el-popover>
		</div>
	</div>
</template>

<script>
	export default {
		data() {
			return {
				id: this.$route.params.id,
				message: "",
				visible: false,
				index: "main"
			}
		},
		methods : {
			setting: function() {
				var width = document.body.clientWidth
				var height = parseInt(window.getComputedStyle(document
					.getElementById("area"), null).getPropertyValue("height"))

				var el = document.getElementById("toast")
				el.setAttribute("style", "position:absolute;left:" + (width/2-80).toString() + "px;")

				var _self = this
				this.handleSelect(_self.index, "")
			},

			showToast: function() {
				this.visible = true
				var _self = this
				setTimeout(function(){
					_self.visible = false
				}, 2000)
			},
			
			handleOpen: function() {

			},

			handleClose: function() {

			},

			handleSelect: function(index, indexPath) {
				this.index = index
				this.$router.push({name:index, params: {
					id: this.id
				}})
				.catch(err => {
					console.log(err)
				})
			}
		},
		beforeMount() {
			var s = sessionStorage.getItem('data')
			if (s != null && s != "{}") {
				var o = JSON.parse(s)
				this.id = o.id
				this.index = o.index
			}

		},
		mounted() {
			this.setting()
			var _self = this
			window.addEventListener("beforeunload",()=>{
				sessionStorage.setItem('data',JSON.stringify({id : _self.id, index : _self.index}))
			})
		}
}
</script>

<style scoped>
#area {
  height:100%;
  width:100%;
  position: fixed;
  margin-top: -8px;
  margin-left: -8px;
  text-align: center;
}

#content {
	width: 100%;
	height: 100%;
}

#background {
	float: left;
	height: 100%;
	width: 100%;
		 background: url("../assets/background.jpg") center center no-repeat;
	background-size: 100% 100%;
	  filter:alpha(opacity=50);  
	  -moz-opacity:0.5;  
	  -khtml-opacity: 0.5;  
	  opacity: 0.7;
}

#main {
	width: 100%;
	height: 100%;

	position: absolute;
}

.el-col {
	border: 1px solid;
	height: 100%;
	/*height: 400px;*/
}

.el-menu-item {
	text-align: left;
}

#tac {
	background-color: white;	
}

#head {
	height: 80px;
	width: 100%;
}

</style>