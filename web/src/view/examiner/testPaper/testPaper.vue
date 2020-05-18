<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
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
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="考题名称" prop="testPaperName" width="120"></el-table-column>

      <el-table-column label="出题人" prop="testPaperAuthor" width="120"></el-table-column>
      <el-table-column label="考生须知" prop="testPaperNote" width="120"></el-table-column>
      <el-table-column label="允许多次提交" prop="testPaperNote" width="120">
        <template slot-scope="scope">
          <div>
            {{scope.row.testPaperSubmitTimes?"是":"否"}}
          </div>
        </template>
      </el-table-column>

      <el-table-column label="考试起止时间" width="300">
        <template slot-scope="scope">
          <div>{{scope.row.testPaperStartTime|formatDate}} 至 {{scope.row.testPaperEndTime|formatDate}}</div>
        </template>
      </el-table-column>
      <el-table-column label="答题模板" prop="testPaperMould" width="120">
        <template slot-scope="scope">
          <div>
            <el-button
              v-if="scope.row.testPaperMould"
              type="text"
              size="mini"
              @click="getTestPaperMould(scope.row)"
            >下载</el-button>
            <span v-else>暂未上传模板</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="460">
        <template slot-scope="scope">
          <el-button
            v-if="scope.row.testPaperSvg"
            @click="viewTestPaper(scope.row)"
            size="small"
            type="text"
          >查看SVG</el-button>
          <el-button @click="updateTestPaper(scope.row)" size="small" type="text">变更</el-button>
          <el-button
            @click="publicTestPaper(scope.row,true)"
            v-if="!scope.row.testPaperStatus"
            size="small"
            type="text"
          >发布考题</el-button>
          <el-button @click="publicTestPaper(scope.row,false)" v-else size="small" type="text">撤销发布</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteTestPaper(scope.row)">确定</el-button>
            </div>
            <el-button type="text" size="mini" slot="reference">删除</el-button>
          </el-popover>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-form-item label="考题名称" prop="testPaperName">
          <el-input
            v-model="formData.testPaperName"
            placeholder="请输入考题名称"
            clearable
            :style="{width: '100%'}"
          ></el-input>
        </el-form-item>
        <el-form-item label="考生须知" prop="testPaperNote">
          <el-input
            v-model="formData.testPaperNote"
            type="textarea"
            :rows="2"
            placeholder="请输入考生须知"
            clearable
            :style="{width: '100%'}"
          ></el-input>
        </el-form-item>
        <el-form-item label="是否允许多次提交" prop="testPaperSubmitTimes">
          <el-checkbox
            v-model="formData.testPaperSubmitTimes"
          ></el-checkbox>
        </el-form-item>
        <el-form-item label="开考时间">
          <el-date-picker
            @change="timeChange"
            v-model="timeArea"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
          ></el-date-picker>
        </el-form-item>

        <el-form-item label="答题模板" prop="testPaperMould">
          <el-upload
            ref="TestPaperMould"
            :action="`${path}/testPaper/uploadMould`"
            :headers="{'x-token':token}"
            name="mould"
            :file-list="mouldFileList"
            :limit="1"
            :before-remove="beforeRemoveMould"
            :on-exceed="handleExceed"
            :on-success="handleMouldSuccess"
            :show-file-list="true"
          >
            <el-button size="small" type="primary" icon="el-icon-upload">点击上传</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="上传考题" prop="testPaperSvg">
          <el-upload
            ref="testPaperSvg"
            :action="`${path}/testPaper/uploadSvg`"
            :headers="{'x-token':token}"
            name="svg"
            :limit="1"
            :before-remove="beforeRemoveSvg"
            :file-list="svgFileList"
            :on-exceed="handleExceed"
            :on-success="handleSvgSuccess"
            :show-file-list="true"
          >
            <el-button size="small" type="primary" icon="el-icon-upload">点击上传</el-button>
          </el-upload>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
//UTF字符转换
var UTFTranslate = {
  Change: function(pValue) {
    return pValue.replace(/[^\u0000-\u00FF]/g, function($0) {
      return escape($0).replace(/(%u)(\w{4})/gi, "&#x$2;");
    });
  },
  ReChange: function(pValue) {
    return unescape(
      pValue
        .replace(/&#x/g, "%u")
        .replace(/\\u/g, "%u")
        .replace(/;/g, "")
    );
  }
};
import {
  createTestPaper,
  deleteTestPaper,
  updateTestPaper,
  findTestPaper,
  getTestPaperList,
  getTestPaperMould,
  removeTestPaperFile,
  publicTestPaper
} from "@/api/testPaper"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "TestPaper",
  mixins: [infoList],
  data() {
    return {
      listApi: getTestPaperList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      path: path,
      timeArea: [],
      mouldFileList: [],
      svgFileList: [],
      formData: {
        testPaperStartTime: null,
        testPaperEndTime: null,
        testPaperNote: "",
        testPaperName: "",
        testPaperAuthor: "",
        testPaperMould: "",
        testPaperSvg: "",
        testPaperStatus: false,
        testPaperSubmitTimes:false,
      },
      rules: {
        testPaperName: [
          {
            required: true,
            message: "请输入考题名称",
            trigger: "blur"
          }
        ]
      }
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"])
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
    async publicTestPaper(row, flag) {
      row.testPaperStatus = flag;
      const res = await publicTestPaper(row);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: flag ? "发布成功" : "取消成功"
        });
        this.getTableData();
      }
    },
    async beforeRemoveSvg(file) {
      this.$confirm("此操作将永久删除该文件, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await removeTestPaperFile({
            ID: this.formData.ID,
            testPaperSvg: file.url
          });
          if (res.code == 0) {
            this.formData.testPaperSvg = "";
            this.$message({
              type: "success",
              message: "删除成功"
            });
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
          this.svgFileList = [{ url: file.url, name: file.name }];
        });
    },
    async beforeRemoveMould(file) {
      this.$confirm("此操作将永久删除该文件, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await removeTestPaperFile({
            ID: this.formData.ID,
            testPaperMould: file.url
          });
          if (res.code == 0) {
            this.formData.testPaperMould = "";
            this.$message({
              type: "success",
              message: "删除成功"
            });
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
          this.mouldFileList = [{ url: file.url, name: file.name }];
        });
    },
    handleExceed() {
      this.$message({
        type: "error",
        message: "最多上传1个文件"
      });
    },
    async getTestPaperMould(row) {
      const res = await getTestPaperMould({ ID: row.ID });
      const blob = new Blob([res]);
      const fileName = row.testPaperMould.split("/mould/")[1];
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
    },
    timeChange() {
      this.formData.testPaperStartTime = this.timeArea[0];
      this.formData.testPaperEndTime = this.timeArea[1];
    },
    viewTestPaper(row) {
      this.$router.push({
        name: "testPaperInfo",
        query: {
          id: row.ID
        }
      });
    },
    handleSvgSuccess(res) {
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "上传成功"
        });
        this.formData.testPaperSvg = res.data.path;
      } else {
        this.$message({
          type: "error",
          message: res.msg
        });
      }
    },
    handleMouldSuccess(res) {
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "上传成功"
        });
        this.formData.testPaperMould = res.data.path;
      } else {
        this.$message({
          type: "error",
          message: res.msg
        });
      }
    },
    async updateTestPaper(row) {
      const res = await findTestPaper({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        if (res.data.retestPaper.testPaperSvg) {
          this.svgFileList = [
            {
              name: res.data.retestPaper.testPaperSvg.split("/svg/")[1],
              url: res.data.retestPaper.testPaperSvg
            }
          ];
        }
        if (res.data.retestPaper.testPaperMould) {
          this.mouldFileList = [
            {
              name: res.data.retestPaper.testPaperMould.split("/mould/")[1],
              url: res.data.retestPaper.testPaperMould
            }
          ];
        }
        this.timeArea = [
          new Date(res.data.retestPaper.testPaperStartTime),
          new Date(res.data.retestPaper.testPaperEndTime)
        ];
        this.formData = res.data.retestPaper;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.timeArea = [];
      this.formData = {
        testPaperStartTime: null,
        testPaperEndTime: null,
        testPaperNote: "",
        testPaperName: "",
        testPaperAuthor: "",
        testPaperMould: "",
        testPaperSvg: "",
        testPaperSubmitTimes:false,
        testPaperStatus: false
      };
      this.mouldFileList = [];
      this.svgFileList = [];
    },
    async deleteTestPaper(row) {
      this.visible = false;
      const res = await deleteTestPaper(row);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      if (!this.formData.testPaperMould || !this.formData.testPaperSvg) {
        this.$message({
          type: "error",
          message: "请完整填写内容"
        });
        return;
      }
      this.$refs["elForm"].validate(async valid => {
        if (!valid) return;
        // TODO 提交表单
        let res;
        this.formData.testPaperAuthor = this.userInfo.nickName;
        switch (this.type) {
          case "create":
            res = await createTestPaper(this.formData);
            break;
          case "update":
            res = await updateTestPaper(this.formData);
            break;
          default:
            res = await createTestPaper(this.formData);
            break;
        }
        if (res.code == 0) {
          this.closeDialog();
          this.getTableData();
        }
      });
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.getTableData();
  }
};
</script>

<style>
</style>