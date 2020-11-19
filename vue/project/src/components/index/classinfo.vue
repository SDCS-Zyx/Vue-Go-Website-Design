<template>
	<div id="area">
		<div id="first">
			<div id="content">
				<el-table id="table1" :data="classInfo" style="width: 100%" max-height="400">
					<el-table-column prop="name" label="课程名称" width="150"></el-table-column>
					<el-table-column prop="school" label="所属学校" width="150"></el-table-column>
					<el-table-column prop="dep" label="所属院系" width="150"></el-table-column>
					<el-table-column prop="time" label="课程时间" width="150"></el-table-column>
					<el-table-column prop="teacher" label="课程教师" width="150"></el-table-column>
					<el-table-column prop="tid" label="教师编号" width="150"></el-table-column>
				</el-table>

				<el-table id="table2" :data="typelist" :row-style="rowStyle" :cell-style="cellStyle" style="width: 100%" max-height="400" border>
					<el-table-column prop="day" label="周" width="100"></el-table-column>
					<el-table-column prop="first" label="第一节" width="100"></el-table-column>
					<el-table-column prop="second" label="第二节" width="100"></el-table-column>
					<el-table-column prop="third" label="第三节" width="100"></el-table-column>
					<el-table-column prop="forth" label="第四节" width="100"></el-table-column>
					<el-table-column prop="fifth" label="第五节" width="100"></el-table-column>
					<el-table-column prop="sixth" label="第六节" width="100"></el-table-column>
					<el-table-column prop="seventh" label="第七节" width="100"></el-table-column>
					<el-table-column prop="eighth" label="第八节" width="100"></el-table-column>
				</el-table>
			</div>
			<div id="button">
				<el-button id="b1" @click="openClass">open class</el-button>
				<el-button id="b2" @click="changeType">class table</el-button>
				<el-button id="b3" @click="back">back</el-button>
			</div>
		</div>

		<div id="upload">
			<el-button id="close" @click="closeUpload">X</el-button>
			<div id="upload_area">
				<el-upload class="upload" drag action="http://localhost:8081/upload">
					<i class="el-icon-upload"></i>
					<div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
				</el-upload>

				<div class="tip">
					<el-row>只能上传txt文件,内容格式如下：</el-row>
					<el-row>1.文件名为课程名_教师id_课程时间</el-row>
					<el-row>(例：教师a_2020_1,3.txt)</el-row>
					<el-row>2.txt中每一行为学生id</el-row>
				</div>
			</div>

		</div>
	</div>
</template>

<script>

	import func from "../../js/function.js"

	export default {
		data() {
			return {
				id: this.$route.params.id,
				type: "student",
				classInfo: [],
				typelist:[{
					day:"Monday",
					first:"-",
					second:"-",
					third:"-",
					forth:"-",
					fifth:"-",
					sixth:"-",
					seventh:"-",
					eighth:"-"
				},{
					day:"Tuesday",
					first:"-",
					second:"-",
					third:"-",
					forth:"-",
					fifth:"-",
					sixth:"-",
					seventh:"-",
					eighth:"-"
				},{
					day:"Wednesday",
					first:"-",
					second:"-",
					third:"-",
					forth:"-",
					fifth:"-",
					sixth:"-",
					seventh:"-",
					eighth:"-"
				},{
					day:"Thursday",
					first:"-",
					second:"-",
					third:"-",
					forth:"-",
					fifth:"-",
					sixth:"-",
					seventh:"-",
					eighth:"-"
				},{
					day:"Friday",
					first:"-",
					second:"-",
					third:"-",
					forth:"-",
					fifth:"-",
					sixth:"-",
					seventh:"-",
					eighth:"-"
				},{
					day:"Saturday",
					first:"-",
					second:"-",
					third:"-",
					forth:"-",
					fifth:"-",
					sixth:"-",
					seventh:"-",
					eighth:"-"
				},{
					day:"Sunday",
					first:"-",
					second:"-",
					third:"-",
					forth:"-",
					fifth:"-",
					sixth:"-",
					seventh:"-",
					eighth:"-"
				}]
			}
		},

		methods: {
			setting: function() {
				this.getClassInfo()
			},

			getClassInfo: function() {
				var _self = this
				this.axios.post("http://localhost:8081"+'/Home/classInfo/getClassInfo', {
					id: _self.id,
				})
				.then((response)=>{ 
					this.message = response.data
				// console.log(this.message)
				if (this.message.slice(0, 7) == "success") {
					//success,
					this.resolveClassInfo(this.message.slice(8))
					this.setTypelist()
				} else {
					// this.showToast()
				}
			})
				.catch((error)=> {
					console.log(error)
				})
			},

			resolveClassInfo: function(info) {
				var infoList = info.split(";")
				this.type = infoList.pop()
				if (this.type == "teacher") {
					this.showOpenClass()
				}

				for (var i in infoList) {
					var l = infoList[i].split(",")
					var dic = {name:l[0], school:l[1]=="sysu" ? "中山大学" : "", dep:l[2], time:func.Time2Week(l[3],l[4]), teacher:l[5], tid:l[6]}
					this.classInfo.push(dic)
				}
			},

			changeType: function() {
				var e1 = document.getElementById("table1")
				var e2 = document.getElementById("table2")
				e1.setAttribute("style","display:none")
				e2.setAttribute("style", "display:block")
			},

			back: function() {
				var e1 = document.getElementById("table1")
				var e2 = document.getElementById("table2")
				e1.setAttribute("style","display:block")
				e2.setAttribute("style", "display:none")
			},			

			setTypelist: function() {
				if (this.classInfo.length == 0) {
					return
				}

				var l = this.classInfo
				for (var i in l) {
					var n = l[i]["name"]
					var t = func.Week2Time(l[i]["time"])
					var temp = this.typelist[t[0]][t[1]]
					temp = temp == "-" ? n : temp + "," + n
					this.typelist[t[0]][t[1]] = temp
				}
			},

			showOpenClass: function() {
				var e = document.getElementById("b1")
				b1.setAttribute("style", "display:inline-block")
			},

			openClass: function() {
				var e = document.getElementById("upload")
				e.setAttribute("style", "display:block")
			},

			closeUpload: function() {
				var e = document.getElementById("upload")
				e.setAttribute("style", "display:none")
			},

		},
		computed: {
			rowStyle: function(row, index) {
				return {"height" : "80px"}
			},

			cellStyle: function(row, col, rowIndex, colIndex) {
				return {
					"text-align" : "center",
					"white-space" : "pre-wrap"
				}
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
			this.setting()
			var _self = this
			window.addEventListener("beforeunload",()=>{
				sessionStorage.setItem('data',JSON.stringify({id : _self.id, index : "classinfo"}))
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

#first {
	position: absolute;
}

#content{
	width: 920px;
	height: 400px;
	border: 1px solid #000000;
	margin-left: 20px;
}

#table2 {
	display: none;
	border-bottom: 1px solid #000000;
}

#button {
	margin-top: 20px;
	float: right;
}

.el-table {
	filter:alpha(opacity=50);  
	-moz-opacity:0.5;  
	-khtml-opacity: 0.5;  
	opacity: 0.7;
}

#b1, #b2, #b3 {
	margin-right: 20px;
}

#upload {
	display: none;
	border: 1px solid;
	border-radius: 5px;
	position: absolute;
	width: 360px;
	height: 300px;
	margin-left: 250px;
	margin-top: 80px;
	z-index: 0;
}

#upload_area {
	position: absolute;
	z-index: 1;
}

#close {
	position: absolute;
	margin-left: 130px;
	border: 1px solid;
	z-index: 2;
}

.tip {
	text-align: left;
	font-weight: bold;
}

#b1 {
	display: none;
}

</style>