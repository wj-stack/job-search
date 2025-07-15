<script setup>
import { ref } from 'vue'

const searchKeyword = ref('')
const region = ref('')
const jobType = ref('')

const handleSearch = () => {
  console.log('搜索关键词:', searchKeyword.value)
  console.log('地区:', region.value)
  console.log('职位类别:', jobType.value)
  // 可添加实际搜索逻辑
}
</script>

<template>
  <div class="home-container">
    <h2>职位搜索</h2>
    <input
      v-model="searchKeyword"
      type="text"
      placeholder="请输入搜索关键词"
      class="search-input"
    />
    <select v-model="region" class="filter-select">
      <option value="">选择地区</option>
      <option value="北京">北京</option>
      <option value="上海">上海</option>
      <option value="广州">广州</option>
    </select>
    <select v-model="jobType" class="filter-select">
      <option value="">选择职位类别</option>
      <option value="前端开发">前端开发</option>
      <option value="后端开发">后端开发</option>
      <option value="测试工程师">测试工程师</option>
    </select>
    <button @click="handleSearch" class="search-button">搜索</button>
  </div>
</template>

<style scoped>
.home-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  text-align: center;
}

.search-input {
  width: 300px;
  padding: 10px;
  margin: 10px;
}

.filter-select {
  width: 150px;
  padding: 10px;
  margin: 10px;
}

.search-button {
  padding: 10px 20px;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.search-button:hover {
  background-color: #3aa876;
}
</style>