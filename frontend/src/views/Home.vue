<template>


  <div class="text-blue-400 text-4xl font-bold animate-pulse flex justify-center mt-1">Powerful Tools</div>
  <main class="w-full max-w-6xl mx-auto px-4 py-12 md:px-6 md:py-16 lg:py-12 mt-0">
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">

      <a href="#">
        <div class="group relative rounded-lg overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300">
          <div class="aspect-w-16 aspect-h-9">
            <img
                src="@/assets/github.png"
                width="600"
                height="338"
                alt="Calendar Preview"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                style="aspect-ratio: 600 / 338; object-fit: cover;"
            />
          </div>
          <div class="p-4 bg-[#F0FEFF] dark:bg-[#0E7490]">
            <h2 class="text-lg font-semibold text-[#0C4A6E] dark:text-[#BEF8FD] mb-2">Github</h2>
            <p class="text-[#164E63] dark:text-[#A5F3FC] mb-4">
              赞一个star吧⭐
            </p>
            <div @click="handleClick" class="inline-flex items-center text-[#0284C7] hover:text-[#0369A1] dark:text-[#06B6D4] dark:hover:text-[#0891B2] transition-colors duration-300">
              Learn more
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
                  class="ml-1 w-4 h-4"
              >
                <path d="M5 12h14"></path>
                <path d="m12 5 7 7-7 7"></path>
              </svg>
            </div>
          </div>
        </div>
      </a>
      <a href="#">
        <div class="group relative rounded-lg overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300">
          <div class="aspect-w-16 aspect-h-9">
            <img
                src="@/assets/todolist.png"
                width="600"
                height="338"
                alt="Notes Preview"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                style="aspect-ratio: 600 / 338; object-fit: cover;"
            />
          </div>
          <div class="p-4 bg-[#FEF9C3] dark:bg-[#A16207]">
            <h2 class="text-lg font-semibold text-[#713F12] dark:text-[#FEFCE8] mb-2">Todolist</h2>
            <p class="text-[#92400E] dark:text-[#FDE68A] mb-4">
              What's to plan?.
            </p>
            <div @click="$router.replace('todolist')" class="inline-flex items-center text-[#A16207] hover:text-[#C2410C] dark:text-[#FBBF24] dark:hover:text-[#F59E0B] transition-colors duration-300">
              Learn more
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
                  class="ml-1 w-4 h-4"
              >
                <path d="M5 12h14"></path>
                <path d="m12 5 7 7-7 7"></path>
              </svg>
            </div>
          </div>
        </div>
      </a>
      <a href="#">
        <div class="group relative rounded-lg overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300">
          <div class="aspect-w-16 aspect-h-9">
            <img
                src="@/assets/more.png"
                width="600"
                height="338"
                alt="Chat Preview"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                style="aspect-ratio: 600 / 338; object-fit: cover;"
            />
          </div>
          <div class="p-4 bg-[#E5E7EB] dark:bg-[#1E293B]">
            <h2 class="text-lg font-semibold text-[#374151] dark:text-[#F3F4F6] mb-2">More</h2>
            <p class="text-[#6B7280] dark:text-[#94A3B8] mb-4">
             敬请期待!
            </p>
            <div class="inline-flex items-center text-[#6366F1] hover:text-[#4F46E5] dark:text-[#818CF8] dark:hover:text-[#6366F1] transition-colors duration-300">
              Learn more
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
                  class="ml-1 w-4 h-4"
              >
                <path d="M5 12h14"></path>
                <path d="m12 5 7 7-7 7"></path>
              </svg>
            </div>
          </div>
        </div>
      </a>

    </div>
  </main>
</template>

<script setup>
const handleClick = () =>{
  window.location.href = 'https://github.com/JaeHua';
}

</script>
<style scoped>
#title{
  width: 100%;
  font: bold 14px Arial;
  font-family: 'Helvetica Neue Bold', Helvetica, Arial, sans-serif;
}
#card:hover { cursor: pointer; background-color: #e8f0fe; }
.flex-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}


</style>