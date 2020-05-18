<template>
  <div class="svgInfo">
    <input @change="handlerUpload($event)" ref="fileupload" type="file" v-show="false" />
    <!-- 没有考试 -->
    <div v-if="status==0">没有进行中的考试</div>
<!-- 考试开始并且未同意考试须知 -->
    <div v-if="status==1&&!examinationRecord.agreement">
      <div style="padding:24px">{{testPaper.testPaperNote}}</div>
      <div style="text-align:center">
        <el-button @click="agree">同意并开始考试</el-button>
      </div>
    </div>

<!-- 考试开始并且已同意考试须知 -->

    <div v-if="status==1&&examinationRecord.agreement">
      <el-button
        @click="getTestPaperMould(testPaper)"
        style="float:right;margin-right:20px;"
        type="primary"
      >下载答题模板</el-button>
      <el-button @click="submitTest" style="float:right;margin-right:20px;" type="primary">提交答卷</el-button>
      <div v-html="svg"></div>
    </div>

<!-- 考试未开始 -->
    <div v-if="status==2">距离考试开始还有：{{countDownList}}</div>


    <!--自定义右键菜单html代码-->
    <ul :style="{left:left+'px',top:top+'px'}" class="contextmenu" v-show="contextMenuVisible">
      <li @click="download('flow')" v-if="activeNode.flow">下载流量文件</li>
      <li @click="download('configuration')" v-if="activeNode.configuration">下载配置文件</li>
      <li @click="download('log')" v-if="activeNode.log">下载日志文件</li>
      <li @click="download('sourceCode')" v-if="activeNode.sourceCode">下载源码文件</li>
    </ul>
  </div>
</template>
<script>
import {
  getTestPaperSvg,
  findAndCreateTestPaperSvgNode,
  downloadTestPaperSvgNode,
  getActiveTestPaper,
  getTestPaperMould
} from '@/api/testPaper'
import {
  createExaminationRecord,
  findExaminationRecord
} from '@/api/examinationRecord'

import { mapGetters } from 'vuex'

import axios from 'axios' // 引入axios

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 99999
})

export default {
  name: 'TestPaperInfo',
  data() {
    return {
      examinationRecord: {},
      testPaper: {},
      svg: '',
      contextMenuVisible: false,
      left: 0,
      top: 0,
      active: {},
      activeNode: {},
      formData: null,
      status: 0,
      agreement: false,
      param: null,
      isCollapse: false,
      isMobile: false,
      countDownList: '00天00时00分00秒'
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  methods: {
    timeFormat(param) {
      return param < 10 ? '0' + param : param
    },
     countDown() {
       const timefunc = async () => {
        let newTime = new Date().getTime() // 对结束时间进行处理渲染到页面
        let endTime = new Date(this.testPaper.testPaperStartTime).getTime()
        let obj = null // 如果活动未结束，对时间进行处理
        if (endTime - newTime > 0) {
          let time = (endTime - newTime) / 1000 // 获取天、时、分、秒
          let day = parseInt(time / (60 * 60 * 24))
          let hou = parseInt((time % (60 * 60 * 24)) / 3600)
          let min = parseInt(((time % (60 * 60 * 24)) % 3600) / 60)
          let sec = parseInt(((time % (60 * 60 * 24)) % 3600) % 60)
          obj = {
            day: this.timeFormat(day),
            hou: this.timeFormat(hou),
            min: this.timeFormat(min),
            sec: this.timeFormat(sec)
          }
        } else {
          // 活动已结束，全部设置为'00'
          obj = {
            day: '00',
            hou: '00',
            min: '00',
            sec: '00'
          }
          clearInterval(interval)
          const res = await getActiveTestPaper()
          this.status = res.data.status
          this.testPaper = res.data.testPaper
          if (this.status == 1) {
            await this.getExamination()
            if (this.examinationRecord.agreement) {
              this.getTest(res.data.testPaper.ID)
            }
          }
        }
        this.countDownList =
          obj.day + '天' + obj.hou + '时' + obj.min + '分' + obj.sec + '秒'
      }
      var interval = setInterval(timefunc, 1000)
    },
    submitTest() {
      this.param = new FormData()
      this.$refs.fileupload.click()
    },
    async handlerUpload(e) {
      this.param.append('file', e.target.files[0])
      this.param.append('testPaperId', Number(this.testPaper.ID))
      this.param.append('sysUserId', Number(this.userInfo.ID))
      const token = this.$store.getters['user/token']
      const res = await service.put(
        '/examinationRecord/submitTest',
        this.param,
        {
          headers: {
            'x-token': token
          }
        }
      )
      if (res.data.code == 0) {
        this.$message({
          type: 'success',
          message: '上传成功'
        })
      } else {
        this.$message({
          type: 'error',
          message: res.data.message
        })
      }
    },
    async getTestPaperMould(row) {
      const res = await getTestPaperMould({ ID: row.ID })
      const blob = new Blob([res])
      const fileName = row.testPaperMould.split('/mould/')[1]
      if ('download' in document.createElement('a')) {
        // 不是IE浏览器
        let url = window.URL.createObjectURL(blob)
        let link = document.createElement('a')
        link.style.display = 'none'
        link.href = url
        link.setAttribute('download', fileName)
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link) // 下载完成移除元素
        window.URL.revokeObjectURL(url) // 释放掉blob对象
      } else {
        // IE 10+
        window.navigator.msSaveBlob(blob, fileName)
      }
    },
    async agree() {
      const res = await createExaminationRecord({
        sysUserId: Number(this.userInfo.ID),
        testPaperId: Number(this.testPaper.ID),
        agreement: true
      })
      if (res.code == 0) {
        await this.getExamination()
        this.getTest(this.testPaper.ID)
      }
    },
    async download(type) {
      const res = await downloadTestPaperSvgNode({
        path: this.activeNode[type]
      })
      const blob = new Blob([res])
      const fileName = this.activeNode[type].split('/' + type + '/')[1]
      if ('download' in document.createElement('a')) {
        // 不是IE浏览器
        let url = window.URL.createObjectURL(blob)
        let link = document.createElement('a')
        link.style.display = 'none'
        link.href = url
        link.setAttribute('download', fileName)
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link) // 下载完成移除元素
        window.URL.revokeObjectURL(url) // 释放掉blob对象
      } else {
        // IE 10+
        window.navigator.msSaveBlob(blob, fileName)
      }
    },
    async openContextMenu(e) {
      window.event.returnValue = false
      var oEvent = e || event
      this.active = e
      const res = await findAndCreateTestPaperSvgNode({
        nodeId: this.active.target.id,
        testPaperID: Number(this.testPaper.ID)
      })
      if (res.code == 0) {
        this.activeNode = res.data.node
        this.contextMenuVisible = true
        let width
        if (this.isCollapse) {
          width = 54
        } else {
          width = 220
        }
        if (this.isMobile) {
          width = 0
        }
        this.left = oEvent.clientX - width
        this.top = oEvent.clientY + 10
      }
    },
    async getTest(id) {
      const res = await getTestPaperSvg({ ID: Number(id) })
      if (res.code == 0) {
        this.svg = res.data.retestPaper
        setTimeout(() => {
          const tspans = document.getElementsByTagName('tspan')
          for (let k = 0; k < tspans.length; k++) {
            const tspan = tspans[k].getElementsByTagName('tspan')
            if (tspan[0]) {
              tspan[0].oncontextmenu = this.openContextMenu
            }
          }
        }, 0)
      }
    },
    async getExamination() {
      const res = await findExaminationRecord({
        sysUserId: this.userInfo.ID,
        testPaperId: this.testPaper.ID
      })
      this.examinationRecord = res.data.examinationRecord
    }
  },
  async created() {
    this.countDown()
    this.$bus.on('mobile', isMobile => {
      this.isMobile = isMobile
    })
    this.$bus.on('collapse', isCollapse => {
      this.isCollapse = isCollapse
    })
    const res = await getActiveTestPaper()
    this.status = res.data.status
    this.testPaper = res.data.testPaper
    if (this.status == 1) {
      await this.getExamination()
      if (this.examinationRecord.agreement) {
        this.getTest(res.data.testPaper.ID)
      }
    }
  },
  watch: {
    contextMenuVisible() {
      if (this.contextMenuVisible) {
        document.body.addEventListener('click', () => {
          this.contextMenuVisible = false
        })
      } else {
        document.body.removeEventListener('click', () => {
          this.contextMenuVisible = false
        })
      }
    }
  }
}
</script>
<style lang="scss">
.svgInfo {
  display: flex;
  justify-content: center;
  height: calc(100vh - 195px);
  overflow: auto;
}

.contextmenu {
  width: 120px;
  margin: 0;
  border: 1px solid #ccc;
  background: #fff;
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 5px 0;
  border-radius: 4px;
  font-size: 14px;
  color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.2);
}
.contextmenu li {
  margin: 0;
  padding: 7px 16px;
}
.contextmenu li:hover {
  background: #f2f2f2;
  cursor: pointer;
}

tspan {
  tspan {
    cursor: pointer;
  }
}
</style>