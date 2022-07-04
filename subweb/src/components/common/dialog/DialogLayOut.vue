<template>
  <el-dialog v-model="dialogVisible" top="30vh" width="40%">
    <span>
      <p style="text-align: center">{{ message }}</p>
    </span>
  </el-dialog>
</template>

<script>
import { ElDialog } from 'element-plus';
import 'element-plus/theme-chalk/index.css';

export default {
  name: 'DialogLayOut',
  components: {
    ElDialog,
  },
  props: {
    dialogVisible: {
      type: Boolean,
      default: false,
    },
    message: {
      type: String,
      default: '这是一个封装好的弹窗组建',
    },
  },
};
</script>

<style scoped></style>
