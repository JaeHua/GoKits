
<template>
    <!--    布局对齐tips-->
<!--    <el-container class="container">-->
<!--      <el-header class="_img">-->
<!--        <h1 >-->
<!--          <img-->
<!--              src="@/assets/profile.png"-->
<!--              style="width: 100px "-->
<!--          />-->
<!--        </h1>-->
<!--      </el-header>-->
<!--      <el-main class="parent-container">-->
<!--        <div class="todolist_box">-->
<!--          <div class="bar-message">-->
<!--            <div class="bar-message-text">-->
<!--              我的个人信息🤏🤏-->
<!--            </div>-->
<!--          </div>-->
<!--          <ul class="empty-tips">-->
<!--            <li>✔️ 用户名:{{userInfo.name}}</li>-->
<!--            <li>✔️ QQ:{{userInfo.telephone}}</li>-->

<!--            <li>🔒</li></ul>-->

<!--        </div>-->
<!--      </el-main>-->
<!--    </el-container>-->
  <div class="flex justify-center items-center h-screen bg-gradient-to-r from-[#9575cd] to-[#ce93d8] dark:from-[#1a237e] dark:to-[#673ab7]">
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg w-[80%] md:w-[60%] lg:w-[40%] p-8 h-auto min-h-[400px] flex flex-col items-center justify-center">
      <div class="flex flex-col items-center">
      <span class="relative flex shrink-0 overflow-hidden rounded-full h-32 w-32 mb-6">
        <img src="@/assets/tou.png"  />
      </span>
        <div class="text-center">
          <h2 class="text-3xl font-bold mb-2"> {{userInfo.name}}</h2>
          <p class="text-[#7b1fa2] dark:text-[#e1bee7] mb-8">
            "Live life to the fullest and make the most of every opportunity."
          </p>
        </div>
        <div class="flex justify-center w-full mb-6">
          <div class="flex items-center">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="mr-2 text-[#7b1fa2] dark:text-[#e1bee7]"
            >
              <path d="M22 17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9.5C2 7 4 5 6.5 5H18c2.2 0 4 1.8 4 4v8Z"></path>
              <polyline points="15,9 18,9 18,11"></polyline>
              <path d="M6.5 5C9 5 11 7 11 9.5V17a2 2 0 0 1-2 2v0"></path>
              <line x1="6" x2="7" y1="10" y2="10"></line>
            </svg>
            <span class="text-[#7b1fa2] dark:text-[#e1bee7]">{{userInfo.telephone}}@qq.com</span>
          </div>
        </div>
        <p class="mb-6 flex items-center text-[#7b1fa2] dark:text-[#e1bee7]">
          <svg @click="$router.push({name:'Todolist'})"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="mr-2"
          >
            <path d="M22 12h-2.48a2 2 0 0 0-1.93 1.46l-2.35 8.36a.25.25 0 0 1-.48 0L9.24 2.18a.25.25 0 0 0-.48 0l-2.35 8.36A2 2 0 0 1 4.49 12H2"></path>
          </svg>
          I have {{completedTodos}} outstanding tasks to complete.
        </p>
      </div>
    </div>
    <div class="absolute top-0 left-0 h-full w-1/4 bg-gradient-to-b from-[#9575cd] to-[#ce93d8] dark:from-[#1a237e] dark:to-[#673ab7] flex flex-col justify-between items-center p-4">
      <div class="text-white text-2xl font-bold animate-bounce">Explore the Possibilities</div>
      <div class="text-white text-lg">"Life is like riding a bicycle. To keep your balance, you must keep moving." - Albert Einstein</div>
      <button   @click="$router.push({name:'Home'})" class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:text-accent-foreground h-10 px-4 py-2 text-[#000080] dark:text-[#191970] hover:bg-white/20">
        Back Home
      </button>
    </div>
    <div class="absolute top-0 right-0 h-full w-1/4 bg-gradient-to-b from-[#9575cd] to-[#ce93d8] dark:from-[#1a237e] dark:to-[#673ab7] flex flex-col justify-between items-center p-4">
      <div class="text-white text-2xl font-bold animate-pulse">Embrace the Future</div>
      <div class="text-white text-lg">"Do one thing every day that scares you." - Eleanor Roosevelt</div>
      <div class="text-white text-lg"></div>

    </div>
  </div>
</template>

<script setup>
import { useStore } from 'vuex'
import {computed, onMounted, ref} from "vue";
import getTodo from "@/service/todoListService.js";
const store = useStore()
const userInfo = computed(()=>{
  return store.state.userModule.userInfo
})
const todos = ref([]);
onMounted(()=>{
  getTodo.getTodo().then((response) => {
    todos.value = response.data.data.todos;
  });

})
const completedTodos = computed(() => {
  let count = 0
  for(let todo of todos.value) {
    console.log(todo)
    if(!todo.Status) {
      count++
    }
  }
  return count
})
</script>

<style scoped>

p {
  cursor: pointer;
}

svg {
  transition: transform 0.2s;
}

p:hover svg {
  transform: scale(1.1);
}
</style>