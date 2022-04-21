<template>
  <div style="background-color: #FFFFFF">
    <br/>
    <a-space style="padding-left: 10px">
      Container：
      <a-select v-model:value="data.container" style="min-width: 200px; max-width: 600px" @change="handleChange">
        <a-select-opt-group label="容器">
          <a-select-option :value=k v-for="(k, index) in data.containerData" :key="index">
            {{ k }}
          </a-select-option>
        </a-select-opt-group>

        <a-select-opt-group label="初始容器">
          <a-select-option :value=k v-for="(k, index) in data.initContainerData" :key="index">
            {{ k }}
          </a-select-option>
        </a-select-opt-group>
      </a-select>
      Pod：
      <a-select v-model:value="data.pod" placeholder="请选择容器" style="min-width: 200px; max-width: 800px" @change="handleLogChange">
        <a-select-option :value=k v-for="(k, index) in data.podNameData" :key="index">
          {{ k }}
        </a-select-option>
      </a-select>

      <a-button type="primary" @click="GetLog()">刷新</a-button>
      自动刷新/1秒：<a-switch v-model:checked="data.autoRefresh" checked-children="开" un-checked-children="关" @change="refreshLog" />

    </a-space>

    <a-tooltip color="#ffffff" :overlayStyle="{'font-size': '12px', 'max-width': '400px'}" placement="bottom">
      <template #title>
          <span style="color: #666">下载日志</span>
      </template>
      <DownloadOutlined  :style="{fontSize: '26px', float: 'right', paddingRight: '20px'}" @click="downLoadLogFile()"/>
    </a-tooltip>

    <br/><br/>
    <div id="filelog-container" style="height: 600px; overflow-y: scroll; background: #404040; color: #dedede; padding: 10px;">
      <div id="viewLog" v-if="data.logData=='' || data.logData==null">
        <p>暂无日志</p>
      </div>
      <div id="viewLog" v-else>
        <div style="white-space: pre-wrap;" v-for="(log, i) in data.logData" :key="i">{{ log.content }}</div>
      </div>
    </div>
  </div>

</template>

<script>
import { UserOutlined,AppleOutlined } from '@ant-design/icons-vue';
import {defineComponent, onBeforeUnmount, onMounted, onUnmounted, reactive, ref, watch} from 'vue';
import {queryHostGroups} from "../../api/group";

import { Terminal } from "xterm"
import { FitAddon } from 'xterm-addon-fit'
import { WebLinksAddon } from 'xterm-addon-web-links';
import { AttachAddon } from 'xterm-addon-attach';
import { SearchAddon } from 'xterm-addon-search';
import "xterm/css/xterm.css"
import {useRouter, useRoute} from "vue-router";
import {message} from "ant-design-vue";
export default defineComponent({
  components: {

  },

  setup() {
    const welcome =
        '\n' +
        '    __           __                     _       __       __       ______                        _                __\n' +
        '   / /   __  __ / /_   ____ _ ____     | |     / /___   / /_     /_  __/___   _____ ____ ___   (_)____   ____ _ / /\n' +
        '  / /   / / / // __ \\ / __ `// __ \\    | | /| / // _ \\ / __ \\     / /  / _ \\ / ___// __ `__ \\ / // __ \\ / __ `// / \n' +
        ' / /___/ /_/ // /_/ // /_/ // / / /    | |/ |/ //  __// /_/ /    / /  /  __// /   / / / / / // // / / // /_/ // /  \n' +
        '/_____/\\__,_//_.___/ \\__,_//_/ /_/     |__/|__/ \\___//_.___/    /_/   \\___//_/   /_/ /_/ /_//_//_/ /_/ \\__,_//_/   \n' +
        '                                                                                                                   \n'



    let store = reactive({
      treeData: [],
      shellWs: ref(null),
      rows: ref(null),
      cols: ref(null),
      ws: ref(),
      webSSHUrl: "",
      instanceId: ref(null),
    })


    // 创建Terminal实例
    const term = new Terminal({
      rendererType: 'canvas',
      cursorBlink: true,
      // 光标样式
      cursorStyle: "underline",
      // 光标设置为下一行开始
      convertEol: true,
      allowTransparency: true,
      fontFamily: 'operator mono,SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace',
      fontSize: 15,
      disableStdin: false,
      // 终端回滚量
      scrollback: 10000,
      tabStopWidth: 8,
      bellStyle: "sound",
      rightClickSelectsWord: true,
      // row: 2000,
      // alert(rows) //42.625
      // alert(cols) // 130
      row: 42,
      cols: 188,
      rows: 42,

      theme: {
        // 字体
        foreground: "white",
        background: "#060101",
        cursor: "help",
      }
    })
    const webLinksPlugin = new WebLinksAddon();
    const SearchPlugin = new SearchAddon();
    const fitPlugin = new FitAddon();
    const fitAddon = new FitAddon();
    // canvas背景全屏
    term.loadAddon(fitAddon)
    fitAddon.fit();
    const router = useRouter();
    const route = useRoute()

    // 获取主机实例id
    store.instanceId = route.params.uuid
    if (store.instanceId && store.instanceId !== "null") {
      let webConsole = {
        "state": true,
        "uuid": route.params.uuid,
        "hostName": route.params.hostName,
        "privateAddr": route.params.privateAddr,
      }
      localStorage.setItem("webConsole", JSON.stringify(webConsole))
      const token = localStorage.getItem('token');
      const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
      store.webSSHUrl = `${protocol}//localhost:8999/api/v1/ws/webssh?instanceId=${store.instanceId}&token=${token}`
      // store.webSSHUrl = `${protocol}//${window.location.host}/api/v1/ws/webssh?instanceId=${instanceId}&token=${token}`
      store.ws = new WebSocket(store.webSSHUrl);

      term.loadAddon(fitPlugin);
      term.loadAddon(webLinksPlugin); //链接检测
      term.loadAddon(SearchPlugin);
    }

    const hosts = []
    const expandedKeys = ref(['0-0']);
    const selectedKeys = ref(['0-0']);
    const panes = ref([
      {
        title: route.params.hostName,
        key: "1",
      },
    ]);

    const activeKey = ref(panes.value[0].key);

    // 新建ssh标签页
    const newSSHTab = () => {
      panes.value.push({
        title: route.params.hostName,
        key: route.params.uuid,
      });
    };

    // 关闭ssh标签页
    const closeSSHTab = targetKey => {
      let lastIndex = 0;
      panes.value.forEach((pane, i) => {
        if (pane.key === targetKey) {
          lastIndex = i - 1;
          // 关闭所有标签显示欢迎信息
          if (lastIndex === -1) {
            store.instanceId = null
          }
        }
      });
      panes.value = panes.value.filter(pane => pane.key !== targetKey);

      if (panes.value.length && activeKey.value === targetKey) {
        if (lastIndex >= 0) {
          activeKey.value = panes.value[lastIndex].key;
        } else {
          activeKey.value = panes.value[0].key;
        }
      }

    };

    const onEdit = targetKey => {
      closeSSHTab(targetKey);
    };
    // 获取分组
    const getGroups = async () => {
      const result = await queryHostGroups()
      if (result.errCode !== 0){
        message.error('获取资产分组失败')
      }else {
        store.treeData = result.data.treeData
      }
    }

    /***
     * 页面刷新时关闭websocket
     *
     */
    const closeWs = () => {
      if (store.ws) {
        const closed = JSON.stringify({
          type: 'closePty',
          close: true,
        });
        store.ws.send(closed);
      }
    }
    const fitListener = () => {
      fitPlugin.fit()
    };

    const terminalRef = ref()

    /***
     * 判断websocket是否建立连接
     * @returns {boolean}
     */
    function isWebSocketOnline() {
      if(store.ws){
        return store.ws.readyState === WebSocket.OPEN;
      }else {
        return false;
      }
    }
    function fullScreen() {
      console.log("full screen")
    }
    onMounted(() => {
      getGroups()
      if (store.instanceId) {
        store.ws.onopen = function (e){
          console.log("WS 通信已建立连接")
          const attachPlugin = new AttachAddon(store.ws);
          term.loadAddon(attachPlugin);
          term.open(document.getElementById('terminal'));
          term.writeln("Connecting...");
          term.clear()
          fitAddon.fit()
          term.focus()
        }
        // 接受数据
        store.ws.onmessage = function (e) {
          const message = e.data
          if (message.indexOf) {
            if (message.indexOf('Anew-Sec-WebSocket-Key') !== -1) {
              let secKey = message.substring(message.lastIndexOf(':') + 1, message.length).replace(/[\r\n]/g, "")
              localStorage.setItem('TABS_TTY_HOSTS', secKey);
            }
          }
        }
        store.ws.onerror = function (e) {
          message.error("连接异常, 请刷新后重试！")
        }
        store.ws.onclose = function (e) {
          setTimeout(() => term.write('\x1b[1;1;31m\r\nConnection is closed.\x1b[0m\r\n'), 500)
        }

        // 监听窗口大小
        term.onResize((size) => {
          const resize = JSON.stringify({
            type: 'resizePty',
            cols: size.cols,
            rows: size.rows,
          });
          if (isWebSocketOnline()){
            store.ws.send(resize);
            term.resize(size.cols, size.rows);
          }

        });
        window.addEventListener("resize", fitListener)
      }
    })

    onUnmounted(()=>{
      closeWs()
      localStorage.removeItem("webConsole")
      window.removeEventListener("resize", fitListener)
    })

    return {
      term,
      panes,
      activeKey,
      onEdit,
      newSSHTab,
      hosts,
      welcome,
      selectedKeys1: ref(['2']),
      selectedKeys2: ref(['1']),
      expandedKeys,
      selectedKeys,
      store,
      terminalRef,
      fullScreen,

    };
  },

});
</script>
<style scoped>
.web-terminal {
  display: flex;
  background-color: #060101;
  padding-left: 240px;
  width: 100%;
  height: 100%;
  justify-content: center;

}
.luban-pre {
  display: flex;
  background-color: #060101;
  color: #fff;
  justify-content: center;
  padding-top: 240px;
}
.logo {
  height: 42px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f0f2f5;
  font-size: 15px;
}

.terminal {
  display: flex;
  background-color: #060101;
  padding-left: 5px;
  height: calc(100vh - 120px);
}
</style>
