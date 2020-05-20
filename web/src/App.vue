<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import { keepOnline } from '@/api/user'
export default {
  name: 'app',
  components: {},
  created() {
    const token = this.$store.getters['user/token']
    if (token) {
      keepOnline()
    }
    setInterval(() => {
      if (token) {
        keepOnline()
      }
    }, 1000 * 60 * 10)
  }
}
</script>

<style lang="scss">
// 引入初始化样式
@import '@/style/main.scss';
@import '@/style/base.scss';
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
}
</style>
