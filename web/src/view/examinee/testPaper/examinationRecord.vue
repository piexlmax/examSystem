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
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
     <el-table-column label="同意协议" prop="agreement" width="120"></el-table-column>
    
     <el-table-column label="答题文件" prop="testPath" width="120"></el-table-column>
    
     <el-table-column label="得分" prop="score" width="120"></el-table-column>
    
      <el-table-column label="按钮组">
        
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
import {
    getExaminationRecordList
} from "@/api/examinationRecord";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "ExaminationRecord ",
  mixins: [infoList],
  data() {
    return {
      listApi: getExaminationRecordList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      formData: {
        agreement:null,testPath:null,score:null,
      }
    };
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
    this.getTableData();
  }
};
</script>

<style>
</style>