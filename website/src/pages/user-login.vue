<template>
  <div class="login">
    <h1>Login</h1>
    <div>
      用户:<input type="text" placeholder="请输入用户名" v-model="UserName">
      <br/>
      密码:<input type="password" placeholder="请输入密码" v-model="Password">
    </div>
    <button @click="login">登陆</button>
  </div>
</template>

<script>
import axios from "axios";
import cryptoJs from "crypto-js";

export default {
    name : "user-login",
    data() {
      return {
        UserName:"",
        Password:""
      }
    },
    methods: {
      login(){
        let hash = cryptoJs.MD5(this.Password).toString()
        let UserLogin = axios.create({
          baseURL: this.$config.serverUrl,
          timeout: 1000
        })
        UserLogin.post("/login",{"username":this.UserName,"password":hash})
          .then((responce) => {this.$router.push({
            name:"blogList",
            params:{"uid":responce.data.id}
          })})
          .catch((response) => { console.log(response)})
      }
    },
    
}
</script>

<style scoped>

.login{
  display: flex;
  flex-direction: column;
  align-items: center;
}

</style>