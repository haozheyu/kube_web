<template>
  <div>

    <div style="margin-bottom: 16px">
      <a-space>
        <a-button type="primary" @click="addK8SCluster">注册集群</a-button>
        <a-button type="primary" danger :disabled="state.selectedRows.length<=0" @click="removeCluster()">批量删除</a-button>
      </a-space>
    </div>

    <a-table
        :row-selection="rowSelection"
        :columns="columns"
        :data-source="state.data"
        :pagination="false"
        rowKey="id"
        :locale="{emptyText: '暂无数据'}"
    >
      <template #name="{ text }">
        <span>
          {{ text }}
        </span>
      </template>
      <template #ClusterVersion="{ text }">
        <span>
          <a-tag color="cyan">{{ text }}</a-tag>
        </span>
      </template>

      <template #managerNode="{ text }">
        <span>
          <a-tag color="orange">{{ text }}</a-tag>
        </span>
      </template>

      <template #kubeConfig="{ text }">
        <a-tooltip placement="topLeft" title="查看凭证" color="orange">
          <a @click="ViewClusterConfig( text )"><IconFont type="pigs-icon-pingzheng"/></a>
        </a-tooltip>
      </template>

      <template #action="{text }">
        <span>
          <a @click="clusterDetail(text)">查看</a>
        </span>
      </template>

    </a-table>

    <a-modal v-model:visible="state.ClusterConfigVisible" title="查看集群凭证" :footer="null" :keyboard="false" :maskClosable="false">

      <a-textarea v-model:value="state.ClusterConfig" placeholder="KubeConfig内容" style="width: 100%; height: 600px"/>
    </a-modal>

    <a-modal v-model:visible="createK8SClusterVisible" title="添加新集群" @ok="onSubmit" @cancel="resetForm" cancelText="取消"
             okText="确定" :keyboard="false" :maskClosable="false">
      <a-form
          ref="formRef"
          :model="formState"
          :rules="rules"
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
      >
        <a-form-item ref="name" label="集群名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入集群名称"/>
        </a-form-item>
        <a-form-item  ref="kubeConfig" label="集群凭证" name="kubeConfig">
          <a-textarea v-model:value="formState.kubeConfig" placeholder="请粘贴KubeConfig内容"
                      style="width: 100%; height: 300px"/>
        </a-form-item>

      </a-form>
    </a-modal>

    <div class="float-right" style="padding: 10px 0;">

      <a-pagination size="md" :show-total="total => `共 ${state.total} 条数据`" :v-model="state.total"
                    :page-size-options="state.pageSizeOptions"
                    :total="state.total"

                    :pageSize="state.pageSize"
                    show-less-items align="right"
                    @showSizeChange="onShowSizeChange"
                    @change="onChange"
      >
      </a-pagination>
    </div>


  </div>
</template>

<script>
import {defineComponent, inject, onMounted, reactive, ref} from 'vue';
import {fetchK8SCluster, k8sCluster, delK8SCluster, clusterSecret} from '../../api/k8s'
import {createFromIconfontCN} from "@ant-design/icons-vue";
import router from "../../router";

const columns = [
  {
    title: '集群名称',
    dataIndex: 'name',
    slots: {customRender: 'name'},
  },
  {
    title: '集群版本',
    dataIndex: 'version',
    slots: {customRender: 'ClusterVersion'}
  },
  {
    title: '管理节点',
    dataIndex: 'master',
    slots: {customRender: 'managerNode'}
  },
  {
    title: '集群凭证',
    dataIndex: 'kubeConfig',
    slots: {customRender: 'kubeConfig'},
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
  },
  {
    title: '操作',
    // dataIndex: 'action',
    slots: {customRender: 'action'},
  }
];
const IconFont = createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/font_2828790_vphs1aik0kn.js',
});
export default defineComponent({
  name: "Manage",
  setup() {
    let selectedRowKeys;
    const state = reactive({
      selectedRows: [],
      selectedRowKeys,
      loading: false,
      data: [],
      pageSize: 10,
      current: null,
      total: 0,
      pageSizeOptions: ['10', '20', '30', '40'],
      ClusterConfigVisible: false,
      ClusterConfig: "",
    });

    const createK8SClusterVisible = ref(false);

    let addK8SCluster = () => {
      createK8SClusterVisible.value = true;
    }

    const formRef = ref();
    const formState = reactive({
      name: "",
      kubeConfig: "",
    });
    const rules = {
      name: [
        {
          required: true,
          message: '请输入集群名称',
          trigger: 'blur',
        },
        {
          min: 3,
          max: 25,
          message: '集群名称长度应为3~25',
          trigger: 'blur',
        },
      ],
      kubeConfig: [
        {
          required: true,
          message: '请粘贴KubeConfig内容',
          trigger: 'blur',
        },
      ],
    };

    const onSubmit = () => {
      formRef.value
          .validate()
          .then(() => {
            console.log(formState.name, formState.kubeConfig, "--------------")
            k8sCluster({
              "name": formState.name,
              "kubeConfig": formState.kubeConfig,
            }).then(res => {
                message.success("创建成功")
                createK8SClusterVisible.value = false;
                resetForm()
                getK8SCluster()
            }).catch(err => {
              message.warning("请检查输入配置是否正确")
            })
          })
    };

    const resetForm = () => {
      formRef.value.resetFields();
    };
    const message = inject('$message');
    // 获取集群信息
    const getK8SCluster = async (pageSize) => {
      const {data} = await fetchK8SCluster({size: pageSize})
      state.data = data.list
      state.total = data.totalCount
      state.pageSize = data.pageSize
    }
    // 翻页
    const onChange = async (pageNumber) => {
      fetchK8SCluster({
        pageNo: pageNumber,
        pageSize: state.pageSize
      }).then(res => {
          state.data = res.data.list
      });

    };
    // 显示条数
    const onShowSizeChange = async (current, pageSize) => {
      const {data} = await fetchK8SCluster({
        pageNo: pageSize,
      })
      state.data = data.list
      state.total = data.totalCount
      state.pageSize = data.pageSize
      state.current = data.pageNo
    };

    const rowSelection = {
      onChange: (selectedRowKeys, selectedRows) => {
        state.selectedRows = selectedRows
        state.selectedRowKeys = selectedRowKeys
      },
    };
    // 批量删除集群
    const removeCluster = async () => {
      let clusterIds = [];
      for (let i=0; i < state.selectedRowKeys.length; i++){
        clusterIds.push(state.selectedRowKeys[i])
      }
      await delK8SCluster({"clusterIds": clusterIds}).then(res => {
        if (res.errCode === 0){
          message.success(res.msg)
          getK8SCluster()
        }else {
          message.error(res.errMsg)
        }
      });

    };
    // 查看集群凭证
    const ViewClusterConfig = (text) => {
      state.ClusterConfig = text
      state.ClusterConfigVisible = true
    }
    // 查看集群详情
    const clusterDetail = (text) => {
      router.push({path: `/k8s/cluster/detail/${text.name}`})
    }
    onMounted(getK8SCluster)


    return {
      columns,
      state,

      createK8SClusterVisible,
      addK8SCluster,
      formRef,
      formState,
      rules,
      onSubmit,
      resetForm,

      labelCol: {
        span: 4,
      },
      wrapperCol: {
        span: 19,
      },

      onShowSizeChange,
      onChange,
      rowSelection,
      removeCluster,
      ViewClusterConfig,
      clusterDetail,
    };
  },
  components: {
    IconFont,
  }
});
</script>

<style>

</style>