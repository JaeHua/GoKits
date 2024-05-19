<template>
<!--  <el-row :gutter="20">-->
<!--    &lt;!&ndash;    布局对齐tips&ndash;&gt;-->
<!--    <el-col :span="8"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--    <el-col :span="15"><div class="grid-content ep-bg-purple" />-->
<!--      <img src="@/assets/logo.png" width="15%" style="margin-left: 21%">-->
<!--      <el-card style="max-width: 500px">-->
<!--        <template #header>-->
<!--          <div class="card-header">-->
<!--            <el-text tag="b" size="large">登陆</el-text>-->
<!--          </div>-->
<!--        </template>-->
<!--        <el-form-->
<!--            ref="ruleFormRef"-->
<!--            style="max-width: 100%"-->
<!--            :model="ruleForm"-->
<!--            status-icon-->
<!--            :rules="rules"-->
<!--            label-width="auto"-->
<!--            class="demo-ruleForm"-->
<!--        >-->

<!--          <el-form-item label="QQ" prop="telephone" >-->
<!--            <el-input v-model.number="ruleForm.telephone" />-->
<!--          </el-form-item>-->
<!--          <el-form-item label="Password" prop="pass">-->
<!--            <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />-->
<!--          </el-form-item>-->


<!--          <el-form-item>-->
<!--            <el-button type="primary" @click="submitForm(ruleFormRef)">-->
<!--              Submit-->
<!--            </el-button>-->
<!--            <el-button @click="resetForm(ruleFormRef)">Reset</el-button>-->
<!--          </el-form-item>-->
<!--        </el-form>-->
<!--        <template #footer>没有账号？ <el-link type="primary" href="/register">前往注册</el-link></template>-->
<!--      </el-card>-->
<!--    </el-col>-->

<!--  </el-row>-->



  <div class="min-h-screen flex items-center justify-center bg-[#f4f7f6]">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8 p-12 bg-white shadow-lg rounded-lg max-w-6xl">
      <div class="flex flex-col justify-between">
        <div class="space-y-6">
          <img
              src="@/assets/logo.png"
              alt="Connexial Logistics"
              class="h-15 w-40"
              width="200"
              height="60"
              style="aspect-ratio: 200 / 60; object-fit: cover;"
          />
          <el-form
              ref="ruleFormRef"
              style="max-width: 100%"
              :model="ruleForm"
              status-icon
              :rules="rules"
              label-width="auto"
              class="demo-ruleForm"

          >
          <div class="space-y-4">
            <div class="space-y-2">
              <label
                  class="text-sm font-bold leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              >
                QQ Number
              </label>
              <el-form-item  prop="telephone" >
                <el-input v-model.number="ruleForm.telephone" />
              </el-form-item>
            </div>
            <div class="space-y-2">
              <label
                  class="text-sm font-bold leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              >
                Password
              </label>
              <el-form-item prop="pass">
                <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
              </el-form-item>
            </div>
            <div style="margin-top: 10%">
            <button type = "button" @click="submitForm(ruleFormRef)" class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-opacity-80 h-10 px-4 py-2 bg-[#e53e3e] text-white w-full">
              Log In
            </button>
            </div>
          </div>
          </el-form>
          <a class="text-[#e53e3e] block text-center" @click="$router.push({name:'Register'})">
            Register New User?
          </a>
        </div>
        <div class="text-center text-xs text-gray-500">
          © 2024 GoKits, Inc. All rights reserved.
          <a class="text-[#e53e3e]" href="#">
            Privacy Policy
          </a>
        </div>
      </div>
      <div class="flex flex-col justify-center bg-[#2c5282] text-white p-12 rounded-lg">
        <h2 class="text-3xl font-bold mb-4">Welcome to GoKits</h2>
        <p class="mb-6">
          GoKits front-end and back-end separation project demo project, is gradually developing new features.
        </p>

        <img
            alt="Cityscape"
            class="mt-8 rounded-lg shadow-lg"
            src="@/assets/blue.jpg"
            width="500px"
            style="aspect-ratio: 400 / 200; object-fit: cover;"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import {useStore} from "vuex";
const ruleFormRef = ref<FormInstance>()

//电话验证
const checkTelephone = (rule: any, value: any, callback: any) => {
  if (!value) {
    return callback(new Error('Please input the QQ'))
  }
  setTimeout(() => {
    if (!Number.isInteger(value)) {
      callback(new Error('Please input digits'))
    } else {
      value = value + '' //转变为字符串
      if (value.length < 6 || value.length>11) {
        // console.log("length:"+value.length)
        callback(new Error('QQ格式错误'))
      } else {
        callback()
      }
    }
  }, 100)

}
//密码验证
const validatePass = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('Please input the password'))
  } else {
    if (value.length < 6) {
      // console.log("length:"+value.length)
      callback(new Error('密码不得少于六位'))
    } else {
      callback()
    }
  }
}

const ruleForm = reactive({
  pass: '',
  telephone: '',
})

const rules = reactive<FormRules<typeof ruleForm>>({
  pass: [{ validator: validatePass, trigger: 'blur' }],
  telephone: [{ validator: checkTelephone, trigger: 'blur' }],
})

const router = useRouter()
const store = useStore()

//提交

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return

  await formEl.validate(async (valid) => {
    if (valid) {
      // console.log('user:'+ruleForm.name+' '+ruleForm.telephone+' '+ruleForm.pass)
      //请求
      const data = {"telephone":ruleForm.telephone.toString(),"password":ruleForm.pass}
      console.log(data)
      //dispatch触发action=>mutation
      store.dispatch('userModule/login',data).then(()=>{
        //跳转主页
        router.replace('home')
      }).catch((err)=>{
            (() => {
              ElMessage.error('Oops,'+err.response.data.msg)
            })()
          })
    } else {
      //数据验证失败的提示
      (() => {
        ElMessage({
          message: 'Warning, the data you input is invalid.',
          type: 'warning',
        })
      })()
      // throw new Error('Validation failed')
    }
  })
}

</script>

