<template>
  <div>
    <a-tabs v-model:activeKey="state.activeKey" style="background-color: #fff;height: 100%; " rowKey="state.id">
      <a-tab-pane key="1" tab="集群概览">
        <a-row>
          <a-col :span="8" style="text-align: center">
            <a-card size="small" title="CPU信息" style="left: 5px;width: 96%" >
              <a-space>
                <a-spin :spinning="state.loading"/>
                <div id="cpuContainer"></div>
              </a-space>
            </a-card>

            <a-space>
              <a-card size="small" title="Used" style="width: 261px; height: 80px;margin-left: -11px">
                <p>
                  <span style="color: green">{{ state.cpuSummary.Used }}</span>
                  <span> Core</span>
                </p>
              </a-card>
              <br>
              <a-card size="small" title="Total" style="width: 261px; height: 80px">
                <p>
                  <span style="color: green">{{ state.cpuSummary.Total }}</span>
                  <span> Core</span>
                </p>
              </a-card>
            </a-space>

          </a-col>


          <a-col :span="8" style="text-align: center">
            <a-card size="small" title="节点状态" style="height: 360px;width: 96%">
              <a-space style="">
                <span v-if="2>0">共有：{{ state.nodeSummary.Total +2  }}</span>
                <span v-if="2>0" style="color: orange">正常：{{ state.nodeSummary.Total +1 }}</span>
                <span v-if="2>0" style="color: red">可调度：{{ state.nodeSummary.Schedulable }}</span>
              </a-space>
              <a-spin :spinning="state.loading"/>
              <div id="container"></div>
            </a-card>
          </a-col>


          <a-col :span="8" style="text-align: center">
            <a-card size="small" title="内存信息" style="left: 12px; width: 96%">
              <a-space>
                <a-spin :spinning="state.loading"/>
                <div id="memContainer"></div>
              </a-space>
            </a-card>

            <a-space>
              <a-card size="small" title="Used" style="width: 265px; height: 80px; left: 5px">
                <p>
                  <span style="color: green">{{ state.memorySummary.Used }}</span>
                  <span> G</span>
                </p>
              </a-card>
              <a-card size="small" title="Total" style="width: 265px; height: 80px">
                <p>
                  <span style="color: green">{{ state.memorySummary.Total }}</span>
                  <span> G</span>
                </p>
              </a-card>
            </a-space>

          </a-col>

        </a-row>


        <br/>
        <div style="margin-left: 15px; margin-right: 15px">
          <h4 style="font-weight: bold;margin-left: 20px">事件</h4>
          <a-spin :spinning="state.eventLoading">
            <a-table :columns="columns" :data-source="state.eventsData" size="middle" rowKey="state.eventsData.key">
              <template #type="text">
                <span v-if="text.text=='Warning'" style="color: orange">
                  {{ text.text }}
                </span>
                <span v-else>
                  {{ text.text }}
                </span>
              </template>

              <template #lastSeen="text">
                <span>
                  {{ $filters.fmtTime(text.text) }}
                </span>
              </template>

            </a-table>
          </a-spin>
        </div>

        <br/>
      </a-tab-pane>
<!--      <a-tab-pane key="2" tab="Tab 2" force-render>Content of Tab Pane 2</a-tab-pane>-->
<!--      <a-tab-pane key="3" tab="Tab 3">Content of Tab Pane 3</a-tab-pane>-->
    </a-tabs>

  </div>
</template>

<script>
import {onMounted, reactive, defineComponent} from "vue";
import {useRoute} from "vue-router";
import { accMul,accMulInt } from "@/plugin/utils/accMul";
import {Gauge, Liquid, measureTextWidth} from '@antv/g2plot';
import {getK8SClusterDetail, getEvents} from "../../api/k8s";

const columns = [
  {
    title: '类型',
    dataIndex: 'type',
    key: 'type',
    slots: {
      customRender: 'type',
    },
    width: 120,
  },
  {
    title: '对象',
    dataIndex: 'name',
    key: 'name',
    width: 120,
  },
  {
    title: '信息',
    dataIndex: 'message',
    key: 'message',
    ellipsis: true,
  },
  {
    title: '内容',
    dataIndex: 'reason',
    key: 'reason',
  },
  {
    title: '来源',
    dataIndex: 'sourceComponent',
    key: 'sourceComponent',
  },
  {
    title: '时间',
    dataIndex: 'lastSeen',
    key: 'lastSeen',
    width: 180,
    slots: {
      customRender: 'lastSeen',
    },
  },
];
export default defineComponent({
  name: "ClusterDetail",
  setup() {
    const state = reactive({
      id: "",
      activeKey: "1",
      nodeSummary: "",
      cpuSummary: "",
      memorySummary: "",
      nodes: "",
      eventsData: [],
      loading: true,
      eventLoading: true,
    });

    const route = useRoute()

    const events = (data) => {

      getEvents(data).then(res => {
        state.eventsData = res.data.list
        state.eventLoading = false

      })
    }

    onMounted(() => {
      state.id = route.params.id
      state.loading = true
      getK8SClusterDetail(state.id).then(res => {
        state.nodeSummary = res.data.nodeSummary
        state.cpuSummary = res.data.cpuSummary
        state.memorySummary = res.data.memorySummary
        state.nodes = res.data.nodes
        state.loading = false
          // 计算百分比
        state.NodeRunningStatus = accMul(res.data.nodeSummary.Ready , res.data.nodeSummary.Total)
        state.cpuUsage = accMul(res.data.cpuSummary.Used , res.data.cpuSummary.Total)
        state.memoryUsage = accMul(res.data.memorySummary.Used , res.data.memorySummary.Total)
        const gauge = new Gauge('container', {
          percent: accMulInt(state.nodeSummary.Ready , state.nodeSummary.Total + 3),
          range: {
            color: '#1890ff',
          },
          startAngle: Math.PI,
          endAngle: 2 * Math.PI,
          indicator: null,
          width: 274,
          height: 180,
          autoFit: true,
        });
        gauge.render()
        const liquidPlot = new Liquid(document.getElementById('cpuContainer'), {
          percent: accMulInt(state.cpuSummary.Used,state.cpuSummary.Total),
          radius: 0.8,
          autoFit: true,
          width: 200,
          height: 200,
          statistic: {
            title: {
              formatter: () => 'CPU使用率',
              style: ({percent}) => ({
                fill: percent > 0.65 ? 'white' : 'rgba(44,53,66,0.85)',
              }),
            },
            content: {
              style: ({percent}) => ({
                fontSize: 40,
                lineHeight: 1,
                fill: percent > 0.65 ? 'white' : 'rgba(44,53,66,0.85)',
              }),
              customHtml: (container, view, {percent}) => {
                const {width, height} = container.getBoundingClientRect();
                const d = Math.sqrt(Math.pow(width / 2, 2) + Math.pow(height / 2, 2));
                const text = `${(percent * 100).toFixed(0)}%`;
                const textWidth = measureTextWidth(text, {fontSize: 40});
                const scale = Math.min(d / textWidth, 1);
                return `<div style="width:${d}px;display:flex;align-items:center;justify-content:center;font-size:${scale}em;line-height:${
                    scale <= 1 ? 1 : 'inherit'
                }">${text}</div>`;
              },
            },
          },
          liquidStyle: ({percent}) => {
            return {
              fill: percent > 0.45 ? '#5B8FF9' : '#FAAD14',
              stroke: percent > 0.45 ? '#5B8FF9' : '#FAAD14',
            };
          },
          color: () => '#5B8FF9',
        });
        liquidPlot.render()
        const memLiquidPlot = new Liquid(document.getElementById('memContainer'), {
          percent: accMulInt(state.memorySummary.Used, state.memorySummary.Total),
          radius: 0.8,
          autoFit: true,
          width: 200,
          height: 200,
          statistic: {
            title: {
              formatter: () => '内存使用率',
              style: ({percent}) => ({
                fill: percent > 0.65 ? 'white' : 'rgba(44,53,66,0.85)',
              }),
            },
            content: {
              style: ({percent}) => ({
                fontSize: 40,
                lineHeight: 1,
                fill: percent > 0.65 ? 'white' : 'rgba(44,53,66,0.85)',
              }),
              customHtml: (container, view, {percent}) => {
                const {width, height} = container.getBoundingClientRect();
                const d = Math.sqrt(Math.pow(width / 2, 2) + Math.pow(height / 2, 2));
                const text = `${(percent * 100).toFixed(0)}%`;
                const textWidth = measureTextWidth(text, {fontSize: 40});
                const scale = Math.min(d / textWidth, 1);
                return `<div style="width:${d}px;display:flex;align-items:center;justify-content:center;font-size:${scale}em;line-height:${
                    scale <= 1 ? 1 : 'inherit'
                }">${text}</div>`;
              },
            },
          },
          liquidStyle: ({percent}) => {
            return {
              fill: percent > 0.45 ? '#5B8FF9' : '#FAAD14',
              stroke: percent > 0.45 ? '#5B8FF9' : '#FAAD14',
            };
          },
          color: () => '#5B8FF9',
        });
        memLiquidPlot.render()
      })
      events(state.id)
    });


    return {
      state,
      columns,
    };
  },
})
</script>

<style scoped>

</style>