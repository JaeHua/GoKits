<template>
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
                  用户名
                </label>
                          <el-form-item  prop="name" >
                            <el-input v-model="ruleForm.name" placeholder="选填" />
                          </el-form-item>
              </div>
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
              <div class="space-y-2">
              <label
                  class="text-sm font-bold leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              >
                Confirm
              </label>
                <el-form-item prop="checkPass">
                  <el-input
                      v-model="ruleForm.checkPass"
                      type="password"
                      autocomplete="off"
                  />
                </el-form-item>
              </div>
                <div class="space-y-2">
                  <label
                      class="text-sm font-bold leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  >
                    验证码
                  </label>
                  <el-form-item prop="">
                    <el-input
                        v-model="ruleForm.vCode"
                        type="number"
                        autocomplete="off"
                        style="width: 40%"
                    />
                    <el-button
                        type="primary"
                        :disabled="disableResend"
                        @click="submitVerify"
                    >
                      <span v-if="disableResend">
                        {{ countdown }}s
                      </span>
                      <span v-else>
                        Send
                      </span>
                    </el-button>
                  </el-form-item>
                </div>

              <div style="margin-top: 10%">
                <button type = "button" @click="submitForm(ruleFormRef)" class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-opacity-80 h-10 px-4 py-2 bg-[#e53e3e] text-white w-full">
                  Register
                </button>
              </div>
            </div>
          </el-form>
          <a class="text-[#e53e3e] block text-center" @click="$router.push({name:'Login'})">
            Already have an account?
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
            src="@/assets/blue1.jpg"
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
import { useStore, } from 'vuex'
import sendVerifyCode from '@/service/emailService'
import verifyTheCode from '@/service/emailService'

const ruleFormRef = ref<FormInstance>()

//QQ验证
const checkTelephone = (rule: any, value: any, callback: any) => {
  if (!value) {
    return callback(new Error('请输入QQ号'))
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
//密码重复验证
const validatePass2 = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('Please input the password again'))
  } else if (value !== ruleForm.pass) {
    callback(new Error("Two inputs don't match!"))
  } else {
    callback()
  }
}

//验证验证码
const checkvCode = (rule: any, value: any, callback: any) =>{
  if (value === '') {
    callback(new Error('请输入验证码'))
  } else if (value.length !== 6) {
    callback(new Error("验证码应该为6位"))
  } else {
    callback()
  }
}

const ruleForm = reactive({
  pass: '',
  checkPass: '',
  telephone: '',
  name:'',
  vCode:'',
})

//验证码逻辑
const countdown = ref(0);
const disableResend = ref(false);
let timer = null;

//验证码发送倒计时
const startCountdown = () => {
  timer = setInterval(() => {
    countdown.value--;
    if (countdown.value === 0) {
      clearInterval(timer);
      disableResend.value = false;
    }
  }, 1000);
};

const submitVerify= () => {
  // 发送验证码逻辑
  sendVerifyCode.sendVerifyCode({"mail":ruleForm.telephone.toString()+"@qq.com"})
  // 启动倒计时
  countdown.value = 60;
  disableResend.value = true;
  startCountdown();
};

//验证规则
const rules = reactive<FormRules<typeof ruleForm>>({
  pass: [{ validator: validatePass, trigger: 'blur' }],
  checkPass: [{ validator: validatePass2, trigger: 'blur' }],
  telephone: [{ validator: checkTelephone, trigger: 'blur' }],
  name:[{trigger:'blur'}], //trigger之后才能计算值
  vCode:[{validator:checkvCode,trigger:'blur'}]
})

//路由
const router = useRouter()

//vuex中store
const store = useStore()

//提交注册
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return

  await formEl.validate(async (valid) => {
    if (valid) {
      // console.log('user:'+ruleForm.name+' '+ruleForm.telephone+' '+ruleForm.pass)
      //请求
      // console.log("vcode",Vcode.value)

      verifyTheCode.verifyTheCode({"mail":ruleForm.telephone.toString()+"@qq.com","vcode":ruleForm.vCode}).then((response)=>{
        //验证码验证成功
        if (response.status===200)
        {
          const data = {"name":ruleForm.name,"telephone":ruleForm.telephone.toString(),"password":ruleForm.pass}
          //dispatch触发action=>mutation
          store.dispatch('userModule/register',data).then(()=>{
            //跳转主页
            router.replace('home')
          })
              .catch((err)=>{
            (() => {
              ElMessage.error('Ops,'+err.response.data.msg)
            })()

          })
        }else
        {
          (() => {
            ElMessage({
              message: '验证码有误.',
              type: 'warning',
            })
          })()
        }
      })

    } else {
      //数据验证失败的提示
      (() => {
        ElMessage({
          message: 'Warning, the data you input is invalid.',
          type: 'warning',
        })
      })()
    }
  })
}

</script>

