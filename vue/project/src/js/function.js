
var FuncLib = {
	CheckUsernameAndPassword : function(un, pw) {
		var message = ""
		if (un == "") {
			message = "username cannot be null"
		} else if (pw == "") {
			message = "password cannot be null"
		} else if (this.checkInput(un)) {
			message = "illegal username"
		} else if (this.checkInput(pw)) {
			message = "illegal password"
		} else if (pw.length < 5) {
			message = "the length of password at less five"
		}
		return message
	},

	checkInput: function(n) {
		return n != n.replace(/[^\w\d]/ig, '') 
	},


	RandomChar: function() {
		var charType = [[48, 57], [65, 90], [97, 122]]
		var typeNum = Math.floor(Math.random() * 3)
		var t = charType[typeNum]
		var num = Math.floor(Math.random() * (t[1] - t[0])) + t[0]
		var char = String.fromCharCode(num)
		return char
	},

	RandomUsername: function() {
		var res = ""
		for (var i = 0; i < 8; ++ i) {
			res += this.RandomChar()
		}
		return res
	},

	Time2Week: function(d, c) {
		var day = ["周一", "周二", "周三", "周四", "周五", "周六", "周日"]
		var cs = ["第一节","第二节","第三节","第四节","第五节","第六节","第七节","第八节"]
		return day[d-1] + cs[c-1]
	},

	Week2Time: function(s) {
		var day = ["周一", "周二", "周三", "周四", "周五", "周六", "周日"]
		var cs = ["第一节","第二节","第三节","第四节","第五节","第六节","第七节","第八节"]
		var count = ["first", "second", "third", "forth", "fifth", "sixth", "seventh", "eighth"]
		return [day.indexOf(s.slice(0,2)),count[cs.indexOf(s.slice(2,5))]]
	},

	StringToByte: function(str) {
		var re = [], idx
		for (var i = 0; i < str.length; i++){
			idx = str.charCodeAt(i)
			if (idx & 0xff00) {
				re.push(idx >> 8)
				re.push(idx & 0xff)
			} else {
				re.push(idx)
			}
		}
		return re;
	},

	ArraySum: function(a) {
		var r = 0
		for (var i = 0; i < a.length; ++ i) {
			r += a[i]
		}
		return r
	},

	RandomCode: function(s, ctime) {
		var date = new Date()
		var t = Number(date.toLocaleDateString().replace(/\//g, ''))
		var c = this.ArraySum(this.StringToByte(s))
		var o = this.ArraySum(this.StringToByte(ctime))
		return Math.sin(t+c+o).toString().substr(4,6)
	},

	ResolveJSON: function(s) {
		//format: '{"..." : "...", "..." : "..."}'
		return JSON.parse(s)
	},

	GetTime: function(t) {
		var date = new Date()

		if (t != "") {
			date = new Date(date.getTime() + Number(t) * 60 * 1000)
		}

		var y = date.getFullYear().toString()
		var m = date.getMonth() + 1 < 10 ? "0" + (date.getMonth() + 1).toString() : (date.getMonth() + 1).toString()
		var d = date.getDate() < 10 ? "0" + date.getDate().toString() : date.getDate().toString()
		var h = date.getHours() < 10 ? "0" + date.getHours().toString() : date.getHours().toString()
		var mi = date.getMinutes() < 10 ? "0" + date.getMinutes().toString() : date.getMinutes().toString()
		return y + m + d + h + mi
	},

	Before: function(d1, d2) {
		return Number(d1) <= Number(d2)
	},

	CheckClassTime: function(t1, t2) {
		var date = new Date()
		var b1 = (date.getDay() == 0 ? 7 : date.getDay()) == t1
		if (!b1) { return false }

		//8:00-9:40    10:00-11:40    14:20-16:00    16:10-18:00    19:00-20:40
		var classTimeList = [[800, 940], [1000, 1140], [1420, 1600], [1610, 1800], [1900, 2040]]
		var num = date.getHours() * 100 + date.getMinutes()
		return num <= classTimeList[t2-1][1] && num >= classTimeList[t2-1][0]
	},

	CheckNum: function(s) {
		var n = Number(s)
		return !isNaN(n) && n > 0
	},
}

export default FuncLib
