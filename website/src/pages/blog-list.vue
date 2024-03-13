<template>
  <div class="content">
    <div v-for="b in blogs" :key="b.blogID">

      <router-link :to="`/blog/${uid}/${b.blogID}`">{{b.title}}</router-link>
      <!-- <router-link :to="{
        name:'blogContent',
        params:{
						uid:b.uid,
            blogid:b.blogID
					}
        }">{{ b.title }}</router-link>
      <br> -->
        创建时间:{{ b.createTime }}
        <br>
        <br>
    </div>
    <button @click="CreateBlog">创建博客</button>
  </div>
</template>

<script>
export default {
  name: 'WebsiteBlogList',
  props:["uid"],

  data() {
    return {
        blogs:[]
    };
  },

  created(){
    this.baseaxios.get(`/blog/${this.uid}`)
        .then((response) => {response.data.map((item) => {this.blogs.push(item)})})
        .catch((response) => {console.log(response);})
  },
  methods: {
    CreateBlog(){
      this.$router.push({
        name:"blogCreate",
        params:{"uid":this.uid}
      })
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
</style>