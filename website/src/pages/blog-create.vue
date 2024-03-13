<template>
  <div class="content">
    标题: <input type="text" name="title" v-model="title"><br>
    博客内容<br>
    <textarea name="" id="" cols="30" rows="10" v-model="content"></textarea>
    <button @click="Save">保存</button>
  </div>
</template>

<script>

export default {
  name: 'WebsiteBlogCreate',

  props:["uid"],

  data() {
    return {
      title:"",
      content:""
    };
  },

  mounted() {

  },
  methods: {
    Save(){
      this.baseaxios.post(`/blog/create/${this.uid}`,{
            "uid":this.uid,
            "content":this.content,
            "title":this.title,
        })
        .then(() => {
          this.$router.back()
        })
        .catch((response) => {console.log(response);})
    }
  },
};
</script>
<style>
.content{
    display: flex;
    flex-direction: column;
    align-items:center ;
}
.blog{
/* background: grey; */
width: 50vh;
height: 30vh;
overflow: auto;

}
</style>