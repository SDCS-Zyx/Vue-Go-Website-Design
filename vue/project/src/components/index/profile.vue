<template>
	<div id="area">
		<div id="content">
			<el-form ref="Info" :model="Info" :rules="rules" label-width="120px" label-position="left">
				<el-row type="flex">
					<el-col :span="12">
						<el-form-item label="姓名：" prop="name">
							<el-input class="inp" v-model="Info.name" :disabled="chg.name"></el-input>
						</el-form-item>
						<el-form-item label="学号/职工号：" prop="id">
							<el-input class="inp" v-model="Info.id" disabled></el-input>
						</el-form-item>
					</el-col>
					<el-col :span="12">
						<img id="img" :src="Info.pic" alt="avatar" height="80" width="80">
					</el-col>
				</el-row>
				<el-row type="flex">
					<el-col :span="12">
						<el-form-item label="类型：" prop="type">
							<el-select v-model="Info.type" placeholder="请选择" disabled>
								<el-option label="教师" value="teacher"></el-option>
								<el-option label="学生" value="student"></el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col :span="12">
						<el-form-item label="性别：" prop="sex">
							<el-radio-group v-model="Info.sex" :disabled="chg.sex">
								<el-radio label="男"></el-radio>
								<el-radio label="女"></el-radio>
							</el-radio-group>
						</el-form-item>
					</el-col>
				</el-row>
				<el-row type="flex">
					<el-col :span="12">
						<el-form-item label="学校：" prop="school">
							<el-select v-model="Info.school" placeholder="请选择" :disabled="chg.school">
								<el-option label="中山大学" value="sysu"></el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col :span="12">
						<el-form-item label="院系：" prop="dep">
							<el-select v-model="Info.dep" placeholder="请选择" :disabled="chg.dep" @change="changeDep">
								<el-option label="中国语言文学系" value="中国语言文学系"></el-option>
								<el-option label="历史学系" value="历史学系"></el-option>
								<el-option label="哲学系" value="哲学系"></el-option>
								<el-option label="社会学与人类学学院" value="社会学与人类学学院"></el-option>
								<el-option label="岭南学院" value="岭南学院"></el-option>
								<el-option label="外国语学院" value="外国语学院"></el-option>
							</el-select>
						</el-form-item>
					</el-col>
				</el-row>
				<el-row type="flex" id="optional">
					<el-col :span="12">
						<el-form-item id="date" label="入学日期：">
							<el-col :span="11">
								<el-date-picker type="date" placeholder="入学日期" v-model="Info.date1" :disabled="chg.date" style="width: 100%;"></el-date-picker>
							</el-col>
							<el-col class="line" :span="1">-</el-col>
							<el-col :span="11">
								<el-date-picker placeholder="毕业日期" v-model="Info.date2" :disabled="chg.date" style="width: 100%;"></el-date-picker>
							</el-col>
						</el-form-item>
					</el-col>
					<el-col :span="12">
						<el-form-item id="subject" label="专业：" prop="subject">
							<el-select v-model="Info.subject" placeholder="请选择" :disabled="chg.subject">
								<el-option
								v-for="item in options[Info.dep]"
								:key="item.value"
								:label="item.label"
								:value="item.value">
							</el-option>
						</el-select>
					</el-form-item>
				</el-col>
			</el-row>
		</el-form>
	</div>
	<div id="buttons"> 
		<el-button @click="modify" id="modify">modify</el-button>
		<el-button @click="confirm" id="confirm">confirm</el-button>
		<el-button @click="back" id="back">back</el-button>
	</div>
</div>
</template>

<script>
	export default {
		data() {
			return {
				message: "",
				modifying: false,
				confirming: false,
				confirmed: false,
				Info: {
					name: '',
					id: this.$route.params.id,
					type: 'student',
					school: '',
					dep: '',
					subject: '',
					date1: '',
					date2: '',
					sex: '',
					pic: require('../../assets/user.jpg')
				},
				chg: {
					name: true,
					sex: true,
					school: true,
					dep: true,
					subject: true,
					date: true
				},
				rules1: {
					name: [
					{required: true, message: "name should not be null", trigger: 'blur'},
					{min: 1, max: 10, message: "name cannot be so long", trigger: 'blur'}
					],
					id: [
					{validator: this.validateNumber, required: true, message: "id should be number", trigger: 'blur', min: 0, max: 10000}
					],
					type: [{required: true, message: "type cannot be null", trigger: 'blur'}],
					school: [{required: true, message: "school cannot be null", trigger: 'change'}],
					dep: [{required: true, message: "department cannot be null", trigger: 'change'}],
					subject: [{required: true, message: "subject cannot be null", trigger: 'change'}],
					sex: [{required: true, message: "sex cannot be null", trigger: 'change'}]
				},
				rules2: {
					name: [
					{required: true, message: "name should not be null", trigger: 'blur'},
					{min: 1, max: 10, message: "name cannot be so long", trigger: 'blur'}
					],
					id: [
					{validator: this.validateNumber, required: true, message: "id should be number", trigger: 'blur', min: 0, max: 10000}
					],
					type: [{required: true, message: "type cannot be null", trigger: 'blur'}],
					school: [{required: true, message: "school cannot be null", trigger: 'change'}],
					dep: [{required: true, message: "department cannot be null", trigger: 'change'}],
					sex: [{required: true, message: "sex cannot be null", trigger: 'change'}]
				},
				options: {
					'': [],
					'中国语言文学系': [{
						label:'古代文体学理论建构与文献整理研究',
						value:'古代文体学理论建构与文献整理研究'
					},{
						value:'古代戏曲与非物质文化遗产',
						label:'古代戏曲与非物质文化遗产'
					},{
						value:'汉语言与古文字研究',
						label:'汉语言与古文字研究'
					},{
						value:'中国现当代文学与视觉文化批评',
						label:'中国现当代文学与视觉文化批评'
					}],
					'历史学系': [{
						value:'中国古代史',
						label:'中国古代史'
					},{
						value:'中国近现代史',
						label:'中国近现代史'
					}],
					'哲学系': [{
						value:'马克思主义哲学',
						label:'马克思主义哲学'
					},{
						value:'逻辑学',
						label:'逻辑学'
					},{
						value:'中国哲学',
						label:'中国哲学'
					}],
					'社会学与人类学学院': [{
						value:'社会学',
						label:'社会学'
					},{
						value:'人类学',
						label:'人类学'
					},{
						value:'考古学',
						label:'考古学'
					},{
						value:'社会工作',
						label:'社会工作'
					}],
					'岭南学院': [{
						value:'经济学系',
						label:'经济学系'
					},{
						value:'金融学系',
						label:'金融学系'
					},{
						value:'商务管理系',
						label:'商务管理系'
					}],
					'外国语学院': [{
						value:'英语专业',
						label:'英语专业'
					},{
						value:'法语专业',
						label:'法语专业'
					},{
						value:'德语专业',
						label:'德语专业'
					},{
						value:'日语专业',
						label:'日语专业'
					}]
				}
			}
		},
		methods : {
			setting: function() {
				this.getInfo()
				this.getPic()
			},

			getPic: function() {
				var _self = this
				this.axios.post("http://localhost:8081"+'/Home/profile/getPic', {
					id: _self.Info.id,
				})
				.then((response)=>{
					let data = response.data
					let imgUrl = "data:image/png;base64," + data
					this.Info.pic = imgUrl
				})
				.catch((error)=> {
					console.log(error)
				})
			},

			getInfo: function() {
				var _self = this
				this.axios.post("http://localhost:8081"+'/Home/profile/getInfo', {
					id: _self.Info.id,
				})
				.then((response)=>{ 
					this.message = response.data
					// console.log(this.message)
					if (this.message.slice(0, 7) == "success") {
						//success,
						this.resolveInfo(this.message.slice(8))
						this.showInfo()
					} else {
						// this.showToast()
					}
				})
				.catch((error)=> {
					console.log(error)
				})
			},

			resolveInfo: function(s) {
				var list = s.split(',')
				this.Info.name = list[0]
				this.Info.type = list[2]
				this.Info.school = list[3]
				this.Info.dep = list[4]
				if (this.Info.type == "student") {
					this.Info.subject = list[5]
					this.Info.date1 = list[6]
					this.Info.date2 = list[7]
					this.Info.sex = list[8] == "male" ? "男" : "女"
				} else {
					this.Info.sex = list[5] == "male" ? "男" : "女"
				}
			},

			showInfo: function() {
				var e = document.getElementById("optional")
				if (this.Info.type == "teacher") {
					e.setAttribute("style", "display:none")
				} else {
					e.setAttribute("style", "display:")
				}
			},

			modify: function() {
				if (!this.modifying) {
					this.change()
					this.show_button()
					this.modifying = true
				}
			},

			back: function() {
				if (!this.confirmed) {
					this.getInfo()
				}
				this.change()
				this.hide_button()
				this.confirmed = false
				this.modifying = false
			},

			show_button: function() {
				var e1 = document.getElementById("confirm")
				var e2 = document.getElementById("back")
				e1.setAttribute("style", "visibility:visible;")
				e2.setAttribute("style", "visibility:visible;")
			},

			hide_button: function() {
				var e1 = document.getElementById("confirm")
				var e2 = document.getElementById("back")
				e1.setAttribute("style", "visibility:hidden;")
				e2.setAttribute("style", "visibility:hidden;")
			},

			change: function() {
				this.chg.name = !this.chg.name
				this.chg.sex = !this.chg.sex
				this.chg.school = !this.chg.school
				this.chg.dep = !this.chg.dep
				this.chg.subject = !this.chg.subject
				this.chg.date = !this.chg.date
			},

			checkData: function(formName) {
				var res = false
				this.$refs[formName].validate(function(valid) {
					res = valid
				})
				return res
			},

			confirm: function() {
				var checkRes = this.checkData('Info')
				if (!checkRes) {
					return
				}
				if (this.confirming) {
					return
				}

				this.confirming = true
				var _self = this

				this.axios.post("http://localhost:8081"+'/Home/profile/changeInfo', {
					info: _self.Info
				})
				.then((response)=>{ 
					this.message = response.data
					if (this.message.slice(0,7) == "success") {
						//success,
					} else {
						// this.showToast()
						this.getInfo()
					}
					this.confirmed = true
					this.confirming = false
					this.back()
				})
				.catch((error)=> {
					console.log(error)
				})
			},

			changeDep: function() {
				this.Info.subject = this.options[this.Info.dep].value
			}
		},
		computed: {
			rules: function() {
				if (this.Info["type"] == "teacher") {
					return this.rules2
				} else {
					return this.rules1
				}
			}
		},
		beforeMount() {
			var s = sessionStorage.getItem('data')
			if (s != null && s != "{}") {
				var o = JSON.parse(s)
				this.Info.id = o.id
			}

		},
		mounted() {
			this.setting()
			var _self = this
			window.addEventListener("beforeunload",()=>{
				sessionStorage.setItem('data',JSON.stringify({id : _self.Info.id, index : "profile"}))
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
	border: 1px solid #000000;
	padding-top: 40px;
	padding-left: 20px;
	padding-bottom: 40px;

	margin-left: 20px;
}

.el-form-item /deep/ .el-form-item__label{
	text-align: justify;
	/*color: #ffffff;*/
	font-weight: bold;
	font-size: 14px;
}

.el-radio /deep/ .el-radio__label {
	text-align: justify;
	/*color: #ffffff;*/
	font-weight: bold;
}

.el-form-item {
	width: 400px;
}

#img {
	margin:0 auto;

}

#buttons {
	width: 400px;
	height: wrap-content;
	margin-top: 40px;
	margin-left: 200px;
}

.el-button {
	margin-left: 40px;
}

#confirm, #back {
	visibility: hidden;
}

</style>