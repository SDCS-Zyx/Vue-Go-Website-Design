<template>
  <div id="area">
    <div id="content">
      <div id="header"><p>Please enter your personal information</p></div>
      <el-form label-position="left" :rules="rules" ref="form" :model="form" label-width="120px">
        <el-form-item label="姓名：" prop="name">
          <el-input class="inp" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="学号/职工号：" prop="id">
          <el-input class="inp" v-model="form.id"></el-input>
        </el-form-item>
        <el-form-item label="类型：" prop="type">
          <el-select v-model="form.type" placeholder="请选择">
            <el-option label="教师" value="teacher"></el-option>
            <el-option label="学生" value="student"></el-option>

          </el-select>
        </el-form-item>
        <el-form-item label="学校：" prop="school">
          <el-select v-model="form.school" placeholder="请选择">
            <el-option label="中山大学" value="student"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="院系：" prop="dep">
          <el-select v-model="form.dep" placeholder="请选择" @change="changeDep">
            <el-option label="中国语言文学系" value="中国语言文学系"></el-option>
            <el-option label="历史学系" value="历史学系"></el-option>
            <el-option label="哲学系" value="哲学系"></el-option>
            <el-option label="社会学与人类学学院" value="社会学与人类学学院"></el-option>
            <el-option label="岭南学院" value="岭南学院"></el-option>
            <el-option label="外国语学院" value="外国语学院"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item id="subject" label="专业：" prop="subject">
          <el-select v-model="form.subject" placeholder="请选择">
            <el-option
            v-for="item in options[form.dep]"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item id="date" label="入学日期：">
        <el-col :span="11">
          <el-date-picker type="date" placeholder="入学日期" v-model="form.date1" style="width: 100%;"></el-date-picker>
        </el-col>
        <el-col class="line" :span="1">-</el-col>
        <el-col :span="11">
          <el-date-picker placeholder="毕业日期" v-model="form.date2" style="width: 100%;"></el-date-picker>
        </el-col>
      </el-form-item>
      <el-form-item label="性别：" prop="sex">
        <el-radio-group v-model="form.sex">
          <el-radio label="male"></el-radio>
          <el-radio label="female"></el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>

    <p>
      <el-button @click="enterAccount('form')">sign up</el-button>
    </p>
  </div>

  <div id="reg">
    <div id="form_acc">
      <el-form :model="account" ref="account" label-width="100px" class="demo-ruleForm">
       <el-form-item label="username:">
        <el-input class="inp" v-model="account.username" autocomplete="off" onkeyup="value=value.replace(/[^\w\d]/ig,'')"></el-input>
      </el-form-item> 
      <el-form-item label="password:">
        <el-input class="inp" v-model="account.password" autocomplete="off" onkeyup="value=value.replace(/[^\w\d]/ig,'')" show-password></el-input>
      </el-form-item>
    </el-form>
  </div>

  <div id="button_row">
    <el-button @click="back">back</el-button>
    <el-button @click="randomName">random username</el-button>
    <el-button @click="register">confirm</el-button>
  </div>

</div>

<div id="toast">
  <el-popover popper-class="pop" width="250" v-model="visible">
    <p id="toast_msg">{{message}}</p>
  </el-popover>
</div>
</div>
</template>

<script>

  import func from "../js/function.js"

  export default {
    data() {
      return {
        form: {
          name: '',
          id: null,
          type: '',
          school: '',
          dep: '',
          subject: '',
          date1: '',
          date2: '',
          sex: ''
        },
        account: {
          username: '',
          password: ''            
        },
        rules1: {
          name: [
          {required: true, message: "name should not be null", trigger: 'blur'},
          {min: 1, max: 10, message: "name cannot be so long", trigger: 'blur'}
          ],
          id: [
          {validator: this.validateNumber, required: true, message: "id should be number", trigger: 'blur', min: 1, max: 10000}
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
          {validator: this.validateNumber, required: true, message: "id should be number", trigger: 'blur', min: 1, max: 10000}
          ],
          type: [{required: true, message: "type cannot be null", trigger: 'blur'}],
          school: [{required: true, message: "school cannot be null", trigger: 'change'}],
          dep: [{required: true, message: "department cannot be null", trigger: 'change'}],
          sex: [{required: true, message: "sex cannot be null", trigger: 'change'}]
        },
        change: false,
        message: "",
        visible: false,
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
    watch: {
      'form.type' : {
        handler(nt, ot) {
          var s = document.getElementById("subject")
          var d = document.getElementById("date")
          if (nt == "teacher") {
            s.style.display = "none"
            d.style.display = "none"
          } else if (nt == "student") {
            s.style.display = ""
            d.style.display = ""
          }
        },
        deep:true,
        immediate:true
      }
    },
    methods: {
      validateNumber: function(rule, value, callback) {
        const val = Number(value)
        if (!val && val != 0 || value == '') {
          rule.message = 'Please enter number'
          callback('')
        } else if ((val < rule.min) || (val > rule.max)) {
          rule.message = 'Please enter true id'
          callback('')
        } else {
          callback()
        }
      },

      setting: function() {
        var width = document.body.clientWidth
        var height = parseInt(window.getComputedStyle(document
          .getElementById("area"), null).getPropertyValue("height"))

        var e = document.getElementById("content")
        var w2 = window.getComputedStyle(e, null).getPropertyValue("width")
        e.setAttribute("style", "margin-left:" + ((width-parseInt(w2))/2).toString() +"px;")
      },

      showToast: function() {
        this.visible = true
        var _self = this
        setTimeout(function(){
          _self.visible = false
        }, 2000)
      },

      randomName: function() {
        var name = func.RandomUsername()
        this.account["username"] = name
      },

      saveData: function() {
        var _self = this
        this.axios.post("http://localhost:8081"+'/Register/Register', {
          form: _self.form,
          account: _self.account
        })
        .then((response)=>{
          this.message = response.data
          if (this.message == "success") {
            this.$router.push({
              name: 'Home',
              params: {
                id: _self.form["id"]
              }
            })
          } else {
            this.showToast()
          }
        })
        .catch((error)=> {
          console.log(error)
        })
      },

      checkData: function(formName) {
        var res = false
        this.$refs[formName].validate(function(valid) {
          res = valid
        })
        return res
      },

      enterAccount: function(formName) {
        var checkRes = this.checkData(formName)
        if (!checkRes) {
          return
        }

        this.checkId(this.form["id"])
      },

      checkAccount: function(un) {
        this.axios.post("http://localhost:8081"+'/Register/checkAccount', {
          username: un
        })
        .then((response)=>{
          this.message = response.data 
          if (this.message == "not use") {
            this.saveData()
          } else {
            this.showToast()
          }
        })
        .catch((error)=> {
          console.log(error)
        })

      },

      checkId: function(i) {
        this.axios.post("http://localhost:8081"+'/Register/checkId', {
          id: i
        })
        .then((response)=>{
          this.message = response.data 
          if (this.message == "not use") {
            var cot = document.getElementById("content")
            cot.style.display = "none"

            var acc = document.getElementById("reg")
            acc.style.display = "block"
          } else {
            this.showToast()
          }
        })
        .catch((error)=> {
          console.log(error)
        })

      },

      register: function() {
        var un = this.account["username"]
        var pw = this.account["password"]

        this.message = func.CheckUsernameAndPassword(un, pw)
        if (this.message != "") {
          this.showToast()
          return
        }

        this.checkAccount(un)

      },

      back: function() {
       var cot = document.getElementById("content")
       cot.style.display = "block"

       var acc = document.getElementById("reg")
       acc.style.display = "none"
     },

     changeDep: function() {
        this.form.subject = this.options[this.form.dep].value
      }
   },
   computed: {
    rules: function() {
      if (this.form["type"] == "teacher") {
        return this.rules2
      } else {
        return this.rules1
      }
    }
  },
  mounted() {
    this.setting()
  }
}
</script>

<style scoped>
#area {  
  background: url("../assets/background.jpg") center center no-repeat;
  background-size: 100% 100%;
  position: absolute;
  width: 100%;
  height: 100%;
  text-align: center;
  margin-left: -8px;
  margin-top: -8px;
}

#content {
  position: absolute;
  height: wrap-content;
  width: 400px;
  border: 1px solid #000000;
  padding-top: 20px;
  padding-left: 20px;
  padding-right: 20px;
}

#header {
  font-size: 16px;
  border: 1px solid #eee;
  margin-bottom: 20px;  
}

.inp /deep/ .el-input__inner {
  /*width: 250px;*/
  font-size: 20px;
}

#reg {
  display: none;
  border: 1px solid #eee;
  width: 400px;
  height: wrap-content;
  margin:0 auto;
  padding-top: 30px;
  margin-top: 150px;
  padding-right: 20px;
  padding-bottom: 20px;
}

#toast {
  width: 200px;
  margin:0 auto;
}

#toast_msg{
  text-align: center;
  font-size: 20px;
}

#button_row {
  margin-left: 10px;
}

</style>