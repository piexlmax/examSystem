<template>
  <div class="big">
    <el-row v-if="userInfo.authorityId !== '777'">
      <el-col :span="12">
        <el-card class="box-card">
          <div slot="header">
            <span>考题概览</span>
          </div>
          <div>
            <ul>
              <li>共有考试{{testPaper.total}}场</li>
              <li>已进行{{doing}}场</li>
              <li>未开始{{waiting}}场</li>
            </ul>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="box-card">
          <div slot="header">
            <span>上次考试</span>
          </div>
          <div v-if="lastTestPaper.ID">
            <ul>
              <li>考试时间：{{lastTestPaper.testPaperStartTime|formatDate}} 至 {{lastTestPaper.testPaperEndTime|formatDate}}</li>
              <li>参考人数：{{lastRecord.total}}</li>
              <li>交卷人数：{{lastRecord.list&&lastRecord.list.filter(item=>item.testPath).length}}</li>
              <li>平均分:{{mean}}</li>
            </ul>
          </div>
          <div v-else>暂无上一场考试</div>
        </el-card>
      </el-col>
    </el-row>
    <el-row v-if="userInfo.authorityId == '777'">
      <el-col :span="12">
        <el-card class="box-card">
          <div slot="header">
            <span>考题概览</span>
          </div>
          <div>
            <ul>
              <li>共有考试{{testPaper.total}}场</li>
              <li>已进行{{doing}}场</li>
              <li>未开始{{waiting}}场</li>
            </ul>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="box-card">
          <div slot="header">
            <span>上次考试</span>
          </div>
          <div v-if="userLastRecord.sysUserId">
            <ul>
              <li>{{userLastRecord.testPaper.testPaperName}}</li>
              <li>考试时间：{{userLastRecord.testPaper.testPaperStartTime|formatDate}} 至 {{userLastRecord.testPaper.testPaperEndTime|formatDate}}</li>
              <li>得分:{{userLastRecord.score}}</li>
            </ul>
          </div>
          <div v-else>暂无上一场考试</div>
        </el-card>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-card class="box-card">
          <div slot="header">
            <span>正在考试</span>
          </div>
          <div v-if="activeTestPaper.ID">
            <ul>
              <li>{{activeTestPaper.testPaperName}}</li>
              <li v-if="userInfo.authorityId !== '777'">{{activeTestPaper.testPaperAuthor}}</li>
              <li>{{activeTestPaper.testPaperSubmitTimes?"允许多次提交":"只允许单次提交"}}</li>
              <li>考试时间：{{activeTestPaper.testPaperStartTime|formatDate}} 至 {{activeTestPaper.testPaperEndTime|formatDate}}</li>
            </ul>
          </div>
          <div v-else>暂无正在进行的考试</div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="box-card">
          <div slot="header">
            <span>即将开考</span>
          </div>
          <div v-if="nextTestPaper.ID">
            <ul>
              <li>{{nextTestPaper.testPaperName}}</li>
              <li v-if="userInfo.authorityId !== '777'">{{nextTestPaper.testPaperAuthor}}</li>
              <li>{{nextTestPaper.testPaperSubmitTimes?"允许多次提交":"只允许单次提交"}}</li>
              <li>考试时间：{{nextTestPaper.testPaperStartTime|formatDate}} 至 {{nextTestPaper.testPaperEndTime|formatDate}}</li>
            </ul>
          </div>
          <div v-else>暂无正在进行的考试</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { getTestPaperList } from '@/api/testPaper' //  此处请自行替换地址
import { getExaminationRecordList } from '@/api/examinationRecord'
import { formatTimeToStr } from '@/utils/data'
import { mapGetters } from 'vuex'
export default {
  name: 'Dashboard',
  data() {
    return {
      testPaper: {},
      doing: 0,
      waiting: 0,
      lastTestPaper: {},
      lastRecord: {},
      activeTestPaper: {},
      nextTestPaper: {},
      userLastRecord: {}
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo']),
    mean() {
      let sum = 0
      this.lastRecord &&
        this.lastRecord.list.map(item => {
          sum += item.score
        })

      sum = sum.toFixed(2)
      return sum
    }
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    }
  },
  async created() {
    const res = await getTestPaperList({ page: 1, pageSize: 99999999 })
    this.testPaper = res.data
    this.testPaper.list &&
      this.testPaper.list.sort(function(a, b) {
        return (
          new Date(a.testPaperEndTime).getTime() -
          new Date(b.testPaperEndTime).getTime()
        )
      })
    this.testPaper.list &&
      this.testPaper.list.map(item => {
        if (item.testPaperStatus) {
          this.activeTestPaper = item
        }
      })

    this.testPaper.list &&
      this.testPaper.list.map(item => {
        if (new Date() > new Date(item.testPaperEndTime)) {
          this.waiting += 1
          if(this.activeTestPaper.ID !==item.ID){
            this.lastTestPaper = item
          }
        } else {
          this.doing += 1
        }
        if (
          new Date(item.testPaperEndTime) >
            new Date(this.activeTestPaper.testPaperEndTime) &&
          !this.nextTestPaper.ID
        ) {
          this.nextTestPaper = item
        }
      })
    if (this.lastTestPaper.ID) {
      const laseRecord = await getExaminationRecordList({
        page: 1,
        pageSize: 9999999,
        testPaperId: this.lastTestPaper.ID
      })
      this.lastRecord = laseRecord.data
      if (this.userInfo.authorityId === '777') {
        this.userLastRecord =
          laseRecord.data.list.filter(
            item => item.sysUserId == this.userInfo.ID
          )[0] || {}
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.big {
  .box-card {
    margin: 5px;
    height: 200px;
    li {
      padding: 5px 10px;
    }
  }
  .item {
    padding: 18px 0;
  }
}
</style>
