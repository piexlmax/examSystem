<template>
  <div>
      <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="考题">
          <el-select v-model="searchInfo.testPaperId" @change="onSubmit" filterable placeholder="请选择考题">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label+(item.status?'(进行中)':'(未进行)')"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="offlineAppraise">离线审阅当前考题</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
     <el-table-column label="考题编号" width="100">
         <template slot-scope="scope">{{scope.row.testPaperId}}</template>
    </el-table-column>
    <el-table-column label="考题" width="180">
         <template slot-scope="scope">{{scope.row.testPaper.testPaperName}}</template>
    </el-table-column>
    <el-table-column label="考生编号" width="100">
         <template slot-scope="scope">{{scope.row.sysUserId}}</template>
    </el-table-column>
    <el-table-column label="考生" width="180">
         <template slot-scope="scope">{{scope.row.sysUser.nickName}}</template>
    </el-table-column>
    
     <el-table-column label="同意协议" prop="agreement" width="120">
       <template slot-scope="scope">
         <div>
           {{scope.row.agreement?"已同意":"未同意"}}
         </div>
       </template>
     </el-table-column>
    
     <el-table-column label="答题状况" prop="testPath" width="120">
       <template slot-scope="scope">
         <div>
           <el-button v-if="scope.row.testPath" @click="downloadTestPaperSvgNode(scope.row)">
             下载答卷
           </el-button>
           <span v-else>
             未作答
           </span>
         </div>
       </template>
     </el-table-column>
    
     <el-table-column label="得分" prop="score" width="120"></el-table-column>
     <el-table-column label="评语" prop="appraise" min-width="120"></el-table-column>
    
      <el-table-column label="按钮组" width="80">
        <template slot-scope="scope">
          <div>
           <el-button type="text" @click="openAppraise(scope.row)">评阅</el-button>
          </div>
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

    
    <el-dialog
     title="评阅"
    :visible.sync="dialogVisible"
    width="30%"
    >
    <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
      <el-form-item label="在线审阅">
        <el-button type="primary" @click="viewTest">开发中</el-button>
      </el-form-item>
      <el-form-item label="得分" prop="score">
        <el-input-number v-model="formData.score" placeholder="得分" :precision='1' controls-position=right>
        </el-input-number>
      </el-form-item>
      <el-form-item label="评语" prop="appraise">
        <el-input v-model="formData.appraise" type="textarea" placeholder="请输入评语"
          :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
      </el-form-item>
    </el-form>
  <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="enterAppraise">确 定</el-button>
  </span>
    </el-dialog>
  </div>
</template>

<script>
import {
    getExaminationRecordList,
    updateExaminationRecord,
    offlineAppraise
} from "@/api/examinationRecord";  //  此处请自行替换地址

import {
  downloadTestPaperSvgNode,
  getTestPaperList
} from '@/api/testPaper'
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { mapGetters } from "vuex";

export default {
  name: "ExaminerRecord",
  mixins: [infoList],
  data() {
    return {
      listApi: getExaminationRecordList,
      type: "",
      options:[],
      searchInfo:{
        testPaperId:0
      },
       dialogVisible:false,
      formData: {
        ID:0,
        score: 0,
        appraise: '',
      },
    rules: {
        score: [{
          required: true,
          message: '得分',
          trigger: 'blur'
        }],
        appraise: [{
          required: true,
          message: '请输入评语',
          trigger: 'blur'
        }],
      },
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"])
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    }
  },
  methods: {
    viewTest(){
      console.logg("开发中")
    },
    async offlineAppraise(){
      const res = await offlineAppraise({testPaperId:String(this.searchInfo.testPaperId)})
       const blob = new Blob([res]);
      if ("download" in document.createElement("a")) {
        // 不是IE浏览器
        let url = window.URL.createObjectURL(blob);
        let link = document.createElement("a");
        link.style.display = "none";
        link.href = url;
        link.setAttribute("download", this.activeTestName+".zip");
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link); // 下载完成移除元素
        window.URL.revokeObjectURL(url); // 释放掉blob对象
      } else {
        // IE 10+
        window.navigator.msSaveBlob(blob,  this.activeTestName+".zip");
      }
    },
    onSubmit(index){
      this.activeTestName = this.options.filter(item=>item.value==index)[0].label
      this.getTableData()
    },
    closeAppraise(){
       this.formData = {
        ID:0,
        score: 0,
        appraise: '',
      }
      this.dialogVisible = false
    },
    async enterAppraise(){
      const res = await updateExaminationRecord(this.formData)
      if(res.code == 0){
        this.$message({
          type:"success",
          message:"评分成功"
        })
      this.closeAppraise()
      this.getTableData()
      }
    },
    openAppraise(row){
      if(!row.testPath){
        this.$message({
          type:"error",
          message:"考生尚未作答"
        })
        return
      }
      this.formData.ID = row.ID
      this.formData.appraise = row.appraise
      this.formData.score = row.score
      this.dialogVisible = true
    },
      async downloadTestPaperSvgNode(row){
        const res = await downloadTestPaperSvgNode({path:row.testPath})
        const blob = new Blob([res]);
        const fileName = row.testPath.split("/"+row.testPaperId+"/")[1];
      if ("download" in document.createElement("a")) {
        // 不是IE浏览器
        let url = window.URL.createObjectURL(blob);
        let link = document.createElement("a");
        link.style.display = "none";
        link.href = url;
        link.setAttribute("download", fileName);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link); // 下载完成移除元素
        window.URL.revokeObjectURL(url); // 释放掉blob对象
      } else {
        // IE 10+
        window.navigator.msSaveBlob(blob, fileName);
      }
      }
  },
  async created() {
     const res = await getTestPaperList({page:1,pageSize:9999})
     const arr = []
     if(res.code == 0){
       res.data.list.map(item=>{
         let obj = {
           value:item.ID,
           label:item.testPaperName,
           status:item.testPaperStatus
         }
         if(item.testPaperStatus){
           this.searchInfo.testPaperId = item.ID
           this.activeTestName = item.testPaperName
         }
         arr.push(obj)
       })
        arr.unshift({
           value:0,
           label:"全部",
           status:false
         })
       this.options = arr
     }
    this.getTableData();
  }
};
</script>

<style>
</style>