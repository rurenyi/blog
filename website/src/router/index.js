import VueRouter from "vue-router";

import userLogin from "@/pages/user-login.vue";
import blogList from "@/pages/blog-list.vue";
import blogContent from "@/pages/blog-content.vue";
import blogCreate from "@/pages/blog-create.vue"

export default new VueRouter({
	routes:[
        {
            name:"login",
            path:"/",
            component:userLogin
        },
        {
            name:"blogList",
            path:"/blog/:uid",
            component:blogList,
            props:true
        },
        {
            name:"blogContent",
            path:"/blog/:uid/:blogid",
            component:blogContent,
            props:true
        },
        {
            name:"blogCreate",
            path:"/blog/create/:uid",
            component:blogCreate,
            props:true
        }
    ]
})