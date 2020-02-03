<template>
  <div class="nav">
    <el-menu class="el-menu-demo" mode="horizontal" router>
      <el-menu-item index="/">首页</el-menu-item>
      <el-menu-item index="/postvideo">视频投稿</el-menu-item>
      <el-menu-item index="/about">关于我们</el-menu-item>
    </el-menu>
    <!--<a href="https://www.ele.me" target="_blank">1</a>-->
  </div>
</template>

<script>
export default {
  name: 'NavBar'
}
</script>

<style>
.nav {
  margin-bottom: 20px;
}
</style>
