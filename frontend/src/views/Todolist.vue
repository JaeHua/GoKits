<template>
  <el-container class="container "style="margin-left: 4%">
    <el-header class="_img">
    <h1 >

      <img
          src="@/assets/todo.svg"
          style="width: 190px "
      />
    </h1>
      <div class="pendulums">
        <div class="pendulum">
          <div class="bar"></div>
          <div class="motion">
            <div class="string"></div>
            <div class="weight"></div>
          </div>
        </div>
        <div class="pendulum shadow">
          <div class="bar"></div>
          <div class="motion">
            <div class="string"></div>
            <div class="weight"></div>
          </div>
        </div>
      </div>
  </el-header>
    <el-main class="parent-container">
      <input type="text"  class="add-content" v-model="data.title" placeholder="新增待办事项..." ref="inputRef"  @keyup.enter="handleSubmit"/>
      <button class="btn submit-btn" type="button"  @click="handleSubmit">提交</button>
    </el-main>
    <el-main  class="parent-container">
    <div class="todolist_box">
      <div class="bar-message">
        <div class="bar-message-text">
          今日事今日毕，勿将今事待明日! 🤏🤏
        </div>
      </div>

      <ul class="todo-list">
        <ul class="empty-tips" v-if="!todos.length">
          <li> 添加你的第一个代办事项！📝</li>
          <li>食用方法💡：</li>
          <li>✔️ 所有提交操作支持 Enter 回车键提交~</li>
          <li>✔️ 双击已有的 Todo 可进行编辑~</li>
          <li>🔒 所有的 Todo 数据存储在远程数据库~</li>
        </ul>
        <li  v-for="todo in todos" class='todo-item'>
          <div class="todo-content" :class='{completed:todo.Status}'>{{todo.Title}}</div>
          <div class="todo-btn btn-unfinish" v-if="todo.Status" @click="handlePut(todo.ID)">
            <img
                src="@/assets/finish.svg"
                style="width: 100%;margin-top: 15% "
            />
          </div>
          <div class="todo-btn btn-finish" v-if="!todo.Status"  @click="handlePut(todo.ID)">

          </div>

          <div class="todo-btn btn-delete" @click="handleDelete(todo.ID)">
            <img
                src="@/assets/x.svg"
                style="width: 75% "
            /></div>
        </li>
      </ul>
    </div>
    </el-main>
    <div  class="side-bar" @click="$router.replace('home')" style="cursor: pointer;">
      回到主页
    </div>
  </el-container>

</template>
<script setup>
import {onBeforeUnmount, onMounted, reactive, ref} from "vue";
import {useStore} from "vuex";
import postTodo from "@/service/todoListService"
import getTodo from "@/service/todoListService"
import deleteTodo from "@/service/todoListService"
import updateTodo from "@/service/todoListService"
import { h } from 'vue'
import { ElNotification } from 'element-plus'
const store = useStore()
const todos = ref([]);

//侧边弹窗
const open1 = () => {
  ElNotification({
    title: 'Tips',
    message: h('i', { style: 'color: teal' }, 'todo不能为空哦'),
  })
}
//todo结构体
const data = reactive({
  title:'',
  status:0,
  telephone:store.state.userModule.userInfo.telephone,
})

//绑定输入框的值
const inputRef = ref(null);

//点击post操作
const handleSubmit = () => {
  if (data.title==='')
  {
    open1()
    return
  }
  console.log(data)
  postTodo.postTodo(data)
      .then(()=>{
        getTodo.getTodo().then((response) => {
          todos.value = response.data.data.todos;
        });
      });
  inputRef.value.value = ""; // 清空输入框值
  data.title = ""; //数据也不要了
};

//delete操作
const handleDelete = (id) =>{
  deleteTodo.deleteTodo(id).then(() => {
    getTodo.getTodo().then((response) => {
      todos.value = response.data.data.todos;
    });
  })
}

//put操作
const handlePut = (id)=>{
  updateTodo.updateTodo(id).then(() => {
    getTodo.getTodo().then((response) => {
      todos.value = response.data.data.todos;
    });
  })
}
onMounted(() => {
  //背景设置
  document
      .querySelector("body")
      .setAttribute("style", "background-color: #e8f6ff");
  document.querySelector("body").classList.add("background-dots");

  // 挂载获取数据
  getTodo.getTodo().then((response) => {
    todos.value = response.data.data.todos;
  });

  //获得光标焦点
  inputRef.value.focus()

  //背景移除
  onBeforeUnmount(() => {
    document.querySelector("body").removeAttribute("style")
  })
});
</script>
<style scoped>

.container{
  margin-top: 20px;
}
._img{
  margin:auto;
}
.parent-container {
  display: flex;
  justify-content: center;
  align-items: center;

}

.add-content {
  margin-top: 15px;
  width: 40%;
  font-size: 18px;
  cursor: text;
  border: 2px solid #33322E;
  height: 55px;
  line-height: 64px;
  text-indent: 12px;
  transition: all 0.35s;
  border-radius: 12px 0 0 12px;
  box-shadow: 4px 4px 0 #33322E;
}

.submit-btn:hover,
.add-content:hover {
  box-shadow: 0 0 0 0 #333;
}
.submit-btn {

  height: 60px;
  width: 96px;
  padding: 0;
  text-align: center;
  position: relative;
  right: 2px;
  top: 7.5px;
  border: 0;
  border-left: 2px solid #33322E;
  border-radius: 0 12px 12px 0;
  cursor: pointer;
  background: #ffd6e9;
  transition: all .25s;
  font-size: 20px;
  text-indent: 12px;
  box-shadow: 4px 4px 0 #33322E;
}

.btn {
  border: 2px solid #33322E;
  text-align: center;
  display: block;
  position: relative;
}
.todolist_box {
  position: relative;
  background: #fff;
  box-shadow: 4px 4px 0 #33322E;
  border: 2px solid #33322E;
  border-radius: 12px;
  margin: 0;
  padding-top: 0;
  transition: all .5s ease;
  max-width: 100%;
  width: 48%;
  height: auto;
  overflow: hidden;
}
.bar-message {
  position: relative;
  display: block;
  align-items: center;
  justify-self: flex-start;
  border-bottom: 2px solid #33322E;

}
.bar-message-text {
  display: inline-block;
  text-align: center;
  font-weight: 600;
  padding: 12px;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
.empty-tips {
  text-align: left;
  max-width: 480px;
  width: 90%;
  margin: 10px auto  ;

  top: 15%;
  left: 2%;
  font-size: 16px;
  color: #606060;
  line-height: 10px;
  letter-spacing: 1.8px;
}
ul, li {
  list-style-type: none;
  margin: 0;
  padding: 0;
  line-height: 1;
}
.side-bar {
  right: 100px;
  top: 0;
  height: 30px;
  font-size: 18px;
  width: 85px;
  margin-top: -20%;
  margin-left: 80%;
  justify-content: flex-start;
  align-items: flex-start;
  border: 2px solid #33322E;
  background: #dceebc;
  box-shadow: 4px 4px 0 #33322E;
  text-align: center;
  transition: all .3s;
  z-index: 999;
}
.side-bar:hover{
  box-shadow: 0 0 0 0 #33322E;

}
.todo-list {
  padding: 32px 36px 36px;
  min-height: 300px;
  transition: all .5s ease;

}
.todo-content {
  border: 2px solid #33322E;
  box-shadow: 4px 4px 0 #33322E;

  background: #F9F3E5;
  border-radius: 12px;
  width: 100%;
  position: relative;
  box-sizing: border-box;
  display: block;
  line-height: 1;
  overflow-wrap: break-word;
  padding: 20px 56px;
  cursor: pointer;
  min-height: 60px;
  transition: all 0.35s;
}
.todo-content:hover{
  box-shadow: 0 0 0 0 #33322E;

}
.btn-finish {
  position: relative;
  margin-top: -55px;
  left: 3%;
  top: 10px;
  width: 25px;
  height: 25px;
  border-radius: 20px;
  z-index: 9999;
}
.btn-finish:hover{
  box-shadow: -4px 4px 0 #33322E;
  transform: translate(5px, -5px);

}
.todo-btn {
  background: none;
  display: block;
  cursor: pointer;
  border: 2px solid #33322E;
  transition: all .2s;
  background: #fff;
}
.todo-btn:hover {
  background: #8cd4cb;
  border: 2px solid #33322E;


}
.todo-list li {
  position: relative;
  border-radius: 12px;
  width: 100%;
  transition: all .5s ease;
  display: block;
  margin-bottom: 18px;
}
.btn-delete {
  position: relative;
  background: #fff;
  top: -17px;

  right: -90%;
  height: 25px;
  width: 25px;
  border-radius: 12px;
  font-size: 13px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;

}
.btn-delete:hover{
  box-shadow: 4px 4px 0 #33322E;
  transform: translate(-5px, -5px);
}
.completed {
  background: #D0F4F0;
  color: rgba(51,50,46,.535);
  text-decoration: line-through;
}
.btn-unfinish {
  position: relative;
  margin-top: -55px;
  left: 3%;
  top: 10px;
  width: 25px;
  height: 25px;
  border-radius: 20px;
  z-index: 9999;
  background: #8CD4CB;
}
.btn-unfinish:hover{
  box-shadow: -4px 4px 0 #33322E;
  transform: translate(5px, -5px);

}
:root {
  --body-bg: #e8f6ff;
  --border-radius: 12px;
  --border: 2px solid #33322E;
  --box-shadow: 4px 4px 0 #33322E;
  --box-shadow-reverse: -4px 4px 0 #33322E;
  --padding: 20px 24px;
  --btn-padding: 12px 24px;
  --btn-small-padding: 10px 20px;
  --font-color: #33322E;
  --font-size-base: 16px;
  --placeholder: rgba(51,50,46,.535);
  --font-color-complete: rgba(51,50,46,.535);
  --bg-normal: #F9F3E5;
  --bg-submit: #ffd6e9;
  --bg-completed: #D0F4F0;
  --bg-discard: #FFF0EE;
  --bg-deleted: #ddd;
  --bg-edit: #fbeef3;
  --normal: #f5d99e;
  --completed: #8CD4CB;
  --deleted: #F6A89E;
  --black: #33322E;
  --font: "PingFang SC", "Lantinghei SC", "Microsoft YaHei", "HanHei SC", "Helvetica Neue", "Open Sans", Arial, "Hiragino Sans GB", 微软雅黑, STHeiti, "WenQuanYi Micro Hei", SimSun, sans-serif, HYWenHei-GEW !important;
}


</style>