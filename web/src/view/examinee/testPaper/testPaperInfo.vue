<template>
  <div class="svgInfo">
    <div v-if="status==0">
      没有进行中的考试
    </div>
    <div v-if="status==1&&!examinationRecord.agreement">
      <div style="padding:24px">考试协议需要同意</div>
      <el-button @click="agree">同意并开始考试</el-button>
    </div>
    <div v-if="status==1&&examinationRecord.agreement" v-html="svg"></div>
    <div v-if="status==2">
      倒计时
    </div>
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
import { getTestPaperSvg ,findAndCreateTestPaperSvgNode,downloadTestPaperSvgNode,getActiveTestPaper} from "@/api/testPaper";
import { createExaminationRecord,findExaminationRecord } from "@/api/examinationRecord";


import { mapGetters } from "vuex";

export default {
  name: "TestPaperInfo",
  data() {
    return {
      examinationRecord:{},
      testPaper:{},
      svg: "",
      contextMenuVisible: false,
      left: 0,
      top: 0,
      active:{},
      activeNode:{},
      formData:null,
      status:0,
      agreement:false,
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"])
  },
  methods: {
    async agree(){
      const res = await createExaminationRecord({sysUserId:Number(this.userInfo.ID),testPaperId:Number(this.testPaper.ID),agreement:true})
       if(res.code == 0){
        this.getExamination()
      }
    },
      async download(type){
        const res = await downloadTestPaperSvgNode({path:this.activeNode[type]})
        const blob = new Blob([res]);
        const fileName = this.activeNode[type].split("/"+type+"/")[1];
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
    async openContextMenu(e) {
      window.event.returnValue = false;
      var oEvent = e || event;
      this.active = e
      const res = await findAndCreateTestPaperSvgNode({nodeId:this.active.target.id,testPaperID:Number(this.testPaper.ID)})
      if (res.code == 0){
      this.activeNode = res.data.node
      this.contextMenuVisible = true;
      let width;
      if (this.isCollapse) {
        width = 54;
      } else {
        width = 220;
      }
      if (this.isMobile) {
        width = 0;
      }
      this.left = oEvent.clientX - width;
      this.top = oEvent.clientY + 10;
      }
    },
    async getTest(id){
      const res = await getTestPaperSvg({ ID: Number(id) });
      if (res.code == 0) {
        this.svg = res.data.retestPaper;
        setTimeout(() => {
          const tspans = document.getElementsByTagName("tspan");
          for (let k = 0; k < tspans.length; k++) {
            const tspan = tspans[k].getElementsByTagName("tspan");
            if (tspan[0]) {
              tspan[0].oncontextmenu = this.openContextMenu;
            }
          }
        }, 0);
      }
    },
    async getExamination(){
      const res = await findExaminationRecord({sysUserId:this.userInfo.ID,testPaperId:this.testPaper.ID})
      this.examinationRecord = res.data.examinationRecord
    }
  },
  async created() {
    const res = await getActiveTestPaper()
    this.status = res.data.status
    if(this.status == 1){
      this.testPaper = res.data.testPaper
      this.getExamination()
      this.getTest(res.data.testPaper.ID)
    }
  },
  watch: {
    contextMenuVisible() {
      if (this.contextMenuVisible) {
        document.body.addEventListener("click", () => {
          this.contextMenuVisible = false;
        });
      } else {
        document.body.removeEventListener("click", () => {
          this.contextMenuVisible = false;
        });
      }
    }
  }
};
</script>
<style lang="scss">
.svgInfo {
  display: flex;
  justify-content: center;
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

tspan{
  tspan{
    cursor: pointer;
  }
}
</style>