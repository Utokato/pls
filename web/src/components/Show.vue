<script setup>
import {ref} from "vue";
import {useRoute} from "vue-router";
import {show} from "@/api/index.js";
import {ElMessage} from "element-plus";

const markdownText = ref('');

const command = useRoute().params.command;
show(command).then(res => {
  const r = res.data;
  if (r.code !== 0) {
    console.log("show error", r.message)
    ElMessage.error('请求失败，请稍后再试')
    return
  }
  if (r.data === null) {
    ElMessage('相关文档为空')
    return
  }

  markdownText.value = r.data
})

</script>

<template>
  <el-container>
    <el-main>
      <v-md-preview :text="markdownText"></v-md-preview>
    </el-main>
  </el-container>
</template>

<style scoped>

</style>