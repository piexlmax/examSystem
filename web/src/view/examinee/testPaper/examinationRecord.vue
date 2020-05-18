<template>
  <div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="考题编号" width="80">
         <template slot-scope="scope">{{scope.row.testPaperId}}</template>
    </el-table-column>
    <el-table-column label="考题" width="180">
         <template slot-scope="scope">{{scope.row.testPaper.testPaperName}}</template>
    </el-table-column>
    <el-table-column label="考试时段" width="320">
        <template slot-scope="scope">
          <div>{{scope.row.testPaper.testPaperStartTime|formatDate}} 至 {{scope.row.testPaper.testPaperEndTime|formatDate}}</div>
        </template>
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
           <span v-if="scope.row.testPath">
             已提交
           </span>
           <span v-else>
             未作答
           </span>
         </div>
       </template>
     </el-table-column>
    
     <el-table-column label="得分" prop="score" width="120"></el-table-column>
     <el-table-column label="评语" prop="appraise" min-width="120"></el-table-column>
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
import {
    getExaminationRecordList
} from "@/api/examinationRecord";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { mapGetters } from "vuex";

export default {
  name: "ExaminationRecord",
  mixins: [infoList],
  data() {
    return {
      listApi: getExaminationRecordList,
      dialogFormVisible: false,
      visible: false,
      type: "",
     
      
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
    
  },
  created() {
    this.searchInfo.sysUserId = this.userInfo.ID
    this.getTableData();
  }
};
</script>

<style>
</style>