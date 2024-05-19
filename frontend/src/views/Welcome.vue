<script setup>
import {ref, onMounted, watch, onBeforeUnmount} from 'vue'
import { DArrowRight } from '@element-plus/icons-vue'

const handleClick = () =>{
  window.location.href = 'https://github.com/JaeHua';
}
const message = ref('')
const message1 = ref('')
const index = ref(0)
const index1 = ref(0)
const showButton = ref(false)

const greetingMessage = 'A Powerful Tool Synthesis Site Under Development!'
const greetingMessage1 = '  This is a demo project for development to learn vue+go and some backend knowledge!'

function typeWriter() {
  if (index.value < greetingMessage.length) {
    message.value += greetingMessage.charAt(index.value);
    index.value++;
    setTimeout(typeWriter, 100);
  } else if (index1.value < greetingMessage1.length) {
    message1.value += greetingMessage1.charAt(index1.value);
    index1.value++;
    setTimeout(typeWriter, 100);
  } else {
    showButton.value = true; // 当文字打印完毕时,显示按钮
  }
}


watch(message, () => {
  if (message.value.length === greetingMessage.length) {
    showButton.value = true
  }
})
onMounted(() => {
  typeWriter()

})

</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8 items-center">
      <div>
        <h1 class="text-5xl font-bold leading-tight">{{message}}</h1>
        <p class="mt-4 text-lg text-gray-600">
        {{message1}}
        </p>
        <div class="mt-8 flex gap-4">
          <button v-if="showButton" @click="$router.push({name:'Home'})" class="inline-flex items-center justify-center whitespace-nowrap text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 bg-blue-600 text-white px-6 py-2 rounded-md shadow-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50">
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
                class="mr-2 h-4 w-4"
            >
              <path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09z"></path>
              <path d="m12 15-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.35 22.35 0 0 1-4 2z"></path>
              <path d="M9 12H4s.55-3.03 2-4c1.62-1.08 5 0 5 0"></path>
              <path d="M12 15v5s3.03-.55 4-2c1.08-1.62 0-5 0-5"></path>
            </svg>
            Get Started
          </button>
          <button @click="handleClick" v-if="showButton" class="inline-flex items-center justify-center whitespace-nowrap text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 bg-gray-300 text-gray-800 px-6 py-2 rounded-md shadow-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-300 focus:ring-opacity-50">
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
                class="mr-2 h-4 w-4"
            >
              <path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"></path>
              <path d="M9 18c-4.51 2-5-2-7-2"></path>
            </svg>
            Github
          </button>
        </div>
      </div>
      <div class="flex justify-center">
        <img
            src="@/assets/tool.jpg"
            alt="Tool Synthesis"
            class="rounded-lg shadow-lg"
            width="600"
            height="400"
            style="height: 400px; width: 600px; aspect-ratio: 600 / 400; object-fit: cover;"
        />
      </div>
    </div>
  </div>
</template>
<style scoped>


</style>