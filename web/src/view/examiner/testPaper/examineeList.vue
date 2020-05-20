<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="考生编号" min-width="80" prop="ID"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="userName"></el-table-column>
      <el-table-column label="考生姓名" min-width="150" prop="nickName"></el-table-column>
      <el-table-column label="登录ip" min-width="150" prop="ip"></el-table-column>
      <el-table-column label="在线状态" min-width="150" prop="onlineStatus">
        <template slot-scope="scope">
          <el-tag effect="dark" size="mini" type="success" v-if="scope.row.onlineStatus">在线</el-tag>
          <el-tag effect="dark" size="mini" type="info" v-else>离线</el-tag>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API
import { getUserList } from '@/api/user'
import infoList from '@/components/mixins/infoList'
import { mapGetters } from 'vuex'
export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getUserList,
      path: path,
      authOptions: [],
      addUserDialog: false,
      userInfo: {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入用户密码', trigger: 'blur' }
        ],
        nickName: [
          { required: true, message: '请输入用户昵称', trigger: 'blur' }
        ],
        authorityId: [
          { required: true, message: '请选择用户角色', trigger: 'blur' }
        ]
      }
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  methods: {},
  async created() {
    this.searchInfo = {
      authorityId: '777'
    }
    this.getTableData()
  }
}
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>