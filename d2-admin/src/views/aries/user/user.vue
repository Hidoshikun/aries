<template>
  <d2-container>
    <template slot="header">用户 / 用户信息</template>
    <el-tabs :tab-position="tabPosition" type="border-card">
      <el-tab-pane label="个人信息" style="width: 500px">
        <el-form :model="userForm" ref="userForm" :rules="userRules" label-width="100px">
          <el-form-item label="用户名" prop="username">
            <el-input size="small" type="text" autocomplete="off" placeholder="用户名" readonly
                      v-model="userForm.username"></el-input>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input size="small" type="text" autocomplete="off" placeholder="邮箱"
                      v-model="userForm.email"></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </d2-container>
</template>

<script>
import { mapActions } from 'vuex'
import { getAllUsers, updateUser, updateUserPwd } from '@/api/aries/user'

export default {
  name: 'user',
  data () {
    return {
      tabPosition: 'top',
      btn: {
        userInfoLoading: false,
        pwdLoading: false
      },
      userForm: {
        ID: null,
        username: '',
        email: '',
      },
      userRules: {
        username: [
          { required: true, trigger: 'blur', message: '请输入用户名' },
          { min: 3, max: 30, trigger: 'blur', message: '用户名长度必须为 3 到 30 位' }
        ],
        email: [
          { required: true, trigger: 'blur', message: '请输入邮箱' },
          { max: 30, trigger: 'blur', message: '邮箱长度不能超过 30 位' },
          { type: 'email', trigger: 'blur', message: '邮箱格式不正确' }
        ]
      },
    }
  },
  created () {
    this.fetchUserInfo()
  },
  methods: {
    ...mapActions('d2admin/account', [
      'logout'
    ]),
    // 获取用户信息
    fetchUserInfo () {
      getAllUsers()
        .then(res => {
          this.userForm.ID = res.data[0].ID
          this.userForm.username = res.data[0].username
          this.userForm.email = res.data[0].email
        })
        .catch(() => {
        })
    },
  }
}
</script>

<style scoped>
</style>
