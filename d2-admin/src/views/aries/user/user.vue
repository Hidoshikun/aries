<template>
  <d2-container>
    <template slot="header">用户 / 用户信息</template>
  </d2-container>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'user',
  data () {
    return {
      tabPosition: 'top',
      btn: {
        userInfoLoading: false,
        pwdLoading: false
      }
    }
  },
  created () {
  },
  methods: {
    ...mapActions('d2admin/account', [
      'logout'
    ])
  }
}
</script>

<style scoped>
</style>
