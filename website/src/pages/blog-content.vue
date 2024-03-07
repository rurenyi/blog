<template>
  <div class="content">
    <h1>{{ title }}</h1>
    <div class="blog"> {{ content }} </div>
    <div>
      <button @click="back">后退</button>
      <button @click="deleteBlog">删除</button>
    </div>

  </div>

</template>

<script>
import axios from 'axios';
export default {
  name: 'WebsiteBlogContent',
  props:["uid","blogid"],

  data() {
    return {
      title:"",
      content:""
    };
  },

  created(){
    let blogContent = axios.create({
        baseURL: this.$config.serverUrl,
        timeout:1000
    })
    blogContent.get(`/blog/${this.uid}/${this.blogid}`)
        .then((response) => {
          this.title = response.data.title
          this.content = response.data.content
        })
        .catch((response) => {console.log(response);})
  },
  methods: {
    back(){
      this.$router.back()
    },
    deleteBlog(){
      let delet = axios.create({
        baseURL: this.$config.serverUrl,
        timeout:1000
      })
      delet.delete(`/blog/delete/${this.uid}/${this.blogid}`)
      .then(() => {
          this.$router.back()
        })
        .catch((response) => {console.log(response);})
    }
  },
};
</script>

<style scoped>
.content{
    display: flex;
    flex-direction: column;
    align-items:center ;
}

.blog{
/* background: grey; */
width: 50vh;
height: 70vh;
overflow: auto;
}
</style>