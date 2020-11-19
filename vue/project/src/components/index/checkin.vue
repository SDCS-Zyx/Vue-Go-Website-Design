<template>
	<div id="area">
		<div id="content">
			<div id="select_row">
				<el-row id="select">
					<el-col :span="6">
						<el-select v-model="SearchBlock.cs" placeholder="请选择课程">
							<el-option v-for="item in SearchBlock.classList" :key="item"
							:label="item" :value="item"></el-option>
						</el-select>
					</el-col>
					<el-col :span="6">
						<el-button @click="search">查询</el-button>
					</el-col>
				</el-row>
			</div>
			<div id="table">
				<el-table id="table1" :data="SearchBlock.checkList" style="width: 100%" height="400" max-height="400">
					<el-table-column align="center" prop="name" label="课程名称" width="140"></el-table-column>
					<el-table-column align="center" prop="ctime" label="签到时间" width="140"></el-table-column>
					<el-table-column align="center" prop="type" label="签到状况" width="135"></el-table-column>
					<el-table-column align="center" v-if="showDownload" label="操作" width="80">
						<template slot-scope="scope">
							<el-button @click="showDetail(scope.row)" type="text" size="small">查看详情</el-button>
							<el-button @click="download(scope.row)" type="text" size="small">下载</el-button>
						</template>
					</el-table-column>
				</el-table>
			</div>
		</div>
		<div id="check">
			<div id="have">
				<el-row>
					<el-col :span="12" :push="6">
						<p>You have a checkin mask</p>
					</el-col>
				</el-row>
				<el-row>
					<el-col :span="12" :push="6">
						<p>当前课程：{{CheckBlock.nc}}</p>
					</el-col>
				</el-row>
				<el-row id="need_checkin">
					<el-col :span="12" :push="3">
						<el-input id="input" v-model="CheckBlock.incode" placeholder="请输入签到码"></el-input>
					</el-col>
					<el-col :span="6" :push="3">
						<el-button @click="studentCheckin">confirm</el-button>
					</el-col>
				</el-row>
				<el-row id="had_checkin">
					<el-col :span="12" :push="6"><p>Checkin Success!</p></el-col>
				</el-row>
			</div>
			<div id="noc">
				<p>You don't have checkin mask now.</p>
			</div>
			<div id="openCheckin">
				<el-row>
					<el-col :span="12" :push="6">
						<p>当前课程：{{CheckBlock.nc}}</p>
					</el-col>
				</el-row>
				<el-row>
					<el-col :span="16" :push="4">
						<p>当前课程{{opened ? "已" : "未"}}开启签到</p>
						<p id="show_ddl">
							已签到人数：{{CheckBlock.count}}<br>截止时间：{{resolveDdl(CheckBlock.ddl)}}
						</p>
					</el-col>
				</el-row>				
				<el-row>
					<el-col :span="8" :offset="2">
						<p>签到码：</p>
					</el-col>
					<el-col :span="12">
						<el-input v-model="CheckBlock.code" disabled></el-input>
					</el-col>
				</el-row>
				<el-row>
					<el-col :span="8" :offset="2">
						<p>设置签到时限：</p>
					</el-col>
					<el-col :span="12">
						<el-input v-model="CheckBlock.timer" :disabled="opened"></el-input>
					</el-col>
				</el-row>
				<el-row>
					<el-col :span="6" :offset="5">
						<el-button @click="openCheckin">开启签到</el-button>
					</el-col>
					<el-col :span="6" :offset="2">
						<el-button @click="cancelCheckin">取消签到</el-button>
					</el-col>
				</el-row>
			</div>
		</div>
		<div id="detail">
			<div id="detailHead">
				<el-button id="closedl" @click="closeDetail">关闭</el-button>
				<p id="detailP">课程签到详情</p>
			</div>
			<el-table id="table2" :data="DetailBlock.infoList" style="width: 100%" height="400" max-height="400">
				<el-table-column align="center" prop="name" label="学生名称" width="140"></el-table-column>
				<el-table-column align="center" prop="id" label="学生id" width="140"></el-table-column>
				<el-table-column align="center" prop="ctime" label="签到时间" width="140"></el-table-column>
				<el-table-column align="center" prop="type" label="签到状况" width="135"></el-table-column>			
				<el-table-column align="center" prop="pic" label="个人/签到图片" width="200">
					<template   slot-scope="scope">            
						<img v-for="item in scope.row.pic" :src="item" width="80" height="80"/>
					</template> 
				</el-table-column>
			</el-table>
		</div>
	</div>
</template>

<script>

	import func from "../../js/function.js"

	export default {
		data() {
			return {
				Info: {
					id: this.$route.params.id,
					type: "student"
				},
				CheckBlock: {
					nc: "none",
					incode: "",
					code: "",
					ddl: "",
					ctype: "", //学生是否签到
					timer: "10",
					count: "0",
					ctime: ""
				},
				SearchBlock: {
					cs: "",
					classList: [],
					checkList: []
				},
				DetailBlock: {
					infoList: []
				},
				opened: false,	//教师是否开启签到
				haveCheckin: false, //是否有签到
				message: "",
				ws: null,
				setTimer: null,
				showDownload:false
			}
		},
		methods: {
			setting: function() {
				this.getConfig()
			},

			getConfig: function() {
				var _self = this
				var t = func.GetTime("")
				this.axios.post("http://localhost:8081"+'/Home/checkin/getConfig', {
					id: _self.Info.id,
					t: t
				})
				.then((response)=>{ 
					this.message = response.data
					// console.log(this.message)
					if (this.message.slice(0, 7) == "success") {
						//success,
						this.resolveConfig(this.message.slice(8))
					} else {
						// this.showToast()
					}
				})
				.catch((error)=> {
					console.log(error)
				})
			},

			resolveConfig: function(s) {
				var obj = func.ResolveJSON(s)
				// console.log(obj)

				this.Info.type = obj.type

				if (this.Info.type == "student") {
					this.SearchBlock.classList = obj.classList.split(",")
					this.SearchBlock.classList.forEach(function(v, i, a) {
						a[i] = v.slice(1,-1)
					})
				} else if (this.Info.type == "teacher") {
					var classTime = obj.classList.split(";")
					for (var i in classTime) {
						var sp = classTime[i].split(",")
						this.SearchBlock.classList.push(sp[0])
						this.checkClassTime(sp)
					}
				}

				if (obj.checkRes != "") {
					this.haveCheckin = true

					var r = obj.checkRes.split(",")
					this.CheckBlock.ddl = r[0]
					this.CheckBlock.nc = r[1]
					this.CheckBlock.ctype = r[2]

					if (this.Info.type == "student" && this.CheckBlock.ctype != "0") {
						this.checkinComplete()
					} else if (this.Info.type == "teacher") {
						this.showCheckinInfo()
						this.setWs()
					}
				}

				this.showCheckin()
			},

			checkClassTime: function(a) {
				if (func.CheckClassTime(a[1], a[2])) {
					this.CheckBlock.nc = a[0]
					this.CheckBlock.ctime = a[1] + "," + a[2]
				}
			},

			search: function() {
				var _self = this
				this.axios.post("http://localhost:8081"+'/Home/checkin/getCheckList', {
					id: _self.Info.id,
					class: _self.SearchBlock.cs,
					type: _self.Info.type
				})
				.then((response)=>{ 
					this.message = response.data
					// console.log(this.message)
					if (this.message.slice(0, 7) == "success") {
						//success,
						this.resolveCheckList(this.message.slice(8))
					} else {
						// this.showToast()
					}
				})
				.catch((error)=> {
					console.log(error)
				})
			},

			resolveCheckList: function(s) {
				this.SearchBlock.checkList = []
				if (s.length == 0) {
					return
				}

				var l = s.split(";")
				if (this.Info.type == "student") {
					for (var i in l) {
						var info = l[i].split(",")
						var dic = {name:this.SearchBlock.cs, ctime:this.resolveDdl(info[1]), type:this.resolveCheckinType(info[0], "")}
						this.SearchBlock.checkList.push(dic)
					} 
				} else if (this.Info.type == "teacher") {
					if (l[0] == "") { return }
						
						var num = l.pop()
					this.showDownload = true
					for (var i in l) {
						var info = l[i].split(",")
						var dic = {name:this.SearchBlock.cs, ctime:this.resolveDdl(info[1])+"-"+this.resolveDdl(info[0]), type:this.resolveCheckinType(info[2], num)}
						this.SearchBlock.checkList.push(dic)
					}
				}

			},

			showCheckin: function() {
				var e1 = document.getElementById("have")
				var e2 = document.getElementById("noc")

				if (this.Info.type == "student") {
					var e3 = document.getElementById("openCheckin")
					e3.setAttribute("style", "display:none")

					if (this.haveCheckin) {
						e1.setAttribute("style", "display:block")
						e2.setAttribute("style", "display:none")
					} else {
						e1.setAttribute("style", "display:none")
						e2.setAttribute("style", "display:block")
					}
				} else if (this.Info.type == "teacher") {
					e1.setAttribute("style", "display:none")
					e2.setAttribute("style", "display:none")
				}

			},

			showCheckinInfo: function() {
				this.opened = true
				this.CheckBlock.code = func.RandomCode(this.CheckBlock.nc, this.CheckBlock.ddl)

				var e1 = document.getElementById("show_ddl")
				e1.setAttribute("style", "display:block")
			},

			closeCheckinInfo: function() {
				var e1 = document.getElementById("show_ddl")
				e1.setAttribute("style", "display:none")
			},

			resolveDdl: function(s) {
				if (s == "" || !s) return "/"
					return s.slice(0,4) + "/" + s.slice(4,6) + "/" + s.slice(6,8)
				+ "  " + s.slice(8,10) + ":" + s.slice(10)
			},

			ddlToString: function(t) {
				return t.replace(/[\/\ \:]/g, "")
			},

			resolveCheckinType: function(s, c) {
				if (this.Info.type == "student") {
					return s == "1" ? "已签到" : "未签到"
				} else if (this.Info.type == "teacher") {
					return "已签到人数： " + (Number(c) - Number(s) + 1).toString() + ",\n未签到人数：" + (Number(s) - 1).toString()
				}
			},

			studentCheckin: function() {
				if (this.CheckBlock.ctype != "1" && func.Before(func.GetTime(""), this.CheckBlock.ddl)) {
					//开始签到，验证code
					var inc = this.CheckBlock.incode
					// var inc = func.RandomCode(this.CheckBlock.nc, this.CheckBlock.ddl)

					if (inc == "" || inc != func.RandomCode(this.CheckBlock.nc, this.CheckBlock.ddl)) {
						return
					}

					var _self = this
					var stime = func.GetTime("")
					this.axios.post("http://localhost:8081"+'/Home/checkin/checkin', {
						id: _self.Info.id,
						cname: _self.CheckBlock.nc,
						ddl: _self.CheckBlock.ddl,
						stime: stime
					})
					.then((response)=>{ 
						this.message = response.data
						// console.log(this.message)
						if (this.message.slice(0, 7) == "success") {
							//success
							this.checkinComplete()
						} else {
							// this.showToast()
						}
					})
					.catch((error)=> {
						console.log(error)
					})
				}
			},

			checkinComplete: function() {
				var e1 = document.getElementById("need_checkin")
				var e2 = document.getElementById("had_checkin")
				e1.setAttribute("style", "display:none")
				e2.setAttribute("style", "display:block")
			},

			openCheckin: function() {
				if (this.CheckBlock.nc == "none" || this.opened || !func.CheckNum(this.CheckBlock.timer)) {
					return
				}

				this.opened = true
				this.haveCheckin = true
				this.CheckBlock.ddl = func.GetTime(this.CheckBlock.timer)
				this.code = func.RandomCode(this.CheckBlock.nc, this.CheckBlock.ddl)

				this.showCheckinInfo()

				var _self = this
				var stime = func.GetTime("")
				this.axios.post("http://localhost:8081"+'/Home/checkin/openCheckin', {
					id: _self.Info.id,
					cname: _self.CheckBlock.nc,
					ddl: _self.CheckBlock.ddl,
					stime: stime
				})
				.then((response)=>{ 
					this.message = response.data
					// console.log(this.message)
					if (this.message.slice(0,7) == "success") {
							//success
							this.setWs()

							sessionStorage.setItem('timer',JSON.stringify({timer:_self.CheckBlock.timer}))
						} else {
							// this.showToast()
						}
					})
				.catch((error)=> {
					console.log(error)
				})
			},

			setWs: function() {
				this.ws = new WebSocket("ws://localhost:8081/Home/checkin/getCount")
				var _self = this
				this.ws.onopen = function(evt) {  
					// console.log("Connection open ...");  
					_self.ws.send("open," + _self.CheckBlock.nc + "," + _self.CheckBlock.ddl);  
					_self.setTimer = setInterval(_self.count, 5000)
				};  
				this.ws.onmessage = function(evt) {  
					// console.log("Received Message: " + evt.data);  
					if(evt.data == "failed") {
						_self.closeWs()
					} else {
						_self.CheckBlock.count = evt.data
					}
				};  
				this.ws.onclose = function(evt) {  
					// console.log("Connection closed.");  
				};  

			},

			count: function() {
				this.ws.send("count")
			},

			closeWs: function() {
				this.ws.send("close")
				this.ws = null

				clearInterval(this.setTimer)
				this.setTimer = null
			},

			cancelCheckin: function() {
				this.opened = false
				this.haveCheckin = false
				this.CheckBlock.code = ""

				this.closeCheckinInfo()

				var _self = this
				this.axios.post("http://localhost:8081"+'/Home/checkin/closeCheckin', {
					cname: _self.CheckBlock.nc,
					ddl: _self.CheckBlock.ddl
				})
				.then((response)=>{ 
					this.message = response.data
					// console.log(this.message)
					if (this.message.slice(0,7) == "success") {
							//success

							this.CheckBlock.ddl = ""
							this.closeWs()

							sessionStorage.setItem('timer',JSON.stringify({timer:"10"}))
						} else {
							// this.showToast()
						}
					})
				.catch((error)=> {
					console.log(error)
				})
			},

			download: function(row) {
				var cname = row.name
				var te = this.ddlToString(row.ctime.split("-")[1])
				var id = this.Info.id
				this.axios.post("http://localhost:8081"+'/download', {
					id: id,
					cname: cname,
					te: te,
					tp: "download"
				})
				.then((response)=>{ this.resoleData(response.data) })
				.catch((error)=> {
					console.log(error)
				})
			},

			formatData: function(d) {
				d = "姓名\tid\t签到状况\t签到时间\n" + d
				d = d.replace(/\,/g, "\t")
				return d			
			},

			resoleData: function(data) {
				data = this.formatData(data)
				let blob = new Blob([data], {type: 'text/plain;charset=UTF-8', endings: 'native'})
				let aTag = document.createElement('a')
				aTag.download = "test"
				aTag.href = URL.createObjectURL(blob)
				aTag.click()
				URL.revokeObjectURL(aTag.href)
			},

			closeDetail: function() {
				var e = document.getElementById("detail")
				e.setAttribute("style", "display:none")
			},

			showDetail: function(row) {
				var e = document.getElementById("detail")
				e.setAttribute("style", "display:block")

				var _self = this
				var ctime = this.ddlToString(row.ctime.split("-")[1])
				this.axios.post("http://localhost:8081"+'/download', {
					id: _self.Info.id,
					cname: row.name,
					te: ctime,
					tp: "detail"
				})
				.then((response)=>{ 
					this.resolveDetail(response.data, row.name, ctime)
				})
				.catch((error)=> {
					console.log(error)
				})
			},

			resolveDetail: function(data, cname, ddl) {
				this.DetailBlock.infoList = []
				if (data == "") return

					data = data.split(";")
				for (var i in data) {
					if (data[i] == "") break
						var info = data[i].split(",")
					var dic = {name:info[0], id:info[1], ctime:info[3], type:info[2], pic:[]}
					this.getPic(info[1], dic)
					if (info[2] == "人脸识别签到") {
						this.getFace(ddl, cname, info[1], dic)
					}
					this.DetailBlock.infoList.push(dic)
				}
			},

			getPic: function(id, dic) {
				this.axios.post("http://localhost:8081"+'/Home/profile/getPic', {
					id: id,
				})
				.then((response)=>{
					let data = response.data
					let imgUrl = "data:image/png;base64," + data
					dic.pic.push(imgUrl)
				})
				.catch((error)=> {
					console.log(error)
				})
			},

			getFace: function(ddl, cname, id, dic) {
				this.axios.post("http://localhost:8081"+'/Home/checkin/getFace', {
					id: id,
					ddl: ddl,
					cname: cname
				})
				.then((response)=>{
					let data = response.data
					let imgUrl = "data:image/png;base64," + data
					dic.pic.push(imgUrl)
				})
				.catch((error)=> {
					console.log(error)
				})
			}

		},
		beforeMount() {
			var s = sessionStorage.getItem('data')
			if (s != null && s != "{}") {
				var o = JSON.parse(s)
				this.Info.id = o.id
			}

			var t = sessionStorage.getItem('timer')
			if (t != null && t != "{}") {
				var c = JSON.parse(t)
				this.CheckBlock.timer = c["timer"]
			}
		},
		mounted() {
			this.setting()
			var _self = this
			window.addEventListener("beforeunload",()=>{
				sessionStorage.setItem('data',JSON.stringify({id : _self.Info.id, index : "checkin"}))
			})
		}
	}
</script>

<style scoped>
#area {
	margin-top: 0px;
	margin-left: 0px;
	position: absolute;
	width: wrap-content;
	height: 100%;
	text-align: center;
}

#content{
	width: wrap-content;
	height: wrap-content;
	border: 1px solid #000000;
	float: left;
	margin-left: 10px;
}

#select {
	width: 100%;
	height: wrap-content;
	margin-top: 20px;
	margin-left: 40px;
}

#table {
	margin-top: 40px;
	margin-left: 40px;
	margin-right: 40px;
	margin-bottom: 20px;

	filter:alpha(opacity=50);  
	-moz-opacity:0.5;  
	-khtml-opacity: 0.5;  
	opacity: 0.7;
}

#check {
	width: 400px;
	height: wrap-content;
	float: left;
	padding-top: 20px;
	padding-bottom: 20px;
	border: 1px solid;

	margin-left: 10px;
}

#had_checkin, #show_ddl {
	display: none;
}

#table1 .cell {
	white-space: pre-wrap;
}

#detail {
	display: none;
	position: absolute;
	border: 1px solid;
	margin-left: 100px;
}

#detailHead {
	background-color: #66CCFF;
	height: 30px;
	width: 100%;
	text-align: center;
}

#detailP {
	display: inline;
	height: 30px;
	font-size: 20px;
	font-weight: bold;
	text-align: center;
}

#closedl {
	text-align: center;
	float: left;
	width: 70px;
	height: 30px;
}

</style>