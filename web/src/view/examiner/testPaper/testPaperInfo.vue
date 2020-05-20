<template>
  <div class="svgInfo" v-loading.fullscreen.lock="loading">
    <input
      @change="handlerUpload($event)"
      ref="fileupload"
      type="file"
      v-if="!loading"
      v-show="false"
    />
    <div v-html="svg"></div>
    <!--自定义右键菜单html代码-->
    <ul :style="{left:left+'px',top:top+'px'}" class="contextmenu" v-show="contextMenuVisible">
      <li @click="upload('flow')" v-if="!activeNode.flow">流量文件:上传</li>
      <li @click="download('flow')" v-if="activeNode.flow">流量文件:下载</li>
      <li @click="clear('flow')" v-if="activeNode.flow">流量文件:清除</li>
      <li @click="upload('configuration')" v-if="!activeNode.configuration">配置文件:上传</li>
      <li @click="download('configuration')" v-if="activeNode.configuration">配置文件:下载</li>
      <li @click="clear('configuration')" v-if="activeNode.configuration">配置文件:清除</li>
      <li @click="upload('log')" v-if="!activeNode.log">日志文件:上传</li>
      <li @click="download('log')" v-if="activeNode.log">日志文件:下载</li>
      <li @click="clear('log')" v-if="activeNode.log">日志文件:清除</li>
      <li @click="upload('sourceCode')" v-if="!activeNode.sourceCode">源码文件:上传</li>
      <li @click="download('sourceCode')" v-if="activeNode.sourceCode">源码文件:下载</li>
      <li @click="clear('sourceCode')" v-if="activeNode.sourceCode">源码文件:清除</li>
    </ul>
  </div>
</template>
<script>
import {
  getTestPaperSvg,
  findAndCreateTestPaperSvgNode,
  downloadTestPaperSvgNode,
  clearTestPaperSvgNode
} from "@/api/testPaper";
import axios from "axios"; // 引入axios
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 99999
});
export default {
  name: 'TestPaperInfo',
  data() {
    return {
      svg: '',
      contextMenuVisible: false,
      left: 0,
      top: 0,
      loading: false,
      active: {},
      activeNode: {},
      formData: null,
      isCollapse: false,
      isMobile: false
    }
  },
  methods: {
    upload(type) {
      this.param = new FormData();
      this.param.append("type", type);
      this.$refs.fileupload.click();
    },
    async handlerUpload(e) {
      this.loading = true;
      this.param.append("file", e.target.files[0]);
      this.param.append("nodeId", this.active.id);
      this.param.append("testPaperID", Number(this.$route.query.id));
      const token = this.$store.getters["user/token"];
      const res = await service.post(
        "/testPaper/uploadTestPaperSvgNode",
        this.param,
        {
          headers: {
            "x-token": token
          }
        }
      );
      this.loading = false;
      if (res.data.code == 0) {
        this.$message({
          type: "success",
          message: "上传成功"
        });
      } else {
        this.$message({
          type: "error",
          message: res.data.msg
        });
      }
    },
    async download(type) {
      const res = await downloadTestPaperSvgNode({
        path: this.activeNode[type]
      });
      const blob = new Blob([res]);
      const fileName = this.activeNode[type].split("/" + type + "/")[1];
      if ("download" in document.createElement("a")) {
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
    async clear(type) {
      const res = await clearTestPaperSvgNode({
        nodeId: this.active.id,
        testPaperID: Number(this.$route.query.id),
        type: type
      });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "清除成功"
        });
      }
    },
    async openContextMenu(dom) {
      window.event.returnValue = false;
      var oEvent = window.event;
      this.active = dom;
      const res = await findAndCreateTestPaperSvgNode({
        nodeId: this.active.id,
        testPaperID: Number(this.$route.query.id)
      });
      if (res.code == 0) {
        this.activeNode = res.data.node;
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
    }
  },
  async created() {
    this.$bus.on("mobile", isMobile => {
      this.isMobile = isMobile;
    });
    this.$bus.on("collapse", isCollapse => {
      this.isCollapse = isCollapse;
    });
    const res = await getTestPaperSvg({ ID: Number(this.$route.query.id) });
    if (res.code == 0) {
      this.svg = res.data.retestPaper
      setTimeout(() => {
        const g = document.getElementsByTagName("g");
        for (let key = 0; key < g.length; key++) {
          const paths = g[key].getElementsByTagName("path");
          if (paths.length > 0) {
            paths[0].oncontextmenu = () => {
              this.openContextMenu(paths[0]);
            };
          }
        }
      }, 0)
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

g {
  cursor: pointer;
}
</style>