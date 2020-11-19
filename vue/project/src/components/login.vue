<template>
  <div id="area">
    <div id="content">
      <div id="header"><p>用户登录</p></div>
      <form>
        <el-row id="un">
          <el-col :span="2"><img src="../assets/adm.png" alt="username" id="un_icon"></el-col>
          <el-col :span="8">
            <el-input class="inp" placeholder="username" v-model="input1" ref="n1" onkeyup="value=value.replace(/[^\w\d]/ig,'')"></el-input>
          </el-col>
        </el-row>

        <el-row id="pw">
          <el-col :span="2"><img src="../assets/key.png" alt="password" id="pw_icon"></el-col>
          <el-col :span="8">
            <el-input class="inp" placeholder="password" v-model="input2" ref="n2" onkeyup="value=value.replace(/[^\w\d]/ig,'')" show-password>
            </el-input>
          </el-col>
        </el-row>
      </form>
      

      <el-row id="two">
        <el-col :span="8">
          <el-button @click='clickLogin'>login</el-button>
        </el-col>
        <el-col :span="8">
          <el-button @click='clickRegister'>sign up</el-button>
        </el-col>
      </el-row>

    </div>

    <div id="toast">
      <el-popover popper-class="pop" width="250" v-model="visible">
        <p id="toast_msg">{{message}}</p>
      </el-popover>
    </div>

  </div>
</template>
<script>

  import axios from 'axios'
  import func from '../js/function.js'

  export default {
    data(){
      return{
        input1:"",
        input2:"",
        visible:false,
        message:""
      }
    },
    methods: {
      setting: function() {
        var el = document.getElementById("toast")
        var width = document.body.clientWidth
        var height = parseInt(window.getComputedStyle(document
          .getElementById("area"), null).getPropertyValue("height"))

        var w1 = window.getComputedStyle(el, null).getPropertyValue("width")
        el.setAttribute("style", "margin-top:" + (height/2).toString() + "px;margin-left:" + ((width-parseInt(w1))/2).toString() + "px;")

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

      clickLogin: function() {
        var un = this.$refs.n1.value
        var pw = this.$refs.n2.value
        this.message = func.CheckUsernameAndPassword(un, pw)
        if (this.message != "") {
          this.showToast()
          return
        }

        this.axios.post("http://localhost:8081"+'/Login', {
          username : un,
          password : pw
        })
        .then((response)=>{ 
          this.message = response.data
          // console.log(this.message)
          if (this.message.slice(0, 7) == "success") {
            var id = this.message.slice(11)
            this.$router.push({
              name: 'Home',
              params: {
                id: id
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

      clickRegister: function() {
        this.$router.push({
          name: 'Register', 
          params: {
            username : this.$refs.n1.value,
            password : this.$refs.n2.value
          }
        })
      }
    },
    mounted(){
      this.setting()
    }
  }
</script>

<style scoped>
* {
  z-index: 1;
}

#area {
  background: url("../assets/background.jpg") center center no-repeat;
  background-size: 100% 100%;
/*  filter:alpha(opacity=50);  
  -moz-opacity:0.5;  
  -khtml-opacity: 0.5;  
  opacity: 0.5;*/
  height:100%;
  width:100%;
  position: absolute;
  z-index: 0;
  margin-top: -8px;
  margin-left: -8px;
  text-align: center;
}

#content {
  margin-top: 100px;
  width: 450px;
  height: wrap-content;
  background-color: #696969;
  border-radius: 5px 5px 5px 5px;
  position: absolute;
  display: block;
  z-index: 1;
    filter:alpha(opacity=50);  
  -moz-opacity:0.5;  
  -khtml-opacity: 0.5;  
  opacity: 0.8;
}

#header {
  font-size: 50px;
  text-align: center;
  height: wrap-content;
  margin-bottom: 30px;
  margin-top: 50px;
  overflow: hidden;
}

#un, #pw {
  border-bottom: 1px solid #eee;
}

.inp /deep/ .el-input__inner {
  width: 250px;
  border: 0px;
  color: #FFFFFF;
  font-size: 20px;
  background-color: #696969;
}

.el-col {
  margin-left: 20px;
}

.el-row {
  margin-top: 20px;
  margin-left: 40px;
  margin-right: 40px;
}

#two {
  margin-top: 40px;
  padding-bottom: 40px;
  padding-left: 10px;
}

.el-button {
  font-size: 20px;
  margin-left: 30px;
}

#un_icon {
  margin-top: 30%;
}

#pw_icon {
  margin-top: 50%;
}

#toast_msg{
  text-align: center;
  font-size: 20px;
}

#toast {
  position: absolute;
  z-index: 2;
  width: 250px;
  filter:alpha(opacity=1);  
  -moz-opacity:0.5;  
  -khtml-opacity: 0.5;  
  opacity: 0.8;
}

</style>